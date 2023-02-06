// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func TestTriggerAPIs(t *testing.T) {
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
			apiName: core.SetContextAPIName,
			apis: []core.API{
				{
					Name:    core.SetContextAPIName,
					Version: core.Version0280,
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
				core.SetContextAPIName: {
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
					Name:    core.SetContextAPIName,
					Version: core.Version0280,
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
					Version: core.Version0280,
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
							ResponseBody: &configapi.Context{
								Name:   "context-one",
								Target: "kubernetes",
								GlobalOpts: &configapi.GlobalServer{
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
			apiName: core.RemoveContextAPIName,
			apis: []core.API{
				{
					Name:    core.SetContextAPIName,
					Version: core.Version0280,
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
					Name:    core.RemoveContextAPIName,
					Version: core.Version0280,
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
				core.RemoveContextAPIName: {
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
			apiName: core.DeleteContextAPIName,
			apis: []core.API{
				{
					Name:    core.SetContextAPIName,
					Version: core.Version0280,
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
					Name:    core.DeleteContextAPIName,
					Version: core.Version0280,
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
				core.DeleteContextAPIName: {
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
			apiName: core.SetCurrentContextAPIName,
			apis: []core.API{
				{
					Name:    core.SetContextAPIName,
					Version: core.Version0280,
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
					Name:    core.SetCurrentContextAPIName,
					Version: core.Version0280,
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
				core.SetCurrentContextAPIName: {
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
			apiName: core.GetCurrentContextAPIName,
			apis: []core.API{
				{
					Name:    core.SetContextAPIName,
					Version: core.Version0280,
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
					Name:    core.SetCurrentContextAPIName,
					Version: core.Version0280,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetCurrentContextAPIName,
					Version: core.Version0280,
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
				core.GetCurrentContextAPIName: {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: &configapi.Context{
								Name:   "context-one",
								Target: "kubernetes",
								GlobalOpts: &configapi.GlobalServer{
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
			apiName: core.RemoveCurrentContextAPIName,
			apis: []core.API{
				{
					Name:    core.SetContextAPIName,
					Version: core.Version0280,
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
					Name:    core.SetCurrentContextAPIName,
					Version: core.Version0280,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ContextName: "context-one",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.RemoveCurrentContextAPIName,
					Version: core.Version0280,
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
				core.RemoveCurrentContextAPIName: {
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
