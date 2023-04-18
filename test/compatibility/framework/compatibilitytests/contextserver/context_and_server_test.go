// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package contextserver_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/server"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
)

var _ = ginkgo.Describe("Combination Tests for Context and Server APIs", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("Get/Set/Delete Context, CurrentContext, Server and CurrentServer API methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, latest")

	// Setup Data
	var contextTestHelper context.Helper
	var serverTestHelper server.Helper
	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
		serverTestHelper.SetUpDefaultData()
		contextTestHelper.SetUpDefaultData()
	})

	ginkgo.Context("using single context and server", func() {

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.25.4", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.28.0", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.25.4", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.28.0", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.25.4", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.28.0", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using two different contexts and servers", func() {

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetServerTwoCmdForRuntime0254).Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntimeLatest).Add(contextTestHelper.SetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.SetCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetServerTwoCmdForRuntime0280).Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetContextTwoCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetContextTwoCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetServerTwoCmdForRuntime0254).Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntime0280).Add(contextTestHelper.SetContextTwoCmdForRuntime0280).Add(contextTestHelper.SetCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetServerTwoCmdForRuntime0280).Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetContextTwoCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntimeLatest).Add(serverTestHelper.SetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.SetCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetContextTwoCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0254).Add(serverTestHelper.SetServerTwoCmdForRuntime0254).Add(serverTestHelper.SetCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextCmdForRuntime0280WithError).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatestWithError).Add(contextTestHelper.GetContextTwoCmdForRuntime0280WithError).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatestWithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280WithError).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(contextTestHelper.SetContextCmdForRuntime0254).Add(contextTestHelper.SetContextTwoCmdForRuntime0254).Add(contextTestHelper.SetCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(serverTestHelper.SetServerCmdForRuntime0280).Add(serverTestHelper.SetServerTwoCmdForRuntime0280).Add(serverTestHelper.SetCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetContextCmdForRuntimeLatest).Add(contextTestHelper.GetContextCmdForRuntime0280).Add(contextTestHelper.GetContextCmdForRuntime0254)
			testCase.Add(contextTestHelper.GetContextTwoCmdForRuntimeLatest).Add(contextTestHelper.GetContextTwoCmdForRuntime0280).Add(contextTestHelper.GetContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetServerCmdForRuntimeLatest).Add(serverTestHelper.GetServerCmdForRuntime0280).Add(serverTestHelper.GetServerCmdForRuntime0254)
			testCase.Add(serverTestHelper.GetServerTwoCmdForRuntimeLatest).Add(serverTestHelper.GetServerTwoCmdForRuntime0280).Add(serverTestHelper.GetServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(contextTestHelper.GetCurrentContextCmdForRuntimeLatest).Add(contextTestHelper.GetCurrentContextCmdForRuntime0280).Add(contextTestHelper.GetCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(serverTestHelper.GetCurrentServerCmdForRuntimeLatest).Add(serverTestHelper.GetCurrentServerCmdForRuntime0280).Add(serverTestHelper.GetCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
	})

})
