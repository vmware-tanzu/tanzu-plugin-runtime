// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package envflags_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/envflags"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version Env Flags APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("GetEnv, GetEnvConfigurations, SetEnv, DeleteEnv methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")
	var multipleTestEnvs map[string]string
	ginkgo.BeforeEach(func() {
		multipleTestEnvs = map[string]string{
			envflags.CompatibilityTestsEnvZero: envflags.CompatibilityTestsEnvVal,
			envflags.CompatibilityTestsEnvOne:  envflags.CompatibilityTestsEnvVal,
		}
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using single env flag", func() {

		ginkgo.It("Run SetEnv latest - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.VersionLatest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v1.0.2 ", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.VersionLatest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.VersionLatest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v1.0.2 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version102))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0.2 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version102))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0.2 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version102))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version090))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v1.0.2", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version090))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version090))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0280))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.VersionLatest))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v1.0.2", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0280))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version102))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase().Add(envflags.DefaultSetEnvCommand(core.Version0280))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version090))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(map[string]string{}), envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using multiple env flags", func() {

		ginkgo.It("Run SetEnv latest - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.VersionLatest))
			testCase.Add(envflags.DefaultSetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.VersionLatest))
			testCase.Add(envflags.DefaultSetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv latest - DeleteEnv v1.0.2", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.VersionLatest))
			testCase.Add(envflags.DefaultSetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v1.0.2 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version102))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0.2 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version102))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v1.0.2 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version102))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v0.28.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version090))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version090))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.90.0 - DeleteEnv v1.0.2", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version090))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv latest Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv latest", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0280))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v1.0.2", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0280))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.28.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetEnv v0.28.0 - DeleteEnv v0.90.0", func() {
			// Build test case with commands

			// Add SetEnv Commands of Runtime v0.28.0
			testCase := core.NewTestCase()

			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0280))
			testCase.Add(envflags.DefaultSetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithValue(envflags.CompatibilityTestsEnvVal), envflags.WithStrictValidationStrategy()))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithEnvs(multipleTestEnvs), envflags.WithStrictValidationStrategy()))

			// Add DeleteEnv v0.90.0 Command
			testCase.Add(envflags.DefaultDeleteEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne)))

			// Add GetEnv latest, v0.90.0, v0.28.0 Commands
			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithStrictValidationStrategy()))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithStrictValidationStrategy()))

			testCase.Add(envflags.DefaultGetEnvCommand(core.VersionLatest, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version102, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version090, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))
			testCase.Add(envflags.DefaultGetEnvCommand(core.Version0280, envflags.WithKey(envflags.CompatibilityTestsEnvOne), envflags.WithError(common.ErrNotFound)))

			// Add GetEnvConfigurations v0.25.4, v0.11.6 Commands
			testCase.Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0254, envflags.WithStrictValidationStrategy())).Add(envflags.DefaultGetEnvConfigurationsCommand(core.Version0116, envflags.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})
	})
})
