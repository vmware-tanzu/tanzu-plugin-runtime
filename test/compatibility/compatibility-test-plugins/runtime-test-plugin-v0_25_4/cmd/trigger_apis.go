// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerAPIs trigger runtime apis and construct logs
func triggerAPIs(apis []compatibilitytestingframework.API) map[compatibilitytestingframework.RuntimeAPIName][]compatibilitytestingframework.APILog {
	// Variable used to store all the logging related to runtime api responses
	logs := make(map[compatibilitytestingframework.RuntimeAPIName][]compatibilitytestingframework.APILog)

	// Loop through array of commands
	for index := range apis {
		// Route to runtime API method call based on passed command value
		triggerContextAPIs(&apis[index], logs)
	}

	return logs
}
