package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"gopkg.in/yaml.v3"
)

// triggerContextAPIs trigger context related runtime apis and construct logs
func triggerContextAPIs(api *compatibilitytestingframework.API, logs map[compatibilitytestingframework.RuntimeAPIName][]compatibilitytestingframework.APILog) {
	// If API name is SetContext then trigger SetContext() API
	if api.Name == compatibilitytestingframework.SetContextAPIName {
		log := triggerSetContextAPI(api)
		logs[compatibilitytestingframework.SetContextAPIName] = append(logs[compatibilitytestingframework.SetContextAPIName], log)
	}
	// If API name is GetContext then trigger GetContext() API
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
		fmt.Println("triggerGetContextAPI1", err)
	}

	//Call runtime GetContext() API function
	ctx, err := configlib.GetContext(ctxName)

	// Construct logging
	log := compatibilitytestingframework.APILog{}
	if err != nil {
		log.APIError = err.Error()
		log.APIResponse = &compatibilitytestingframework.APIResponse{
			ResponseType: compatibilitytestingframework.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	if ctx != nil {
		log.APIResponse = &compatibilitytestingframework.APIResponse{
			ResponseBody: ctx,
			ResponseType: compatibilitytestingframework.MapResponse,
		}
	}

	return log
}

// triggerSetContextAPI trigger set context runtime api
func triggerSetContextAPI(api *compatibilitytestingframework.API) compatibilitytestingframework.APILog {
	// Parse arguments needed to trigger the runtime api
	ctx, err := parseContext(api.Arguments["context"].(string))
	if err != nil {
		fmt.Println("triggerGetContextAPI2", err)
	}
	isCurrent := api.Arguments["isCurrent"].(bool)

	// Call the runtime SetContextAPIName API function
	err = configlib.SetContext(ctx, isCurrent)

	// Construct logging
	log := compatibilitytestingframework.APILog{}
	if err != nil {
		log.APIError = err.Error()
		log.APIResponse = &compatibilitytestingframework.APIResponse{
			ResponseType: compatibilitytestingframework.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &compatibilitytestingframework.APIResponse{
			ResponseBody: "",
			ResponseType: compatibilitytestingframework.StringResponse,
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
