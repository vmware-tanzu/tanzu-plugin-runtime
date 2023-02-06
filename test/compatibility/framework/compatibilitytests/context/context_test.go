// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package context_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/context"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version Context APIs compatibility tests", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("GetContext, SetContext, DeleteContext, GetCurrentContext, SetCurrentContext, RemoveCurrentContext methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, latest")

	// Setup Data
	var contextTestHelper context.Helper
	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
		contextTestHelper.SetUpDefaultData()
	})

	ginkgo.Context("using single context object on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime latest
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add GetClientConfig Commands on all supported runtime versions
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntime0280)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime v0.28.0
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add GetClientConfig Commands on all supported runtime versions
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext latest Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntimeLatest)

			// Add DeleteContext latest Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext latestCommand
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntimeLatestWithError)

			// Add DeleteContext latest Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntimeLatestWithError)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add DeleteContext v0.25.4 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using multiple context objects on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext on Runtime latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntime0280)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254WithError)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetContextTwoCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntime0280WithError)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0280WithError)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextTwoCmdForRuntime0280WithError).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetContextTwoCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add DeleteContext v0.25.4 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254WithError)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

	})
})
