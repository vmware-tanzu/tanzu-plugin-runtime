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
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name                string
		plugin              string
		key                 string
		value               string
		errStr              string
		errStrForSetFeature string
	}{
		{
			name:                "empty plugin",
			plugin:              "",
			key:                 "feature1",
			value:               "true",
			errStr:              "plugin cannot be empty",
			errStrForSetFeature: "plugin cannot be empty",
		},
		{
			name:                "empty key",
			plugin:              "global",
			key:                 "",
			value:               "",
			errStr:              "key cannot be empty",
			errStrForSetFeature: "key cannot be empty",
		},
		{
			name:                "empty value",
			plugin:              "global",
			key:                 "feature1",
			value:               "",
			errStr:              "not found",
			errStrForSetFeature: "value cannot be empty",
		},
		{
			name:   "success feature1",
			plugin: "global",
			key:    "feature1",
			value:  "true",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SetFeature(tc.plugin, tc.key, tc.value)
			if tc.errStrForSetFeature != "" {
				assert.Equal(t, tc.errStrForSetFeature, err.Error())
			} else {
				assert.NoError(t, err)
			}

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			if tc.errStr != "" {
				assert.Equal(t, tc.errStr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, ok, true)
			}
		})
	}
}

func TestSetAndDeleteFeature(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name   string
		plugin string
		key    string
		value  bool
	}{
		{
			name:   "success feature1",
			plugin: "global",
			key:    "feature1",
			value:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, tc.value)

			err = DeleteFeature(tc.plugin, tc.key)
			assert.NoError(t, err)

			ok, err = IsFeatureEnabled(tc.plugin, tc.key)
			assert.Equal(t, err.Error(), "not found")
			assert.Equal(t, ok, tc.value)
		})
	}
}

func TestSetFeature(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name         string
		plugin       string
		key          string
		value        bool
		errStrForSet string
		errStrForGet string
	}{
		{
			name:         "empty plugin",
			plugin:       "",
			key:          "feature1",
			value:        false,
			errStrForSet: "plugin cannot be empty",
			errStrForGet: "plugin cannot be empty",
		},
		{
			name:         "empty key",
			plugin:       "global",
			key:          "",
			value:        false,
			errStrForSet: "key cannot be empty",
			errStrForGet: "key cannot be empty",
		},
		{
			name:   "success feature1",
			plugin: "global",
			key:    "feature1",
			value:  false,
		},
		{
			name:   "success feature1",
			plugin: "global",
			key:    "feature1",
			value:  false,
		},
		{
			name:   "should not update the same feature value",
			plugin: "global",
			key:    "feature1",
			value:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			if tc.errStrForSet != "" {
				assert.Equal(t, tc.errStrForSet, err.Error())
			} else {
				assert.NoError(t, err)
			}

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			if tc.errStrForGet != "" {
				assert.Equal(t, tc.errStrForGet, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, ok, tc.value)
			}
		})
	}
}

func TestIsFeatureActivated(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	defaultFlags := map[string]bool{
		"features.plugin1.feature1": true,
		"features.plugin2.feature2": false,
		"features.plugin3.feature3": false,
	}

	testCases := []struct {
		name            string
		feature         string
		expectActivated bool
	}{
		{
			name:            "FeatureActivated",
			feature:         "features.plugin1.feature1",
			expectActivated: true,
		},
		{
			name:            "FeatureDeactivated",
			feature:         "features.plugin2.feature2",
			expectActivated: false,
		},
		{
			name:            "ConfigRetrievalFails",
			feature:         "features.plugin3.feature3",
			expectActivated: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ConfigureFeatureFlags(defaultFlags)
			assert.NoError(t, err)

			activated := IsFeatureActivated(tc.feature)
			assert.Equal(t, tc.expectActivated, activated)
		})
	}
}

func TestConfigureDefaultFeatureFlags(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	testCases := []struct {
		name          string
		defaultPlugin string
		defaultKey    string
		defaultValue  string
		flags         map[string]bool
		opts          []Options
		resultFlags   map[string]configtypes.FeatureMap
		expectErr     bool
	}{
		{
			name:          "SuccessfulConfiguration of global flags",
			defaultPlugin: "global",
			defaultKey:    "feature1",
			defaultValue:  "false",
			flags: map[string]bool{
				"features.global.feature1": true,
				"features.global.feature2": false,
				"features.global.feature3": false,
			},
			opts: nil,
			resultFlags: map[string]configtypes.FeatureMap{
				"global": {
					"feature1": "true",
					"feature2": "false",
					"feature3": "false",
				},
			},
			expectErr: false,
		},
		{
			name:          "SuccessfulConfiguration of global flags and skip is exists",
			defaultPlugin: "global",
			defaultKey:    "feature1",
			defaultValue:  "false",
			flags: map[string]bool{
				"features.global.feature1": true,
				"features.global.feature2": false,
				"features.global.feature3": false,
			},
			opts: []Options{SkipIfExists()},
			resultFlags: map[string]configtypes.FeatureMap{
				"global": {
					"feature1": "false",
					"feature2": "false",
					"feature3": "false",
				},
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set default feature flag
			_ = SetFeature(tc.defaultPlugin, tc.defaultKey, tc.defaultValue)

			// Configure test feature flags
			err := ConfigureFeatureFlags(tc.flags, tc.opts...)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				flags, err := GetAllFeatureFlags()
				assert.NoError(t, err)
				assert.Equal(t, tc.resultFlags, flags)
			}
		})
	}
}

func TestGetAllFeatureFlags(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	testCases := []struct {
		name          string
		plugin        string
		key           string
		value         string
		expectedFlags map[string]configtypes.FeatureMap
	}{
		{
			name:   "FeatureActivated",
			plugin: "plugin",
			key:    "feature1",
			value:  "true",

			expectedFlags: map[string]configtypes.FeatureMap{
				"plugin": {
					"feature1": "true",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetFeature(tc.plugin, tc.key, tc.value)
			assert.NoError(t, err)

			flags, err := GetAllFeatureFlags()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedFlags, flags)
		})
	}
}
