// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ucp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/ucp/internal/kubeconfig"
)

const ConfigFilePermissions = 0o600

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
				Name:   "test-ucp",
				Target: configtypes.TargetUCP,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "https://api.tanzu.cloud.vmware.com:443/org/fake-org-id",
					Path:     "test-path",
					Context:  "test-context",
				},
				AdditionalMetadata: map[string]interface{}{
					OrgID: "fake-org-id",
				},
			},
		},
		CurrentContext: map[configtypes.Target]string{
			configtypes.TargetK8s: "test-mc-2",
			configtypes.TargetTMC: "test-tmc",
			configtypes.TargetUCP: "test-ucp",
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

	c, err := config.GetContext("test-ucp")
	assert.NoError(t, err)
	c.ClusterOpts.Path = kubeconfigFilePath.Name()
	c.ClusterOpts.Context = "tanzu-cli-myucp"
	err = config.SetContext(c, false)
	assert.NoError(t, err)

	// Test getting the kubeconfig for an arbitrary UCP resource
	kubeconfigBytes, err := GetKubeconfigForContext(c.Name, "project1", "space1")
	assert.NoError(t, err)
	c, err = config.GetContext("test-ucp")
	assert.NoError(t, err)
	var kc kubeconfig.Config
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster := kubeconfig.GetCluster(&kc, "tanzu-cli-myucp/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project1/space/space1")

	// Test getting the kubeconfig for an arbitrary UCP resource
	kubeconfigBytes, err = GetKubeconfigForContext(c.Name, "project2", "")
	assert.NoError(t, err)
	c, err = config.GetContext("test-ucp")
	assert.NoError(t, err)
	err = yaml.Unmarshal(kubeconfigBytes, &kc)
	assert.NoError(t, err)
	cluster = kubeconfig.GetCluster(&kc, "tanzu-cli-myucp/current")
	assert.Equal(t, cluster.Cluster.Server, c.ClusterOpts.Endpoint+"/project/project2")

	// Test getting the kubeconfig for an arbitrary UCP resource for non UCP context
	nonUCPCtx, err := config.GetContext("test-mc")
	assert.NoError(t, err)
	_, err = GetKubeconfigForContext(nonUCPCtx.Name, "project2", "")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "context must be of type: ucp")
}
