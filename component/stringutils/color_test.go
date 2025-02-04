// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package stringutils provides helpers to format strings
package stringutils

import (
	"testing"
)

// Test StripColor function
func TestStripColor(t *testing.T) {
	coloredText := SuccessColor.Sprintf("Success: %s", "operation completed")
	expected := "Success: operation completed"

	strippedText := StripColor(coloredText)
	if strippedText != expected {
		t.Errorf("StripColor(%q) = %q; want %q", coloredText, strippedText, expected)
	}
}
