// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestUsageFunc(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	cmd := &cobra.Command{
		Use:   "Sub1",
		Short: "Sub1 description",
	}
	err = UsageFunc(cmd)
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := <-c
	assert.Contains(t, string(got), "Usage:")
}

func usageTestPlugin(t *testing.T, target types.Target) *Plugin {
	var pluginsCmd = &cobra.Command{
		Use:   "plugin",
		Short: "Plugin tests",
	}

	var fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch the plugin tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("fetch")
			return nil
		},
	}

	var pushCmd = &cobra.Command{
		Use:     "push SOMESTUFF",
		Aliases: []string{"psh"},
		Short:   "Push the plugin tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("push")
			return nil
		},
	}

	var pushMoreCmd = &cobra.Command{
		Use:     "more",
		Short:   "Push more",
		Aliases: []string{"mo"},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("push more")
			return nil
		},
	}

	pushCmd.AddCommand(pushMoreCmd)

	var descriptor = PluginDescriptor{
		Name:        "testNotUserVisible",
		Target:      target,
		Aliases:     []string{"t"},
		Description: "Test the CLI",
		Group:       AdminCmdGroup,
		Version:     "v1.1.0",
		BuildSHA:    "1234567",
		CommandMap: []CommandMapEntry{
			{
				DestinationCommandPath: "test",
			},
			{
				DestinationCommandPath: "fetch",
				SourceCommandPath:      "fetch",
			},
		},
	}

	var local string
	var url string

	fetchCmd.Flags().StringVarP(&local, "local", "l", "", "path to local repository")
	_ = fetchCmd.MarkFlagRequired("local")

	fetchCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "url to remote repository")
	_ = fetchCmd.MarkFlagRequired("url")

	fetchCmd.Example = "sample example usage of the fetch command"

	pushCmd.Example = "sample example usage of the push command"

	p, err := NewPlugin(&descriptor)
	assert.Nil(t, err)

	var env string
	p.Cmd.PersistentFlags().StringVarP(&env, "env", "e", "", "env to test")

	p.Cmd.Example = "sample example usage of the test command"
	p.AddCommands(
		fetchCmd,
		pushCmd,
		pluginsCmd,
	)

	return p
}

func TestGlobalTestPluginCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with Global target
	p := usageTestPlugin(t, types.TargetGlobal)

	// Set the arguments as if the user typed them in the command line
	p.Cmd.SetArgs([]string{"--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	// note: reference to the unmapped name, as in
	//
	// '-h, --help         help for testNotUserVisible'
	//
	// is a known bug in cobra 1.8.0 that should be fixed in the next patch or
	// minor release
	expected := `Test the CLI

Usage:
  tanzu test [command]

Aliases:
  test, t

Examples:
  sample example usage of the test command

Available Commands:
  fetch         Fetch the plugin tests
  push          Push the plugin tests

Flags:
  -e, --env string   env to test
  -h, --help         help for testNotUserVisible

Additional help topics:
  test plugin        Plugin tests

Use "tanzu test [command] --help" for more information about a command.
`
	assert.Equal(t, expected, got)
}

func TestKubernetesTestPluginCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with Kubernetes target
	p := usageTestPlugin(t, types.TargetK8s)

	// Set the arguments as if the user typed them in the command line
	p.Cmd.SetArgs([]string{"--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Test the CLI

Usage:
  tanzu test [command]
  tanzu kubernetes test [command]

Aliases:
  test, t

Examples:
  sample example usage of the test command

Available Commands:
  fetch         Fetch the plugin tests
  push          Push the plugin tests

Flags:
  -e, --env string   env to test
  -h, --help         help for testNotUserVisible

Additional help topics:
  test plugin        Plugin tests

Use "tanzu test [command] --help" for more information about a command.
Use "tanzu kubernetes test [command] --help" for more information about a command.
`
	assert.Equal(t, expected, got)
}

func TestMissionControlTestPluginCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with MissionControl target
	p := usageTestPlugin(t, types.TargetTMC)

	// Set the arguments as if the user typed them in the command line
	p.Cmd.SetArgs([]string{"--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Test the CLI

Usage:
  tanzu mission-control test [command]

Aliases:
  test, t

Examples:
  sample example usage of the test command

Available Commands:
  fetch         Fetch the plugin tests
  push          Push the plugin tests

Flags:
  -e, --env string   env to test
  -h, --help         help for testNotUserVisible

Additional help topics:
  test plugin        Plugin tests

Use "tanzu mission-control test [command] --help" for more information about a command.
`
	assert.Equal(t, expected, got)
}

func TestGlobalTestPluginFetchCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with Global target
	p := usageTestPlugin(t, types.TargetGlobal)

	// Set the arguments as if the user typed them in the command line
	p.Cmd.SetArgs([]string{"fetch", "--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)
	expected := `Fetch the plugin tests

Usage:
  tanzu test fetch [flags]

Examples:
  sample example usage of the fetch command

Flags:
  -h, --help           help for fetch
  -l, --local string   path to local repository
  -u, --url string     url to remote repository

Global Flags:
  -e, --env string   env to test
`
	assert.Equal(t, expected, got)
}

func TestGlobalTestFetchCommandHelpTextWithInvocationContext(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr

	os.Setenv("TANZU_CLI_COMMAND_MAPPED_FROM", "fetch")
	os.Setenv("TANZU_CLI_INVOKED_COMMAND", "fetch")
	os.Setenv("TANZU_CLI_INVOKED_GROUP", "")
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		os.Unsetenv("TANZU_CLI_COMMAND_MAPPED_FROM")
		os.Unsetenv("TANZU_CLI_INVOKED_COMMAND")
		os.Unsetenv("TANZU_CLI_INVOKED_GROUP")
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with Global target
	p := usageTestPlugin(t, types.TargetGlobal)

	p.Cmd.SetArgs([]string{"fetch", "--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Fetch the plugin tests

Usage:
  tanzu fetch

Examples:
  sample example usage of the fetch command

Flags:
  -h, --help           help for fetch
  -l, --local string   path to local repository
  -u, --url string     url to remote repository

Global Flags:
  -e, --env string   env to test
`

	assert.Equal(t, expected, got)
}

func TestKubernetesTestPluginFetchCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with Kubernetes target
	p := usageTestPlugin(t, types.TargetK8s)

	// Set the arguments as if the user typed them in the command line
	p.Cmd.SetArgs([]string{"fetch", "--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Fetch the plugin tests

Usage:
  tanzu test fetch [flags]
  tanzu kubernetes test fetch [flags]

Examples:
  sample example usage of the fetch command

Flags:
  -h, --help           help for fetch
  -l, --local string   path to local repository
  -u, --url string     url to remote repository

Global Flags:
  -e, --env string   env to test
`
	assert.Equal(t, expected, got)
}

func TestMissionControlTestPluginFetchCommandHelpText(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with MissionControl target
	p := usageTestPlugin(t, types.TargetTMC)

	// Set the arguments as if the user typed them in the command line
	p.Cmd.SetArgs([]string{"fetch", "--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Fetch the plugin tests

Usage:
  tanzu mission-control test fetch [flags]

Examples:
  sample example usage of the fetch command

Flags:
  -h, --help           help for fetch
  -l, --local string   path to local repository
  -u, --url string     url to remote repository

Global Flags:
  -e, --env string   env to test
`
	assert.Equal(t, expected, got)
}

func TestCommandMappedCommandWithInvocationContext(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr

	os.Setenv("TANZU_CLI_COMMAND_MAPPED_FROM", "push")
	os.Setenv("TANZU_CLI_INVOKED_COMMAND", "pu")
	os.Setenv("TANZU_CLI_INVOKED_GROUP", "")
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		os.Unsetenv("TANZU_CLI_COMMAND_MAPPED_FROM")
		os.Unsetenv("TANZU_CLI_INVOKED_COMMAND")
		os.Unsetenv("TANZU_CLI_INVOKED_GROUP")
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with Global target
	p := usageTestPlugin(t, types.TargetGlobal)

	p.Cmd.SetArgs([]string{"push", "--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Push the plugin tests

Usage:
  tanzu pu SOMESTUFF

  tanzu pu [command]

Aliases:
  pu, psh

Examples:
  sample example usage of the push command

Available Commands:
  more        Push more

Flags:
  -h, --help   help for push

Global Flags:
  -e, --env string   env to test

Use "tanzu pu [command] --help" for more information about a command.
`

	assert.Equal(t, expected, got)
}

func TestCommandMappedCommandSubCommandWithInvocationContext(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr

	os.Setenv("TANZU_CLI_COMMAND_MAPPED_FROM", "push")
	os.Setenv("TANZU_CLI_INVOKED_COMMAND", "pu")
	os.Setenv("TANZU_CLI_INVOKED_GROUP", "")
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		os.Unsetenv("TANZU_CLI_COMMAND_MAPPED_FROM")
		os.Unsetenv("TANZU_CLI_INVOKED_COMMAND")
		os.Unsetenv("TANZU_CLI_INVOKED_GROUP")
	}()
	os.Stdout = w
	os.Stderr = w

	// Prepare the root command with Global target
	p := usageTestPlugin(t, types.TargetGlobal)

	p.Cmd.SetArgs([]string{"push", "more", "--help"})

	// Execute the command which will trigger the help
	err = p.Execute()
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	got := string(<-c)

	expected := `Push more

Usage:
  tanzu pu more

Aliases:
  more, mo

Flags:
  -h, --help   help for more

Global Flags:
  -e, --env string   env to test
`

	assert.Equal(t, expected, got)
}
