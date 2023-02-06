// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func (opts *ContextOpts) ShouldNotIncludeTarget() bool {
	return opts.Target == ""
}

func (opts *ContextOpts) ShouldNotIncludeContextType() bool {
	return opts.Type == ""
}

func (opts *GetCurrentContextInputOptions) ShouldNotIncludeTarget() bool {
	return opts.Target == ""
}

func (opts *GetCurrentContextInputOptions) ShouldNotIncludeContextType() bool {
	return opts.ContextType == ""
}

func (opts *RemoveCurrentContextInputOptions) ShouldNotIncludeTarget() bool {
	return opts.Target == ""
}

func (opts *ContextOpts) ValidName() bool {
	return opts.Name != ""
}

func (opts *ContextOpts) ValidTarget() bool {
	return opts.Target != "" && (opts.Target == TargetK8s || opts.Target == TargetTMC)
}

func (opts *ContextOpts) ValidContextType() bool {
	return opts.Type != "" && (opts.Type == CtxTypeK8s || opts.Type == CtxTypeTMC)
}

func (opts *ContextOpts) ValidGlobalOptsOrClusterOpts() bool {
	return (opts.GlobalOpts != nil && opts.GlobalOpts.Endpoint != "") || (opts.ClusterOpts != nil && opts.ClusterOpts.Endpoint != "")
}

func (opts *ContextOpts) ValidDiscoverySources() bool {
	return opts.DiscoverySources != nil || len(opts.DiscoverySources) == 0
}

// Validate  the setContextInputOptions as per runtime version i.e. check whether mandatory fields are set and throw error if missing
func (opts *SetContextInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version0280:
		if !opts.ValidName() {
			return false, fmt.Errorf("invalid 'name' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if !opts.ValidTarget() {
			return false, fmt.Errorf("invalid 'target' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if !opts.ValidGlobalOptsOrClusterOpts() {
			return false, fmt.Errorf("invalid 'global or clusterOpts' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	case core.Version0254:
		if !opts.ValidName() {
			return false, fmt.Errorf("invalid 'Name' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if !opts.ValidContextType() {
			return false, fmt.Errorf("invalid 'ContextType' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		if !opts.ValidGlobalOptsOrClusterOpts() {
			return false, fmt.Errorf("invalid 'GlobalOpts or ClusterOpts' for set context input options for the specified runtime version %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, errors.New("SetContext API is not supported for the specified runtime version")
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetContextInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	if opts.ContextName == "" {
		return false, errors.New("context name is required")
	}
	return true, nil
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetContextOutputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version0280:
		if !opts.ShouldNotIncludeContextType() {
			return false, fmt.Errorf("invalid get context output options for the specified runtime version contextType is not supported %v", opts.RuntimeVersion)
		}
		return true, nil
	case core.Version0254:
		if !opts.ShouldNotIncludeTarget() {
			return false, fmt.Errorf("invalid get context output options for the specified runtime version Target is not supported %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, errors.New("GetContext API is not supported for the specified runtime version")
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *DeleteContextInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}
	if opts.ContextName == "" {
		return false, errors.New("context name is required")
	}
	return true, nil
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetCurrentContextInputOptions) Validate() (bool, error) {
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version0280:
		if !opts.ShouldNotIncludeContextType() {
			return false, fmt.Errorf("invalid get current context input options for the specified runtime version contextType is not supported %v", opts.RuntimeVersion)
		}
		return true, nil
	case core.Version0254:
		if !opts.ShouldNotIncludeTarget() {
			return false, fmt.Errorf("invalid get current context input options for the specified runtime version Target is not supported %v", opts.RuntimeVersion)
		}
		return true, nil
	default:
		return false, errors.New("GetCurrentContext API is not supported for the specified runtime version")
	}
}

// Validate the opts as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *SetCurrentContextInputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	if opts.ContextName == "" {
		return false, errors.New("context name is required")
	}
	return true, nil
}

// Validate the getContextOutputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (opts *GetCurrentContextOutputOptions) Validate() (bool, error) {
	// Run Core Validators
	_, err := opts.RuntimeAPIVersion.Validate()
	if err != nil {
		return false, err
	}

	var valid bool
	switch opts.RuntimeVersion {
	case core.VersionLatest, core.Version0280:
		valid = opts.ContextOpts.ShouldNotIncludeContextType()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid get context output options for the specified runtime version contextType is not supported %v", opts.RuntimeVersion)
	case core.Version0254:
		valid = opts.ContextOpts.ShouldNotIncludeTarget()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid get context output options for the specified runtime version Target is not supported %v", opts.RuntimeVersion)

	default:
		return false, errors.New("GetCurrentContext API is not supported for the specified runtime version")
	}
}
