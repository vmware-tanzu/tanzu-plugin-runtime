// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
)

// PopulateContexts converts the known servers that are missing in contexts.
// This is needed when reading the config file persisted by an older core or plugin,
// so that it is forwards compatible with a new core plugin.
// Returns true if there was any delta.
func PopulateContexts(cfg *configtypes.ClientConfig) bool {
	if cfg == nil || len(cfg.KnownServers) == 0 {
		return false
	}

	var delta bool
	if len(cfg.KnownContexts) == 0 {
		cfg.KnownContexts = make([]*configtypes.Context, 0, len(cfg.KnownServers))
	}
	for _, s := range cfg.KnownServers {
		if cfg.HasContext(s.Name) {
			// server already present in known contexts; skip
			continue
		}

		delta = true
		// convert and append the server to the list of known contexts
		c := convertServerToContext(s)
		cfg.KnownContexts = append(cfg.KnownContexts, c)

		if s.Name == cfg.CurrentServer {
			err := cfg.SetCurrentContext(c.Target, c.Name)
			if err != nil {
				log.Warningf(err.Error())
			}
		}
	}

	return delta
}

func convertServerToContext(s *configtypes.Server) *configtypes.Context {
	if s == nil {
		return nil
	}

	return &configtypes.Context{
		Name:             s.Name,
		Target:           convertServerTypeToTarget(s.Type),
		GlobalOpts:       s.GlobalOpts,
		ClusterOpts:      convertMgmtClusterOptsToClusterOpts(s.ManagementClusterOpts),
		DiscoverySources: s.DiscoverySources,
	}
}

func convertServerTypeToTarget(t configtypes.ServerType) string {
	switch t {
	case configtypes.ManagementClusterServerType:
		return configtypes.TargetK8s
	case configtypes.GlobalServerType:
		return configtypes.TargetTMC
	}
	// no other server type is supported in v0
	return string(t)
}

func convertMgmtClusterOptsToClusterOpts(s *configtypes.ManagementClusterServer) *configtypes.ClusterServer {
	if s == nil {
		return nil
	}

	return &configtypes.ClusterServer{
		Endpoint:            s.Endpoint,
		Path:                s.Path,
		Context:             s.Context,
		IsManagementCluster: true,
	}
}

// populateServers converts the known contexts that are missing in servers.
// This is needed when writing the config file from the newer core or plugin,
// so that it is backwards compatible with an older core or plugin.
func populateServers(cfg *configtypes.ClientConfig) {
	if cfg == nil {
		return
	}

	if len(cfg.KnownServers) == 0 {
		cfg.KnownServers = make([]*configtypes.Server, 0, len(cfg.KnownContexts))
	}
	for _, c := range cfg.KnownContexts {
		if cfg.HasServer(c.Name) {
			// context already present in known servers; skip
			continue
		}

		// convert and append the context to the list of known servers
		s := convertContextToServer(c)
		cfg.KnownServers = append(cfg.KnownServers, s)

		if cfg.CurrentServer == "" && (c.IsManagementCluster() || c.Target == configtypes.TargetTMC) && c.Name == cfg.CurrentContext[c.Target] {
			// This is lossy because only one server can be active at a time in the older CLI.
			// Using the K8s context for a management cluster or TMC, since these are the two
			// available publicly at the time of deprecation.
			cfg.CurrentServer = cfg.CurrentContext[configtypes.TargetK8s]
		}
	}
}

func convertContextToServer(c *configtypes.Context) *configtypes.Server {
	if c == nil {
		return nil
	}

	return &configtypes.Server{
		Name:                  c.Name,
		Type:                  convertTargetToServerType(c.Target),
		GlobalOpts:            c.GlobalOpts,
		ManagementClusterOpts: convertClusterOptsToMgmtClusterOpts(c.ClusterOpts),
		DiscoverySources:      c.DiscoverySources,
	}
}

func convertTargetToServerType(t string) configtypes.ServerType {
	switch t {
	case configtypes.TargetK8s:
		// This is lossy because only management cluster servers are supported by the older CLI.
		return configtypes.ManagementClusterServerType
	case configtypes.TargetTMC:
		return configtypes.GlobalServerType
	}
	// no other context type is supported in v1 yet
	return configtypes.ServerType(t)
}

func convertClusterOptsToMgmtClusterOpts(o *configtypes.ClusterServer) *configtypes.ManagementClusterServer {
	if o == nil || !o.IsManagementCluster {
		return nil
	}

	return &configtypes.ManagementClusterServer{
		Endpoint: o.Endpoint,
		Path:     o.Path,
		Context:  o.Context,
	}
}

// convertNodeToClientConfig converts yaml node to client config type
func convertNodeToClientConfig(node *yaml.Node) (obj *configtypes.ClientConfig, err error) {
	err = node.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert node to ClientConfig")
	}
	if obj == nil {
		return &configtypes.ClientConfig{}, err
	}
	return obj, err
}

// convertNodeToMetadata converts yaml node to client config type
func convertNodeToMetadata(node *yaml.Node) (obj *configtypes.Metadata, err error) {
	err = node.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert node to Metadata")
	}
	return obj, err
}

// convertObjectToNode converts a typed object to yaml node
func convertObjectToNode[
	T *configtypes.ClientConfig |
		*configtypes.Metadata |
		*configtypes.Server |
		*configtypes.PluginRepository |
		*configtypes.Context |
		*configtypes.Cert |
		*configtypes.TelemetryOptions |
		*configtypes.PluginDiscovery](obj T) (*yaml.Node, error) {

	bytes, err := yaml.Marshal(obj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert obj to node")
	}
	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal bytes to node")
	}
	return &node, nil
}
