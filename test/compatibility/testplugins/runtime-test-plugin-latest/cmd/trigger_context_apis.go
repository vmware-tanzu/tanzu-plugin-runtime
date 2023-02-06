// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
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

// triggerSetContextAPI trigger Runtime SetContext API
func triggerSetContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime SetContext API
	ctx, err := parseContext(api.Arguments[core.Context].(string))
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse context from argument %v with error %v ", core.Context, err.Error()),
		}
	}
	setCurrent := api.Arguments[core.SetCurrent].(bool)
	// Trigger SetContext API
	return setContext(ctx, setCurrent)
}

// triggerDeleteContextAPI trigger Runtime DeleteContext API
func triggerDeleteContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime DeleteContext API
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ContextName, err.Error()),
		}
	}
	// Trigger DeleteContext API
	return deleteContext(ctxName)
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
	target, err := core.ParseStr(api.Arguments[core.Target])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.Target, err.Error()),
		}
	}
	// Trigger GetCurrentContext API
	return getCurrentContext(configtypes.Target(target))
}

// triggerRemoveCurrentContextAPI trigger Runtime RemoveCurrentContext API
func triggerRemoveCurrentContextAPI(api *core.API) *core.APIResponse {
	// Parse arguments needed to trigger the Runtime RemoveCurrentContext API
	target, err := core.ParseStr(api.Arguments[core.Target])
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.Target, err.Error()),
		}
	}
	// Trigger RemoveCurrentContext
	return removeCurrentContext(configtypes.Target(target))
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

func setContext(context *configtypes.Context, setCurrent bool) *core.APIResponse {
	err := configlib.SetContext(context, setCurrent)
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

func deleteContext(contextName string) *core.APIResponse {
	err := configlib.DeleteContext(contextName)
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

func getCurrentContext(target configtypes.Target) *core.APIResponse {
	ctx, err := configlib.GetCurrentContext(target)
	if err != nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}
	if ctx == nil {
		return &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("context %s not found", target),
		}
	}
	return &core.APIResponse{
		ResponseType: core.MapResponse,
		ResponseBody: ctx,
	}
}

func removeCurrentContext(target configtypes.Target) *core.APIResponse {
	err := configlib.RemoveCurrentContext(target)
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
