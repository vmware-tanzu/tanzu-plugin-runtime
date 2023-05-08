// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package envflags

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// ConstructGetEnvConfigurationsInputOptions creates a GetEnvConfigurationsInputOptions
func ConstructGetEnvConfigurationsInputOptions(version core.RuntimeVersion) *GetEnvConfigurationsInputOptions {
	return &GetEnvConfigurationsInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}
}

// ConstructGetEnvConfigurationsOutputOptions creates a GetEnvConfigurationsOutputOptions with val
func ConstructGetEnvConfigurationsOutputOptions(version core.RuntimeVersion, val map[string]string) *GetEnvConfigurationsOutputOptions {
	return &GetEnvConfigurationsOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Value: val,
	}
}

// ConstructGetEnvInputOptions creates a GetEnvInputOptions object with key
func ConstructGetEnvInputOptions(version core.RuntimeVersion, key string) *GetEnvInputOptions {
	return &GetEnvInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		KeyName: key,
	}
}

// ConstructGetEnvOutputOptions creates GetEnvOutputOptions with val
func ConstructGetEnvOutputOptions(version core.RuntimeVersion, val string) *GetEnvOutputOptions {
	return &GetEnvOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Value: val,
	}
}

// ConstructGetEnvOutputOptionsWithError creates a GetEnvOutputOptions with err string
func ConstructGetEnvOutputOptionsWithError(version core.RuntimeVersion, err string) *GetEnvOutputOptions {
	return &GetEnvOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Error: err,
	}
}

// ConstructSetEnvInputOptions creates SetEnvInputOptions with key and value
func ConstructSetEnvInputOptions(version core.RuntimeVersion, key, value string) *SetEnvInputOptions {
	return &SetEnvInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		KeyName:   key,
		ValueName: value,
	}
}

// ConstructDeleteEnvInputOptions creates DeleteEnvInputOptions with key
func ConstructDeleteEnvInputOptions(version core.RuntimeVersion, key string) *DeleteEnvInputOptions {
	return &DeleteEnvInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		KeyName: key,
	}
}
