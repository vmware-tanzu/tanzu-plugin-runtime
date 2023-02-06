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

		It("using SetContextAPIName and GetContextAPIName", func() {

			apis := []compatibilitytestingframework.API{
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
			}

			expectedLogs := map[compatibilitytestingframework.RuntimeAPIName][]compatibilitytestingframework.APILog{
				"SetContextAPIName": {
					{
						APIResponse: &compatibilitytestingframework.APIResponse{
							ResponseBody: "",
							ResponseType: compatibilitytestingframework.StringResponse,
						},
						APIError: "",
					},
				},
				"GetContextAPIName": {
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
			}

			logs := triggerAPIs(apis)

			Expect(expectedLogs[compatibilitytestingframework.SetContextAPIName]).To(Equal(logs[compatibilitytestingframework.SetContextAPIName]))
			Expect(expectedLogs[compatibilitytestingframework.GetContextAPIName]).To(Equal(logs[compatibilitytestingframework.GetContextAPIName]))
		})
	})
})
