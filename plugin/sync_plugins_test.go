// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

const (
	fakePluginScriptFmtString string = `#!/bin/bash
# Fake tanzu core binary

# fake alternate command that simply echos args provided
bad() { echo "bad command failed"; exit 1; }
plugin() {
	# error to stderr
	>&2 echo "$@ failed"

	# regular output to stdout
	echo "$@ succeeded"

	exit %s
}

newcommand() {
	# error to stderr
	>&2 echo "$@ failed"

	# regular output to stdout
	echo "$@ succeeded"

	exit %s
}

case "$1" in
    _custom_command) "newcommand $@";;
    newcommand)   $1 "$@";;
    plugin)   $1 "$@";;
    bad)   $1 "$@";;
    *) cat << EOF
Tanzu Core CLI mock

Usage:
  tanzu [command]

Available Commands:
  newcommand  fake new command
  _custom_command provide alternate command to invoke, if available
  bad     (non-working)
EOF
       exit 0
       ;;
esac
`
)

func setupFakeCLI(dir string, exitStatus string, newCommandExitStatus string) (string, error) {
	filePath := filepath.Join(dir, "tanzu")

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fmt.Fprintf(f, fakePluginScriptFmtString, exitStatus, newCommandExitStatus)

	return filePath, nil
}

func TestSyncPlugins(t *testing.T) {
	tests := []struct {
		test            string
		exitStatus      string
		expectedOutput  string
		expectedFailure bool
	}{
		{
			test:            "with no alternate command, sync successfully",
			exitStatus:      "0",
			expectedOutput:  "plugin sync succeeded\n",
			expectedFailure: false,
		},
		{
			test:            "with no alternate command, sync unsuccessfully",
			exitStatus:      "1",
			expectedOutput:  "plugin sync failed\n",
			expectedFailure: true,
		},
	}

	for _, spec := range tests {
		dir, err := os.MkdirTemp("", "tanzu-cli-sync-api")
		assert.Nil(t, err)
		defer os.RemoveAll(dir)
		t.Run(spec.test, func(t *testing.T) {
			assert := assert.New(t)

			cliPath, err := setupFakeCLI(dir, spec.exitStatus, spec.exitStatus)
			assert.Nil(err)
			os.Setenv("TANZU_BIN", cliPath)

			output, err := SyncPluginsForTarget(types.TargetK8s)

			if spec.expectedFailure {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
			}
			assert.Equal(spec.expectedOutput, output)

			os.Unsetenv("TANZU_BIN")
		})
	}
}
