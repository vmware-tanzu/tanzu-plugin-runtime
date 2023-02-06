package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

func TestGetTestData(t *testing.T) {

	apiYaml := `- name: SetContextAPIName
  version: v1.0.0
  arguments:
    context: |
        name: context-one
        target: kubernetes
        globalOpts:
          endpoint: test-endpoint
    isCurrent: false
  output:
    result: success
    content: ""
`

	temp, err := os.CreateTemp("", "api")
	assert.Nil(t, err)

	err = os.WriteFile(temp.Name(), []byte(apiYaml), 0644)
	apis := GetTestData(temp.Name())

	expectedAPIs := []compatibilitytestingframework.API{
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
	}
	assert.Equal(t, expectedAPIs, apis)
}
