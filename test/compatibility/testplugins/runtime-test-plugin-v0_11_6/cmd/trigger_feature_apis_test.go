// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func TestTriggerFeatureAPIs(t *testing.T) {
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
			name:    "Trigger IsFeatureEnabled API",
			apiName: core.IsFeatureEnabledAPI,
			apis: []core.API{
				{
					Name:    core.IsFeatureEnabledAPI,
					Version: core.VersionLatest,
					Arguments: map[core.APIArgumentType]interface{}{
						core.Plugin: "test-plugin",
						core.Key:    "compatibility-tests",
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
			},

			expectedLogs: map[core.RuntimeAPIName][]core.APILog{
				core.IsFeatureEnabledAPI: {
					{
						APIResponse: &core.APIResponse{
							ResponseType: core.BooleanResponse,
							ResponseBody: false,
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
