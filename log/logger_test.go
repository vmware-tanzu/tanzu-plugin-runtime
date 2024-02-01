// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"bytes"
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		title                 string
		initialLogLevel       string
		logLevel              string
		containStrings        []string
		doesNotContainStrings []string
	}{

		{
			title:                 "when TANZU_CLI_LOG_LEVEL is not configured and then set to 3",
			initialLogLevel:       "",
			logLevel:              "3",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3"},
			doesNotContainStrings: []string{"log-5", "log-8"},
		},
		{
			title:                 "when TANZU_CLI_LOG_LEVEL is set to 1 and 5",
			initialLogLevel:       "1",
			logLevel:              "5",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-5"},
			doesNotContainStrings: []string{"log-3", "log-8"},
		},
		{
			title:                 "when TANZU_CLI_LOG_LEVEL is set to 3 and 6",
			initialLogLevel:       "3",
			logLevel:              "6",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3", "log-5"},
			doesNotContainStrings: []string{"log-8"},
		},
		{
			title:                 "when TANZU_CLI_LOG_LEVEL is set to 3 and 8",
			initialLogLevel:       "3",
			logLevel:              "8",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3", "log-5", "log-8"},
			doesNotContainStrings: []string{},
		},
		{
			title:                 "when TANZU_CLI_LOG_LEVEL is set to 5 and 9",
			initialLogLevel:       "5",
			logLevel:              "9",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3", "log-5", "log-8"},
			doesNotContainStrings: []string{},
		},
		{
			title:                 "when TANZU_CLI_LOG_LEVEL is configured with incorrect value",
			initialLogLevel:       "a",
			logLevel:              "a",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3"},
			doesNotContainStrings: []string{"log-5", "log-8"},
		},
	}

	for _, spec := range tests {
		t.Run(spec.title, func(t *testing.T) {
			defer os.Setenv(EnvTanzuCLILogLevel, "")

			l = NewLogger()
			stderr := &bytes.Buffer{}
			SetStderr(stderr)

			os.Setenv(EnvTanzuCLILogLevel, spec.initialLogLevel)

			Info("log-default")
			V(0).Info("log-0")
			V(1).Info("log-1")
			V(3).Info("log-3")

			os.Setenv(EnvTanzuCLILogLevel, spec.logLevel)

			V(5).Info("log-5")
			V(8).Info("log-8")

			for _, logStr := range spec.containStrings {
				assert.Contains(t, stderr.String(), logStr, spec.title)
			}
			for _, logStr := range spec.doesNotContainStrings {
				assert.NotContains(t, stderr.String(), logStr, spec.title)
			}
		})
	}
}
