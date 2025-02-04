// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package stringutils

import (
	"fmt"
	"strings"
	"unicode"
)

// TrimRightSpace removes space (including all unicode space character) from right
func TrimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// Rpad adds padding to the right of a string.
// from https://github.com/spf13/cobra/blob/4ba5566f5704a9c0d205e1ef3efc4896156d33fa/cobra.go#L173-L177
func Rpad(s string, padding int) string {
	tmpl := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(tmpl, s)
}
