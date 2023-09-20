// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package legacyclientconfig_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
)

var _ = ginkgo.Describe("Cross-version Legacy Client Config APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("GetClientConfig, StoreClientConfig methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")

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

	ginkgo.Context("involving context", func() {

		ginkgo.It("Run StoreClientConfig latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime latest
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntime0280)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0280)

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime latest
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.90.0 Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntime090)

			// Add DeleteContext v0.90.0 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime090)

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime latest
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.Version090)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext latest Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntimeLatest)

			// Add DeleteContext latest Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntimeLatest)

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime latest
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.Version090)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntime0280)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0280)

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add DeleteContext v0.25.4 Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime0254)

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext latest Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntimeLatestWithError)

			// Add DeleteContext latest Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntimeLatestWithError)

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext latest Command
			testCase.Add(contextTestHelper.RemoveCurrentContextCmdForRuntime090WithError)

			// Add DeleteContext latest Command
			testCase.Add(contextTestHelper.DeleteContextCmdForRuntime090WithError)

			// Add GetContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.90.0, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime090WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError)
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

	})

})
