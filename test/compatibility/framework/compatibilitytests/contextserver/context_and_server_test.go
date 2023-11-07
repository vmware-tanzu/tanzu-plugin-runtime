// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package contextserver_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/server"
)

var _ = ginkgo.Describe("Combination Tests for Context - Server APIs", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("Get/Set/Delete Context, CurrentContext, Server and CurrentServer API methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, latest")

	// Setup Data
	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using single context- Server", func() {

		ginkgo.It("Set Context@latest - Set Server@v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v1.0.2 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v1.0.2 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v1.0.2 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v1.0.2 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.90.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.28.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.25.4 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

			addTestCasesToVerifyContextAndServer(testCase)

			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using two different contexts- Servers", func() {

		ginkgo.It("Set Context@v0.90.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.90.0 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@latest - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@latest - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.28.0 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.28.0 - Set Server@v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})

		ginkgo.It("Set Context@v0.25.4 - Set Server@latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
		ginkgo.It("Set Context@v0.25.4 - Set Server@v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

			addTestCasesToVerifyTwoContextsAndServers(testCase)

			executer.Execute(testCase)
		})
	})

})

func addTestCasesToVerifyContextAndServer(testCase *core.TestCase) {
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultContextAndServer(core.Version102)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.Version090)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
	testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

	testCase.Add(context.GetContextCommand())
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

	testCase.Add(context.GetCurrentContextCommand())
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

	testCase.Add(server.GetServerCommand())
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116)))

	testCase.Add(server.GetCurrentServerCommand())
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116)))
}

func addTestCasesToVerifyTwoContextsAndServers(testCase *core.TestCase) {
	testCase.Add(context.GetContextCommand())
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

	testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
	testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

	testCase.Add(context.GetCurrentContextCommand())
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
	testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

	testCase.Add(server.GetServerCommand())
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116)))

	testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
	testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

	testCase.Add(server.GetCurrentServerCommand())
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))
	testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116)))
}
