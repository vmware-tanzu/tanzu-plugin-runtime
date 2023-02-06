// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// GetClientConfigInputOptions used to generate GetClientConfig command
type GetClientConfigInputOptions struct {
	*core.RuntimeAPIVersion // required
}

// GetClientConfigOutputOptions used to generate GetClientConfig command
type GetClientConfigOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ClientConfigOpts        *ClientConfigOpts       // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// StoreClientConfigInputOptions used to generate StoreClientConfig command
type StoreClientConfigInputOptions struct {
	*core.RuntimeAPIVersion                   // required
	ClientConfigOpts        *ClientConfigOpts // required

}

// StoreClientConfigOutputOptions used to generate StoreClientConfig command
type StoreClientConfigOutputOptions struct {
	ValidationStrategy core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error              string                  // expected error message could be the sub string of actual error message
}
