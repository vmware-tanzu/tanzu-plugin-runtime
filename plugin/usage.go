// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
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

func matchOnCommandName(cmd *cobra.Command, value string) bool {
	return cmd.Name() == value
}

func findSubCommandByHierarchy(cmd *cobra.Command, hierarchy []string, matcher func(*cobra.Command, string) bool) (*cobra.Command, *cobra.Command) {
	if len(hierarchy) == 0 {
		parent := cmd.Parent()
		return cmd, parent
	}

	childCmds := cmd.Commands()
	for i := range childCmds {
		if len(hierarchy) == 1 {
			if matcher(childCmds[i], hierarchy[0]) {
				return childCmds[i], childCmds[i].Parent()
			}
		} else {
			if childCmds[i].Name() == hierarchy[0] {
				return findSubCommandByHierarchy(childCmds[i], hierarchy[1:], matcher)
			}
		}
	}
	if len(hierarchy) == 1 {
		return nil, cmd
	}
	return nil, nil
}

// hierarchyFromMappedCommand returns list of command names relative to the mapped
// command (as indicated by mappedCommandPath) of subcommands to traverse to cmd
func hierarchyFromMappedCommandPath(mappedCommandPath string, cmd *cobra.Command) ([]string, error) {
	var fromCmd *cobra.Command
	var additionalCmdNames []string

	rootCmd := cmd.Root()

	if mappedCommandPath == "" {
		fromCmd = rootCmd
	} else {
		hierarchy := strings.Fields(mappedCommandPath)
		fromCmd, _ = findSubCommandByHierarchy(rootCmd, hierarchy, matchOnCommandName)
	}

	for cmd != fromCmd && cmd.HasParent() {
		additionalCmdNames = append([]string{cmd.Name()}, additionalCmdNames...)
		cmd = cmd.Parent()
	}

	if cmd != fromCmd {
		return []string{}, fmt.Errorf("fail to locate mapped command path '%s'", mappedCommandPath)
	}

	return additionalCmdNames, nil
}

func useLineEx(cmd *cobra.Command, ic *InvocationContext) string {
	useline := getInvocationStringForUseLine(cmd, ic)

	if cmd.DisableFlagsInUseLine {
		return useline
	}

	if cmd.HasAvailableFlags() && !strings.Contains(useline, "[flags]") {
		useline += " [flags]"
	}

	return useline
}

func getInvocationStringForUseLine(cmd *cobra.Command, ic *InvocationContext) string {
	// by checking sourceCommandPath we limit the use of the InvocationContext
	// to only command-level (not plugin level) mapping
	if ic == nil || ic.sourceCommandPath == "" {
		return cmd.UseLine()
	}

	hierarchy, err := hierarchyFromMappedCommandPath(ic.sourceCommandPath, cmd)
	if err != nil {
		return ic.CLIInvocationString()
	}

	// we are executing the actual mapped command
	if len(hierarchy) == 0 {
		found := strings.HasPrefix(cmd.Use, cmd.Name())

		// in this case, it is acceptable to leverage command's .Use string if
		// it is prefixed by the command name as long as the prefix is stripped
		if found {
			tail := strings.TrimPrefix(cmd.Use, cmd.Name())
			return ic.CLIInvocationString() + tail
		}
		return ic.CLIInvocationString()
	}

	// if command is deeper than the mapped command, its .Use string is
	// considered safe to use in replacement of the command name as long as the
	// former is prefixed by the latter
	lastCommandName := hierarchy[len(hierarchy)-1]
	if strings.HasPrefix(cmd.Use, lastCommandName) {
		hierarchy[len(hierarchy)-1] = cmd.Use
	}

	return ic.CLIInvocationString() + " " + strings.Join(hierarchy, " ")
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
	// if root cmd
	if v, ok := cmd.Annotations[cobra.CommandDisplayNameAnnotation]; ok {
		cmdName = v
	} else {
		ic := GetInvocationContext()
		if ic != nil && ic.MappedSourceCommandPath() != "" {
			hierarchy, err := hierarchyFromMappedCommandPath(ic.MappedSourceCommandPath(), cmd)
			// cmd is actually the one being command-level mapped
			if err == nil && len(hierarchy) == 0 {
				cmdName = ic.InvokedCommandName()
			}
		}
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

		// matches cobra default help template's behavior of not indenting the
		// Example value, which has the added benefit of ensuring multiline
		// .Example values are aligned
		output.WriteString(cmd.Example + "\n")
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
