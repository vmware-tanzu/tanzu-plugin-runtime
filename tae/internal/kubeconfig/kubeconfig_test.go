// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package kubeconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const ConfigFilePermissions = 0o600

func copyFile(t *testing.T, sourceFile, destFile string) {
	input, err := os.ReadFile(sourceFile)
	assert.NoError(t, err)
	err = os.WriteFile(destFile, input, ConfigFilePermissions)
	assert.NoError(t, err)
}

func TestReadAndMinifyKubeConfig(t *testing.T) {
	expectedMytaeKubeconfig := `
kind: Config
apiVersion: v1
preferences: {}
clusters:
    - name: tanzu-cli-mytae/current
      cluster:
        server: https://api.tanzu.cloud.vmware.com:443/org/fake-org-id
        insecure-skip-tls-verify: true
users:
    - name: tanzu-cli-mytae-user
      user:
        exec:
            command: tanzu
            args:
                - context
                - get-token
                - mytae
            env: []
            apiVersion: client.authentication.k8s.io/v1beta1
            provideClusterInfo: false
contexts:
    - name: tanzu-cli-mytae
      context:
        cluster: tanzu-cli-mytae/current
        user: tanzu-cli-mytae-user
current-context: tanzu-cli-mytae
`

	expectedFooContextKubeconfig := `
kind: Config
apiVersion: v1
preferences: {}
clusters:
    - name: foo-cluster
      cluster:
        server: https://foo.org:443
        insecure-skip-tls-verify: true
users:
    - name: blue-user
      user:
        token: blue-token
contexts:
    - name: foo-context
      context:
        cluster: foo-cluster
        user: blue-user
        namespace: saw-ns
current-context: foo-context
`

	testKubeconfiFilePath := "../../../fakes/config/kubeconfig-1.yaml"
	kubeconfigFilePath, err := os.CreateTemp("", "config")
	assert.NoError(t, err)
	copyFile(t, testKubeconfiFilePath, kubeconfigFilePath.Name())

	defer func() {
		_ = os.RemoveAll(kubeconfigFilePath.Name())
	}()

	config, err := ReadKubeConfig(kubeconfigFilePath.Name())
	assert.NoError(t, err)

	// Test with non-existing context
	_, err = MinifyKubeConfig(config, "non-existing-context")
	assert.Equal(t, err.Error(), `context "non-existing-context" missing in the kubeconfig`)

	// Test reading and minifying the kubeconfig using Token as user(AuthInfo)
	minifiedKubeconfig, err := MinifyKubeConfig(config, "foo-context")
	assert.NoError(t, err)
	wantKubeConfig := &Config{}
	err = yaml.Unmarshal([]byte(expectedFooContextKubeconfig), wantKubeConfig)
	assert.NoError(t, err)
	assert.Equal(t, minifiedKubeconfig, wantKubeConfig)

	// Test reading and minifying the kubeconfig having ExecConfig as user(AuthInfo)
	minifiedKubeconfig, err = MinifyKubeConfig(config, "tanzu-cli-mytae")
	assert.NoError(t, err)
	wantKubeConfig = &Config{}
	err = yaml.Unmarshal([]byte(expectedMytaeKubeconfig), wantKubeConfig)
	assert.NoError(t, err)
	assert.Equal(t, minifiedKubeconfig, wantKubeConfig)
}
