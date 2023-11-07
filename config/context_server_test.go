// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestContextsServersSyncWhenNoServersExist(t *testing.T) {
	// Setup config data
	cfg, expectedCfg, cfg2, expectedCfg2 := func() (string, string, string, string) {
		cfg := ``
		expectedCfg := `servers:
    - name: test-mc2
      type: managementcluster
      managementClusterOpts:
        endpoint: test-endpoint
        path: test-path
current: test-mc2
`
		cfg2 := `contexts:
  - name: test-mc
    target: kubernetes
    clusterOpts:
      isManagementCluster: true
      path: test-path
      endpoint: test-endpoint
currentContext:
  kubernetes: test-mc
`
		expectedCfg2 := `contexts:
    - name: test-mc
      target: kubernetes
      clusterOpts:
        isManagementCluster: true
        path: test-path
        endpoint: test-endpoint
    - name: test-mc2
      target: kubernetes
      clusterOpts:
        endpoint: test-endpoint
        path: test-path
        isManagementCluster: true
currentContext:
    kubernetes: test-mc2
`
		return cfg, expectedCfg, cfg2, expectedCfg2
	}()

	cfgTestFiles, cleanUp := setupTestConfig(t, &CfgTestData{cfg: cfg, cfgNextGen: cfg2})

	defer func() {
		cleanUp()
	}()

	// Get Context
	context, err := GetContext("test-mc")
	expected := &configtypes.Context{
		Name:   "test-mc",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Endpoint:            "test-endpoint",
			Path:                "test-path",
			IsManagementCluster: true,
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, context)

	// Get Server
	server, err := GetServer("test-mc")
	assert.Nil(t, server)
	assert.Equal(t, "could not find server \"test-mc\"", err.Error())

	// Add new Context
	newCtx := &configtypes.Context{
		Name:   "test-mc2",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Path:                "test-path",
			Endpoint:            "test-endpoint",
			IsManagementCluster: true,
		},
	}

	newServer := &configtypes.Server{
		Name: "test-mc2",
		Type: configtypes.ManagementClusterServerType,
		ManagementClusterOpts: &configtypes.ManagementClusterServer{
			Endpoint: "test-endpoint",
			Path:     "test-path",
		},
	}

	err = SetContext(newCtx, true)
	assert.NoError(t, err)

	ctx, err := GetContext("test-mc2")
	assert.Nil(t, err)
	assert.Equal(t, newCtx, ctx)

	server, err = GetServer("test-mc2")
	assert.Nil(t, err)
	assert.Equal(t, newServer, server)

	//Read config files
	file, err := os.ReadFile(cfgTestFiles[0].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg, string(file))

	file, err = os.ReadFile(cfgTestFiles[1].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg2, string(file))

	// Delete context
	err = DeleteContext("test-mc2")
	assert.NoError(t, err)
	ctx, err = GetContext("test-mc2")
	assert.Equal(t, "context test-mc2 not found", err.Error())
	assert.Nil(t, ctx)
}
func TestContextsServersSyncWhenNoContextsExist(t *testing.T) {
	// Setup config data
	cfg, expectedCfg, cfg2, expectedCfg2 := func() (string, string, string, string) {
		cfg := `servers:
    - name: test-mc
      type: managementcluster
      managementClusterOpts:
        endpoint: test-endpoint
        path: test-path
current: test-mc
`
		expectedCfg := `servers:
    - name: test-mc
      type: managementcluster
      managementClusterOpts:
        endpoint: test-endpoint
        path: test-path
    - name: test-mc2
      type: managementcluster
      managementClusterOpts:
        endpoint: test-endpoint
        path: test-path
current: test-mc2
`
		cfg2 := ``
		expectedCfg2 := `contexts:
    - name: test-mc2
      target: kubernetes
      clusterOpts:
        endpoint: test-endpoint
        path: test-path
        isManagementCluster: true
currentContext:
    kubernetes: test-mc2
`
		return cfg, expectedCfg, cfg2, expectedCfg2
	}()

	cfgTestFiles, cleanUp := setupTestConfig(t, &CfgTestData{cfg: cfg, cfgNextGen: cfg2})

	defer func() {
		cleanUp()
	}()

	// Get Server
	expectedServer := &configtypes.Server{
		Name: "test-mc",
		Type: configtypes.ManagementClusterServerType,
		ManagementClusterOpts: &configtypes.ManagementClusterServer{
			Endpoint: "test-endpoint",
			Path:     "test-path",
		},
	}

	server, err := GetServer("test-mc")
	assert.Nil(t, err)
	assert.Equal(t, expectedServer, server)

	// Get Context
	context, err := GetContext("test-mc")
	assert.Nil(t, context)
	assert.Equal(t, "context test-mc not found", err.Error())

	// Add new Server
	newServer := &configtypes.Server{
		Name: "test-mc2",
		Type: configtypes.ManagementClusterServerType,
		ManagementClusterOpts: &configtypes.ManagementClusterServer{
			Endpoint: "test-endpoint",
			Path:     "test-path",
		},
	}
	newCtx := &configtypes.Context{
		Name:   "test-mc2",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Path:                "test-path",
			Endpoint:            "test-endpoint",
			IsManagementCluster: true,
		},
	}

	err = SetServer(newServer, true)
	assert.NoError(t, err)

	server, err = GetServer("test-mc2")
	assert.Nil(t, err)
	assert.Equal(t, newServer, server)

	ctx, err := GetContext("test-mc2")
	assert.Nil(t, err)
	assert.Equal(t, newCtx, ctx)

	//Read config files
	file, err := os.ReadFile(cfgTestFiles[0].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg, string(file))

	file, err = os.ReadFile(cfgTestFiles[1].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg2, string(file))

	// Delete context
	err = DeleteContext("test-mc2")
	assert.NoError(t, err)
	ctx, err = GetContext("test-mc2")
	assert.Equal(t, "context test-mc2 not found", err.Error())
	assert.Nil(t, ctx)
}
