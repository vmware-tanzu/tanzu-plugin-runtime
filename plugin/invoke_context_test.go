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

func SetupTestPlugin(t *testing.T) *Plugin {
	var topCmd *cobra.Command

	// build long Command chain to attach to the root command of the plugin
	for i := 7; i >= 0; i-- {
		cmdName := fmt.Sprintf("c%d", i)
		cmd := &cobra.Command{
			Use: cmdName,
		}
		if topCmd == nil {
			topCmd = cmd
		} else {
			cmd.AddCommand(topCmd)
			topCmd = cmd
		}
	}

	var descriptor = PluginDescriptor{
		Name:        "test",
		Target:      types.TargetGlobal,
		Aliases:     []string{"t"},
		Description: "Test CLI invocation context",
		Group:       AdminCmdGroup,
		Version:     "v1.1.0",
		BuildSHA:    "1234567",
	}

	p, err := NewPlugin(&descriptor)
	assert.Nil(t, err)

	p.AddCommands(
		topCmd,
	)

	return p
}

func subCommandAtLevel(cmd *cobra.Command, level int) *cobra.Command {
	curr := cmd
	for l := 0; l <= level; l++ {
		wantCommandName := fmt.Sprintf("c%d", l)
		var found bool
		for _, cmd := range curr.Commands() {
			cname := cmd.Name()
			//if cmd.Name() == wantCommandName {
			if cname == wantCommandName {
				curr = cmd
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return curr
}

func TestInvocationContext(t *testing.T) {
	defer func() {
		os.Unsetenv("TANZU_CLI_COMMAND_MAPPED_FROM")
		os.Unsetenv("TANZU_CLI_INVOKED_COMMAND")
		os.Unsetenv("TANZU_CLI_INVOKED_GROUP")
	}()

	p := SetupTestPlugin(t)
	assert.NotNil(t, p)

	os.Setenv("TANZU_CLI_COMMAND_MAPPED_FROM", "c0 c1 c2 c3")
	os.Setenv("TANZU_CLI_INVOKED_COMMAND", "To3")
	os.Setenv("TANZU_CLI_INVOKED_GROUP", "jump")

	ic := GetInvocationContext()
	assert.NotNil(t, ic)

	result := ic.MappedSourceCommandPath()
	assert.Equal(t, "c0 c1 c2 c3", result)

	result = ic.InvokedCommandName()
	assert.Equal(t, "To3", result)

	result = ic.InvokedGroupPath()
	assert.Equal(t, "jump", result)

	result = ic.CLIInvocationString()
	assert.Equal(t, "jump To3", result)

	// right at mapped command
	cmd := subCommandAtLevel(p.Cmd, 3)
	assert.NotNil(t, cmd)
	result = ic.CLIInvocationStringForCommand(cmd)
	assert.Equal(t, "jump To3", result)

	// command not found below mapped command
	cmd = subCommandAtLevel(p.Cmd, 2)
	assert.NotNil(t, cmd)
	result = ic.CLIInvocationStringForCommand(cmd)
	assert.Equal(t, "jump To3", result)

	// deeper command
	cmd = subCommandAtLevel(p.Cmd, 5)
	assert.NotNil(t, cmd)
	result = ic.CLIInvocationStringForCommand(cmd)
	assert.Equal(t, "jump To3 c4 c5", result)

	// same command under different invocation context
	os.Setenv("TANZU_CLI_COMMAND_MAPPED_FROM", "c0 c1")
	os.Setenv("TANZU_CLI_INVOKED_COMMAND", "To1")
	os.Setenv("TANZU_CLI_INVOKED_GROUP", "")
	ic = GetInvocationContext()
	assert.NotNil(t, ic)
	result = ic.CLIInvocationStringForCommand(cmd)
	assert.Equal(t, "To1 c2 c3 c4 c5", result)
}
