// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/component"
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
const CmdTemplate = `
{{- bold "Usage:" -}}

{{- if .Runnable -}}
{{- $target := index .Annotations "target" -}}
{{- /* For kubernetes, k8s, global, or no target display tanzu command path without target*/ -}}
{{- if or (eq $target "kubernetes") (eq $target "k8s") (eq $target "global") (eq $target "") }}
 tanzu {{.UseLine}}
{{- end -}}
{{- /* For non global, or no target display tanzu command path with target*/ -}}
{{- if and (ne $target "global") (ne $target "") }}
 tanzu {{ $target }} {{.UseLine}}
{{- end -}}
{{- print "\n" -}}
{{- end -}}

{{- if .HasAvailableSubCommands -}}
{{- $target := index .Annotations "target" -}}
{{- /* For kubernetes, k8s, global, or no target display tanzu command path without target*/ -}}
{{- if or (eq $target "kubernetes") (eq $target "k8s") (eq $target "global") (eq $target "") }}
 tanzu {{.CommandPath}} [command]
{{- end -}}
{{- /* For non global, or no target display tanzu command path with target*/ -}}
{{- if and (ne $target "global") (ne $target "") }}
 tanzu {{ $target }} {{.CommandPath}} [command]
{{- end -}}
{{- print "\n" -}}
{{- end -}}

{{- /* Display Aliases for the plugin if specified*/ -}}
{{ if gt (len .Aliases) 0 }}
{{ bold "Aliases:" }}
  {{.NameAndAliases}}
{{- print "\n" -}}
{{- end -}}

{{- /* Display Examples for the plugin if specified*/ -}}
{{ if .HasExample }}
{{ bold "Examples:" }}
  {{.Example}}
{{- print "\n" -}}
{{- end -}}

{{- /* Display Available Commands for the plugin*/ -}}
{{ if .HasAvailableSubCommands }}
{{ bold "Available Commands:" }}{{range .Commands}}
{{- if .IsAvailableCommand }}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end -}}
{{- end -}}
{{- print "\n" -}}
{{- end -}}

{{- /* Display Flags of the plugin*/ -}}
{{ if .HasAvailableLocalFlags}}
{{ bold "Flags:" }}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{- print "\n" -}}
{{- end -}}

{{- /* Display Global Flags of the plugin*/ -}}
{{ if .HasAvailableInheritedFlags}}
{{ bold "Global Flags:" }}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}
{{- print "\n" -}}
{{- end -}}

{{- /* Display Additional help topics of the plugin*/ -}}
{{ if .HasHelpSubCommands}}
{{ bold "Additional help topics:" }}{{range .Commands}}
{{- if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}
{{- end -}}
{{- end -}}
{{- print "\n" -}}
{{- end -}}

{{ if .HasAvailableSubCommands}}
{{- $target := index .Annotations "target" }}
{{- if or (eq $target "kubernetes") (eq $target "k8s") (eq $target "global") (eq $target "") }}
Use "{{- if beginsWith .CommandPath "tanzu "}}{{.CommandPath}}{{- else}}tanzu {{.CommandPath}}{{- end}} [command] --help" for more information about a command.
{{- end -}}
{{- if and (ne $target "global") (ne $target "") }}
Use "{{- if beginsWith .CommandPath "tanzu "}}{{.CommandPath}}{{- else}}tanzu {{ $target }} {{.CommandPath}}{{- end}} [command] --help" for more information about a command.
{{- end -}}
{{- end }}
`

// TemplateFuncs are the template usage funcs.
var TemplateFuncs = template.FuncMap{
	"rpad":                    component.Rpad,
	"bold":                    component.Bold,
	"underline":               component.Underline,
	"trimTrailingWhitespaces": component.TrimRightSpace,
	"beginsWith":              component.BeginsWith,
}
