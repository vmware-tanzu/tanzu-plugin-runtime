// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerAPIs trigger runtime apis and construct logs
func triggerAPIs(apis []core.API) map[core.RuntimeAPIName][]core.APILog {
	// Variable used to store all the logging related to runtime api responses
	logs := make(map[core.RuntimeAPIName][]core.APILog)

	// Loop through array of commands
	for index := range apis {
		// Route to runtime API method call based on passed command value
		triggerContextAPIs(&apis[index], logs)
	}

	return logs
}

// triggerContextAPIs trigger context related runtime apis and construct logs
func triggerContextAPIs(api *core.API, logs map[core.RuntimeAPIName][]core.APILog) {

	switch api.Name {
	case core.SetContextAPIName:
		log := triggerSetContextAPI(api)
		logs[core.SetContextAPIName] = append(logs[core.SetContextAPIName], log)

	case core.GetContextAPIName:
		log := triggerGetContextAPI(api)
		logs[core.GetContextAPIName] = append(logs[core.GetContextAPIName], log)

	case core.RemoveContextAPIName:
		log := triggerRemoveContextAPI(api)
		logs[core.RemoveContextAPIName] = append(logs[core.RemoveContextAPIName], log)

	case core.DeleteContextAPIName:
		log := triggerRemoveContextAPI(api)
		logs[core.DeleteContextAPIName] = append(logs[core.DeleteContextAPIName], log)

	case core.SetCurrentContextAPIName:
		log := triggerSetCurrentContextAPI(api)
		logs[core.SetCurrentContextAPIName] = append(logs[core.SetCurrentContextAPIName], log)

	case core.GetCurrentContextAPIName:
		log := triggerGetCurrentContextAPI(api)
		logs[core.GetCurrentContextAPIName] = append(logs[core.GetCurrentContextAPIName], log)

	default:
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("command %v not found", api.Name),
		}
		logs[api.Name] = append(logs[api.Name], log)
	}
}
