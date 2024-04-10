// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"strings"
)

// InvocationContext provides details regarding how a plugin's command is being
// called by the Tanzu CLI. These details allow the plugin to, among other things,
// construct proper help information, and learn if the command being invoked
// is done via a command-level mapping or not.
type InvocationContext struct {
	// invokedGroup is a space-delimited portion of the Tanzu CLI command
	// invocation between the CLI binary and the command name itself.
	// Empty when invoking a top-level command, e.g. "tanzu apply".
	invokedGroup string

	// invokedCommand is the name of the command in a Tanzu CLI command invocation
	invokedCommand string

	// sourceCommandPath is a space-delimited path relative to the plugin's
	// root command of the command being invoked.
	// This value is empty when the CLI command invoked is mapped to the
	// plugin's root command,
	sourceCommandPath string
}

func (ic *InvocationContext) String() string {
	return strings.TrimSpace(ic.invokedGroup + " " + ic.invokedCommand)
}

// GetInvocationContext returns information about how a Tanzu CLI command is invoked.
// Note that at the moment a valid InvocationContext is only returned when the
// invoked plugin command (or its ancestor) has been remapped
func GetInvocationContext() *InvocationContext {
	invokedGroup := os.Getenv("TANZU_CLI_INVOKED_GROUP")
	invokedCommand := os.Getenv("TANZU_CLI_INVOKED_COMMAND")
	sourceCommandPath := os.Getenv("TANZU_CLI_COMMAND_MAPPED_FROM")

	if invokedGroup != "" || invokedCommand != "" {
		return &InvocationContext{
			invokedGroup:      invokedGroup,
			invokedCommand:    invokedCommand,
			sourceCommandPath: sourceCommandPath,
		}
	}

	return nil
}
