// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"bytes"
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

# fake command that simulates a plugin lcm operation
plugin() {
	if [ "%s" -eq "0" ]; then
		# regular output to stderr
		>&2 echo "$@ succeeded"
	else
		# error to stderr
		>&2 echo "$@ failed"
	fi

	exit %s
}

# fake alternate command to use
newcommand() {
	if [ "%s" -eq "0" ]; then
		# regular output to stdout
		echo "$@ succeeded"
	else
		# error to stderr
		>&2 echo "$@ failed"
	fi

	exit %s
}

case "$1" in
    # simulate returning an alternative set of args to invoke with, which
    # translates to running the command 'newcommand'
    %s) shift && shift && echo "newcommand $@";;
    newcommand)   $1 "$@";;
    plugin)   $1 "$@";;
    *) cat << EOF
Tanzu Core CLI Fake

Usage:
  tanzu [command]

Available Commands:
  plugin          fake command
  newcommand      fake new command
  _custom_command provide alternate command to invoke, if available
EOF
       exit 1
       ;;
esac
`
)

func setupFakeCLI(dir string, exitStatus string, newCommandExitStatus string, enableCustomCommand bool) (string, error) {
	filePath := filepath.Join(dir, "tanzu")

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fakeCustomCommandName := "unused_command"
	// when enabled, the fake CLI script generated will be capable of
	// returning an alternate set of args for a provided set of args
	if enableCustomCommand {
		fakeCustomCommandName = customCommandName
	}

	fmt.Fprintf(f, fakePluginScriptFmtString, exitStatus, exitStatus, newCommandExitStatus, newCommandExitStatus, fakeCustomCommandName)

	return filePath, nil
}

func TestSyncPlugins(t *testing.T) {
	tests := []struct {
		test                 string
		exitStatus           string
		newCommandExitStatus string
		expectedOutput       string
		expectedFailure      bool
		enableCustomCommand  bool
	}{
		{
			test:            "with no alternate command and sync successfully",
			exitStatus:      "0",
			expectedOutput:  "plugin sync succeeded\n",
			expectedFailure: false,
		},
		{
			test:            "with no alternate command and sync unsuccessfully",
			exitStatus:      "1",
			expectedOutput:  "plugin sync failed\n",
			expectedFailure: true,
		},
		{
			test:                 "with alternate command and sync successfully",
			newCommandExitStatus: "0",
			expectedOutput:       "newcommand sync --target kubernetes succeeded\n",
			expectedFailure:      false,
			enableCustomCommand:  true,
		},
		{
			test:                 "with alternate command and sync unsuccessfully",
			newCommandExitStatus: "1",
			expectedOutput:       "newcommand sync --target kubernetes failed\n",
			expectedFailure:      true,
			enableCustomCommand:  true,
		},
	}

	for _, spec := range tests {
		dir, err := os.MkdirTemp("", "tanzu-cli-sync-api")
		assert.Nil(t, err)
		defer os.RemoveAll(dir)
		t.Run(spec.test, func(t *testing.T) {
			assert := assert.New(t)

			// Set up stdout and stderr for our test
			r, w, err := os.Pipe()
			if err != nil {
				t.Error(err)
			}
			c := make(chan []byte)
			go readOutput(t, r, c)
			stdout := os.Stdout
			stderr := os.Stderr
			defer func() {
				os.Stdout = stdout
				os.Stderr = stderr
			}()
			os.Stdout = w
			os.Stderr = w

			cliPath, err := setupFakeCLI(dir, spec.exitStatus, spec.newCommandExitStatus, spec.enableCustomCommand)
			assert.Nil(err)
			os.Setenv("TANZU_BIN", cliPath)

			// Test-1:
			// - verify correct combinedOutput string returned as part of the output
			// - verify correct string gets printed to default stdout and stderr
			combinedOutput, err := SyncPluginsForTarget(types.TargetK8s)
			w.Close()
			stdoutRecieved := <-c

			if spec.expectedFailure {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
			}
			assert.Equal(spec.expectedOutput, combinedOutput, "incorrect combinedOutput result")
			assert.Equal(spec.expectedOutput, string(stdoutRecieved), "incorrect combinedOutput result")

			// Test-2: when external stdout and stderr are provided with WithStdout, WithStderr options,
			// verify correct string gets printed to provided custom stdout/stderr
			var combinedOutputBuff bytes.Buffer
			combinedOutput, err = SyncPluginsForTarget(types.TargetK8s, WithOutputWriter(&combinedOutputBuff), WithErrorWriter(&combinedOutputBuff))
			if spec.expectedFailure {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
			}
			assert.Equal(spec.expectedOutput, combinedOutput, "incorrect combinedOutput result when external stdout/stderr is provided")
			assert.Equal(spec.expectedOutput, combinedOutputBuff.String(), "incorrect combinedOutputBuff result")

			// Test-3: when user asks to discard the stdout and stderr, it should not print it to any stdout/stderr by default
			// but still return the combinedOutput string as part of the function return value
			combinedOutput, err = SyncPluginsForTarget(types.TargetK8s, WithNoStdout(), WithNoStderr())
			if spec.expectedFailure {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
			}
			assert.Equal(spec.expectedOutput, combinedOutput, "incorrect combinedOutput result when external stdout/stderr is provided")

			os.Unsetenv("TANZU_BIN")
		})
	}
}
