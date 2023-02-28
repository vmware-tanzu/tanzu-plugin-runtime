// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

type GetContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

type GetContextOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	*ContextOpts                                    // For specific version options look into ContextOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

type SetContextInputOptions struct {
	*core.RuntimeAPIVersion      // required
	*ContextOpts                 // required
	SetCurrentContext       bool // required
}

type SetContextOutputOptions struct {
	ValidationStrategy core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error              string                  // expected error message could be the sub string of actual error message
}

type DeleteContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

type DeleteContextOutputOptions struct {
	ValidationStrategy core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error              string                  // expected error message could be the sub string of actual error message
}

type SetCurrentContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

type SetCurrentContextOutputOptions struct {
	*core.RuntimeAPIVersion        // required
	Error                   string // expected error message could be the sub string of actual error message
}

type GetCurrentContextInputOptions struct {
	*core.RuntimeAPIVersion             // required
	Target                  Target      // required for v1.0.0 - v0.28.0
	ContextType             ContextType // required for v0.25.4
}

type GetCurrentContextOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	*ContextOpts                                    // For specific version options look into ContextOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}
type RemoveCurrentContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	Target                  Target // required
}

type RemoveCurrentContextOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

func (s *ContextOpts) ShouldNotIncludeTarget() bool {
	return s.Target == ""
}

func (s *ContextOpts) ShouldNotIncludeContextType() bool {
	return s.Type == ""
}

func (s *GetCurrentContextInputOptions) ShouldNotIncludeTarget() bool {
	return s.Target == ""
}

func (s *GetCurrentContextInputOptions) ShouldNotIncludeContextType() bool {
	return s.ContextType == ""
}

func (s *SetContextInputOptions) ValidName() bool {
	return s.Name != ""
}

func (s *SetContextInputOptions) ValidTarget() bool {
	return s.Target != "" && (s.Target == TargetK8s || s.Target == TargetTMC)
}

func (s *SetContextInputOptions) ValidContextType() bool {
	return s.Type != "" && (s.Type == CtxTypeK8s || s.Type == CtxTypeTMC)
}

func (s *SetContextInputOptions) ValidGlobalOptsOrClusterOpts() bool {
	return (s.GlobalOpts != nil && s.GlobalOpts.Endpoint != "") || (s.ClusterOpts != nil && s.ClusterOpts.Endpoint != "")
}

func (s *SetContextInputOptions) ValidDiscoverySources() bool {
	return s.DiscoverySources != nil || len(s.DiscoverySources) == 0
}

// ValidateContextOutputOptionsAsPerRuntimeVersion validate the getContextOutputOptions as per runtime version i.e. check whether the expected fields are supported for the runtime version specified
func (s *ContextOpts) ValidateContextOutputOptionsAsPerRuntimeVersion(version core.RuntimeVersion) (bool, error) {
	var valid bool
	switch version {
	case core.Version100, core.Version0280:
		valid = s.ShouldNotIncludeContextType()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid get context output options for the specified runtime version contextType is not supported %v", version)
	case core.Version0254:
		valid = s.ShouldNotIncludeTarget()
		if valid {
			return valid, nil
		}
		return valid, fmt.Errorf("invalid get context output options for the specified runtime version Target is not supported %v", version)

	default:
		return false, errors.New("GetContext API is not supported for the specified runtime version")
	}
}
