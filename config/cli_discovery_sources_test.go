// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestGetCLIDiscoverySources(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	discoveries := []configtypes.PluginDiscovery{
		{
			OCI: &configtypes.OCIDiscovery{
				Name:  "test",
				Image: "image",
			},
		},
	}

	tests := []struct {
		name    string
		in, out []configtypes.PluginDiscovery
		errStr  string
	}{
		{
			name: "success get all",
			in:   discoveries,
			out:  discoveries,
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLIDiscoverySources(spec.in)
			assert.NoError(t, err)
			c, err := GetCLIDiscoverySources()
			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)
		})
	}
}

func TestGetCLIDiscoverySource(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	discovery := &configtypes.PluginDiscovery{
		OCI: &configtypes.OCIDiscovery{
			Name:  "test",
			Image: "image",
		},
	}

	tests := []struct {
		name    string
		in, out *configtypes.PluginDiscovery
	}{
		{
			name: "success get",
			in:   discovery,
			out:  discovery,
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLIDiscoverySource(*spec.in)
			assert.NoError(t, err)
			c, err := GetCLIDiscoverySource("test")
			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)
		})
	}
}

func TestSetCLIDiscoverySources(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name  string
		input []configtypes.PluginDiscovery
		total int
	}{
		{
			name: "success add test",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test",
						Image: "image",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default",
						Path: "standalone",
					},
				},
			},
			total: 2,
		},
		{
			name: "success add test",
			input: []configtypes.PluginDiscovery{
				{
					Local: &configtypes.LocalDiscovery{
						Name: "admin-local",
						Path: "admin",
					},
				},
			},
			total: 3,
		},
		{
			name: "success add test",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "test-image",
					},
				},
			},
			total: 3,
		},
		{
			name: "success update test",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test",
						Image: "updatedImage",
					},
				},
			},
			total: 3,
		},
		{
			name: "should not persist same test",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test",
						Image: "updatedImage",
					},
				},
			},
			total: 3,
		},
		{
			name: "success add default oci",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "image",
					},
				},
			},
			total: 3,
		},
		{
			name: "success add default-local oci",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default-local",
						Image: "localImage",
					},
				},
			},
			total: 4,
		},
		{
			name: "success add default-local local",
			input: []configtypes.PluginDiscovery{
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default-local",
						Path: "test-path",
					},
				},
			},
			total: 4,
		},
		{
			name: "success add default-local local",
			input: []configtypes.PluginDiscovery{
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default-local",
						Path: "test-path",
					},
				},
			},
			total: 4,
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLIDiscoverySources(spec.input)
			assert.NoError(t, err)
			sources, err := GetCLIDiscoverySources()
			assert.NoError(t, err)
			assert.Equal(t, spec.total, len(sources))
		})
	}
}

func TestDeleteCLIDiscoverySource(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name   string
		src    []configtypes.PluginDiscovery
		input  string
		count  int
		errStr string
	}{
		{
			name: "should return err on deleting non existing source",
			src: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test",
						Image: "image",
					},
				},
			},
			input:  "test-notfound",
			count:  1,
			errStr: "cli discovery source not found",
		},
		{
			name: "should delete existing test source",
			src: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test",
						Image: "image",
					},
				},
			},
			input: "test",
			count: 0,
		},
		{
			name: "should delete test2 source",
			src: []configtypes.PluginDiscovery{
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
			},
			count: 1,
			input: "test2",
		},
		{
			name: "should delete local default source",
			src: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test",
						Image: "image",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default",
						Path: "standalone",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "admin-local",
						Path: "admin",
					},
				},
			},
			count: 2,
			input: "default",
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLIDiscoverySources(spec.src)
			assert.NoError(t, err)
			err = DeleteCLIDiscoverySource(spec.input)
			if spec.errStr != "" {
				assert.Equal(t, err.Error(), spec.errStr)
			} else {
				assert.NoError(t, err)
			}
			sources, err := GetCLIDiscoverySources()
			assert.NoError(t, err)
			assert.Equal(t, spec.count, len(sources))
		})
	}
}

func TestIntegrationSetGetDeleteCLIDiscoverySource(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	sources := []configtypes.PluginDiscovery{
		{
			OCI: &configtypes.OCIDiscovery{
				Name:  "default",
				Image: "image",
			},
		},
	}

	// Get from the empty config
	ds, err := GetCLIDiscoverySource("test")
	assert.Equal(t, "cli discovery source not found", err.Error())
	assert.Nil(t, ds)

	// Add source to empty config
	err = SetCLIDiscoverySources(sources)
	assert.NoError(t, err)

	ds, err = GetCLIDiscoverySource("default")
	assert.Nil(t, err)
	assert.Equal(t, sources[0], *ds)

	// Delete existing source
	err = DeleteCLIDiscoverySource("default")
	assert.NoError(t, err)

	ds, err = GetCLIDiscoverySource("default")
	assert.Equal(t, "cli discovery source not found", err.Error())
	assert.Nil(t, ds)

	err = DeleteCLIDiscoverySource("default-local")
	assert.Equal(t, "cli discovery source not found", err.Error())

	ds, err = GetCLIDiscoverySource("default-local")
	assert.Equal(t, "cli discovery source not found", err.Error())
	assert.Nil(t, ds)

	ds, err = GetCLIDiscoverySource("default")
	assert.Equal(t, "cli discovery source not found", err.Error())
	assert.Nil(t, ds)
}

func TestSetCLIDiscoverySourceLocalMulti(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	input := configtypes.PluginDiscovery{
		Local: &configtypes.LocalDiscovery{
			Name: "admin-local",
			Path: "admin",
		},
	}
	input2 := configtypes.PluginDiscovery{
		Local: &configtypes.LocalDiscovery{
			Name: "default-local",
			Path: "standalone",
		},
	}
	updateInput2 := configtypes.PluginDiscovery{
		Local: &configtypes.LocalDiscovery{
			Name: "default-local",
			Path: "standalone-updated",
		},
	}

	// Actions
	err := SetCLIDiscoverySource(input)
	assert.NoError(t, err)
	c, err := GetCLIDiscoverySource("admin-local")
	assert.Equal(t, input.Local, c.Local)
	assert.NoError(t, err)
	err = SetCLIDiscoverySource(input2)
	assert.NoError(t, err)
	c, err = GetCLIDiscoverySource("default-local")
	assert.Equal(t, input2.Local, c.Local)
	assert.NoError(t, err)
	// Update Input2
	err = SetCLIDiscoverySource(updateInput2)
	assert.NoError(t, err)
	c, err = GetCLIDiscoverySource("default-local")
	assert.Equal(t, updateInput2.Local, c.Local)
	assert.NoError(t, err)
}

func TestSetCLIDiscoverySourceWithDefaultAndDefaultLocal(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name         string
		input        []configtypes.PluginDiscovery
		totalSources int
	}{
		{
			name: "success add default-test source",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default-test",
						Image: "image",
					},
				},
			},
			totalSources: 1,
		},
		{
			name: "success add default source",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "image",
					},
				},
			},
			totalSources: 2,
		},
		{
			name: "success add default-local source",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default-local",
						Image: "image",
					},
				},
			},

			totalSources: 3,
		},
		{
			name: "success update default",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "updatedImage",
					},
				},
			},

			totalSources: 3,
		},
		{
			name: "success update default-local",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default-local",
						Image: "updatedImage",
					},
				},
			},

			totalSources: 3,
		},
		{
			name: "success add default",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "updatedImage",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "test-local",
						Path: "test-local-path",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default",
						Path: "default-local-path",
					},
				},
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "updatedImage2",
					},
				},
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test-oci1",
						Image: "updatedImage",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default-local",
						Path: "default-local-path",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "test-oci1",
						Path: "default-local-path",
					},
				},
			},
			totalSources: 5,
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			for _, ds := range spec.input {
				err := SetCLIDiscoverySource(ds)
				assert.NoError(t, err)
			}

			if spec.totalSources != 0 {
				sources, err := GetCLIDiscoverySources()
				assert.NoError(t, err)
				assert.Equal(t, spec.totalSources, len(sources))
			}
		})
	}
}

func TestSetCLIDiscoverySourceMultiTypes(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name         string
		input        []configtypes.PluginDiscovery
		totalSources int
	}{

		{
			name: "success add multiple discovery source types",
			input: []configtypes.PluginDiscovery{
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "defaultImage",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "test-local",
						Path: "test-local-path",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default",
						Path: "default-local-path",
					},
				},
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "default",
						Image: "defaultImage2",
					},
				},
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test-oci1",
						Image: "updatedImage",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "default-local",
						Path: "default-local-path",
					},
				},
				{
					Local: &configtypes.LocalDiscovery{
						Name: "test-oci1",
						Path: "default-local-path",
					},
				},
				{
					OCI: &configtypes.OCIDiscovery{
						Name:  "test-oci2",
						Image: "updatedImage",
					},
				},
			},
			totalSources: 5,
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			for _, ds := range spec.input {
				err := SetCLIDiscoverySource(ds)
				assert.NoError(t, err)
			}

			if spec.totalSources != 0 {
				sources, err := GetCLIDiscoverySources()
				assert.NoError(t, err)
				assert.Equal(t, spec.totalSources, len(sources))
			}
		})
	}
}
