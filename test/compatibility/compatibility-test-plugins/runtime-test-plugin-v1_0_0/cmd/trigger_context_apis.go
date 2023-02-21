// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	compatibilitytestingcore "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"gopkg.in/yaml.v3"
)

// triggerContextAPIs trigger context related runtime apis and construct logs
func triggerContextAPIs(api *compatibilitytestingcore.API, logs map[compatibilitytestingcore.RuntimeAPIName][]compatibilitytestingcore.APILog) {
	// If API name is SetContext then trigger SetContext() API
	if api.Name == compatibilitytestingcore.SetContextAPIName {
		log := triggerSetContextAPI(api)
		logs[compatibilitytestingcore.SetContextAPIName] = append(logs[compatibilitytestingcore.SetContextAPIName], log)
	}
	// If API name is GetContext then trigger GetContext() API
	if api.Name == compatibilitytestingcore.GetContextAPIName {
		log := triggerGetContextAPI(api)
		logs[compatibilitytestingcore.GetContextAPIName] = append(logs[compatibilitytestingcore.GetContextAPIName], log)
	}
}

// triggerGetContextAPI trigger get context runtime api
func triggerGetContextAPI(api *compatibilitytestingcore.API) compatibilitytestingcore.APILog {
	// Parse arguments needed to trigger the runtime api
	ctxName, err := compatibilitytestingcore.ParseStr(api.Arguments["contextName"])
	if err != nil {
		fmt.Println("triggerGetContextAPI1", err)
	}

	//Call runtime GetContext() API function
	ctx, err := configlib.GetContext(ctxName)

	// Construct logging
	log := compatibilitytestingcore.APILog{}
	if err != nil {
		log.APIError = err.Error()
		log.APIResponse = &compatibilitytestingcore.APIResponse{
			ResponseType: compatibilitytestingcore.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	if ctx != nil {
		log.APIResponse = &compatibilitytestingcore.APIResponse{
			ResponseBody: ctx,
			ResponseType: compatibilitytestingcore.MapResponse,
		}
	}

	return log
}

// triggerSetContextAPI trigger set context runtime api
func triggerSetContextAPI(api *compatibilitytestingcore.API) compatibilitytestingcore.APILog {
	// Parse arguments needed to trigger the runtime api
	ctx, err := parseContext(api.Arguments["context"].(string))
	if err != nil {
		fmt.Println("triggerGetContextAPI2", err)
	}
	isCurrent := api.Arguments["isCurrent"].(bool)

	// Call the runtime SetContextAPIName API function
	err = configlib.SetContext(ctx, isCurrent)

	// Construct logging
	log := compatibilitytestingcore.APILog{}
	if err != nil {
		log.APIError = err.Error()
		log.APIResponse = &compatibilitytestingcore.APIResponse{
			ResponseType: compatibilitytestingcore.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &compatibilitytestingcore.APIResponse{
			ResponseBody: "",
			ResponseType: compatibilitytestingcore.StringResponse,
		}
	}

	return log
}

// parseContext unmarshalls string to Context struct
func parseContext(context string) (*configtypes.Context, error) {
	var ctx configtypes.Context
	err := yaml.Unmarshal([]byte(context), &ctx)
	if err != nil {
		return nil, err
	}
	return &ctx, nil
}
