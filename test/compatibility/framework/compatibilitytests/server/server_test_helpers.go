// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package server provides cross-version Server API compatibility tests
package server

import (
	"fmt"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
)

const ServerNotFound = "current server \"\" not found in tanzu config"

// Helper struct provides input and output options and api commands to be used in test cases
type Helper struct {
	// SetServer Input Options
	SetServerInputOptionsForRuntime0116   *framework.SetServerInputOptions
	SetServerInputOptionsForRuntime0254   *framework.SetServerInputOptions
	SetServerInputOptionsForRuntime0280   *framework.SetServerInputOptions
	SetServerInputOptionsForRuntimeLatest *framework.SetServerInputOptions

	SetServerTwoInputOptionsForRuntime0116   *framework.SetServerInputOptions
	SetServerTwoInputOptionsForRuntime0254   *framework.SetServerInputOptions
	SetServerTwoInputOptionsForRuntime0280   *framework.SetServerInputOptions
	SetServerTwoInputOptionsForRuntimeLatest *framework.SetServerInputOptions

	// SetCurrentServer Input Options
	SetCurrentServerInputOptionsForRuntime0116   *framework.SetCurrentServerInputOptions
	SetCurrentServerInputOptionsForRuntime0254   *framework.SetCurrentServerInputOptions
	SetCurrentServerInputOptionsForRuntime0280   *framework.SetCurrentServerInputOptions
	SetCurrentServerInputOptionsForRuntimeLatest *framework.SetCurrentServerInputOptions

	// GetServer Input Options
	GetServerInputOptionsForRuntimeLatest *framework.GetServerInputOptions
	GetServerInputOptionsForRuntime0280   *framework.GetServerInputOptions
	GetServerInputOptionsForRuntime0254   *framework.GetServerInputOptions
	GetServerInputOptionsForRuntime0116   *framework.GetServerInputOptions

	GetServerTwoInputOptionsForRuntimeLatest *framework.GetServerInputOptions
	GetServerTwoInputOptionsForRuntime0280   *framework.GetServerInputOptions
	GetServerTwoInputOptionsForRuntime0254   *framework.GetServerInputOptions
	GetServerTwoInputOptionsForRuntime0116   *framework.GetServerInputOptions

	// GetServer Output Options
	GetServerOutputOptionsForRuntime0116   *framework.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0254   *framework.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0280   *framework.GetServerOutputOptions
	GetServerOutputOptionsForRuntimeLatest *framework.GetServerOutputOptions

	GetServerTwoOutputOptionsForRuntime0116   *framework.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime0254   *framework.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime0280   *framework.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntimeLatest *framework.GetServerOutputOptions

	// GetServer Output Options with expected error
	GetServerOutputOptionsForRuntimeLatestWithError *framework.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0280WithError   *framework.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0254WithError   *framework.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0116WithError   *framework.GetServerOutputOptions

	GetServerTwoOutputOptionsForRuntimeLatestWithError *framework.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime0280WithError   *framework.GetServerOutputOptions

	// GetCurrentServer Input Options
	GetCurrentServerInputOptionsForRuntime0116   *framework.GetCurrentServerInputOptions
	GetCurrentServerInputOptionsForRuntime0254   *framework.GetCurrentServerInputOptions
	GetCurrentServerInputOptionsForRuntime0280   *framework.GetCurrentServerInputOptions
	GetCurrentServerInputOptionsForRuntimeLatest *framework.GetCurrentServerInputOptions

	// GetCurrentServer Output Options
	GetCurrentServerOutputOptionsForRuntime0116   *framework.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0254   *framework.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0280   *framework.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntimeLatest *framework.GetCurrentServerOutputOptions

	// GetCurrentServer Output Options with expected error
	GetCurrentServerOutputOptionsForRuntimeLatestWithError *framework.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0280WithError   *framework.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0254WithError   *framework.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0116WithError   *framework.GetCurrentServerOutputOptions

	// DeleteServer Input Options
	DeleteServerInputOptionsForRuntime0254   *framework.DeleteServerInputOptions
	DeleteServerInputOptionsForRuntime0280   *framework.DeleteServerInputOptions
	DeleteServerInputOptionsForRuntimeLatest *framework.DeleteServerInputOptions
	DeleteServerInputOptionsForRuntime0116   *framework.DeleteServerInputOptions

	// DeleteServer Output Options with expected error
	DeleteServerOutputOptionsForRuntime0280WithError   *framework.DeleteServerOutputOptions
	DeleteServerOutputOptionsForRuntimeLatestWithError *framework.DeleteServerOutputOptions

	// RemoveCurrentServer Input Options
	RemoveCurrentServerInputOptionsForRuntime0280   *framework.RemoveCurrentServerInputOptions
	RemoveCurrentServerInputOptionsForRuntimeLatest *framework.RemoveCurrentServerInputOptions

	// RemoveCurrentServer Output Options with expected error
	RemoveCurrentServerOutputOptionsForRuntimeLatestWithError *framework.RemoveCurrentServerOutputOptions
	RemoveCurrentServerOutputOptionsForRuntime0280WithError   *framework.RemoveCurrentServerOutputOptions

	// Server API Commands
	// SetServer API Commands
	SetServerCmdForRuntimeLatest *core.Command
	SetServerCmdForRuntime0280   *core.Command
	SetServerCmdForRuntime0254   *core.Command
	SetServerCmdForRuntime0116   *core.Command

	SetServerTwoCmdForRuntimeLatest *core.Command
	SetServerTwoCmdForRuntime0280   *core.Command
	SetServerTwoCmdForRuntime0254   *core.Command
	SetServerTwoCmdForRuntime0116   *core.Command

	// SetCurrentServer API Commands
	SetCurrentServerCmdForRuntime0116   *core.Command
	SetCurrentServerCmdForRuntime0254   *core.Command
	SetCurrentServerCmdForRuntime0280   *core.Command
	SetCurrentServerCmdForRuntimeLatest *core.Command

	// GetServer API Commands
	GetServerCmdForRuntimeLatest *core.Command
	GetServerCmdForRuntime0280   *core.Command
	GetServerCmdForRuntime0254   *core.Command
	GetServerCmdForRuntime0116   *core.Command

	GetServerTwoCmdForRuntimeLatest *core.Command
	GetServerTwoCmdForRuntime0280   *core.Command
	GetServerTwoCmdForRuntime0254   *core.Command
	GetServerTwoCmdForRuntime0116   *core.Command

	GetServerCmdForRuntimeLatestWithError *core.Command
	GetServerCmdForRuntime0280WithError   *core.Command
	GetServerCmdForRuntime0254WithError   *core.Command
	GetServerCmdForRuntime0116WithError   *core.Command

	GetServerTwoCmdForRuntimeLatestWithError *core.Command
	GetServerTwoCmdForRuntime0280WithError   *core.Command

	// GetCurrentServer API Commands
	GetCurrentServerCmdForRuntimeLatest *core.Command
	GetCurrentServerCmdForRuntime0280   *core.Command
	GetCurrentServerCmdForRuntime0254   *core.Command
	GetCurrentServerCmdForRuntime0116   *core.Command

	GetCurrentServerCmdForRuntimeLatestWithError *core.Command
	GetCurrentServerCmdForRuntime0280WithError   *core.Command
	GetCurrentServerCmdForRuntime0254WithError   *core.Command
	GetCurrentServerCmdForRuntime0116WithError   *core.Command

	// DeleteServer API Commands
	DeleteServerCmdForRuntime0116   *core.Command
	DeleteServerCmdForRuntime0280   *core.Command
	DeleteServerCmdForRuntime0254   *core.Command
	DeleteServerCmdForRuntimeLatest *core.Command

	DeleteServerCmdForRuntime0280WithError   *core.Command
	DeleteServerCmdForRuntimeLatestWithError *core.Command

	// RemoveCurrentServer API Commands
	RemoveCurrentServerCmdForRuntime0280   *core.Command
	RemoveCurrentServerCmdForRuntimeLatest *core.Command

	RemoveCurrentServerCmdForRuntimeLatestWithError *core.Command
	RemoveCurrentServerCmdForRuntime0280WithError   *core.Command
}

// SetUpDefaultData sets up the Helper struct with default input/output options and api commands
func (b *Helper) SetUpDefaultData() {
	ginkgo.By("Setup Input and Output Options for Servers APIs")
	b.SetupSetServerTestInputAndOutputOptions()
	b.CreateSetServerAPICommands()

	b.SetupSetCurrentServerTestInputAndOutputOptions()
	b.CreateSetCurrentServerAPICommands()

	b.SetupGetCurrentServerTestInputAndOutputOptions()
	b.CreateGetCurrentServerAPICommands()

	b.SetupGetServerTestInputAndOutputOptions()
	b.CreateGetServerAPICommands()

	b.SetupDeleteServerTestInputAndOutputOptions()
	b.CreateDeleteServerAPICommands()

	b.SetupRemoveCurrentServerTestInputAndOutputOptions()
	b.CreateRemoveCurrentServerAPICommands()
}

// SetupRemoveCurrentServerTestInputAndOutputOptions sets input and output options and api commands for RemoveCurrentServer API
func (b *Helper) SetupRemoveCurrentServerTestInputAndOutputOptions() {
	// Input and Output Options for RemoveCurrentServer
	ginkgo.By("Setup Input and Output Options for RemoveCurrentServer")

	b.RemoveCurrentServerInputOptionsForRuntime0280 = DefaultRemoveCurrentServerInputOptions(core.Version0280)
	b.RemoveCurrentServerInputOptionsForRuntimeLatest = DefaultRemoveCurrentServerInputOptions(core.VersionLatest)

	b.RemoveCurrentServerOutputOptionsForRuntimeLatestWithError = DefaultRemoveCurrentServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
	b.RemoveCurrentServerOutputOptionsForRuntime0280WithError = DefaultRemoveCurrentServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
}

// CreateRemoveCurrentServerAPICommands sets api commands for RemoveCurrentServer API
func (b *Helper) CreateRemoveCurrentServerAPICommands() {
	// Create RemoveCurrentServer Commands with input and output options
	ginkgo.By("Create RemoveCurrentServer API Commands")

	removeCurrentServerCmdForRuntime0280, err := framework.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntime0280 = removeCurrentServerCmdForRuntime0280

	removeCurrentServerCmdForRuntimeLatest, err := framework.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntimeLatest = removeCurrentServerCmdForRuntimeLatest

	removeCurrentServerCmdForRuntimeLatestWithError, err := framework.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntimeLatest, b.RemoveCurrentServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntimeLatestWithError = removeCurrentServerCmdForRuntimeLatestWithError

	removeCurrentServerCmdForRuntime0280WithError, err := framework.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntime0280, b.RemoveCurrentServerOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntime0280WithError = removeCurrentServerCmdForRuntime0280WithError
}

// SetupDeleteServerTestInputAndOutputOptions sets api commands for DeleteServer API
func (b *Helper) SetupDeleteServerTestInputAndOutputOptions() {
	// Input and Output Options for DeleteServer
	ginkgo.By("Setup Input and Output Options for DeleteServer")

	b.DeleteServerInputOptionsForRuntime0116 = DefaultDeleteServerInputOptions(core.Version0116, common.CompatibilityTestOne)
	b.DeleteServerInputOptionsForRuntime0280 = DefaultDeleteServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.DeleteServerInputOptionsForRuntime0254 = DefaultDeleteServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.DeleteServerInputOptionsForRuntimeLatest = DefaultDeleteServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)

	b.DeleteServerOutputOptionsForRuntime0280WithError = DefaultDeleteServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
	b.DeleteServerOutputOptionsForRuntimeLatestWithError = DefaultDeleteServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
}

// CreateDeleteServerAPICommands sets api commands for DeleteServer API
func (b *Helper) CreateDeleteServerAPICommands() {
	// Create DeleteServer Commands with input and output options
	ginkgo.By("Create DeleteServer API Commands")

	deleteServerCmdForRuntimeLatest, err := framework.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntimeLatest = deleteServerCmdForRuntimeLatest

	deleteServerCmdForRuntime0280, err := framework.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0280 = deleteServerCmdForRuntime0280

	deleteServerCmdForRuntime0254, err := framework.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0254 = deleteServerCmdForRuntime0254

	deleteServerCmdForRuntime0116, err := framework.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0116 = deleteServerCmdForRuntime0116

	deleteServerCmdForRuntime0280WithError, err := framework.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0280, b.DeleteServerOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0280WithError = deleteServerCmdForRuntime0280WithError

	deleteServerCmdForRuntimeLatestWithError, err := framework.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntimeLatest, b.DeleteServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntimeLatestWithError = deleteServerCmdForRuntimeLatestWithError
}

// SetupGetServerTestInputAndOutputOptions sets input and output options for GetServer API
func (b *Helper) SetupGetServerTestInputAndOutputOptions() {
	// Input and Output params for GetServer
	ginkgo.By("Setup Input and Output Options for GetServer")

	b.GetServerInputOptionsForRuntimeLatest = DefaultGetServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.GetServerInputOptionsForRuntime0280 = DefaultGetServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetServerInputOptionsForRuntime0254 = DefaultGetServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetServerInputOptionsForRuntime0116 = DefaultGetServerInputOptions(core.Version0116, common.CompatibilityTestOne)

	b.GetServerTwoInputOptionsForRuntimeLatest = DefaultGetServerInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetServerTwoInputOptionsForRuntime0280 = DefaultGetServerInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.GetServerTwoInputOptionsForRuntime0254 = DefaultGetServerInputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.GetServerTwoInputOptionsForRuntime0116 = DefaultGetServerInputOptions(core.Version0116, common.CompatibilityTestTwo)

	b.GetServerOutputOptionsForRuntime0280 = DefaultGetServerOutputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0254 = DefaultGetServerOutputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntimeLatest = DefaultGetServerOutputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0116 = DefaultGetServerOutputOptions(core.Version0116, common.CompatibilityTestOne)

	b.GetServerTwoOutputOptionsForRuntimeLatest = DefaultGetServerOutputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime0280 = DefaultGetServerOutputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime0254 = DefaultGetServerOutputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime0116 = DefaultGetServerOutputOptions(core.Version0116, common.CompatibilityTestTwo)

	b.GetServerOutputOptionsForRuntimeLatestWithError = DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0280WithError = DefaultGetServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0254WithError = DefaultGetServerOutputOptionsWithError(core.Version0254, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0116WithError = DefaultGetServerOutputOptionsWithError(core.Version0116, common.CompatibilityTestOne)

	b.GetServerTwoOutputOptionsForRuntimeLatestWithError = DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime0280WithError = DefaultGetServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestTwo)
}

// CreateGetServerAPICommands sets api commands for GetServer API
func (b *Helper) CreateGetServerAPICommands() {
	// Create GetServer Commands with input and output options
	ginkgo.By("Create GetServer API Commands")

	getServerCmdForRuntimeLatest, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntimeLatest, b.GetServerOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntimeLatest = getServerCmdForRuntimeLatest

	getServerCmdForRuntime0280, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntime0280, b.GetServerOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0280 = getServerCmdForRuntime0280

	getServerCmdForRuntime0254, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntime0254, b.GetServerOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0254 = getServerCmdForRuntime0254

	getServerCmdForRuntime0116, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntime0116, b.GetServerOutputOptionsForRuntime0116)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0116 = getServerCmdForRuntime0116

	getServerTwoCmdForRuntimeLatest, err := framework.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntimeLatest, b.GetServerTwoOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntimeLatest = getServerTwoCmdForRuntimeLatest

	getServerTwoCmdForRuntime0280, err := framework.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0280, b.GetServerTwoOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0280 = getServerTwoCmdForRuntime0280

	getServerTwoCmdForRuntime0254, err := framework.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0254, b.GetServerTwoOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0254 = getServerTwoCmdForRuntime0254

	getServerTwoCmdForRuntime0116, err := framework.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0116, b.GetServerTwoOutputOptionsForRuntime0116)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0116 = getServerTwoCmdForRuntime0116

	getServerCmdForRuntimeLatestWithError, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntimeLatest, b.GetServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntimeLatestWithError = getServerCmdForRuntimeLatestWithError

	getServerCmdForRuntime0280WithError, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntime0280, b.GetServerOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0280WithError = getServerCmdForRuntime0280WithError

	getServerCmdForRuntime0254WithError, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntime0254, b.GetServerOutputOptionsForRuntime0254WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0254WithError = getServerCmdForRuntime0254WithError

	getServerCmdForRuntime0116WithError, err := framework.NewGetServerCommand(b.GetServerInputOptionsForRuntime0116, b.GetServerOutputOptionsForRuntime0116WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0116WithError = getServerCmdForRuntime0116WithError

	getServerTwoCmdForRuntimeLatestWithError, err := framework.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntimeLatest, b.GetServerTwoOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntimeLatestWithError = getServerTwoCmdForRuntimeLatestWithError

	getServerTwoCmdForRuntime0280WithError, err := framework.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0280, b.GetServerTwoOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0280WithError = getServerTwoCmdForRuntime0280WithError
}

// SetupGetCurrentServerTestInputAndOutputOptions sets input and output options for GetCurrentServer API
func (b *Helper) SetupGetCurrentServerTestInputAndOutputOptions() {
	// Input and Output Parameters for GetCurrentServer
	ginkgo.By("Setup Input and Output Options for GetCurrentServer")

	b.GetCurrentServerInputOptionsForRuntimeLatest = DefaultGetCurrentServerInputOptions(core.VersionLatest)
	b.GetCurrentServerInputOptionsForRuntime0280 = DefaultGetCurrentServerInputOptions(core.Version0280)
	b.GetCurrentServerInputOptionsForRuntime0254 = DefaultGetCurrentServerInputOptions(core.Version0254)
	b.GetCurrentServerInputOptionsForRuntime0116 = DefaultGetCurrentServerInputOptions(core.Version0116)

	b.GetCurrentServerOutputOptionsForRuntime0280 = DefaultGetCurrentServerOutputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetCurrentServerOutputOptionsForRuntime0254 = DefaultGetCurrentServerOutputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetCurrentServerOutputOptionsForRuntimeLatest = DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.GetCurrentServerOutputOptionsForRuntime0116 = DefaultGetCurrentServerOutputOptions(core.Version0116, common.CompatibilityTestOne)

	b.GetCurrentServerOutputOptionsForRuntimeLatestWithError = DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)
	b.GetCurrentServerOutputOptionsForRuntime0280WithError = DefaultGetCurrentServerOutputOptionsWithError(core.Version0280)
	b.GetCurrentServerOutputOptionsForRuntime0254WithError = DefaultGetCurrentServerOutputOptionsWithError(core.Version0254)
	b.GetCurrentServerOutputOptionsForRuntime0116WithError = DefaultGetCurrentServerOutputOptionsWithError(core.Version0116)
}

// CreateGetCurrentServerAPICommands sets api commands for GetCurrentServer API
func (b *Helper) CreateGetCurrentServerAPICommands() {
	// Create GetCurrentServer Commands with input and output options
	ginkgo.By("Create GetCurrentServer API Commands")

	getCurrentServerCmdForRuntimeLatest, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntimeLatest, b.GetCurrentServerOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntimeLatest = getCurrentServerCmdForRuntimeLatest

	getCurrentServerCmdForRuntime0280, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0280, b.GetCurrentServerOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0280 = getCurrentServerCmdForRuntime0280

	getCurrentServerCmdForRuntime0254, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0254, b.GetCurrentServerOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0254 = getCurrentServerCmdForRuntime0254

	getCurrentServerCmdForRuntime0116, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0116, b.GetCurrentServerOutputOptionsForRuntime0116)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0116 = getCurrentServerCmdForRuntime0116

	getCurrentServerCmdForRuntimeLatestWithError, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntimeLatest, b.GetCurrentServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntimeLatestWithError = getCurrentServerCmdForRuntimeLatestWithError

	getCurrentServerCmdForRuntime0280WithError, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0280, b.GetCurrentServerOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0280WithError = getCurrentServerCmdForRuntime0280WithError

	getCurrentServerCmdForRuntime0254WithError, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0254, b.GetCurrentServerOutputOptionsForRuntime0254WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0254WithError = getCurrentServerCmdForRuntime0254WithError

	getCurrentServerCmdForRuntime0116WithError, err := framework.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0116, b.GetCurrentServerOutputOptionsForRuntime0116WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0116WithError = getCurrentServerCmdForRuntime0116WithError
}

// SetupSetCurrentServerTestInputAndOutputOptions sets input and output options for SetCurrentServer API
func (b *Helper) SetupSetCurrentServerTestInputAndOutputOptions() {
	// Input and Output Parameters for SetCurrentServer
	ginkgo.By("Setup Input and Output Options for SetCurrentServer")

	b.SetCurrentServerInputOptionsForRuntimeLatest = DefaultSetCurrentServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.SetCurrentServerInputOptionsForRuntime0280 = DefaultSetCurrentServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetCurrentServerInputOptionsForRuntime0254 = DefaultSetCurrentServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.SetCurrentServerInputOptionsForRuntime0116 = DefaultSetCurrentServerInputOptions(core.Version0116, common.CompatibilityTestOne)
}

// CreateSetCurrentServerAPICommands sets api commands for SetCurrentServer API
func (b *Helper) CreateSetCurrentServerAPICommands() {
	// Create SetCurrentServer Commands with input and output options
	ginkgo.By("Create SetCurrentServer API Commands")

	setCurrentServerCmdForRuntimeLatest, err := framework.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntimeLatest = setCurrentServerCmdForRuntimeLatest

	setCurrentServerCmdForRuntime0280, err := framework.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntime0280 = setCurrentServerCmdForRuntime0280

	setCurrentServerCmdForRuntime0254, err := framework.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntime0254 = setCurrentServerCmdForRuntime0254

	setCurrentServerCmdForRuntime0116, err := framework.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntime0116 = setCurrentServerCmdForRuntime0116
}

// SetupSetServerTestInputAndOutputOptions sets input and output options for SetServer API
func (b *Helper) SetupSetServerTestInputAndOutputOptions() {
	// Input and Output Parameters for SetServer
	ginkgo.By("Setup Input and Output Options for SetServer")

	b.SetServerInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestOne)

	b.SetServerTwoInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestTwo)

	// Input and Output Parameters for SetServer
	b.SetServerInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestOne)

	b.SetServerTwoInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestTwo)

	b.CreateSetServerAPICommands()
}

// CreateSetServerAPICommands sets api commands for SetServer API
func (b *Helper) CreateSetServerAPICommands() {
	// Create SetServer Commands with input and output options
	ginkgo.By("Create SetServer API Commands")

	setServerCmdForRuntimeLatest, err := framework.NewSetServerCommand(b.SetServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntimeLatest = setServerCmdForRuntimeLatest

	setServerCmdForRuntime0254, err := framework.NewSetServerCommand(b.SetServerInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntime0254 = setServerCmdForRuntime0254

	setServerCmdForRuntime0280, err := framework.NewSetServerCommand(b.SetServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntime0280 = setServerCmdForRuntime0280

	setServerCmdForRuntime0116, err := framework.NewSetServerCommand(b.SetServerInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntime0116 = setServerCmdForRuntime0116

	setServerTwoCmdForRuntimeLatest, err := framework.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntimeLatest = setServerTwoCmdForRuntimeLatest

	setServerTwoCmdForRuntime0254, err := framework.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntime0254 = setServerTwoCmdForRuntime0254

	setServerTwoCmdForRuntime0280, err := framework.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntime0280 = setServerTwoCmdForRuntime0280

	setServerTwoCmdForRuntime0116, err := framework.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntime0116 = setServerTwoCmdForRuntime0116
}

// DefaultSetServerInputOptions helper method to construct SetServer API input options
func DefaultSetServerInputOptions(version core.RuntimeVersion, serverName string) *framework.SetServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.SetServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}

	return nil
}

// DefaultGetServerInputOptions helper method to construct GetServer API input options
func DefaultGetServerInputOptions(version core.RuntimeVersion, serverName string) *framework.GetServerInputOptions {
	return &framework.GetServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultGetServerOutputOptions helper method to construct GetServer API output options
func DefaultGetServerOutputOptions(version core.RuntimeVersion, serverName string) *framework.GetServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &framework.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0254, core.Version0116:
		return &framework.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}
	return nil
}

// DefaultGetServerOutputOptionsWithError helper method to construct GetServer API output options with error
func DefaultGetServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *framework.GetServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.VersionLatest,
			},
			Error: fmt.Sprintf("could not find server \"%v\"", serverName),
		}
	}
	return nil
}

// DefaultSetCurrentServerInputOptions helper method to construct SetCurrentServer API input options
func DefaultSetCurrentServerInputOptions(version core.RuntimeVersion, serverName string) *framework.SetCurrentServerInputOptions {
	return &framework.SetCurrentServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultGetCurrentServerInputOptions helper method to construct GetCurrentServer API input options
func DefaultGetCurrentServerInputOptions(version core.RuntimeVersion) *framework.GetCurrentServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.GetCurrentServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.VersionLatest,
			},
		}
	}
	return nil
}

// DefaultGetCurrentServerOutputOptions helper method to construct GetCurrentServer API output options
func DefaultGetCurrentServerOutputOptions(version core.RuntimeVersion, serverName string) *framework.GetCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0254, core.Version0116:
		return &framework.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0280:
		return &framework.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0280,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	}
	return nil
}

// DefaultGetCurrentServerOutputOptionsWithError helper method to construct GetCurrentServer API output options with error
func DefaultGetCurrentServerOutputOptionsWithError(version core.RuntimeVersion) *framework.GetCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: ServerNotFound,
		}
	}
	return nil
}

// DefaultRemoveCurrentServerInputOptions helper method to construct RemoveCurrentServer API input options
func DefaultRemoveCurrentServerInputOptions(version core.RuntimeVersion) *framework.RemoveCurrentServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &framework.RemoveCurrentServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerName: common.CompatibilityTestOne,
		}
	}
	return nil
}

// DefaultRemoveCurrentServerOutputOptionsWithError helper method to construct RemoveCurrentServer API output option
func DefaultRemoveCurrentServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *framework.RemoveCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.RemoveCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", serverName),
		}
	}
	return nil
}

// DefaultDeleteServerInputOptions helper method to construct DeleteServer API input options
func DefaultDeleteServerInputOptions(version core.RuntimeVersion, serverName string) *framework.DeleteServerInputOptions {
	return &framework.DeleteServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultDeleteServerOutputOptionsWithError helper method to construct DeleteServer API output options
func DefaultDeleteServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *framework.DeleteServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.DeleteServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", serverName),
		}
	}
	return nil
}
