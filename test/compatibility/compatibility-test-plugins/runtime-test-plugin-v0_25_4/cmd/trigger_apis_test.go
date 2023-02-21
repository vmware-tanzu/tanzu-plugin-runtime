package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	configapi "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = Describe("Test RunAPIs method", func() {

	BeforeEach(func() {
		compatibilitytestingframework.SetupTempCfgFiles()
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
						Result:  compatibilitytestingframework.Success,
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
						Result: compatibilitytestingframework.Success,
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
							ResponseBody: &configapi.Context{
								Name: "context-one",
								Type: configapi.CtxTypeK8s,
								GlobalOpts: &configapi.GlobalServer{
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
