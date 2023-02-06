// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"fmt"
)

// Validate validated the whether passed version is valid and supported
func (r *RuntimeAPIVersion) Validate() (bool, error) {
	if r == nil || r.RuntimeVersion == "" {
		return false, fmt.Errorf("runtime version is mandatory")
	}

	if !r.RuntimeVersion.IsRuntimeVersionSupported() {
		return false, fmt.Errorf("runtime version %v is not supported", r.RuntimeVersion)
	}
	return true, nil
}

// IsRuntimeVersionSupported check whether passed version is currently supported
func (version RuntimeVersion) IsRuntimeVersionSupported() bool {
	for _, v := range SupportedRuntimeVersions {
		if v == version {
			return true
		}
	}
	return false
}
