// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package cmd contains root and test command to trigger apis
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "runtime-test-plugin-v1_0_2",
	Short: "runtime cross-version api compatibility test plugin",
	Long:  `A test plugin with runtime lib latest used for cross-version api compatibility testing of runtime config apis`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
