// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package tabwriter exposes an tabwriter functionality
// from https://github.com/kubernetes/cli-runtime/blob/v0.28.1/pkg/printers/tabwriter.go
package tabwriter

import (
	"io"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/component/tabwriter/internal"
)

const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
	tabwriterFlags    = internal.RememberWidths | internal.IgnoreAnsiCodes
)

var (
	tabwriterPaddingStart int
)

// GetNewTabWriter returns a tabwriter that translates tabbed columns in input into properly aligned text.
func GetNewTabWriter(output io.Writer) *internal.Writer {
	return internal.NewWriter(output, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterPaddingStart, tabwriterFlags)
}
