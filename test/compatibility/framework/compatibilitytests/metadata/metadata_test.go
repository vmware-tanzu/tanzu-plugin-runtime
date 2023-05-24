// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package metadata_test

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/executer"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/metadata"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Cross-version ConfigMetadata Flags APIs compatibility tests", func() {
	ginkgo.GinkgoWriter.Println("SetConfigMetadataPatchStrategy, SetConfigMetadataSetting, GetMetadata, GetConfigMetadata, GetConfigMetadataPatchStrategy, GetConfigMetadataSettings, GetConfigMetadataSetting,IsConfigMetadataSettingsEnabled, UseUnifiedConfig,  DeleteConfigMetadataSetting methods are tested for cross-version API compatibility with supported Runtime versions v0.28.0, latest")

	ginkgo.BeforeEach(func() {
		// Setup mock temporary config files for testing
		_, cleanup := core.SetupTempCfgFiles()
		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("using default metadata", func() {

		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting latest then Getters for v0.28.0, latest then DeleteConfigMetadataSetting v0.28.0 then Getters for v0.28.0, latest", func() {
			// Build test case with commands

			// Add SetConfigMetadata Commands of Runtime Latest
			testCase := core.NewTestCase()

			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.VersionLatest))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.VersionLatest))

			// Add GetConfigMetadata latest, v0.28.0 Commands
			testCase.Add(metadata.DefaultGetMetadataCommand(core.VersionLatest)).Add(metadata.DefaultGetMetadataCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.Version0280))
			testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.VersionLatest)).Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0280))
			testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.VersionLatest)).Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0280))

			// Add DeleteConfigMetadata v0.28.0 Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.Version0280))

			// Add GetConfigMetadata latest, v0.28.0 Commands
			metadataOpts := &framework.MetadataOpts{
				ConfigMetadata: &framework.ConfigMetadataOpts{
					Settings: map[string]string{},
				},
			}

			testCase.Add(metadata.DefaultGetMetadataCommand(core.VersionLatest, metadata.WithMetadataOpts(metadataOpts))).Add(metadata.DefaultGetMetadataCommand(core.Version0280, metadata.WithMetadataOpts(metadataOpts)))
			testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.VersionLatest, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata))).Add(metadata.DefaultGetConfigMetadataCommand(core.Version0280, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))
			testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.VersionLatest, metadata.WithError(common.ErrNotFound))).Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0280, metadata.WithError(common.ErrNotFound)))
			testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.VersionLatest, metadata.WithError(common.ErrNotFound))).Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0280, metadata.WithError(common.ErrNotFound)))
			testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.VersionLatest, metadata.WithError(common.ErrNotFound))).Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0280, metadata.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})

		ginkgo.It("Run SetConfigMetadataPatchStrategy, SetConfigMetadataSetting v0.28.0 then Getters for v0.28.0, latest then DeleteConfigMetadataSetting latest then Getters for v0.28.0, latest", func() {
			// Build test case with commands
			testCase := core.NewTestCase()

			// Add SetConfigMetadata Commands of Runtime v0.28.0
			testCase.Add(metadata.DefaultSetConfigMetadataPatchStrategyCommand(core.Version0280))
			testCase.Add(metadata.DefaultSetConfigMetadataSettingCommand(core.Version0280))

			// Add GetConfigMetadata latest, v0.28.0 Commands
			testCase.Add(metadata.DefaultGetMetadataCommand(core.VersionLatest)).Add(metadata.DefaultGetMetadataCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataSettingsCommand(core.Version0280))
			testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.VersionLatest)).Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0280))
			testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.VersionLatest)).Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0280))

			// Add DeleteConfigMetadataSetting latest Command
			testCase.Add(metadata.DefaultDeleteConfigMetadataSettingCommand(core.VersionLatest))

			// Add GetConfigMetadata latest, v0.28.0 Commands
			metadataOpts := &framework.MetadataOpts{
				ConfigMetadata: &framework.ConfigMetadataOpts{
					Settings: map[string]string{},
				},
			}

			testCase.Add(metadata.DefaultGetMetadataCommand(core.VersionLatest, metadata.WithMetadataOpts(metadataOpts))).Add(metadata.DefaultGetMetadataCommand(core.Version0280, metadata.WithMetadataOpts(metadataOpts)))
			testCase.Add(metadata.DefaultGetConfigMetadataCommand(core.VersionLatest, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata))).Add(metadata.DefaultGetConfigMetadataCommand(core.Version0280, metadata.WithConfigMetadataOpts(metadataOpts.ConfigMetadata)))
			testCase.Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.VersionLatest)).Add(metadata.DefaultGetConfigMetadataPatchStrategyCommand(core.Version0280))
			testCase.Add(metadata.DefaultGetConfigMetadataSettingCommand(core.VersionLatest, metadata.WithError(common.ErrNotFound))).Add(metadata.DefaultGetConfigMetadataSettingCommand(core.Version0280, metadata.WithError(common.ErrNotFound)))
			testCase.Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.VersionLatest, metadata.WithError(common.ErrNotFound))).Add(metadata.DefaultIsConfigMetadataSettingsEnabledCommand(core.Version0280, metadata.WithError(common.ErrNotFound)))
			testCase.Add(metadata.DefaultUseUnifiedConfigCommand(core.VersionLatest, metadata.WithError(common.ErrNotFound))).Add(metadata.DefaultUseUnifiedConfigCommand(core.Version0280, metadata.WithError(common.ErrNotFound)))

			// Run all the commands
			executer.Execute(testCase)
		})
	})

})
