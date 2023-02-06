// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package featureflags contains all the cross version api compatibility tests for context apis
package featureflags

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/featureflags"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
)

// DefaultIsFeatureEnabledCommand creates a IsFeatureEnabled Command with default input and output options
func DefaultIsFeatureEnabledCommand(version core.RuntimeVersion, key string, expectedFeatureEnabled bool) *core.Command {
	cmd, err := featureflags.NewIsFeatureEnabledCommand(featureflags.ConstructIsFeatureEnabledInputOptions(version, common.CompatibilityTestsPlugin, key), featureflags.ConstructIsFeatureEnabledOutputOptions(version, expectedFeatureEnabled))
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultIsFeatureEnabledCommandWithOutputError creates a DefaultIsFeatureEnabledCommandWithOutputError Command with default input options and output with error options
func DefaultIsFeatureEnabledCommandWithOutputError(version core.RuntimeVersion, key string) *core.Command {
	cmd, err := featureflags.NewIsFeatureEnabledCommand(featureflags.ConstructIsFeatureEnabledInputOptions(version, common.CompatibilityTestsPlugin, key), featureflags.ConstructIsFeatureEnabledOutputOptionsWithError(version, "not found"))
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultDeleteFeatureCommand creates a DeleteFeature Command with default input and output options
func DefaultDeleteFeatureCommand(version core.RuntimeVersion, key string) *core.Command {
	cmd, err := featureflags.NewDeleteFeatureCommand(featureflags.ConstructDeleteFeatureInputOptions(version, common.CompatibilityTestsPlugin, key), nil)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultSetFeatureCommand creates a DeleteFeature Command with default input and output options
func DefaultSetFeatureCommand(version core.RuntimeVersion, key, value string) *core.Command {
	cmd, err := featureflags.NewSetFeatureCommand(featureflags.ConstructSetFeatureInputOptions(version, common.CompatibilityTestsPlugin, key, value), nil)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}
