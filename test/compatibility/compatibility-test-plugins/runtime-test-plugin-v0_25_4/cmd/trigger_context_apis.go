// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configapi "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"gopkg.in/yaml.v3"
)

// triggerContextAPIs trigger context related runtime apis and construct logs
func triggerContextAPIs(api *compatibilitytestingframework.API, logs map[compatibilitytestingframework.RuntimeAPIName][]compatibilitytestingframework.APILog) {
	if api.Name == compatibilitytestingframework.SetContextAPIName {
		log := triggerSetContextAPI(api)
		logs[compatibilitytestingframework.SetContextAPIName] = append(logs[compatibilitytestingframework.SetContextAPIName], log)
	}
	if api.Name == compatibilitytestingframework.GetContextAPIName {
		log := triggerGetContextAPI(api)
		logs[compatibilitytestingframework.GetContextAPIName] = append(logs[compatibilitytestingframework.GetContextAPIName], log)
	}
}

// triggerGetContextAPI trigger get context runtime api
func triggerGetContextAPI(api *compatibilitytestingframework.API) compatibilitytestingframework.APILog {
	// Parse arguments needed to trigger the runtime api
	ctxName, err := compatibilitytestingframework.ParseStr(api.Arguments["contextName"])
	if err != nil {
		fmt.Println(err)
	}

	// Call runtime GetContext() API function
	ctx, err := configlib.GetContext(ctxName)

	// Construct logging
	log := compatibilitytestingframework.APILog{}
	if err != nil {
		log.APIError = err.Error()
	}
	log.APIResponse = &compatibilitytestingframework.APIResponse{
		ResponseBody: ctx,
		ResponseType: compatibilitytestingframework.MapResponse,
	}
	return log
}

// triggerSetContextAPI trigger set context runtime api
func triggerSetContextAPI(api *compatibilitytestingframework.API) compatibilitytestingframework.APILog {
	// Parse arguments needed to trigger the runtime api
	ctx, err := parseContext(api.Arguments["context"].(string))
	if err != nil {
		fmt.Println(err)
	}
	isCurrent := api.Arguments["isCurrent"].(bool)

	// Call the runtime SetContext() API function
	err = configlib.AddContext(ctx, isCurrent)

	// Construct logging
	log := compatibilitytestingframework.APILog{}
	if err != nil {
		log.APIError = err.Error()
	}
	log.APIResponse = &compatibilitytestingframework.APIResponse{
		ResponseBody: "",
		ResponseType: compatibilitytestingframework.StringResponse,
	}
	return log
}

// parseContext unmarshalls string to Context struct
func parseContext(context string) (*configapi.Context, error) {
	var ctx configapi.Context
	err := yaml.Unmarshal([]byte(context), &ctx)
	if err != nil {
		return nil, err
	}
	return &ctx, nil
}
