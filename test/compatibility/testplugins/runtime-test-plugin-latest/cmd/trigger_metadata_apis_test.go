// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

const (
	CompatibilityTestsMetadataPatchStrategyKey   = "compatibility-tests.contexts.name"
	CompatibilityTestsMetadataPatchStrategyValue = "replace"
	CompatibilityTestsMetadataSettingsKey        = "useUnifiedConfig"
	CompatibilityTestsMetadataSettingsValue      = "true"
)

func TestTriggerMetadataAPIs(t *testing.T) {
	_, cleanup := core.SetupTempCfgFiles()
	defer func() {
		cleanup()
	}()

	var tests = []struct {
		name         string
		apiName      core.RuntimeAPIName
		apis         []core.API
		expectedLogs map[core.RuntimeAPIName][]core.APILog
	}{
		{
			name:    "Trigger SetConfigMetadataSetting API",
			apiName: core.SetConfigMetadataSettingAPI,
			apis: []core.API{
				{
					Name:    core.SetConfigMetadataSettingAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Key:   CompatibilityTestsMetadataSettingsKey,
						core.Value: CompatibilityTestsMetadataSettingsValue,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetConfigMetadataSettingAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
				},
			},
		},

		{
			name:    "Trigger SetConfigMetadataPatchStrategy API",
			apiName: core.SetConfigMetadataPatchStrategyAPI,
			apis: []core.API{
				{
					Name:    core.SetConfigMetadataPatchStrategyAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Key:   CompatibilityTestsMetadataPatchStrategyKey,
						core.Value: CompatibilityTestsMetadataPatchStrategyValue,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetConfigMetadataPatchStrategyAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
				},
			},
		},

		{
			name:    "Trigger GetMetadata API",
			apiName: core.GetMetadataAPI,
			apis: []core.API{
				{
					Name:    core.GetMetadataAPI,
					Version: core.VersionLatest,
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetMetadataAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.MapResponse,
							ResponseBody: &types.Metadata{
								ConfigMetadata: &types.ConfigMetadata{
									PatchStrategy: map[string]string{CompatibilityTestsMetadataPatchStrategyKey: CompatibilityTestsMetadataPatchStrategyValue},
									Settings:      map[string]string{CompatibilityTestsMetadataSettingsKey: CompatibilityTestsMetadataSettingsValue},
								},
							},
						},
					},
				},
			},
		},

		{
			name:    "Trigger GetConfigMetadata API",
			apiName: core.GetConfigMetadataAPI,
			apis: []core.API{
				{
					Name:    core.GetConfigMetadataAPI,
					Version: core.VersionLatest,
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetConfigMetadataAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.MapResponse,
							ResponseBody: &types.ConfigMetadata{
								PatchStrategy: map[string]string{CompatibilityTestsMetadataPatchStrategyKey: CompatibilityTestsMetadataPatchStrategyValue},
								Settings:      map[string]string{CompatibilityTestsMetadataSettingsKey: CompatibilityTestsMetadataSettingsValue},
							},
						},
					},
				},
			},
		},

		{
			name:    "Trigger GetConfigMetadataPatchStrategy API",
			apiName: core.GetConfigMetadataPatchStrategyAPI,
			apis: []core.API{
				{
					Name:    core.GetConfigMetadataPatchStrategyAPI,
					Version: core.VersionLatest,
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetConfigMetadataPatchStrategyAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.MapResponse,
							ResponseBody: map[string]string{CompatibilityTestsMetadataPatchStrategyKey: CompatibilityTestsMetadataPatchStrategyValue},
						},
					},
				},
			},
		},

		{
			name:    "Trigger GetConfigMetadataSettings API",
			apiName: core.GetConfigMetadataSettingsAPI,
			apis: []core.API{
				{
					Name:    core.GetConfigMetadataSettingsAPI,
					Version: core.VersionLatest,
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetConfigMetadataSettingsAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.MapResponse,
							ResponseBody: map[string]string{CompatibilityTestsMetadataSettingsKey: CompatibilityTestsMetadataSettingsValue},
						},
					},
				},
			},
		},

		{
			name:    "Trigger GetConfigMetadataSetting API",
			apiName: core.GetConfigMetadataSettingAPI,
			apis: []core.API{
				{
					Name:    core.GetConfigMetadataSettingAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Key: CompatibilityTestsMetadataSettingsKey,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetConfigMetadataSettingAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.StringResponse,
							ResponseBody: "true",
						},
					},
				},
			},
		},

		{
			name:    "Trigger IsConfigMetadataSettingsEnabledAPI API",
			apiName: core.IsConfigMetadataSettingsEnabledAPI,
			apis: []core.API{
				{
					Name:    core.IsConfigMetadataSettingsEnabledAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Key: CompatibilityTestsMetadataSettingsKey,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.IsConfigMetadataSettingsEnabledAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.BooleanResponse,
							ResponseBody: true,
						},
					},
				},
			},
		},

		{
			name:    "Trigger UseUnifiedConfigAPI API",
			apiName: core.UseUnifiedConfigAPI,
			apis: []core.API{
				{
					Name:    core.UseUnifiedConfigAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Key: CompatibilityTestsMetadataSettingsKey,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.UseUnifiedConfigAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.BooleanResponse,
							ResponseBody: true,
						},
					},
				},
			},
		},

		{
			name:    "Trigger DeleteConfigMetadataSettingAPI API",
			apiName: core.DeleteConfigMetadataSettingAPI,
			apis: []core.API{
				{
					Name:    core.DeleteConfigMetadataSettingAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Key: CompatibilityTestsMetadataSettingsKey,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.DeleteConfigMetadataSettingAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
				},
			},
		},

		{
			name:    "Trigger UseUnifiedConfigAPI API",
			apiName: core.UseUnifiedConfigAPI,
			apis: []core.API{
				{
					Name:    core.UseUnifiedConfigAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Key: CompatibilityTestsMetadataSettingsKey,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.UseUnifiedConfigAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.ErrorResponse,
							ResponseBody: "not found",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualLogs := triggerAPIs(tt.apis)
			assert.Equal(t, tt.expectedLogs[tt.apiName], actualLogs[tt.apiName])
		})
	}
}
