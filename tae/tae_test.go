// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tae

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/tae/internal/kubeconfig"
)

const ConfigFilePermissions = 0o600
const (
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
	p, err := config.LocalDir()
	assert.NoError(t, err)
	err = os.RemoveAll(p)
	assert.NoError(t, err)
}

func copyFile(t *testing.T, sourceFile, destFile string) {
	input, err := os.ReadFile(sourceFile)
	assert.NoError(t, err)
	err = os.WriteFile(destFile, input, ConfigFilePermissions)
	assert.NoError(t, err)
}
func readOutput(t *testing.T, r io.Reader, c chan<- []byte) {
	data, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	c <- data
}

func setupForGetContext(t *testing.T) {
	// setup
	cfg := &configtypes.ClientConfig{
		KnownContexts: []*configtypes.Context{
			{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
			{
				Name:   "test-mc-2",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint-2",
					Path:                "test-path-2",
					Context:             "test-context-2",
					IsManagementCluster: true,
				},
			},
			{
				Name:   "test-tmc",
				Target: configtypes.TargetTMC,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
			{
				Name:   "test-tae",
				Target: configtypes.TargetTAE,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "https://api.tanzu.cloud.vmware.com:443/org/fake-org-id",
					Path:     "test-path",
					Context:  "test-context",
				},
				AdditionalMetadata: map[string]interface{}{
					OrgIDKey: "fake-org-id",
				},
			},
		},
		CurrentContext: map[configtypes.ContextType]string{
			configtypes.ContextTypeK8s: "test-mc-2",
			configtypes.ContextTypeTMC: "test-tmc",
			configtypes.ContextTypeTAE: "test-tae",
		},
	}
	func() {
		config.LocalDirName = config.TestLocalDirName
		err := config.StoreClientConfig(cfg)
		assert.NoError(t, err)
	}()
}

func TestGetKubeconfigForContext(t *testing.T) {
	setupForGetContext(t)

	testKubeconfiFilePath := "../fakes/config/kubeconfig-1.yaml"
	kubeconfigFilePath, err := os.CreateTemp("", "config")
	assert.NoError(t, err)
	copyFile(t, testKubeconfiFilePath, kubeconfigFilePath.Name())

	defer func() {
		cleanupTestingDir(t)
		_ = os.RemoveAll(kubeconfigFilePath.Name())
	}()

	c, err := config.GetContext("test-tae")
	assert.NoError(t, err)
	c.ClusterOpts.Path = kubeconfigFilePath.Name()
	c.ClusterOpts.Context = "tanzu-cli-mytae"
	err = config.SetContext(c, false)
	assert.NoError(t, err)

	// Test getting the kubeconfig for an arbitrary TAE resource
	kubeconfigBytes, err := GetKubeconfigForContext(c.Name, "project1", "space1")
	assert.NoError(t, err)
	c, err = config.GetContext("test-tae")
	assert.NoError(t, err)
	var kc kubeconfig.Config
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster := kubeconfig.GetCluster(&kc, "tanzu-cli-mytae/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project1/space/space1")

	// Test getting the kubeconfig for an arbitrary TAE resource
	kubeconfigBytes, err = GetKubeconfigForContext(c.Name, "project2", "")
	assert.NoError(t, err)
	c, err = config.GetContext("test-tae")
	assert.NoError(t, err)
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster = kubeconfig.GetCluster(&kc, "tanzu-cli-mytae/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project2")

	// Test getting the kubeconfig for an arbitrary TAE resource for non TAE context
	nonTAECtx, err := config.GetContext("test-mc")
	assert.NoError(t, err)
	_, err = GetKubeconfigForContext(nonTAECtx.Name, "project2", "")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "context must be of type: application-engine")
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

func TestSetTAEContextActiveResource(t *testing.T) {
	tests := []struct {
		test                 string
		exitStatus           string
		newCommandExitStatus string
		expectedOutput       string
		expectedFailure      bool
		enableCustomCommand  bool
	}{
		{
			test:            "with no alternate command and tae active resource update successfully",
			exitStatus:      "0",
			expectedOutput:  "context update tae-active-resource test-context --project projectA --space spaceA succeeded\n",
			expectedFailure: false,
		},
		{
			test:            "with no alternate command and tae active resource update unsuccessfully",
			exitStatus:      "1",
			expectedOutput:  "context update tae-active-resource test-context --project projectA --space spaceA failed\n",
			expectedFailure: true,
		},
		{
			test:                 "with alternate command and tae active resource update successfully",
			newCommandExitStatus: "0",
			expectedOutput:       "newcommand update tae-active-resource test-context --project projectA --space spaceA succeeded\n",
			expectedFailure:      false,
			enableCustomCommand:  true,
		},
		{
			test:                 "with alternate command and tae active resource update unsuccessfully",
			newCommandExitStatus: "1",
			expectedOutput:       "newcommand update tae-active-resource test-context --project projectA --space spaceA failed\n",
			expectedFailure:      true,
			enableCustomCommand:  true,
		},
	}

	for _, spec := range tests {
		dir, err := os.MkdirTemp("", "tanzu-set-tae-active-resource-api")
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
			err = SetTAEContextActiveResource("test-context", "projectA", "spaceA")
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
			err = SetTAEContextActiveResource("test-context", "projectA", "spaceA", WithOutputWriter(&combinedOutputBuff), WithErrorWriter(&combinedOutputBuff))
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
