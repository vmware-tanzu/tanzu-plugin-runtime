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

// GetServer retrieves server by name
//
// Deprecated: This API is deprecated. Use GetContext instead.
func GetServer(name string) (*configtypes.Server, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getServer(node, name)
}

// ServerExists checks if server by specified name is present in config
//
// Deprecated: This API is deprecated. Use ContextExists instead.
func ServerExists(name string) (bool, error) {
	exists, _ := GetServer(name)
	return exists != nil, nil
}

// GetCurrentServer retrieves the current server
//
// Deprecated: This API is deprecated. Use GetCurrentContext instead.
func GetCurrentServer() (*configtypes.Server, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getCurrentServer(node)
}

// SetCurrentServer add or update current server
//
// Deprecated: This API is deprecated. Use SetCurrentContext instead.
func SetCurrentServer(name string) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	s, err := getServer(node, name)
	if err != nil {
		return err
	}
	persist, err := setCurrentServer(node, name)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}
	// Front fill CurrentContext
	c := convertServerToContext(s)
	persist, err = setCurrentContext(node, c.Name, c.ContextType)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveCurrentServer removes the current server if server exists by specified name
//
// Deprecated: This API is deprecated. Use RemoveCurrentContext instead.
func RemoveCurrentServer(name string) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	_, err = getServer(node, name)
	if err != nil {
		return err
	}
	err = removeCurrentServer(node, name)
	if err != nil {
		return err
	}

	// Front fill Context and CurrentContext
	c, err := getContext(node, name)
	if err != nil {
		return err
	}
	err = removeCurrentContext(node, c.Name, c.ContextType)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

// PutServer add or update server and currentServer
//
// Deprecated: This API is deprecated. Use AddContext or SetContext instead.
func PutServer(s *configtypes.Server, setCurrent bool) error {
	return SetServer(s, setCurrent)
}

// AddServer add or update server and currentServer
//
// Deprecated: This API is deprecated. Use AddContext or SetContext instead.
func AddServer(s *configtypes.Server, setCurrent bool) error {
	return SetServer(s, setCurrent)
}

// SetServer add or update server and currentServer
//
// Deprecated: This API is deprecated. Use AddContext or SetContext instead.
func SetServer(s *configtypes.Server, setCurrent bool) error {
	// Acquire tanzu config lock
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	persist, err := setServer(node, s)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}
	if setCurrent && s.Type == configtypes.ManagementClusterServerType {
		persist, err = setCurrentServer(node, s.Name)
		if err != nil {
			return err
		}
		if persist {
			err = persistConfig(node)
			if err != nil {
				return err
			}
		}
	}

	err = frontFillContexts(s, setCurrent, node)
	if err != nil {
		return err
	}

	return nil
}

func frontFillContexts(s *configtypes.Server, setCurrent bool, node *yaml.Node) error {
	// Front fill Context and CurrentContext
	c := convertServerToContext(s)
	persist, err := setContext(node, c)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}
	if setCurrent {
		persist, err = setCurrentContext(node, c.Name, c.ContextType)
		if err != nil {
			return err
		}
		if persist {
			err = persistConfig(node)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteServer deletes the server specified by name
//
// Deprecated: This API is deprecated. Use DeleteContext instead.
func DeleteServer(name string) error {
	return RemoveServer(name)
}

// RemoveServer removed the server by name
//
// Deprecated: This API is deprecated. Use DeleteContext instead.
func RemoveServer(name string) error {
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	_, err = getServer(node, name)
	if err != nil {
		return err
	}
	err = removeCurrentServer(node, name)
	if err != nil {
		return err
	}
	err = removeServer(node, name)
	if err != nil {
		return err
	}
	// Front fill Context and CurrentContext
	c, err := getContext(node, name)
	if err != nil {
		return err
	}
	err = removeCurrentContext(node, c.Name, c.ContextType)
	if err != nil {
		return err
	}
	err = removeContext(node, name)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

func setCurrentServer(node *yaml.Node, name string) (persist bool, err error) {
	s, err := getServer(node, name)
	if err != nil {
		return false, err
	}

	if s.Type != configtypes.ManagementClusterServerType {
		return false, errors.Errorf("cannot set non management-cluster server as current server")
	}

	// find current server node
	keys := []nodeutils.Key{
		{Name: KeyCurrentServer, Type: yaml.ScalarNode, Value: ""},
	}
	currentServerNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if currentServerNode == nil {
		return persist, nodeutils.ErrNodeNotFound
	}
	if currentServerNode.Value != name {
		currentServerNode.Value = name
		persist = true
	}
	return persist, err
}

func getServer(node *yaml.Node, name string) (*configtypes.Server, error) {
	// check if name is empty
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	for _, server := range cfg.KnownServers {
		if server.Name == name {
			return server, nil
		}
	}
	return nil, fmt.Errorf("could not find server %q", name)
}

func getCurrentServer(node *yaml.Node) (s *configtypes.Server, err error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	for _, server := range cfg.KnownServers {
		if server.Name == cfg.CurrentServer {
			return server, nil
		}
	}
	return s, fmt.Errorf("current server %q not found in tanzu config", cfg.CurrentServer)
}

func removeCurrentServer(node *yaml.Node, name string) error {
	// check if name is empty
	if name == "" {
		return errors.New("name cannot be empty")
	}

	// find current server node
	keys := []nodeutils.Key{
		{Name: KeyCurrentServer},
	}
	currentServerNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if currentServerNode == nil {
		return nil
	}
	if currentServerNode.Value == name {
		currentServerNode.Value = ""
	}
	return nil
}

//nolint:dupl
func removeServer(node *yaml.Node, name string) error {
	// check if name is empty
	if name == "" {
		return errors.New("name cannot be empty")
	}

	// find servers node
	keys := []nodeutils.Key{
		{Name: KeyServers},
	}
	serversNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if serversNode == nil {
		return nil
	}
	var servers []*yaml.Node
	for _, serverNode := range serversNode.Content {
		if index := nodeutils.GetNodeIndex(serverNode.Content, "name"); index != -1 && serverNode.Content[index].Value == name {
			continue
		}
		servers = append(servers, serverNode)
	}
	serversNode.Content = servers
	return nil
}

func setServers(node *yaml.Node, servers []*configtypes.Server) error {
	for _, server := range servers {
		_, err := setServer(node, server)
		if err != nil {
			return err
		}
	}
	return nil
}

func setServer(node *yaml.Node, s *configtypes.Server) (persist bool, err error) {
	// check if name is empty
	if s.Name == "" {
		return false, errors.New("server name cannot be empty")
	}

	// Get Patch Strategies
	patchStrategies, err := GetConfigMetadataPatchStrategy()
	if err != nil {
		patchStrategies = make(map[string]string)
	}
	var persistDiscoverySources bool

	// convert server to node
	newServerNode, err := convertObjectToNode(s)
	if err != nil {
		return persist, err
	}

	// find servers node
	keys := []nodeutils.Key{
		{Name: KeyServers, Type: yaml.SequenceNode},
	}
	serversNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if serversNode == nil {
		return persist, nodeutils.ErrNodeNotFound
	}
	exists := false
	var result []*yaml.Node
	//nolint: dupl
	for _, serverNode := range serversNode.Content {
		if index := nodeutils.GetNodeIndex(serverNode.Content, "name"); index != -1 &&
			serverNode.Content[index].Value == s.Name {
			exists = true
			_, err = nodeutils.DeleteNodes(newServerNode.Content[0], serverNode, nodeutils.WithPatchStrategyKey(KeyServers), nodeutils.WithPatchStrategies(patchStrategies))
			if err != nil {
				return false, err
			}
			persist, err = nodeutils.MergeNodes(newServerNode.Content[0], serverNode)
			if err != nil {
				return false, err
			}
			// add or update discovery sources of server
			persistDiscoverySources, err = setDiscoverySources(serverNode, s.DiscoverySources, nodeutils.WithPatchStrategyKey(fmt.Sprintf("%v.%v", KeyServers, KeyDiscoverySources)), nodeutils.WithPatchStrategies(patchStrategies))
			if err != nil {
				return false, err
			}
			if persistDiscoverySources {
				_, err = nodeutils.MergeNodes(newServerNode.Content[0], serverNode)
				if err != nil {
					return false, err
				}
			}
			result = append(result, serverNode)
			continue
		}
		result = append(result, serverNode)
	}
	if !exists {
		result = append(result, newServerNode.Content[0])
		persist = true
	}
	serversNode.Content = result
	return persistDiscoverySources || persist, err
}

// EndpointFromServer returns the endpoint from server.
//
// Deprecated: This API is deprecated. Use EndpointFromContext instead.
func EndpointFromServer(s *configtypes.Server) (endpoint string, err error) {
	switch s.Type {
	case configtypes.ManagementClusterServerType:
		return s.ManagementClusterOpts.Endpoint, nil
	case configtypes.GlobalServerType:
		return s.GlobalOpts.Endpoint, nil
	default:
		return endpoint, fmt.Errorf("unknown server type %q", s.Type)
	}
}
