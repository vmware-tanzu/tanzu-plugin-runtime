// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewSetContextCommand constructs a command to make a call to specific runtime version SetContext API
// Input Parameter setContextInputOptions has all input parameters which are required for Runtime SetContextAPIName API
// Input Parameter: setContextOutputOptions has details about expected output from Runtime SetContextAPIName API call
// Return: command to execute or error if any validations fails for SetContextInputOptions or SetContextOutputOptions
// This method does validate the input parameters  SetContextInputOptions/SetContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetContextInputOptions/SetContextOutputOptions definition(and CtxOptions struct, which is embedded)
func NewSetContextCommand(setContextInputOptions *SetContextInputOptions, setContextOutputOptions *SetContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}
	api.Name = core.SetContextAPIName

	// Run Core Validators
	_, err := core.ValidateRuntimeVersion(setContextInputOptions.RuntimeAPIVersion)
	if err != nil {
		return nil, err
	}

	api.Version = setContextInputOptions.RuntimeVersion

	// Validate the SetContext Input Options
	_, err = ValidateSetContextInputOptionsAsPerRuntimeVersion(setContextInputOptions)
	if err != nil {
		return nil, err
	}

	// Construct the set context api arguments and output
	bytes, err := yaml.Marshal(setContextInputOptions.ContextOpts)
	if err != nil {
		return nil, err
	}

	// Construct the setCurrent Argument
	var setCurrent bool
	if setContextInputOptions.SetCurrentContext {
		setCurrent = true
	} else {
		setCurrent = false
	}

	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Context:    string(bytes),
		core.SetCurrent: setCurrent,
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if setContextOutputOptions != nil && setContextOutputOptions.Error != "" {
		res = core.Failed
		content = setContextOutputOptions.Error
	} else {
		res = core.Success
		content = ""
	}
	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	c.APIs = append(c.APIs, api)

	return c, nil
}

// NewGetContextCommand creates a get context command object from inputOptions and outputOptions
// Creates the context specific command based on runtimeVersion passed in inputOptions also validates if the required input and outputOptions are passed
func NewGetContextCommand(getContextInputOptions *GetContextInputOptions, getContextOutputOptions *GetContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}
	// Init the API object
	api := &core.API{}
	api.Name = core.GetContextAPIName

	// Run Core Validators
	_, err := core.ValidateRuntimeVersion(getContextInputOptions.RuntimeAPIVersion)
	if err != nil {
		return nil, err
	}

	api.Version = getContextInputOptions.RuntimeVersion

	// Validate the Input Options
	if getContextInputOptions.ContextName == "" {
		return nil, errors.New("context name is required")
	}

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ContextName: getContextInputOptions.ContextName,
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if getContextOutputOptions.ContextOpts != nil {
		// Run Core Validators
		_, err = core.ValidateRuntimeVersion(getContextOutputOptions.RuntimeAPIVersion)
		if err != nil {
			return nil, err
		}

		// Validate the Output Options
		_, err := ValidateGetContextOutputOptionsAsPerRuntimeVersion(getContextOutputOptions)
		if err != nil {
			return nil, err
		}

		// Construct get context output context opts
		bytes, err := yaml.Marshal(getContextOutputOptions.ContextOpts)
		if err != nil {
			return nil, err
		}

		content = string(bytes)
		res = core.Success
	} else if getContextOutputOptions.Error != "" {
		res = core.Failed
		content = getContextOutputOptions.Error
	}

	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	if getContextOutputOptions.ValidationStrategy != "" {
		api.Output.ValidationStrategy = getContextOutputOptions.ValidationStrategy
	}

	c.APIs = append(c.APIs, api)
	return c, nil
}
