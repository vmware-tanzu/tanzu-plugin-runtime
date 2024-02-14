// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/internal/kubeconfig"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// keys to Context's AdditionalMetadata map
const (
	OrgIDKey            = "tanzuOrgID"
	ProjectNameKey      = "tanzuProjectName"
	SpaceNameKey        = "tanzuSpaceName"
	ClusterGroupNameKey = "tanzuClusterGroupName"
)

const (
	// customCommandName is the name of the command expected to be implemented
	// by the CLI should there be a need to discover and alternative invocation
	// method
	customCommandName string = "_custom_command"
)

// ResourceInfo contains information identifying the Tanzu resource associated with the Context
type ResourceInfo struct {
	// OrgID ID of the Organization
	OrgID string
	// ProjectName name of the Project
	ProjectName string
	// SpaceName name of the Space
	SpaceName string
	// ClusterGroupName name of the ClusterGroup
	ClusterGroupName string
}

// cmdOptions specifies the command options
type cmdOptions struct {
	outWriter io.Writer
	errWriter io.Writer
}

type CommandOptions func(o *cmdOptions)

// WithOutputWriter specifies the CommandOption for configuring Stdout
func WithOutputWriter(outWriter io.Writer) CommandOptions {
	return func(o *cmdOptions) {
		o.outWriter = outWriter
	}
}

// WithErrorWriter specifies the CommandOption for configuring Stderr
func WithErrorWriter(errWriter io.Writer) CommandOptions {
	return func(o *cmdOptions) {
		o.errWriter = errWriter
	}
}

// WithNoStdout specifies to ignore stdout
func WithNoStdout() CommandOptions {
	return func(o *cmdOptions) {
		o.outWriter = io.Discard
	}
}

// WithNoStderr specifies to ignore stderr
func WithNoStderr() CommandOptions {
	return func(o *cmdOptions) {
		o.errWriter = io.Discard
	}
}

func runCommand(commandPath string, args []string, opts *cmdOptions) (bytes.Buffer, bytes.Buffer, error) {
	command := exec.Command(commandPath, args...)

	var stderr bytes.Buffer
	var stdout bytes.Buffer

	wout := io.MultiWriter(&stdout, os.Stdout)
	werr := io.MultiWriter(&stderr, os.Stderr)

	if opts.outWriter != nil {
		wout = io.MultiWriter(&stdout, opts.outWriter)
	}
	if opts.errWriter != nil {
		werr = io.MultiWriter(&stderr, opts.errWriter)
	}

	command.Stdout = wout
	command.Stderr = werr

	return stdout, stderr, command.Run()
}

// resourceOptions specifies the resources to use for kubeconfig generation
type resourceOptions struct {
	// projectName name of the Project
	projectName string
	// spaceName name of the Space
	spaceName string
	// clusterGroupName name of the ClusterGroup
	clusterGroupName string
	// customPath use specified path when constructing kubeconfig
	customPath string
}

type ResourceOptions func(o *resourceOptions)

func ForProject(projectName string) ResourceOptions {
	return func(o *resourceOptions) {
		o.projectName = strings.TrimSpace(projectName)
	}
}
func ForSpace(spaceName string) ResourceOptions {
	return func(o *resourceOptions) {
		o.spaceName = strings.TrimSpace(spaceName)
	}
}
func ForClusterGroup(clusterGroupName string) ResourceOptions {
	return func(o *resourceOptions) {
		o.clusterGroupName = strings.TrimSpace(clusterGroupName)
	}
}
func ForCustomPath(customPath string) ResourceOptions {
	return func(o *resourceOptions) {
		o.customPath = customPath
	}
}

// GetKubeconfigForContext returns the kubeconfig for any arbitrary Tanzu resource in the Tanzu object hierarchy
// referred by the Tanzu context
// Pre-reqs: project, space and clustergroup names should be valid
//
// Notes:
//
// Use Case 1: Get the kubeconfig pointing to Tanzu org
// -> projectName        = ""
// -> spaceName          = ""
// -> clusterGroupName   = ""
// ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid
//
// Use Case 2: Get the kubeconfig pointing to Tanzu project
// -> projectName        = "PROJECTNAME"
// -> spaceName          = ""
// -> clusterGroupName   = ""
// ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid/project/<projectName>
//
// Use Case 3: Get the kubeconfig pointing to Tanzu space
// -> projectName        = "PROJECTNAME"
// -> spaceName          = "SPACENAME"
// -> clusterGroupName   = ""
// ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid/project/<projectName>/space/<spaceName>
//
// Use Case 4: Get the kubeconfig pointing to Tanzu clustergroup
// -> projectName        = "PROJECTNAME"
// -> spaceName          = ""
// -> clusterGroupName   = "CLUSTERGROUPNAME"
// ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid/project/<projectName>/clustergroup/<clustergroupName>
//
// Note: Specifying `spaceName` and `clusterGroupName` both at the same time is incorrect input.
func GetKubeconfigForContext(contextName string, opts ...ResourceOptions) ([]byte, error) {
	ctx, err := GetContext(contextName)
	if err != nil {
		return nil, err
	}

	rOptions := &resourceOptions{}
	for _, opt := range opts {
		opt(rOptions)
	}

	if ctx.ContextType != configtypes.ContextTypeTanzu && ctx.ContextType != configtypes.ContextTypeK8s {
		return nil, errors.Errorf("context must be of type: %s or %s", configtypes.ContextTypeTanzu, configtypes.ContextTypeK8s)
	}

	if ctx.ContextType == configtypes.ContextTypeTanzu && rOptions.spaceName != "" && rOptions.clusterGroupName != "" {
		return nil, errors.Errorf("incorrect resource options provided. Both space and clustergroup are set but only one can be set")
	}

	kc, err := kubeconfig.ReadKubeConfig(ctx.ClusterOpts.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read the Tanzu context kubeconfig")
	}

	kc, err = kubeconfig.MinifyKubeConfig(kc, ctx.ClusterOpts.Context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to minify the kubeconfig")
	}

	if ctx.ContextType == configtypes.ContextTypeTanzu {
		updateKubeconfigServerURL(kc, ctx, rOptions)
	}

	kubeconfigBytes, err := yaml.Marshal(kc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal the kubeconfig")
	}
	return kubeconfigBytes, nil
}

func prepareClusterServerURL(context *configtypes.Context, rOptions *resourceOptions) string {
	serverURL := context.ClusterOpts.Endpoint

	// If customPath is set, append customPath after endpoint to form endpoint URL
	if rOptions.customPath != "" {
		return fmt.Sprintf("%s/%s", strings.TrimRight(context.GlobalOpts.Endpoint, "/"), strings.TrimLeft(rOptions.customPath, "/"))
	}

	if rOptions.projectName == "" {
		return serverURL
	}
	serverURL = serverURL + "/project/" + rOptions.projectName

	if rOptions.spaceName != "" {
		return serverURL + "/space/" + rOptions.spaceName
	}
	if rOptions.clusterGroupName != "" {
		return serverURL + "/clustergroup/" + rOptions.clusterGroupName
	}
	return serverURL
}

func updateKubeconfigServerURL(kc *kubeconfig.Config, cliContext *configtypes.Context, rOptions *resourceOptions) {
	currentContextName := kc.CurrentContext
	context := kubeconfig.GetContext(kc, currentContextName)
	cluster := kubeconfig.GetCluster(kc, context.Context.Cluster)
	cluster.Cluster.Server = prepareClusterServerURL(cliContext, rOptions)
}

// SetTanzuContextActiveResource sets the active Tanzu resource for the given context and also updates
// the kubeconfig referenced by the context of type Tanzu
//
// Pre-reqs: project and space/clustergroup names should be valid
//
// Note: To set
//   - a space as active resource, both project and space names are required
//   - a clustergroup as active resource, both project and clustergroup names are required
//   - a project as active resource, only project name is required (space should be empty string)
//   - org as active resource, project, space and clustergroup names should be empty strings
func SetTanzuContextActiveResource(contextName string, resourceInfo ResourceInfo, opts ...CommandOptions) error {
	// For now, the implementation expects env var TANZU_BIN to be set and
	// pointing to the core CLI binary used to invoke setting the active Tanzu resource.

	options := &cmdOptions{}
	for _, opt := range opts {
		opt(options)
	}

	cliPath := os.Getenv("TANZU_BIN")
	if cliPath == "" {
		return errors.New("the environment variable TANZU_BIN is not set")
	}

	altCommandArgs := []string{customCommandName}
	args := []string{"context", "update", "tanzu-active-resource", contextName, "--project", resourceInfo.ProjectName}
	if resourceInfo.SpaceName != "" {
		args = append(args, "--space", resourceInfo.SpaceName)
	}
	if resourceInfo.ClusterGroupName != "" {
		args = append(args, "--clustergroup", resourceInfo.ClusterGroupName)
	}

	altCommandArgs = append(altCommandArgs, args...)

	// Check if there is an alternate means to set the active Tanzu context active resource
	// operation, if not fall back to `context update tanzu-active-resource`
	stdoutOutput, _, err := runCommand(cliPath, altCommandArgs, &cmdOptions{outWriter: io.Discard, errWriter: io.Discard})
	if err == nil {
		args = strings.Fields(stdoutOutput.String())
	}

	// Runs the actual command
	_, stderrOutput, err := runCommand(cliPath, args, options)
	if err != nil {
		return errors.New(stderrOutput.String())
	}
	return nil
}

// GetTanzuContextActiveResource returns the Tanzu active resource information for the given context
func GetTanzuContextActiveResource(contextName string) (*ResourceInfo, error) {
	ctx, err := GetContext(contextName)
	if err != nil {
		return nil, err
	}
	if ctx.ContextType != configtypes.ContextTypeTanzu {
		return nil, errors.Errorf("context must be of type: %s", configtypes.ContextTypeTanzu)
	}
	if ctx.AdditionalMetadata == nil {
		return nil, errors.New("context is missing the Tanzu metadata")
	}
	activeResourceInfo := &ResourceInfo{
		OrgID:            stringValue(ctx.AdditionalMetadata[OrgIDKey]),
		ProjectName:      stringValue(ctx.AdditionalMetadata[ProjectNameKey]),
		SpaceName:        stringValue(ctx.AdditionalMetadata[SpaceNameKey]),
		ClusterGroupName: stringValue(ctx.AdditionalMetadata[ClusterGroupNameKey]),
	}
	return activeResourceInfo, nil
}

func stringValue(val interface{}) string {
	if val == nil {
		return ""
	}
	str, valid := val.(string)
	if !valid {
		return ""
	}
	return str
}
