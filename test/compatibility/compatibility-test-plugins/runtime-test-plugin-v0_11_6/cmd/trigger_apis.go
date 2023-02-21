// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	compatibilitytestingtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerAPIs trigger all runtime apis and construct logs as per specified array of apis
func triggerAPIs(apis []compatibilitytestingtypes.API) map[compatibilitytestingtypes.RuntimeAPIName][]compatibilitytestingtypes.APILog {
	// Variable used to store all the logging related to runtime api responses
	logs := make(map[compatibilitytestingtypes.RuntimeAPIName][]compatibilitytestingtypes.APILog)

	// Loop through array of commands
	for _, api := range apis {
		// Route to runtime API method call based on passed command value
		triggerServerAPIs(&api, logs)
	}
	return logs
}
