// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFeatureFlags(t *testing.T) {
	testCases := []struct {
		name        string
		config      *ClientConfig
		expectErr   bool
		expectFlags map[string]FeatureMap
	}{
		{
			name: "FeatureFlagsPresent",
			config: &ClientConfig{
				ClientOptions: &ClientOptions{
					Features: map[string]FeatureMap{"plugin1": {"feature1": "true"}},
				},
			},
			expectErr:   false,
			expectFlags: map[string]FeatureMap{"plugin1": {"feature1": "true"}},
		},
		{
			name:        "FeatureFlagsAbsent",
			config:      &ClientConfig{},
			expectErr:   true,
			expectFlags: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flags, err := tc.config.GetAllFeatureFlags()
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectFlags, flags)
			}
		})
	}
}

func TestIsFeatureFlagExists(t *testing.T) {
	clientConfig := &ClientConfig{
		ClientOptions: &ClientOptions{
			Features: map[string]FeatureMap{"plugin1": {"feature1": "true"}},
		},
	}

	testCases := []struct {
		name         string
		plugin       string
		feature      string
		expectExists bool
	}{
		{
			name:         "FeatureFlagExists",
			plugin:       "plugin1",
			feature:      "feature1",
			expectExists: true,
		},
		{
			name:         "FeatureFlagDoesNotExist",
			plugin:       "plugin2",
			feature:      "feature2",
			expectExists: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			exists := clientConfig.IsFeatureFlagSet(tc.plugin, tc.feature)
			assert.Equal(t, tc.expectExists, exists)
		})
	}
}

func TestIsConfigFeatureActivated(t *testing.T) {
	clientConfig := &ClientConfig{
		ClientOptions: &ClientOptions{
			Features: map[string]FeatureMap{
				"plugin1": {"feature1": "true", "feature2": "false", "feature3": "invalid"},
			},
		},
	}

	testCases := []struct {
		name            string
		featurePath     string
		expectActivated bool
		expectErr       bool
	}{
		{
			name:            "FeatureActivated",
			featurePath:     "features.plugin1.feature1",
			expectActivated: true,
			expectErr:       false,
		},
		{
			name:            "FeatureDeactivated",
			featurePath:     "features.plugin1.feature2",
			expectActivated: false,
			expectErr:       false,
		},
		{
			name:            "InvalidFeatureValue",
			featurePath:     "features.plugin1.feature3",
			expectActivated: false,
			expectErr:       true,
		},
		{
			name:            "InvalidFeaturePath",
			featurePath:     "invalid",
			expectActivated: false,
			expectErr:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			activated, err := clientConfig.IsConfigFeatureActivated(tc.featurePath)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectActivated, activated)
			}
		})
	}
}
