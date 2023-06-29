// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"errors"

	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// GetCLITelemetryOptions retrieves the CLI telemetry configuration
func GetCLITelemetryOptions() (*configtypes.TelemetryOptions, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getCLITelemetryOptions(node)
}

// SetCLITelemetryOptions add or update CLI telemetry configuration
func SetCLITelemetryOptions(c *configtypes.TelemetryOptions) error {
	if c == nil {
		return nil
	}
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// Add or update the CLI telemetry options
	persist, err := setCLITelemetryOptions(node, c)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}
	return err
}

// DeleteTelemetryOptions deletes the telemetry options  from the CLI configuration
func DeleteTelemetryOptions() error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	err = deleteTelemetryOptionsNode(node)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

// Pre-reqs: node != nil
func getCLITelemetryOptions(node *yaml.Node) (*configtypes.TelemetryOptions, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	if cfg != nil && cfg.CoreCliOptions != nil && cfg.CoreCliOptions.TelemetryOptions != nil {
		return cfg.CoreCliOptions.TelemetryOptions, nil
	}
	return nil, errors.New("telemetry not found")
}

// Pre-reqs: node != nil and telemetryOptions != nil
func setCLITelemetryOptions(node *yaml.Node, telemetryOptions *configtypes.TelemetryOptions) (persist bool, err error) {
	// Get Patch Strategies from config metadata
	patchStrategies, err := GetConfigMetadataPatchStrategy()
	if err != nil {
		patchStrategies = make(map[string]string)
	}

	// Convert telemetryOptions to node
	newTelemetryNode, err := convertObjectToNode(telemetryOptions)
	if err != nil {
		return persist, err
	}

	// Find the telemetry node from the root node
	keys := []nodeutils.Key{
		{Name: KeyCLI, Type: yaml.MappingNode},
		{Name: KeyTelemetry, Type: yaml.MappingNode},
	}
	telemetryOptionsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if telemetryOptionsNode == nil {
		return persist, err
	}

	// replace the nodes as per patch strategy
	_, err = nodeutils.DeleteNodes(newTelemetryNode.Content[0], telemetryOptionsNode, nodeutils.WithPatchStrategyKey(KeyTelemetry), nodeutils.WithPatchStrategies(patchStrategies))
	if err != nil {
		return false, err
	}
	persist, err = nodeutils.MergeNodes(newTelemetryNode.Content[0], telemetryOptionsNode)
	if err != nil {
		return false, err
	}
	return persist, err
}

// deleteTelemetryOptionsNode removes the telemetry options in the configuration
// Pre-reqs: node != nil
func deleteTelemetryOptionsNode(node *yaml.Node) error {
	// Find the telemetry node from the root node
	keys := []nodeutils.Key{
		{Name: KeyCLI, Type: yaml.MappingNode},
	}
	cliOptionsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if cliOptionsNode == nil {
		return nil
	}
	targetNodeIndex := nodeutils.GetNodeIndex(cliOptionsNode.Content, KeyTelemetry)
	if targetNodeIndex == -1 {
		return nil
	}
	targetNodeIndex--
	cliOptionsNode.Content = append(cliOptionsNode.Content[:targetNodeIndex], cliOptionsNode.Content[targetNodeIndex+2:]...)

	return nil
}
