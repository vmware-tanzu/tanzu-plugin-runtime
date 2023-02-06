// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package server_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/server"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version Server APIs Compatibility Tests for supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("Get/Set/Delete Server and CurrentServer API methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")

	// Setup Data
	var serverTestHelper server.Helper
	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})

		serverTestHelper.SetUpDefaultData()
	})

	ginkgo.Context("using single server", func() {

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime Latest version
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(serverTestHelper.RemoveCurrentServerCmdForRuntime0280)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetServerCmdForRuntime0280WithError).Add(serverTestHelper.GetServerCmdForRuntime0254WithError).Add(serverTestHelper.GetServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime v0.11.6 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime Latest version
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.11.6 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetServerCmdForRuntime0280WithError).Add(serverTestHelper.GetServerCmdForRuntime0254WithError).Add(serverTestHelper.GetServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetServerCmdForRuntime0280WithError).Add(serverTestHelper.GetServerCmdForRuntime0254WithError).Add(serverTestHelper.GetServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.25.4
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(serverTestHelper.RemoveCurrentServerCmdForRuntime0280WithError)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0280WithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime 0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0116).Add(serverTestHelper.SetCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetServerCmdForRuntime0280WithError).Add(serverTestHelper.GetServerCmdForRuntime0254WithError).Add(serverTestHelper.GetServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0116).Add(serverTestHelper.SetCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add DeleteServer latest Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntimeLatestWithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using multiple servers", func() {

		ginkgo.It("Run two SetServer of Runtime latest then SetCurrentServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime Latest
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetServerTwoCmdForRuntimeLatest)

			// Add SetCurrentServer Command of Runtime Latest
			testCase.Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(serverTestHelper.RemoveCurrentServerCmdForRuntime0280)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetServerCmdForRuntime0280WithError).Add(serverTestHelper.GetServerCmdForRuntime0254WithError).Add(serverTestHelper.GetServerCmdForRuntime0116WithError)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.25.4 then SetCurrentServer of Runtime v0.25.4 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime v0.25.4
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetServerTwoCmdForRuntime0254)

			// Add SetCurrentServer Command of Runtime v0.25.4
			testCase.Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(serverTestHelper.RemoveCurrentServerCmdForRuntime0280WithError)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0280WithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.28.0 then SetCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {

			// Add two SetServer Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetServerTwoCmdForRuntime0280)

			// Add SetCurrentServer Command of Runtime v0.28.0
			testCase.Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetServerCmdForRuntime0280WithError).Add(serverTestHelper.GetServerCmdForRuntime0254WithError).Add(serverTestHelper.GetServerCmdForRuntime0116WithError)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.11.6 then SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0116).Add(serverTestHelper.SetServerTwoCmdForRuntime0116)

			// Add SetCurrentServer Command of Runtime v0.11.6
			testCase.Add(serverTestHelper.SetCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add DeleteServer latest Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntimeLatestWithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.11.6 then SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(serverTestHelper.SetServerCmdForRuntime0116).Add(serverTestHelper.SetServerTwoCmdForRuntime0116)

			// Add SetCurrentServer Command of Runtime v0.11.6
			testCase.Add(serverTestHelper.SetCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254).Add(serverTestHelper.GetServerCmdForRuntime0116)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(serverTestHelper.DeleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetServerCmdForRuntime0280WithError).Add(serverTestHelper.GetServerCmdForRuntime0254WithError).Add(serverTestHelper.GetServerCmdForRuntime0116WithError)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254).Add(serverTestHelper.GetServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatestWithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254WithError).Add(serverTestHelper.GetCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})
	})
})
