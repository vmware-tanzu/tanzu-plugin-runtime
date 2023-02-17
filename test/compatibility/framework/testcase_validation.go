// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"
	"reflect"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

// ValidateAPIsOutput validate all the api expected output with actual output logs
func ValidateAPIsOutput(apis []*API, stdout string) {
	// Construct logs
	logs := unmarshallStdout(stdout)
	for _, api := range apis {
		for _, log := range logs[api.Name] {
			if log.APIResponse.ResponseType == StringResponse {
				actual := fmt.Sprintf("%v", log.APIResponse.ResponseBody)
				expected := api.Output.Content
				Expect(actual).To(Equal(expected))
			} else if log.APIResponse.ResponseType == MapResponse {
				actual := log.APIResponse.ResponseBody
				expected := strToMap(api.Output.Content)
				Expect(actual).To(Equal(expected))
				Expect(reflect.DeepEqual(actual, expected)).To(Equal(true))
			} else if log.APIResponse.ResponseType == ErrorResponse {
				//Check for errors
				actual := log.APIError
				expected := api.Output.Content
				Expect(actual).To(Equal(expected))
			}
		}
	}
}

// unmarshallStdout convert the string represented std out log into map structure
func unmarshallStdout(s string) map[RuntimeAPIName][]APILog {
	var logs map[RuntimeAPIName][]APILog

	err := yaml.Unmarshal([]byte(s), &logs)
	Expect(err).To(BeNil())

	return logs
}
