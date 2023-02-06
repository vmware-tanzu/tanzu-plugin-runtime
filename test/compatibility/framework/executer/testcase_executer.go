// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package executer provides function to execute the test case commands
package executer

import (
	"fmt"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/validators"
)

// Execute the list of commands from the testcase and validate the expected output with actual output and return err if output doesn't match
func Execute(t *core.TestCase) {
	// Loop through each command
	for _, cmd := range t.Commands {
		for _, api := range cmd.APIs {
			ginkgo.By(fmt.Sprintf("Runnning %v - %v", api.Name, api.Version))
			// Construct the runtime-test-plugin-x_xx_xx command to execute
			pluginCommand, err := core.ConstructTestPluginCmd(api.Version, cmd.APIs)
			gomega.Expect(err).To(gomega.BeNil())

			// Execute the constructed runtime-test-plugin-x_xx_xx command
			stdout, stderr, err := core.Exec(pluginCommand)

			if stderr != nil && stderr.String() != "" {
				gomega.Expect(stderr.String()).To(gomega.BeNil())
			}
			gomega.Expect(err).To(gomega.BeNil())

			// Validate the expected API Output with actual API Output
			outputLogs := stdout.String()
			validators.ValidateAPIsOutput(cmd.APIs, outputLogs)
			ginkgo.By(fmt.Sprintf("Successful Running %v - %v", api.Name, api.Version))
		}
	}
}
