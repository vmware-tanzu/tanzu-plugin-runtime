// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package envflags contains all the cross version api compatibility tests for context apis
package envflags

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/envflags"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// DefaultGetEnvConfigurationsCommand creates a GetEnvConfigurations Command with default input and output options
func DefaultGetEnvConfigurationsCommand(version core.RuntimeVersion, value map[string]string) *core.Command {
	cmd, err := envflags.NewGetEnvConfigurationsCommand(envflags.ConstructGetEnvConfigurationsInputOptions(version), envflags.ConstructGetEnvConfigurationsOutputOptions(version, value))
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetEnvCommand creates a GetEnv Command with default input and output options
func DefaultGetEnvCommand(version core.RuntimeVersion, key, value string) *core.Command {
	cmd, err := envflags.NewGetEnvCommand(envflags.ConstructGetEnvInputOptions(version, key), envflags.ConstructGetEnvOutputOptions(version, value))
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultGetEnvCommandWithOutputError creates a GetEnvConfigurations Command with error and default input and output options
func DefaultGetEnvCommandWithOutputError(version core.RuntimeVersion, key, errStr string) *core.Command {
	cmd, err := envflags.NewGetEnvCommand(envflags.ConstructGetEnvInputOptions(version, key), envflags.ConstructGetEnvOutputOptionsWithError(version, errStr))
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultDeleteEnvCommand creates a DeleteEnv Command with default input and output options
func DefaultDeleteEnvCommand(version core.RuntimeVersion, key string) *core.Command {
	cmd, err := envflags.NewDeleteEnvCommand(envflags.ConstructDeleteEnvInputOptions(version, key), nil)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultSetEnvCommand creates a SetEnv Command with default input and output options
func DefaultSetEnvCommand(version core.RuntimeVersion, key, value string) *core.Command {
	cmd, err := envflags.NewSetEnvCommand(envflags.ConstructSetEnvInputOptions(version, key, value), nil)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}
