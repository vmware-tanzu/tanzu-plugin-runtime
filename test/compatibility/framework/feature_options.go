// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// IsFeatureEnabledInputOptions used to generate GetFeature command
type IsFeatureEnabledInputOptions struct {
	*core.RuntimeAPIVersion        // required
	PluginName              string // required
	KeyName                 string // required
}

// IsFeatureEnabledOutputOptions used to generate GetFeature command
type IsFeatureEnabledOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	FeatureEnabled          bool                    // For specific version options look into FeatureOpts definition
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// SetFeatureInputOptions used to generate SetFeature command
type SetFeatureInputOptions struct {
	*core.RuntimeAPIVersion        // required
	PluginName              string // required
	KeyName                 string // required
	ValueName               string // required
}

// SetFeatureOutputOptions used to generate SetFeature command
type SetFeatureOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}

// DeleteFeatureInputOptions used to generate DeleteFeature command
type DeleteFeatureInputOptions struct {
	*core.RuntimeAPIVersion        // required
	PluginName              string // required
	KeyName                 string // required

}

// DeleteFeatureOutputOptions used to generate DeleteFeature command
type DeleteFeatureOutputOptions struct {
	*core.RuntimeAPIVersion                         // required
	ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
	Error                   string                  // expected error message could be the sub string of actual error message
}
