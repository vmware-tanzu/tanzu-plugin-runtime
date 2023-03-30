// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"
)

// GetCEIPOptIn retrieves ClientOptions ceipOptIn
func GetCEIPOptIn() (string, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return "", err
	}
	return getCEIPOptIn(node)
}

// SetCEIPOptIn adds or updates ceipOptIn value
func SetCEIPOptIn(val string) (err error) {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Add or Update ceipOptIn in the yaml node
	persist := setCEIPOptIn(node, val)

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}
	return err
}

func setCEIPOptIn(node *yaml.Node, val string) (persist bool) {
	ceipOptInNode := getNGCLIOptionsChildNode(KeyCEIPOptIn, node)
	if ceipOptInNode != nil && ceipOptInNode.Value != val {
		ceipOptInNode.Value = val
		persist = true
	}
	return persist
}

func getCEIPOptIn(node *yaml.Node) (string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}
	if cfg != nil && cfg.CoreCliOptions != nil {
		return cfg.CoreCliOptions.CEIPOptIn, nil
	}
	return "", errors.New("ceipOptIn not found")
}

// getNGCLIOptionsChildNode parses the yaml node and returns the matched node based on configOptions
func getNGCLIOptionsChildNode(key string, node *yaml.Node) *yaml.Node {
	configOptions := func(c *nodeutils.CfgNode) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyCLI, Type: yaml.MappingNode},
			{Name: key, Type: yaml.ScalarNode, Value: ""},
		}
	}
	keyNode := nodeutils.FindNode(node.Content[0], configOptions)
	return keyNode
}
