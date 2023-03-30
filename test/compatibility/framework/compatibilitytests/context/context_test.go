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
	var getContextOutputOptionsForRuntime0254WithError *framework.GetContextOutputOptions
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
	var getCurrentContextOutputOptionsForRuntime0254WithError *framework.GetCurrentContextOutputOptions

	// DeleteContext Input Options
	var deleteContextInputOptionsForRuntime0254 *framework.DeleteContextInputOptions
	var deleteContextInputOptionsForRuntime0280 *framework.DeleteContextInputOptions
	var deleteContextInputOptionsForRuntimeLatest *framework.DeleteContextInputOptions

	// DeleteContext Output Options with expected error
	var deleteContextOutputOptionsForRuntime0280WithError *framework.DeleteContextOutputOptions
	var deleteContextOutputOptionsForRuntimeLatestWithError *framework.DeleteContextOutputOptions

	// RemoveCurrentContext Input Options
	var removeCurrentContextInputOptionsForRuntime0280 *framework.RemoveCurrentContextInputOptions
	var removeCurrentContextInputOptionsForRuntimeLatest *framework.RemoveCurrentContextInputOptions

	// RemoveCurrentContext Output Options with expected error
	var removeCurrentContextOutputOptionsForRuntimeLatestWithError *framework.RemoveCurrentContextOutputOptions
	var removeCurrentContextOutputOptionsForRuntime0280WithError *framework.RemoveCurrentContextOutputOptions

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
	var getContextCmdForRuntime0254WithError *core.Command

	var getContextTwoCmdForRuntimeLatestWithError *core.Command
	var getContextTwoCmdForRuntime0280WithError *core.Command

	// GetCurrentContext API Commands
	var getCurrentContextCmdForRuntimeLatest *core.Command
	var getCurrentContextCmdForRuntime0280 *core.Command
	var getCurrentContextCmdForRuntime0254 *core.Command

	var getCurrentContextCmdForRuntimeLatestWithError *core.Command
	var getCurrentContextCmdForRuntime0280WithError *core.Command
	var getCurrentContextCmdForRuntime0254WithError *core.Command

	// DeleteContext API Commands
	var deleteContextCmdForRuntime0280 *core.Command
	var deleteContextCmdForRuntime0254 *core.Command
	var deleteContextCmdForRuntimeLatest *core.Command

	var deleteContextCmdForRuntime0280WithError *core.Command
	var deleteContextCmdForRuntimeLatestWithError *core.Command

	// RemoveCurrentContext API Commands
	var removeCurrentContextCmdForRuntime0280 *core.Command
	var removeCurrentContextCmdForRuntimeLatest *core.Command

	var removeCurrentContextCmdForRuntimeLatestWithError *core.Command
	var removeCurrentContextCmdForRuntime0280WithError *core.Command

	var err error
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
		getCurrentContextOutputOptionsForRuntime0254WithError = context.DefaultGetCurrentContextOutputOptionsWithError(core.Version0254)

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
		getContextOutputOptionsForRuntime0254WithError = context.DefaultGetContextOutputOptionsWithError(core.Version0254, common.CtxCompatibilityOne)

		getContextTwoOutputOptionsForRuntimeLatestWithError = context.DefaultGetContextOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityTwo)
		getContextTwoOutputOptionsForRuntime0280WithError = context.DefaultGetContextOutputOptionsWithError(core.Version0280, common.CtxCompatibilityTwo)

		// Input and Output Options for DeleteContext
		deleteContextInputOptionsForRuntime0280 = context.DefaultDeleteContextInputOptions(core.Version0280, common.CtxCompatibilityOne)
		deleteContextInputOptionsForRuntime0254 = context.DefaultDeleteContextInputOptions(core.Version0254, common.CtxCompatibilityOne)
		deleteContextInputOptionsForRuntimeLatest = context.DefaultDeleteContextInputOptions(core.VersionLatest, common.CtxCompatibilityOne)

		deleteContextOutputOptionsForRuntime0280WithError = context.DefaultDeleteContextOutputOptionsWithError(core.Version0280, common.CtxCompatibilityOne)
		deleteContextOutputOptionsForRuntimeLatestWithError = context.DefaultDeleteContextOutputOptionsWithError(core.VersionLatest, common.CtxCompatibilityOne)

		// Input and Output Options for RemoveCurrentContext
		removeCurrentContextInputOptionsForRuntime0280 = context.DefaultRemoveCurrentContextInputOptions(core.Version0280)
		removeCurrentContextInputOptionsForRuntimeLatest = context.DefaultRemoveCurrentContextInputOptions(core.VersionLatest)

		removeCurrentContextOutputOptionsForRuntimeLatestWithError = context.DefaultRemoveCurrentContextOutputOptionsWithError(core.VersionLatest)
		removeCurrentContextOutputOptionsForRuntime0280WithError = context.DefaultRemoveCurrentContextOutputOptionsWithError(core.Version0280)

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
		getContextCmdForRuntime0254WithError, err = framework.NewGetContextCommand(getContextInputOptionsForRuntime0254, getContextOutputOptionsForRuntime0254WithError)
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
		getCurrentContextCmdForRuntime0254WithError, err = framework.NewGetCurrentContextCommand(getCurrentContextInputOptionsForRuntime0254, getCurrentContextOutputOptionsForRuntime0254WithError)
		gomega.Expect(err).To(gomega.BeNil())

		// Create DeleteContext Commands with input and output options
		deleteContextCmdForRuntimeLatest, err = framework.NewDeleteContextCommand(deleteContextInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())
		deleteContextCmdForRuntime0280, err = framework.NewDeleteContextCommand(deleteContextInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		deleteContextCmdForRuntime0254, err = framework.NewDeleteContextCommand(deleteContextInputOptionsForRuntime0254, nil)
		gomega.Expect(err).To(gomega.BeNil())

		deleteContextCmdForRuntime0280WithError, err = framework.NewDeleteContextCommand(deleteContextInputOptionsForRuntime0280, deleteContextOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
		deleteContextCmdForRuntimeLatestWithError, err = framework.NewDeleteContextCommand(deleteContextInputOptionsForRuntimeLatest, deleteContextOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())

		// Create RemoveCurrentContext Commands with input and output options
		removeCurrentContextCmdForRuntime0280, err = framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptionsForRuntime0280, nil)
		gomega.Expect(err).To(gomega.BeNil())
		removeCurrentContextCmdForRuntimeLatest, err = framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptionsForRuntimeLatest, nil)
		gomega.Expect(err).To(gomega.BeNil())

		removeCurrentContextCmdForRuntimeLatestWithError, err = framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptionsForRuntimeLatest, removeCurrentContextOutputOptionsForRuntimeLatestWithError)
		gomega.Expect(err).To(gomega.BeNil())

		removeCurrentContextCmdForRuntime0280WithError, err = framework.NewRemoveCurrentContextCommand(removeCurrentContextInputOptionsForRuntime0280, removeCurrentContextOutputOptionsForRuntime0280WithError)
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.Context("using single context object on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime latest
			testCase := core.NewTestCase().Add(setContextCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(removeCurrentContextCmdForRuntime0280)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(deleteContextCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands for Runtime v0.28.0
			testCase := core.NewTestCase().Add(setContextCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext latest Command
			testCase.Add(removeCurrentContextCmdForRuntimeLatest)

			// Add DeleteContext latest Command
			testCase.Add(deleteContextCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext latestCommand
			testCase.Add(removeCurrentContextCmdForRuntimeLatestWithError)

			// Add DeleteContext latest Command
			testCase.Add(deleteContextCmdForRuntimeLatestWithError)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})
		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add DeleteContext v0.25.4 Command
			testCase.Add(deleteContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254WithError)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	ginkgo.Context("using multiple context objects on supported Runtime API versions", func() {

		ginkgo.It("Run SetContext, SetCurrentContext on Runtime latest then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntimeLatest).Add(setContextTwoCmdForRuntimeLatest).Add(setCurrentContextCmdForRuntimeLatest)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(removeCurrentContextCmdForRuntime0280)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(deleteContextCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254WithError)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {
			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0254).Add(setContextTwoCmdForRuntime0254).Add(setCurrentContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Add RemoveCurrentContext v0.28.0 Command
			testCase.Add(removeCurrentContextCmdForRuntime0280WithError)

			// Add DeleteContext v0.28.0 Command
			testCase.Add(deleteContextCmdForRuntime0280WithError)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatestWithError).Add(getContextCmdForRuntime0280WithError).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatestWithError).Add(getContextTwoCmdForRuntime0280WithError).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatestWithError).Add(getCurrentContextCmdForRuntime0280WithError).Add(getCurrentContextCmdForRuntime0254)

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest then DeleteContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, latest", func() {

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCmdForRuntime0280).Add(setContextTwoCmdForRuntime0280).Add(setCurrentContextCmdForRuntime0280)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254)

			// Add DeleteContext v0.25.4 Command
			testCase.Add(deleteContextCmdForRuntime0254)

			// Add GetContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getContextCmdForRuntimeLatest).Add(getContextCmdForRuntime0280).Add(getContextCmdForRuntime0254WithError)
			testCase.Add(getContextTwoCmdForRuntimeLatest).Add(getContextTwoCmdForRuntime0280).Add(getContextTwoCmdForRuntime0254)

			// Add GetCurrentContext latest, v0.28.0, v0.25.4 Commands
			testCase.Add(getCurrentContextCmdForRuntimeLatest).Add(getCurrentContextCmdForRuntime0280).Add(getCurrentContextCmdForRuntime0254WithError)

			// Run all the commands
			executer.Execute(testCase)
		})

	})
})
