// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
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
				[]core.API{
					{
						Name:    core.SetContextAPIName,
						Version: core.Version100,
						Arguments: map[core.APIArgumentType]interface{}{
							core.Context: `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
							core.SetCurrent: false,
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
							core.ContextName: "context-one",
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
				map[core.RuntimeAPIName][]core.APILog{
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
		}

		ginkgo.It("using SetContext and GetContext APIs", func() {
			for _, tt := range tests {
				actualLogs := triggerAPIs(tt.apis)

				gomega.Expect(tt.expectedLogs[core.SetContextAPIName]).To(gomega.Equal(actualLogs[core.SetContextAPIName]))
				gomega.Expect(tt.expectedLogs[core.GetContextAPIName]).To(gomega.Equal(actualLogs[core.GetContextAPIName]))

			}
		})
	})
})
