// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"gopkg.in/yaml.v3"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerContextAPIs trigger context related runtime apis and construct logs
func triggerContextAPIs(api *core.API, logs map[core.RuntimeAPIName][]core.APILog) {
	// If API name is SetContext then trigger SetContext() API
	if api.Name == core.SetContextAPIName {
		log := triggerSetContextAPI(api)
		logs[core.SetContextAPIName] = append(logs[core.SetContextAPIName], log)
	}
	// If API name is GetContext then trigger GetContext() API
	if api.Name == core.GetContextAPIName {
		log := triggerGetContextAPI(api)
		logs[core.GetContextAPIName] = append(logs[core.GetContextAPIName], log)
	}
}

// triggerGetContextAPI trigger get context runtime api
func triggerGetContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		fmt.Println(err)
	}

	// Call runtime GetContext API
	ctx, err := configlib.GetContext(ctxName)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	if ctx != nil {
		log.APIResponse = &core.APIResponse{
			ResponseBody: ctx,
			ResponseType: core.MapResponse,
		}
	}
	return log
}

// triggerSetContextAPI trigger set context runtime api
func triggerSetContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	ctx, err := parseContext(api.Arguments[core.Context].(string))
	if err != nil {
		fmt.Println(err)
	}
	setCurrent := api.Arguments[core.SetCurrent].(bool)

	// Call runtime SetContext API
	err = configlib.SetContext(ctx, setCurrent)

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

// parseContext unmarshalls string to Context struct
func parseContext(context string) (*configtypes.Context, error) {
	var ctx configtypes.Context
	err := yaml.Unmarshal([]byte(context), &ctx)
	if err != nil {
		return nil, err
	}
	return &ctx, nil
}
