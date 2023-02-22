// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// ValidateSetContextInputOptionsAsPerRuntimeVersion validate the setContextInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
func ValidateSetContextInputOptionsAsPerRuntimeVersion(setContextInputOptions *SetContextInputOptions) (bool, error) {
	var valid bool
	switch setContextInputOptions.RuntimeVersion {
	case core.Version100, core.Version0280:
		valid = setContextInputOptions.ValidName() && setContextInputOptions.ValidTarget() && setContextInputOptions.ValidGlobalOptsOrClusterOpts()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)

	case core.Version0254:
		valid = setContextInputOptions.ValidName() && setContextInputOptions.ValidContextType() && setContextInputOptions.ValidGlobalOptsOrClusterOpts()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)
	default:
		return false, errors.New("SetContext API is not supported for the specified runtime version")
	}
}

// ValidateGetContextOutputOptionsAsPerRuntimeVersion validate the getContextOutputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func ValidateGetContextOutputOptionsAsPerRuntimeVersion(getContextOutputOptions *GetContextOutputOptions) (bool, error) {
	var valid bool
	switch getContextOutputOptions.RuntimeVersion {
	case core.Version100, core.Version0280:
		valid = getContextOutputOptions.ShouldNotIncludeContextType()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid get context output options for the specified runtime version contextType is not supported %v", getContextOutputOptions.RuntimeVersion)
	case core.Version0254:
		valid = getContextOutputOptions.ShouldNotIncludeTarget()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid get context output options for the specified runtime version Target is not supported %v", getContextOutputOptions.RuntimeVersion)

	default:
		return false, errors.New("GetContext API is not supported for the specified runtime version")
	}
}
