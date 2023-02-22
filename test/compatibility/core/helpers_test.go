// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
    setCurrent: false
  output:
    result: success
    content: ""
`

	temp, err := os.CreateTemp("", "api")
	assert.Nil(t, err)

	err = os.WriteFile(temp.Name(), []byte(apiYaml), 0644)
	assert.Nil(t, err)
	apis, err := ParseRuntimeAPIsFromFile(temp.Name())
	assert.Nil(t, err)

	expectedAPIs := []API{
		{
			Name:    SetContextAPIName,
			Version: Version100,
			Arguments: map[APIArgumentType]interface{}{
				Context: `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
				SetCurrent: false,
			},
			Output: &Output{
				Result:  Success,
				Content: "",
			},
		},
	}
	assert.Equal(t, expectedAPIs, apis)
}
