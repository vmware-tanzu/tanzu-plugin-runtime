// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clidiscoverysources

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

// GetCLIDiscoverySourceInputOptions used to generate GetCLIDiscoverySource command
type GetCLIDiscoverySourceInputOptions struct {
	*core.RuntimeAPIVersion        // required
	DiscoverySourceName     string // required
}

// GetCLIDiscoverySourceOutputOptions used to generate GetCLIDiscoverySource command
type GetCLIDiscoverySourceOutputOptions struct {
	*core.RuntimeAPIVersion                            // required
	PluginDiscoveryOpts     *types.PluginDiscoveryOpts // For specific version options look into DiscoverySourceOpts definition
	ValidationStrategy      core.ValidationStrategy    // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                     // expected error message could be the sub string of actual error message
}

// SetCLIDiscoverySourceInputOptions used to generate SetCLIDiscoverySource command
type SetCLIDiscoverySourceInputOptions struct {
	*core.RuntimeAPIVersion                            // required
	PluginDiscoveryOpts     *types.PluginDiscoveryOpts // For specific version options look into DiscoverySourceOpts definition
}

// SetCLIDiscoverySourceOutputOptions used to generate SetCLIDiscoverySource command
type SetCLIDiscoverySourceOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	Error                   string                  // expected error message could be the sub string of actual error message
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// DeleteCLIDiscoverySourceInputOptions used to generate DeleteCLIDiscoverySource command
type DeleteCLIDiscoverySourceInputOptions struct {
	*core.RuntimeAPIVersion        // required
	DiscoverySourceName     string // required
}

// DeleteCLIDiscoverySourceOutputOptions used to generate DeleteCLIDiscoverySource command
type DeleteCLIDiscoverySourceOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// CfgCLIDiscoverySourcesArgs used to construct input and output options
type CfgCLIDiscoverySourcesArgs struct {
	*core.RuntimeAPIVersion
	DiscoverySourceName string
	PluginDiscoveryOpts *types.PluginDiscoveryOpts // For specific version options look into DiscoverySourceOpts definition
	ValidationStrategy  core.ValidationStrategy
	Error               string
}

type CfgCLIDiscoverySourcesArgsOption func(*CfgCLIDiscoverySourcesArgs)

func WithRuntimeAPIVersion(v *core.RuntimeAPIVersion) CfgCLIDiscoverySourcesArgsOption {
	return func(c *CfgCLIDiscoverySourcesArgs) {
		c.RuntimeAPIVersion = v
	}
}
func WithDiscoverySourceName(name string) CfgCLIDiscoverySourcesArgsOption {
	return func(c *CfgCLIDiscoverySourcesArgs) {
		c.DiscoverySourceName = name
	}
}
func WithError(e string) CfgCLIDiscoverySourcesArgsOption {
	return func(c *CfgCLIDiscoverySourcesArgs) {
		c.Error = e
	}
}

func NewCfgCLIDiscoverySourcesArgs(options ...CfgCLIDiscoverySourcesArgsOption) *CfgCLIDiscoverySourcesArgs {
	// Default Value
	p := &CfgCLIDiscoverySourcesArgs{
		PluginDiscoveryOpts: &types.PluginDiscoveryOpts{},
	}

	for _, opt := range options {
		opt(p)
	}

	return p
}

func WithGCPDiscoveryOpts(opts *types.GCPDiscoveryOpts) CfgCLIDiscoverySourcesArgsOption {
	return func(p *CfgCLIDiscoverySourcesArgs) {
		p.PluginDiscoveryOpts.GCP = opts
	}
}

func WithOCIDiscoveryOpts(opts *types.OCIDiscoveryOpts) CfgCLIDiscoverySourcesArgsOption {
	return func(p *CfgCLIDiscoverySourcesArgs) {
		p.PluginDiscoveryOpts.OCI = opts
	}
}

func WithGenericRESTDiscoveryOpts(opts *types.GenericRESTDiscoveryOpts) CfgCLIDiscoverySourcesArgsOption {
	return func(p *CfgCLIDiscoverySourcesArgs) {
		p.PluginDiscoveryOpts.REST = opts
	}
}

func WithKubernetesDiscoveryOpts(opts *types.KubernetesDiscoveryOpts) CfgCLIDiscoverySourcesArgsOption {
	return func(p *CfgCLIDiscoverySourcesArgs) {
		p.PluginDiscoveryOpts.Kubernetes = opts
	}
}

func WithLocalDiscoveryOpts(opts *types.LocalDiscoveryOpts) CfgCLIDiscoverySourcesArgsOption {
	return func(p *CfgCLIDiscoverySourcesArgs) {
		p.PluginDiscoveryOpts.Local = opts
	}
}

func WithContextType(contextType types.ContextType) CfgCLIDiscoverySourcesArgsOption {
	return func(p *CfgCLIDiscoverySourcesArgs) {
		p.PluginDiscoveryOpts.ContextType = contextType
	}
}
