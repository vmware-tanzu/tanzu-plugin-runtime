// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package server provides cross-version Server API compatibility tests
package server

import (
	"fmt"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/server"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

const ServerNotFound = "current server \"\" not found in tanzu config"

// Helper struct provides input and output options and api commands to be used in test cases
type Helper struct {
	// SetServer Input Options
	SetServerInputOptionsForRuntime0116   *server.SetServerInputOptions
	SetServerInputOptionsForRuntime0254   *server.SetServerInputOptions
	SetServerInputOptionsForRuntime0280   *server.SetServerInputOptions
	SetServerInputOptionsForRuntimeLatest *server.SetServerInputOptions
	SetServerInputOptionsForRuntime090    *server.SetServerInputOptions

	SetServerTwoInputOptionsForRuntime0116   *server.SetServerInputOptions
	SetServerTwoInputOptionsForRuntime0254   *server.SetServerInputOptions
	SetServerTwoInputOptionsForRuntime0280   *server.SetServerInputOptions
	SetServerTwoInputOptionsForRuntimeLatest *server.SetServerInputOptions
	SetServerTwoInputOptionsForRuntime090    *server.SetServerInputOptions

	// SetCurrentServer Input Options
	SetCurrentServerInputOptionsForRuntime0116   *server.SetCurrentServerInputOptions
	SetCurrentServerInputOptionsForRuntime0254   *server.SetCurrentServerInputOptions
	SetCurrentServerInputOptionsForRuntime0280   *server.SetCurrentServerInputOptions
	SetCurrentServerInputOptionsForRuntimeLatest *server.SetCurrentServerInputOptions
	SetCurrentServerInputOptionsForRuntime090    *server.SetCurrentServerInputOptions

	// GetServer Input Options
	GetServerInputOptionsForRuntimeLatest *server.GetServerInputOptions
	GetServerInputOptionsForRuntime090    *server.GetServerInputOptions
	GetServerInputOptionsForRuntime0280   *server.GetServerInputOptions
	GetServerInputOptionsForRuntime0254   *server.GetServerInputOptions
	GetServerInputOptionsForRuntime0116   *server.GetServerInputOptions

	GetServerTwoInputOptionsForRuntimeLatest *server.GetServerInputOptions
	GetServerTwoInputOptionsForRuntime090    *server.GetServerInputOptions
	GetServerTwoInputOptionsForRuntime0280   *server.GetServerInputOptions
	GetServerTwoInputOptionsForRuntime0254   *server.GetServerInputOptions
	GetServerTwoInputOptionsForRuntime0116   *server.GetServerInputOptions

	// GetServer Output Options
	GetServerOutputOptionsForRuntime0116   *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0254   *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0280   *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntimeLatest *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntime090    *server.GetServerOutputOptions

	GetServerTwoOutputOptionsForRuntime0116   *server.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime0254   *server.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime0280   *server.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntimeLatest *server.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime090    *server.GetServerOutputOptions

	// GetServer Output Options with expected error
	GetServerOutputOptionsForRuntimeLatestWithError *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntime090WithError    *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0280WithError   *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0254WithError   *server.GetServerOutputOptions
	GetServerOutputOptionsForRuntime0116WithError   *server.GetServerOutputOptions

	GetServerTwoOutputOptionsForRuntimeLatestWithError *server.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime090WithError    *server.GetServerOutputOptions
	GetServerTwoOutputOptionsForRuntime0280WithError   *server.GetServerOutputOptions

	// GetCurrentServer Input Options
	GetCurrentServerInputOptionsForRuntime0116   *server.GetCurrentServerInputOptions
	GetCurrentServerInputOptionsForRuntime0254   *server.GetCurrentServerInputOptions
	GetCurrentServerInputOptionsForRuntime0280   *server.GetCurrentServerInputOptions
	GetCurrentServerInputOptionsForRuntimeLatest *server.GetCurrentServerInputOptions
	GetCurrentServerInputOptionsForRuntime090    *server.GetCurrentServerInputOptions

	// GetCurrentServer Output Options
	GetCurrentServerOutputOptionsForRuntime0116   *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0254   *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0280   *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime090    *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntimeLatest *server.GetCurrentServerOutputOptions

	// GetCurrentServer Output Options with expected error
	GetCurrentServerOutputOptionsForRuntimeLatestWithError *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime090WithError    *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0280WithError   *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0254WithError   *server.GetCurrentServerOutputOptions
	GetCurrentServerOutputOptionsForRuntime0116WithError   *server.GetCurrentServerOutputOptions

	// DeleteServer Input Options
	DeleteServerInputOptionsForRuntime0254   *server.DeleteServerInputOptions
	DeleteServerInputOptionsForRuntime0280   *server.DeleteServerInputOptions
	DeleteServerInputOptionsForRuntime090    *server.DeleteServerInputOptions
	DeleteServerInputOptionsForRuntimeLatest *server.DeleteServerInputOptions
	DeleteServerInputOptionsForRuntime0116   *server.DeleteServerInputOptions

	// DeleteServer Output Options with expected error
	DeleteServerOutputOptionsForRuntime0280WithError   *server.DeleteServerOutputOptions
	DeleteServerOutputOptionsForRuntimeLatestWithError *server.DeleteServerOutputOptions
	DeleteServerOutputOptionsForRuntime090WithError    *server.DeleteServerOutputOptions

	// RemoveCurrentServer Input Options
	RemoveCurrentServerInputOptionsForRuntime0280   *server.RemoveCurrentServerInputOptions
	RemoveCurrentServerInputOptionsForRuntimeLatest *server.RemoveCurrentServerInputOptions
	RemoveCurrentServerInputOptionsForRuntime090    *server.RemoveCurrentServerInputOptions

	// RemoveCurrentServer Output Options with expected error
	RemoveCurrentServerOutputOptionsForRuntimeLatestWithError *server.RemoveCurrentServerOutputOptions
	RemoveCurrentServerOutputOptionsForRuntime090WithError    *server.RemoveCurrentServerOutputOptions
	RemoveCurrentServerOutputOptionsForRuntime0280WithError   *server.RemoveCurrentServerOutputOptions

	// Server API Commands
	// SetServer API Commands
	SetServerCmdForRuntimeLatest *core.Command
	SetServerCmdForRuntime090    *core.Command
	SetServerCmdForRuntime0280   *core.Command
	SetServerCmdForRuntime0254   *core.Command
	SetServerCmdForRuntime0116   *core.Command

	SetServerTwoCmdForRuntimeLatest *core.Command
	SetServerTwoCmdForRuntime090    *core.Command
	SetServerTwoCmdForRuntime0280   *core.Command
	SetServerTwoCmdForRuntime0254   *core.Command
	SetServerTwoCmdForRuntime0116   *core.Command

	// SetCurrentServer API Commands
	SetCurrentServerCmdForRuntime0116   *core.Command
	SetCurrentServerCmdForRuntime0254   *core.Command
	SetCurrentServerCmdForRuntime0280   *core.Command
	SetCurrentServerCmdForRuntime090    *core.Command
	SetCurrentServerCmdForRuntimeLatest *core.Command

	// GetServer API Commands
	GetServerCmdForRuntimeLatest *core.Command
	GetServerCmdForRuntime090    *core.Command
	GetServerCmdForRuntime0280   *core.Command
	GetServerCmdForRuntime0254   *core.Command
	GetServerCmdForRuntime0116   *core.Command

	GetServerTwoCmdForRuntimeLatest *core.Command
	GetServerTwoCmdForRuntime090    *core.Command
	GetServerTwoCmdForRuntime0280   *core.Command
	GetServerTwoCmdForRuntime0254   *core.Command
	GetServerTwoCmdForRuntime0116   *core.Command

	GetServerCmdForRuntimeLatestWithError *core.Command
	GetServerCmdForRuntime090WithError    *core.Command
	GetServerCmdForRuntime0280WithError   *core.Command
	GetServerCmdForRuntime0254WithError   *core.Command
	GetServerCmdForRuntime0116WithError   *core.Command

	GetServerTwoCmdForRuntimeLatestWithError *core.Command
	GetServerTwoCmdForRuntime090WithError    *core.Command
	GetServerTwoCmdForRuntime0280WithError   *core.Command

	// GetCurrentServer API Commands
	GetCurrentServerCmdForRuntimeLatest *core.Command
	GetCurrentServerCmdForRuntime090    *core.Command
	GetCurrentServerCmdForRuntime0280   *core.Command
	GetCurrentServerCmdForRuntime0254   *core.Command
	GetCurrentServerCmdForRuntime0116   *core.Command

	GetCurrentServerCmdForRuntimeLatestWithError *core.Command
	GetCurrentServerCmdForRuntime090WithError    *core.Command
	GetCurrentServerCmdForRuntime0280WithError   *core.Command
	GetCurrentServerCmdForRuntime0254WithError   *core.Command
	GetCurrentServerCmdForRuntime0116WithError   *core.Command

	// DeleteServer API Commands
	DeleteServerCmdForRuntime0116   *core.Command
	DeleteServerCmdForRuntime0280   *core.Command
	DeleteServerCmdForRuntime0254   *core.Command
	DeleteServerCmdForRuntime090    *core.Command
	DeleteServerCmdForRuntimeLatest *core.Command

	DeleteServerCmdForRuntime0280WithError   *core.Command
	DeleteServerCmdForRuntime090WithError    *core.Command
	DeleteServerCmdForRuntimeLatestWithError *core.Command

	// RemoveCurrentServer API Commands
	RemoveCurrentServerCmdForRuntime0280   *core.Command
	RemoveCurrentServerCmdForRuntime090    *core.Command
	RemoveCurrentServerCmdForRuntimeLatest *core.Command

	RemoveCurrentServerCmdForRuntimeLatestWithError *core.Command
	RemoveCurrentServerCmdForRuntime090WithError    *core.Command
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
	b.RemoveCurrentServerInputOptionsForRuntime090 = DefaultRemoveCurrentServerInputOptions(core.Version090)

	b.RemoveCurrentServerOutputOptionsForRuntime090WithError = DefaultRemoveCurrentServerOutputOptionsWithError(core.Version090, common.CompatibilityTestOne)
	b.RemoveCurrentServerOutputOptionsForRuntimeLatestWithError = DefaultRemoveCurrentServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
	b.RemoveCurrentServerOutputOptionsForRuntime0280WithError = DefaultRemoveCurrentServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
}

// CreateRemoveCurrentServerAPICommands sets api commands for RemoveCurrentServer API
func (b *Helper) CreateRemoveCurrentServerAPICommands() {
	// Create RemoveCurrentServer Commands with input and output options
	ginkgo.By("Create RemoveCurrentServer API Commands")

	removeCurrentServerCmdForRuntime0280, err := server.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntime0280 = removeCurrentServerCmdForRuntime0280

	removeCurrentServerCmdForRuntimeLatest, err := server.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntimeLatest = removeCurrentServerCmdForRuntimeLatest

	removeCurrentServerCmdForRuntime090, err := server.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntime090, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntime090 = removeCurrentServerCmdForRuntime090

	removeCurrentServerCmdForRuntimeLatestWithError, err := server.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntimeLatest, b.RemoveCurrentServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntimeLatestWithError = removeCurrentServerCmdForRuntimeLatestWithError

	removeCurrentServerCmdForRuntime090WithError, err := server.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntime090, b.RemoveCurrentServerOutputOptionsForRuntime090WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.RemoveCurrentServerCmdForRuntime090WithError = removeCurrentServerCmdForRuntime090WithError

	removeCurrentServerCmdForRuntime0280WithError, err := server.NewRemoveCurrentServerCommand(b.RemoveCurrentServerInputOptionsForRuntime0280, b.RemoveCurrentServerOutputOptionsForRuntime0280WithError)
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
	b.DeleteServerInputOptionsForRuntime090 = DefaultDeleteServerInputOptions(core.Version090, common.CompatibilityTestOne)
	b.DeleteServerInputOptionsForRuntimeLatest = DefaultDeleteServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)

	b.DeleteServerOutputOptionsForRuntime0280WithError = DefaultDeleteServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
	b.DeleteServerOutputOptionsForRuntimeLatestWithError = DefaultDeleteServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
	b.DeleteServerOutputOptionsForRuntime090WithError = DefaultDeleteServerOutputOptionsWithError(core.Version090, common.CompatibilityTestOne)
}

// CreateDeleteServerAPICommands sets api commands for DeleteServer API
func (b *Helper) CreateDeleteServerAPICommands() {
	// Create DeleteServer Commands with input and output options
	ginkgo.By("Create DeleteServer API Commands")

	deleteServerCmdForRuntimeLatest, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntimeLatest = deleteServerCmdForRuntimeLatest

	deleteServerCmdForRuntime090, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime090, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime090 = deleteServerCmdForRuntime090

	deleteServerCmdForRuntime0280, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0280 = deleteServerCmdForRuntime0280

	deleteServerCmdForRuntime0254, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0254 = deleteServerCmdForRuntime0254

	deleteServerCmdForRuntime0116, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0116 = deleteServerCmdForRuntime0116

	deleteServerCmdForRuntime0280WithError, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime0280, b.DeleteServerOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime0280WithError = deleteServerCmdForRuntime0280WithError

	deleteServerCmdForRuntimeLatestWithError, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntimeLatest, b.DeleteServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntimeLatestWithError = deleteServerCmdForRuntimeLatestWithError

	deleteServerCmdForRuntime090WithError, err := server.NewDeleteServerCommand(b.DeleteServerInputOptionsForRuntime090, b.DeleteServerOutputOptionsForRuntime090WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.DeleteServerCmdForRuntime090WithError = deleteServerCmdForRuntime090WithError
}

// SetupGetServerTestInputAndOutputOptions sets input and output options for GetServer API
func (b *Helper) SetupGetServerTestInputAndOutputOptions() {
	// Input and Output params for GetServer
	ginkgo.By("Setup Input and Output Options for GetServer")

	b.GetServerInputOptionsForRuntimeLatest = DefaultGetServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.GetServerInputOptionsForRuntime090 = DefaultGetServerInputOptions(core.Version090, common.CompatibilityTestOne)

	b.GetServerInputOptionsForRuntime0280 = DefaultGetServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetServerInputOptionsForRuntime0254 = DefaultGetServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetServerInputOptionsForRuntime0116 = DefaultGetServerInputOptions(core.Version0116, common.CompatibilityTestOne)

	b.GetServerTwoInputOptionsForRuntimeLatest = DefaultGetServerInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetServerTwoInputOptionsForRuntime090 = DefaultGetServerInputOptions(core.Version090, common.CompatibilityTestTwo)

	b.GetServerTwoInputOptionsForRuntime0280 = DefaultGetServerInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.GetServerTwoInputOptionsForRuntime0254 = DefaultGetServerInputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.GetServerTwoInputOptionsForRuntime0116 = DefaultGetServerInputOptions(core.Version0116, common.CompatibilityTestTwo)

	b.GetServerOutputOptionsForRuntime0280 = DefaultGetServerOutputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0254 = DefaultGetServerOutputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntimeLatest = DefaultGetServerOutputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime090 = DefaultGetServerOutputOptions(core.Version090, common.CompatibilityTestOne)

	b.GetServerOutputOptionsForRuntime0116 = DefaultGetServerOutputOptions(core.Version0116, common.CompatibilityTestOne)

	b.GetServerTwoOutputOptionsForRuntimeLatest = DefaultGetServerOutputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime090 = DefaultGetServerOutputOptions(core.Version090, common.CompatibilityTestTwo)

	b.GetServerTwoOutputOptionsForRuntime0280 = DefaultGetServerOutputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime0254 = DefaultGetServerOutputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime0116 = DefaultGetServerOutputOptions(core.Version0116, common.CompatibilityTestTwo)

	b.GetServerOutputOptionsForRuntimeLatestWithError = DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime090WithError = DefaultGetServerOutputOptionsWithError(core.Version090, common.CompatibilityTestOne)

	b.GetServerOutputOptionsForRuntime0280WithError = DefaultGetServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0254WithError = DefaultGetServerOutputOptionsWithError(core.Version0254, common.CompatibilityTestOne)
	b.GetServerOutputOptionsForRuntime0116WithError = DefaultGetServerOutputOptionsWithError(core.Version0116, common.CompatibilityTestOne)

	b.GetServerTwoOutputOptionsForRuntimeLatestWithError = DefaultGetServerOutputOptionsWithError(core.VersionLatest, common.CompatibilityTestTwo)
	b.GetServerTwoOutputOptionsForRuntime090WithError = DefaultGetServerOutputOptionsWithError(core.Version090, common.CompatibilityTestTwo)

	b.GetServerTwoOutputOptionsForRuntime0280WithError = DefaultGetServerOutputOptionsWithError(core.Version0280, common.CompatibilityTestTwo)
}

// CreateGetServerAPICommands sets api commands for GetServer API
//
//nolint:funlen
func (b *Helper) CreateGetServerAPICommands() {
	// Create GetServer Commands with input and output options
	ginkgo.By("Create GetServer API Commands")

	getServerCmdForRuntimeLatest, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntimeLatest, b.GetServerOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntimeLatest = getServerCmdForRuntimeLatest

	getServerCmdForRuntime090, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime090, b.GetServerOutputOptionsForRuntime090)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime090 = getServerCmdForRuntime090

	getServerCmdForRuntime0280, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime0280, b.GetServerOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0280 = getServerCmdForRuntime0280

	getServerCmdForRuntime0254, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime0254, b.GetServerOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0254 = getServerCmdForRuntime0254

	getServerCmdForRuntime0116, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime0116, b.GetServerOutputOptionsForRuntime0116)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0116 = getServerCmdForRuntime0116

	getServerTwoCmdForRuntimeLatest, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntimeLatest, b.GetServerTwoOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntimeLatest = getServerTwoCmdForRuntimeLatest

	getServerTwoCmdForRuntime090, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime090, b.GetServerTwoOutputOptionsForRuntime090)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime090 = getServerTwoCmdForRuntime090

	getServerTwoCmdForRuntime0280, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0280, b.GetServerTwoOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0280 = getServerTwoCmdForRuntime0280

	getServerTwoCmdForRuntime0254, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0254, b.GetServerTwoOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0254 = getServerTwoCmdForRuntime0254

	getServerTwoCmdForRuntime0116, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0116, b.GetServerTwoOutputOptionsForRuntime0116)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0116 = getServerTwoCmdForRuntime0116

	getServerCmdForRuntimeLatestWithError, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntimeLatest, b.GetServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntimeLatestWithError = getServerCmdForRuntimeLatestWithError

	getServerCmdForRuntime090WithError, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime090, b.GetServerOutputOptionsForRuntime090WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime090WithError = getServerCmdForRuntime090WithError

	getServerCmdForRuntime0280WithError, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime0280, b.GetServerOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0280WithError = getServerCmdForRuntime0280WithError

	getServerCmdForRuntime0254WithError, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime0254, b.GetServerOutputOptionsForRuntime0254WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0254WithError = getServerCmdForRuntime0254WithError

	getServerCmdForRuntime0116WithError, err := server.NewGetServerCommand(b.GetServerInputOptionsForRuntime0116, b.GetServerOutputOptionsForRuntime0116WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerCmdForRuntime0116WithError = getServerCmdForRuntime0116WithError

	getServerTwoCmdForRuntimeLatestWithError, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntimeLatest, b.GetServerTwoOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntimeLatestWithError = getServerTwoCmdForRuntimeLatestWithError

	getServerTwoCmdForRuntime090WithError, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime090, b.GetServerTwoOutputOptionsForRuntime090WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime090WithError = getServerTwoCmdForRuntime090WithError

	getServerTwoCmdForRuntime0280WithError, err := server.NewGetServerCommand(b.GetServerTwoInputOptionsForRuntime0280, b.GetServerTwoOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetServerTwoCmdForRuntime0280WithError = getServerTwoCmdForRuntime0280WithError
}

// SetupGetCurrentServerTestInputAndOutputOptions sets input and output options for GetCurrentServer API
func (b *Helper) SetupGetCurrentServerTestInputAndOutputOptions() {
	// Input and Output Parameters for GetCurrentServer
	ginkgo.By("Setup Input and Output Options for GetCurrentServer")

	b.GetCurrentServerInputOptionsForRuntimeLatest = DefaultGetCurrentServerInputOptions(core.VersionLatest)

	b.GetCurrentServerInputOptionsForRuntime090 = DefaultGetCurrentServerInputOptions(core.Version090)

	b.GetCurrentServerInputOptionsForRuntime0280 = DefaultGetCurrentServerInputOptions(core.Version0280)
	b.GetCurrentServerInputOptionsForRuntime0254 = DefaultGetCurrentServerInputOptions(core.Version0254)
	b.GetCurrentServerInputOptionsForRuntime0116 = DefaultGetCurrentServerInputOptions(core.Version0116)

	b.GetCurrentServerOutputOptionsForRuntime0280 = DefaultGetCurrentServerOutputOptions(core.Version0280, common.CompatibilityTestOne)
	b.GetCurrentServerOutputOptionsForRuntime0254 = DefaultGetCurrentServerOutputOptions(core.Version0254, common.CompatibilityTestOne)
	b.GetCurrentServerOutputOptionsForRuntimeLatest = DefaultGetCurrentServerOutputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.GetCurrentServerOutputOptionsForRuntime090 = DefaultGetCurrentServerOutputOptions(core.Version090, common.CompatibilityTestOne)

	b.GetCurrentServerOutputOptionsForRuntime0116 = DefaultGetCurrentServerOutputOptions(core.Version0116, common.CompatibilityTestOne)

	b.GetCurrentServerOutputOptionsForRuntimeLatestWithError = DefaultGetCurrentServerOutputOptionsWithError(core.VersionLatest)
	b.GetCurrentServerOutputOptionsForRuntime090WithError = DefaultGetCurrentServerOutputOptionsWithError(core.Version090)

	b.GetCurrentServerOutputOptionsForRuntime0280WithError = DefaultGetCurrentServerOutputOptionsWithError(core.Version0280)
	b.GetCurrentServerOutputOptionsForRuntime0254WithError = DefaultGetCurrentServerOutputOptionsWithError(core.Version0254)
	b.GetCurrentServerOutputOptionsForRuntime0116WithError = DefaultGetCurrentServerOutputOptionsWithError(core.Version0116)
}

// CreateGetCurrentServerAPICommands sets api commands for GetCurrentServer API
func (b *Helper) CreateGetCurrentServerAPICommands() {
	// Create GetCurrentServer Commands with input and output options
	ginkgo.By("Create GetCurrentServer API Commands")

	getCurrentServerCmdForRuntimeLatest, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntimeLatest, b.GetCurrentServerOutputOptionsForRuntimeLatest)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntimeLatest = getCurrentServerCmdForRuntimeLatest

	getCurrentServerCmdForRuntime090, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime090, b.GetCurrentServerOutputOptionsForRuntime090)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime090 = getCurrentServerCmdForRuntime090

	getCurrentServerCmdForRuntime0280, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0280, b.GetCurrentServerOutputOptionsForRuntime0280)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0280 = getCurrentServerCmdForRuntime0280

	getCurrentServerCmdForRuntime0254, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0254, b.GetCurrentServerOutputOptionsForRuntime0254)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0254 = getCurrentServerCmdForRuntime0254

	getCurrentServerCmdForRuntime0116, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0116, b.GetCurrentServerOutputOptionsForRuntime0116)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0116 = getCurrentServerCmdForRuntime0116

	getCurrentServerCmdForRuntimeLatestWithError, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntimeLatest, b.GetCurrentServerOutputOptionsForRuntimeLatestWithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntimeLatestWithError = getCurrentServerCmdForRuntimeLatestWithError

	getCurrentServerCmdForRuntime090WithError, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime090, b.GetCurrentServerOutputOptionsForRuntime090WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime090WithError = getCurrentServerCmdForRuntime090WithError

	getCurrentServerCmdForRuntime0280WithError, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0280, b.GetCurrentServerOutputOptionsForRuntime0280WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0280WithError = getCurrentServerCmdForRuntime0280WithError

	getCurrentServerCmdForRuntime0254WithError, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0254, b.GetCurrentServerOutputOptionsForRuntime0254WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0254WithError = getCurrentServerCmdForRuntime0254WithError

	getCurrentServerCmdForRuntime0116WithError, err := server.NewGetCurrentServerCommand(b.GetCurrentServerInputOptionsForRuntime0116, b.GetCurrentServerOutputOptionsForRuntime0116WithError)
	gomega.Expect(err).To(gomega.BeNil())
	b.GetCurrentServerCmdForRuntime0116WithError = getCurrentServerCmdForRuntime0116WithError
}

// SetupSetCurrentServerTestInputAndOutputOptions sets input and output options for SetCurrentServer API
func (b *Helper) SetupSetCurrentServerTestInputAndOutputOptions() {
	// Input and Output Parameters for SetCurrentServer
	ginkgo.By("Setup Input and Output Options for SetCurrentServer")

	b.SetCurrentServerInputOptionsForRuntimeLatest = DefaultSetCurrentServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)
	b.SetCurrentServerInputOptionsForRuntime090 = DefaultSetCurrentServerInputOptions(core.Version090, common.CompatibilityTestOne)

	b.SetCurrentServerInputOptionsForRuntime0280 = DefaultSetCurrentServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetCurrentServerInputOptionsForRuntime0254 = DefaultSetCurrentServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.SetCurrentServerInputOptionsForRuntime0116 = DefaultSetCurrentServerInputOptions(core.Version0116, common.CompatibilityTestOne)
}

// CreateSetCurrentServerAPICommands sets api commands for SetCurrentServer API
func (b *Helper) CreateSetCurrentServerAPICommands() {
	// Create SetCurrentServer Commands with input and output options
	ginkgo.By("Create SetCurrentServer API Commands")

	setCurrentServerCmdForRuntimeLatest, err := server.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntimeLatest = setCurrentServerCmdForRuntimeLatest

	setCurrentServerCmdForRuntime090, err := server.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntime090, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntime090 = setCurrentServerCmdForRuntime090

	setCurrentServerCmdForRuntime0280, err := server.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntime0280 = setCurrentServerCmdForRuntime0280

	setCurrentServerCmdForRuntime0254, err := server.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntime0254 = setCurrentServerCmdForRuntime0254

	setCurrentServerCmdForRuntime0116, err := server.NewSetCurrentServerCommand(b.SetCurrentServerInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetCurrentServerCmdForRuntime0116 = setCurrentServerCmdForRuntime0116
}

// SetupSetServerTestInputAndOutputOptions sets input and output options for SetServer API
func (b *Helper) SetupSetServerTestInputAndOutputOptions() {
	// Input and Output Parameters for SetServer
	ginkgo.By("Setup Input and Output Options for SetServer")

	b.SetServerInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)

	b.SetServerInputOptionsForRuntime090 = DefaultSetServerInputOptions(core.Version090, common.CompatibilityTestOne)

	b.SetServerInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestOne)

	b.SetServerTwoInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestTwo)

	b.SetServerTwoInputOptionsForRuntime090 = DefaultSetServerInputOptions(core.Version090, common.CompatibilityTestTwo)

	b.SetServerTwoInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestTwo)

	// Input and Output Parameters for SetServer
	b.SetServerInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestOne)

	b.SetServerInputOptionsForRuntime090 = DefaultSetServerInputOptions(core.Version090, common.CompatibilityTestOne)

	b.SetServerInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestOne)
	b.SetServerInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestOne)

	b.SetServerTwoInputOptionsForRuntimeLatest = DefaultSetServerInputOptions(core.VersionLatest, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime090 = DefaultSetServerInputOptions(core.Version090, common.CompatibilityTestTwo)

	b.SetServerTwoInputOptionsForRuntime0280 = DefaultSetServerInputOptions(core.Version0280, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0254 = DefaultSetServerInputOptions(core.Version0254, common.CompatibilityTestTwo)
	b.SetServerTwoInputOptionsForRuntime0116 = DefaultSetServerInputOptions(core.Version0116, common.CompatibilityTestTwo)

	b.CreateSetServerAPICommands()
}

// CreateSetServerAPICommands sets api commands for SetServer API
func (b *Helper) CreateSetServerAPICommands() {
	// Create SetServer Commands with input and output options
	ginkgo.By("Create SetServer API Commands")

	setServerCmdForRuntimeLatest, err := server.NewSetServerCommand(b.SetServerInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntimeLatest = setServerCmdForRuntimeLatest

	setServerCmdForRuntime090, err := server.NewSetServerCommand(b.SetServerInputOptionsForRuntime090, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntime090 = setServerCmdForRuntime090

	setServerCmdForRuntime0254, err := server.NewSetServerCommand(b.SetServerInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntime0254 = setServerCmdForRuntime0254

	setServerCmdForRuntime0280, err := server.NewSetServerCommand(b.SetServerInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntime0280 = setServerCmdForRuntime0280

	setServerCmdForRuntime0116, err := server.NewSetServerCommand(b.SetServerInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerCmdForRuntime0116 = setServerCmdForRuntime0116

	setServerTwoCmdForRuntimeLatest, err := server.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntimeLatest, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntimeLatest = setServerTwoCmdForRuntimeLatest

	setServerTwoCmdForRuntime090, err := server.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntime090, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntime090 = setServerTwoCmdForRuntime090

	setServerTwoCmdForRuntime0254, err := server.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntime0254, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntime0254 = setServerTwoCmdForRuntime0254

	setServerTwoCmdForRuntime0280, err := server.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntime0280, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntime0280 = setServerTwoCmdForRuntime0280

	setServerTwoCmdForRuntime0116, err := server.NewSetServerCommand(b.SetServerTwoInputOptionsForRuntime0116, nil)
	gomega.Expect(err).To(gomega.BeNil())
	b.SetServerTwoCmdForRuntime0116 = setServerTwoCmdForRuntime0116
}

// DefaultSetServerInputOptions helper method to construct SetServer API input options
func DefaultSetServerInputOptions(version core.RuntimeVersion, serverName string) *server.SetServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		return &server.SetServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &types.ServerOpts{
				Name: serverName,
				Type: types.ManagementClusterServerType,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}

	return nil
}

// DefaultGetServerInputOptions helper method to construct GetServer API input options
func DefaultGetServerInputOptions(version core.RuntimeVersion, serverName string) *server.GetServerInputOptions {
	return &server.GetServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultGetServerOutputOptions helper method to construct GetServer API output options
func DefaultGetServerOutputOptions(version core.RuntimeVersion, serverName string) *server.GetServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280:
		return &server.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &types.ServerOpts{
				Name: serverName,
				Type: types.ManagementClusterServerType,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0254, core.Version0116:
		return &server.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &types.ServerOpts{
				Name: serverName,
				Type: types.ManagementClusterServerType,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}
	return nil
}

// DefaultGetServerOutputOptionsWithError helper method to construct GetServer API output options with error
func DefaultGetServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *server.GetServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		return &server.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.VersionLatest,
			},
			Error: fmt.Sprintf("could not find server \"%v\"", serverName),
		}
	}
	return nil
}

// DefaultSetCurrentServerInputOptions helper method to construct SetCurrentServer API input options
func DefaultSetCurrentServerInputOptions(version core.RuntimeVersion, serverName string) *server.SetCurrentServerInputOptions {
	return &server.SetCurrentServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultGetCurrentServerInputOptions helper method to construct GetCurrentServer API input options
func DefaultGetCurrentServerInputOptions(version core.RuntimeVersion) *server.GetCurrentServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		return &server.GetCurrentServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.VersionLatest,
			},
		}
	}
	return nil
}

// DefaultGetCurrentServerOutputOptions helper method to construct GetCurrentServer API output options
func DefaultGetCurrentServerOutputOptions(version core.RuntimeVersion, serverName string) *server.GetCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0254, core.Version0116:
		return &server.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &types.ServerOpts{
				Name: serverName,
				Type: types.ManagementClusterServerType,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0280:
		return &server.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0280,
			},
			ServerOpts: &types.ServerOpts{
				Name: serverName,
				Type: types.ManagementClusterServerType,
				GlobalOpts: &types.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	}
	return nil
}

// DefaultGetCurrentServerOutputOptionsWithError helper method to construct GetCurrentServer API output options with error
func DefaultGetCurrentServerOutputOptionsWithError(version core.RuntimeVersion) *server.GetCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		return &server.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: ServerNotFound,
		}
	}
	return nil
}

// DefaultRemoveCurrentServerInputOptions helper method to construct RemoveCurrentServer API input options
func DefaultRemoveCurrentServerInputOptions(version core.RuntimeVersion) *server.RemoveCurrentServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280:
		return &server.RemoveCurrentServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerName: common.CompatibilityTestOne,
		}
	}
	return nil
}

// DefaultRemoveCurrentServerOutputOptionsWithError helper method to construct RemoveCurrentServer API output option
func DefaultRemoveCurrentServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *server.RemoveCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		return &server.RemoveCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", serverName),
		}
	}
	return nil
}

// DefaultDeleteServerInputOptions helper method to construct DeleteServer API input options
func DefaultDeleteServerInputOptions(version core.RuntimeVersion, serverName string) *server.DeleteServerInputOptions {
	return &server.DeleteServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultDeleteServerOutputOptionsWithError helper method to construct DeleteServer API output options
func DefaultDeleteServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *server.DeleteServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		return &server.DeleteServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", serverName),
		}
	}
	return nil
}
