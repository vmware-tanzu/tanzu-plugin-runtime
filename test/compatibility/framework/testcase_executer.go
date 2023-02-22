// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package framework Package provides all the helper methods to write cross-version API compatibility tests
package framework

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// Execute the list of commands from the testcase and validate the expected output with actual output and return err if output doesn't match
func Execute(t *core.TestCase) {
	// Loop through each command
	for _, cmd := range t.Commands {
		for _, api := range cmd.APIs {
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
			ValidateAPIsOutput(cmd.APIs, outputLogs)
		}
	}
}
