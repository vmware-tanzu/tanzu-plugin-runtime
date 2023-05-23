// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	configtypes "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

	"gopkg.in/yaml.v3"
)

// parseContext unmarshalls string to Context struct
func parseContext(context string) (*configtypes.Context, error) {
	var ctx configtypes.Context
	err := yaml.Unmarshal([]byte(context), &ctx)
	if err != nil {
		return nil, err
	}
	return &ctx, nil
}

// parseServer unmarshalls string to Server struct
func parseServer(server string) (*configtypes.Server, error) {
	var s configtypes.Server
	err := yaml.Unmarshal([]byte(server), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// parseCLIDiscoverySource unmarshalls string to PluginDiscovery struct
func parseCLIDiscoverySource(source string) (*configtypes.PluginDiscovery, error) {
	var pluginDiscovery configtypes.PluginDiscovery
	err := yaml.Unmarshal([]byte(source), &pluginDiscovery)
	if err != nil {
		return nil, err
	}
	return &pluginDiscovery, nil
}
