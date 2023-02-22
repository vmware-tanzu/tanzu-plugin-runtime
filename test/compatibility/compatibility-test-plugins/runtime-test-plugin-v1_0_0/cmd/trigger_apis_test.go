// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	compatibilitytestingcore "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Test RunAPIs method", func() {

	ginkgo.BeforeEach(func() {
		compatibilitytestingcore.SetupTempCfgFiles()
	})

	ginkgo.Context("Test TriggerAPIs", func() {

		var tests = []struct {
			apis         []compatibilitytestingcore.API
			expectedLogs map[compatibilitytestingcore.RuntimeAPIName][]compatibilitytestingcore.APILog
		}{
			{
				apis: []compatibilitytestingcore.API{
					{
						Name:    compatibilitytestingcore.SetContextAPIName,
						Version: compatibilitytestingcore.Version100,
						Arguments: map[string]interface{}{
							"context": `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
							"isCurrent": false,
						},
						Output: &compatibilitytestingcore.Output{
							Result:  "success",
							Content: "",
						},
					},
					{
						Name:    compatibilitytestingcore.GetContextAPIName,
						Version: compatibilitytestingcore.Version100,
						Arguments: map[string]interface{}{
							"contextName": "context-one",
						},
						Output: &compatibilitytestingcore.Output{
							Result: "success",
							Content: `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
						},
					},
				},

				expectedLogs: map[compatibilitytestingcore.RuntimeAPIName][]compatibilitytestingcore.APILog{
					compatibilitytestingcore.SetContextAPIName: {
						{
							APIResponse: &compatibilitytestingcore.APIResponse{
								ResponseBody: "",
								ResponseType: compatibilitytestingcore.StringResponse,
							},
							APIError: "",
						},
					},
					compatibilitytestingcore.GetContextAPIName: {
						{
							APIResponse: &compatibilitytestingcore.APIResponse{
								ResponseBody: &configtypes.Context{
									Name:   "context-one",
									Target: "kubernetes",
									GlobalOpts: &configtypes.GlobalServer{
										Endpoint: "test-endpoint",
									},
								},
								ResponseType: compatibilitytestingcore.MapResponse,
							},
							APIError: "",
						},
					},
				},
			},
		}

		ginkgo.It("using SetContextAPIName and GetContextAPIName", func() {
			for _, tt := range tests {
				actualLogs := triggerAPIs(tt.apis)

				gomega.Expect(tt.expectedLogs[compatibilitytestingcore.SetContextAPIName]).To(gomega.Equal(actualLogs[compatibilitytestingcore.SetContextAPIName]))
				gomega.Expect(tt.expectedLogs[compatibilitytestingcore.GetContextAPIName]).To(gomega.Equal(actualLogs[compatibilitytestingcore.GetContextAPIName]))

			}
		})
	})
})
