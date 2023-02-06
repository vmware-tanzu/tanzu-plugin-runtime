// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package context_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/context"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version Context APIs compatibility tests", func() {
	// Description on the Tests
	ginkgo.GinkgoWriter.Println("GetContext, SetContext, DeleteContext, GetCurrentContext, SetCurrentContext, RemoveCurrentContext methods are tested for cross-version API compatibility with supported Runtime versions v0.25.4, v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using single context object on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Input and Output Parameters for SetContext latest
			setContextInputOptions := context.DefaultSetContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)

			// Input and Output Parameters for SetCurrentContext latest
			setCurrentContextInputOptions := context.DefaultSetCurrentContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentContext
			getCurrentContextInputOptionsForRuntime100 := context.DefaultGetCurrentContextInputOptions(core.VersionLatest)
			getCurrentContextOutputOptionsForRuntime100 := context.DefaultGetCurrentContextOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime100WithError := context.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)

			getCurrentContextInputOptionsForRuntime0280 := context.DefaultGetCurrentContextInputOptions(core.Version0280)
			getCurrentContextOutputOptionsForRuntime0280 := context.DefaultGetCurrentContextOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0280WithError := context.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)

			getCurrentContextInputOptionsForRuntime0254 := context.DefaultGetCurrentContextInputOptions(core.Version0254)
			getCurrentContextOutputOptionsForRuntime0254 := context.DefaultGetCurrentContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0254WithError := context.DefaultGetCurrentContextOutputOptionsWithError(core.Version0254)

			// Input and Output params for GetContext
			getContextInputOptionsForRuntime100 := context.DefaultGetContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime100 := context.DefaultGetContextOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime100WithError := context.DefaultGetContextOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)

			getContextInputOptionsForRuntime0280 := context.DefaultGetContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280 := context.DefaultGetContextOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280WithError := context.DefaultGetContextOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			getContextInputOptionsForRuntime0254 := context.DefaultGetContextInputOptions(core.Version0254, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254 := context.DefaultGetContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254WithError := context.DefaultGetContextOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)

			// Input params for DeleteContext v0.28.0
			deleteContextInputOptions := context.DefaultDeleteContextInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input params for RemoveCurrentContext v0.28.0
			removeCurrentContextInputOptions := context.DefaultRemoveCurrentContextInputOptions(core.Version0280)

			// Create SetContext latest Command with input and output options
			setContextCmd, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext latest Command with input and output options
			setCurrentContextCmd, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			getContextCmdForRuntime100, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime100, getContextOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime100WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime100, getContextOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			getCurrentContextCmdForRuntime100, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime100WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext Command
			deleteCtxCmd, err := framework.NewDeleteContextCommand(deleteContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentContext Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmd).Add(setCurrentContextCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100WithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100WithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Setting up the input and output parameters data for various APIs
			// Input and Output Parameters for SetContext v0.25.4
			setContextInputOptions := context.DefaultSetContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input Parameters for SetCurrentContext v0.25.4
			setCurrentContextInputOptions := context.DefaultSetCurrentContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentContext
			getCurrentContextInputOptionsForRuntime100 := context.DefaultGetCurrentContextInputOptions(core.VersionLatest)
			getCurrentContextInputOptionsForRuntime0280 := context.DefaultGetCurrentContextInputOptions(core.Version0280)
			getCurrentContextInputOptionsForRuntime0254 := context.DefaultGetCurrentContextInputOptions(core.Version0254)

			getCurrentContextOutputOptionsForRuntime100WithError := context.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)
			getCurrentContextOutputOptionsForRuntime0280WithError := context.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)
			getCurrentContextOutputOptionsForRuntime0254 := context.DefaultGetCurrentContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input and Output params for GetContext
			getContextInputOptionsForRuntime100 := context.DefaultGetContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0280 := context.DefaultGetContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0254 := context.DefaultGetContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

			getContextOutputOptionsForRuntime100WithError := context.DefaultGetContextOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280WithError := context.DefaultGetContextOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254 := context.DefaultGetContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Input and Output params for RemoveCurrentContext v0.28.0
			removeCurrentContextInputOptions := context.DefaultRemoveCurrentContextInputOptions(core.Version0280)
			removeCurrentContextOutputOptionsWithError := context.DefaultRemoveCurrentContextOutputOptionsWithError(core.Version0280)

			// Input and Output params for DeleteContext v0.28.0
			deleteContextInputOptions := context.DefaultDeleteContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
			deleteContextOutputOptionsWithError := context.DefaultDeleteContextOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)

			// Creating Commands to trigger Runtime APIs

			// Create SetContext latest Command with input and output options
			setContextCmd, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext latest Command with input and output options
			setCurrentContextCmd, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			getContextCmdForRuntime0254, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime100WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime100, getContextOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			getCurrentContextCmdForRuntime0254, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime100WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext Command
			deleteCtxCmd, err := framework.NewDeleteContextCommand(deleteContextInputOptions, deleteContextOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentContext Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptions, removeCurrentContextOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmd).Add(setCurrentContextCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100WithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100WithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100WithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100WithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Execute the test case
			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Setting up the input and output parameters data for various APIs

			// Input Parameters for SetContext v0.28.0
			setContextInputOptions := context.DefaultSetContextInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input Parameters for SetCurrentContext v0.28.0
			setCurrentContextInputOptions := context.DefaultSetCurrentContextInputOptions(core.Version0280, common.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentContext
			getCurrentContextInputOptionsForRuntime100 := context.DefaultGetCurrentContextInputOptions(core.VersionLatest)
			getCurrentContextInputOptionsForRuntime0280 := context.DefaultGetCurrentContextInputOptions(core.Version0280)
			getCurrentContextInputOptionsForRuntime0254 := context.DefaultGetCurrentContextInputOptions(core.Version0254)

			getCurrentContextOutputOptionsForRuntime100 := context.DefaultGetCurrentContextOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0280 := context.DefaultGetCurrentContextOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0254 := context.DefaultGetCurrentContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0254WithError := context.DefaultGetCurrentContextOutputOptionsWithError(core.Version0254)

			// Input and Output params for GetContext
			getContextInputOptionsForRuntime100 := context.DefaultGetContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0280 := context.DefaultGetContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0254 := context.DefaultGetContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

			getContextOutputOptionsForRuntime100 := context.DefaultGetContextOutputOptions(core.VersionLatest, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280 := context.DefaultGetContextOutputOptions(core.Version0280, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254 := context.DefaultGetContextOutputOptions(core.Version0254, common.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254WithError := context.DefaultGetContextOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)

			// Input params for DeleteContext v0.25.4
			deleteContextInputOptions := context.DefaultDeleteContextInputOptions(core.Version0254, common.CtxCompatibilityOne)

			// Creating Commands to trigger Runtime APIs

			// Create SetContext v0.28.0 Command with input and output options
			setContextCmd, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext v0.28.0 Command with input and output options
			setCurrentContextCmd, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			getContextCmdForRuntime100, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime100, getContextOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			getCurrentContextCmdForRuntime100, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext v0.25.4 Command
			deleteCtxCmd, err := framework.NewDeleteContextCommand(deleteContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmd).Add(setCurrentContextCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add DeleteContext v0.25.4 Command
			testCase.Add(deleteCtxCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using multiple context objects on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext on Runtime latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Input and Output Parameters for SetContext latest
			setContextOneInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			setContextTwoInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)

			// Input and Output Parameters for SetCurrentContext latest
			setCurrentContextInputOptions := compatibilitytests.DefaultSetCurrentContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentContext
			getCurrentContextInputOptionsForRuntime100 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.VersionLatest)

			getCurrentContextOutputOptionsForRuntime100 := compatibilitytests.DefaultGetCurrentContextOutputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)

			getCurrentContextOutputOptionsForRuntime100WithError := compatibilitytests.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)

			getCurrentContextInputOptionsForRuntime0280 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.Version0280)

			getCurrentContextOutputOptionsForRuntime0280 := compatibilitytests.DefaultGetCurrentContextOutputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)

			getCurrentContextOutputOptionsForRuntime0280WithError := compatibilitytests.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)

			getCurrentContextInputOptionsForRuntime0254 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.Version0254)

			getCurrentContextOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetCurrentContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			getCurrentContextOutputOptionsForRuntime0254WithError := compatibilitytests.DefaultGetCurrentContextOutputOptionsWithError(core.Version0254)

			// Input and Output params for GetContext
			getContextOneInputOptionsForRuntime100 := compatibilitytests.DefaultGetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextOneOutputOptionsForRuntime100 := compatibilitytests.DefaultGetContextOutputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextOneOutputOptionsForRuntime100WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)

			getContextTwoInputOptionsForRuntime100 := compatibilitytests.DefaultGetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoOutputOptionsForRuntime100 := compatibilitytests.DefaultGetContextOutputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)
			// getContextTwoOutputOptionsForRuntime100WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)

			getContextOneInputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextOneOutputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextOneOutputOptionsForRuntime0280WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0280, compatibilitytests.CtxCompatibilityOne)

			getContextTwoInputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoOutputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0280, compatibilitytests.CtxCompatibilityTwo)
			// getContextTwoOutputOptionsForRuntime0280WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0280, compatibilitytests.CtxCompatibilityTwo)

			getContextOneInputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)
			getContextOneOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)
			getContextOneOutputOptionsForRuntime0254WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			getContextTwoInputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityTwo)
			// getContextTwoOutputOptionsForRuntime0254WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0254, compatibilitytests.CtxCompatibilityTwo)

			// Input params for DeleteContext v0.28.0
			deleteContextInputOptions := compatibilitytests.DefaultDeleteContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)

			// Input params for RemoveCurrentContext v0.28.0
			removeCurrentContextInputOptions := compatibilitytests.DefaultRemoveCurrentContextInputOptions(core.Version0280)

			// Create SetContext latest Command with input and output options
			setContextOneCmd, err := framework.NewSetContextCommand(setContextOneInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			setContextTwoCmd, err := framework.NewSetContextCommand(setContextTwoInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext latest Command with input and output options
			setCurrentContextCmd, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			getContextCmdForRuntime100, err := framework.NewGetContextCommand(getContextOneInputOptionsForRuntime100, getContextOneOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280, err := framework.NewGetContextCommand(getContextOneInputOptionsForRuntime0280, getContextOneOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254, err := framework.NewGetContextCommand(getContextOneInputOptionsForRuntime0254, getContextOneOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime100WithError, err := framework.NewGetContextCommand(getContextOneInputOptionsForRuntime100, getContextOneOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280WithError, err := framework.NewGetContextCommand(getContextOneInputOptionsForRuntime0280, getContextOneOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254WithError, err := framework.NewGetContextCommand(getContextOneInputOptionsForRuntime0254, getContextOneOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			getContextTwoCmdForRuntime100, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime100, getContextTwoOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getContextTwoCmdForRuntime0280, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0280, getContextTwoOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getContextTwoCmdForRuntime0254, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0254, getContextTwoOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			//getContextTwoCmdForRuntime100WithError, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime100, getContextTwoOutputOptionsForRuntime100WithError)
			//gomega.Expect(err).To(gomega.BeNil())
			//getContextTwoCmdForRuntime0280WithError, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0280, getContextTwoOutputOptionsForRuntime0280WithError)
			//gomega.Expect(err).To(gomega.BeNil())
			//getContextTwoCmdForRuntime0254WithError, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0254, getContextTwoOutputOptionsForRuntime0254WithError)
			//gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			getCurrentContextCmdForRuntime100, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime100WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext Command
			deleteCtxCmd, err := framework.NewDeleteContextCommand(deleteContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentContext Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextOneCmd).Add(setContextTwoCmd).Add(setCurrentContextCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntime100).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100WithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254WithError)
			testCase.Add(getContextTwoCmdForRuntime100).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100WithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Setting up the input and output parameters data for various APIs
			// Input and Output Parameters for SetContext v0.25.4
			setContextOneInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)
			setContextTwoInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityTwo)

			// Input Parameters for SetCurrentContext v0.25.4
			setCurrentContextInputOptions := compatibilitytests.DefaultSetCurrentContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentContext
			getCurrentContextInputOptionsForRuntime100 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.VersionLatest)
			getCurrentContextInputOptionsForRuntime0280 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.Version0280)
			getCurrentContextInputOptionsForRuntime0254 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.Version0254)

			getCurrentContextOutputOptionsForRuntime100WithError := compatibilitytests.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)
			getCurrentContextOutputOptionsForRuntime0280WithError := compatibilitytests.DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)
			getCurrentContextOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetCurrentContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			// Input and Output params for GetContext
			getContextInputOptionsForRuntime100 := compatibilitytests.DefaultGetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			getContextOutputOptionsForRuntime100WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			getContextTwoInputOptionsForRuntime100 := compatibilitytests.DefaultGetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoInputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoInputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityTwo)

			getContextTwoOutputOptionsForRuntime100WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoOutputOptionsForRuntime0280WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0280, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityTwo)

			// Input and Output params for RemoveCurrentContext v0.28.0
			removeCurrentContextInputOptions := compatibilitytests.DefaultRemoveCurrentContextInputOptions(core.Version0280)
			removeCurrentContextOutputOptionsWithError := compatibilitytests.DefaultRemoveCurrentContextOutputOptionsWithError(core.Version0280)

			// Input and Output params for DeleteContext v0.28.0
			deleteContextInputOptions := compatibilitytests.DefaultDeleteContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			deleteContextOutputOptionsWithError := compatibilitytests.DefaultDeleteContextOutputOptionsWithError(core.Version0280, compatibilitytests.CtxCompatibilityOne)

			// Creating Commands to trigger Runtime APIs

			// Create SetContext latest Command with input and output options
			setContextOneCmd, err := framework.NewSetContextCommand(setContextOneInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())
			setContextTwoCmd, err := framework.NewSetContextCommand(setContextTwoInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext latest Command with input and output options
			setCurrentContextCmd, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			getContextCmdForRuntime0254, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime100WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime100, getContextOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextTwoCmdForRuntime0254, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0254, getContextTwoOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getContextTwoCmdForRuntime100WithError, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime100, getContextTwoOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getContextTwoCmdForRuntime0280WithError, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0280, getContextTwoOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			getCurrentContextCmdForRuntime0254, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime100WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100WithError)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext Command
			deleteCtxCmd, err := framework.NewDeleteContextCommand(deleteContextInputOptions, deleteContextOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentContext Command
			removeCurrentCtxCmd, err := framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptions, removeCurrentContextOutputOptionsWithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextOneCmd).Add(setContextTwoCmd).Add(setCurrentContextCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100WithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100WithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(removeCurrentCtxCmd)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(deleteCtxCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100WithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntime100WithError).Add(getContextTwoCmdForRuntime0280WithError).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100WithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Execute the test case
			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Setting up the input and output parameters data for various APIs

			// Input Parameters for SetContext v0.28.0
			setContextOneInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			setContextTwoInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityTwo)

			// Input Parameters for SetCurrentContext v0.28.0
			setCurrentContextInputOptions := compatibilitytests.DefaultSetCurrentContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)

			// Input and Output Parameters for GetCurrentContext
			getCurrentContextInputOptionsForRuntime100 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.VersionLatest)
			getCurrentContextInputOptionsForRuntime0280 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.Version0280)
			getCurrentContextInputOptionsForRuntime0254 := compatibilitytests.DefaultGetCurrentContextInputOptions(core.Version0254)

			getCurrentContextOutputOptionsForRuntime100 := compatibilitytests.DefaultGetCurrentContextOutputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0280 := compatibilitytests.DefaultGetCurrentContextOutputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetCurrentContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)
			getCurrentContextOutputOptionsForRuntime0254WithError := compatibilitytests.DefaultGetCurrentContextOutputOptionsWithError(core.Version0254)

			// Input and Output params for GetContext
			getContextInputOptionsForRuntime100 := compatibilitytests.DefaultGetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextInputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			getContextOutputOptionsForRuntime100 := compatibilitytests.DefaultGetContextOutputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			getContextTwoInputOptionsForRuntime100 := compatibilitytests.DefaultGetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoInputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoInputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityTwo)

			getContextTwoOutputOptionsForRuntime100 := compatibilitytests.DefaultGetContextOutputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoOutputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0280, compatibilitytests.CtxCompatibilityTwo)
			getContextTwoOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityTwo)
			// getContextTwoOutputOptionsForRuntime0254WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0254, compatibilitytests.CtxCompatibilityTwo)

			// Input params for DeleteContext v0.25.4
			deleteContextInputOptions := compatibilitytests.DefaultDeleteContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			// Creating Commands to trigger Runtime APIs

			// Create SetContext v0.28.0 Command with input and output options
			setContextOneCmd, err := framework.NewSetContextCommand(setContextOneInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())
			setContextTwoCmd, err := framework.NewSetContextCommand(setContextTwoInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext v0.28.0 Command with input and output options
			setCurrentContextCmd, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			getContextCmdForRuntime100, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime100, getContextOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0280, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0280, getContextOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCmdForRuntime0254WithError, err := framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			getContextTwoCmdForRuntime100, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime100, getContextTwoOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getContextTwoCmdForRuntime0280, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0280, getContextTwoOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getContextTwoCmdForRuntime0254, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0254, getContextTwoOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			//getContextTwoCmdForRuntime0254WithError, err := framework.NewGetContextCommand(getContextTwoInputOptionsForRuntime0254, getContextTwoOutputOptionsForRuntime0254WithError)
			//gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			getCurrentContextCmdForRuntime100, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime100, getCurrentContextOutputOptionsForRuntime100)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0280, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0280, getCurrentContextOutputOptionsForRuntime0280)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254)
			gomega.Expect(err).To(gomega.BeNil())
			getCurrentContextCmdForRuntime0254WithError, err := framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254WithError)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext v0.25.4 Command
			deleteCtxCmd, err := framework.NewDeleteContextCommand(deleteContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Build test case with commands

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextOneCmd).Add(setContextTwoCmd).Add(setCurrentContextCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntime100).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add DeleteContext v0.25.4 Command
			testCase.Add(deleteCtxCmd)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntime100).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254WithError)
			testCase.Add(getContextTwoCmdForRuntime100).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntime100).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	// TODO: Add tests with different context targets global, tmc

})
