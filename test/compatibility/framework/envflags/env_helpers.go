// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package envflags

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

const (
	CompatibilityTestsEnvZero string = "compatibility-tests-env0"
	CompatibilityTestsEnvOne  string = "compatibility-tests-env1"
	CompatibilityTestsEnvVal  string = "compatibility-tests-env-val"
)

// DefaultGetEnvConfigurationsCommand creates a GetEnvConfigurations Command with default input and output options
func DefaultGetEnvConfigurationsCommand(version core.RuntimeVersion, opts ...CfgEnvOptionArgs) *core.Command {
	args := &CfgEnvArgs{
		Envs: map[string]string{
			CompatibilityTestsEnvZero: CompatibilityTestsEnvVal,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetEnvConfigurationsInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}

	outputOpts := &GetEnvConfigurationsOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Envs:               args.Envs,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewGetEnvConfigurationsCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetEnvCommand creates a GetEnv Command with default input and output options
func DefaultGetEnvCommand(version core.RuntimeVersion, opts ...CfgEnvOptionArgs) *core.Command {
	args := &CfgEnvArgs{
		Key:   CompatibilityTestsEnvZero,
		Value: CompatibilityTestsEnvVal,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetEnvInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key: args.Key,
	}

	outputOpts := &GetEnvOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Value:              args.Value,
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewGetEnvCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultDeleteEnvCommand creates a DeleteEnv Command with default input and output options
func DefaultDeleteEnvCommand(version core.RuntimeVersion, opts ...CfgEnvOptionArgs) *core.Command {
	args := &CfgEnvArgs{
		Key: CompatibilityTestsEnvZero,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &DeleteEnvInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key: args.Key,
	}

	outputOpts := &DeleteEnvOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Error: args.Error,
	}

	cmd, err := NewDeleteEnvCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultSetEnvCommand creates a SetEnv Command with default input and output options
func DefaultSetEnvCommand(version core.RuntimeVersion, opts ...CfgEnvOptionArgs) *core.Command {
	args := &CfgEnvArgs{
		Key:   CompatibilityTestsEnvZero,
		Value: CompatibilityTestsEnvVal,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetEnvInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Key:   args.Key,
		Value: args.Value,
	}

	outputOpts := &SetEnvOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Error: args.Error,
	}

	cmd, err := NewSetEnvCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}
