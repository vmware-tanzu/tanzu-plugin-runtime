// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package legacyclientconfig_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"
)

var _ = ginkgo.Describe("Cross-version Legacy Client Config APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("GetClientConfig, StoreClientConfig methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("involving context", func() {

		ginkgo.It("Run StoreClientConfig latest - DeleteContext, RemoveCurrentContext v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig latest - DeleteContext, RemoveCurrentContext v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig latest - DeleteContext, RemoveCurrentContext v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v1.0.2 - DeleteContext, RemoveCurrentContext v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultContextAndServer(core.Version102)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v1.0.2 - DeleteContext, RemoveCurrentContext v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultContextAndServer(core.Version102)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v1.0.2 - DeleteContext, RemoveCurrentContext latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version102, legacyclientconfig.WithDefaultContextAndServer(core.Version102)))

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

			testCase.Add(context.RemoveCurrentContextCommand())
			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.90.0 - DeleteContext, RemoveCurrentContext latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.Version090)))

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

			testCase.Add(context.RemoveCurrentContextCommand())
			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.90.0 - DeleteContext, RemoveCurrentContext v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.Version090)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.90.0 - DeleteContext, RemoveCurrentContext v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.Version090)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.28.0 - DeleteContext v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand())
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetCurrentContextCommand())
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteContext, RemoveCurrentContext latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteContext, RemoveCurrentContext v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			executer.Execute(testCase)
		})
		ginkgo.It("Run StoreClientConfig v0.25.4 - DeleteContext, RemoveCurrentContext v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(legacyclientconfig.DefaultStoreClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			executer.Execute(testCase)
		})

	})

})
