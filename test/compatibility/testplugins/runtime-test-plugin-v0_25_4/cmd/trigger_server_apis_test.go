// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func TestTriggerServerAPIs(t *testing.T) {
	_, cleanup := core.SetupTempCfgFiles()
	defer func() {
		cleanup()
	}()
	server := `name: compatibility-test-one
type: managementcluster
globalOpts:
    endpoint: test-endpoint
`
	var tests = []struct {
		name         string
		apiName      core.RuntimeAPIName
		apis         []core.API
		expectedLogs map[core.RuntimeAPIName][]core.APILog
	}{
		{
			name:    "Trigger SetServerAPI",
			apiName: core.SetServerAPIName,
			apis: []core.API{
				{
					Name:    core.SetServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Server:     server,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetServerAPIName: {
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
			name:    "Trigger GetServerAPI",
			apiName: core.GetServerAPIName,
			apis: []core.API{
				{
					Name:    core.SetServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Server:     server,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-test-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: server,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetServerAPIName: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: &configtypes.Server{
								Name: "compatibility-test-one",
								Type: configtypes.ManagementClusterServerType,
								GlobalOpts: &configtypes.GlobalServer{
									Endpoint: "test-endpoint",
								},
							},
							ResponseType: core.MapResponse,
						},
					},
				},
			},
		},

		{
			name:    "Trigger RemoveServerAPI",
			apiName: core.RemoveServerAPIName,
			apis: []core.API{
				{
					Name:    core.SetServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Server:     server,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.RemoveServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-test-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.RemoveServerAPIName: {
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
			name:    "Trigger DeleteServerAPI",
			apiName: core.DeleteServerAPIName,
			apis: []core.API{
				{
					Name:    core.SetServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Server:     server,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.DeleteServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-test-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.DeleteServerAPIName: {
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
			name:    "Trigger SetCurrentServerAPI",
			apiName: core.SetCurrentServerAPIName,
			apis: []core.API{
				{
					Name:    core.SetServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Server:     server,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.SetCurrentServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-test-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetCurrentServerAPIName: {
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
			name:    "Trigger GetCurrentServerAPI",
			apiName: core.GetCurrentServerAPIName,
			apis: []core.API{
				{
					Name:    core.SetServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Server:     server,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.SetCurrentServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-test-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetCurrentServerAPIName,
					Version: core.Version0254,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Target: "kubernetes",
					},
					Output: &core.Output{
						Result:  "success",
						Content: server,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetCurrentServerAPIName: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: &configtypes.Server{
								Name: "compatibility-test-one",
								Type: configtypes.ManagementClusterServerType,
								GlobalOpts: &configtypes.GlobalServer{
									Endpoint: "test-endpoint",
								},
							},
							ResponseType: core.MapResponse,
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
