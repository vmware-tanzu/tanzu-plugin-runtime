package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

var _ = Describe("Test RunAPIs method", func() {

	BeforeEach(func() {
		helpers.SetupTempCfgFiles()
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
								ResponseBody: &types.Context{
									Name:   "context-one",
									Target: "kubernetes",
									GlobalOpts: &types.GlobalServer{
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
