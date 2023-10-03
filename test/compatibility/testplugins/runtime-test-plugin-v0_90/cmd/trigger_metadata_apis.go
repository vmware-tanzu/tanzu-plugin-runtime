// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetMetadataAPI trigger Runtime GetMetadata API
func triggerGetMetadataAPI(_ *core.API) *core.APIResponse {
	// Trigger GetMetadata API
	return getMetadata()
}

// triggerGetConfigMetadataAPI trigger Runtime GetConfigMetadata API
func triggerGetConfigMetadataAPI(_ *core.API) *core.APIResponse {
	// Trigger GetConfigMetadata API
	return getConfigMetadata()
}

// triggerGetConfigMetadataPatchStrategyAPI trigger Runtime GetConfigMetadataPatchStrategy API
func triggerGetConfigMetadataPatchStrategyAPI(_ *core.API) *core.APIResponse {
	// Trigger GetConfigMetadataPatchStrategy API
	return getConfigMetadataPatchStrategy()
}

// triggerGetConfigMetadataSettingsAPI trigger Runtime GetConfigMetadataSettings API
func triggerGetConfigMetadataSettingsAPI(_ *core.API) *core.APIResponse {
	// Trigger GetConfigMetadataSettings API
	return getConfigMetadataSettings()
}

// triggerGetConfigMetadataSettingAPI trigger Runtime GetConfigMetadataSetting API
func triggerGetConfigMetadataSettingAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime GetConfigMetadataSetting API
	key, err := core.ParseStr(api.Arguments[core.Key])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Key, err),
		}
	}
	// Trigger GetConfigMetadataSetting API
	return getConfigMetadataSetting(key)
}

// triggerIsConfigMetadataSettingsEnabledAPI trigger Runtime IsConfigMetadataSettingsEnabled API
func triggerIsConfigMetadataSettingsEnabledAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime IsConfigMetadataSettingsEnabled API
	key, err := core.ParseStr(api.Arguments[core.Key])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Key, err),
		}
	}
	// Trigger IsConfigMetadataSettingsEnabled API
	return isConfigMetadataSettingsEnabled(key)
}

// triggerUseUnifiedConfigAPI trigger Runtime UseUnifiedConfig API
func triggerUseUnifiedConfigAPI(_ *core.API) *core.APIResponse {
	// Trigger UseUnifiedConfig API
	return useUnifiedConfig()
}

// triggerSetConfigMetadataSettingAPI trigger Runtime SetConfigMetadataSetting API
func triggerSetConfigMetadataSettingAPI(api *core.API) *core.APIResponse {
	keyName, valueName, response, done := parseMetadataArguments(api)
	if done {
		return response
	}
	// Trigger SetConfigMetadataSetting API
	return setConfigMetadataSetting(keyName, valueName)
}

// triggerSetConfigMetadataPatchStrategyAPI trigger Runtime SetConfigMetadataPatchStrategy API
func triggerSetConfigMetadataPatchStrategyAPI(api *core.API) *core.APIResponse {
	keyName, valueName, response, done := parseMetadataArguments(api)
	if done {
		return response
	}
	// Trigger SetConfigMetadataPatchStrategy API
	return setConfigMetadataPatchStrategy(keyName, valueName)
}

// triggerDeleteConfigMetadataSettingAPI trigger Runtime DeleteConfigMetadataSetting API
func triggerDeleteConfigMetadataSettingAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime DeleteConfigMetadataSetting API
	key, err := core.ParseStr(api.Arguments[core.Key])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.Key, err.Error()),
		}
	}
	// Trigger DeleteConfigMetadataSetting API
	return deleteConfigMetadataSetting(key)
}

func getMetadata() *core.APIResponse {
	metadata, err := configlib.GetMetadata()
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if metadata == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("metadata %v not found", metadata),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: metadata,
	}
}

func getConfigMetadata() *core.APIResponse {
	cfgMetadata, err := configlib.GetConfigMetadata()
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if cfgMetadata == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("config metadata %s not found", cfgMetadata),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: cfgMetadata,
	}
}

func getConfigMetadataPatchStrategy() *core.APIResponse {
	patchStrategy, err := configlib.GetConfigMetadataPatchStrategy()
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if len(patchStrategy) == 0 {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("config metadata patch strategy %s not found", patchStrategy),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: patchStrategy,
	}
}

func getConfigMetadataSettings() *core.APIResponse {
	settings, err := configlib.GetConfigMetadataSettings()
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if len(settings) == 0 {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("metadata settings not found"),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: settings,
	}
}

func getConfigMetadataSetting(key string) *core.APIResponse {
	val, err := configlib.GetConfigMetadataSetting(key)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	return &core.APIResponse{
		ResponseType: core.StringResponse,
		ResponseBody: val,
	}
}

func isConfigMetadataSettingsEnabled(key string) *core.APIResponse {
	enabled, err := configlib.IsConfigMetadataSettingsEnabled(key)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	return &core.APIResponse{
		ResponseType: core.BooleanResponse,
		ResponseBody: enabled,
	}
}

func useUnifiedConfig() *core.APIResponse {
	enabled, err := configlib.UseUnifiedConfig()
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	return &core.APIResponse{
		ResponseType: core.BooleanResponse,
		ResponseBody: enabled,
	}
}

func setConfigMetadataSetting(key, value string) *core.APIResponse {
	err := configlib.SetConfigMetadataSetting(key, value)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	return &core.APIResponse{
		ResponseBody: "",
		ResponseType: core.StringResponse,
	}
}

func setConfigMetadataPatchStrategy(key, value string) *core.APIResponse {
	err := configlib.SetConfigMetadataPatchStrategy(key, value)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	return &core.APIResponse{
		ResponseBody: "",
		ResponseType: core.StringResponse,
	}
}

func deleteConfigMetadataSetting(key string) *core.APIResponse {
	err := configlib.DeleteConfigMetadataSetting(key)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	return &core.APIResponse{
		ResponseBody: "",
		ResponseType: core.StringResponse,
	}
}

func parseMetadataArguments(api *core.API) (string, string, *core.APIResponse, bool) {
	keyName, err := core.ParseStr(api.Arguments[core.Key])
	if err != nil {
		return "", "", &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Key, err),
		}, true
	}

	valueName, err := core.ParseStr(api.Arguments[core.Value])
	if err != nil {
		return "", "", &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Value, err),
		}, true
	}
	return keyName, valueName, nil, false
}
