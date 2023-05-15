// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerIsFeatureEnabledAPI trigger Runtime IsFeatureEnabled API
func triggerIsFeatureEnabledAPI(api *core.API) *core.APIResponse {
	pluginName, keyName, err := parseFeatureArguments(api)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err,
		}
	}

	// Trigger IsFeatureEnabled API
	return isFeatureEnabled(pluginName, keyName)
}

// triggerSetFeatureAPI trigger Runtime SetFeature API
func triggerSetFeatureAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime SetFeature API
	pluginName, err := core.ParseStr(api.Arguments[core.Plugin])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Plugin, err),
		}
	}

	keyName, err := core.ParseStr(api.Arguments[core.Key])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Key, err),
		}
	}

	valueName, err := core.ParseStr(api.Arguments[core.Value])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Key, err),
		}
	}
	// Trigger SetFeature API
	return setFeature(pluginName, keyName, valueName)
}

// triggerDeleteFeatureAPI trigger Runtime DeleteFeature API
func triggerDeleteFeatureAPI(api *core.API) *core.APIResponse {
	pluginName, keyName, err := parseFeatureArguments(api)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err,
		}
	}

	// Trigger DeleteFeature API
	return deleteFeature(pluginName, keyName)
}

func isFeatureEnabled(pluginName, keyName string) *core.APIResponse {
	enabled, err := configlib.IsFeatureEnabled(pluginName, keyName)
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

func setFeature(pluginName, keyName, valueName string) *core.APIResponse {
	err := configlib.SetFeature(pluginName, keyName, valueName)
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

func deleteFeature(pluginName, keyName string) *core.APIResponse {
	err := configlib.DeleteFeature(pluginName, keyName)
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

func parseFeatureArguments(api *core.API) (string, string, error) {
	// Parse arguments needed to trigger the Runtime IsFeatureEnabled API
	pluginName, err := core.ParseStr(api.Arguments[core.Plugin])
	if err != nil {
		return "", "",
			fmt.Errorf("failed to parse string from argument %v with error %v", core.Plugin, err)
	}

	keyName, err := core.ParseStr(api.Arguments[core.Key])
	if err != nil {
		return "", "",
			fmt.Errorf("failed to parse string from argument %v with error %v", core.Key, err)
	}

	return pluginName, keyName, nil
}
