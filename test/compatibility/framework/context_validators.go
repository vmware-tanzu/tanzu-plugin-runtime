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
	switch setContextInputOptions.RuntimeVersion {
	case core.Version100, core.Version0280:
		if !setContextInputOptions.ValidName() {
			return false, fmt.Errorf("invalid 'name' for set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)
		}
		if !setContextInputOptions.ValidTarget() {
			return false, fmt.Errorf("invalid 'target' for set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)
		}
		if !setContextInputOptions.ValidGlobalOptsOrClusterOpts() {
			return false, fmt.Errorf("invalid 'global or clusterOpts' for set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)
		}
		return true, nil
	case core.Version0254:
		if !setContextInputOptions.ValidName() {
			return false, fmt.Errorf("invalid 'Name' for set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)
		}
		if !setContextInputOptions.ValidContextType() {
			return false, fmt.Errorf("invalid 'ContextType' for set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)
		}
		if !setContextInputOptions.ValidGlobalOptsOrClusterOpts() {
			return false, fmt.Errorf("invalid 'GlobalOpts or ClusterOpts' for set context input options for the specified runtime version %v", setContextInputOptions.RuntimeVersion)
		}
		return true, nil
	default:
		return false, errors.New("SetContext API is not supported for the specified runtime version")
	}
}

// ValidateGetContextOutputOptionsAsPerRuntimeVersion validate the getContextOutputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func ValidateGetContextOutputOptionsAsPerRuntimeVersion(getContextOutputOptions *GetContextOutputOptions) (bool, error) {
	switch getContextOutputOptions.RuntimeVersion {
	case core.Version100, core.Version0280:
		if !getContextOutputOptions.ShouldNotIncludeContextType() {
			return false, fmt.Errorf("invalid get context output options for the specified runtime version contextType is not supported %v", getContextOutputOptions.RuntimeVersion)
		}
		return true, nil
	case core.Version0254:
		if !getContextOutputOptions.ShouldNotIncludeTarget() {
			return false, fmt.Errorf("invalid get context output options for the specified runtime version Target is not supported %v", getContextOutputOptions.RuntimeVersion)
		}
		return true, nil
	default:
		return false, errors.New("GetContext API is not supported for the specified runtime version")
	}
}
