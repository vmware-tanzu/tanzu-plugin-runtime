// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// Validate  the setServerInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
func (opts *SetServerInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		if !opts.ValidName() {
			return false, fmt.Errorf("invalid 'name' for set server input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if !opts.ValidServerType() {
			return false, fmt.Errorf("invalid 'type' for set server input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if !opts.ValidGlobalOptsOrManagementClusterOpts() {
			return false, fmt.Errorf("invalid 'global or clusterOpts' for set server input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, errors.New("SetServer API is not supported for the specified runtime version")
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetServerInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	if opts.ServerName == "" {
		return false, errors.New("server name is required")
	}
	return true, nil
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetServerOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *DeleteServerInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	if opts.ServerName == "" {
		return false, errors.New("server name is required")
	}
	return true, nil
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetCurrentServerInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *SetCurrentServerInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	if opts.ServerName == "" {
		return false, errors.New("server name is required")
	}
	return true, nil
}

// Validate the getServerOutputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetCurrentServerOutputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	return true, nil
}
