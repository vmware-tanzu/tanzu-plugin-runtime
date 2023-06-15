// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package metadata contains all the cross version api compatibility tests for context apis
package metadata

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

// Metadata related constants

const (
	CompatibilityTestsMetadataPatchStrategyKey   = "compatibility-tests.contexts.name"
	CompatibilityTestsMetadataPatchStrategyValue = "replace"
	CompatibilityTestsMetadataSettingsKey        = "useUnifiedConfig"
	CompatibilityTestsMetadataSettingsValue      = "true"
)

// DefaultSetConfigMetadataPatchStrategyCommand creates a SetConfigMetadataPatchStrategy Command with default input and output options
//
//nolint:dupl
func DefaultSetConfigMetadataPatchStrategyCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	args := &CfgMetadataArgs{
		Key:   CompatibilityTestsMetadataPatchStrategyKey,
		Value: CompatibilityTestsMetadataPatchStrategyValue,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetConfigMetadataPatchStrategyInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key:   args.Key,
		Value: args.Value,
	}

	outputOpts := &SetConfigMetadataPatchStrategyOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewSetConfigMetadataPatchStrategyCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultSetConfigMetadataSettingCommand creates a SetConfigMetadataSetting Command with default input and output options
//
//nolint:dupl
func DefaultSetConfigMetadataSettingCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	args := &CfgMetadataArgs{
		Key:   CompatibilityTestsMetadataSettingsKey,
		Value: CompatibilityTestsMetadataSettingsValue,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetConfigMetadataSettingInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key:   args.Key,
		Value: args.Value,
	}

	outputOpts := &SetConfigMetadataSettingOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewSetConfigMetadataSettingCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultDeleteConfigMetadataSettingCommand creates a DeleteConfigMetadataSetting Command with default input and output options
func DefaultDeleteConfigMetadataSettingCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	args := &CfgMetadataArgs{
		Key: CompatibilityTestsMetadataSettingsKey,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &DeleteConfigMetadataSettingInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key: args.Key,
	}

	outputOpts := &DeleteConfigMetadataSettingOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewDeleteConfigMetadataSettingCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetMetadataCommand creates a GetMetadataCommand Command with default input and output options
func DefaultGetMetadataCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	patchStrategies := map[string]string{
		CompatibilityTestsMetadataPatchStrategyKey: CompatibilityTestsMetadataPatchStrategyValue,
	}

	settings := map[string]string{
		CompatibilityTestsMetadataSettingsKey: CompatibilityTestsMetadataSettingsValue,
	}

	args := &CfgMetadataArgs{
		MetadataOpts: &types.MetadataOpts{
			ConfigMetadata: &types.ConfigMetadataOpts{
				PatchStrategy: patchStrategies,
				Settings:      settings,
			},
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetMetadataInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}

	outputOpts := &GetMetadataOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		MetadataOpts:       args.MetadataOpts,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewGetMetadataCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetConfigMetadataCommand creates a GetConfigMetadata Command with default input and output options
func DefaultGetConfigMetadataCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	patchStrategies := map[string]string{
		CompatibilityTestsMetadataPatchStrategyKey: CompatibilityTestsMetadataPatchStrategyValue,
	}

	settings := map[string]string{
		CompatibilityTestsMetadataSettingsKey: CompatibilityTestsMetadataSettingsValue,
	}

	args := &CfgMetadataArgs{
		ConfigMetadataOpts: &types.ConfigMetadataOpts{
			PatchStrategy: patchStrategies,
			Settings:      settings,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetConfigMetadataInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}

	outputOpts := &GetConfigMetadataOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ConfigMetadataOpts: args.ConfigMetadataOpts,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewGetConfigMetadataCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetConfigMetadataPatchStrategyCommand creates a GetConfigMetadataPatchStrategy Command with default input and output options
//
//nolint:dupl
func DefaultGetConfigMetadataPatchStrategyCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	patchStrategies := map[string]string{
		CompatibilityTestsMetadataPatchStrategyKey: CompatibilityTestsMetadataPatchStrategyValue,
	}

	args := &CfgMetadataArgs{
		PatchStrategy: patchStrategies,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetConfigMetadataPatchStrategyInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}

	outputOpts := &GetConfigMetadataPatchStrategyOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		PatchStrategy:      args.PatchStrategy,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewGetConfigMetadataPatchStrategyCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetConfigMetadataSettingCommand creates a GetConfigMetadataSetting Command with default input and output options
func DefaultGetConfigMetadataSettingCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	args := &CfgMetadataArgs{
		Key:   CompatibilityTestsMetadataSettingsKey,
		Value: CompatibilityTestsMetadataSettingsValue,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetConfigMetadataSettingInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key: args.Key,
	}

	outputOpts := &GetConfigMetadataSettingOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Value:              args.Value,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewGetConfigMetadataSettingCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetConfigMetadataSettingsCommand creates a GetConfigMetadataSettings Command with default input and output options
//
//nolint:dupl
func DefaultGetConfigMetadataSettingsCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	settings := map[string]string{
		CompatibilityTestsMetadataSettingsKey: CompatibilityTestsMetadataSettingsValue,
	}

	args := &CfgMetadataArgs{
		Settings: settings,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetConfigMetadataSettingsInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}

	outputOpts := &GetConfigMetadataSettingsOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		MetadataSettings:   args.Settings,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewGetConfigMetadataSettingsCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultIsConfigMetadataSettingsEnabledCommand creates a IsConfigMetadataSettingsEnabled Command with default input and output options
func DefaultIsConfigMetadataSettingsEnabledCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	args := &CfgMetadataArgs{
		Key:     CompatibilityTestsMetadataSettingsKey,
		Enabled: true,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &IsConfigMetadataSettingsEnabledInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key: args.Key,
	}

	outputOpts := &IsConfigMetadataSettingsEnabledOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Enabled:            args.Enabled,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}
	cmd, err := NewIsConfigMetadataSettingsEnabledCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultUseUnifiedConfigCommand creates a UseUnifiedConfig Command with default input and output options
func DefaultUseUnifiedConfigCommand(version core.RuntimeVersion, opts ...CfgMetadataOptionArgs) *core.Command {
	args := &CfgMetadataArgs{
		Enabled: true,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &UseUnifiedConfigInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}

	outputOpts := &UseUnifiedConfigOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Enabled:            args.Enabled,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewUseUnifiedConfigCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}
