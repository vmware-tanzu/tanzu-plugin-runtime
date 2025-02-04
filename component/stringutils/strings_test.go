// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package stringutils

import (
	"testing"
)

func TestTrimRightSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello   ", "hello"},
		{"test\n\t", "test"},
		{"  no trim needed", "  no trim needed"},
		{"", ""},
	}

	for _, tt := range tests {
		result := TrimRightSpace(tt.input)
		if result != tt.expected {
			t.Errorf("TrimRightSpace(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestBeginsWith(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
		{"", "", true},
		{"", "nonempty", false},
		{"prefix test", "prefix", true},
	}

	for _, tt := range tests {
		result := BeginsWith(tt.s, tt.prefix)
		if result != tt.expected {
			t.Errorf("BeginsWith(%q, %q) = %v; want %v", tt.s, tt.prefix, result, tt.expected)
		}
	}
}

func TestRpad(t *testing.T) {
	tests := []struct {
		s        string
		padding  int
		expected string
	}{
		{"test", 8, "test    "},
		{"hello", 3, "hello"},
		{"pad", 6, "pad   "},
		{"", 4, "    "},
	}

	for _, tt := range tests {
		result := Rpad(tt.s, tt.padding)
		if result != tt.expected {
			t.Errorf("Rpad(%q, %d) = %q; want %q", tt.s, tt.padding, result, tt.expected)
		}
	}
}
