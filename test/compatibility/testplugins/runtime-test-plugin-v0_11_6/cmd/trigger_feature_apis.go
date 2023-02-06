// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerIsFeatureActivatedAPI trigger Runtime IsFeatureActivated API
func triggerIsFeatureActivatedAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime IsFeatureActivated API
	feature, err := core.ParseStr(api.Arguments[core.Feature])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.Plugin, err),
		}
	}

	// Trigger IsFeatureActivated API
	return isFeatureActivated(feature)
}

func isFeatureActivated(feature string) *core.APIResponse {
	enabled := configlib.IsFeatureActivated(feature)

	return &core.APIResponse{
		ResponseType: core.BooleanResponse,
		ResponseBody: enabled,
	}
}
