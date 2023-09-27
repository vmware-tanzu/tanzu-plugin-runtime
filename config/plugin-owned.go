// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	// PluginsBaseDir is the name of the plugins owned base directory in which plugin owned settings is stored.
	PluginsBaseDir = "plugins"
)

// GetTanzuPluginConfigDir Retrieve the tanzu configuration directory that can be used by the plugins to // create a plugin specific directory to manage plugin owned configurations.
// .config/tanzu/plugins
func GetTanzuPluginConfigDir() (string, error) {
	// Fetch the base tanzu config directory
	tanzuDir, err := LocalDir()
	if err != nil {
		return "", errors.Wrap(err, "could not find local tanzu dir for OS")
	}

	// Create plugins directory in tanzu config
	pluginsBaseDir := filepath.Join(tanzuDir, PluginsBaseDir)
	if err := os.MkdirAll(pluginsBaseDir, 0755); err != nil {
		return "", errors.Wrap(err, "could not make local tanzu plugins directory")
	}

	return pluginsBaseDir, nil
}
