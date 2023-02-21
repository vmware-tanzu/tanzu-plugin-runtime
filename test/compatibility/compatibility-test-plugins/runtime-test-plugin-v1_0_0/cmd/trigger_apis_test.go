// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = Describe("Test RunAPIs method", func() {

	BeforeEach(func() {
		compatibilitytestingframework.SetupTempCfgFiles()
	})

	Context("Test TriggerAPIs", func() {

		var tests = []struct {
			apis         []compatibilitytestingframework.API
			expectedLogs map[compatibilitytestingframework.RuntimeAPIName][]compatibilitytestingframework.APILog
		}{
			{
				apis: []compatibilitytestingframework.API{
					{
						Name:    compatibilitytestingframework.SetContextAPIName,
						Version: compatibilitytestingframework.Version100,
						Arguments: map[string]interface{}{
							"context": `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
							"isCurrent": false,
						},
						Output: &compatibilitytestingframework.Output{
							Result:  "success",
							Content: "",
						},
					},
					{
						Name:    compatibilitytestingframework.GetContextAPIName,
						Version: compatibilitytestingframework.Version100,
						Arguments: map[string]interface{}{
							"contextName": "context-one",
						},
						Output: &compatibilitytestingframework.Output{
							Result: "success",
							Content: `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
						},
					},
				},

				expectedLogs: map[compatibilitytestingframework.RuntimeAPIName][]compatibilitytestingframework.APILog{
					compatibilitytestingframework.SetContextAPIName: {
						{
							APIResponse: &compatibilitytestingframework.APIResponse{
								ResponseBody: "",
								ResponseType: compatibilitytestingframework.StringResponse,
							},
							APIError: "",
						},
					},
					compatibilitytestingframework.GetContextAPIName: {
						{
							APIResponse: &compatibilitytestingframework.APIResponse{
								ResponseBody: &configtypes.Context{
									Name:   "context-one",
									Target: "kubernetes",
									GlobalOpts: &configtypes.GlobalServer{
										Endpoint: "test-endpoint",
									},
								},
								ResponseType: compatibilitytestingframework.MapResponse,
							},
							APIError: "",
						},
					},
				},
			},
		}

		It("using SetContextAPIName and GetContextAPIName", func() {
			for _, tt := range tests {
				actualLogs := triggerAPIs(tt.apis)

				Expect(tt.expectedLogs[compatibilitytestingframework.SetContextAPIName]).To(Equal(actualLogs[compatibilitytestingframework.SetContextAPIName]))
				Expect(tt.expectedLogs[compatibilitytestingframework.GetContextAPIName]).To(Equal(actualLogs[compatibilitytestingframework.GetContextAPIName]))

			}
		})
	})
})
