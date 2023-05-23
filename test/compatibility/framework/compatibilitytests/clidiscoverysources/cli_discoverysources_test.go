// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clidiscoverysources_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/clidiscoverysources"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version CLI Discovery Source APIs Compatibility Tests for supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest", func() {
	ginkgo.GinkgoWriter.Println("Get/Set/Delete CLI Discovery Source API methods are tested for cross-version API compatibility with supported Runtime versions v0.11.6, v0.25.4, v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})

	})

	ginkgo.Context("Run SetCLIDiscoverySource, GetCLIDiscoverySource, DeleteCLIDiscoverySource on all supported Runtime library versions latest, v0.28.0", func() {

		ginkgo.It("Run SetCLIDiscoverySource of Runtime latest then GetCLIDiscoverySource on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource on all supported Runtime library versions", func() {
			// Add SetCLIDiscoverySource Commands of Runtime Latest version
			testCase := core.NewTestCase().Add(clidiscoverysources.DefaultSetCLIDiscoverySourceCommand(core.VersionLatest))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.VersionLatest, clidiscoverysources.WithStrictValidationStrategy())).Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0280, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Add DeleteCLIDiscoverySource v0.28.0 Command
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.Version0280))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.VersionLatest, clidiscoverysources.WithStrictValidationStrategy())).Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0280, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetCLIDiscoverySource of Runtime v0.28.0 then GetCLIDiscoverySource, GetCurrentCLIDiscoverySource on all supported Runtime library versions and then DeleteCLIDiscoverySource of Runtime v0.11.6 then GetCLIDiscoverySource, GetCurrentCLIDiscoverySource on all supported Runtime library versions", func() {
			// Add SetCLIDiscoverySource Commands of Runtime Latest version
			testCase := core.NewTestCase().Add(clidiscoverysources.DefaultSetCLIDiscoverySourceCommand(core.Version0280))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.VersionLatest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound))).Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0280, clidiscoverysources.WithStrictValidationStrategy()))

			// Add DeleteCLIDiscoverySource v0.28.0 Command
			testCase.Add(clidiscoverysources.DefaultDeleteCLIDiscoverySourceCommand(core.VersionLatest))

			// Add GetCLIDiscoverySource Commands on all supported Runtime library versions
			testCase.Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.VersionLatest, clidiscoverysources.WithError(clidiscoverysources.CLIDiscoverySourceNotFound))).Add(clidiscoverysources.DefaultGetCLIDiscoverySourceCommand(core.Version0280, clidiscoverysources.WithStrictValidationStrategy()))

			// Run all the commands
			executer.Execute(testCase)
		})

	})

	// TODO: Additional tests to be added usign StoreClientConfig and GetClientConfig

})
