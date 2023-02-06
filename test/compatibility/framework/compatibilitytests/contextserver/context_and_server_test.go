// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package contextserver_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/server"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
)

var _ = ginkgo.Describe("Combination Tests for Context and Server APIs", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("Get/Set/Delete Context, CurrentContext, Server and CurrentServer API methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	// Input and Output Options for Server APIs
	var setServerInputOptionsForRuntime0254 *framework.SetServerInputOptions
	var setServerInputOptionsForRuntime0280 *framework.SetServerInputOptions
	var setServerInputOptionsForRuntimeLatest *framework.SetServerInputOptions

	var setServerTwoInputOptionsForRuntime0254 *framework.SetServerInputOptions
	var setServerTwoInputOptionsForRuntime0280 *framework.SetServerInputOptions
	var setServerTwoInputOptionsForRuntimeLatest *framework.SetServerInputOptions

	var setCurrentServerInputOptionsForRuntime0254 *framework.SetCurrentServerInputOptions
	var setCurrentServerInputOptionsForRuntime0280 *framework.SetCurrentServerInputOptions
	var setCurrentServerInputOptionsForRuntimeLatest *framework.SetCurrentServerInputOptions

	var getServerInputOptionsForRuntimeLatest *framework.GetServerInputOptions
	var getServerInputOptionsForRuntime0280 *framework.GetServerInputOptions
	var getServerInputOptionsForRuntime0254 *framework.GetServerInputOptions

	var getServerTwoInputOptionsForRuntimeLatest *framework.GetServerInputOptions
	var getServerTwoInputOptionsForRuntime0280 *framework.GetServerInputOptions
	var getServerTwoInputOptionsForRuntime0254 *framework.GetServerInputOptions

	var getServerOutputOptionsForRuntime0254 *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntime0280 *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntimeLatest *framework.GetServerOutputOptions

	var getServerTwoOutputOptionsForRuntime0254 *framework.GetServerOutputOptions
	var getServerTwoOutputOptionsForRuntime0280 *framework.GetServerOutputOptions
	var getServerTwoOutputOptionsForRuntimeLatest *framework.GetServerOutputOptions

	var getCurrentServerInputOptionsForRuntime0254 *framework.GetCurrentServerInputOptions
	var getCurrentServerInputOptionsForRuntime0280 *framework.GetCurrentServerInputOptions
	var getCurrentServerInputOptionsForRuntimeLatest *framework.GetCurrentServerInputOptions

	var getCurrentServerOutputOptionsForRuntime0254 *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntime0280 *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntimeLatest *framework.GetCurrentServerOutputOptions

	// Server API Commands
	var setServerCmdForRuntimeLatest *core.Command
	var setServerCmdForRuntime0280 *core.Command
	var setServerCmdForRuntime0254 *core.Command

	var setServerTwoCmdForRuntimeLatest *core.Command
	var setServerTwoCmdForRuntime0280 *core.Command
	var setServerTwoCmdForRuntime0254 *core.Command

	var setCurrentServerCmdForRuntime0254 *core.Command
	var setCurrentServerCmdForRuntime0280 *core.Command
	var setCurrentServerCmdForRuntimeLatest *core.Command

	var getServerCmdForRuntimeLatest *core.Command
	var getServerCmdForRuntime0280 *core.Command
	var getServerCmdForRuntime0254 *core.Command

	var getServerTwoCmdForRuntimeLatest *core.Command
	var getServerTwoCmdForRuntime0280 *core.Command
	var getServerTwoCmdForRuntime0254 *core.Command

	var getCurrentServerCmdForRuntimeLatest *core.Command
	var getCurrentServerCmdForRuntime0280 *core.Command
	var getCurrentServerCmdForRuntime0254 *core.Command

	var err error

	// Input and Output Options for Context APIs

	// SetContext Input Options
	var setContextInputOptionsForRuntime0254 *framework.SetContextInputOptions
	var setContextInputOptionsForRuntime0280 *framework.SetContextInputOptions
	var setContextInputOptionsForRuntimeLatest *framework.SetContextInputOptions
	var setContextTwoInputOptionsForRuntime0254 *framework.SetContextInputOptions
	var setContextTwoInputOptionsForRuntime0280 *framework.SetContextInputOptions
	var setContextTwoInputOptionsForRuntimeLatest *framework.SetContextInputOptions

	// SetCurrentContext Input Options
	var setCurrentContextInputOptionsForRuntime0254 *framework.SetCurrentContextInputOptions
	var setCurrentContextInputOptionsForRuntime0280 *framework.SetCurrentContextInputOptions
	var setCurrentContextInputOptionsForRuntimeLatest *framework.SetCurrentContextInputOptions

	// GetContext Input Options
	var getContextInputOptionsForRuntimeLatest *framework.GetContextInputOptions
	var getContextInputOptionsForRuntime0280 *framework.GetContextInputOptions
	var getContextInputOptionsForRuntime0254 *framework.GetContextInputOptions
	var getContextTwoInputOptionsForRuntimeLatest *framework.GetContextInputOptions
	var getContextTwoInputOptionsForRuntime0280 *framework.GetContextInputOptions
	var getContextTwoInputOptionsForRuntime0254 *framework.GetContextInputOptions

	// GetContext Output Options
	var getContextOutputOptionsForRuntime0254 *framework.GetContextOutputOptions
	var getContextOutputOptionsForRuntime0280 *framework.GetContextOutputOptions
	var getContextOutputOptionsForRuntimeLatest *framework.GetContextOutputOptions
	var getContextTwoOutputOptionsForRuntime0254 *framework.GetContextOutputOptions
	var getContextTwoOutputOptionsForRuntime0280 *framework.GetContextOutputOptions
	var getContextTwoOutputOptionsForRuntimeLatest *framework.GetContextOutputOptions

	// GetContext Output Options with expected error
	var getContextOutputOptionsForRuntimeLatestWithError *framework.GetContextOutputOptions
	var getContextOutputOptionsForRuntime0280WithError *framework.GetContextOutputOptions
	var getContextTwoOutputOptionsForRuntimeLatestWithError *framework.GetContextOutputOptions
	var getContextTwoOutputOptionsForRuntime0280WithError *framework.GetContextOutputOptions

	// GetCurrentContext Input Options
	var getCurrentContextInputOptionsForRuntime0254 *framework.GetCurrentContextInputOptions
	var getCurrentContextInputOptionsForRuntime0280 *framework.GetCurrentContextInputOptions
	var getCurrentContextInputOptionsForRuntimeLatest *framework.GetCurrentContextInputOptions

	// GetCurrentContext Output Options
	var getCurrentContextOutputOptionsForRuntime0254 *framework.GetCurrentContextOutputOptions
	var getCurrentContextOutputOptionsForRuntime0280 *framework.GetCurrentContextOutputOptions
	var getCurrentContextOutputOptionsForRuntimeLatest *framework.GetCurrentContextOutputOptions

	// GetCurrentContext Output Options with expected error
	var getCurrentContextOutputOptionsForRuntimeLatestWithError *framework.GetCurrentContextOutputOptions
	var getCurrentContextOutputOptionsForRuntime0280WithError *framework.GetCurrentContextOutputOptions

	// Context API Commands
	// SetContext API Commands
	var setContextCmdForRuntimeLatest *core.Command
	var setContextCmdForRuntime0280 *core.Command
	var setContextCmdForRuntime0254 *core.Command

	var setContextTwoCmdForRuntimeLatest *core.Command
	var setContextTwoCmdForRuntime0280 *core.Command
	var setContextTwoCmdForRuntime0254 *core.Command

	// SetCurrentContext API Commands
	var setCurrentContextCmdForRuntime0254 *core.Command
	var setCurrentContextCmdForRuntime0280 *core.Command
	var setCurrentContextCmdForRuntimeLatest *core.Command

	// GetContext API Commands
	var getContextCmdForRuntimeLatest *core.Command
	var getContextCmdForRuntime0280 *core.Command
	var getContextCmdForRuntime0254 *core.Command

	var getContextTwoCmdForRuntimeLatest *core.Command
	var getContextTwoCmdForRuntime0280 *core.Command
	var getContextTwoCmdForRuntime0254 *core.Command

	var getContextCmdForRuntimeLatestWithError *core.Command
	var getContextCmdForRuntime0280WithError *core.Command
	var getContextTwoCmdForRuntimeLatestWithError *core.Command
	var getContextTwoCmdForRuntime0280WithError *core.Command

	// GetCurrentContext API Commands
	var getCurrentContextCmdForRuntimeLatest *core.Command
	var getCurrentContextCmdForRuntime0280 *core.Command
	var getCurrentContextCmdForRuntime0254 *core.Command

	var getCurrentContextCmdForRuntimeLatestWithError *core.Command
	var getCurrentContextCmdForRuntime0280WithError *core.Command

	ginkgo.BeforeEach(func() {
		ginkgo.By("Setup Input and Output Options for Contexts APIs")
		// Input and Output Parameters for SetContext
		setContextInputOptionsForRuntimeLatest = context.DefaultSetContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		setContextInputOptionsForRuntime0280 = context.DefaultSetContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
		setContextInputOptionsForRuntime0254 = context.DefaultSetContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

		setContextTwoInputOptionsForRuntimeLatest = context.DefaultSetContextInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		setContextTwoInputOptionsForRuntime0280 = context.DefaultSetContextInputOptions(core.Version0280, common.CtxCompatibilityTwo)
		setContextTwoInputOptionsForRuntime0254 = context.DefaultSetContextInputOptions(core.Version0254, common.CtxCompatibilityTwo)

		// Input and Output Parameters for SetCurrentContext
		setCurrentContextInputOptionsForRuntimeLatest = context.DefaultSetCurrentContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		setCurrentContextInputOptionsForRuntime0280 = context.DefaultSetCurrentContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
		setCurrentContextInputOptionsForRuntime0254 = context.DefaultSetCurrentContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

		// Input and Output Parameters for GetCurrentContext
		getCurrentContextInputOptionsForRuntimeLatest = context.DefaultGetCurrentContextInputOptions(core.VersionLatest)
		getCurrentContextInputOptionsForRuntime0280 = context.DefaultGetCurrentContextInputOptions(core.Version0280)
		getCurrentContextInputOptionsForRuntime0254 = context.DefaultGetCurrentContextInputOptions(core.Version0254)

		getCurrentContextOutputOptionsForRuntime0280 = context.DefaultGetCurrentContextOutputOptions(core.Version0280, common.CtxCompatibilityOne)
		getCurrentContextOutputOptionsForRuntime0254 = context.DefaultGetCurrentContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)
		getCurrentContextOutputOptionsForRuntimeLatest = context.DefaultGetCurrentContextOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		getCurrentContextOutputOptionsForRuntimeLatestWithError = context.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)
		getCurrentContextOutputOptionsForRuntime0280WithError = context.DefaultGetCurrentContextOutputOptionsWithError(core.Version0280)

		// Input and Output params for GetContext
		getContextInputOptionsForRuntimeLatest = context.DefaultGetContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		getContextInputOptionsForRuntime0280 = context.DefaultGetContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
		getContextInputOptionsForRuntime0254 = context.DefaultGetContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

		getContextTwoInputOptionsForRuntimeLatest = context.DefaultGetContextInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		getContextTwoInputOptionsForRuntime0280 = context.DefaultGetContextInputOptions(core.Version0280, common.CtxCompatibilityTwo)
		getContextTwoInputOptionsForRuntime0254 = context.DefaultGetContextInputOptions(core.Version0254, common.CtxCompatibilityTwo)

		getContextOutputOptionsForRuntime0280 = context.DefaultGetContextOutputOptions(core.Version0280, common.CtxCompatibilityOne)
		getContextOutputOptionsForRuntime0254 = context.DefaultGetContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)
		getContextOutputOptionsForRuntimeLatest = context.DefaultGetContextOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		getContextTwoOutputOptionsForRuntimeLatest = context.DefaultGetContextOutputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		getContextTwoOutputOptionsForRuntime0280 = context.DefaultGetContextOutputOptions(core.Version0280, common.CtxCompatibilityTwo)
		getContextTwoOutputOptionsForRuntime0254 = context.DefaultGetContextOutputOptions(core.Version0254, common.CtxCompatibilityTwo)

		getContextOutputOptionsForRuntimeLatestWithError = context.DefaultGetContextOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)
		getContextOutputOptionsForRuntime0280WithError = context.DefaultGetContextOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

		getContextTwoOutputOptionsForRuntimeLatestWithError = context.DefaultGetContextOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityTwo)
		getContextTwoOutputOptionsForRuntime0280WithError = context.DefaultGetContextOutputOptionsWithError(core.Version0280, common.CtxCompatibilityTwo)

		ginkgo.By("Setup Context API commands")

		// Create SetContext Commands with input and output options
		setContextCmdForRuntimeLatest, err = framework.NewSetContextCommand(setContextInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setContextCmdForRuntime0254, err = framework.NewSetContextCommand(setContextInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setContextCmdForRuntime0280, err = framework.NewSetContextCommand(setContextInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())

		setContextTwoCmdForRuntimeLatest, err = framework.NewSetContextCommand(setContextTwoInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setContextTwoCmdForRuntime0254, err = framework.NewSetContextCommand(setContextTwoInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setContextTwoCmdForRuntime0280, err = framework.NewSetContextCommand(setContextTwoInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())

		// Create SetCurrentContext Commands with input and output options
		setCurrentContextCmdForRuntimeLatest, err = framework.NewSetCurrentContextCommand(setCurrentContextInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setCurrentContextCmdForRuntime0280, err = framework.NewSetCurrentContextCommand(setCurrentContextInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setCurrentContextCmdForRuntime0254, err = framework.NewSetCurrentContextCommand(setCurrentContextInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())

		// Create GetContext Commands with input and output options
		getContextCmdForRuntimeLatest, err = framework.NewGetContextCommand(getContextInputOptionsForRuntimeLatest, getContextOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getContextCmdForRuntime0280, err = framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getContextCmdForRuntime0254, err = framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())

		getContextTwoCmdForRuntimeLatest, err = framework.NewGetContextCommand(getContextTwoInputOptionsForRuntimeLatest, getContextTwoOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getContextTwoCmdForRuntime0280, err = framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0280, getContextTwoOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getContextTwoCmdForRuntime0254, err = framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0254, getContextTwoOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())

		getContextCmdForRuntimeLatestWithError, err = framework.NewGetContextCommand(getContextInputOptionsForRuntimeLatest, getContextOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())
		getContextCmdForRuntime0280WithError, err = framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
		getContextTwoCmdForRuntimeLatestWithError, err = framework.NewGetContextCommand(getContextTwoInputOptionsForRuntimeLatest, getContextTwoOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())
		getContextTwoCmdForRuntime0280WithError, err = framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0280, getContextTwoOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())

		// Create GetCurrentContext Commands with input and output options
		getCurrentContextCmdForRuntimeLatest, err = framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntimeLatest, getCurrentContextOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentContextCmdForRuntime0280, err = framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentContextCmdForRuntime0254, err = framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())

		getCurrentContextCmdForRuntimeLatestWithError, err = framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntimeLatest, getCurrentContextOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentContextCmdForRuntime0280WithError, err = framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.BeforeEach(func() {
		ginkgo.By("Setup Input and Output Options for Servers APIs")
		// Input and Output Parameters for SetServer
		setServerInputOptionsForRuntimeLatest = server.DefaultSetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		setServerInputOptionsForRuntime0280 = server.DefaultSetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
		setServerInputOptionsForRuntime0254 = server.DefaultSetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

		setServerTwoInputOptionsForRuntimeLatest = server.DefaultSetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		setServerTwoInputOptionsForRuntime0280 = server.DefaultSetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)
		setServerTwoInputOptionsForRuntime0254 = server.DefaultSetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)

		// Input and Output Parameters for SetCurrentServer
		setCurrentServerInputOptionsForRuntimeLatest = server.DefaultSetCurrentServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		setCurrentServerInputOptionsForRuntime0280 = server.DefaultSetCurrentServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
		setCurrentServerInputOptionsForRuntime0254 = server.DefaultSetCurrentServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

		// Input and Output Parameters for GetCurrentServer
		getCurrentServerInputOptionsForRuntimeLatest = server.DefaultGetCurrentServerInputOptions(core.VersionLatest)
		getCurrentServerInputOptionsForRuntime0280 = server.DefaultGetCurrentServerInputOptions(core.Version0280)
		getCurrentServerInputOptionsForRuntime0254 = server.DefaultGetCurrentServerInputOptions(core.Version0254)

		getCurrentServerOutputOptionsForRuntime0280 = server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
		getCurrentServerOutputOptionsForRuntime0254 = server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
		getCurrentServerOutputOptionsForRuntimeLatest = server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		// Input and Output params for GetServer
		getServerInputOptionsForRuntimeLatest = server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		getServerInputOptionsForRuntime0280 = server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
		getServerInputOptionsForRuntime0254 = server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

		getServerTwoInputOptionsForRuntimeLatest = server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		getServerTwoInputOptionsForRuntime0280 = server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)
		getServerTwoInputOptionsForRuntime0254 = server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)

		getServerTwoOutputOptionsForRuntimeLatest = server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		getServerTwoOutputOptionsForRuntime0280 = server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityTwo)
		getServerTwoOutputOptionsForRuntime0254 = server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityTwo)

		getServerOutputOptionsForRuntime0280 = server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntime0254 = server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntimeLatest = server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		ginkgo.By("Setup Server API commands")

		// Create SetServer Commands with input and output options
		setServerCmdForRuntimeLatest, err = framework.NewSetServerCommand(setServerInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerCmdForRuntime0254, err = framework.NewSetServerCommand(setServerInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerCmdForRuntime0280, err = framework.NewSetServerCommand(setServerInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())

		setServerTwoCmdForRuntimeLatest, err = framework.NewSetServerCommand(setServerTwoInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerTwoCmdForRuntime0254, err = framework.NewSetServerCommand(setServerTwoInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerTwoCmdForRuntime0280, err = framework.NewSetServerCommand(setServerTwoInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())

		// Create SetCurrentServer Commands with input and output options
		setCurrentServerCmdForRuntimeLatest, err = framework.NewSetCurrentServerCommand(setCurrentServerInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setCurrentServerCmdForRuntime0280, err = framework.NewSetCurrentServerCommand(setCurrentServerInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setCurrentServerCmdForRuntime0254, err = framework.NewSetCurrentServerCommand(setCurrentServerInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())

		// Create GetServer Commands with input and output options
		getServerCmdForRuntimeLatest, err = framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0280, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0254, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())

		getServerTwoCmdForRuntimeLatest, err = framework.NewGetServerCommand(getServerTwoInputOptionsForRuntimeLatest, getServerTwoOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getServerTwoCmdForRuntime0280, err = framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0280, getServerTwoOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getServerTwoCmdForRuntime0254, err = framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0254, getServerTwoOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())

		// Create GetCurrentServer Commands with input and output options
		getCurrentServerCmdForRuntimeLatest, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0280, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0254, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.Context("using single context and server", func() {

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.25.4", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0254).Add(setCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.28.0", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0280).Add(setCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.25.4", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0254).Add(setCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.28.0", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0280).Add(setCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0254).Add(setCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.25.4", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0254).Add(setCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0254).Add(setCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.28.0", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0254).Add(setCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0280).Add(setCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using two different contexts and servers", func() {

		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntimeLatest).Add(setContextTwoCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntimeLatest).Add(setServerTwoCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntimeLatest).Add(setContextTwoCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0254).Add(setServerTwoCmdForRuntime0254).Add(setCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime Latest and Set Server with Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntimeLatest).Add(setContextTwoCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0280).Add(setServerTwoCmdForRuntime0280).Add(setCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntime0280).Add(setContextTwoCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntimeLatest).Add(setServerTwoCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntime0280).Add(setContextTwoCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0254).Add(setServerTwoCmdForRuntime0254).Add(setCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.28.0 and Set Server with Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntime0280).Add(setContextTwoCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0280).Add(setServerTwoCmdForRuntime0280).Add(setCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime latest", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntime0254).Add(setContextTwoCmdForRuntime0254).Add(setCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntimeLatest).Add(setServerTwoCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.25.4", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntime0254).Add(setContextTwoCmdForRuntime0254).Add(setCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0254).Add(setServerTwoCmdForRuntime0254).Add(setCurrentServerCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatestWithError).Add(getContextTwoCmdForRuntime0280WithError).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Set Context with Runtime v0.25.4 and Set Server with Runtime v0.28.0", func() {
			testCase := core.NewTestCase()

			// Add SetContext and SetCurrentContext Commands
			testCase.Add(setContextCmdForRuntime0254).Add(setContextTwoCmdForRuntime0254).Add(setCurrentContextCmdForRuntime0254)

			// Add SetServer and SetCurrentServer Commands
			testCase.Add(setServerCmdForRuntime0280).Add(setServerTwoCmdForRuntime0280).Add(setCurrentServerCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
	})

})
