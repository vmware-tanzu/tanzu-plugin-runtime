// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"

	"github.com/stretchr/testify/assert"
)

func TestStoreClientConfig(t *testing.T) {
	cfg, expectedCfg, cfg2, expectedCfg2, c := setupStoreClientConfigData()

	// Setup config data
	cfgTestFiles, cleanUp := setupTestConfig(t, &CfgTestData{cfg: cfg, cfgNextGen: cfg2})

	defer func() {
		cleanUp()
	}()

	// Action
	err := StoreClientConfig(c)
	assert.NoError(t, err)

	file, err := os.ReadFile(cfgTestFiles[0].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg, string(file))

	file, err = os.ReadFile(cfgTestFiles[1].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg2, string(file))
}

func setupStoreClientConfigData() (string, string, string, string, *types.ClientConfig) {
	cfg := `clientOptions:
  cli:
    discoverySources:
      - oci:
          name: default
          image: "/:"
          unknown: cli-unknown
        contextType: k8s
      - local:
          name: default-local
        contextType: k8s
      - local:
          name: admin-local
          path: admin
servers:
  - name: test-mc
    type: managementcluster
    managementClusterOpts:
      endpoint: updated-test-endpoint
      path: updated-test-path
      context: updated-test-context
      annotation: one
      required: true
    discoverySources:
      - gcp:
          name: test
          bucket: updated-test-bucket
          manifestPath: updated-test-manifest-path
          annotation: one
          required: true
        contextType: tmc
current: test-mc
`

	cfg2 := `contexts:
  - name: test-mc
    target: kubernetes
    group: one
    clusterOpts:
      isManagementCluster: true
      annotation: one
      required: true
      annotationStruct:
        one: one
      endpoint: test-endpoint
      path: test-path
      context: test-context
    discoverySources:
      - gcp:
          name: test
          bucket: test-bucket
          manifestPath: test-manifest-path
          annotation: one
          required: true
        contextType: tmc
currentContext:
  kubernetes: test-mc
`
	expectedCfg := `clientOptions:
    cli:
        discoverySources:
            - oci:
                name: default
                image: "/:"
                unknown: cli-unknown
              contextType: k8s
            - local:
                name: default-local
              contextType: k8s
            - local:
                name: admin-local
                path: admin
        repositories:
            - gcpPluginRepository:
                name: test
                bucketName: bucket
                rootPath: root-path
        unstableVersionSelector: unstable-version
        edition: test=tkg
        bomRepo: test-bomrepo
        compatibilityFilePath: test-compatibility-file-path
servers:
    - name: test-mc
      type: managementcluster
      managementClusterOpts:
        endpoint: test-endpoint
        path: test-path
        context: test-context
        annotation: one
        required: true
      discoverySources:
        - gcp:
            name: test
            bucket: test-bucket
            manifestPath: test-manifest-path
            annotation: one
            required: true
          contextType: tmc
current: test-mc
`

	expectedCfg2 := `contexts:
    - name: test-mc
      target: kubernetes
      group: one
      clusterOpts:
        isManagementCluster: true
        annotation: one
        required: true
        annotationStruct:
            one: one
        endpoint: test-context-endpoint
        path: test-context-path
        context: test-context
      discoverySources:
        - local:
            name: test
            path: test-local-path
        - gcp:
            name: test2
            bucket: ctx-test-bucket
            manifestPath: ctx-test-manifest-path
currentContext:
    kubernetes: test-mc
`

	c := &types.ClientConfig{
		KnownServers: []*types.Server{
			{
				Name: "test-mc",
				Type: types.ManagementClusterServerType,
				ManagementClusterOpts: &types.ManagementClusterServer{
					Endpoint: "test-endpoint",
					Context:  "test-context",
					Path:     "test-path",
				},
				DiscoverySources: []types.PluginDiscovery{
					{
						GCP: &types.GCPDiscovery{
							Name:         "test",
							Bucket:       "test-bucket",
							ManifestPath: "test-manifest-path",
						},
					},
				},
			},
		},
		CurrentServer: "test-mc",
		KnownContexts: []*types.Context{
			{
				Name:   "test-mc",
				Target: types.TargetK8s,
				ClusterOpts: &types.ClusterServer{
					Endpoint:            "test-context-endpoint",
					Path:                "test-context-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
				DiscoverySources: []types.PluginDiscovery{
					{
						GCP: &types.GCPDiscovery{
							Name:         "test2",
							Bucket:       "ctx-test-bucket",
							ManifestPath: "ctx-test-manifest-path",
						},
					},
					{
						Local: &types.LocalDiscovery{
							Name: "test",
							Path: "test-local-path",
						},
					},
				},
			},
		},
		CurrentContext: map[types.Target]string{
			types.TargetK8s: "test-mc",
		},
		ClientOptions: &types.ClientOptions{
			CLI: &types.CLIOptions{
				Repositories: []types.PluginRepository{
					{
						GCPPluginRepository: &types.GCPPluginRepository{
							Name:       "test",
							BucketName: "bucket",
							RootPath:   "root-path",
						},
					},
				},
				DiscoverySources: []types.PluginDiscovery{
					{
						GCP: &types.GCPDiscovery{
							Name:         "test",
							Bucket:       "ctx-test-bucket",
							ManifestPath: "ctx-test-manifest-path",
						},
					},
				},
				UnstableVersionSelector: types.VersionSelectorLevel("unstable-version"),
				Edition:                 types.EditionSelector("test=tkg"),
				BOMRepo:                 "test-bomrepo",
				CompatibilityFilePath:   "test-compatibility-file-path",
			},
		},
	}
	return cfg, expectedCfg, cfg2, expectedCfg2, c
}
