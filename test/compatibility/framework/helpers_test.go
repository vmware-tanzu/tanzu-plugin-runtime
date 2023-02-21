// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
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
	assert.Nil(t, err)
	apis, err := core.ParseRuntimeAPIsFromFile(temp.Name())
	assert.Nil(t, err)

	expectedAPIs := []core.API{
		{
			Name:    core.SetContextAPIName,
			Version: core.Version100,
			Arguments: map[string]interface{}{
				"context": `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
				"isCurrent": false,
			},
			Output: &core.Output{
				Result:  core.Success,
				Content: "",
			},
		},
	}
	assert.Equal(t, expectedAPIs, apis)
}
