// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

const sourceOne string = "compatibility-tests-source-one"
const sourceImage string = "compatibility-tests-source-image"

func TestTriggerCLIDiscoverySourceAPIs(t *testing.T) {
	_, cleanup := core.SetupTempCfgFiles()
	defer func() {
		cleanup()
	}()
	source := `
oci:
    name: compatibility-tests-source-one
    image: compatibility-tests-source-image
`
	var tests = []struct {
		name         string
		apiName      core.RuntimeAPIName
		apis         []core.API
		expectedLogs map[core.RuntimeAPIName][]core.APILog
	}{
		{
			name:    "Trigger SetCLIDiscoverySourceAPI",
			apiName: core.SetCLIDiscoverySourceAPI,
			apis: []core.API{
				{
					Name:    core.SetCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.DiscoverySource: source,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.SetCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.DiscoverySource: source,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetCLIDiscoverySourceAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
				},
			},
		},

		{
			name:    "Trigger GetCLIDiscoverySourceAPI",
			apiName: core.GetCLIDiscoverySourceAPI,
			apis: []core.API{
				{
					Name:    core.SetCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.DiscoverySource: source,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Name: sourceOne,
					},
					Output: &core.Output{
						Result:  "success",
						Content: source,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetCLIDiscoverySourceAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: &configtypes.PluginDiscovery{
								OCI: &configtypes.OCIDiscovery{
									Name:  sourceOne,
									Image: sourceImage,
								},
							},
							ResponseType: core.MapResponse,
						},
					},
				},
			},
		},

		{
			name:    "Trigger DeleteCLIDiscoverySourceAPI",
			apiName: core.DeleteCLIDiscoverySourceAPI,
			apis: []core.API{
				{
					Name:    core.SetCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.DiscoverySource: source,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.DeleteCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Name: sourceOne,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.DeleteCLIDiscoverySourceAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
				},
			},
		},

		{
			name:    "Trigger DeleteCLIDiscoverySourceAPI",
			apiName: core.DeleteCLIDiscoverySourceAPI,
			apis: []core.API{
				{
					Name:    core.SetCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.DiscoverySource: source,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.DeleteCLIDiscoverySourceAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Name: sourceOne,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.DeleteCLIDiscoverySourceAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualLogs := triggerAPIs(tt.apis)
			assert.Equal(t, tt.expectedLogs[tt.apiName], actualLogs[tt.apiName])
		})
	}
}
