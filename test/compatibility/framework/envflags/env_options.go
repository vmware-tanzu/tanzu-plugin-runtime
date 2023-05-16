// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package envflags

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// GetEnvConfigurationsInputOptions used to generate GetEnvConfigurations command
type GetEnvConfigurationsInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetEnvConfigurationsOutputOptions used to generate GetEnvConfigurations command
type GetEnvConfigurationsOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	Envs                    map[string]string       // For specific version options look into EnvOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// GetEnvInputOptions used to generate GetEnv command
type GetEnvInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Key                     string // required
}

// GetEnvOutputOptions used to generate GetEnv command
type GetEnvOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	Value                   string                  // For specific version options look into EnvOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetEnvInputOptions used to generate SetEnv command
type SetEnvInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Key                     string // required
	Value                   string // required
}

// SetEnvOutputOptions used to generate SetEnv command
type SetEnvOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// DeleteEnvInputOptions used to generate DeleteEnv command
type DeleteEnvInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Key                     string // required

}

// DeleteEnvOutputOptions used to generate DeleteEnv command
type DeleteEnvOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// CfgEnvArgs arguments struct that are used to create Env API commands
type CfgEnvArgs struct {
	*core.RuntimeAPIVersion
	Key                string
	Value              string
	Envs               map[string]string
	ValidationStrategy core.ValidationStrategy
	Error              string
}

type CfgEnvOptionArgs func(*CfgEnvArgs)

func WithRuntimeAPIVersion(v *core.RuntimeAPIVersion) CfgEnvOptionArgs {
	return func(c *CfgEnvArgs) {
		c.RuntimeAPIVersion = v
	}
}

func WithEnvs(v map[string]string) CfgEnvOptionArgs {
	return func(c *CfgEnvArgs) {
		c.Envs = v
	}
}

func WithKey(k string) CfgEnvOptionArgs {
	return func(c *CfgEnvArgs) {
		c.Key = k
	}
}

func WithValue(v string) CfgEnvOptionArgs {
	return func(c *CfgEnvArgs) {
		c.Value = v
	}
}

func WithValidationStrategy(vs core.ValidationStrategy) CfgEnvOptionArgs {
	return func(c *CfgEnvArgs) {
		c.ValidationStrategy = vs
	}
}

func WithError(e string) CfgEnvOptionArgs {
	return func(c *CfgEnvArgs) {
		c.Error = e
	}
}
