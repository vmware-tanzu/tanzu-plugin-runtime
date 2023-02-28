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

// NewDeleteContextCommand deletes a context command object
// Creates the context specific command based on runtimeVersion passed in inputOptions also validates if the required input and outputOptions are passed
func NewDeleteContextCommand(deleteContextInputOptions *DeleteContextInputOptions, deleteContextOutputOptions *DeleteContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}
	// Init the API object
	api := &core.API{}
	api.Name = core.DeleteContextAPIName
	api.Version = deleteContextInputOptions.RuntimeVersion

	// Validate the Input Options
	if deleteContextInputOptions.ContextName == "" {
		return nil, errors.New("context name is required")
	}

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ContextName: deleteContextInputOptions.ContextName,
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if deleteContextOutputOptions != nil && deleteContextOutputOptions.Error != "" {
		res = core.Failed
		content = deleteContextOutputOptions.Error
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

func NewSetCurrentContextCommand(setCurrentContextInputOptions *SetCurrentContextInputOptions, setCurrentContextOutputOptions *SetCurrentContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}
	api.Name = core.SetCurrentContextAPIName
	api.Version = setCurrentContextInputOptions.RuntimeVersion

	// Validate the Input Options
	if setCurrentContextInputOptions.ContextName == "" {
		return nil, errors.New("context name is required")
	}

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ContextName: setCurrentContextInputOptions.ContextName,
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if setCurrentContextOutputOptions != nil && setCurrentContextOutputOptions.Error != "" {
		res = core.Failed
		content = setCurrentContextOutputOptions.Error
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

// NewGetCurrentContextCommand creates a get context command object from inputOptions and outputOptions
// Creates the context specific command based on runtimeVersion passed in inputOptions also validates if the required input and outputOptions are passed
func NewGetCurrentContextCommand(getCurrentContextInputOptions *GetCurrentContextInputOptions, getCurrentContextOutputOptions *GetCurrentContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}
	// Init the API object
	api := &core.API{}
	api.Name = core.GetCurrentContextAPIName

	// Run Core Validators
	_, err := core.ValidateRuntimeVersion(getCurrentContextInputOptions.RuntimeAPIVersion)
	if err != nil {
		return nil, err
	}

	api.Version = getCurrentContextInputOptions.RuntimeVersion

	// Validate the Input Options
	_, err = ValidateGetCurrentContextOutputOptionsAsPerRuntimeVersion(getCurrentContextInputOptions)
	if err != nil {
		return nil, err
	}

	// Construct the context api arguments and output
	api.Arguments = make(map[core.APIArgumentType]interface{})

	if getCurrentContextInputOptions.Target != "" {
		api.Arguments[core.Target] = getCurrentContextInputOptions.Target
	} else if getCurrentContextInputOptions.ContextType != "" {
		api.Arguments[core.ContextType] = getCurrentContextInputOptions.ContextType
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if getCurrentContextOutputOptions.ContextOpts != nil {
		// Run Core Validators
		_, err = core.ValidateRuntimeVersion(getCurrentContextOutputOptions.RuntimeAPIVersion)
		if err != nil {
			return nil, err
		}

		// Validate the Output Options
		_, err := getCurrentContextOutputOptions.ContextOpts.ValidateContextOutputOptionsAsPerRuntimeVersion(getCurrentContextOutputOptions.RuntimeVersion)
		if err != nil {
			return nil, err
		}

		// Construct get context output context opts
		bytes, err := yaml.Marshal(getCurrentContextOutputOptions.ContextOpts)
		if err != nil {
			return nil, err
		}

		content = string(bytes)
		res = core.Success
	} else if getCurrentContextOutputOptions.Error != "" {
		res = core.Failed
		content = getCurrentContextOutputOptions.Error
	}

	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	if getCurrentContextOutputOptions.ValidationStrategy != "" {
		api.Output.ValidationStrategy = getCurrentContextOutputOptions.ValidationStrategy
	}

	c.APIs = append(c.APIs, api)
	return c, nil
}

//nolint: dupl
func NewRemoveCurrentContextCommand(removeCurrentContextInputOptions *RemoveCurrentContextInputOptions, removeCurrentContextOutputOptions *RemoveCurrentContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}
	api.Name = core.RemoveCurrentContextAPIName
	api.Version = removeCurrentContextInputOptions.RuntimeVersion

	// Validate the Input Options
	if removeCurrentContextInputOptions.Target == "" {
		return nil, errors.New("context target is required")
	}

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Target: removeCurrentContextInputOptions.Target,
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if removeCurrentContextOutputOptions != nil && removeCurrentContextOutputOptions.Error != "" {
		res = core.Failed
		content = removeCurrentContextOutputOptions.Error
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
