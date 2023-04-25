// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func setupConfigData() (string, string, string, string) {
	cfg := `servers:
  - name: test-mc
    type: managementcluster
    managementClusterOpts:
      endpoint: test-endpoint
      path: test-path
      context: test-context
      annotation: one
      required: true
current: test-mc
contexts:
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
currentContext:
  kubernetes: test-mc
`
	expectedCfg := `servers:
    - name: test-mc
      type: managementcluster
      managementClusterOpts:
        endpoint: test-endpoint
        path: test-path
        context: test-context
        annotation: one
        required: true
current: test-mc
`

	cfg2 := `cli:
  discoverySources:
    - oci:
        name: test
        image: image
        annotation: one
        required: true
    - oci:
        name: test2
        image: image2
        annotation: one
        required: true
    - local:
        name: test-local
        bucket: test-bucket2
        manifestPath: test-manifest-path2
        annotation: one
        required: true
contexts:
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
currentContext:
    kubernetes: test-mc
`
	expectedCfg2 := `cli:
    discoverySources:
        - oci:
            name: test
            image: image
            annotation: one
        - oci:
            name: test2
            image: image2
            annotation: one
            required: true
        - oci:
            name: test-local
            image: test-local-image-path
contexts:
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
currentContext:
    kubernetes: test-mc
`

	return cfg, expectedCfg, cfg2, expectedCfg2
}
func setupConfigMetadata() string {
	metadata := `configMetadata:
  patchStrategy:
    contexts.group: replace
    contexts.clusterOpts.endpoint: replace
    contexts.clusterOpts.annotation: replace
    cli.discoverySources.oci.required: replace
`
	return metadata
}

func TestIntegrationWithReplacePatchStrategy(t *testing.T) {
	// Setup config data
	cfg, expectedCfg, cfg2, expectedCfg2 := setupConfigData()
	cfgTestFiles, cleanUp := setupTestConfig(t, &CfgTestData{cfg: cfg, cfgNextGen: cfg2, cfgMetadata: setupConfigMetadata()})

	defer func() {
		cleanUp()
	}()

	// Actions

	// Get CLI discovery sources
	expectedSources := []configtypes.PluginDiscovery{
		{
			OCI: &configtypes.OCIDiscovery{
				Name:  "test",
				Image: "image",
			},
		},
		{
			OCI: &configtypes.OCIDiscovery{
				Name:  "test2",
				Image: "image2",
			},
		},
		{
			Local: &configtypes.LocalDiscovery{
				Name: "test-local",
			},
		},
	}

	sources, err := GetCLIDiscoverySources()
	assert.NoError(t, err)
	assert.Equal(t, expectedSources, sources)

	// Get CLI Discovery Source
	expectedSource := &configtypes.PluginDiscovery{
		OCI: &configtypes.OCIDiscovery{
			Name:  "test",
			Image: "image",
		},
	}

	source, err := GetCLIDiscoverySource("test")
	assert.NoError(t, err)
	assert.Equal(t, expectedSource, source)

	// Update CLI discovery sources
	updatedSources := []configtypes.PluginDiscovery{
		{
			OCI: &configtypes.OCIDiscovery{
				Name:  "test",
				Image: "image",
			},
		},
		{
			OCI: &configtypes.OCIDiscovery{
				Name:  "test-local",
				Image: "test-local-image-path",
			},
		},
	}

	err = SetCLIDiscoverySources(updatedSources)
	assert.NoError(t, err)

	// Expectations on file content
	file, err := os.ReadFile(cfgTestFiles[0].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg, string(file))

	file, err = os.ReadFile(cfgTestFiles[1].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg2, string(file))
}
