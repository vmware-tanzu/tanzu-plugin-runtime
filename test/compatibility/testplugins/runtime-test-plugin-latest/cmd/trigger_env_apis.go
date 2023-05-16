// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetEnvConfigurationsAPI trigger Runtime GetEnvConfigurations API
func triggerGetEnvConfigurationsAPI(_ *core.API) *core.APIResponse {
	// Trigger GetEnvConfigurations API
	return getEnvConfigurations()
}

// triggerGetEnvAPI trigger Runtime GetEnv API
func triggerGetEnvAPI(api *core.API) *core.APIResponse {
	keyName, err := parseEnvArguments(api)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err,
		}
	}

	// Trigger GetEnv API
	return getEnv(keyName)
}

// triggerSetEnvAPI trigger Runtime SetEnv API
func triggerSetEnvAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime SetEnv API
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
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Value, err),
		}
	}
	// Trigger SetEnv API
	return setEnv(keyName, valueName)
}

// triggerDeleteEnvAPI trigger Runtime DeleteEnv API
func triggerDeleteEnvAPI(api *core.API) *core.APIResponse {
	keyName, err := parseEnvArguments(api)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err,
		}
	}

	// Trigger DeleteEnv API
	return deleteEnv(keyName)
}

func getEnvConfigurations() *core.APIResponse {
	envs := configlib.GetEnvConfigurations()

	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: envs,
	}
}

func getEnv(keyName string) *core.APIResponse {
	enabled, err := configlib.GetEnv(keyName)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	return &core.APIResponse{
		ResponseType: core.StringResponse,
		ResponseBody: enabled,
	}
}

func setEnv(keyName, valueName string) *core.APIResponse {
	err := configlib.SetEnv(keyName, valueName)
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

func deleteEnv(keyName string) *core.APIResponse {
	err := configlib.DeleteEnv(keyName)
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

func parseEnvArguments(api *core.API) (string, error) {
	// Parse arguments needed to trigger the Runtime GetEnv API
	keyName, err := core.ParseStr(api.Arguments[core.Key])
	if err != nil {
		return "",
			fmt.Errorf("failed to parse string from argument %v with error %v", core.Key, err)
	}

	return keyName, nil
}
