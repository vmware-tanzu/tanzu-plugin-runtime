package cmd

import (
	compatibilitytestingtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

// triggerAPIs trigger runtime apis and construct logs
func triggerAPIs(apis []compatibilitytestingtypes.API) map[compatibilitytestingtypes.RuntimeAPIName][]compatibilitytestingtypes.APILog {
	// Variable used to store all the logging related to runtime api responses
	logs := make(map[compatibilitytestingtypes.RuntimeAPIName][]compatibilitytestingtypes.APILog)

	// Loop through array of commands
	for _, api := range apis {
		// Route to runtime API method call based on passed command value
		triggerContextAPIs(&api, logs)
	}

	return logs
}
