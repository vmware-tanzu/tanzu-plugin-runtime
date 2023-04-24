// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// GetCLIDiscoverySources retrieves cli discovery sources
func GetCLIDiscoverySources() ([]configtypes.PluginDiscovery, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIDiscoverySources(node)
}

// GetCLIDiscoverySource retrieves cli discovery source by name assuming that there should only be one source with the name, returns the first match
func GetCLIDiscoverySource(name string) (*configtypes.PluginDiscovery, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIDiscoverySource(node, name)
}

// SetCLIDiscoverySources Add/Update array of cli discovery sources to the yaml node
func SetCLIDiscoverySources(discoverySources []configtypes.PluginDiscovery) (err error) {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Loop through each discovery source and add or update existing node
	for _, discoverySource := range discoverySources {
		persist, err := setCLIDiscoverySource(node, discoverySource)
		if err != nil {
			return err
		}
		// Persist the config node to the file
		if persist {
			err = persistConfig(node)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// SetCLIDiscoverySource add or update a cli discoverySource
func SetCLIDiscoverySource(discoverySource configtypes.PluginDiscovery) (err error) {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Add/Update cli discovery source in the yaml node
	persist, err := setCLIDiscoverySource(node, discoverySource)
	if err != nil {
		return err
	}

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}

	return err
}

// DeleteCLIDiscoverySource delete cli discoverySource by name
func DeleteCLIDiscoverySource(name string) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Delete the matching cli discovery source from the yaml node
	err = deleteCLIDiscoverySource(node, name)
	if err != nil {
		return err
	}

	// Persist the config node to the file
	return persistConfig(node)
}

func getCLIDiscoverySources(node *yaml.Node) ([]configtypes.PluginDiscovery, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	if cfg.CoreCliOptions != nil && cfg.CoreCliOptions.DiscoverySources != nil {
		return cfg.CoreCliOptions.DiscoverySources, nil
	}
	return nil, errors.New("cli discovery sources not found")
}

func getCLIDiscoverySource(node *yaml.Node, name string) (*configtypes.PluginDiscovery, error) {
	// check if context name is empty
	if name == "" {
		return nil, errors.New("discovery source name cannot be empty")
	}

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	if cfg.CoreCliOptions != nil && cfg.CoreCliOptions.DiscoverySources != nil {
		for _, discoverySource := range cfg.CoreCliOptions.DiscoverySources {
			_, discoverySourceName, err := getDiscoverySourceTypeAndName(discoverySource)
			if err != nil {
				return nil, err
			}
			if discoverySourceName == name {
				return &discoverySource, nil
			}
		}
	}
	return nil, errors.New("cli discovery source not found")
}

// setCLIDiscoverySource Add/Update cli discovery source in the yaml node
func setCLIDiscoverySource(node *yaml.Node, discoverySource configtypes.PluginDiscovery) (persist bool, err error) {
	// Retrieve the patch strategies from config metadata
	patchStrategies, err := GetConfigMetadataPatchStrategy()
	if err != nil {
		patchStrategies = make(map[string]string)
	}

	// Find the cli discovery sources node
	keys := []nodeutils.Key{
		{Name: KeyCLI, Type: yaml.MappingNode},
		{Name: KeyDiscoverySources, Type: yaml.SequenceNode},
	}
	discoverySourcesNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if discoverySourcesNode == nil {
		return persist, nodeutils.ErrNodeNotFound
	}

	// Add or Update cli discovery source to discovery sources node based on patch strategy
	key := fmt.Sprintf("%v.%v", KeyCLI, KeyDiscoverySources)
	return setDiscoverySource(discoverySourcesNode, discoverySource, nodeutils.WithPatchStrategyKey(key), nodeutils.WithPatchStrategies(patchStrategies))
}

func deleteCLIDiscoverySource(node *yaml.Node, name string) error {
	// Find cli discovery sources node in the yaml node
	keys := []nodeutils.Key{
		{Name: KeyCLI},
		{Name: KeyDiscoverySources},
	}
	cliDiscoverySourcesNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if cliDiscoverySourcesNode == nil {
		return nil
	}

	// Get matching cli discovery source from the yaml node
	discoverySource, err := getCLIDiscoverySource(node, name)
	if err != nil {
		return err
	}

	discoverySourceType, discoverySourceName, err := getDiscoverySourceTypeAndName(*discoverySource)
	if err != nil {
		return err
	}

	var result []*yaml.Node
	for _, discoverySourceNode := range cliDiscoverySourcesNode.Content {
		// Find discovery source matched by discoverySourceType
		if discoverySourceIndex := nodeutils.GetNodeIndex(discoverySourceNode.Content, discoverySourceType); discoverySourceIndex != -1 {
			// Find matching discovery source
			if discoverySourceFieldIndex := nodeutils.GetNodeIndex(discoverySourceNode.Content[discoverySourceIndex].Content, "name"); discoverySourceFieldIndex != -1 && discoverySourceNode.Content[discoverySourceIndex].Content[discoverySourceFieldIndex].Value == discoverySourceName {
				continue
			}
		}
		result = append(result, discoverySourceNode)
	}
	cliDiscoverySourcesNode.Style = 0
	cliDiscoverySourcesNode.Content = result
	return nil
}
