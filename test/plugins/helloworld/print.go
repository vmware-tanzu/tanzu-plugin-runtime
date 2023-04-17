// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PrintCmd is for printing the "hello world"
var PrintCmd = &cobra.Command{
	Use:   "print",
	Short: "prints hello-world",
	Args:  cobra.ExactArgs(0),
	Example: `
	# Prints the "Hello world" from the test plugin
	tanzu helloworld-test print`,
	RunE: printHelloWorld,
}

func printHelloWorld(_ *cobra.Command, _ []string) error {
	fmt.Printf("Hello world from test plugin")
	return nil
}
