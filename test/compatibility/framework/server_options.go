// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

// GetServerInputOptions used to generate GetServer command
type GetServerInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ServerName              string // required
}

// GetServerOutputOptions used to generate GetServer command
type GetServerOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	*types.ServerOpts                               // For specific version options look into ServerOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetServerInputOptions used to generate SetServer command
type SetServerInputOptions struct {
	*core.RuntimeAPIVersion      // required
	*types.ServerOpts            // required
	SetCurrentServer        bool // required
}

// SetServerOutputOptions used to generate SetServer command
type SetServerOutputOptions struct {
	ValidationStrategy core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error              string                  // expected error message could be the sub string of actual error message
}

// DeleteServerInputOptions used to generate DeleteServer command
type DeleteServerInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ServerName              string // required
}

// DeleteServerOutputOptions used to generate DeleteServer command
type DeleteServerOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetCurrentServerInputOptions used to generate SetCurrentServer command
type SetCurrentServerInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ServerName              string // required
}

// SetCurrentServerOutputOptions used to generate SetCurrentServer command
type SetCurrentServerOutputOptions struct {
	*core.RuntimeAPIVersion        // required
	Error                   string // expected error message could be the sub string of actual error message
}

// GetCurrentServerInputOptions used to generate GetCurrentServer command
type GetCurrentServerInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetCurrentServerOutputOptions used to generate GetCurrentServer command
type GetCurrentServerOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	*types.ServerOpts                               // For specific version options look into ServerOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// RemoveCurrentServerInputOptions used to generate RemoveCurrentServer command
type RemoveCurrentServerInputOptions struct {
	*core.RuntimeAPIVersion        // required
	ServerName              string // required for v1.0.0 - v0.28.0
}

// RemoveCurrentServerOutputOptions used to generate RemoveCurrentServer command
type RemoveCurrentServerOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}
