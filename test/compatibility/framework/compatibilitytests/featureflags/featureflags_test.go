// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package featureflags_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/featureflags"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version Feature Flags APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("IsFeatureEnabled, SetFeature, DeleteFeature methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using default single feature flag", func() {

		ginkgo.It("Run SetFeature latest - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version090))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v1.0.2", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version102))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v1.0.2 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version102))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0.2 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version102))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version090))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0.2 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version102))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version090))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version090))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v1.0.2", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version090))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version102))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0280))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v1.0.2", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0280))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version102))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0280))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version090))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature latest", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature v1.0.2", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version102))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version090))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			testCase := core.NewTestCase()

			// Add StoreClientConfig Command for Runtime v0.25.4
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultFeatureFlags()))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultFeatureFlags()))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using default multiple feature flag", func() {

		ginkgo.It("Run SetFeature latest - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version090))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature latest - DeleteFeature v1.0.2", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version102))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v1.0.2 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version102))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0.2 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version102))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version090))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v1.0.2 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version102))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v0.28.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version090))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature v1.0.2", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version090))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version102))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.90.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version090))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig latest Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0280))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig v0.28.0 Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v1.0.2", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0280))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig v0.28.0 Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version102))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetFeature v0.28.0 - DeleteFeature v0.90.0", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0280))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add StoreClientConfig v0.28.0 Commands
			features := map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey:  "true",
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add DeleteFeature v0.90.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version090))

			// Add IsFeatureEnabled latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version102, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version090, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false"), featureflags.WithStrictValidationStrategy()))

			// Add GetClientConfig latest, v0.90.0, v0.28.0, v0.25.4, v0.11.6 Commands
			features = map[string]types.FeatureMap{
				featureflags.CompatibilityTestsPlugin: map[string]string{
					featureflags.CompatibilityTestsPluginKey0: "false",
					featureflags.CompatibilityTestsPluginKey1: "true",
				},
			}
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithFeatureFlags(features)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithFeatureFlags(features)))

			// Run all the commands
			executer.Execute(testCase)
		})

	})

})
