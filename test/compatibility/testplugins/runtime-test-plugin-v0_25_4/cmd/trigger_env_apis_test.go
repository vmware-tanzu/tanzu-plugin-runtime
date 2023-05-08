// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func TestTriggerEnvAPIs(t *testing.T) {
	_, cleanup := core.SetupTempCfgFiles()
	defer func() {
		cleanup()
	}()

	var tests = []struct {
		name         string
		apiName      core.RuntimeAPIName
		apis         []core.API
		expectedLogs map[core.RuntimeAPIName][]core.APILog
	}{
		{
			name:    "Trigger SetEnv API",
			apiName: core.SetEnvAPI,
			apis: []core.API{
				{
					Name:    core.SetEnvAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.KeyName:   "compatibility-tests",
						core.ValueName: "default-env-val",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.SetEnvAPI: {
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
			name:    "Trigger GenEnvConfigurations API",
			apiName: core.GetEnvConfigurationsAPI,
			apis: []core.API{
				{
					Name:    core.GetEnvConfigurationsAPI,
					Version: core.VersionLatest,
					Output: &core.Output{
						Result:  "success",
						Content: `compatibility-tests: default-env-val`,
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.GetEnvConfigurationsAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.MapResponse,
							ResponseBody: map[string]string{
								"compatibility-tests": "default-env-val",
							},
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
