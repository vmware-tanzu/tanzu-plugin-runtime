// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

// GetMetadataInputOptions used to generate GetEnv command
type GetMetadataInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetMetadataOutputOptions used to generate GetEnv command
type GetMetadataOutputOptions struct {
	*core.RuntimeAPIVersion // required
	MetadataOpts            *framework.MetadataOpts
	Error                   string
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// GetConfigMetadataInputOptions used to generate GetEnv command
type GetConfigMetadataInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetConfigMetadataOutputOptions used to generate GetEnv command
type GetConfigMetadataOutputOptions struct {
	*core.RuntimeAPIVersion // required
	ConfigMetadataOpts      *framework.ConfigMetadataOpts
	Error                   string
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// GetConfigMetadataPatchStrategyInputOptions used to generate GetEnv command
type GetConfigMetadataPatchStrategyInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetConfigMetadataPatchStrategyOutputOptions used to generate GetEnv command
type GetConfigMetadataPatchStrategyOutputOptions struct {
	*core.RuntimeAPIVersion // required
	PatchStrategy           map[string]string
	Error                   string
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// GetConfigMetadataSettingsInputOptions used to generate GetEnv command
type GetConfigMetadataSettingsInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetConfigMetadataSettingsOutputOptions used to generate GetEnv command
type GetConfigMetadataSettingsOutputOptions struct {
	*core.RuntimeAPIVersion // required
	MetadataSettings        map[string]string
	Error                   string
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// GetConfigMetadataSettingInputOptions used to generate GetEnv command
type GetConfigMetadataSettingInputOptions struct {
	*core.RuntimeAPIVersion // required
	Key                     string
}

// GetConfigMetadataSettingOutputOptions used to generate GetEnv command
type GetConfigMetadataSettingOutputOptions struct {
	*core.RuntimeAPIVersion // required
	Value                   string
	Error                   string
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// IsConfigMetadataSettingsEnabledInputOptions used to generate GetEnv command
type IsConfigMetadataSettingsEnabledInputOptions struct {
	*core.RuntimeAPIVersion // required
	Key                     string
}

// IsConfigMetadataSettingsEnabledOutputOptions used to generate GetEnv command
type IsConfigMetadataSettingsEnabledOutputOptions struct {
	*core.RuntimeAPIVersion // required
	Enabled                 bool
	Error                   string
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// UseUnifiedConfigInputOptions used to generate GetEnv command
type UseUnifiedConfigInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// UseUnifiedConfigOutputOptions used to generate GetEnv command
type UseUnifiedConfigOutputOptions struct {
	*core.RuntimeAPIVersion // required
	Enabled                 bool
	Error                   string
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
}

// DeleteConfigMetadataSettingInputOptions used to generate DeleteEnv command
type DeleteConfigMetadataSettingInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Key                     string // required
}

// DeleteConfigMetadataSettingOutputOptions used to generate DeleteEnv command
type DeleteConfigMetadataSettingOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetConfigMetadataSettingInputOptions used to generate SetEnv command
type SetConfigMetadataSettingInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Key                     string // required
	Value                   string // required
}

// SetConfigMetadataSettingOutputOptions used to generate SetEnv command
type SetConfigMetadataSettingOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetConfigMetadataPatchStrategyInputOptions used to generate SetEnv command
type SetConfigMetadataPatchStrategyInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Key                     string // required
	Value                   string // required
}

// SetConfigMetadataPatchStrategyOutputOptions used to generate SetEnv command
type SetConfigMetadataPatchStrategyOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

type CfgMetadataArgs struct {
	*core.RuntimeAPIVersion // required
	MetadataOpts            *framework.MetadataOpts
	ConfigMetadataOpts      *framework.ConfigMetadataOpts
	PatchStrategy           map[string]string
	Settings                map[string]string
	Key                     string // required
	Value                   string // required
	Enabled                 bool
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

type CfgMetadataOptionArgs func(*CfgMetadataArgs)

func WithRuntimeAPIVersion(v *core.RuntimeAPIVersion) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.RuntimeAPIVersion = v
	}
}

func WithMetadataOpts(m *framework.MetadataOpts) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.MetadataOpts = m
	}
}

func WithConfigMetadataOpts(cm *framework.ConfigMetadataOpts) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.ConfigMetadataOpts = cm
	}
}

func WithPatchStrategy(ps map[string]string) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.PatchStrategy = ps
	}
}

func WithMetadataSettings(ms map[string]string) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.Settings = ms
	}
}

func WithKey(k string) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.Key = k
	}
}

func WithValue(v string) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.Value = v
	}
}

func WithEnabled(e bool) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.Enabled = e
	}
}

func WithValidationStrategy(vs core.ValidationStrategy) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.ValidationStrategy = vs
	}
}

func WithError(e string) CfgMetadataOptionArgs {
	return func(c *CfgMetadataArgs) {
		c.Error = e
	}
}

func NewCfgMetadataArgs(opts ...CfgMetadataOptionArgs) *CfgMetadataArgs {
	c := &CfgMetadataArgs{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
