// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"

	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// ValidateAPIsOutput validate all the api expected output with actual output logs
func ValidateAPIsOutput(apis []*core.API, stdout string) {
	// Parse stdout to logs
	logs, err := core.ParseStdout(stdout)
	gomega.Expect(err).To(gomega.BeNil())

	for _, api := range apis {
		for _, log := range logs[api.Name] {
			if log.APIResponse.ResponseType == core.StringResponse {
				actual := fmt.Sprintf("%v", log.APIResponse.ResponseBody)
				expected := api.Output.Content
				gomega.Expect(actual).To(gomega.Equal(expected))
			} else if log.APIResponse.ResponseType == core.MapResponse {
				actual := log.APIResponse.ResponseBody
				expected := core.StrToMap(api.Output.Content)

				if api.Output.ValidationStrategy == core.ValidationStrategyStrict {
					gomega.Expect(actual).To(gomega.Equal(expected))
				} else {
					gomega.Expect(core.ValidateMaps(actual.(map[string]interface{}), expected)).To(gomega.Equal(true))
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