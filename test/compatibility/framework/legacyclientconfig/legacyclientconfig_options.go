// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package legacyclientconfig

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/clidiscoverysources"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/featureflags"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

// GetClientConfigInputOptions used to generate GetClientConfig command
type GetClientConfigInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetClientConfigOutputOptions used to generate GetClientConfig command
type GetClientConfigOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ClientConfigOpts        *types.ClientConfigOpts // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// StoreClientConfigInputOptions used to generate StoreClientConfig command
type StoreClientConfigInputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ClientConfigOpts        *types.ClientConfigOpts // required
}

// StoreClientConfigOutputOptions used to generate StoreClientConfig command
type StoreClientConfigOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// CfgClientConfigArgs used to build the arguments for client config commands
type CfgClientConfigArgs struct {
	*core.RuntimeAPIVersion
	ClientConfigOpts   *types.ClientConfigOpts
	ValidationStrategy core.ValidationStrategy
	Error              string
}

type CfgClientConfigArgsOption func(*CfgClientConfigArgs)

// WithCLIDiscoverySources used to build the cli discovery sources arguments for client config commands
func WithCLIDiscoverySources(version core.RuntimeVersion, sources []types.PluginDiscoveryOpts) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		switch version {
		case core.VersionLatest, core.Version102, core.Version090:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				CoreCliOptions: &types.CoreCliOptionsOpts{
					DiscoverySources: sources,
				},
			}
		default:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				ClientOptions: &types.ClientOptionsOpts{
					CLI: &types.CLIOptionsOpts{
						DiscoverySources: sources,
					},
				},
			}
		}
	}
}

// WithDefaultCLIDiscoverySource used to build the default cli discovery sources arguments for client config commands
func WithDefaultCLIDiscoverySource(version core.RuntimeVersion) CfgClientConfigArgsOption {
	defaultPluginDiscoverySource := clidiscoverysources.DefaultCLIDiscoverySourcePerVersion(version)
	return func(c *CfgClientConfigArgs) {
		switch version {
		case core.VersionLatest, core.Version102, core.Version090:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				CoreCliOptions: &types.CoreCliOptionsOpts{
					DiscoverySources: []types.PluginDiscoveryOpts{
						*defaultPluginDiscoverySource.PluginDiscoveryOpts,
					},
				},
			}
		default:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				ClientOptions: &types.ClientOptionsOpts{
					CLI: &types.CLIOptionsOpts{
						DiscoverySources: []types.PluginDiscoveryOpts{
							*defaultPluginDiscoverySource.PluginDiscoveryOpts,
						},
					},
				},
			}
		}
	}
}

// WithDefaultServer used to build the default server arguments for client config commands
func WithDefaultServer(version core.RuntimeVersion) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		switch version {
		case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				KnownServers: []*types.ServerOpts{
					{
						Name: common.CompatibilityTestOne,
						Type: types.ManagementClusterServerType,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: common.DefaultEndpoint,
						},
					},
				},
				CurrentServer: common.CompatibilityTestOne,
			}

		case core.Version0254:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				KnownServers: []*types.ServerOpts{
					{
						Name: common.CompatibilityTestOne,
						Type: types.ManagementClusterServerType,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: common.DefaultEndpoint,
						},
					},
				},
				CurrentServer: common.CompatibilityTestOne,
			}
		}
	}
}

// WithDefaultContextAndServer used to build the default server and context arguments for client config commands
func WithDefaultContextAndServer(version core.RuntimeVersion) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		switch version {
		case core.VersionLatest:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				KnownServers: []*types.ServerOpts{
					{
						Name: common.CompatibilityTestOne,
						Type: types.ManagementClusterServerType,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: common.DefaultEndpoint,
						},
					},
				},
				CurrentServer: common.CompatibilityTestOne,
				CurrentContext: map[string]string{
					string(types.TargetK8s): common.CompatibilityTestOne,
				},
				KnownContexts: []*types.ContextOpts{
					{
						Name: common.CompatibilityTestOne,
						// Note: We are not setting Target anymore with the latest CLI because
						// it should be automatically configured by API for backwards compatibility
						ContextType: types.ContextTypeK8s,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: "default-compatibility-test-endpoint",
						},
					},
				},
			}
		case core.Version102, core.Version090, core.Version0280:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				KnownServers: []*types.ServerOpts{
					{
						Name: common.CompatibilityTestOne,
						Type: types.ManagementClusterServerType,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: common.DefaultEndpoint,
						},
					},
				},
				CurrentServer: common.CompatibilityTestOne,
				CurrentContext: map[string]string{
					string(types.TargetK8s): common.CompatibilityTestOne,
				},
				KnownContexts: []*types.ContextOpts{
					{
						Name:   common.CompatibilityTestOne,
						Target: types.TargetK8s,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: "default-compatibility-test-endpoint",
						},
					},
				},
			}
		case core.Version0254:
			c.ClientConfigOpts = &types.ClientConfigOpts{
				KnownServers: []*types.ServerOpts{
					{
						Name: common.CompatibilityTestOne,
						Type: types.ManagementClusterServerType,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: common.DefaultEndpoint,
						},
					},
				},
				CurrentServer: common.CompatibilityTestOne,
				CurrentContext: map[string]string{
					string(types.CtxTypeK8s): common.CompatibilityTestOne,
				},
				KnownContexts: []*types.ContextOpts{
					{
						Name: common.CompatibilityTestOne,
						Type: types.CtxTypeK8s,
						GlobalOpts: &types.GlobalServerOpts{
							Endpoint: "default-compatibility-test-endpoint",
						},
					},
				},
			}
		}
	}
}

// WithDefaultFeatureFlags used to build the default feature flags arguments for client config commands
func WithDefaultFeatureFlags() CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		c.ClientConfigOpts = &types.ClientConfigOpts{
			ClientOptions: &types.ClientOptionsOpts{
				Features: map[string]types.FeatureMap{
					featureflags.CompatibilityTestsPlugin: map[string]string{
						featureflags.CompatibilityTestsPluginKey: "true",
					},
				},
			},
		}
	}
}

// WithFeatureFlags used to build the feature flags arguments for client config commands
func WithFeatureFlags(features map[string]types.FeatureMap) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		c.ClientConfigOpts = &types.ClientConfigOpts{
			ClientOptions: &types.ClientOptionsOpts{
				Features: features,
			},
		}
	}
}

func WithClientConfigOpts(cf *types.ClientConfigOpts) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		c.ClientConfigOpts = cf
	}
}

func WithRuntimeVersion(v core.RuntimeVersion) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		c.RuntimeVersion = v
	}
}

func WithValidationStrategy(vs core.ValidationStrategy) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		c.ValidationStrategy = vs
	}
}

func WithError(e string) CfgClientConfigArgsOption {
	return func(c *CfgClientConfigArgs) {
		c.Error = e
	}
}
