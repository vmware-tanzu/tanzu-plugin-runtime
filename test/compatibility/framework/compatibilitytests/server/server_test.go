// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package server_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/server"
)

var _ = ginkgo.Describe("Cross-version Server APIs Compatibility Tests", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("Get/Set/Delete Server and CurrentServer API methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, v0.90.0, v1.0.2, latest")

	// Setup Data
	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})

	})

	ginkgo.Context("using single server", func() {

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest - DeleteServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer, RemoveCurrentServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer, RemoveCurrentServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.RemoveCurrentServerCommand())
			testCase.Add(server.DeleteServerCommand())

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer, RemoveCurrentServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer, RemoveCurrentServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.RemoveCurrentServerCommand())
			testCase.Add(server.DeleteServerCommand())

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer,RemoveCurrentServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.RemoveCurrentServerCommand())
			testCase.Add(server.DeleteServerCommand())

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer,RemoveCurrentServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer,RemoveCurrentServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using multiple servers", func() {

		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime latest - DeleteServer, RemoveCurrentServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand())
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand())

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer, RemoveCurrentServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.RemoveCurrentServerCommand())
			testCase.Add(server.DeleteServerCommand())

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer, RemoveCurrentServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v1.0.2 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer, RemoveCurrentServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer, RemoveCurrentServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.RemoveCurrentServerCommand())
			testCase.Add(server.DeleteServerCommand())

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.90.0 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.RemoveCurrentServerCommand())
			testCase.Add(server.DeleteServerCommand())

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090)))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer of Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.28.0 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0280)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0280)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer, RemoveCurrentServer of Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.25.4 - DeleteServer of Runtime v0.11.6", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0254)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.11.6 - DeleteServer of Runtime latest", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0116)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.11.6 - DeleteServer of Runtime v1.0.2", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0116)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.11.6 - DeleteServer of Runtime v0.90.0", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0116)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116)))

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

			testCase.Add(server.RemoveCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))

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

			executer.Execute(testCase)
		})
		ginkgo.It("Run two SetServer, SetCurrentServer of Runtime v0.11.6 - DeleteServer of Runtime v0.25.4 ", func() {
			testCase := core.NewTestCase()

			testCase.Add(server.SetServerCommand(server.WithRuntimeVersion(core.Version0116)))
			testCase.Add(server.SetServerCommand(server.WithServerName(common.CompatibilityTestTwo), server.WithRuntimeVersion(core.Version0116)))

			testCase.Add(server.SetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116)))

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

			testCase.Add(server.DeleteServerCommand(server.WithRuntimeVersion(core.Version0254)))

			testCase.Add(server.GetServerCommand(server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			testCase.Add(server.GetServerCommand(server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version102), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version090), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithServerName(common.CompatibilityTestTwo)))
			testCase.Add(server.GetServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithServerName(common.CompatibilityTestTwo)))

			testCase.Add(server.GetCurrentServerCommand(server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version102), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version090), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0280), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0254), server.WithError()))
			testCase.Add(server.GetCurrentServerCommand(server.WithRuntimeVersion(core.Version0116), server.WithError()))

			executer.Execute(testCase)
		})
	})

})
