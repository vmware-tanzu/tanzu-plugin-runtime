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
	assert := assert.New(t)

	tests := []struct {
		test                  string
		logLevel              string
		containStrings        []string
		doesNotContainStrings []string
	}{
		{
			test:                  "when TANZU_CLI_LOG_LEVEL is not configured",
			logLevel:              "",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3"},
			doesNotContainStrings: []string{"log-5", "log-8"},
		},
		{
			test:                  "when TANZU_CLI_LOG_LEVEL is set to 3",
			logLevel:              "",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3"},
			doesNotContainStrings: []string{"log-5", "log-8"},
		},
		{
			test:                  "when TANZU_CLI_LOG_LEVEL is set to 6",
			logLevel:              "6",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3", "log-5"},
			doesNotContainStrings: []string{"log-8"},
		},
		{
			test:                  "when TANZU_CLI_LOG_LEVEL is set to 9",
			logLevel:              "9",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3", "log-5", "log-8"},
			doesNotContainStrings: []string{},
		},
		{
			test:                  "when TANZU_CLI_LOG_LEVEL is configured with incorrect value",
			logLevel:              "a",
			containStrings:        []string{"log-default", "log-0", "log-1", "log-3"},
			doesNotContainStrings: []string{"log-5", "log-8"},
		},
	}

	for _, spec := range tests {
		defer os.Setenv(EnvTanzuCLILogLevel, "")
		os.Setenv(EnvTanzuCLILogLevel, spec.logLevel)

		l = NewLogger()
		stderr := &bytes.Buffer{}
		SetStderr(stderr)

		Info("log-default")
		V(0).Info("log-0")
		V(1).Info("log-1")
		V(3).Info("log-3")
		V(5).Info("log-5")
		V(8).Info("log-8")

		for _, logStr := range spec.containStrings {
			assert.Contains(stderr.String(), logStr, spec.test)
		}
		for _, logStr := range spec.doesNotContainStrings {
			assert.NotContains(stderr.String(), logStr, spec.test)
		}
	}
}
