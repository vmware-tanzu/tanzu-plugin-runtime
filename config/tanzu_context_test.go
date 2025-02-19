// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/internal/kubeconfig"
)

const (
	fakeOrgID            = "fake-org-id"
	fakeOrgName          = "fake-org-name"
	fakeProjectName      = "fake-project"
	fakeProjectID        = "fake-project-id"
	fakeSpace            = "fake-space"
	fakeClusterGroupName = "fake-clustergroup"

	fakePluginScriptFmtString string = `#!/bin/bash
# Fake tanzu core binary

# fake command that simulates a context lcm operation
context() {
	if [ "%s" -eq "0" ]; then
		# regular output to stderr
		>&2 echo "$@ succeeded"
	else
		# error to stderr
		>&2 echo "$@ failed"
	fi

	exit %s
}

# fake alternate command to use
newcommand() {
	if [ "%s" -eq "0" ]; then
		# regular output to stdout
		echo "$@ succeeded"
	else
		# error to stderr
		>&2 echo "$@ failed"
	fi

	exit %s
}

case "$1" in
    # simulate returning an alternative set of args to invoke with, which
    # translates to running the command 'newcommand'
    %s) shift && shift && echo "newcommand $@";;
    newcommand)   $1 "$@";;
    context)   $1 "$@";;
    *) cat << EOF
Tanzu Core CLI Fake

Usage:
  tanzu [command]

Available Commands:
  update          fake command
  newcommand      fake new command
  _custom_command provide alternate command to invoke, if available
EOF
       exit 1
       ;;
esac
`
)

func cleanupTestingDir(t *testing.T) {
	p, err := LocalDir()
	assert.NoError(t, err)
	err = os.RemoveAll(p)
	assert.NoError(t, err)
}

func readOutput(t *testing.T, r io.Reader, c chan<- []byte) {
	data, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	c <- data
}

func TestGetKubeconfigForContext(t *testing.T) {
	err := setupForGetContext()
	assert.NoError(t, err)

	testKubeconfiFilePath := "../fakes/config/kubeconfig-1.yaml"
	kubeconfigFilePath, err := os.CreateTemp("", "config")
	assert.NoError(t, err)
	err = copyFile(testKubeconfiFilePath, kubeconfigFilePath.Name())
	assert.NoError(t, err)

	defer func() {
		cleanupTestingDir(t)
		_ = os.RemoveAll(kubeconfigFilePath.Name())
	}()

	// Cluster context
	myTanzu := "tanzu-cli-mytanzu"

	// Setup tanzu context
	c, err := GetContext("test-tanzu")
	assert.NoError(t, err)
	c.ClusterOpts.Path = kubeconfigFilePath.Name()
	c.ClusterOpts.Context = myTanzu
	err = SetContext(c, false)
	assert.NoError(t, err)

	// Setup kubernetes context
	k8sCtx, err := GetContext("k8s-context")
	assert.NoError(t, err)
	k8sCtx.ClusterOpts.Path = kubeconfigFilePath.Name()
	k8sCtx.ClusterOpts.Context = "k8s-context"
	err = SetContext(k8sCtx, false)
	assert.NoError(t, err)

	// Test getting the kubeconfig for a space within a project
	kubeconfigBytes, err := GetKubeconfigForContext(c.Name, ForProject("project1-id"), ForSpace("space1"))
	assert.NoError(t, err)
	c, err = GetContext("test-tanzu")
	assert.NoError(t, err)
	var kc kubeconfig.Config
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster := kubeconfig.GetCluster(&kc, "tanzu-cli-mytanzu/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project1-id/space/space1")

	// Test getting the kubeconfig for a project
	kubeconfigBytes, err = GetKubeconfigForContext(c.Name, ForProject("project2-id"))
	assert.NoError(t, err)
	c, err = GetContext("test-tanzu")
	assert.NoError(t, err)
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster = kubeconfig.GetCluster(&kc, "tanzu-cli-mytanzu/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project2-id")

	// Test getting the kubeconfig for a clustergroup within a project
	kubeconfigBytes, err = GetKubeconfigForContext(c.Name, ForProject("project2-id"), ForClusterGroup("clustergroup1"))
	assert.NoError(t, err)
	c, err = GetContext("test-tanzu")
	assert.NoError(t, err)
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster = kubeconfig.GetCluster(&kc, "tanzu-cli-mytanzu/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project2-id/clustergroup/clustergroup1")

	// Test getting the kubeconfig for a foundation within a project
	kubeconfigBytes, err = GetKubeconfigForContext(c.Name, ForProject("project2-id"), ForFoundationGroup("foundationgroup1"))
	assert.NoError(t, err)
	c, err = GetContext("test-tanzu")
	assert.NoError(t, err)
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster = kubeconfig.GetCluster(&kc, "tanzu-cli-mytanzu/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project2-id/foundationgroup/foundationgroup1")

	// Test getting the kubeconfig with incorrect resource combination (request kubeconfig for space and clustergroup)
	c, err = GetContext("test-tanzu")
	assert.NoError(t, err)
	_, err = GetKubeconfigForContext(c.Name, ForProject("project2-id"), ForFoundationGroup("foundationgroup1"), ForClusterGroup("clustergroup1"))
	assert.Error(t, err)
	assert.ErrorContains(t, err, "incorrect resource options provided. Only one of space, clustergroup, or foundationgroup can be set. Provided configuration - space: , clustergroup: clustergroup1, foundationgroup: foundationgroup1")

	// Test getting the kubeconfig with incorrect resource combination (request kubeconfig for space and clustergroup)
	c, err = GetContext("test-tanzu")
	assert.NoError(t, err)
	_, err = GetKubeconfigForContext(c.Name, ForProject("project2-id"), ForSpace("space1"), ForClusterGroup("clustergroup1"))
	assert.Error(t, err)
	assert.ErrorContains(t, err, "incorrect resource options provided. Only one of space, clustergroup, or foundationgroup can be set. Provided configuration - space: space1, clustergroup: clustergroup1, foundationgroup: ")

	// Test getting the kubeconfig for an arbitrary Tanzu resource for non Tanzu context
	tmcCtx, err := GetContext("test-tmc")
	assert.NoError(t, err)
	_, err = GetKubeconfigForContext(tmcCtx.Name, ForProject("project2-id"))
	assert.Error(t, err)
	assert.ErrorContains(t, err, "context must be of type: tanzu or kubernetes")

	// Test getting the kubeconfig for a custom endpoint path
	tanzuCtx, err := GetContext("test-tanzu")
	assert.NoError(t, err)
	kubeconfigBytes, err = GetKubeconfigForContext(tanzuCtx.Name, ForCustomPath("/custom-path"))
	assert.NoError(t, err)
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster = kubeconfig.GetCluster(&kc, "tanzu-cli-mytanzu/current")
	assert.Equal(t, cluster.Cluster.Server, tanzuCtx.GlobalOpts.Endpoint+"/custom-path")

	// Test getting the kubeconfig for a kubernetes context
	k8sCtx, err = GetContext("k8s-context")
	assert.NoError(t, err)
	kubeconfigBytes, err = GetKubeconfigForContext(k8sCtx.Name)
	assert.NoError(t, err)
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster = kubeconfig.GetCluster(&kc, "k8s-cluster")
	assert.Equal(t, cluster.Cluster.Server, k8sCtx.ClusterOpts.Endpoint)

	// Test if tanzu context ClusterOpts is set to nil
	tanzuCtx, err = GetContext("test-tanzu")
	assert.NoError(t, err)
	// update context to remove clusterOtps
	tanzuCtx.Name = "test-tanzu-withoutClusterOpts"
	tanzuCtx.ClusterOpts = nil
	err = AddContext(tanzuCtx, false)
	assert.NoError(t, err)
	_, err = GetKubeconfigForContext(tanzuCtx.Name)
	assert.ErrorContains(t, err, "invalid context. context missing kubeconfig details")
	err = DeleteContext(tanzuCtx.Name)
	assert.NoError(t, err)
}

func TestGetTanzuContextActiveResource(t *testing.T) {
	err := setupForGetContext()
	assert.NoError(t, err)
	defer cleanupTestingDir(t)

	c, err := GetContext("test-tanzu")
	assert.NoError(t, err)

	// Test getting the Tanzu active resource of a non-existent context
	_, err = GetTanzuContextActiveResource("non-existent-context")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "context non-existent-context not found")

	// Test getting the Tanzu active resource of a context that is not Tanzu context
	_, err = GetTanzuContextActiveResource("test-mc")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "context must be of type: tanzu")

	// Test getting the Tanzu active resource of a context with active resource as Org only
	activeResources, err := GetTanzuContextActiveResource("test-tanzu")
	assert.NoError(t, err)
	assert.Equal(t, activeResources.OrgID, fakeOrgID)
	assert.Equal(t, activeResources.OrgName, fakeOrgName)
	assert.Empty(t, activeResources.ProjectName)
	assert.Empty(t, activeResources.ProjectID)
	assert.Empty(t, activeResources.SpaceName)

	// Test getting the Tanzu active resource of a context with active resource as space
	c.AdditionalMetadata[ProjectNameKey] = fakeProjectName
	c.AdditionalMetadata[ProjectIDKey] = fakeProjectID
	c.AdditionalMetadata[SpaceNameKey] = fakeSpace
	err = SetContext(c, false)
	assert.NoError(t, err)
	activeResources, err = GetTanzuContextActiveResource("test-tanzu")
	assert.NoError(t, err)
	assert.Equal(t, activeResources.OrgID, fakeOrgID)
	assert.Equal(t, activeResources.OrgName, fakeOrgName)
	assert.Equal(t, activeResources.ProjectName, fakeProjectName)
	assert.Equal(t, activeResources.ProjectID, fakeProjectID)
	assert.Equal(t, activeResources.SpaceName, fakeSpace)
	assert.Equal(t, activeResources.ClusterGroupName, "")

	// Test getting the Tanzu active resource of a context with active resource as clustergroup
	c.AdditionalMetadata[ProjectNameKey] = fakeProjectName
	c.AdditionalMetadata[ProjectIDKey] = fakeProjectID
	c.AdditionalMetadata[SpaceNameKey] = ""
	c.AdditionalMetadata[ClusterGroupNameKey] = fakeClusterGroupName
	err = SetContext(c, false)
	assert.NoError(t, err)
	activeResources, err = GetTanzuContextActiveResource("test-tanzu")
	assert.NoError(t, err)
	assert.Equal(t, activeResources.OrgID, fakeOrgID)
	assert.Equal(t, activeResources.OrgName, fakeOrgName)
	assert.Equal(t, activeResources.ProjectName, fakeProjectName)
	assert.Equal(t, activeResources.ProjectID, fakeProjectID)
	assert.Equal(t, activeResources.ClusterGroupName, fakeClusterGroupName)
	assert.Equal(t, activeResources.SpaceName, "")

	// If context activeMetadata is not set(error case)
	c.AdditionalMetadata = nil
	err = SetContext(c, false)
	assert.NoError(t, err)
	_, err = GetTanzuContextActiveResource("test-tanzu")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "context is missing the Tanzu metadata")
}

func setupFakeCLI(dir string, exitStatus string, newCommandExitStatus string, enableCustomCommand bool) (string, error) {
	filePath := filepath.Join(dir, "tanzu")

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fakeCustomCommandName := "unused_command"
	// when enabled, the fake CLI script generated will be capable of
	// returning an alternate set of args for a provided set of args
	if enableCustomCommand {
		fakeCustomCommandName = customCommandName
	}

	fmt.Fprintf(f, fakePluginScriptFmtString, exitStatus, exitStatus, newCommandExitStatus, newCommandExitStatus, fakeCustomCommandName)

	return filePath, nil
}

func TestSetTanzuContextActiveResource(t *testing.T) {
	tests := []struct {
		test                 string
		exitStatus           string
		newCommandExitStatus string
		expectedOutput       string
		expectedFailure      bool
		enableCustomCommand  bool
	}{
		{
			test:            "with no alternate command and Tanzu active resource update successfully",
			exitStatus:      "0",
			expectedOutput:  "context update tanzu-active-resource test-context --project projectA --project-id projectA-ID --space spaceA succeeded\n",
			expectedFailure: false,
		},
		{
			test:            "with no alternate command and Tanzu active resource update unsuccessfully",
			exitStatus:      "1",
			expectedOutput:  "context update tanzu-active-resource test-context --project projectA --project-id projectA-ID --space spaceA failed\n",
			expectedFailure: true,
		},
		{
			test:                 "with alternate command and Tanzu active resource update successfully",
			newCommandExitStatus: "0",
			expectedOutput:       "newcommand update tanzu-active-resource test-context --project projectA --project-id projectA-ID --space spaceA succeeded\n",
			expectedFailure:      false,
			enableCustomCommand:  true,
		},
		{
			test:                 "with alternate command and Tanzu active resource update unsuccessfully",
			newCommandExitStatus: "1",
			expectedOutput:       "newcommand update tanzu-active-resource test-context --project projectA --project-id projectA-ID --space spaceA failed\n",
			expectedFailure:      true,
			enableCustomCommand:  true,
		},
	}

	for _, spec := range tests {
		dir, err := os.MkdirTemp("", "tanzu-set-tanzu-active-resource-api")
		assert.Nil(t, err)
		defer os.RemoveAll(dir)
		t.Run(spec.test, func(t *testing.T) {
			assert := assert.New(t)

			// Set up stdout and stderr for our test
			r, w, err := os.Pipe()
			if err != nil {
				t.Error(err)
			}
			c := make(chan []byte)
			go readOutput(t, r, c)
			stdout := os.Stdout
			stderr := os.Stderr
			defer func() {
				os.Stdout = stdout
				os.Stderr = stderr
			}()
			os.Stdout = w
			os.Stderr = w

			cliPath, err := setupFakeCLI(dir, spec.exitStatus, spec.newCommandExitStatus, spec.enableCustomCommand)
			assert.Nil(err)
			os.Setenv("TANZU_BIN", cliPath)

			// Test-1:
			// - verify correct string gets printed to default stdout and stderr
			err = SetTanzuContextActiveResource("test-context", ResourceInfo{ProjectName: "projectA", ProjectID: "projectA-ID", SpaceName: "spaceA"})
			w.Close()
			stdoutRecieved := <-c

			if spec.expectedFailure {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
			}

			assert.Equal(spec.expectedOutput, string(stdoutRecieved), "incorrect combinedOutput result")

			// Test-2: when external stdout and stderr are provided with WithStdout, WithStderr options,
			// verify correct string gets printed to provided custom stdout/stderr
			var combinedOutputBuff bytes.Buffer
			err = SetTanzuContextActiveResource("test-context", ResourceInfo{ProjectName: "projectA", ProjectID: "projectA-ID", SpaceName: "spaceA"}, WithOutputWriter(&combinedOutputBuff), WithErrorWriter(&combinedOutputBuff))
			if spec.expectedFailure {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
			}
			assert.Equal(spec.expectedOutput, combinedOutputBuff.String(), "incorrect combinedOutputBuff result")

			os.Unsetenv("TANZU_BIN")
		})
	}
}
