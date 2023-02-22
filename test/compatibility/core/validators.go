// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"fmt"
)

func ValidateRuntimeVersion(version *RuntimeAPIVersion) (bool, error) {
	if version == nil || version.RuntimeVersion == "" {
		return false, fmt.Errorf("runtime version is mandatory")
	}

	if !isRuntimeVersionSupported(version.RuntimeVersion) {
		return false, fmt.Errorf("runtime version %v is not supported", version)
	}
	return true, nil
}

func isRuntimeVersionSupported(version RuntimeVersion) bool {
	for _, v := range SupportedRuntimeVersions {
		if v == version {
			return true
		}
	}
	return false
}
