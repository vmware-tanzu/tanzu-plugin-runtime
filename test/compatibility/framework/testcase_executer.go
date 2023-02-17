// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	. "github.com/onsi/gomega"
)

// Execute the list of commands from the testcase and validate the expected output with actual output and return err if output doesn't match
func (t *TestCase) Execute() {
	// Mock the config files CFG, CFG_NG and META
	_, cleanUp := SetupTempCfgFiles()

	// Clean up the mock config files after execution is complete
	defer func() {
		cleanUp()
	}()

	// Loop through each command
	for _, cmd := range t.Commands {
		for _, api := range cmd.APIs {

			// Construct the runtime-test-plugin-x_xx command to execute
			pluginCommand, err := constructTestPluginCmd(api.Version, cmd.APIs)
			Expect(err).To(BeNil())

			// Execute the constructed runtime-test-plugin-x_xx command
			stdout, stderr, err := Exec(pluginCommand)

			if stderr != nil && len(stderr.String()) != 0 {
				Expect(stderr.String()).To(BeNil())
			}
			Expect(err).To(BeNil())

			// Validate the expected API Output with actual API Output
			outputLogs := stdout.String()
			ValidateAPIsOutput(cmd.APIs, outputLogs)
		}
	}

}
