// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package featureflags_test

import (
	"github.com/onsi/ginkgo/v2"

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

		ginkgo.It("Run SetFeature latest then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest then DeleteFeature v0.28.0 then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.28.0 then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest then DeleteFeature latest then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(featureflags.DefaultSetFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using default multiple feature flag", func() {

		ginkgo.It("Run SetFeature latest then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest then DeleteFeature v0.28.0 then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime Latest
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add DeleteFeature v0.28.0 Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.Version0280))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetFeature v0.28.0 then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest then DeleteFeature latest then IsFeatureEnabled v0.11.6, v0.25.4, v0.28.0, latest", func() {
			// Build test case with commands

			// Add SetFeature Commands of Runtime v0.28.0
			testCase := core.NewTestCase()
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultSetFeatureCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Add DeleteFeature latest Command
			testCase.Add(featureflags.DefaultDeleteFeatureCommand(core.VersionLatest))

			// Add IsFeatureEnabled latest, v0.28.0, v0.25.4, v0.11.6 Commands
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithError(common.ErrNotFound)))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.VersionLatest, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0280, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0254, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))
			testCase.Add(featureflags.DefaultIsFeatureEnabledCommand(core.Version0116, featureflags.WithKey(featureflags.CompatibilityTestsPluginKey0), featureflags.WithValue("false")))

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	// TODO: More tests using GetClientConfig StoreClientConfig will be added
})
