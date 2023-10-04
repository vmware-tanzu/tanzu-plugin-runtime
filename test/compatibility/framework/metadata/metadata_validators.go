// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// Validate  the setEnvInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
func (opts *SetConfigMetadataPatchStrategyInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for SetConfigMetadataPatchStrategyInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		if opts.Value == "" {
			return false, fmt.Errorf("invalid 'value' for SetConfigMetadataPatchStrategyInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("SetConfigMetadataPatchStrategyInputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate  the setEnvInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
func (opts *SetConfigMetadataSettingInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for SetConfigMetadataSettingInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		if opts.Value == "" {
			return false, fmt.Errorf("invalid 'value' for SetConfigMetadataSettingInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("SetConfigMetadataSettingInputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetMetadataOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.MetadataOpts != nil && opts.MetadataOpts.ConfigMetadata == nil && opts.MetadataOpts.ConfigMetadata.Settings == nil && opts.MetadataOpts.ConfigMetadata.PatchStrategy == nil {
			return false, fmt.Errorf("invalid 'key' for GetMetadataOutputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetMetadataOutputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetConfigMetadataOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.ConfigMetadataOpts != nil && opts.ConfigMetadataOpts.Settings == nil && opts.ConfigMetadataOpts.PatchStrategy == nil {
			return false, fmt.Errorf("invalid 'key' for GetConfigMetadataOutputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetConfigMetadataOutputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetConfigMetadataPatchStrategyOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.PatchStrategy == nil {
			return false, fmt.Errorf("invalid 'key' for GetConfigMetadataPatchStrategyOutputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetConfigMetadataPatchStrategyOutputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetConfigMetadataSettingsOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.MetadataSettings == nil {
			return false, fmt.Errorf("invalid 'key' for GetConfigMetadataSettingsOutputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetConfigMetadataSettingsOutputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetConfigMetadataSettingInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for GetConfigMetadataSettingInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetConfigMetadataSettingInputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetConfigMetadataSettingOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Value == "" {
			return false, fmt.Errorf("invalid 'key' for GetConfigMetadataSettingOutputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("GetConfigMetadataSettingOutputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *IsConfigMetadataSettingsEnabledInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for IsConfigMetadataSettingsEnabledInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("IsConfigMetadataSettingsEnabledInputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *DeleteConfigMetadataSettingInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if opts.Key == "" {
			return false, fmt.Errorf("invalid 'key' for DeleteConfigMetadataSettingInputOptions for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, fmt.Errorf("DeleteConfigMetadataSettingInputOptions API is not supported for the specified runtime version %v", opts.RuntimeVersion)
	}
}
