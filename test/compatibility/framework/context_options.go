// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

type GetContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

type GetContextOutputOptions struct {
	*core.RuntimeAPIVersion                        // required
	*ContextOpts                                   // For specific version options look into ContextOpts definition
	ValidationMatcher       core.ValidationMatcher // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                 // expected error message could be the sub string of actual error message
}

type SetContextInputOptions struct {
	*core.RuntimeAPIVersion      // required
	*ContextOpts                 // required
	IsCurrentContext        bool // required
}

type SetContextOutputOptions struct {
	Error string // expected error message could be the sub string of actual error message
}

type DeleteContextInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ContextName             string // required
}

type DeleteContextOutputOptions struct {
	Error string // expected error message could be the sub string of actual error message
}

func (s *GetContextOutputOptions) ShouldNotIncludeTarget() bool {
	return s.Target == ""
}

func (s *GetContextOutputOptions) ShouldNotIncludeContextType() bool {
	return s.Type == ""
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
