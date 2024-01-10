// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"

	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"
)

// GetAllFeatureFlags retrieves all feature flags values from config
func GetAllFeatureFlags() (map[string]types.FeatureMap, error) {
	// Retrieve client config node
	cfg, err := GetClientConfig()
	if err != nil {
		return nil, err
	}
	return cfg.GetAllFeatureFlags()
}

// IsFeatureEnabled checks and returns whether specific plugin and key is true
func IsFeatureEnabled(plugin, key string) (bool, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return false, err
	}
	val, err := getFeature(node, plugin, key)
	if err != nil {
		return false, err
	}
	if strings.EqualFold(val, "true") {
		return true, nil
	}
	return false, nil
}

func getFeature(node *yaml.Node, plugin, key string) (string, error) {
	// check if plugin is empty
	if plugin == "" {
		return "", errors.New("plugin cannot be empty")
	}

	// check if key is empty
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}
	if cfg.ClientOptions == nil || cfg.ClientOptions.Features == nil || cfg.ClientOptions.Features[plugin] == nil {
		return "", errors.New("not found")
	}
	if val, ok := cfg.ClientOptions.Features[plugin][key]; ok {
		return val, nil
	}
	return "", errors.New("not found")
}

// DeleteFeature deletes the specified plugin key
func DeleteFeature(plugin, key string) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	err = deleteFeature(node, plugin, key)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

func deleteFeature(node *yaml.Node, plugin, key string) error {
	// check if plugin is empty
	if plugin == "" {
		return errors.New("plugin cannot be empty")
	}

	// check if key is empty
	if key == "" {
		return errors.New("key cannot be empty")
	}
	// Find plugin node
	keys := []nodeutils.Key{
		{Name: KeyClientOptions},
		{Name: KeyFeatures},
		{Name: plugin},
	}
	pluginNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if pluginNode == nil {
		return nil
	}
	plugins, err := nodeutils.ConvertNodeToMap(pluginNode)
	if err != nil {
		return err
	}
	delete(plugins, key)
	newPluginsNode, err := nodeutils.ConvertMapToNode(plugins)
	if err != nil {
		return err
	}
	pluginNode.Content = newPluginsNode.Content[0].Content
	return nil
}

// SetFeature add or update plugin key value
func SetFeature(plugin, key, value string) (err error) {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// Add or Update Feature plugin
	persist, err := setFeature(node, plugin, key, value)
	if err != nil {
		return err
	}
	if persist {
		return persistConfig(node)
	}
	return err
}

func setFeature(node *yaml.Node, plugin, key, value string) (persist bool, err error) {
	// check if plugin is empty
	if plugin == "" {
		return false, errors.New("plugin cannot be empty")
	}

	// check if key is empty
	if key == "" {
		return false, errors.New("key cannot be empty")
	}

	// check if value is empty
	if value == "" {
		return false, errors.New("value cannot be empty")
	}

	// find plugin node
	keys := []nodeutils.Key{
		{Name: KeyClientOptions, Type: yaml.MappingNode},
		{Name: KeyFeatures, Type: yaml.MappingNode},
		{Name: plugin, Type: yaml.MappingNode},
	}
	pluginNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if pluginNode == nil {
		return persist, nodeutils.ErrNodeNotFound
	}
	if index := nodeutils.GetNodeIndex(pluginNode.Content, key); index != -1 {
		if pluginNode.Content[index].Value != value {
			pluginNode.Content[index].Tag = "!!str"
			pluginNode.Content[index].Value = value
			persist = true
		}
	} else {
		pluginNode.Content = append(pluginNode.Content, nodeutils.CreateScalarNode(key, value)...)
		persist = true
	}
	return persist, err
}

// ConfigureDefaultFeatureFlagsIfMissing add or update plugin features based on specified default feature flags
func ConfigureDefaultFeatureFlagsIfMissing(plugin string, defaultFeatureFlags map[string]bool) error {
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// find plugin node
	keys := []nodeutils.Key{
		{Name: KeyClientOptions, Type: yaml.MappingNode},
		{Name: KeyFeatures, Type: yaml.MappingNode},
		{Name: plugin, Type: yaml.MappingNode},
	}
	pluginNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if pluginNode == nil {
		return nodeutils.ErrNodeNotFound
	}
	for key, value := range defaultFeatureFlags {
		val := strconv.FormatBool(value)
		if index := nodeutils.GetNodeIndex(pluginNode.Content, key); index != -1 {
			pluginNode.Content[index].Value = val
		} else {
			pluginNode.Content = append(pluginNode.Content, nodeutils.CreateScalarNode(key, val)...)
		}
	}
	return nil
}

// IsFeatureActivated returns true if the given feature is activated
// User can set this CLI feature flag using `tanzu config set features.global.<feature> true`
func IsFeatureActivated(feature string) bool {
	cfg, err := GetClientConfig()
	if err != nil {
		return false
	}
	status, err := cfg.IsConfigFeatureActivated(feature)
	if err != nil {
		return false
	}
	return status
}

// FeatureOptions is a struct that defines the options for feature flag configuration.
type FeatureOptions struct {
	SkipIfExists bool // SkipIfExists indicates whether to skip setting the feature flag if it already exists.
}

// Options is a function type that applies configuration options to FeatureOptions.
type Options func(opts *FeatureOptions)

// SkipIfExists returns an Options function that sets the SkipIfExists option to true.
func SkipIfExists() Options {
	return func(opts *FeatureOptions) {
		opts.SkipIfExists = true // Sets the SkipIfExists field of FeatureOptions to true.
	}
}

// ConfigureFeatureFlags sets default feature flags to ClientConfig if they are missing.
// It accepts a map of feature flags and a variadic Options parameter to apply additional settings.
func ConfigureFeatureFlags(defaultFeatureFlags map[string]bool, opts ...Options) error {
	options := new(FeatureOptions) // Initialize FeatureOptions.
	for _, opt := range opts {
		opt(options) // Apply each Options function to the FeatureOptions.
	}

	cfg, err := GetClientConfig() // Retrieves the current client configuration.
	if err != nil {
		return errors.Wrap(err, "error while getting client config") // Returns an error if fetching client config fails.
	}

	return configureFlags(cfg, defaultFeatureFlags, options) // Sets the feature flags based on the provided configuration and options.
}

// configureFlags processes and sets the feature flags based on the provided configuration and options.
// It iterates over each flag in the provided map and sets them using the setFlag function.
func configureFlags(cfg *types.ClientConfig, flags map[string]bool, options *FeatureOptions) error {
	for path, activated := range flags {
		if err := setFlag(cfg, path, activated, options); err != nil {
			return err // Returns an error if setting any feature flag fails.
		}
	}
	return nil
}

// setFlag configures a single feature flag based on the specified path and activation status.
// It extracts the plugin and feature from the path and sets the feature flag accordingly.
func setFlag(cfg *types.ClientConfig, path string, activated bool, options *FeatureOptions) error {
	plugin, feature, err := cfg.SplitFeaturePath(path)
	if err != nil {
		return errors.Wrap(err, "failed to configure feature flags") // Returns an error if splitting the feature path fails.
	}

	// Checks if the feature flag should be skipped if it already exists.
	if options.SkipIfExists && cfg.IsFeatureFlagSet(plugin, feature) {
		return nil // Skips setting the flag.
	}

	// Sets the feature flag with the specified activation status.
	return SetFeature(plugin, feature, strconv.FormatBool(activated))
}
