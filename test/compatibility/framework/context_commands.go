// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewSetContextCommand constructs a command to make a call to specific runtime version SetContext API
// Input Parameter inputOpts has all input parameters which are required for Runtime SetContext API
// Input Parameter: outputOpts has details about expected output from Runtime SetContext API call
// Return: command to execute or error if any validations fails for SetContextInputOptions or SetContextOutputOptions
// This method does validate the input parameters  SetContextInputOptions or SetContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetContextInputOptions or SetContextOutputOptions definition(and ContextOpts struct, which is embedded)
// nolint:dupl
func NewSetContextCommand(inputOpts *SetContextInputOptions, outputOpts *SetContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.SetContextAPIName

	// Validate the SetContext input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetContext API arguments
	bytes, err := yaml.Marshal(inputOpts.ContextOpts)
	if err != nil {
		return nil, err
	}

	// Construct the setCurrent Argument
	var setCurrent = false
	if inputOpts.SetCurrentContext {
		setCurrent = true
	}

	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Context:    string(bytes),
		core.SetCurrent: setCurrent,
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts != nil && outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	}

	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	c.APIs = append(c.APIs, api)

	return c, nil
}

// NewGetContextCommand constructs a command to make a call to specific runtime version GetContext API
// Input Parameter inputOpts has all input parameters which are required for Runtime GetContext API
// Input Parameter: outputOpts has details about expected output from Runtime GetContext API call
// Return: command to execute or error if any validations fails for GetContextInputOptions or GetContextOutputOptions
// This method does validate the input parameters  GetContextInputOptions or GetContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetContextInputOptions or GetContextOutputOptions definition(and ContextOpts struct, which is embedded)
func NewGetContextCommand(inputOpts *GetContextInputOptions, outputOpts *GetContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetContextAPIName

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the GetContext API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ContextName: inputOpts.ContextName,
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.ContextOpts != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		// Construct get context output context opts
		bytes, err := yaml.Marshal(outputOpts.ContextOpts)
		if err != nil {
			return nil, err
		}

		content = string(bytes)
		res = core.Success
	} else if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	}

	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	if outputOpts.ValidationStrategy != "" {
		api.Output.ValidationStrategy = outputOpts.ValidationStrategy
	}

	c.APIs = append(c.APIs, api)
	return c, nil
}

// NewDeleteContextCommand constructs a command to make a call to specific runtime version DeleteContext API
// Input Parameter inputOpts has all input parameters which are required for Runtime DeleteContext API
// Input Parameter: outputOpts has details about expected output from Runtime DeleteContext API call
// Return: command to execute or error if any validations fails for DeleteContextInputOptions or DeleteContextOutputOptions
// This method does validate the input parameters  DeleteContextInputOptions or DeleteContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to DeleteContextInputOptions or DeleteContextOutputOptions definition(and ContextOpts struct, which is embedded)
// nolint: dupl
func NewDeleteContextCommand(inputOpts *DeleteContextInputOptions, outputOpts *DeleteContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.DeleteContextAPIName

	// Validate the input options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ContextName: inputOpts.ContextName,
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts != nil && outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	}

	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	c.APIs = append(c.APIs, api)
	return c, nil
}

// NewSetCurrentContextCommand constructs a command to make a call to specific runtime version SetCurrentContext API
// Input Parameter inputOpts has all input parameters which are required for Runtime SetCurrentContext API
// Input Parameter: outputOpts has details about expected output from Runtime SetCurrentContext API call
// Return: command to execute or error if any validations fails for SetCurrentContextInputOptions or SetCurrentContextOutputOptions
// This method does validate the input parameters  SetCurrentContextInputOptions or SetCurrentContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetCurrentContextInputOptions or SetCurrentContextOutputOptions definition(and ContextOpts struct, which is embedded)
func NewSetCurrentContextCommand(inputOpts *SetCurrentContextInputOptions, outputOpts *SetCurrentContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{Name: core.SetCurrentContextAPIName}

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ContextName: inputOpts.ContextName,
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts != nil && outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	}

	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	c.APIs = append(c.APIs, api)

	return c, nil
}

// NewGetCurrentContextCommand constructs a command to make a call to specific runtime version GetCurrentContext API
// Input Parameter inputOpts has all input parameters which are required for Runtime GetCurrentContext API
// Input Parameter: outputOpts has details about expected output from Runtime GetCurrentContext API call
// Return: command to execute or error if any validations fails for GetCurrentContextInputOptions or GetCurrentContextOutputOptions
// This method does validate the input parameters  GetCurrentContextInputOptions or GetCurrentContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetCurrentContextInputOptions or GetCurrentContextOutputOptions definition(and ContextOpts struct, which is embedded)
func NewGetCurrentContextCommand(inputOpts *GetCurrentContextInputOptions, outputOpts *GetCurrentContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}
	// Init the API object
	api := &core.API{Name: core.GetCurrentContextAPIName}

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the context api arguments and output
	api.Arguments = make(map[core.APIArgumentType]interface{})

	if inputOpts.Target != "" {
		api.Arguments[core.Target] = inputOpts.Target
	} else if inputOpts.ContextType != "" {
		api.Arguments[core.ContextType] = inputOpts.ContextType
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.ContextOpts != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		// Construct get current context output context opts
		bytes, err := yaml.Marshal(outputOpts.ContextOpts)
		if err != nil {
			return nil, err
		}

		content = string(bytes)
		res = core.Success
	} else if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	}

	api.Output = &core.Output{
		Result:  res,
		Content: content,
	}

	if outputOpts.ValidationStrategy != "" {
		api.Output.ValidationStrategy = outputOpts.ValidationStrategy
	}

	c.APIs = append(c.APIs, api)
	return c, nil
}

// NewRemoveCurrentContextCommand constructs a command to make a call to specific runtime version RemoveCurrentContext API
// Input Parameter inputOpts has all input parameters which are required for Runtime RemoveCurrentContext API
// Input Parameter: outputOpts has details about expected output from Runtime RemoveCurrentContext API call
// Return: command to execute or error if any validations fails for RemoveCurrentContextInputOptions or RemoveCurrentContextOutputOptions
// This method does validate the input parameters  RemoveCurrentContextInputOptions/ RemoveCurrentContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to RemoveCurrentContextInputOptions/ RemoveCurrentContextOutputOptions definition(and ContextOpts struct, which is embedded)
func NewRemoveCurrentContextCommand(removeCurrentContextInputOptions *RemoveCurrentContextInputOptions, removeCurrentContextOutputOptions *RemoveCurrentContextOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}
	api.Name = core.RemoveCurrentContextAPIName
	api.Version = removeCurrentContextInputOptions.RuntimeVersion

	// Validate the Input Options
	if removeCurrentContextInputOptions.Target == "" {
		return nil, fmt.Errorf("context target is required")
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
