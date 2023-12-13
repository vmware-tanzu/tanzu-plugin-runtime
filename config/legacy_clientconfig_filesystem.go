// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"path/filepath"
)

var (
	// legacyLocalDirName is the name of the old local directory in which to look for tanzu state. This will be
	// removed in the future in favor of LocalDirName.
	// Deprecated: This legacyLocalDirName is deprecated. Use next get config apis
	legacyLocalDirName = ".tanzu"
)

// Deprecated: This legacyLocalDir is deprecated. Use next get config apis
func legacyLocalDir() (path string, err error) {
	return localDirPath(legacyLocalDirName)
}

// legacyConfigPath returns the legacy tanzu config path, checking for environment overrides.
// Deprecated: This legacyConfigPath is deprecated. Use next get config apis
func legacyConfigPath() (path string, err error) {
	return legacyCfgPath(legacyLocalDir)
}

// legacyCfgPath constructs the full config path
// Deprecated: This legacyCfgPath is deprecated. Use next get config apis
func legacyCfgPath(localDirGetter func() (string, error)) (path string, err error) {
	localDir, err := localDirGetter()
	if err != nil {
		return path, err
	}
	path = filepath.Join(localDir, ConfigName)
	return path, nil
}
