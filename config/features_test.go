// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestIsFeatureEnabled(t *testing.T) {
	// setup
	func() {
		LocalDirName = TestLocalDirName
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tests := []struct {
		name    string
		feature map[string]configtypes.FeatureMap
		plugin  string
		key     string
	}{
		{
			name: "success context-aware-cli-for-plugins",
			feature: map[string]configtypes.FeatureMap{
				"global": {
					"context-aware-cli-for-plugins": "true",
				},
			},
			plugin: "global",
			key:    "context-aware-cli-for-plugins",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &configtypes.ClientConfig{
				ClientOptions: &configtypes.ClientOptions{
					Features: tc.feature,
				},
			}
			err := StoreClientConfig(cfg)
			assert.NoError(t, err)
			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, true)
		})
	}
}

func TestSetAndDeleteFeature(t *testing.T) {
	// setup
	func() {
		LocalDirName = TestLocalDirName
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tests := []struct {
		name    string
		feature map[string]configtypes.FeatureMap
		plugin  string
		key     string
		value   bool
		persist bool
	}{
		{
			name: "success context-aware-cli-for-plugins",
			feature: map[string]configtypes.FeatureMap{
				"global": {
					"sample":                        "true",
					"context-aware-cli-for-plugins": "true",
				},
			},
			plugin:  "global",
			key:     "context-aware-cli-for-plugins",
			value:   false,
			persist: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &configtypes.ClientConfig{
				ClientOptions: &configtypes.ClientOptions{
					Features: tc.feature,
				},
			}
			err := StoreClientConfig(cfg)
			assert.NoError(t, err)
			err = SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)
			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, tc.value)
			err = DeleteFeature(tc.plugin, tc.key)
			assert.NoError(t, err)
			ok, err = IsFeatureEnabled(tc.plugin, tc.key)
			assert.Equal(t, err.Error(), "not found")
			assert.Equal(t, ok, tc.value)
			err = SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)
		})
	}
}

func TestSetFeature(t *testing.T) {
	// setup
	func() {
		LocalDirName = TestLocalDirName
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tests := []struct {
		name   string
		cfg    *configtypes.ClientConfig
		plugin string
		key    string
		value  bool
	}{
		{
			name: "success context-aware-cli-for-plugins",
			cfg: &configtypes.ClientConfig{
				ClientOptions: &configtypes.ClientOptions{
					Features: map[string]configtypes.FeatureMap{
						"global": {
							"context-aware-cli-for-plugins": "true",
						},
					},
				},
			},
			plugin: "global",
			key:    "context-aware-cli-for-plugins",
			value:  false,
		},
		{
			name: "success context-aware-cli-for-plugins",
			cfg: &configtypes.ClientConfig{
				ClientOptions: &configtypes.ClientOptions{
					Features: map[string]configtypes.FeatureMap{
						"global": {
							"context-aware-cli-for-plugins": "true",
						},
					},
				},
			},
			plugin: "global",
			key:    "context-aware-cli-for-plugins",
			value:  false,
		},
		{
			name: "should not update the same feature value",
			cfg: &configtypes.ClientConfig{
				ClientOptions: &configtypes.ClientOptions{
					Features: map[string]configtypes.FeatureMap{
						"global": {
							"context-aware-cli-for-plugins": "true",
						},
					},
				},
			},
			plugin: "global",
			key:    "context-aware-cli-for-plugins",
			value:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := StoreClientConfig(tc.cfg)
			assert.NoError(t, err)
			err = SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)
			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, tc.value)
		})
	}
}
