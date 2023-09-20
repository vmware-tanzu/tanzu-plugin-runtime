// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetClientConfigAPI trigger Runtime GetClientConfig API
func triggerGetClientConfigAPI(_ *core.API) *core.APIResponse {
	// Trigger GetClientConfig API
	return getClientConfig()
}

// triggerStoreClientConfigAPI trigger Runtime StoreClientConfig API
func triggerStoreClientConfigAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime AddContext API
	cfg, err := parseClientConfig(api.Arguments[core.ClientConfig].(string))
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.ClientConfig, err),
		}
	}
	// Trigger AddContext API
	return storeClientConfig(cfg)
}

func getClientConfig() *core.APIResponse {
	cfg, err := configlib.GetClientConfig()
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if cfg == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("clientconfig %v not found", cfg),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: cfg,
	}
}

func storeClientConfig(cfg *configtypes.ClientConfig) *core.APIResponse {
	err := configlib.StoreClientConfig(cfg)
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
