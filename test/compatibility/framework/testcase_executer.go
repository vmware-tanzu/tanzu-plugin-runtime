// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"
)

// Execute the list of commands from the testcase and validate the expected output with actual output and return err if output doesn't match
func (t *TestCase) Execute() []error {
	// SetupTempCfgFiles()
	var errors []error
	// Loop through each command
	for _, cmd := range t.Commands {
		for _, api := range cmd.APIs {
			// Construct the runtime-test-plugin-x_xx command to execute
			pluginCommand, err := constructTestPluginCmd(api.Version, cmd.APIs)
			if err != nil {
				errors = append(errors, err)
			}
			// Execute the constructed runtime-test-plugin-x_xx command
			stdout, _, err := Exec(pluginCommand)
			if err != nil {
				fmt.Println("Failed to Execute :", err)
				errors = append(errors, err)
				continue
			}

			// Validate the expected API Output with actual API Output
			outputLogs := stdout.String()
			ValidateAPIsOutput(cmd.APIs, outputLogs)
		}
	}
	return errors
}
