// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package context

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

// CfgContextArgs used to construct input and output options
type CfgContextArgs struct {
	RuntimeAPIVersion  *core.RuntimeAPIVersion
	ContextName        string
	Target             types.Target
	Type               types.ContextType
	SetCurrentContext  bool
	GlobalOpts         *types.GlobalServerOpts
	ClusterOpts        *types.ClusterServerOpts
	DiscoverySources   []types.PluginDiscoveryOpts
	ValidationStrategy core.ValidationStrategy
	Error              bool
}

type CfgContextArgsOption func(*CfgContextArgs)

func WithContextName(name string) CfgContextArgsOption {
	return func(c *CfgContextArgs) {
		c.ContextName = name
	}
}

func WithRuntimeVersion(version core.RuntimeVersion) CfgContextArgsOption {
	return func(c *CfgContextArgs) {
		c.RuntimeAPIVersion = &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		}
	}
}

func WithError() CfgContextArgsOption {
	return func(c *CfgContextArgs) {
		c.Error = true
	}
}

// GetContextInputOptions used to generate GetContext command
type GetContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

// GetContextOutputOptions used to generate GetContext command
type GetContextOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	*types.ContextOpts                              // For specific version options look into ContextOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetContextInputOptions used to generate SetContext command
type SetContextInputOptions struct {
	*core.RuntimeAPIVersion      // required
	*types.ContextOpts           // required
	SetCurrentContext       bool // required
}

// SetContextOutputOptions used to generate SetContext command
type SetContextOutputOptions struct {
	ValidationStrategy core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error              string                  // expected error message could be the sub string of actual error message
}

// DeleteContextInputOptions used to generate DeleteContext command
type DeleteContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

// DeleteContextOutputOptions used to generate DeleteContext command
type DeleteContextOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetCurrentContextInputOptions used to generate SetCurrentContext command
type SetCurrentContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

// SetCurrentContextOutputOptions used to generate SetCurrentContext command
type SetCurrentContextOutputOptions struct {
	*core.RuntimeAPIVersion        // required
	Error                   string // expected error message could be the sub string of actual error message
}

// GetCurrentContextInputOptions used to generate GetCurrentContext command
type GetCurrentContextInputOptions struct {
	*core.RuntimeAPIVersion                   // required
	Target                  types.Target      // required for v1.0.0 - v0.28.0
	ContextType             types.ContextType // required for v0.25.4
}

// GetCurrentContextOutputOptions used to generate GetCurrentContext command
type GetCurrentContextOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	*types.ContextOpts                              // For specific version options look into ContextOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// RemoveCurrentContextInputOptions used to generate RemoveCurrentContext command
type RemoveCurrentContextInputOptions struct {
	*core.RuntimeAPIVersion              // required
	Target                  types.Target // required for v1.0.0 - v0.28.0
}

// RemoveCurrentContextOutputOptions used to generate RemoveCurrentContext command
type RemoveCurrentContextOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}
