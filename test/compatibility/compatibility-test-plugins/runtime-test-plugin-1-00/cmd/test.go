// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	compatibilitytestingtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers"
	"gopkg.in/yaml.v3"
)

// testCmd represents the test command
var (
	file    string
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "A test command that parse the file and trigger the runtime apis",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse the file into array of apis struct
			apis, err := helpers.GetTestData(file)
			if err != nil {
				fmt.Println(err)
			}

			// mock config files
			_, cleanUp := helpers.SetupTempCfgFiles()
			defer func() {
				cleanUp()
			}()
			runAPIs(apis)
		},
	}
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&file, "file", "f", "", "test file path")
}

// runAPIs loop through the apis and trigger the runtime api methods and print logs to stdout
func runAPIs(apis []compatibilitytestingtypes.API) {
	// Trigger apis and return logs to be printed
	logs := triggerAPIs(apis)

	// Log the output to stdout
	bytes, err := yaml.Marshal(logs)
	if err != nil {
		fmt.Println("runAPIs", err)
	}

	fmt.Println(string(bytes))
}
