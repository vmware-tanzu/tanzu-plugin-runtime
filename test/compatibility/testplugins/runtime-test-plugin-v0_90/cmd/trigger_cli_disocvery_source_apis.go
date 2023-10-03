// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/pkg/errors"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetCLIDiscoverySourceAPI trigger get cli discovery source name runtime api
func triggerGetCLIDiscoverySourceAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	name, err := core.ParseStr(api.Arguments[core.Name])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.Name, err.Error()),
		}
	}
	return getCLIDiscoverySource(name)
}

// triggerSetCLIDiscoverySourceAPI trigger add server runtime api
func triggerSetCLIDiscoverySourceAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	source, err := parseCLIDiscoverySource(api.Arguments[core.DiscoverySource].(string))
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse discovery source from argument %v with error %v ", core.DiscoverySource, err.Error()),
		}
	}
	return setCLIDiscoverySource(source)
}

// triggerDeleteCLIDiscoverySourceAPI trigger remove context runtime api
func triggerDeleteCLIDiscoverySourceAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	name, err := core.ParseStr(api.Arguments[core.Name])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.Name, err.Error()),
		}
	}
	return deleteCLIDiscoverySource(name)
}

func getCLIDiscoverySource(name string) *core.APIResponse {
	// Call runtime GetCLIDiscoverySource API
	source, err := configlib.GetCLIDiscoverySource(name)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if source == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("source %v not found", source),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: source,
	}
}

func setCLIDiscoverySource(source *configtypes.PluginDiscovery) *core.APIResponse {
	// Call runtime SetCLIDiscoverySource API
	err := configlib.SetCLIDiscoverySource(*source)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: errors.Wrap(err, "failed"),
		}
	}
	return &core.APIResponse{
		ResponseBody: "",
		ResponseType: core.StringResponse,
	}
}

func deleteCLIDiscoverySource(name string) *core.APIResponse {
	// Call runtime RemoveCLIDiscoverySource API
	err := configlib.DeleteCLIDiscoverySource(name)

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
