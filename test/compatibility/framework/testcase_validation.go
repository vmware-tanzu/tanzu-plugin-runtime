// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// ValidateAPIsOutput validate all the api expected output with actual output logs
func ValidateAPIsOutput(apis []*core.API, stdout string) {
	// Construct logs
	logs := unmarshallStdout(stdout)
	for _, api := range apis {
		for _, log := range logs[api.Name] {
			if log.APIResponse.ResponseType == core.StringResponse {
				actual := fmt.Sprintf("%v", log.APIResponse.ResponseBody)
				expected := api.Output.Content
				gomega.Expect(actual).To(gomega.Equal(expected))
			} else if log.APIResponse.ResponseType == core.MapResponse {
				actual := log.APIResponse.ResponseBody
				expected := StrToMap(api.Output.Content)

				if api.Output.ValidationStrategy == core.ValidationStrategyStrict {
					gomega.Expect(actual).To(gomega.Equal(expected))
				} else {
					gomega.Expect(validateMaps(actual.(map[string]interface{}), expected)).To(gomega.Equal(true))
				}
			} else if log.APIResponse.ResponseType == core.ErrorResponse {
				// Check for errors
				actual := log.APIResponse.ResponseBody
				expected := api.Output.Content
				gomega.Expect(actual).To(gomega.Equal(expected))
			}
		}
	}
}

// unmarshallStdout convert the string represented std out log into map structure
func unmarshallStdout(s string) map[core.RuntimeAPIName][]core.APILog {
	var logs map[core.RuntimeAPIName][]core.APILog

	err := yaml.Unmarshal([]byte(s), &logs)
	gomega.Expect(err).To(gomega.BeNil())

	return logs
}

func validateMaps(actual, expected map[string]interface{}) bool {
	for k, v := range expected {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			validateMaps(actual[k].(map[string]interface{}), v.(map[string]interface{}))
		} else if !reflect.DeepEqual(actual[k], v) {
			return false
		}
	}
	return true
}
