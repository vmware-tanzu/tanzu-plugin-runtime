// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package context_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests"
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

	ginkgo.Context("For SetContext, SetCurrentContext, GetContext, GetCurrentContext, DeleteContext, RemoveCurrentContext on respective Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Input and Output Parameters for SetContext latest
			setContextInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)

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
			getContextInputOptionsForRuntime100 := compatibilitytests.DefaultGetContextInputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime100 := compatibilitytests.DefaultGetContextOutputOptions(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime100WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.VersionLatest, compatibilitytests.CtxCompatibilityOne)

			getContextInputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0280WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0280, compatibilitytests.CtxCompatibilityOne)

			getContextInputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254 := compatibilitytests.DefaultGetContextOutputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)
			getContextOutputOptionsForRuntime0254WithError := compatibilitytests.DefaultGetContextOutputOptionsWithError(core.Version0254, compatibilitytests.CtxCompatibilityOne)

			// Input params for DeleteContext v0.28.0
			deleteContextInputOptions := compatibilitytests.DefaultDeleteContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)

			// Input params for RemoveCurrentContext v0.28.0
			removeCurrentContextInputOptions := compatibilitytests.DefaultRemoveCurrentContextInputOptions(core.Version0280)

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
			setContextInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

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

			// Input and Output params for RemoveCurrentContext v0.28.0
			removeCurrentContextInputOptions := compatibilitytests.DefaultRemoveCurrentContextInputOptions(core.Version0280)
			removeCurrentContextOutputOptionsWithError := compatibilitytests.DefaultRemoveCurrentContextOutputOptionsWithError(core.Version0280)

			// Input and Output params for DeleteContext v0.28.0
			deleteContextInputOptions := compatibilitytests.DefaultDeleteContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)
			deleteContextOutputOptionsWithError := compatibilitytests.DefaultDeleteContextOutputOptionsWithError(core.Version0280, compatibilitytests.CtxCompatibilityOne)

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
			setContextInputOptions := compatibilitytests.DefaultSetContextInputOptions(core.Version0280, compatibilitytests.CtxCompatibilityOne)

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

			// Input params for DeleteContext v0.25.4
			deleteContextInputOptions := compatibilitytests.DefaultDeleteContextInputOptions(core.Version0254, compatibilitytests.CtxCompatibilityOne)

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

	// TODO: Add tests to set multiple contexts
	// TODO: Add tests with different context targets global, tmc
})
