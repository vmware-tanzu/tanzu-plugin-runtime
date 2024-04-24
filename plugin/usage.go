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
	t, err := template.New("usage").Funcs(TemplateFuncs).Parse(cmdTemplate)
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, c)
}

// CmdTemplate is the template for plugin commands.
// Deprecated: This variable is deprecated.
const CmdTemplate = `{{ bold "Usage:" }}
  {{if .Runnable}}{{ $target := index .Annotations "target" }}{{ if or (eq $target "kubernetes") (eq $target "k8s") }}tanzu {{.UseLine}}{{ end }}{{ if and (ne $target "global") (ne $target "") }}tanzu {{ $target }} {{ else }} {{ end }}{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}{{ $target := index .Annotations "target" }}{{ if or (eq $target "kubernetes") (eq $target "k8s") }}tanzu {{.CommandPath}} [command]{{end}}{{ if and (ne $target "global") (ne $target "") }}tanzu {{ $target }} {{ else }} {{ end }}{{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

{{ bold "Aliases:" }}
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

{{ bold "Examples:" }}
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

{{ bold "Available Commands:" }}{{range .Commands}}{{if .IsAvailableCommand }}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{ bold "Flags:" }}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{ bold "Global Flags:" }}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{ bold "Additional help topics:" }}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

{{ $target := index .Annotations "target" }}{{ if or (eq $target "kubernetes") (eq $target "k8s") }}Use "{{if beginsWith .CommandPath "tanzu "}}{{.CommandPath}}{{else}}tanzu {{.CommandPath}}{{end}} [command] --help" for more information about a command.{{end}}Use "{{if beginsWith .CommandPath "tanzu "}}{{.CommandPath}}{{else}}tanzu{{ $target := index .Annotations "target" }}{{ if and (ne $target "global") (ne $target "") }} {{ $target }} {{ else }} {{ end }}{{.CommandPath}}{{end}} [command] --help" for more information about a command.{{end}}
`

// cmdTemplate is the template for plugin commands.
const cmdTemplate = `{{ printHelp . }}`

// Constants for help text labels
const (
	usageStr                = "Usage:"
	aliasesStr              = "Aliases:"
	examplesStr             = "Examples:"
	availableCommandsStr    = "Available Commands:"
	flagsStr                = "Flags:"
	globalFlagsStr          = "Global Flags:"
	additionalHelpTopicsStr = "Additional help topics:"
	indentStr               = "  "
)

// return space delimited concatenation of each non empty string in the array,
// in the order as provided
func buildInvocationString(parts ...string) string {
	var nonEmptyParts []string
	for _, s := range parts {
		if strings.TrimSpace(s) != "" {
			nonEmptyParts = append(nonEmptyParts, s)
		}
	}
	return strings.Join(nonEmptyParts, " ")
}

func useLineEx(cmd *cobra.Command, ic *InvocationContext) string {
	// by checking sourceCommandPath we limit the use of the InvocationContext
	// to only command-level (not plugin level) mapping
	if ic == nil || ic.sourceCommandPath == "" {
		return cmd.UseLine()
	}

	// TODO(vuil) look into still incorporating relevant parts of UseLine into output
	return ic.CLIInvocationString()
}

func commandPathEx(cmd *cobra.Command, ic *InvocationContext) string {
	if ic == nil || ic.sourceCommandPath == "" {
		return cmd.CommandPath()
	}
	return ic.CLIInvocationString()
}

// Helper to format the usage help section.
func formatUsageHelpSection(cmd *cobra.Command, target types.Target) string {
	var output strings.Builder
	ic := GetInvocationContext()

	output.WriteString(component.Bold(usageStr) + "\n")
	base := indentStr + "tanzu"

	if cmd.Runnable() {
		// For kubernetes, k8s, global, or no target display tanzu command path without target
		if target == types.TargetK8s || target == types.TargetGlobal || target == types.TargetUnknown {
			output.WriteString(buildInvocationString(base, useLineEx(cmd, ic)) + "\n")
		}

		// For non global, or no target ;display tanzu command path with target
		if target != types.TargetGlobal && target != types.TargetUnknown {
			output.WriteString(buildInvocationString(base, string(target), useLineEx(cmd, ic)) + "\n")
		}
	}

	if cmd.HasAvailableSubCommands() {
		if cmd.Runnable() {
			// If the command is both Runnable and has sub-commands, let's insert an empty
			// line between the usage for the Runnable and the one for the sub-commands
			output.WriteString("\n")
		}
		// For kubernetes, k8s, global, or no target display tanzu command path without target
		if target == types.TargetK8s || target == types.TargetGlobal || target == types.TargetUnknown {
			output.WriteString(buildInvocationString(base, commandPathEx(cmd, ic), "[command]") + "\n")
		}

		// For non global, or no target display tanzu command path with target
		if target != types.TargetGlobal && target != types.TargetUnknown {
			output.WriteString(buildInvocationString(base, string(target), commandPathEx(cmd, ic), "[command]") + "\n")
		}
	}
	return output.String()
}

// Helper to format the help footer.
func formatHelpFooter(cmd *cobra.Command, target types.Target) string {
	var footer strings.Builder
	if !cmd.HasAvailableSubCommands() {
		return ""
	}

	footer.WriteString("\n")

	ic := GetInvocationContext()
	base := "Use \""
	if !strings.HasPrefix(cmd.CommandPath(), "tanzu ") {
		base = "Use \"tanzu"
	}

	// For kubernetes, k8s, global, or no target display tanzu command path without target
	if target == types.TargetK8s || target == types.TargetGlobal || target == types.TargetUnknown {
		footer.WriteString(buildInvocationString(base, commandPathEx(cmd, ic), `[command] --help" for more information about a command.`+"\n"))
	}

	// For non global, or no target display tanzu command path with target
	if target != types.TargetGlobal && target != types.TargetUnknown {
		footer.WriteString(buildInvocationString(base, string(target), commandPathEx(cmd, ic), `[command] --help" for more information about a command.`+"\n"))
	}

	return footer.String()
}

func aliasesWithMappedName(cmd *cobra.Command) string {
	cmdName := cmd.Name()
	if v, ok := cmd.Annotations[cobra.CommandDisplayNameAnnotation]; ok {
		cmdName = v
	}
	return strings.Join(append([]string{cmdName}, cmd.Aliases...), ", ")
}

func printHelp(cmd *cobra.Command) string {
	var output strings.Builder
	target := types.StringToTarget(cmd.Annotations["target"])

	output.WriteString(formatUsageHelpSection(cmd, target))

	if len(cmd.Aliases) > 0 {
		output.WriteString("\n" + component.Bold(aliasesStr) + "\n")
		output.WriteString(indentStr + aliasesWithMappedName(cmd) + "\n")
	}

	if cmd.HasExample() {
		output.WriteString("\n" + component.Bold(examplesStr) + "\n")
		output.WriteString(indentStr + cmd.Example + "\n")
	}

	if cmd.HasAvailableSubCommands() {
		output.WriteString("\n" + component.Bold(availableCommandsStr) + "\n")
		for _, c := range cmd.Commands() {
			if c.IsAvailableCommand() {
				output.WriteString(indentStr + component.Rpad(c.Name(), c.NamePadding()) + " " + c.Short + "\n")
			}
		}
	}

	if cmd.HasAvailableLocalFlags() {
		output.WriteString("\n" + component.Bold(flagsStr) + "\n")
		output.WriteString(strings.TrimRight(cmd.LocalFlags().FlagUsages(), " "))
	}

	if cmd.HasAvailableInheritedFlags() {
		output.WriteString("\n" + component.Bold(globalFlagsStr) + "\n")
		output.WriteString(strings.TrimRight(cmd.InheritedFlags().FlagUsages(), " "))
	}

	if cmd.HasHelpSubCommands() {
		output.WriteString("\n" + component.Bold(additionalHelpTopicsStr) + "\n")
		for _, c := range cmd.Commands() {
			if c.IsAdditionalHelpTopicCommand() {
				output.WriteString(indentStr + component.Rpad(c.CommandPath(), c.CommandPathPadding()) + " " + c.Short + "\n")
			}
		}
	}
	output.WriteString(formatHelpFooter(cmd, target))

	return output.String()
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
