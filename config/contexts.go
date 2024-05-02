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

// GetContext retrieves the context by name
func GetContext(name string) (*configtypes.Context, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getContext(node, name)
}

// AddContext add or update context and currentContext
func AddContext(c *configtypes.Context, setCurrent bool) error {
	return SetContext(c, setCurrent)
}

// SetContext add or update context and currentContext
//
//nolint:gocyclo
func SetContext(c *configtypes.Context, setCurrent bool) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// Add or update the context
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
	// Set current context
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

	// Back-fill servers based on contexts
	if c.ContextType == configtypes.ContextTypeTanzu {
		return nil
	}
	s := convertContextToServer(c)

	// Add or update server
	persist, err = setServer(node, s)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}

	// Set current server
	if setCurrent && s.Type == configtypes.ManagementClusterServerType { //nolint:staticcheck
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
	return err
}

// DeleteContext delete a context by name
func DeleteContext(name string) error {
	return RemoveContext(name)
}

// RemoveContext delete a context by name
func RemoveContext(name string) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	ctx, err := getContext(node, name)
	if err != nil {
		return err
	}
	err = removeCurrentContext(node, ctx.Name, ctx.ContextType)
	if err != nil {
		return err
	}
	err = removeContext(node, name)
	if err != nil {
		return err
	}
	err = removeServer(node, name)
	if err != nil {
		return err
	}
	err = removeCurrentServer(node, name)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

// ContextExists checks if context by name already exists
func ContextExists(name string) (bool, error) {
	exists, _ := GetContext(name)
	return exists != nil, nil
}

func validateContext(c *configtypes.Context) error {
	// Check if ctx.Name is empty
	if c.Name == "" {
		return errors.New("context name cannot be empty")
	}
	if c.Target != "" && c.ContextType != "" && c.ContextType != configtypes.ConvertTargetToContextType(c.Target) {
		return errors.Errorf("specified Target(%s) and ContextType(%s) for the Context object does not match", c.Target, c.ContextType)
	}
	return nil
}

// GetCurrentContext retrieves the current context for the specified target
//
// Deprecated: GetCurrentContext is deprecated. Use GetActiveContext instead
func GetCurrentContext(target configtypes.Target) (c *configtypes.Context, err error) {
	return GetActiveContext(configtypes.ConvertTargetToContextType(target))
}

// GetActiveContext retrieves the active context for the specified contextType
func GetActiveContext(contextType configtypes.ContextType) (c *configtypes.Context, err error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getActiveContext(node, contextType)
}

// GetContextsByType retrieves the contexts of a provided context type
func GetContextsByType(contextType configtypes.ContextType) ([]*configtypes.Context, error) {
	var results []*configtypes.Context

	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	for _, ctx := range cfg.KnownContexts {
		if ctx.ContextType == contextType {
			results = append(results, ctx)
		}
	}
	return results, nil
}

// GetAllCurrentContextsMap returns all current context per Target
//
// Deprecated: GetAllCurrentContextsMap is deprecated. Use GetAllActiveContextsMap instead
func GetAllCurrentContextsMap() (map[configtypes.Target]*configtypes.Context, error) {
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return nil, err
	}
	return getAllCurrentContextsMap(node)
}

// GetAllActiveContextsMap returns all active context per ContextType
func GetAllActiveContextsMap() (map[configtypes.ContextType]*configtypes.Context, error) {
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return nil, err
	}
	return getAllActiveContextsMap(node)
}

// GetAllActiveContextsList returns all active context names as list
func GetAllActiveContextsList() ([]string, error) {
	currentContextsMap, err := GetAllActiveContextsMap()
	if err != nil {
		return nil, err
	}
	var serverNames []string
	for _, context := range currentContextsMap {
		serverNames = append(serverNames, context.Name)
	}
	return serverNames, nil
}

// GetAllCurrentContextsList returns all current context names as list
//
// Deprecated: GetAllCurrentContextsList is deprecated. Use GetAllActiveContextsList instead
func GetAllCurrentContextsList() ([]string, error) {
	return GetAllActiveContextsList()
}

// SetCurrentContext sets the current context to the specified name if context is present
//
// Deprecated: SetCurrentContext is deprecated. Use SetActiveContext instead
func SetCurrentContext(name string) error {
	return SetActiveContext(name)
}

// SetActiveContext sets the active context to the specified name if context is present
func SetActiveContext(name string) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	ctx, err := getContext(node, name)
	if err != nil {
		return err
	}
	persist, err := setCurrentContext(node, ctx.Name, ctx.ContextType)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}
	if ctx.ContextType == configtypes.ContextTypeK8s {
		persist, err = setCurrentServer(node, name)
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
	return err
}

// RemoveCurrentContext removed the current context of specified context type
//
// Deprecated: RemoveCurrentContext is deprecated. Use RemoveActiveContext instead
func RemoveCurrentContext(target configtypes.Target) error {
	return RemoveActiveContext(configtypes.ConvertTargetToContextType(target))
}

// RemoveActiveContext removed the current context of specified context type
func RemoveActiveContext(contextType configtypes.ContextType) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	c, err := getActiveContext(node, contextType)
	if err != nil {
		return err
	}
	err = removeCurrentContext(node, "", contextType)
	if err != nil {
		return err
	}
	err = removeCurrentServer(node, c.Name)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

// EndpointFromContext retrieved the endpoint from the specified context
func EndpointFromContext(s *configtypes.Context) (endpoint string, err error) {
	missingFieldsErrMsg := "invalid context. Required fields missing in the context"
	switch s.ContextType {
	case configtypes.ContextTypeK8s:
		if s.ClusterOpts == nil {
			return endpoint, errors.New(missingFieldsErrMsg)
		}
		return s.ClusterOpts.Endpoint, nil
	case configtypes.ContextTypeTMC:
		if s.GlobalOpts == nil {
			return endpoint, errors.New(missingFieldsErrMsg)
		}
		return s.GlobalOpts.Endpoint, nil
	case configtypes.ContextTypeTanzu:
		if s.ClusterOpts == nil {
			return endpoint, errors.New(missingFieldsErrMsg)
		}
		return s.ClusterOpts.Endpoint, nil
	default:
		return endpoint, fmt.Errorf("unknown context type %q", s.ContextType)
	}
}

func getContext(node *yaml.Node, name string) (*configtypes.Context, error) {
	// check if context name is empty
	if name == "" {
		return nil, errors.New("context name cannot be empty")
	}

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	for _, ctx := range cfg.KnownContexts {
		if ctx.Name == name {
			return ctx, nil
		}
	}
	return nil, fmt.Errorf("context %v not found", name)
}

func getActiveContext(node *yaml.Node, contextType configtypes.ContextType) (*configtypes.Context, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	return cfg.GetActiveContext(contextType)
}

// Deprecated: getAllCurrentContextsMap is deprecated. Use getAllActiveContextsMap instead
func getAllCurrentContextsMap(node *yaml.Node) (map[configtypes.Target]*configtypes.Context, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	return cfg.GetAllCurrentContextsMap()
}

func getAllActiveContextsMap(node *yaml.Node) (map[configtypes.ContextType]*configtypes.Context, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	return cfg.GetAllActiveContextsMap()
}

func setContexts(node *yaml.Node, contexts []*configtypes.Context) (err error) {
	for _, c := range contexts {
		_, err = setContext(node, c)
		if err != nil {
			return err
		}
	}
	return err
}

func setContext(node *yaml.Node, ctx *configtypes.Context) (persist bool, err error) {
	// validate ctx object
	err = validateContext(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error while validating the Context object")
	}

	// Fill missing ContextType or Target in the Context object
	fillMissingContextTypeInContext(ctx)
	fillMissingTargetInContext(ctx)

	// Get Patch Strategies
	patchStrategies := constructPatchStrategies()

	var persistDiscoverySources bool

	// Convert context to node
	newContextNode, err := convertObjectToNode(ctx)
	if err != nil {
		return persist, err
	}

	// Find the contexts node from the root node
	keys := []nodeutils.Key{
		{Name: KeyContexts, Type: yaml.SequenceNode},
	}
	contextsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if contextsNode == nil {
		return persist, err
	}

	exists := false
	var result []*yaml.Node
	// Skip duplicate for context and server similar logic
	//nolint:dupl
	for _, contextNode := range contextsNode.Content {
		if index := nodeutils.GetNodeIndex(contextNode.Content, "name"); index != -1 &&
			contextNode.Content[index].Value == ctx.Name {
			exists = true
			// replace the nodes as per patch strategy
			_, err = nodeutils.DeleteNodes(newContextNode.Content[0], contextNode, nodeutils.WithPatchStrategyKey(KeyContexts), nodeutils.WithPatchStrategies(patchStrategies))
			if err != nil {
				return false, err
			}
			persist, err = nodeutils.MergeNodes(newContextNode.Content[0], contextNode)
			if err != nil {
				return false, err
			}
			persistDiscoverySources, err = setDiscoverySources(contextNode, ctx.DiscoverySources, nodeutils.WithPatchStrategyKey(fmt.Sprintf("%v.%v", KeyContexts, KeyDiscoverySources)), nodeutils.WithPatchStrategies(patchStrategies))
			if err != nil {
				return false, err
			}
			// merge the discovery sources to context
			if persistDiscoverySources {
				_, err = nodeutils.MergeNodes(newContextNode.Content[0], contextNode)
				if err != nil {
					return false, err
				}
			}
			result = append(result, contextNode)
			continue
		}
		result = append(result, contextNode)
	}
	if !exists {
		result = append(result, newContextNode.Content[0])
		persist = true
	}
	contextsNode.Content = result
	return persistDiscoverySources || persist, err
}

// Get Patch Strategies from config metadata
// By default;  AdditionalMetadata field will be patched in replace strategy if there are no patch strategies
func constructPatchStrategies() map[string]string {
	patchStrategies, err := GetConfigMetadataPatchStrategy()
	if err != nil {
		patchStrategies = map[string]string{
			"contexts.additionalMetadata": "replace",
		}
	}
	// Verify if there are patch strategies defined for `contexts.additionalMetadata` if not set replace by default
	if patchStrategies != nil && patchStrategies["contexts.additionalMetadata"] != "merge" {
		patchStrategies["contexts.additionalMetadata"] = "replace"
	}
	return patchStrategies
}

func setCurrentContext(node *yaml.Node, ctxName string, ctxType configtypes.ContextType) (persist bool, err error) {
	// Find current context node in the yaml node
	keys := []nodeutils.Key{
		{Name: KeyCurrentContext, Type: yaml.MappingNode},
	}
	currentContextNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if currentContextNode == nil {
		return persist, nodeutils.ErrNodeNotFound
	}
	if index := nodeutils.GetNodeIndex(currentContextNode.Content, string(ctxType)); index != -1 {
		if currentContextNode.Content[index].Value != ctxName {
			currentContextNode.Content[index].Value = ctxName
			currentContextNode.Content[index].Style = 0
			persist = true
		}
	} else {
		currentContextNode.Content = append(currentContextNode.Content, nodeutils.CreateScalarNode(string(ctxType), ctxName)...)
		persist = true
	}
	// maintain mutual exclusive behavior among all the current context types except TMC
	// (i.e. there can only be one active current context among all the context types except TMC.
	//  TMC context type can still be active when other context types are active)
	if persist {
		if err := updateMutualExclusiveCurrentContexts(node, ctxType); err != nil {
			return persist, err
		}
	}
	return persist, err
}

func removeCurrentContext(node *yaml.Node, ctxName string, ctxType configtypes.ContextType) error {
	// Find current context node in the yaml node
	keys := []nodeutils.Key{
		{Name: KeyCurrentContext},
	}

	currentContextNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if currentContextNode == nil {
		return nil
	}
	ctNodeIndex := nodeutils.GetNodeIndex(currentContextNode.Content, string(ctxType))
	if ctNodeIndex == -1 {
		return nil
	}
	if currentContextNode.Content[ctNodeIndex].Value == ctxName || ctxName == "" {
		ctNodeIndex--
		currentContextNode.Content = append(currentContextNode.Content[:ctNodeIndex], currentContextNode.Content[ctNodeIndex+2:]...)
	}
	return nil
}

//nolint:dupl
func removeContext(node *yaml.Node, name string) error {
	// check if context name is empty
	if name == "" {
		return errors.New("context name cannot be empty")
	}

	// Find the contexts node in the yaml node
	keys := []nodeutils.Key{
		{Name: KeyContexts},
	}
	contextsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if contextsNode == nil {
		return nil
	}
	var contexts []*yaml.Node
	for _, contextNode := range contextsNode.Content {
		if index := nodeutils.GetNodeIndex(contextNode.Content, "name"); index != -1 && contextNode.Content[index].Value == name {
			continue
		}
		contexts = append(contexts, contextNode)
	}
	contextsNode.Content = contexts
	return nil
}

// updateMutualExclusiveCurrentContexts updates the current contexts to maintain
// mutual exclusive behavior among the current context types except TMC
func updateMutualExclusiveCurrentContexts(node *yaml.Node, setterCtxType configtypes.ContextType) error {
	if setterCtxType == configtypes.ContextTypeTMC {
		return nil
	}

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return err
	}
	// deactivate all the other existing current contexts that are not TMC
	for contextType, contextName := range cfg.CurrentContext {
		if contextType == setterCtxType || contextType == configtypes.ContextTypeTMC {
			continue
		}

		err = removeCurrentContext(node, "", contextType)
		if err != nil {
			return err
		}
		err = removeCurrentServer(node, contextName)
		if err != nil {
			return err
		}
	}
	return nil
}
