// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package server_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/server"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

var _ = ginkgo.Describe("Cross-version Server APIs Compatibility Tests for supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("Get/Set/Delete Server and CurrentServer API methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})

	})
	// Input and Output Options for Server APIs
	var setServerInputOptionsForRuntime0116 *framework.SetServerInputOptions
	var setServerInputOptionsForRuntime0254 *framework.SetServerInputOptions
	var setServerInputOptionsForRuntime0280 *framework.SetServerInputOptions
	var setServerInputOptionsForRuntimeLatest *framework.SetServerInputOptions

	var setServerTwoInputOptionsForRuntime0116 *framework.SetServerInputOptions
	var setServerTwoInputOptionsForRuntime0254 *framework.SetServerInputOptions
	var setServerTwoInputOptionsForRuntime0280 *framework.SetServerInputOptions
	var setServerTwoInputOptionsForRuntimeLatest *framework.SetServerInputOptions

	var setCurrentServerInputOptionsForRuntime0116 *framework.SetCurrentServerInputOptions
	var setCurrentServerInputOptionsForRuntime0254 *framework.SetCurrentServerInputOptions
	var setCurrentServerInputOptionsForRuntime0280 *framework.SetCurrentServerInputOptions
	var setCurrentServerInputOptionsForRuntimeLatest *framework.SetCurrentServerInputOptions

	var getServerInputOptionsForRuntimeLatest *framework.GetServerInputOptions
	var getServerInputOptionsForRuntime0280 *framework.GetServerInputOptions
	var getServerInputOptionsForRuntime0254 *framework.GetServerInputOptions
	var getServerInputOptionsForRuntime0116 *framework.GetServerInputOptions

	var getServerTwoInputOptionsForRuntimeLatest *framework.GetServerInputOptions
	var getServerTwoInputOptionsForRuntime0280 *framework.GetServerInputOptions
	var getServerTwoInputOptionsForRuntime0254 *framework.GetServerInputOptions
	var getServerTwoInputOptionsForRuntime0116 *framework.GetServerInputOptions

	var getServerOutputOptionsForRuntime0116 *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntime0254 *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntime0280 *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntimeLatest *framework.GetServerOutputOptions

	var getServerTwoOutputOptionsForRuntime0116 *framework.GetServerOutputOptions
	var getServerTwoOutputOptionsForRuntime0254 *framework.GetServerOutputOptions
	var getServerTwoOutputOptionsForRuntime0280 *framework.GetServerOutputOptions
	var getServerTwoOutputOptionsForRuntimeLatest *framework.GetServerOutputOptions

	var getServerOutputOptionsForRuntimeLatestWithError *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntime0280WithError *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntime0254WithError *framework.GetServerOutputOptions
	var getServerOutputOptionsForRuntime0116WithError *framework.GetServerOutputOptions

	var getCurrentServerInputOptionsForRuntime0116 *framework.GetCurrentServerInputOptions
	var getCurrentServerInputOptionsForRuntime0254 *framework.GetCurrentServerInputOptions
	var getCurrentServerInputOptionsForRuntime0280 *framework.GetCurrentServerInputOptions
	var getCurrentServerInputOptionsForRuntimeLatest *framework.GetCurrentServerInputOptions

	var getCurrentServerOutputOptionsForRuntime0116 *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntime0254 *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntime0280 *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntimeLatest *framework.GetCurrentServerOutputOptions

	var getCurrentServerOutputOptionsForRuntimeLatestWithError *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntime0280WithError *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntime0254WithError *framework.GetCurrentServerOutputOptions
	var getCurrentServerOutputOptionsForRuntime0116WithError *framework.GetCurrentServerOutputOptions

	var deleteServerInputOptionsForRuntime0254 *framework.DeleteServerInputOptions
	var deleteServerInputOptionsForRuntime0280 *framework.DeleteServerInputOptions
	var deleteServerInputOptionsForRuntimeLatest *framework.DeleteServerInputOptions
	var deleteServerInputOptionsForRuntime0116 *framework.DeleteServerInputOptions

	var deleteServerOutputOptionsForRuntime0280WithError *framework.DeleteServerOutputOptions
	var deleteServerOutputOptionsForRuntimeLatestWithError *framework.DeleteServerOutputOptions

	var removeCurrentServerInputOptionsForRuntime0280 *framework.RemoveCurrentServerInputOptions
	var removeCurrentServerOutputOptionsForRuntime0280WithError *framework.RemoveCurrentServerOutputOptions

	// Server API Commands
	var setServerCmdForRuntimeLatest *core.Command
	var setServerCmdForRuntime0280 *core.Command
	var setServerCmdForRuntime0254 *core.Command
	var setServerCmdForRuntime0116 *core.Command

	var setServerTwoCmdForRuntimeLatest *core.Command
	var setServerTwoCmdForRuntime0280 *core.Command
	var setServerTwoCmdForRuntime0254 *core.Command
	var setServerTwoCmdForRuntime0116 *core.Command

	var setCurrentServerCmdForRuntime0116 *core.Command
	var setCurrentServerCmdForRuntime0254 *core.Command
	var setCurrentServerCmdForRuntime0280 *core.Command
	var setCurrentServerCmdForRuntimeLatest *core.Command

	var getServerCmdForRuntimeLatest *core.Command
	var getServerCmdForRuntime0280 *core.Command
	var getServerCmdForRuntime0254 *core.Command
	var getServerCmdForRuntime0116 *core.Command

	var getServerTwoCmdForRuntimeLatest *core.Command
	var getServerTwoCmdForRuntime0280 *core.Command
	var getServerTwoCmdForRuntime0254 *core.Command
	var getServerTwoCmdForRuntime0116 *core.Command

	var getServerCmdForRuntimeLatestWithError *core.Command
	var getServerCmdForRuntime0280WithError *core.Command
	var getServerCmdForRuntime0254WithError *core.Command
	var getServerCmdForRuntime0116WithError *core.Command

	var getCurrentServerCmdForRuntimeLatest *core.Command
	var getCurrentServerCmdForRuntime0280 *core.Command
	var getCurrentServerCmdForRuntime0254 *core.Command
	var getCurrentServerCmdForRuntime0116 *core.Command

	var getCurrentServerCmdForRuntimeLatestWithError *core.Command
	var getCurrentServerCmdForRuntime0280WithError *core.Command
	var getCurrentServerCmdForRuntime0254WithError *core.Command
	var getCurrentServerCmdForRuntime0116WithError *core.Command

	var deleteServerCmdForRuntime0280 *core.Command
	var deleteServerCmdForRuntime0116 *core.Command
	var deleteServerCmdForRuntime0254 *core.Command

	var deleteServerCmdForRuntime0280WithError *core.Command
	var deleteServerCmdForRuntimeLatestWithError *core.Command

	var removeCurrentServerCmdForRuntime0280 *core.Command
	var removeCurrentServerCmdForRuntime0280WithError *core.Command

	var err error
	ginkgo.BeforeEach(func() {
		ginkgo.By("Setup Input and Output Options for Servers APIs")

		// Input and Output Parameters for SetServer
		setServerInputOptionsForRuntimeLatest = server.DefaultSetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		setServerInputOptionsForRuntime0280 = server.DefaultSetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
		setServerInputOptionsForRuntime0254 = server.DefaultSetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)
		setServerInputOptionsForRuntime0116 = server.DefaultSetServerInputOptions(core.Version0116, common.CtxCompatibilityOne)

		setServerTwoInputOptionsForRuntimeLatest = server.DefaultSetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		setServerTwoInputOptionsForRuntime0280 = server.DefaultSetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)
		setServerTwoInputOptionsForRuntime0254 = server.DefaultSetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)
		setServerTwoInputOptionsForRuntime0116 = server.DefaultSetServerInputOptions(core.Version0116, common.CtxCompatibilityTwo)

		// Input and Output Parameters for SetCurrentServer
		setCurrentServerInputOptionsForRuntimeLatest = server.DefaultSetCurrentServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		setCurrentServerInputOptionsForRuntime0280 = server.DefaultSetCurrentServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
		setCurrentServerInputOptionsForRuntime0254 = server.DefaultSetCurrentServerInputOptions(core.Version0254, common.CtxCompatibilityOne)
		setCurrentServerInputOptionsForRuntime0116 = server.DefaultSetCurrentServerInputOptions(core.Version0116, common.CtxCompatibilityOne)

		// Input and Output Parameters for GetCurrentServer
		getCurrentServerInputOptionsForRuntimeLatest = server.DefaultGetCurrentServerInputOptions(core.VersionLatest)
		getCurrentServerInputOptionsForRuntime0280 = server.DefaultGetCurrentServerInputOptions(core.Version0280)
		getCurrentServerInputOptionsForRuntime0254 = server.DefaultGetCurrentServerInputOptions(core.Version0254)
		getCurrentServerInputOptionsForRuntime0116 = server.DefaultGetCurrentServerInputOptions(core.Version0116)

		getCurrentServerOutputOptionsForRuntime0280 = server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
		getCurrentServerOutputOptionsForRuntime0254 = server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
		getCurrentServerOutputOptionsForRuntime0116 = server.DefaultGetCurrentServerOutputOptions(core.Version0116, common.CtxCompatibilityOne)
		getCurrentServerOutputOptionsForRuntimeLatest = server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		getCurrentServerOutputOptionsForRuntimeLatestWithError = server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)
		getCurrentServerOutputOptionsForRuntime0280WithError = server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0280)
		getCurrentServerOutputOptionsForRuntime0254WithError = server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0254)
		getCurrentServerOutputOptionsForRuntime0116WithError = server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0116)

		// Input and Output params for GetServer
		getServerInputOptionsForRuntimeLatest = server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
		getServerInputOptionsForRuntime0280 = server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
		getServerInputOptionsForRuntime0254 = server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)
		getServerInputOptionsForRuntime0116 = server.DefaultGetServerInputOptions(core.Version0116, common.CtxCompatibilityOne)

		getServerTwoInputOptionsForRuntimeLatest = server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		getServerTwoInputOptionsForRuntime0280 = server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)
		getServerTwoInputOptionsForRuntime0254 = server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)
		getServerTwoInputOptionsForRuntime0116 = server.DefaultGetServerInputOptions(core.Version0116, common.CtxCompatibilityTwo)

		getServerTwoOutputOptionsForRuntimeLatest = server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
		getServerTwoOutputOptionsForRuntime0280 = server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityTwo)
		getServerTwoOutputOptionsForRuntime0254 = server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityTwo)
		getServerTwoOutputOptionsForRuntime0116 = server.DefaultGetServerOutputOptions(core.Version0116, common.CtxCompatibilityTwo)

		getServerOutputOptionsForRuntime0280 = server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntime0254 = server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntime0116 = server.DefaultGetServerOutputOptions(core.Version0116, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntimeLatest = server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		getServerOutputOptionsForRuntimeLatestWithError = server.DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntime0280WithError = server.DefaultGetServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntime0254WithError = server.DefaultGetServerOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)
		getServerOutputOptionsForRuntime0116WithError = server.DefaultGetServerOutputOptionsWithError(core.Version0116, common.CtxCompatibilityOne)

		// Input and Output Options for DeleteServer
		deleteServerInputOptionsForRuntime0280 = server.DefaultDeleteServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
		deleteServerInputOptionsForRuntime0254 = server.DefaultDeleteServerInputOptions(core.Version0254, common.CtxCompatibilityOne)
		deleteServerInputOptionsForRuntime0116 = server.DefaultDeleteServerInputOptions(core.Version0116, common.CtxCompatibilityOne)
		deleteServerInputOptionsForRuntimeLatest = server.DefaultDeleteServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		deleteServerOutputOptionsForRuntime0280WithError = server.DefaultDeleteServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)
		deleteServerOutputOptionsForRuntimeLatestWithError = server.DefaultDeleteServerOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)

		// Input and Output Options for RemoveCurrentServer
		removeCurrentServerInputOptionsForRuntime0280 = server.DefaultRemoveCurrentServerInputOptions(core.Version0280)

		removeCurrentServerOutputOptionsForRuntime0280WithError = server.DefaultRemoveCurrentServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

		ginkgo.By("Setup Server API commands")

		// Create SetServer Commands with input and output options
		setServerCmdForRuntimeLatest, err = framework.NewSetServerCommand(setServerInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerCmdForRuntime0254, err = framework.NewSetServerCommand(setServerInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerCmdForRuntime0280, err = framework.NewSetServerCommand(setServerInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerCmdForRuntime0116, err = framework.NewSetServerCommand(setServerInputOptionsForRuntime0116, nil)
		gomega.Expect(err).To(gomega.BeNil())

		setServerTwoCmdForRuntimeLatest, err = framework.NewSetServerCommand(setServerTwoInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerTwoCmdForRuntime0254, err = framework.NewSetServerCommand(setServerTwoInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerTwoCmdForRuntime0280, err = framework.NewSetServerCommand(setServerTwoInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setServerTwoCmdForRuntime0116, err = framework.NewSetServerCommand(setServerTwoInputOptionsForRuntime0116, nil)
		gomega.Expect(err).To(gomega.BeNil())

		// Create SetCurrentServer Commands with input and output options
		setCurrentServerCmdForRuntimeLatest, err = framework.NewSetCurrentServerCommand(setCurrentServerInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setCurrentServerCmdForRuntime0280, err = framework.NewSetCurrentServerCommand(setCurrentServerInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setCurrentServerCmdForRuntime0254, err = framework.NewSetCurrentServerCommand(setCurrentServerInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		setCurrentServerCmdForRuntime0116, err = framework.NewSetCurrentServerCommand(setCurrentServerInputOptionsForRuntime0116, nil)
		gomega.Expect(err).To(gomega.BeNil())

		// Create GetServer Commands with input and output options
		getServerCmdForRuntimeLatest, err = framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0280, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0254, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0116, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0116, getServerOutputOptionsForRuntime0116)
		gomega.Expect(err).To(gomega.BeNil())

		getServerTwoCmdForRuntimeLatest, err = framework.NewGetServerCommand(getServerTwoInputOptionsForRuntimeLatest, getServerTwoOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getServerTwoCmdForRuntime0280, err = framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0280, getServerTwoOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getServerTwoCmdForRuntime0254, err = framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0254, getServerTwoOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())
		getServerTwoCmdForRuntime0116, err = framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0116, getServerTwoOutputOptionsForRuntime0116)
		gomega.Expect(err).To(gomega.BeNil())

		getServerCmdForRuntimeLatestWithError, err = framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0280WithError, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0254WithError, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254WithError)
		gomega.Expect(err).To(gomega.BeNil())
		getServerCmdForRuntime0116WithError, err = framework.NewGetServerCommand(getServerInputOptionsForRuntime0116, getServerOutputOptionsForRuntime0116WithError)
		gomega.Expect(err).To(gomega.BeNil())

		// Create GetCurrentServer Commands with input and output options
		getCurrentServerCmdForRuntimeLatest, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0280, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0254, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0116, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0116, getCurrentServerOutputOptionsForRuntime0116)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntimeLatestWithError, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0280WithError, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0254WithError, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254WithError)
		gomega.Expect(err).To(gomega.BeNil())
		getCurrentServerCmdForRuntime0116WithError, err = framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0116, getCurrentServerOutputOptionsForRuntime0116WithError)
		gomega.Expect(err).To(gomega.BeNil())

		// Create DeleteServer Commands with input and output options
		deleteServerCmdForRuntime0280, err = framework.NewDeleteServerCommand(deleteServerInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		deleteServerCmdForRuntime0254, err = framework.NewDeleteServerCommand(deleteServerInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())
		deleteServerCmdForRuntime0116, err = framework.NewDeleteServerCommand(deleteServerInputOptionsForRuntime0116, nil)
		gomega.Expect(err).To(gomega.BeNil())
		deleteServerCmdForRuntime0280WithError, err = framework.NewDeleteServerCommand(deleteServerInputOptionsForRuntime0280, deleteServerOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
		deleteServerCmdForRuntimeLatestWithError, err = framework.NewDeleteServerCommand(deleteServerInputOptionsForRuntimeLatest, deleteServerOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())

		// Create RemoveCurrentServer Commands with input and output options
		removeCurrentServerCmdForRuntime0280, err = framework.NewRemoveCurrentServerCommand(removeCurrentServerInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		removeCurrentServerCmdForRuntime0280WithError, err = framework.NewRemoveCurrentServerCommand(removeCurrentServerInputOptionsForRuntime0280, removeCurrentServerOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.Context("using single server", func() {

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime Latest version
			testCase := core.NewTestCase().Add(setServerCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentServerCmdForRuntime0280)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError).Add(getServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError).Add(getCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime v0.11.6 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime Latest version
			testCase := core.NewTestCase().Add(setServerCmdForRuntimeLatest).Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.11.6 Command
			testCase.Add(deleteServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError).Add(getServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError).Add(getCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0280).Add(setCurrentServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(deleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError).Add(getServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError).Add(getCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.25.4
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0254).Add(setCurrentServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentServerCmdForRuntime0280WithError)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteServerCmdForRuntime0280WithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime 0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0116).Add(setCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(deleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError).Add(getServerCmdForRuntime0116WithError)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError).Add(getCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add SetServer and SetCurrentServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0116).Add(setCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add DeleteServer latest Command
			testCase.Add(deleteServerCmdForRuntimeLatestWithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Run all the commands
			executer.Execute(testCase)
		})
	})

	ginkgo.Context("using multiple servers", func() {

		ginkgo.It("Run two SetServer of Runtime latest then SetCurrentServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions and then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime Latest
			testCase := core.NewTestCase().Add(setServerCmdForRuntimeLatest).Add(setServerTwoCmdForRuntimeLatest)

			// Add SetCurrentServer Command of Runtime Latest
			testCase.Add(setCurrentServerCmdForRuntimeLatest)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentServerCmdForRuntime0280)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError).Add(getServerCmdForRuntime0116WithError)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.25.4 then SetCurrentServer of Runtime v0.25.4 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer, RemoveCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime v0.25.4
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0254).Add(setServerTwoCmdForRuntime0254)

			// Add SetCurrentServer Command of Runtime v0.25.4
			testCase.Add(setCurrentServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentServerCmdForRuntime0280WithError)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteServerCmdForRuntime0280WithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.28.0 then SetCurrentServer of Runtime v0.28.0 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {

			// Add two SetServer Commands of Runtime v0.28.0
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0280).Add(setServerTwoCmdForRuntime0280)

			// Add SetCurrentServer Command of Runtime v0.28.0
			testCase.Add(setCurrentServerCmdForRuntime0280)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(deleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError).Add(getServerCmdForRuntime0116WithError)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError).Add(getCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.11.6 then SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer of Runtime latest then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0116).Add(setServerTwoCmdForRuntime0116)

			// Add SetCurrentServer Command of Runtime v0.11.6
			testCase.Add(setCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add DeleteServer latest Command
			testCase.Add(deleteServerCmdForRuntimeLatestWithError)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run two SetServer of Runtime v0.11.6 then SetCurrentServer of Runtime v0.11.6 then GetServer, GetCurrentServer on v0.11.6, v0.25.4, v0.28.0, latest then DeleteServer of Runtime v0.25.4 then GetServer, GetCurrentServer on all supported Runtime library versions", func() {
			// Add two SetServer Commands of Runtime v0.11.6
			testCase := core.NewTestCase().Add(setServerCmdForRuntime0116).Add(setServerTwoCmdForRuntime0116)

			// Add SetCurrentServer Command of Runtime v0.11.6
			testCase.Add(setCurrentServerCmdForRuntime0116)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254).Add(getServerCmdForRuntime0116)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254).Add(getCurrentServerCmdForRuntime0116)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(deleteServerCmdForRuntime0254)

			// Add GetServer Commands on all supported Runtime library versions
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError).Add(getServerCmdForRuntime0116WithError)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254).Add(getServerTwoCmdForRuntime0116)

			// Add GetCurrentServer Commands on all supported Runtime library versions
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError).Add(getCurrentServerCmdForRuntime0116WithError)

			// Run all the commands
			executer.Execute(testCase)
		})
	})
})
