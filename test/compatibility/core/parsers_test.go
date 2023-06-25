// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRuntimeAPIsFromFile(t *testing.T) {
	tests := []struct {
		apiYaml      string
		expectedAPIs []API
	}{

		{
			`- name: SetContext
  version: latest
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
`,
			[]API{
				{
					Name:    SetContextAPI,
					Version: VersionLatest,
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
			},
		},
	}

	for _, tt := range tests {
		temp, err := os.CreateTemp("", "api")
		assert.Nil(t, err)

		err = os.WriteFile(temp.Name(), []byte(tt.apiYaml), 0644)
		assert.Nil(t, err)
		apis, err := ParseRuntimeAPIsFromFile(temp.Name())
		assert.Nil(t, err)

		assert.Equal(t, tt.expectedAPIs, apis)
	}
}

func TestParseStr(t *testing.T) {
	tests := []struct {
		input  interface{}
		output string
		err    string
	}{
		{
			"compatibility-one",
			"compatibility-one",
			"",
		},
	}

	for _, tt := range tests {
		out, err := ParseStr(tt.input)
		if tt.err == "" {
			assert.Equal(t, tt.output, out)
		} else {
			assert.Equal(t, tt.err, err.Error())
		}
	}
}

func TestParseStdout(t *testing.T) {
	tests := []struct {
		input  string
		output map[RuntimeAPIName][]APILog
		err    string
	}{
		{
			"SetContext:\n    - apiResponse:\n        responseType: str\n        responseBody: \"\"\n\n",
			map[RuntimeAPIName][]APILog{
				"SetContext": {
					{
						APIResponse: &APIResponse{
							ResponseBody: "",
							ResponseType: "str",
						},
					},
				},
			},
			"",
		},
	}

	for _, tt := range tests {
		out, err := ParseStdout(tt.input)
		if tt.err == "" {
			assert.Equal(t, tt.output, out)
		} else {
			assert.Equal(t, tt.err, err.Error())
		}
	}
}
