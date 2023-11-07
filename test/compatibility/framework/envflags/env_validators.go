// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package envflags

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// Validate the SetEnvInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
func (opts *SetEnvInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for SetEnvInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		if opts.Value == "" {
			return false, fmt.Errorf("invalid 'value' for SetEnvInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("SetEnv API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the GetEnvInputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetEnvInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for GetEnvInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetEnv API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the GetEnvConfigurationsOutputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetEnvConfigurationsOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		if opts.Envs == nil {
			return false, fmt.Errorf("invalid 'envs' for GetEnvConfigurationsOutputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetEnvConfigurations API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the DeleteEnvInputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *DeleteEnvInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for DeleteEnvInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("DeleteEnv API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}
