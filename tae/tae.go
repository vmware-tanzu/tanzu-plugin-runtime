// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package tae provides APIs specific to Tanzu Application Engine(TAE)
package tae

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/tae/internal/kubeconfig"
)

// keys to Context's AdditionalMetadata map
const (
	OrgIDKey       = "taeOrgID"
	ProjectNameKey = "taeProjectName"
	SpaceNameKey   = "taeSpaceName"
)

const (
	// customCommandName is the name of the command expected to be implemented
	// by the CLI should there be a need to discover and alternative invocation
	// method
	customCommandName string = "_custom_command"
)

// ResourceInfo resource information of the application engine
type ResourceInfo struct {
	// OrgID ID of the Organization
	OrgID string
	// ProjectName name of the Project
	ProjectName string
	// SpaceName name of the Space
	SpaceName string
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

// GetKubeconfigForContext returns the kubeconfig for any arbitrary TAE resource in the TAE object hierarchy
// referred by the TAE context
// Pre-reqs: project and space names should be valid
//
// Notes:
// If projectName and spaceName is empty string the kubeconfig generated would be pointing to TAE org
//
//	ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid
//
// If projectName is valid projectName and spaceName is empty string the kubeconfig generated would be pointing to TAE project
//
//	ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid/project/<projectName>
//
// similarly if both project and space names are valid names the kubeconfig generated would be pointing to TAE space
//
//	ex: kubeconfig's cluster.server URL:  https://endpoint/org/orgid/project/<projectName>/space/<spaceName>
func GetKubeconfigForContext(contextName, projectName, spaceName string) ([]byte, error) {
	ctx, err := config.GetContext(contextName)
	if err != nil {
		return nil, err
	}
	if ctx.ContextType != configtypes.ContextTypeTAE {
		return nil, errors.Errorf("context must be of type: %s", configtypes.ContextTypeTAE)
	}

	kc, err := kubeconfig.ReadKubeConfig(ctx.ClusterOpts.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read the TAE context kubeconfig")
	}

	kc, err = kubeconfig.MinifyKubeConfig(kc, ctx.ClusterOpts.Context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to minify the kubeconfig")
	}
	updateKubeconfigServerURL(kc, ctx, projectName, spaceName)

	kubeconfigBytes, err := yaml.Marshal(kc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal the kubeconfig")
	}
	return kubeconfigBytes, nil
}

func prepareClusterServerURL(context *configtypes.Context, projectName, spaceName string) string {
	serverURL := context.ClusterOpts.Endpoint
	if projectName == "" {
		return serverURL
	}
	serverURL = serverURL + "/project/" + projectName

	if spaceName == "" {
		return serverURL
	}
	return serverURL + "/space/" + spaceName
}

func updateKubeconfigServerURL(kc *kubeconfig.Config, cliContext *configtypes.Context, projectName, spaceName string) {
	currentContextName := kc.CurrentContext
	context := kubeconfig.GetContext(kc, currentContextName)
	cluster := kubeconfig.GetCluster(kc, context.Context.Cluster)
	cluster.Cluster.Server = prepareClusterServerURL(cliContext, projectName, spaceName)
}

// SetTAEContextActiveResource sets the active TAE resource for the given context and also updates
// the kubeconfig referrenced by the TAE context
//
// Pre-reqs: project and space names should be valid
//
// Note: To set
//   - a space as active resource, both project and space names are required
//   - a project as active resource, only project name is required (space should be empty string)
//   - org as active resource, both project and space names should be empty strings
func SetTAEContextActiveResource(contextName, projectName, spaceName string, opts ...CommandOptions) error {
	// For now, the implementation expects env var TANZU_BIN to be set and
	// pointing to the core CLI binary used to invoke setting the active TAE resource.

	options := &cmdOptions{}
	for _, opt := range opts {
		opt(options)
	}

	cliPath := os.Getenv("TANZU_BIN")
	if cliPath == "" {
		return errors.New("the environment variable TANZU_BIN is not set")
	}

	altCommandArgs := []string{customCommandName}
	args := []string{"context", "update", "tae-active-resource", contextName, "--project", projectName, "--space", spaceName}

	altCommandArgs = append(altCommandArgs, args...)

	// Check if there is an alternate means to set the active TAE context active resource
	// operation, if not fall back to `context update tae-active-resource`
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

// GetTAEContextActiveResource returns the application engine active resource information for the given context
func GetTAEContextActiveResource(contextName string) (*ResourceInfo, error) {
	ctx, err := config.GetContext(contextName)
	if err != nil {
		return nil, err
	}
	if ctx.ContextType != configtypes.ContextTypeTAE {
		return nil, errors.Errorf("context must be of type: %s", configtypes.ContextTypeTAE)
	}
	if ctx.AdditionalMetadata == nil {
		return nil, errors.New("context is missing the TAE metadata")
	}
	activeResourceInfo := &ResourceInfo{
		OrgID:       stringValue(ctx.AdditionalMetadata[OrgIDKey]),
		ProjectName: stringValue(ctx.AdditionalMetadata[ProjectNameKey]),
		SpaceName:   stringValue(ctx.AdditionalMetadata[SpaceNameKey]),
	}
	return activeResourceInfo, nil
}

func stringValue(val interface{}) string {
	if val == nil {
		return ""
	}
	return val.(string)
}
