// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// SyncPluginsForTarget will attempt to install plugins required by the active
// Context of the provided target. This is most useful for any plugin
// implementation which creates a new Context or updates an existing one as
// part of its operation, and prefers that the plugins appropriate for the
// Context are immediately available for use.
//
// Note: This API is considered EXPERIMENTAL. Both the function's signature and
// implementation are subjected to change/removal if an alternative means to
// provide equivalent functionality can be introduced.
//
// Stdout output of the command is returned as a string on successful
// execution of the command, otherwise, the Stderr output is returned instead.
func SyncPluginsForTarget(target types.Target) (string, error) {
	// For now, the implementation expects env var TANZU_BIN to be set and
	// pointing to the core CLI binary used to invoke the plugin sync with.

	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cliPath := os.Getenv("TANZU_BIN")
	command := exec.Command(cliPath, "plugin", "sync")
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()

	if err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}
