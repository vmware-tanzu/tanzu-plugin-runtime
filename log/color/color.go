/*
Copyright 2023 VMware, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package color

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

var (
	FaintColor   = color.New(color.Faint)
	InfoColor    = color.New(color.FgCyan)
	SuccessColor = color.New(color.FgGreen)
	WarnColor    = color.New(color.FgYellow)
	ErrorColor   = color.New(color.FgRed)
	BoldColor    = color.New(color.Bold)
)

func Infof(format string, a ...interface{}) string {
	return InfoColor.Sprintf(format, a...)
}

func Warnf(format string, a ...interface{}) string {
	return WarnColor.Sprintf(format, a...)
}

func Errorf(format string, a ...interface{}) string {
	return ErrorColor.Sprintf(format, a...)
}

func Successf(format string, a ...interface{}) string {
	return SuccessColor.Sprintf(format, a...)
}

func Faintf(format string, a ...interface{}) string {
	return FaintColor.Sprintf(format, a...)
}

func Boldf(format string, a ...interface{}) string {
	return BoldColor.Sprintf(format, a...)
}

func IsTTYEnabled() bool {
	ttyEnabled := true
	if os.Getenv("TANZU_CLI_NO_COLOR") != "" || os.Getenv("NO_COLOR") != "" || strings.EqualFold(os.Getenv("TERM"), "DUMB") || !isatty.IsTerminal(os.Stdout.Fd()) {
		ttyEnabled = false
	}
	return ttyEnabled
}
