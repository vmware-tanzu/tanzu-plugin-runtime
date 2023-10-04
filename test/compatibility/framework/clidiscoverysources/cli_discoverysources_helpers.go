// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clidiscoverysources

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

const (
	CLIDiscoverySourceNotFound    string = "cli discovery source not found"
	CompatibilityTestsSourceName  string = "compatibility-tests-source-name"
	CompatibilityTestsSourceImage string = "compatibility-tests-source-image"
)

// DefaultSetCLIDiscoverySourceCommand constructs a default SetCLIDiscoverySource command without an error.
func DefaultSetCLIDiscoverySourceCommand(version core.RuntimeVersion, opts ...CfgCLIDiscoverySourcesArgsOption) *core.Command {
	// Make input and output options
	defaultPluginDiscoverySource := DefaultCLIDiscoverySourcePerVersion(version)

	args := &CfgCLIDiscoverySourcesArgs{
		PluginDiscoveryOpts: defaultPluginDiscoverySource.PluginDiscoveryOpts,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetCLIDiscoverySourceInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		PluginDiscoveryOpts: args.PluginDiscoveryOpts,
	}

	outputOpts := &SetCLIDiscoverySourceOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Error: args.Error,
	}

	// Construct API command
	cmd, err := NewSetCLIDiscoverySourceCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetCLIDiscoverySourceCommand constructs a default GetCLIDiscoverySource command without an error.
func DefaultGetCLIDiscoverySourceCommand(version core.RuntimeVersion, opts ...CfgCLIDiscoverySourcesArgsOption) *core.Command {
	// Make input and output options
	defaultPluginDiscoverySource := DefaultCLIDiscoverySourcePerVersion(version)

	args := &CfgCLIDiscoverySourcesArgs{
		DiscoverySourceName: CompatibilityTestsSourceName,
		PluginDiscoveryOpts: defaultPluginDiscoverySource.PluginDiscoveryOpts,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetCLIDiscoverySourceInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		DiscoverySourceName: args.DiscoverySourceName,
	}

	outputOpts := &GetCLIDiscoverySourceOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		PluginDiscoveryOpts: args.PluginDiscoveryOpts,
		ValidationStrategy:  args.ValidationStrategy,
		Error:               args.Error,
	}

	// Construct API command
	cmd, err := NewGetCLIDiscoverySourceCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultDeleteCLIDiscoverySourceCommand constructs a default DeleteCLIDiscoverySource command without an error.
func DefaultDeleteCLIDiscoverySourceCommand(version core.RuntimeVersion, opts ...CfgCLIDiscoverySourcesArgsOption) *core.Command {
	args := &CfgCLIDiscoverySourcesArgs{
		DiscoverySourceName: CompatibilityTestsSourceName,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &DeleteCLIDiscoverySourceInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		DiscoverySourceName: args.DiscoverySourceName,
	}

	outputOpts := &DeleteCLIDiscoverySourceOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Error: args.Error,
	}

	// Construct API command
	cmd, err := NewDeleteCLIDiscoverySourceCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultCLIDiscoverySourcePerVersion constructs a default PluginDiscoveryOpts
func DefaultCLIDiscoverySourcePerVersion(version core.RuntimeVersion) *CfgCLIDiscoverySourcesArgs {
	oci := &types.OCIDiscoveryOpts{
		Name:  CompatibilityTestsSourceName,
		Image: CompatibilityTestsSourceImage,
	}

	switch version {
	case core.Version0116:
		return NewCfgCLIDiscoverySourcesArgs(
			WithOCIDiscoveryOpts(oci),
		)
	case core.Version0254:
		return NewCfgCLIDiscoverySourcesArgs(
			WithOCIDiscoveryOpts(oci),
			WithContextType(types.CtxTypeTMC),
		)
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		return NewCfgCLIDiscoverySourcesArgs(
			WithOCIDiscoveryOpts(oci),
		)
	}
	return nil
}
