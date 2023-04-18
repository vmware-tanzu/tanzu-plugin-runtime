// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package validators provides functions to validate the API response
package validators

import (
	"fmt"
	"reflect"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// ValidateAPIsOutput validate all the api expected output with actual output logs
func ValidateAPIsOutput(apis []*core.API, stdout string) {
	// Parse stdout to logs
	logs, err := core.ParseStdout(stdout)
	gomega.Expect(err).To(gomega.BeNil())

	for _, api := range apis {
		for _, log := range logs[api.Name] {
			switch log.APIResponse.ResponseType {
			case core.StringResponse:
				actual := fmt.Sprintf("%v", log.APIResponse.ResponseBody)
				expected := api.Output.Content
				gomega.Expect(actual).To(gomega.Equal(expected))
			case core.MapResponse:
				actual := log.APIResponse.ResponseBody
				// Convert string represented struct to map string interface
				var expected map[string]interface{}
				err := yaml.Unmarshal([]byte(api.Output.Content), &expected)
				if err != nil {
					ginkgo.By("Expected struct response but found string")
				}
				gomega.Expect(err).To(gomega.BeNil())

				// Determine the validation strategy i.e. should compare the exact objects or check whether expected properties match in the actual response
				if api.Output.ValidationStrategy == core.ValidationStrategyStrict {
					gomega.Expect(actual).To(gomega.Equal(expected))
				} else {
					gomega.Expect(ValidateMaps(actual.(map[string]interface{}), expected)).To(gomega.Equal(true))
				}
			case core.ErrorResponse:
				// Check for errors
				actual := log.APIResponse.ResponseBody
				expected := api.Output.Content
				gomega.Expect(actual).To(gomega.Equal(expected))
			}
		}
	}
}

// ValidateMaps recursive equality check on map structs
func ValidateMaps(actual, expected map[string]interface{}) bool {
	for k, v := range expected {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			ValidateMaps(actual[k].(map[string]interface{}), v.(map[string]interface{}))
		} else if !reflect.DeepEqual(actual[k], v) {
			gomega.Expect(actual[k]).To(gomega.Equal(v))
			return false
		}
	}
	return true
}
