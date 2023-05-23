// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package featureflags

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// IsFeatureEnabledInputOptions used to generate GetFeature command
type IsFeatureEnabledInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Plugin                  string // required
	Key                     string // required
}

// IsFeatureEnabledOutputOptions used to generate GetFeature command
type IsFeatureEnabledOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	FeatureEnabled          bool                    // For specific version options look into FeatureOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetFeatureInputOptions used to generate SetFeature command
type SetFeatureInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Plugin                  string // required
	Key                     string // required
	Value                   string // required
}

// SetFeatureOutputOptions used to generate SetFeature command
type SetFeatureOutputOptions struct {
	*core.RuntimeAPIVersion        // required
	Error                   string // expected error message could be the sub string of actual error message
}

// DeleteFeatureInputOptions used to generate DeleteFeature command
type DeleteFeatureInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Plugin                  string // required
	Key                     string // required

}

// DeleteFeatureOutputOptions used to generate DeleteFeature command
type DeleteFeatureOutputOptions struct {
	*core.RuntimeAPIVersion        // required
	Error                   string // expected error message could be the sub string of actual error message
}

type CfgFeatureArgs struct {
	*core.RuntimeAPIVersion
	Plugin             string
	Key                string
	Value              string
	FeatureEnabled     bool
	ValidationStrategy core.ValidationStrategy
	Error              string
}

type CfgFeatureOptionArgs func(*CfgFeatureArgs)

func WithRuntimeAPIVersion(v *core.RuntimeAPIVersion) CfgFeatureOptionArgs {
	return func(c *CfgFeatureArgs) {
		c.RuntimeAPIVersion = v
	}
}

func WithPlugin(p string) CfgFeatureOptionArgs {
	return func(c *CfgFeatureArgs) {
		c.Plugin = p
	}
}

func WithFeatureEnabled(e bool) CfgFeatureOptionArgs {
	return func(c *CfgFeatureArgs) {
		c.FeatureEnabled = e
	}
}

func WithKey(k string) CfgFeatureOptionArgs {
	return func(c *CfgFeatureArgs) {
		c.Key = k
	}
}

func WithValue(v string) CfgFeatureOptionArgs {
	return func(c *CfgFeatureArgs) {
		c.Value = v
	}
}

func WithValidationStrategy(vs core.ValidationStrategy) CfgFeatureOptionArgs {
	return func(c *CfgFeatureArgs) {
		c.ValidationStrategy = vs
	}
}

func WithStrictValidationStrategy() CfgFeatureOptionArgs {
	return WithValidationStrategy(core.ValidationStrategyStrict)
}

func WithError(e string) CfgFeatureOptionArgs {
	return func(c *CfgFeatureArgs) {
		c.Error = e
	}
}
