// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	configtypes "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	compatibilitytestingcore "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = Describe("Test RunAPIs method", func() {

	BeforeEach(func() {
		compatibilitytestingcore.SetupTempCfgFiles()
	})

	Context("Test TriggerAPIs", func() {

		It("using SetContextAPIName and GetContextAPIName", func() {

			apis := []compatibilitytestingcore.API{
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
			}

			expectedLogs := map[compatibilitytestingcore.RuntimeAPIName][]compatibilitytestingcore.APILog{
				"SetContextAPIName": {
					{
						APIResponse: &compatibilitytestingcore.APIResponse{
							ResponseBody: "",
							ResponseType: compatibilitytestingcore.StringResponse,
						},
						APIError: "",
					},
				},
				"GetContextAPIName": {
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
			}

			logs := triggerAPIs(apis)

			Expect(expectedLogs[compatibilitytestingcore.SetContextAPIName]).To(Equal(logs[compatibilitytestingcore.SetContextAPIName]))
			Expect(expectedLogs[compatibilitytestingcore.GetContextAPIName]).To(Equal(logs[compatibilitytestingcore.GetContextAPIName]))
		})
	})
})
