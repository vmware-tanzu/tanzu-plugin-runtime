// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerAPIs trigger runtime apis and construct logs
func triggerAPIs(apis []core.API) map[core.RuntimeAPIName][]core.APILog {
	// Variable used to store all the logging related to runtime api responses
	logs := make(map[core.RuntimeAPIName][]core.APILog)

	// Loop through array of commands
	for _, api := range apis {
		// Route to runtime API method call based on passed command value
		triggerContextAPIs(&api, logs)
	}

	return logs
}
