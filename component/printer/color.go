// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package printer

import (
	"regexp"

	"github.com/fatih/color"
)

var (
	FaintColor   = color.New(color.Faint)
	InfoColor    = color.New(color.FgCyan)
	SuccessColor = color.New(color.FgGreen)
	WarnColor    = color.New(color.FgYellow)
	ErrorColor   = color.New(color.FgRed)
	BoldColor    = color.New(color.Bold)
)

func Sfaintf(format string, a ...interface{}) string {
	return FaintColor.Sprintf(format, a...)
}

func Sinfof(format string, a ...interface{}) string {
	return InfoColor.Sprintf(format, a...)
}

func Ssuccessf(format string, a ...interface{}) string {
	return SuccessColor.Sprintf(format, a...)
}

func Swarnf(format string, a ...interface{}) string {
	return WarnColor.Sprintf(format, a...)
}

func Serrorf(format string, a ...interface{}) string {
	return ErrorColor.Sprintf(format, a...)
}

func Sboldf(format string, a ...interface{}) string {
	return BoldColor.Sprintf(format, a...)
}

// StripColor removes ANSI escape codes from a string
func StripColor(s string) string {
	ansiRegex := regexp.MustCompile("\x1b\\[(\\d+;)*\\d+m")
	return ansiRegex.ReplaceAllString(s, "")
}
