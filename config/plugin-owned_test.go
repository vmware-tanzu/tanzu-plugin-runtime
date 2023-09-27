// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTanzuPluginConfigDir(t *testing.T) {
	// setup
	func() {
		LocalDirName = ".config2/tanzu"
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()

	expectedPath := ".config2/tanzu/plugins"

	dir, err := GetTanzuPluginConfigDir()

	require.NoError(t, err, "Expected no error from GetTanzuPluginConfigDir")
	assert.NotEmpty(t, dir, "Expected a non-empty directory path")

	assert.Contains(t, dir, expectedPath)
}
