// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
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
			apiName: core.SetServerAPI,
			apis: []core.API{
				{
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
				core.SetServerAPI: {
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
			name:    "Trigger GetServerAPI",
			apiName: core.GetServerAPI,
			apis: []core.API{
				{
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
					Name:    core.GetServerAPI,
					Version: core.VersionLatest,
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
				core.GetServerAPI: {
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
			apiName: core.RemoveServerAPI,
			apis: []core.API{
				{
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
					Name:    core.RemoveServerAPI,
					Version: core.VersionLatest,
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
				core.RemoveServerAPI: {
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
			apiName: core.DeleteServerAPI,
			apis: []core.API{
				{
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
					Name:    core.DeleteServerAPI,
					Version: core.VersionLatest,
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
				core.DeleteServerAPI: {
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
			apiName: core.SetCurrentServerAPI,
			apis: []core.API{
				{
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
					Name:    core.SetCurrentServerAPI,
					Version: core.VersionLatest,
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
				core.SetCurrentServerAPI: {
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
			apiName: core.GetCurrentServerAPI,
			apis: []core.API{
				{
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
					Name:    core.SetCurrentServerAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-test-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetCurrentServerAPI,
					Version: core.VersionLatest,
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
				core.GetCurrentServerAPI: {
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
			name:    "Trigger RemoveCurrentServerAPI",
			apiName: core.RemoveCurrentServerAPI,
			apis: []core.API{
				{
					Name:    core.SetServerAPI,
					Version: core.VersionLatest,
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
					Name:    core.SetCurrentServerAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-test-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.RemoveCurrentServerAPI,
					Version: core.VersionLatest,
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
				core.RemoveCurrentServerAPI: {
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
