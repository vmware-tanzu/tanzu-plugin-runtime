// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package context contains all the cross version api compatibility tests for context apis
package context

import (
	"fmt"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/context"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

// Helper struct provides input and output options and api commands to be used in test cases
type Helper struct {
	// SetContext Input Options
	SetContextInputOptionsForRuntime0254      *context.SetContextInputOptions
	SetContextInputOptionsForRuntime0280      *context.SetContextInputOptions
	SetContextInputOptionsForRuntimeLatest    *context.SetContextInputOptions
	SetContextTwoInputOptionsForRuntime0254   *context.SetContextInputOptions
	SetContextTwoInputOptionsForRuntime0280   *context.SetContextInputOptions
	SetContextTwoInputOptionsForRuntimeLatest *context.SetContextInputOptions

	// SetCurrentContext Input Options
	SetCurrentContextInputOptionsForRuntime0254   *context.SetCurrentContextInputOptions
	SetCurrentContextInputOptionsForRuntime0280   *context.SetCurrentContextInputOptions
	SetCurrentContextInputOptionsForRuntimeLatest *context.SetCurrentContextInputOptions

	// GetContext Input Options
	GetContextInputOptionsForRuntimeLatest    *context.GetContextInputOptions
	GetContextInputOptionsForRuntime0280      *context.GetContextInputOptions
	GetContextInputOptionsForRuntime0254      *context.GetContextInputOptions
	GetContextTwoInputOptionsForRuntimeLatest *context.GetContextInputOptions
	GetContextTwoInputOptionsForRuntime0280   *context.GetContextInputOptions
	GetContextTwoInputOptionsForRuntime0254   *context.GetContextInputOptions

	// GetContext Output Options
	GetContextOutputOptionsForRuntime0254      *context.GetContextOutputOptions
	GetContextOutputOptionsForRuntime0280      *context.GetContextOutputOptions
	GetContextOutputOptionsForRuntimeLatest    *context.GetContextOutputOptions
	GetContextTwoOutputOptionsForRuntime0254   *context.GetContextOutputOptions
	GetContextTwoOutputOptionsForRuntime0280   *context.GetContextOutputOptions
	GetContextTwoOutputOptionsForRuntimeLatest *context.GetContextOutputOptions

	// GetContext Output Options with expected error
	GetContextOutputOptionsForRuntimeLatestWithError    *context.GetContextOutputOptions
	GetContextOutputOptionsForRuntime0280WithError      *context.GetContextOutputOptions
	GetContextOutputOptionsForRuntime0254WithError      *context.GetContextOutputOptions
	GetContextTwoOutputOptionsForRuntimeLatestWithError *context.GetContextOutputOptions
	GetContextTwoOutputOptionsForRuntime0280WithError   *context.GetContextOutputOptions

	// GetCurrentContext Input Options
	GetCurrentContextInputOptionsForRuntime0254   *context.GetCurrentContextInputOptions
	GetCurrentContextInputOptionsForRuntime0280   *context.GetCurrentContextInputOptions
	GetCurrentContextInputOptionsForRuntimeLatest *context.GetCurrentContextInputOptions

	// GetCurrentContext Output Options
	GetCurrentContextOutputOptionsForRuntime0254   *context.GetCurrentContextOutputOptions
	GetCurrentContextOutputOptionsForRuntime0280   *context.GetCurrentContextOutputOptions
	GetCurrentContextOutputOptionsForRuntimeLatest *context.GetCurrentContextOutputOptions

	// GetCurrentContext Output Options with expected error
	GetCurrentContextOutputOptionsForRuntimeLatestWithError *context.GetCurrentContextOutputOptions
	GetCurrentContextOutputOptionsForRuntime0280WithError   *context.GetCurrentContextOutputOptions
	GetCurrentContextOutputOptionsForRuntime0254WithError   *context.GetCurrentContextOutputOptions

	// DeleteContext Input Options
	DeleteContextInputOptionsForRuntime0254   *context.DeleteContextInputOptions
	DeleteContextInputOptionsForRuntime0280   *context.DeleteContextInputOptions
	DeleteContextInputOptionsForRuntimeLatest *context.DeleteContextInputOptions

	// DeleteContext Output Options with expected error
	DeleteContextOutputOptionsForRuntime0280WithError   *context.DeleteContextOutputOptions
	DeleteContextOutputOptionsForRuntimeLatestWithError *context.DeleteContextOutputOptions

	// RemoveCurrentContext Input Options
	RemoveCurrentContextInputOptionsForRuntime0280   *context.RemoveCurrentContextInputOptions
	RemoveCurrentContextInputOptionsForRuntimeLatest *context.RemoveCurrentContextInputOptions

	// RemoveCurrentContext Output Options with expected error
	RemoveCurrentContextOutputOptionsForRuntimeLatestWithError *context.RemoveCurrentContextOutputOptions
	RemoveCurrentContextOutputOptionsForRuntime0280WithError   *context.RemoveCurrentContextOutputOptions

	// Context API Commands
	// SetContext API Commands
	SetContextCmdForRuntimeLatest *core.Command
	SetContextCmdForRuntime0280   *core.Command
	SetContextCmdForRuntime0254   *core.Command

	SetContextTwoCmdForRuntimeLatest *core.Command
	SetContextTwoCmdForRuntime0280   *core.Command
	SetContextTwoCmdForRuntime0254   *core.Command

	// SetCurrentContext API Commands
	SetCurrentContextCmdForRuntime0254   *core.Command
	SetCurrentContextCmdForRuntime0280   *core.Command
	SetCurrentContextCmdForRuntimeLatest *core.Command

	// GetContext API Commands
	GetContextCmdForRuntimeLatest *core.Command
	GetContextCmdForRuntime0280   *core.Command
	GetContextCmdForRuntime0254   *core.Command

	GetContextTwoCmdForRuntimeLatest *core.Command
	GetContextTwoCmdForRuntime0280   *core.Command
	GetContextTwoCmdForRuntime0254   *core.Command

	GetContextCmdForRuntimeLatestWithError *core.Command
	GetContextCmdForRuntime0280WithError   *core.Command
	GetContextCmdForRuntime0254WithError   *core.Command

	GetContextTwoCmdForRuntimeLatestWithError *core.Command
	GetContextTwoCmdForRuntime0280WithError   *core.Command

	// GetCurrentContext API Commands
	GetCurrentContextCmdForRuntimeLatest *core.Command
	GetCurrentContextCmdForRuntime0280   *core.Command
	GetCurrentContextCmdForRuntime0254   *core.Command

	GetCurrentContextCmdForRuntimeLatestWithError *core.Command
	GetCurrentContextCmdForRuntime0280WithError   *core.Command
	GetCurrentContextCmdForRuntime0254WithError   *core.Command

	// DeleteContext API Commands
	DeleteContextCmdForRuntime0280   *core.Command
	DeleteContextCmdForRuntime0254   *core.Command
	DeleteContextCmdForRuntimeLatest *core.Command

	DeleteContextCmdForRuntime0280WithError   *core.Command
	DeleteContextCmdForRuntimeLatestWithError *core.Command

	// RemoveCurrentContext API Commands
	RemoveCurrentContextCmdForRuntime0280   *core.Command
	RemoveCurrentContextCmdForRuntimeLatest *core.Command

	RemoveCurrentContextCmdForRuntimeLatestWithError *core.Command
	RemoveCurrentContextCmdForRuntime0280WithError   *core.Command
}

// SetUpDefaultData sets up the Helper struct with default input/output options and api commands
func (b *Helper) SetUpDefaultData() {
	b.SetupSetContextTestInputAndOutputOptions()
	b.CreateSetContextAPICommands()

	b.SetupSetCurrentContextTestInputAndOutputOptions()
	b.CreateSetCurrentContextAPICommands()

	b.SetupGetCurrentContextTestInputAndOutputOptions()
	b.CreateGetCurrentContextAPICommands()

	b.SetupGetContextTestInputAndOutputOptions()
	b.CreateGetContextAPICommands()

	b.SetupDeleteContextTestInputAndOutputOptions()
	b.CreateDeleteContextAPICommands()

	b.SetupRemoveCurrentContextTestInputAndOutputOptions()
	b.CreateRemoveCurrentContextAPICommands()
}

// SetupRemoveCurrentContextTestInputAndOutputOptions sets input and output options for RemoveCurrentContext API
func (b *Helper) SetupRemoveCurrentContextTestInputAndOutputOptions() {
	// Input and Output Options for RemoveCurrentContext
	ginkgo.By("Setup Input and Output Options for RemoveCurrentContext")

	b.RemoveCurrentContextInputOptionsForRuntime0280 = DefaultRemoveCurrentContextInputOptions(core.Version0280)
	b.RemoveCurrentContextInputOptionsForRuntimeLatest = DefaultRemoveCurrentContextInputOptions(core.VersionLatest)

	b.RemoveCurrentContextOutputOptionsForRuntimeLatestWithError = DefaultRemoveCurrentContextOutputOptionsWithError(core.VersionLatest)
	b.RemoveCurrentContextOutputOptionsForRuntime0280WithError = DefaultRemoveCurrentContextOutputOptionsWithError(core.Version0280)
}

// CreateRemoveCurrentContextAPICommands sets api commands for RemoveCurrentContext API
func (b *Helper) CreateRemoveCurrentContextAPICommands() {
	// Create RemoveCurrentContext Commands with input and output options
	ginkgo.By("Create RemoveCurrentContext API Commands")

	removeCurrentContextCmdForRuntime0280, err := context.NewRemoveCurrentContextCommand(b.RemoveCurrentContextInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentContextCmdForRuntime0280 = removeCurrentContextCmdForRuntime0280

	removeCurrentContextCmdForRuntimeLatest, err := context.NewRemoveCurrentContextCommand(b.RemoveCurrentContextInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentContextCmdForRuntimeLatest = removeCurrentContextCmdForRuntimeLatest

	removeCurrentContextCmdForRuntimeLatestWithError, err := context.NewRemoveCurrentContextCommand(b.RemoveCurrentContextInputOptionsForRuntimeLatest, b.RemoveCurrentContextOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentContextCmdForRuntimeLatestWithError = removeCurrentContextCmdForRuntimeLatestWithError

	removeCurrentContextCmdForRuntime0280WithError, err := context.NewRemoveCurrentContextCommand(b.RemoveCurrentContextInputOptionsForRuntime0280, b.RemoveCurrentContextOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentContextCmdForRuntime0280WithError = removeCurrentContextCmdForRuntime0280WithError
}

// SetupDeleteContextTestInputAndOutputOptions sets input and output options for DeleteContext API
func (b *Helper) SetupDeleteContextTestInputAndOutputOptions() {
	// Input and Output Options for DeleteContext
	ginkgo.By("Setup Input and Output Options for DeleteContext")

	b.DeleteContextInputOptionsForRuntime0280 = DefaultDeleteContextInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.DeleteContextInputOptionsForRuntime0254 = DefaultDeleteContextInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.DeleteContextInputOptionsForRuntimeLatest = DefaultDeleteContextInputOptions(core.VersionLatest, common.CompatibilityTestOne)

	b.DeleteContextOutputOptionsForRuntime0280WithError = DefaultDeleteContextOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
	b.DeleteContextOutputOptionsForRuntimeLatestWithError = DefaultDeleteContextOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
}

// CreateDeleteContextAPICommands sets api commands for DeleteContext API
func (b *Helper) CreateDeleteContextAPICommands() {
	// Create DeleteContext Commands with input and output options
	ginkgo.By("Create DeleteContext API Commands")

	deleteContextCmdForRuntimeLatest, err := context.NewDeleteContextCommand(b.DeleteContextInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteContextCmdForRuntimeLatest = deleteContextCmdForRuntimeLatest

	deleteContextCmdForRuntime0280, err := context.NewDeleteContextCommand(b.DeleteContextInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteContextCmdForRuntime0280 = deleteContextCmdForRuntime0280

	deleteContextCmdForRuntime0254, err := context.NewDeleteContextCommand(b.DeleteContextInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteContextCmdForRuntime0254 = deleteContextCmdForRuntime0254

	deleteContextCmdForRuntime0280WithError, err := context.NewDeleteContextCommand(b.DeleteContextInputOptionsForRuntime0280, b.DeleteContextOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteContextCmdForRuntime0280WithError = deleteContextCmdForRuntime0280WithError

	deleteContextCmdForRuntimeLatestWithError, err := context.NewDeleteContextCommand(b.DeleteContextInputOptionsForRuntimeLatest, b.DeleteContextOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteContextCmdForRuntimeLatestWithError = deleteContextCmdForRuntimeLatestWithError
}

// SetupGetContextTestInputAndOutputOptions sets input and output options for GetContext API
func (b *Helper) SetupGetContextTestInputAndOutputOptions() {
	// Input and Output params for GetContext
	ginkgo.By("Setup Input and Output Options for GetContext")

	b.GetContextInputOptionsForRuntimeLatest = DefaultGetContextInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.GetContextInputOptionsForRuntime0280 = DefaultGetContextInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetContextInputOptionsForRuntime0254 = DefaultGetContextInputOptions(core.Version0254, common.CompatibilityTestOne)

	b.GetContextTwoInputOptionsForRuntimeLatest = DefaultGetContextInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetContextTwoInputOptionsForRuntime0280 = DefaultGetContextInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.GetContextTwoInputOptionsForRuntime0254 = DefaultGetContextInputOptions(core.Version0254, common.CompatibilityTestTwo)

	b.GetContextOutputOptionsForRuntime0280 = DefaultGetContextOutputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetContextOutputOptionsForRuntime0254 = DefaultGetContextOutputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetContextOutputOptionsForRuntimeLatest = DefaultGetContextOutputOptions(core.VersionLatest, common.CompatibilityTestOne)

	b.GetContextTwoOutputOptionsForRuntimeLatest = DefaultGetContextOutputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetContextTwoOutputOptionsForRuntime0280 = DefaultGetContextOutputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.GetContextTwoOutputOptionsForRuntime0254 = DefaultGetContextOutputOptions(core.Version0254, common.CompatibilityTestTwo)

	b.GetContextOutputOptionsForRuntimeLatestWithError = DefaultGetContextOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
	b.GetContextOutputOptionsForRuntime0280WithError = DefaultGetContextOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
	b.GetContextOutputOptionsForRuntime0254WithError = DefaultGetContextOutputOptionsWithError(core.Version0254, common.CompatibilityTestOne)

	b.GetContextTwoOutputOptionsForRuntimeLatestWithError = DefaultGetContextOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetContextTwoOutputOptionsForRuntime0280WithError = DefaultGetContextOutputOptionsWithError(core.Version0280, common.CompatibilityTestTwo)
}

// CreateGetContextAPICommands sets api commands for GetContext API
func (b *Helper) CreateGetContextAPICommands() {
	// Create GetContext Commands with input and output options
	ginkgo.By("Create GetContext API Commands")

	getContextCmdForRuntimeLatest, err := context.NewGetContextCommand(b.GetContextInputOptionsForRuntimeLatest, b.GetContextOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextCmdForRuntimeLatest = getContextCmdForRuntimeLatest

	getContextCmdForRuntime0280, err := context.NewGetContextCommand(b.GetContextInputOptionsForRuntime0280, b.GetContextOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextCmdForRuntime0280 = getContextCmdForRuntime0280

	getContextCmdForRuntime0254, err := context.NewGetContextCommand(b.GetContextInputOptionsForRuntime0254, b.GetContextOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextCmdForRuntime0254 = getContextCmdForRuntime0254

	getContextTwoCmdForRuntimeLatest, err := context.NewGetContextCommand(b.GetContextTwoInputOptionsForRuntimeLatest, b.GetContextTwoOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextTwoCmdForRuntimeLatest = getContextTwoCmdForRuntimeLatest

	getContextTwoCmdForRuntime0280, err := context.NewGetContextCommand(b.GetContextTwoInputOptionsForRuntime0280, b.GetContextTwoOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextTwoCmdForRuntime0280 = getContextTwoCmdForRuntime0280

	getContextTwoCmdForRuntime0254, err := context.NewGetContextCommand(b.GetContextTwoInputOptionsForRuntime0254, b.GetContextTwoOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextTwoCmdForRuntime0254 = getContextTwoCmdForRuntime0254

	getContextCmdForRuntimeLatestWithError, err := context.NewGetContextCommand(b.GetContextInputOptionsForRuntimeLatest, b.GetContextOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextCmdForRuntimeLatestWithError = getContextCmdForRuntimeLatestWithError

	getContextCmdForRuntime0280WithError, err := context.NewGetContextCommand(b.GetContextInputOptionsForRuntime0280, b.GetContextOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextCmdForRuntime0280WithError = getContextCmdForRuntime0280WithError

	getContextCmdForRuntime0254WithError, err := context.NewGetContextCommand(b.GetContextInputOptionsForRuntime0254, b.GetContextOutputOptionsForRuntime0254WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextCmdForRuntime0254WithError = getContextCmdForRuntime0254WithError

	getContextTwoCmdForRuntimeLatestWithError, err := context.NewGetContextCommand(b.GetContextTwoInputOptionsForRuntimeLatest, b.GetContextTwoOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextTwoCmdForRuntimeLatestWithError = getContextTwoCmdForRuntimeLatestWithError

	getContextTwoCmdForRuntime0280WithError, err := context.NewGetContextCommand(b.GetContextTwoInputOptionsForRuntime0280, b.GetContextTwoOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetContextTwoCmdForRuntime0280WithError = getContextTwoCmdForRuntime0280WithError
}

// SetupGetCurrentContextTestInputAndOutputOptions sets input and output options for GetCurrentContext API
func (b *Helper) SetupGetCurrentContextTestInputAndOutputOptions() {
	// Input and Output Parameters for GetCurrentContext
	ginkgo.By("Setup Input and Output Options for GetCurrentContext")

	b.GetCurrentContextInputOptionsForRuntimeLatest = DefaultGetCurrentContextInputOptions(core.VersionLatest)
	b.GetCurrentContextInputOptionsForRuntime0280 = DefaultGetCurrentContextInputOptions(core.Version0280)
	b.GetCurrentContextInputOptionsForRuntime0254 = DefaultGetCurrentContextInputOptions(core.Version0254)

	b.GetCurrentContextOutputOptionsForRuntime0280 = DefaultGetCurrentContextOutputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetCurrentContextOutputOptionsForRuntime0254 = DefaultGetCurrentContextOutputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetCurrentContextOutputOptionsForRuntimeLatest = DefaultGetCurrentContextOutputOptions(core.VersionLatest, common.CompatibilityTestOne)

	b.GetCurrentContextOutputOptionsForRuntimeLatestWithError = DefaultGetCurrentContextOutputOptionsWithError(core.VersionLatest)
	b.GetCurrentContextOutputOptionsForRuntime0280WithError = DefaultGetCurrentContextOutputOptionsWithError(core.Version0280)
	b.GetCurrentContextOutputOptionsForRuntime0254WithError = DefaultGetCurrentContextOutputOptionsWithError(core.Version0254)
}

// CreateGetCurrentContextAPICommands sets api commands for GetCurrentContext API
func (b *Helper) CreateGetCurrentContextAPICommands() {
	// Create GetCurrentContext Commands with input and output options
	ginkgo.By("Create GetCurrentContext API Commands")

	getCurrentContextCmdForRuntimeLatest, err := context.NewGetCurrentContextCommand(b.GetCurrentContextInputOptionsForRuntimeLatest, b.GetCurrentContextOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentContextCmdForRuntimeLatest = getCurrentContextCmdForRuntimeLatest

	getCurrentContextCmdForRuntime0280, err := context.NewGetCurrentContextCommand(b.GetCurrentContextInputOptionsForRuntime0280, b.GetCurrentContextOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentContextCmdForRuntime0280 = getCurrentContextCmdForRuntime0280

	getCurrentContextCmdForRuntime0254, err := context.NewGetCurrentContextCommand(b.GetCurrentContextInputOptionsForRuntime0254, b.GetCurrentContextOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentContextCmdForRuntime0254 = getCurrentContextCmdForRuntime0254

	getCurrentContextCmdForRuntimeLatestWithError, err := context.NewGetCurrentContextCommand(b.GetCurrentContextInputOptionsForRuntimeLatest, b.GetCurrentContextOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentContextCmdForRuntimeLatestWithError = getCurrentContextCmdForRuntimeLatestWithError

	getCurrentContextCmdForRuntime0280WithError, err := context.NewGetCurrentContextCommand(b.GetCurrentContextInputOptionsForRuntime0280, b.GetCurrentContextOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentContextCmdForRuntime0280WithError = getCurrentContextCmdForRuntime0280WithError

	getCurrentContextCmdForRuntime0254WithError, err := context.NewGetCurrentContextCommand(b.GetCurrentContextInputOptionsForRuntime0254, b.GetCurrentContextOutputOptionsForRuntime0254WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentContextCmdForRuntime0254WithError = getCurrentContextCmdForRuntime0254WithError
}

// SetupSetCurrentContextTestInputAndOutputOptions sets input and output options for SetCurrentContext API
func (b *Helper) SetupSetCurrentContextTestInputAndOutputOptions() {
	// Input and Output Parameters for SetCurrentContext
	ginkgo.By("Setup Input and Output Options for SetCurrentContext")

	b.SetCurrentContextInputOptionsForRuntimeLatest = DefaultSetCurrentContextInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.SetCurrentContextInputOptionsForRuntime0280 = DefaultSetCurrentContextInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetCurrentContextInputOptionsForRuntime0254 = DefaultSetCurrentContextInputOptions(core.Version0254, common.CompatibilityTestOne)
}

// CreateSetCurrentContextAPICommands sets api commands for SetCurrentContext API
func (b *Helper) CreateSetCurrentContextAPICommands() {
	// Create SetCurrentContext Commands with input and output options
	ginkgo.By("Create SetCurrentContext API Commands")

	setCurrentContextCmdForRuntimeLatest, err := context.NewSetCurrentContextCommand(b.SetCurrentContextInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentContextCmdForRuntimeLatest = setCurrentContextCmdForRuntimeLatest

	setCurrentContextCmdForRuntime0280, err := context.NewSetCurrentContextCommand(b.SetCurrentContextInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentContextCmdForRuntime0280 = setCurrentContextCmdForRuntime0280

	setCurrentContextCmdForRuntime0254, err := context.NewSetCurrentContextCommand(b.SetCurrentContextInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentContextCmdForRuntime0254 = setCurrentContextCmdForRuntime0254
}

// SetupSetContextTestInputAndOutputOptions sets input and output options for SetContext API
func (b *Helper) SetupSetContextTestInputAndOutputOptions() {
	// Input and Output Parameters for SetContext
	ginkgo.By("Setup Input and Output Options for SetContext")

	b.SetContextInputOptionsForRuntimeLatest = DefaultSetContextInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.SetContextInputOptionsForRuntime0280 = DefaultSetContextInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetContextInputOptionsForRuntime0254 = DefaultSetContextInputOptions(core.Version0254, common.CompatibilityTestOne)

	b.SetContextTwoInputOptionsForRuntimeLatest = DefaultSetContextInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.SetContextTwoInputOptionsForRuntime0280 = DefaultSetContextInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.SetContextTwoInputOptionsForRuntime0254 = DefaultSetContextInputOptions(core.Version0254, common.CompatibilityTestTwo)

	// Input and Output Parameters for SetContext
	b.SetContextInputOptionsForRuntimeLatest = DefaultSetContextInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.SetContextInputOptionsForRuntime0280 = DefaultSetContextInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetContextInputOptionsForRuntime0254 = DefaultSetContextInputOptions(core.Version0254, common.CompatibilityTestOne)

	b.SetContextTwoInputOptionsForRuntimeLatest = DefaultSetContextInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.SetContextTwoInputOptionsForRuntime0280 = DefaultSetContextInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.SetContextTwoInputOptionsForRuntime0254 = DefaultSetContextInputOptions(core.Version0254, common.CompatibilityTestTwo)
}

// CreateSetContextAPICommands sets api commands for SetContext API
func (b *Helper) CreateSetContextAPICommands() {
	// Create SetContext Commands with input and output options
	ginkgo.By("Create SetContext API Commands")

	setContextCmdForRuntimeLatest, err := context.NewSetContextCommand(b.SetContextInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetContextCmdForRuntimeLatest = setContextCmdForRuntimeLatest
	setContextCmdForRuntime0254, err := context.NewSetContextCommand(b.SetContextInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetContextCmdForRuntime0254 = setContextCmdForRuntime0254
	setContextCmdForRuntime0280, err := context.NewSetContextCommand(b.SetContextInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetContextCmdForRuntime0280 = setContextCmdForRuntime0280
	setContextTwoCmdForRuntimeLatest, err := context.NewSetContextCommand(b.SetContextTwoInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetContextTwoCmdForRuntimeLatest = setContextTwoCmdForRuntimeLatest
	setContextTwoCmdForRuntime0254, err := context.NewSetContextCommand(b.SetContextTwoInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetContextTwoCmdForRuntime0254 = setContextTwoCmdForRuntime0254
	setContextTwoCmdForRuntime0280, err := context.NewSetContextCommand(b.SetContextTwoInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetContextTwoCmdForRuntime0280 = setContextTwoCmdForRuntime0280
}

// DefaultSetContextInputOptions helper method to construct SetContext API input options
func DefaultSetContextInputOptions(version core.RuntimeVersion, contextName string) *context.SetContextInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.SetContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ContextOpts: &types.ContextOpts{
				Name:   contextName,
				Target: types.TargetK8s,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: "default-compatibility-test-endpoint",
				},
			},
		}
	case core.Version0254:
		return &context.SetContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0254,
			},
			ContextOpts: &types.ContextOpts{
				Name: contextName,
				Type: types.CtxTypeK8s,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: "default-compatibility-test-endpoint",
				},
			},
		}
	}
	return nil
}

// DefaultGetContextInputOptions helper method to construct GetContext API input options
func DefaultGetContextInputOptions(version core.RuntimeVersion, contextName string) *context.GetContextInputOptions {
	return &context.GetContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ContextName: contextName,
	}
}

// DefaultGetContextOutputOptions helper method to construct GetContext API output options
func DefaultGetContextOutputOptions(version core.RuntimeVersion, contextName string) *context.GetContextOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.GetContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ContextOpts: &types.ContextOpts{
				Name:   contextName,
				Target: types.TargetK8s,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0254:
		return &context.GetContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0254,
			},
			ContextOpts: &types.ContextOpts{
				Name: contextName,
				Type: types.CtxTypeK8s,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}
	return nil
}

// DefaultGetContextOutputOptionsWithError helper method to construct GetContext API output options with error
func DefaultGetContextOutputOptionsWithError(version core.RuntimeVersion, contextName string) *context.GetContextOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.GetContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", contextName),
		}
	case core.Version0254:
		return &context.GetContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0254,
			},
			Error: fmt.Sprintf("could not find context \"%v\"", contextName),
		}
	}
	return nil
}

// DefaultSetCurrentContextInputOptions helper method to construct SetCurrentContext API input options
func DefaultSetCurrentContextInputOptions(version core.RuntimeVersion, contextName string) *context.SetCurrentContextInputOptions {
	return &context.SetCurrentContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ContextName: contextName,
	}
}

// DefaultGetCurrentContextInputOptions helper method to construct GetCurrentContext API input options
func DefaultGetCurrentContextInputOptions(version core.RuntimeVersion) *context.GetCurrentContextInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.GetCurrentContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Target: types.TargetK8s,
		}
	case core.Version0254:
		return &context.GetCurrentContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0254,
			},
			ContextType: types.CtxTypeK8s,
		}
	}
	return nil
}

// DefaultGetCurrentContextOutputOptions helper method to construct GetCurrentContext API output options
func DefaultGetCurrentContextOutputOptions(version core.RuntimeVersion, contextName string) *context.GetCurrentContextOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.GetCurrentContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.VersionLatest,
			},
			ContextOpts: &types.ContextOpts{
				Name:   contextName,
				Target: types.TargetK8s,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0254:
		return &context.GetCurrentContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0254,
			},
			ContextOpts: &types.ContextOpts{
				Name: contextName,
				Type: types.CtxTypeK8s,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}
	return nil
}

// DefaultGetCurrentContextOutputOptionsWithError helper method to construct GetCurrentContext API output options with error
func DefaultGetCurrentContextOutputOptionsWithError(version core.RuntimeVersion) *context.GetCurrentContextOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.GetCurrentContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("no current context set for target \"%v\"", types.TargetK8s),
		}
	case core.Version0254:
		return &context.GetCurrentContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("no current context set for type \"%v\"", types.CtxTypeK8s),
		}
	}
	return nil
}

// DefaultRemoveCurrentContextInputOptions helper method to construct RemoveCurrentContext API input options
func DefaultRemoveCurrentContextInputOptions(version core.RuntimeVersion) *context.RemoveCurrentContextInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.RemoveCurrentContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Target: types.TargetK8s,
		}
	}
	return nil
}

// DefaultRemoveCurrentContextOutputOptionsWithError helper method to construct RemoveCurrentContext API output option
func DefaultRemoveCurrentContextOutputOptionsWithError(version core.RuntimeVersion) *context.RemoveCurrentContextOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.RemoveCurrentContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("no current context set for target \"%v\"", types.TargetK8s),
		}
	case core.Version0254:
		return &context.RemoveCurrentContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("no current context set for type \"%v\"", types.CtxTypeK8s),
		}
	}
	return nil
}

// DefaultDeleteContextInputOptions helper method to construct DeleteContext API input options
func DefaultDeleteContextInputOptions(version core.RuntimeVersion, contextName string) *context.DeleteContextInputOptions {
	return &context.DeleteContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ContextName: contextName,
	}
}

// DefaultDeleteContextOutputOptionsWithError helper method to construct DeleteContext API output options
func DefaultDeleteContextOutputOptionsWithError(version core.RuntimeVersion, contextName string) *context.DeleteContextOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &context.DeleteContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", contextName),
		}
	case core.Version0254:
		return &context.DeleteContextOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0254,
			},
			Error: fmt.Sprintf("could not find context \"%v\"", contextName),
		}
	}
	return nil
}
