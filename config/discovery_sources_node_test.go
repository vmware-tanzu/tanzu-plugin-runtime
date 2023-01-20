// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestSetDiscoverySource(t *testing.T) {
	func() {
		LocalDirName = TestLocalDirName
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tests := []struct {
		name            string
		discoverySource configtypes.PluginDiscovery
		contextNode     *yaml.Node
		errStr          string
	}{
		{
			name: "success k8s",
			discoverySource: configtypes.PluginDiscovery{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "updated-test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},

			contextNode: &yaml.Node{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := setDiscoverySource(tc.contextNode, tc.discoverySource, nil)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
		})
	}
}
