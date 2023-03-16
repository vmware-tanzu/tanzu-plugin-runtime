// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerAPIs trigger all runtime apis and construct logs as per specified array of apis
func triggerAPIs(apis []core.API) map[core.RuntimeAPIName][]core.APILog {
	// Variable used to store all the logging related to runtime api responses
	logs := make(map[core.RuntimeAPIName][]core.APILog)

	// Loop through array of commands
	for index := range apis {
		// Route to runtime API method call based on passed command value
		triggerServerAPIs(&apis[index], logs)
	}
	return logs
}

// triggerServerAPIs trigger server related runtime apis and construct response logs
func triggerServerAPIs(api *core.API, logs map[core.RuntimeAPIName][]core.APILog) {
	if api.Name == core.AddServerAPIName {
		log := triggerAddServerAPI(api)
		logs[core.AddServerAPIName] = append(logs[core.AddServerAPIName], log)
	}
	if api.Name == core.GetServerAPIName {
		log := triggerGetServerAPI(api)
		logs[core.GetServerAPIName] = append(logs[core.GetServerAPIName], log)
	}
}
