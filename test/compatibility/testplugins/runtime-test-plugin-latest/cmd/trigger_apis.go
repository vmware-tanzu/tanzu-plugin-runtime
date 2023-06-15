// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// Route to runtime API method call based on passed command value
var apiHandlers = map[core.RuntimeAPIName]func(*core.API) *core.APIResponse{
	// Context APIs
	core.SetContextAPI:           triggerSetContextAPI,
	core.GetContextAPIName:       triggerGetContextAPI,
	core.RemoveContextAPI:        triggerDeleteContextAPI,
	core.DeleteContextAPI:        triggerDeleteContextAPI,
	core.SetCurrentContextAPI:    triggerSetCurrentContextAPI,
	core.GetCurrentContextAPI:    triggerGetCurrentContextAPI,
	core.RemoveCurrentContextAPI: triggerRemoveCurrentContextAPI,

	// Server APIs
	core.SetServerAPI:           triggerSetServerAPI,
	core.AddServerAPI:           triggerSetServerAPI,
	core.PutServerAPI:           triggerSetServerAPI,
	core.GetServerAPI:           triggerGetServerAPI,
	core.RemoveServerAPI:        triggerRemoveServerAPI,
	core.DeleteServerAPI:        triggerRemoveServerAPI,
	core.SetCurrentServerAPI:    triggerSetCurrentServerAPI,
	core.GetCurrentServerAPI:    triggerGetCurrentServerAPI,
	core.RemoveCurrentServerAPI: triggerRemoveCurrentServerAPI,

	// Feature APIs
	core.SetFeatureAPI:       triggerSetFeatureAPI,
	core.IsFeatureEnabledAPI: triggerIsFeatureEnabledAPI,
	core.DeleteFeatureAPI:    triggerDeleteFeatureAPI,

	// Env APIs
	core.SetEnvAPI:               triggerSetEnvAPI,
	core.GetEnvAPI:               triggerGetEnvAPI,
	core.GetEnvConfigurationsAPI: triggerGetEnvConfigurationsAPI,
	core.DeleteEnvAPI:            triggerDeleteEnvAPI,

	// CLI Discovery Source APIs
	core.SetCLIDiscoverySourceAPI:    triggerSetCLIDiscoverySourceAPI,
	core.GetCLIDiscoverySourceAPI:    triggerGetCLIDiscoverySourceAPI,
	core.DeleteCLIDiscoverySourceAPI: triggerDeleteCLIDiscoverySourceAPI,

	// Metadata APIs
	core.SetConfigMetadataSettingAPI:        triggerSetConfigMetadataSettingAPI,
	core.SetConfigMetadataPatchStrategyAPI:  triggerSetConfigMetadataPatchStrategyAPI,
	core.DeleteConfigMetadataSettingAPI:     triggerDeleteConfigMetadataSettingAPI,
	core.GetMetadataAPI:                     triggerGetMetadataAPI,
	core.GetConfigMetadataAPI:               triggerGetConfigMetadataAPI,
	core.GetConfigMetadataPatchStrategyAPI:  triggerGetConfigMetadataPatchStrategyAPI,
	core.GetConfigMetadataSettingsAPI:       triggerGetConfigMetadataSettingsAPI,
	core.GetConfigMetadataSettingAPI:        triggerGetConfigMetadataSettingAPI,
	core.IsConfigMetadataSettingsEnabledAPI: triggerIsConfigMetadataSettingsEnabledAPI,
	core.UseUnifiedConfigAPI:                triggerUseUnifiedConfigAPI,

	// Global APIs
	core.GetClientConfigAPI:   triggerGetClientConfigAPI,
	core.StoreClientConfigAPI: triggerStoreClientConfigAPI,
}

// triggerAPIs trigger runtime apis and construct logs
func triggerAPIs(apis []core.API) map[core.RuntimeAPIName][]core.APILog {
	// Variable used to store all the logging related to runtime api responses
	logs := make(map[core.RuntimeAPIName][]core.APILog)

	// Loop through array of commands
	for index := range apis {
		api := &apis[index]
		handler, ok := apiHandlers[api.Name]
		if !ok {
			log := core.APILog{
				APIResponse: &core.APIResponse{
					ResponseType: core.ErrorResponse,
					ResponseBody: fmt.Errorf("command %v not found", api.Name),
				},
			}
			logs[api.Name] = append(logs[api.Name], log)
			continue
		}

		// Trigger the API handler
		apiResponse := handler(api)

		// Construct the logs
		log := core.APILog{
			APIResponse: apiResponse,
		}
		logs[api.Name] = append(logs[api.Name], log)
	}

	return logs
}
