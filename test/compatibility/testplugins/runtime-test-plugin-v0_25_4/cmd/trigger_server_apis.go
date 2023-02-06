// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configapi "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetServerAPI trigger get server runtime api
func triggerGetServerAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	serverName, err := core.ParseStr(api.Arguments[core.ServerName])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ServerName, err.Error()),
		}
	}
	return getServer(serverName)
}

// triggerSetServerAPI trigger add server runtime api
func triggerSetServerAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	server, err := parseServer(api.Arguments[core.Server].(string))
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse server from argument %v with error %v ", core.Server, err.Error()),
		}
	}
	setCurrent := api.Arguments[core.SetCurrent].(bool)
	return setServer(server, setCurrent)
}

// triggerAddServerAPI trigger add server runtime api
func triggerAddServerAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	server, err := parseServer(api.Arguments[core.Server].(string))
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse server from argument %v with error %v ", core.Server, err.Error()),
		}
	}
	setCurrent := api.Arguments[core.SetCurrent].(bool)
	return addServer(server, setCurrent)
}

// triggerRemoveServerAPI trigger remove context runtime api
func triggerRemoveServerAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	serverName, err := core.ParseStr(api.Arguments[core.ServerName])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ServerName, err.Error()),
		}
	}
	return removeServer(serverName)
}

// triggerSetCurrentServerAPI trigger remove context runtime api
func triggerSetCurrentServerAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the runtime api
	serverName, err := core.ParseStr(api.Arguments[core.ServerName])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ServerName, err.Error()),
		}
	}
	return setCurrentServer(serverName)
}

// triggerGetCurrentServerAPI trigger remove context runtime api
func triggerGetCurrentServerAPI(*core.API) *core.APIResponse {
	return getCurrentServer()
}

func setServer(server *configapi.Server, setCurrent bool) *core.APIResponse {
	// Call runtime PutServer API
	err := configlib.PutServer(server, setCurrent)
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

func addServer(server *configapi.Server, setCurrent bool) *core.APIResponse {
	// Call runtime AddServer API
	err := configlib.AddServer(server, setCurrent)
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

func removeServer(serverName string) *core.APIResponse {
	// Call runtime RemoveServer API
	err := configlib.RemoveServer(serverName)

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

func getCurrentServer() *core.APIResponse {
	server, err := configlib.GetCurrentServer()
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if server == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("server not found"),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: server,
	}
}

func setCurrentServer(serverName string) *core.APIResponse {
	// Call runtime SetCurrentServer API
	err := configlib.SetCurrentServer(serverName)
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

func getServer(serverName string) *core.APIResponse {
	// Call runtime GetServer API
	server, err := configlib.GetServer(serverName)

	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if server == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("server %v not found", server),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: server,
	}
}
