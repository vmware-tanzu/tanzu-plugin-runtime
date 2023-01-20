// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestSetGetRepository(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name string
		cfg  *configtypes.ClientConfig
		in   configtypes.PluginRepository
		out  configtypes.PluginRepository
	}{
		{
			name: "should persist repository",
			cfg:  &configtypes.ClientConfig{},
			in: configtypes.PluginRepository{
				GCPPluginRepository: &configtypes.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
			out: configtypes.PluginRepository{
				GCPPluginRepository: &configtypes.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
		},
		{
			name: "should not persist same repo",
			cfg: &configtypes.ClientConfig{
				ClientOptions: &configtypes.ClientOptions{
					CLI: &configtypes.CLIOptions{
						Repositories: []configtypes.PluginRepository{
							{
								GCPPluginRepository: &configtypes.GCPPluginRepository{
									Name:       "test",
									BucketName: "bucket",
									RootPath:   "root-path",
								},
							},
						},
					},
				},
			},
			in: configtypes.PluginRepository{
				GCPPluginRepository: &configtypes.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
			out: configtypes.PluginRepository{
				GCPPluginRepository: &configtypes.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := StoreClientConfig(tc.cfg)
			assert.NoError(t, err)
			err = SetCLIRepository(tc.in)
			assert.NoError(t, err)
			r, err := GetCLIRepository(tc.out.GCPPluginRepository.Name)
			assert.NoError(t, err)
			assert.Equal(t, tc.out.GCPPluginRepository.Name, r.GCPPluginRepository.Name)
		})
	}
}
