// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"gopkg.in/yaml.v3"

	configlib "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerGetContextAPI trigger get context runtime api
func triggerGetContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ContextName, err.Error()),
		}
		return log
	}

	// Call runtime GetContext API
	ctx, err := configlib.GetContext(ctxName)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	if ctx != nil {
		log.APIResponse = &core.APIResponse{
			ResponseBody: ctx,
			ResponseType: core.MapResponse,
		}
	}
	return log
}

// triggerSetContextAPI trigger set context runtime api
func triggerSetContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	ctx, err := parseContext(api.Arguments[core.Context].(string))
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse context from argument %v with error %v ", core.Context, err.Error()),
		}
		return log
	}
	setCurrent := api.Arguments[core.SetCurrent].(bool)

	// Call runtime SetContext API
	err = configlib.SetContext(ctx, setCurrent)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &core.APIResponse{
			ResponseBody: "",
			ResponseType: core.StringResponse,
		}
	}

	return log
}

// triggerDeleteContextAPI trigger remove context runtime api
func triggerDeleteContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ContextName, err.Error()),
		}
		return log
	}

	// Call runtime DeleteContext API
	err = configlib.DeleteContext(ctxName)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &core.APIResponse{
			ResponseBody: "",
			ResponseType: core.StringResponse,
		}
	}

	return log
}

// triggerSetCurrentContextAPI trigger remove context runtime api
func triggerSetCurrentContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	ctxName, err := core.ParseStr(api.Arguments[core.ContextName])
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.ContextName, err.Error()),
		}
		return log
	}

	// Call runtime AddContext API
	err = configlib.SetCurrentContext(ctxName)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &core.APIResponse{
			ResponseBody: "",
			ResponseType: core.StringResponse,
		}
	}

	return log
}

// triggerGetCurrentContextAPI trigger remove context runtime api
func triggerGetCurrentContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	target, err := core.ParseStr(api.Arguments[core.Target])
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.Target, err.Error()),
		}
		return log
	}

	// Call runtime AddContext API
	ctx, err := configlib.GetCurrentContext(configtypes.Target(target))

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	if ctx != nil {
		log.APIResponse = &core.APIResponse{
			ResponseBody: ctx,
			ResponseType: core.MapResponse,
		}
	}
	return log
}

// triggerRemoveCurrentContextAPI trigger remove context runtime api
func triggerRemoveCurrentContextAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	target, err := core.ParseStr(api.Arguments[core.Target])
	if err != nil {
		log := core.APILog{}
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: fmt.Errorf("failed to parse string from argument %v with error %v ", core.Target, err.Error()),
		}
		return log
	}

	// Call runtime AddContext API
	err = configlib.RemoveCurrentContext(configtypes.Target(target))

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &core.APIResponse{
			ResponseBody: "",
			ResponseType: core.StringResponse,
		}
	}
	return log
}

// parseContext unmarshalls string to Context struct
func parseContext(context string) (*configtypes.Context, error) {
	var ctx configtypes.Context
	err := yaml.Unmarshal([]byte(context), &ctx)
	if err != nil {
		return nil, err
	}
	return &ctx, nil
}
