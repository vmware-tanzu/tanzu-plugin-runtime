// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package featureflags

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

const (
	CompatibilityTestsPlugin     string = "compatibility-tests-plugin"
	CompatibilityTestsPluginKey  string = "compatibility-tests-plugin-key"
	CompatibilityTestsPluginKey0 string = "compatibility-tests-plugin-key0"
)

// DefaultIsFeatureEnabledCommand creates a IsFeatureEnabled Command with default input and output options
func DefaultIsFeatureEnabledCommand(version core.RuntimeVersion, opts ...CfgFeatureOptionArgs) *core.Command {
	args := &CfgFeatureArgs{
		Plugin:         CompatibilityTestsPlugin,
		Key:            CompatibilityTestsPluginKey,
		FeatureEnabled: true,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &IsFeatureEnabledInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Plugin: args.Plugin,
		Key:    args.Key,
	}

	outputOpts := &IsFeatureEnabledOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		FeatureEnabled:     args.FeatureEnabled,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewIsFeatureEnabledCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultDeleteFeatureCommand creates a DeleteFeature Command with default input and output options
func DefaultDeleteFeatureCommand(version core.RuntimeVersion, opts ...CfgFeatureOptionArgs) *core.Command {
	args := &CfgFeatureArgs{
		Plugin: CompatibilityTestsPlugin,
		Key:    CompatibilityTestsPluginKey,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &DeleteFeatureInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Plugin: args.Plugin,
		Key:    args.Key,
	}

	outputOpts := &DeleteFeatureOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewDeleteFeatureCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultSetFeatureCommand creates a SetFeature Command with default input and output options
func DefaultSetFeatureCommand(version core.RuntimeVersion, opts ...CfgFeatureOptionArgs) *core.Command {
	args := &CfgFeatureArgs{
		Plugin: CompatibilityTestsPlugin,
		Key:    CompatibilityTestsPluginKey,
		Value:  "true",
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetFeatureInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Plugin: args.Plugin,
		Key:    args.Key,
		Value:  args.Value,
	}

	outputOpts := &SetFeatureOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewSetFeatureCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}
