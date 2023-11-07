// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package context_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/legacyclientconfig"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

var _ = ginkgo.Describe("Cross-version Context APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("GetContext, SetContext, DeleteContext, GetCurrentContext, SetCurrentContext, RemoveCurrentContext methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, v0.90.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using single context object on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

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

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version090, legacyclientconfig.WithDefaultContextAndServer(core.Version090)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0280, legacyclientconfig.WithDefaultContextAndServer(core.Version0280)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0254, legacyclientconfig.WithDefaultContextAndServer(core.Version0254)))
			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.Version0116, legacyclientconfig.WithDefaultContextAndServer(core.Version0116)))

			testCase.Add(legacyclientconfig.DefaultGetClientConfigCommand(core.VersionLatest, legacyclientconfig.WithDefaultContextAndServer(core.VersionLatest)))
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

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetCurrentContextCommand())

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0.2 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0.2 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0.2 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

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

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

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

	ginkgo.Context("using multiple context objects on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v1.0.2 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand())
			testCase.Add(context.SetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand())
			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v1.0.2 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand())

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0280)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand())
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand())
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(context.DeleteContextCommand())

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext v1.0.2 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version090)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.RemoveCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v0.90.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

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

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext v1.0.2 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

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

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0 latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, v0.90.0, latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254)))
			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254)))

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

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

			testCase.Add(context.GetContextCommand(context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version090), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithContextName(common.CompatibilityTestTwo), context.WithError()))
			testCase.Add(context.GetContextCommand(context.WithRuntimeVersion(core.Version0254), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetCurrentContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version090), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0280), context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version0254)))

			executer.Execute(testCase)
		})
	})
	ginkgo.Context("Using tanzu and k8s context types objects on supported Runtime API versions to validate mutual exclusion behavior", func() {
		ginkgo.It("Run SetContext, SetCurrentContext, GetActiveContext latest then SetContext,SetCurrentContext, GetCurrentContext ,GetActiveContext, v1.0.2 then DeleteContext latest then SetContext, SetCurrentContext, GetCurrentContext v1.0.2 then SetContext, SetCurrentContext, GetActiveContext latest then GetCurrentContextCommand v1.0.2", func() {
			testCase := core.NewTestCase()
			// When latest plugin version sets tanzu context as current and later old plugin API version sets k8s context as active,
			// the k8s context type and tanzu context types should be active which CLI would inspect and remove the tanzu context type from current contexts
			testCase.Add(context.SetContextCommand(context.WithContextType(types.ContextTypeTanzu)))
			testCase.Add(context.SetCurrentContextCommand())
			testCase.Add(context.GetActiveContextCommand(context.WithContextType(types.ContextTypeTanzu)))

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.GetActiveContextCommand(context.WithContextType(types.ContextTypeTanzu)))
			testCase.Add(context.GetActiveContextCommand(context.WithContextName(common.CompatibilityTestTwo)))

			// When old plugin API version sets k8s context as active and later if latest plugin version sets tanzu context as current,
			// the k8s context type should be removed from the current context list automatically
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102)))
			testCase.Add(context.DeleteContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.SetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithContextName(common.CompatibilityTestTwo)))

			testCase.Add(context.SetContextCommand(context.WithContextType(types.ContextTypeTanzu)))
			testCase.Add(context.SetCurrentContextCommand())
			testCase.Add(context.GetActiveContextCommand(context.WithContextType(types.ContextTypeTanzu)))
			testCase.Add(context.GetActiveContextCommand(context.WithError()))
			testCase.Add(context.GetCurrentContextCommand(context.WithRuntimeVersion(core.Version102), context.WithError()))

			executer.Execute(testCase)
		})
	})
})
