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

	ginkgo.Context("using single server", func() {

		ginkgo.It("SetServer, SetCurrentServer latest then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest then DeleteServer, RemoveCurrentServer v0.28.0 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest", func() {
			ginkgo.By("Setting up the input and output parameters data for various APIs")
			// Input and Output Parameters for SetServer latest
			setServerInputOptions := server.DefaultSetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)

			// Input and Output Parameters for SetCurrentServer latest
			setCurrentServerInputOptions := server.DefaultSetCurrentServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentServer
			getCurrentServerInputOptionsForRuntimeLatest := server.DefaultGetCurrentServerInputOptions(core.VersionLatest)
			getCurrentServerOutputOptionsForRuntimeLatest := server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)

			getCurrentServerInputOptionsForRuntime0280 := server.DefaultGetCurrentServerInputOptions(core.Version0280)
			getCurrentServerOutputOptionsForRuntime0280 := server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0280WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)

			getCurrentServerInputOptionsForRuntime0254 := server.DefaultGetCurrentServerInputOptions(core.Version0254)
			getCurrentServerOutputOptionsForRuntime0254 := server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0254WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0254)

			// Input and Output params for GetServer
			getServerInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)

			getServerInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			getServerInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0254WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)

			// Input params for DeleteServer v0.28.0
			deleteServerInputOptions := server.DefaultDeleteServerInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input params for RemoveCurrentServer v0.28.0
			removeCurrentServerInputOptions := server.DefaultRemoveCurrentServerInputOptions(core.Version0280)

			ginkgo.By("Creating Commands to trigger Runtime APIs")

			// Create SetServer latest Command with input and output options
			setServerCmd, err := framework.NewSetServerCommand(setServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentServer latest Command with input and output options
			setCurrentServerCmd, err := framework.NewSetCurrentServerCommand(setCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetServer Commands with input and output options
			getServerCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntimeLatestWithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280WithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254WithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentServer Commands
			getCurrentServerCmdForRuntimeLatest, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntimeLatestWithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteServer Command
			deleteCtxCmd, err := framework.NewDeleteServerCommand(deleteServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentServer Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentServerCommand(removeCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Build test case with commands")

			// Add SetServer and SetCurrentServer Commands
			testCase := core.NewTestCase().Add(setServerCmd).Add(setCurrentServerCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError)

			ginkgo.By("Execute the testcase")

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("SetServer, SetCurrentServer v0.25.4 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest then DeleteServer, RemoveCurrentServer v0.28.0 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest", func() {
			ginkgo.By("Setting up the input and output parameters data for various APIs")
			// Input and Output Parameters for SetServer v0.25.4
			setServerInputOptions := server.DefaultSetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input Parameters for SetCurrentServer v0.25.4
			setCurrentServerInputOptions := server.DefaultSetCurrentServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentServer
			getCurrentServerInputOptionsForRuntimeLatest := server.DefaultGetCurrentServerInputOptions(core.VersionLatest)
			getCurrentServerInputOptionsForRuntime0280 := server.DefaultGetCurrentServerInputOptions(core.Version0280)
			getCurrentServerInputOptionsForRuntime0254 := server.DefaultGetCurrentServerInputOptions(core.Version0254)

			getCurrentServerOutputOptionsForRuntimeLatest := server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0280 := server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0254 := server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)

			//getCurrentServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)
			//getCurrentServerOutputOptionsForRuntime0280WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)

			// Input and Output params for GetServer
			getServerInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			getServerOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)

			//getServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)
			//getServerOutputOptionsForRuntime0280WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			// Input and Output params for RemoveCurrentServer v0.28.0
			removeCurrentServerInputOptions := server.DefaultRemoveCurrentServerInputOptions(core.Version0280)
			removeCurrentServerOutputOptionsWithError := server.DefaultRemoveCurrentServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			// Input and Output params for DeleteServer v0.28.0
			deleteServerInputOptions := server.DefaultDeleteServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			deleteServerOutputOptionsWithError := server.DefaultDeleteServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			ginkgo.By("Creating Commands to trigger Runtime APIs")

			// Create SetServer latest Command with input and output options
			setServerCmd, err := framework.NewSetServerCommand(setServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentServer latest Command with input and output options
			setCurrentServerCmd, err := framework.NewSetCurrentServerCommand(setCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetServer Commands with input and output options
			getServerCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			//getServerCmdForRuntimeLatestWithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatestWithError)
			//gomega.Expect(err).To(gomega.BeNil())
			//getServerCmdForRuntime0280WithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280WithError)
			//gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentServer Commands
			getCurrentServerCmdForRuntimeLatest, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			//getCurrentServerCmdForRuntimeLatestWithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatestWithError)
			//gomega.Expect(err).To(gomega.BeNil())
			//getCurrentServerCmdForRuntime0280WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280WithError)
			//gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteServer Command
			deleteCtxCmd, err := framework.NewDeleteServerCommand(deleteServerInputOptions, deleteServerOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentServer Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentServerCommand(removeCurrentServerInputOptions, removeCurrentServerOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Build test case with commands")

			// Add SetServer and SetCurrentServer Commands
			testCase := core.NewTestCase().Add(setServerCmd).Add(setCurrentServerCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			ginkgo.By("Execute the test case")
			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("SetServer, SetCurrentServer v0.28.0 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest then DeleteServer v0.25.4 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest", func() {
			ginkgo.By("Setting up the input and output parameters data for various APIs")

			// Input Parameters for SetServer v0.28.0
			setServerInputOptions := server.DefaultSetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input Parameters for SetCurrentServer v0.28.0
			setCurrentServerInputOptions := server.DefaultSetCurrentServerInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentServer
			getCurrentServerInputOptionsForRuntimeLatest := server.DefaultGetCurrentServerInputOptions(core.VersionLatest)
			getCurrentServerOutputOptionsForRuntimeLatest := server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)

			getCurrentServerInputOptionsForRuntime0280 := server.DefaultGetCurrentServerInputOptions(core.Version0280)
			getCurrentServerOutputOptionsForRuntime0280 := server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0280WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)

			getCurrentServerInputOptionsForRuntime0254 := server.DefaultGetCurrentServerInputOptions(core.Version0254)
			getCurrentServerOutputOptionsForRuntime0254 := server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0254WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0254)

			// Input and Output params for GetServer
			getServerInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			getServerOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			getServerOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0254WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)

			// Input params for DeleteServer v0.25.4
			deleteServerInputOptions := server.DefaultDeleteServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			ginkgo.By("Creating Commands to trigger Runtime APIs")

			// Create SetServer v0.28.0 Command with input and output options
			setServerCmd, err := framework.NewSetServerCommand(setServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentServer v0.28.0 Command with input and output options
			setCurrentServerCmd, err := framework.NewSetCurrentServerCommand(setCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetServer Commands with input and output options
			getServerCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			getServerCmdForRuntimeLatestWithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280WithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254WithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentServer Commands
			getCurrentServerCmdForRuntimeLatest, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			getCurrentServerCmdForRuntimeLatestWithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteServer v0.25.4 Command
			deleteCtxCmd, err := framework.NewDeleteServerCommand(deleteServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Build test case with commands")

			// Add SetServer and SetCurrentServer Commands
			testCase := core.NewTestCase().Add(setServerCmd).Add(setCurrentServerCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(deleteCtxCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError)

			ginkgo.By("Execute the test case")
			// Run all the commands
			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using multiple servers", func() {

		ginkgo.It("Run SetServer, SetCurrentServer on Runtime latest then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest then DeleteServer, RemoveCurrentServer v0.28.0 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest", func() {
			// Input and Output Parameters for SetServer latest
			setServerOneInputOptions := server.DefaultSetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			setServerTwoInputOptions := server.DefaultSetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)

			// Input and Output Parameters for SetCurrentServer latest
			setCurrentServerInputOptions := server.DefaultSetCurrentServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentServer
			getCurrentServerInputOptionsForRuntimeLatest := server.DefaultGetCurrentServerInputOptions(core.VersionLatest)

			getCurrentServerOutputOptionsForRuntimeLatest := server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)

			getCurrentServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)

			getCurrentServerInputOptionsForRuntime0280 := server.DefaultGetCurrentServerInputOptions(core.Version0280)

			getCurrentServerOutputOptionsForRuntime0280 := server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)

			getCurrentServerOutputOptionsForRuntime0280WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)

			getCurrentServerInputOptionsForRuntime0254 := server.DefaultGetCurrentServerInputOptions(core.Version0254)

			getCurrentServerOutputOptionsForRuntime0254 := server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)

			getCurrentServerOutputOptionsForRuntime0254WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0254)

			// Input and Output params for GetServer
			getServerOneInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOneOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOneOutputOptionsForRuntimeLatestWithError := server.DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)

			getServerTwoInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
			getServerTwoOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityTwo)

			getServerOneInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerOneOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerOneOutputOptionsForRuntime0280WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			getServerTwoInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)
			getServerTwoOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityTwo)

			getServerOneInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)
			getServerOneOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getServerOneOutputOptionsForRuntime0254WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)

			getServerTwoInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)
			getServerTwoOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityTwo)

			// Input params for DeleteServer v0.28.0
			deleteServerInputOptions := server.DefaultDeleteServerInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input params for RemoveCurrentServer v0.28.0
			removeCurrentServerInputOptions := server.DefaultRemoveCurrentServerInputOptions(core.Version0280)

			// Create SetServer latest Command with input and output options
			setServerOneCmd, err := framework.NewSetServerCommand(setServerOneInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			setServerTwoCmd, err := framework.NewSetServerCommand(setServerTwoInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentServer latest Command with input and output options
			setCurrentServerCmd, err := framework.NewSetCurrentServerCommand(setCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetServer Commands with input and output options
			getServerCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerOneInputOptionsForRuntimeLatest, getServerOneOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280, err := framework.NewGetServerCommand(getServerOneInputOptionsForRuntime0280, getServerOneOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254, err := framework.NewGetServerCommand(getServerOneInputOptionsForRuntime0254, getServerOneOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntimeLatestWithError, err := framework.NewGetServerCommand(getServerOneInputOptionsForRuntimeLatest, getServerOneOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280WithError, err := framework.NewGetServerCommand(getServerOneInputOptionsForRuntime0280, getServerOneOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254WithError, err := framework.NewGetServerCommand(getServerOneInputOptionsForRuntime0254, getServerOneOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			getServerTwoCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntimeLatest, getServerTwoOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getServerTwoCmdForRuntime0280, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0280, getServerTwoOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getServerTwoCmdForRuntime0254, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0254, getServerTwoOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentServer Commands
			getCurrentServerCmdForRuntimeLatest, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntimeLatestWithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteServer Command
			deleteCtxCmd, err := framework.NewDeleteServerCommand(deleteServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentServer Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentServerCommand(removeCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetServer and SetCurrentServer Commands
			testCase := core.NewTestCase().Add(setServerOneCmd).Add(setServerTwoCmd).Add(setCurrentServerCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer v0.25.4 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest then DeleteServer, RemoveCurrentServer v0.28.0 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest", func() {
			// Setting up the input and output parameters data for various APIs
			// Input and Output Parameters for SetServer v0.25.4
			setServerOneInputOptions := server.DefaultSetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)
			setServerTwoInputOptions := server.DefaultSetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)

			// Input Parameters for SetCurrentServer v0.25.4
			setCurrentServerInputOptions := server.DefaultSetCurrentServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentServer
			getCurrentServerInputOptionsForRuntimeLatest := server.DefaultGetCurrentServerInputOptions(core.VersionLatest)
			getCurrentServerInputOptionsForRuntime0280 := server.DefaultGetCurrentServerInputOptions(core.Version0280)
			getCurrentServerInputOptionsForRuntime0254 := server.DefaultGetCurrentServerInputOptions(core.Version0254)

			getCurrentServerOutputOptionsForRuntimeLatest := server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0280 := server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0254 := server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input and Output params for GetServer
			getServerInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			getServerOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)

			getServerTwoInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
			getServerTwoInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)
			getServerTwoInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)

			getServerTwoOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
			getServerTwoOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityTwo)
			getServerTwoOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityTwo)

			// Input and Output params for RemoveCurrentServer v0.28.0
			removeCurrentServerInputOptions := server.DefaultRemoveCurrentServerInputOptions(core.Version0280)
			removeCurrentServerOutputOptionsWithError := server.DefaultRemoveCurrentServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			// Input and Output params for DeleteServer v0.28.0
			deleteServerInputOptions := server.DefaultDeleteServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			deleteServerOutputOptionsWithError := server.DefaultDeleteServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			// Creating Commands to trigger Runtime APIs

			// Create SetServer latest Command with input and output options
			setServerOneCmd, err := framework.NewSetServerCommand(setServerOneInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())
			setServerTwoCmd, err := framework.NewSetServerCommand(setServerTwoInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentServer latest Command with input and output options
			setCurrentServerCmd, err := framework.NewSetCurrentServerCommand(setCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetServer Commands with input and output options
			getServerCmdForRuntime0254, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			getServerCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())

			getServerCmdForRuntime0280, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())

			getServerTwoCmdForRuntime0254, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0254, getServerTwoOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getServerTwoCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntimeLatest, getServerTwoOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())

			getServerTwoCmdForRuntime0280, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0280, getServerTwoOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentServer Commands
			getCurrentServerCmdForRuntime0254, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			getCurrentServerCmdForRuntimeLatest, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())

			getCurrentServerCmdForRuntime0280, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteServer Command
			deleteCtxCmd, err := framework.NewDeleteServerCommand(deleteServerInputOptions, deleteServerOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentServer Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentServerCommand(removeCurrentServerInputOptions, removeCurrentServerOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetServer and SetCurrentServer Commands
			testCase := core.NewTestCase().Add(setServerOneCmd).Add(setServerTwoCmd).Add(setCurrentServerCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Add RemoveCurrentServer v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteServer v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Execute the test case
			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetServer, SetCurrentServer v0.28.0 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest then DeleteServer v0.25.4 then GetServer, GetCurrentServer v0.25.4, v0.28.0, latest", func() {
			// Setting up the input and output parameters data for various APIs

			// Input Parameters for SetServer v0.28.0
			setServerOneInputOptions := server.DefaultSetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			setServerTwoInputOptions := server.DefaultSetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)

			// Input Parameters for SetCurrentServer v0.28.0
			setCurrentServerInputOptions := server.DefaultSetCurrentServerInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentServer
			getCurrentServerInputOptionsForRuntimeLatest := server.DefaultGetCurrentServerInputOptions(core.VersionLatest)
			getCurrentServerInputOptionsForRuntime0280 := server.DefaultGetCurrentServerInputOptions(core.Version0280)
			getCurrentServerInputOptionsForRuntime0254 := server.DefaultGetCurrentServerInputOptions(core.Version0254)

			getCurrentServerOutputOptionsForRuntimeLatest := server.DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0280 := server.DefaultGetCurrentServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0254 := server.DefaultGetCurrentServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getCurrentServerOutputOptionsForRuntime0254WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0254)
			getCurrentServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)
			getCurrentServerOutputOptionsForRuntime0280WithError := server.DefaultGetCurrentServerOutputOptionsWithError(core.Version0280)

			// Input and Output params for GetServer
			getServerInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			getServerOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0254WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntimeLatestWithError := server.DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)
			getServerOutputOptionsForRuntime0280WithError := server.DefaultGetServerOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			getServerTwoInputOptionsForRuntimeLatest := server.DefaultGetServerInputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
			getServerTwoInputOptionsForRuntime0280 := server.DefaultGetServerInputOptions(core.Version0280, common.CtxCompatibilityTwo)
			getServerTwoInputOptionsForRuntime0254 := server.DefaultGetServerInputOptions(core.Version0254, common.CtxCompatibilityTwo)

			getServerTwoOutputOptionsForRuntimeLatest := server.DefaultGetServerOutputOptions(core.VersionLatest, common.CtxCompatibilityTwo)
			getServerTwoOutputOptionsForRuntime0280 := server.DefaultGetServerOutputOptions(core.Version0280, common.CtxCompatibilityTwo)
			getServerTwoOutputOptionsForRuntime0254 := server.DefaultGetServerOutputOptions(core.Version0254, common.CtxCompatibilityTwo)

			// Input params for DeleteServer v0.25.4
			deleteServerInputOptions := server.DefaultDeleteServerInputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Creating Commands to trigger Runtime APIs

			// Create SetServer v0.28.0 Command with input and output options
			setServerOneCmd, err := framework.NewSetServerCommand(setServerOneInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())
			setServerTwoCmd, err := framework.NewSetServerCommand(setServerTwoInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentServer v0.28.0 Command with input and output options
			setCurrentServerCmd, err := framework.NewSetCurrentServerCommand(setCurrentServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetServer Commands with input and output options
			getServerCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0254WithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0254, getServerOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			getServerCmdForRuntimeLatestWithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntimeLatest, getServerOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getServerCmdForRuntime0280WithError, err := framework.NewGetServerCommand(getServerInputOptionsForRuntime0280, getServerOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())

			getServerTwoCmdForRuntimeLatest, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntimeLatest, getServerTwoOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getServerTwoCmdForRuntime0280, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0280, getServerTwoOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getServerTwoCmdForRuntime0254, err := framework.NewGetServerCommand(getServerTwoInputOptionsForRuntime0254, getServerTwoOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentServer Commands
			getCurrentServerCmdForRuntimeLatest, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatest)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0254WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0254, getCurrentServerOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntimeLatestWithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntimeLatest, getCurrentServerOutputOptionsForRuntimeLatestWithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentServerCmdForRuntime0280WithError, err := framework.NewGetCurrentServerCommand(getCurrentServerInputOptionsForRuntime0280, getCurrentServerOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteServer v0.25.4 Command
			deleteServerCmd, err := framework.NewDeleteServerCommand(deleteServerInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetServer and SetCurrentServer Commands
			testCase := core.NewTestCase().Add(setServerOneCmd).Add(setServerTwoCmd).Add(setCurrentServerCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatest).Add(getServerCmdForRuntime0280).Add(getServerCmdForRuntime0254)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatest).Add(getCurrentServerCmdForRuntime0280).Add(getCurrentServerCmdForRuntime0254)

			// Add DeleteServer v0.25.4 Command
			testCase.Add(deleteServerCmd)

			// Add GetServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getServerCmdForRuntimeLatestWithError).Add(getServerCmdForRuntime0280WithError).Add(getServerCmdForRuntime0254WithError)
			testCase.Add(getServerTwoCmdForRuntimeLatest).Add(getServerTwoCmdForRuntime0280).Add(getServerTwoCmdForRuntime0254)

			// Add GetCurrentServer latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentServerCmdForRuntimeLatestWithError).Add(getCurrentServerCmdForRuntime0280WithError).Add(getCurrentServerCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})
	})
})
