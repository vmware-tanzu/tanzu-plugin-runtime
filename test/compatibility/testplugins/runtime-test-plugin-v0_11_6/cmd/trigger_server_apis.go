// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetServerAPI trigger get server runtime api
func triggerGetServerAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	serverName, err := core.ParseStr(api.Arguments[core.ServerName])
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ServerName, err.Error()),
		}
		return log
	}

	// Call runtime GetServer API
	server, err := configlib.GetServer(serverName)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	if server != nil {
		log.APIResponse = &core.APIResponse{
			ResponseBody: server,
			ResponseType: core.MapResponse,
		}
	}

	return log
}

// triggerAddServerAPI trigger add server runtime api
func triggerAddServerAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	server, err := parseServer(api.Arguments[core.Server].(string))
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse server from argument %v with error %v ", core.Server, err.Error()),
		}
		return log
	}

	setCurrent := api.Arguments[core.SetCurrent].(bool)

	// Call runtime SetServer API
	err = configlib.AddServer(server, setCurrent)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &core.APIResponse{
			ResponseBody: "",
			ResponseType: core.StringResponse,
		}
	}
	return log
}
