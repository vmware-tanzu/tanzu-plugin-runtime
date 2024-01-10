// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strconv"

	"github.com/pkg/errors"
)

// GetAllFeatureFlags returns all feature flags from the config file
func (c *ClientConfig) GetAllFeatureFlags() (map[string]FeatureMap, error) {
	if c.ClientOptions != nil && c.ClientOptions.Features != nil {
		return c.ClientOptions.Features, nil
	}
	return nil, errors.New("not found")
}

// IsFeatureFlagSet returns true if the features section in the configuration object contains any value for the plugin.feature combination
func (c *ClientConfig) IsFeatureFlagSet(plugin, feature string) bool {
	return c.ClientOptions != nil &&
		c.ClientOptions.Features != nil &&
		c.ClientOptions.Features[plugin] != nil &&
		c.ClientOptions.Features[plugin][feature] != ""
}

// IsConfigFeatureActivated return true if the feature is activated, false if not. An error if the featurePath is malformed
func (c *ClientConfig) IsConfigFeatureActivated(featurePath string) (bool, error) {
	plugin, flag, err := c.SplitFeaturePath(featurePath)
	if err != nil {
		return false, err
	}

	if c.ClientOptions == nil || c.ClientOptions.Features == nil ||
		c.ClientOptions.Features[plugin] == nil || c.ClientOptions.Features[plugin][flag] == "" {
		return false, nil
	}

	booleanValue, err := strconv.ParseBool(c.ClientOptions.Features[plugin][flag])
	if err != nil {
		errMsg := "error converting " + featurePath + " entry '" + c.ClientOptions.Features[plugin][flag] + "' to boolean value: " + err.Error()
		return false, errors.New(errMsg)
	}
	return booleanValue, nil
}
