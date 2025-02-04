// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	auroraPackage "github.com/logrusorgru/aurora"
	"github.com/mattn/go-isatty"
)

var aurora auroraPackage.Aurora

func init() {
	NewAurora()
}

// Deprecated: NewAurora is being deprecated and will be removed in favor of using the `github.com/fatih/color` package
func NewAurora() auroraPackage.Aurora {
	if aurora == nil {
		aurora = auroraPackage.NewAurora(IsTTYEnabled())
	}
	return aurora
}

func IsTTYEnabled() bool {
	ttyEnabled := true
	if os.Getenv("TANZU_CLI_NO_COLOR") != "" || os.Getenv("NO_COLOR") != "" || strings.EqualFold(os.Getenv("TERM"), "DUMB") || !isatty.IsTerminal(os.Stdout.Fd()) {
		ttyEnabled = false
	}
	return ttyEnabled
}

// Rpad adds padding to the right of a string.
// from https://github.com/spf13/cobra/blob/4ba5566f5704a9c0d205e1ef3efc4896156d33fa/cobra.go#L173-L177
//
// Deprecated: Rpad is being moved under `github.com/vmware-tanzu/tanzu-plugin-runtime/component/stringutils` package
func Rpad(s string, padding int) string {
	tmpl := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(tmpl, s)
}

// Deprecated: Underline is being moved under `github.com/vmware-tanzu/tanzu-plugin-runtime/component/stringutils` package as `Sunderlinef`
func Underline(s string) string {
	return aurora.Underline(s).String()
}

// Deprecated: Bold is being moved under `github.com/vmware-tanzu/tanzu-plugin-runtime/component/stringutils` package as `Sboldf`
func Bold(s string) string {
	return aurora.Bold(s).String()
}

// Deprecated: TrimRightSpace is being moved under `github.com/vmware-tanzu/tanzu-plugin-runtime/component/stringutils` package
func TrimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// Deprecated: BeginsWith is being deprecated. Use strings.HasPrefix instead
func BeginsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// Deprecated: Green is being moved under `github.com/vmware-tanzu/tanzu-plugin-runtime/component/stringutils` package as `Ssuccessf`
func Green(s string) string {
	return aurora.Green(s).String()
}
