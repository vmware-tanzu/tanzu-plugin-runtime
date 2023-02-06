// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configapi "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetContextAPI trigger Runtime GetContext API
func triggerGetContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime GetContext API
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v", core.ContextName, err),
		}
	}
	// Trigger GetContext API
	return getContext(ctxName)
}

// triggerAddContextAPI trigger Runtime AddContext API
func triggerAddContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime AddContext API
	ctx, err := parseContext(api.Arguments[core.Context].(string))
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse context from argument %v with error %v ", core.Context, err.Error()),
		}
	}
	setCurrent := api.Arguments[core.SetCurrent].(bool)
	// Trigger AddContext API
	return addContext(ctx, setCurrent)
}

// triggerRemoveContextAPI trigger Runtime RemoveContext API
func triggerRemoveContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime RemoveContext API
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ContextName, err.Error()),
		}
	}
	// Trigger RemoveContext API
	return removeContext(ctxName)
}

// triggerSetCurrentContextAPI trigger Runtime SetCurrentContext API
func triggerSetCurrentContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime SetCurrentContext API
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ContextName, err.Error()),
		}
	}
	// Trigger SetCurrentContext API
	return setCurrentContext(ctxName)
}

// triggerGetCurrentContextAPI trigger Runtime GetCurrentContext API
func triggerGetCurrentContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime GetCurrentContext API
	contextType, err := core.ParseStr(api.Arguments[core.ContextType])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ContextType, err.Error()),
		}
	}
	// Trigger GetCurrentContext API
	return getCurrentContext(configapi.ContextType(contextType))
}

func getContext(ctxName string) *core.APIResponse {
	ctx, err := configlib.GetContext(ctxName)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if ctx == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("context %s not found", ctxName),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: ctx,
	}
}

func addContext(context *configapi.Context, setCurrent bool) *core.APIResponse {
	err := configlib.AddContext(context, setCurrent)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	return &core.APIResponse{
		ResponseBody: "",
		ResponseType: core.StringResponse,
	}
}

func removeContext(contextName string) *core.APIResponse {
	err := configlib.RemoveContext(contextName)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	return &core.APIResponse{
		ResponseBody: "",
		ResponseType: core.StringResponse,
	}
}

func setCurrentContext(contextName string) *core.APIResponse {
	err := configlib.SetCurrentContext(contextName)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	return &core.APIResponse{
		ResponseBody: "",
		ResponseType: core.StringResponse,
	}
}

func getCurrentContext(contextType configapi.ContextType) *core.APIResponse {
	ctx, err := configlib.GetCurrentContext(contextType)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if ctx == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("context %s not found", contextType),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: ctx,
	}
}
