// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func TestTriggerContextAPIs(t *testing.T) {
	_, cleanup := core.SetupTempCfgFiles()
	defer func() {
		cleanup()
	}()
	ctx := `name: context-one
target: kubernetes
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
			name:    "Trigger SetContext API",
			apiName: core.SetContextAPI,
			apis: []core.API{
				{
					Name:    core.SetContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Context:    ctx,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetContextAPI: {
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
			name:    "Trigger GetContext API",
			apiName: core.GetContextAPIName,
			apis: []core.API{
				{
					Name:    core.SetContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Context:    ctx,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetContextAPIName,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: ctx,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetContextAPIName: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: &configtypes.Context{
								Name:   "context-one",
								Target: "kubernetes",
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
			name:    "Trigger RemoveContext API",
			apiName: core.RemoveContextAPI,
			apis: []core.API{
				{
					Name:    core.SetContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Context:    ctx,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.RemoveContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.RemoveContextAPI: {
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
			name:    "Trigger DeleteContext API",
			apiName: core.DeleteContextAPI,
			apis: []core.API{
				{
					Name:    core.SetContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Context:    ctx,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.DeleteContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.DeleteContextAPI: {
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
			name:    "Trigger SetCurrentContext API",
			apiName: core.SetCurrentContextAPI,
			apis: []core.API{
				{
					Name:    core.SetContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Context:    ctx,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.SetCurrentContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetCurrentContextAPI: {
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
			name:    "Trigger GetCurrentContext API",
			apiName: core.GetCurrentContextAPI,
			apis: []core.API{
				{
					Name:    core.SetContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Context:    ctx,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.SetCurrentContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetCurrentContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Target: "kubernetes",
					},
					Output: &core.Output{
						Result:  "success",
						Content: ctx,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetCurrentContextAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: &configtypes.Context{
								Name:   "context-one",
								Target: "kubernetes",
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
			name:    "Trigger RemoveCurrentContext API",
			apiName: core.RemoveCurrentContextAPI,
			apis: []core.API{
				{
					Name:    core.SetContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Context:    ctx,
						core.SetCurrent: false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.SetCurrentContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.RemoveCurrentContextAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Target: "kubernetes",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.RemoveCurrentContextAPI: {
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
