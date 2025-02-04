// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"io"
	"strings"
)

// From https://github.com/kubernetes/cli-runtime/blob/v0.28.1/pkg/printers/terminal.go"

// terminalEscaper replaces ANSI escape sequences and other terminal special
// characters to avoid terminal escape character attacks (https://github.com/kubernetes/kubernetes/issues/101695).
// Add "\x1b", "^[" to the `NewReplacer` params to escape color
var terminalEscaper = strings.NewReplacer("\x1b", "^[", "\r", "\\r")

// WriteEscaped replaces unsafe terminal characters with replacement strings
// and writes them to the given writer.
func WriteEscaped(writer io.Writer, output string) error {
	_, err := terminalEscaper.WriteString(writer, output)
	return err
}

// EscapeTerminal escapes terminal special characters in a human readable (but
// non-reversible) format.
func EscapeTerminal(in string) string {
	return terminalEscaper.Replace(in)
}
