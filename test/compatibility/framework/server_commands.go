// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewSetServerCommand constructs a command to make a call to specific runtime version SetServer API
// Input Parameter inputOpts has all input parameters which are required for Runtime SetServer API
// Input Parameter: outputOpts has details about expected output from Runtime SetServer API call
// Return: command to execute or error if any validations fails for SetServerInputOptions or SetServerOutputOptions
// This method does validate the input parameters  SetServerInputOptions or SetServerOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetServerInputOptions or SetServerOutputOptions definition(and ServerOpts struct, which is embedded)
// nolint:dupl
func NewSetServerCommand(inputOpts *SetServerInputOptions, outputOpts *SetServerOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.SetServerAPIName

	// Validate the SetServer input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetServer API arguments
	bytes, err := yaml.Marshal(inputOpts.ServerOpts)
	if err != nil {
		return nil, err
	}

	// Construct the setCurrent Argument
	var setCurrent = false
	if inputOpts.SetCurrentServer {
		setCurrent = true
	}

	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Server:     string(bytes),
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

// NewGetServerCommand constructs a command to make a call to specific runtime version GetServer API
// Input Parameter inputOpts has all input parameters which are required for Runtime GetServer API
// Input Parameter: outputOpts has details about expected output from Runtime GetServer API call
// Return: command to execute or error if any validations fails for GetServerInputOptions or GetServerOutputOptions
// This method does validate the input parameters  GetServerInputOptions or GetServerOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetServerInputOptions or GetServerOutputOptions definition(and ServerOpts struct, which is embedded)
func NewGetServerCommand(inputOpts *GetServerInputOptions, outputOpts *GetServerOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetServerAPIName

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the GetServer API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ServerName: inputOpts.ServerName,
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.ServerOpts != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		// Construct get server output server opts
		bytes, err := yaml.Marshal(outputOpts.ServerOpts)
		if err != nil {
			return nil, err
		}

		content = string(bytes)
		res = core.Success
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

// NewDeleteServerCommand constructs a command to make a call to specific runtime version DeleteServer API
// Input Parameter inputOpts has all input parameters which are required for Runtime DeleteServer API
// Input Parameter: outputOpts has details about expected output from Runtime DeleteServer API call
// Return: command to execute or error if any validations fails for DeleteServerInputOptions or DeleteServerOutputOptions
// This method does validate the input parameters  DeleteServerInputOptions or DeleteServerOutputOptions based on Runtime API Version
// For more details about supported parameters refer to DeleteServerInputOptions or DeleteServerOutputOptions definition(and ServerOpts struct, which is embedded)
// nolint: dupl
func NewDeleteServerCommand(inputOpts *DeleteServerInputOptions, outputOpts *DeleteServerOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.DeleteServerAPIName

	// Validate the input options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the server api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ServerName: inputOpts.ServerName,
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

// NewSetCurrentServerCommand constructs a command to make a call to specific runtime version SetCurrentServer API
// Input Parameter inputOpts has all input parameters which are required for Runtime SetCurrentServer API
// Input Parameter: outputOpts has details about expected output from Runtime SetCurrentServer API call
// Return: command to execute or error if any validations fails for SetCurrentServerInputOptions or SetCurrentServerOutputOptions
// This method does validate the input parameters  SetCurrentServerInputOptions or SetCurrentServerOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetCurrentServerInputOptions or SetCurrentServerOutputOptions definition(and ServerOpts struct, which is embedded)
func NewSetCurrentServerCommand(inputOpts *SetCurrentServerInputOptions, outputOpts *SetCurrentServerOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{Name: core.SetCurrentServerAPIName}

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the server api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ServerName: inputOpts.ServerName,
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

// NewGetCurrentServerCommand constructs a command to make a call to specific runtime version GetCurrentServer API
// Input Parameter inputOpts has all input parameters which are required for Runtime GetCurrentServer API
// Input Parameter: outputOpts has details about expected output from Runtime GetCurrentServer API call
// Return: command to execute or error if any validations fails for GetCurrentServerInputOptions or GetCurrentServerOutputOptions
// This method does validate the input parameters  GetCurrentServerInputOptions or GetCurrentServerOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetCurrentServerInputOptions or GetCurrentServerOutputOptions definition(and ServerOpts struct, which is embedded)
func NewGetCurrentServerCommand(inputOpts *GetCurrentServerInputOptions, outputOpts *GetCurrentServerOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}
	// Init the API object
	api := &core.API{Name: core.GetCurrentServerAPIName}

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the server api arguments and output
	api.Arguments = make(map[core.APIArgumentType]interface{})

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.ServerOpts != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		// Construct get current server output server opts
		bytes, err := yaml.Marshal(outputOpts.ServerOpts)
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

// NewRemoveCurrentServerCommand constructs a command to make a call to specific runtime version RemoveCurrentServer API
// Input Parameter inputOpts has all input parameters which are required for Runtime RemoveCurrentServer API
// Input Parameter: outputOpts has details about expected output from Runtime RemoveCurrentServer API call
// Return: command to execute or error if any validations fails for RemoveCurrentServerInputOptions or RemoveCurrentServerOutputOptions
// This method does validate the input parameters  RemoveCurrentServerInputOptions/ RemoveCurrentServerOutputOptions based on Runtime API Version
// For more details about supported parameters refer to RemoveCurrentServerInputOptions/ RemoveCurrentServerOutputOptions definition(and ServerOpts struct, which is embedded)
func NewRemoveCurrentServerCommand(inputOpts *RemoveCurrentServerInputOptions, outputOpts *RemoveCurrentServerOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{Name: core.RemoveCurrentServerAPIName}

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the server api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ServerName: inputOpts.ServerName,
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if outputOpts != nil && outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
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
