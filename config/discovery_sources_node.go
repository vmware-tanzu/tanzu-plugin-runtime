// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/collectionutils"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// DiscoveryType constants
const (
	DiscoveryTypeOCI        = "oci"
	DiscoveryTypeLocal      = "local"
	DiscoveryTypeGCP        = "gcp"
	DiscoveryTypeKubernetes = "kubernetes"
	DiscoveryTypeREST       = "rest"
)

const (
	Default = "default"
)

// setDiscoverySources adds or updates the node discoverySources
func setDiscoverySources(node *yaml.Node, discoverySources []configtypes.PluginDiscovery, patchStrategyOpts ...nodeutils.PatchStrategyOpts) (persist bool, err error) {
	var anyPersists []bool
	isTrue := func(item bool) bool { return item }
	// Find the discovery sources node in the specific yaml node
	keys := []nodeutils.Key{
		{Name: KeyDiscoverySources, Type: yaml.SequenceNode},
	}
	discoverySourcesNode := nodeutils.FindNode(node, nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if discoverySourcesNode == nil {
		return persist, err
	}
	// Add or update discovery sources in the discovery sources node
	for _, discoverySource := range discoverySources {
		persist, err = setDiscoverySource(discoverySourcesNode, discoverySource, patchStrategyOpts...)
		anyPersists = append(anyPersists, persist)
		if err != nil {
			return persist, err
		}
	}
	persist = collectionutils.SomeBool(anyPersists, isTrue)
	return persist, err
}

//nolint:gocyclo
func setDiscoverySource(discoverySourcesNode *yaml.Node, discoverySource configtypes.PluginDiscovery, patchStrategyOpts ...nodeutils.PatchStrategyOpts) (persist bool, err error) {
	// Convert discoverySource change obj to yaml node
	newNode, err := convertObjectToNode(&discoverySource)
	if err != nil {
		return persist, err
	}

	exists := false
	var result []*yaml.Node

	// Get discovery source type and name
	newOrUpdatedDiscoverySourceType, newOrUpdatedDiscoverySourceName, err := getDiscoverySourceTypeAndName(discoverySource)
	if err != nil {
		return persist, err
	}

	// Loop through each discovery source node
	for _, discoverySourceNode := range discoverySourcesNode.Content {
		// Find discovery source by weak match
		discoverySourceTypeOfAnyType, discoverySourceIndexOfAnyType := findDiscoverySourceTypeAndIndexByWeakMatch(discoverySourceNode.Content)

		// Find discovery source by exact match
		discoverySourceIndexOfExactType := nodeutils.GetNodeIndex(discoverySourceNode.Content, newOrUpdatedDiscoverySourceType)

		// check if same name already exists
		nameIdx := nodeutils.GetNodeIndex(discoverySourceNode.Content[discoverySourceIndexOfAnyType].Content, "name")
		isSameNameAlreadyExists := discoverySourceNode.Content[discoverySourceIndexOfAnyType].Content[nameIdx].Value == newOrUpdatedDiscoverySourceName

		// If it's an exact match i.e. change discovery source type and current discovery source type is of same type proceed with regular merge
		if discoverySourceIndexOfAnyType != -1 && discoverySourceIndexOfExactType != -1 {
			if isSameNameAlreadyExists {
				// match found proceed with regular merge
				exists = true
				// Delete nodes as per patch strategy defined in config-metadata.yaml
				_, err = nodeutils.DeleteNodes(newNode.Content[0], discoverySourceNode, patchStrategyOpts...)
				if err != nil {
					return false, err
				}
				// Merge the new node into discovery source node
				persist, err = nodeutils.MergeNodes(newNode.Content[0], discoverySourceNode)
				if err != nil {
					return false, err
				}
			}
			// If not an exact match i.e. change discovery source type is of different current discovery type
		} else if discoverySourceIndexOfAnyType != -1 || discoverySourceIndexOfExactType != -1 {
			if isSameNameAlreadyExists {
				exists = true
				// Since merging discovery sources of different discovery source types we need to replace the nodes of different discovery type
				options := &nodeutils.PatchStrategyOptions{}
				for _, opt := range patchStrategyOpts {
					opt(options)
				}
				replaceDiscoverySourceTypeKey := fmt.Sprintf("%v.%v", options.Key, discoverySourceTypeOfAnyType)
				replaceDiscoverySourceContextTypeKey := fmt.Sprintf("%v.%v", options.Key, "contextType")
				options.PatchStrategies[replaceDiscoverySourceTypeKey] = nodeutils.PatchStrategyReplace
				options.PatchStrategies[replaceDiscoverySourceContextTypeKey] = nodeutils.PatchStrategyReplace

				// Delete nodes as per patch strategy defined in config-metadata.yaml
				_, err = nodeutils.DeleteNodes(newNode.Content[0], discoverySourceNode, patchStrategyOpts...)
				if err != nil {
					return false, err
				}
				// Merge the new node into discovery source node
				persist, err = nodeutils.MergeNodes(newNode.Content[0], discoverySourceNode)
				if err != nil {
					return false, err
				}
			}
		}
		result = append(result, discoverySourceNode)
	}
	if !exists {
		result = append(result, newNode.Content[0])
		persist = true
	}
	discoverySourcesNode.Style = 0
	discoverySourcesNode.Content = result
	return persist, err
}

func getDiscoverySourceTypeAndName(discoverySource configtypes.PluginDiscovery) (string, string, error) {
	var discoverySourceType, discoverySourceName string

	switch {
	//nolint:staticcheck // Deprecated
	case discoverySource.GCP != nil:
		discoverySourceType = DiscoveryTypeGCP
		discoverySourceName = discoverySource.GCP.Name
	case discoverySource.OCI != nil:
		discoverySourceType = DiscoveryTypeOCI
		discoverySourceName = discoverySource.OCI.Name
	case discoverySource.Local != nil:
		discoverySourceType = DiscoveryTypeLocal
		discoverySourceName = discoverySource.Local.Name
	case discoverySource.Kubernetes != nil:
		discoverySourceType = DiscoveryTypeKubernetes
		discoverySourceName = discoverySource.Kubernetes.Name
	case discoverySource.REST != nil:
		discoverySourceType = DiscoveryTypeREST
		discoverySourceName = discoverySource.REST.Name
	default:
		return "", "", errors.New("discovery source type cannot be empty")
	}

	if discoverySourceName == "" {
		return "", "", errors.New("discovery source name cannot be empty")
	}

	return discoverySourceType, discoverySourceName, nil
}

// Find the matching discovery source type and index from accepted discovery sources
func findDiscoverySourceTypeAndIndexByWeakMatch(discoverySourceContentNodes []*yaml.Node) (string, int) {
	acceptedDiscoverySources := []string{DiscoveryTypeOCI, DiscoveryTypeLocal, DiscoveryTypeGCP, DiscoveryTypeKubernetes, DiscoveryTypeREST}
	for _, discoverySourceType := range acceptedDiscoverySources {
		idx := nodeutils.GetNodeIndex(discoverySourceContentNodes, discoverySourceType)
		if idx != -1 {
			return discoverySourceType, idx
		}
	}
	return "", -1
}
