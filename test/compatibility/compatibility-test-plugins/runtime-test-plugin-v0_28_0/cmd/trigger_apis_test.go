// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	configtypes "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Test RunAPIs method", func() {

	ginkgo.BeforeEach(func() {
		core.SetupTempCfgFiles()
	})

	ginkgo.Context("Test TriggerAPIs", func() {

		ginkgo.It("using SetContextAPIName and GetContextAPIName", func() {

			apis := []core.API{
				{
					Name:    core.SetContextAPIName,
					Version: core.Version100,
					Arguments: map[core.APIArgumentType]interface{}{
						"context": `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
						"isCurrent": false,
					},
					Output: &core.Output{
						Result:  "success",
						Content: "",
					},
				},
				{
					Name:    core.GetContextAPIName,
					Version: core.Version100,
					Arguments: map[core.APIArgumentType]interface{}{
						"contextName": "context-one",
					},
					Output: &core.Output{
						Result: "success",
						Content: `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
					},
				},
			}

			expectedLogs := map[core.RuntimeAPIName][]core.APILog{
				"SetContextAPIName": {
					{
						APIResponse: &core.APIResponse{
							ResponseBody: "",
							ResponseType: core.StringResponse,
						},
					},
				},
				"GetContextAPIName": {
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
			}

			logs := triggerAPIs(apis)

			gomega.Expect(expectedLogs[core.SetContextAPIName]).To(gomega.Equal(logs[core.SetContextAPIName]))
			gomega.Expect(expectedLogs[core.GetContextAPIName]).To(gomega.Equal(logs[core.GetContextAPIName]))
		})
	})
})
