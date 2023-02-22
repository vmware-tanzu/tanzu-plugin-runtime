// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Test RunAPIs method", func() {

	ginkgo.BeforeEach(func() {
		core.SetupTempCfgFiles()
	})

	ginkgo.Context("Test TriggerAPIs", func() {

		var tests = []struct {
			apis         []core.API
			expectedLogs map[core.RuntimeAPIName][]core.APILog
		}{
			{
				apis: []core.API{
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
		}

		ginkgo.It("using SetContextAPIName and GetContextAPIName", func() {
			for _, tt := range tests {
				actualLogs := triggerAPIs(tt.apis)

				gomega.Expect(tt.expectedLogs[core.SetContextAPIName]).To(gomega.Equal(actualLogs[core.SetContextAPIName]))
				gomega.Expect(tt.expectedLogs[core.GetContextAPIName]).To(gomega.Equal(actualLogs[core.GetContextAPIName]))

			}
		})
	})
})
