// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetEnvConfigurationsAPI trigger Runtime GetEnvConfigurations API
func triggerGetEnvConfigurationsAPI(_ *core.API) *core.APIResponse {
	// Trigger GetEnvConfigurations API
	return getEnvConfigurations()
}

func getEnvConfigurations() *core.APIResponse {
	envs := configlib.GetEnvConfigurations()

	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: envs,
	}
}
