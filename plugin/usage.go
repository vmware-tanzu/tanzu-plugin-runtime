// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/component"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// UsageFunc is the usage func for a plugin.
var UsageFunc = func(c *cobra.Command) error {
	t, err := template.New("usage").Funcs(TemplateFuncs).Parse(CmdTemplate)
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, c)
}

// CmdTemplate is the template for plugin commands.
const CmdTemplate = `{{ printHelp . }}`

//nolint:all
func printHelp(cmd cobra.Command) string {
	var output string
	target := types.StringToTarget(cmd.Annotations["target"])

	output += component.Bold("Usage:") + "\n"

	// Display usage for commands that are runnable
	if cmd.Runnable() {
		// For kubernetes, k8s, global, or no target display tanzu command path without target
		if target == types.TargetK8s || target == types.TargetGlobal || target == types.TargetUnknown {
			output += "  tanzu " + cmd.UseLine() + "\n"
		}

		// For non global, or no target display tanzu command path with target
		if target != types.TargetGlobal && target != types.TargetUnknown {
			output += "  tanzu " + string(target) + " " + cmd.UseLine() + "\n"
		}
	}

	// Display usage for commands that have sub-commands
	if cmd.HasAvailableSubCommands() {
		if cmd.Runnable() {
			// If the command is both Runnable and has sub-commands, let's insert an empty
			// line between the usage for the Runnable and the one for the sub-commands
			output += "\n"
		}
		// For kubernetes, k8s, global, or no target display tanzu command path without target
		if target == types.TargetK8s || target == types.TargetGlobal || target == types.TargetUnknown {
			output += "  tanzu " + cmd.CommandPath() + " [command]\n"
		}

		// For non global, or no target display tanzu command path with target
		if target != types.TargetGlobal && target != types.TargetUnknown {
			output += "  tanzu " + string(target) + " " + cmd.CommandPath() + " [command]\n"
		}
	}

	// Display Aliases for the plugin if specified
	if len(cmd.Aliases) > 0 {
		output += "\n" + component.Bold("Aliases:") + "\n"
		output += "  " + cmd.NameAndAliases() + "\n"
	}

	// Display Examples for the plugin if specified
	if cmd.HasExample() {
		output += "\n" + component.Bold("Examples:") + "\n"
		output += cmd.Example + "\n"
	}

	// Display Available Commands for the plugin
	if cmd.HasAvailableSubCommands() {
		output += "\n" + component.Bold("Available Commands:") + "\n"
		for _, c := range cmd.Commands() {
			if c.IsAvailableCommand() {
				output += "  " + component.Rpad(c.Name(), c.NamePadding()) + " " + c.Short + "\n"
			}
		}
	}

	// Display Flags of the plugin
	if cmd.HasAvailableLocalFlags() {
		output += "\n" + component.Bold("Flags:") + "\n"
		output += strings.TrimRight(cmd.LocalFlags().FlagUsages(), " ")
	}

	// Display Global Flags of the plugin
	if cmd.HasAvailableInheritedFlags() {
		output += "\n" + component.Bold("Global Flags:") + "\n"
		output += strings.TrimRight(cmd.InheritedFlags().FlagUsages(), " ")
	}

	// Display Additional help topics of the plugin
	if cmd.HasHelpSubCommands() {
		output += "\n" + component.Bold("Additional help topics:") + "\n"
		for _, c := range cmd.Commands() {
			if c.IsAdditionalHelpTopicCommand() {
				output += "  " + component.Rpad(c.CommandPath(), c.CommandPathPadding()) + " " + c.Short + "\n"
			}
		}
	}

	// Display a note at the very bottom to indicate how to get more help
	if cmd.HasAvailableSubCommands() {
		output += "\n"

		// For kubernetes, k8s, global, or no target display tanzu command path without target
		if target == types.TargetK8s || target == types.TargetGlobal || target == types.TargetUnknown {
			output += `Use "` //nolint:goconst
			if !strings.HasSuffix(cmd.CommandPath(), "tanzu ") {
				output += "tanzu " //nolint:goconst
			}
			output += cmd.CommandPath() + ` [command] --help" for more information about a command.` + "\n"
		}

		// For non global, or no target display tanzu command path with target
		if target != types.TargetGlobal && target != types.TargetUnknown {
			output += `Use "`
			if !strings.HasSuffix(cmd.CommandPath(), "tanzu ") {
				output += "tanzu "
			}
			output += string(target) + " " + cmd.CommandPath() + ` [command] --help" for more information about a command.` + "\n"
		}
	}

	return output
}

// TemplateFuncs are the template usage funcs.
var TemplateFuncs = template.FuncMap{
	"printHelp": printHelp,
	// The below are not needed but are kept for backwards-compatibility
	// in case it is being used through the API
	"rpad":                    component.Rpad,
	"bold":                    component.Bold,
	"underline":               component.Underline,
	"trimTrailingWhitespaces": component.TrimRightSpace,
	"beginsWith":              component.BeginsWith,
}
