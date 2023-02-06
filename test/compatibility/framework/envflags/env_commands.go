// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package envflags provides api command helpers and validators to write compatibility tests for env apis
package envflags

import (
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewSetEnvCommand constructs a command to make a call to specific runtime version SetEnv API
// Input Parameter: inputOpts has all input parameters which are required for Runtime SetEnv API
// Input Parameter: outputOpts has details about expected output from Runtime SetEnv API call
// Return: command to execute or error if any validations fails for SetEnvInputOptions or SetEnvOutputOptions
// This method does validate the input parameters SetEnvInputOptions or SetEnvOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetEnvInputOptions or SetEnvOutputOptions definition(and EnvOpts struct, which is embedded)
func NewSetEnvCommand(inputOpts *SetEnvInputOptions, outputOpts *SetEnvOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.SetEnvAPI

	// Validate the SetEnv input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetEnv API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Key:   inputOpts.Key,
		core.Value: inputOpts.Value,
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

// NewGetEnvCommand constructs a command to make a call to specific runtime version GetEnv API
// Input Parameter inputOpts has all input parameters which are required for Runtime GetEnv API
// Input Parameter: outputOpts has details about expected output from Runtime GetEnv API call
// Return: command to execute or error if any validations fails for GetEnvInputOptions or GetEnvOutputOptions
// This method does validate the input parameters GetEnvInputOptions or GetEnvOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetEnvInputOptions or GetEnvOutputOptions definition(and EnvOpts struct, which is embedded)
func NewGetEnvCommand(inputOpts *GetEnvInputOptions, outputOpts *GetEnvOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetEnvAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the GetEnv API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Key: inputOpts.Key,
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.Value != "" {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		content = outputOpts.Value
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

// NewDeleteEnvCommand constructs a command to make a call to specific runtime version DeleteEnv API
// Input Parameter inputOpts has all input parameters which are required for Runtime DeleteEnv API
// Input Parameter: outputOpts has details about expected output from Runtime DeleteEnv API call
// Return: command to execute or error if any validations fails for DeleteEnvInputOptions or DeleteEnvOutputOptions
// This method does validate the input parameters DeleteEnvInputOptions or DeleteEnvOutputOptions based on Runtime API Version
// For more details about supported parameters refer to DeleteEnvInputOptions or DeleteEnvOutputOptions definition(and EnvOpts struct, which is embedded)
func NewDeleteEnvCommand(inputOpts *DeleteEnvInputOptions, outputOpts *DeleteEnvOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.DeleteEnvAPI

	// Validate the input options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Key: inputOpts.Key,
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

// NewGetEnvConfigurationsCommand constructs a command to make a call to specific runtime version GetEnvConfigurations API
// Input Parameter inputOpts has all input parameters which are required for Runtime GetEnv API
// Input Parameter: outputOpts has details about expected output from Runtime GetEnv API call
// Return: command to execute or error if any validations fails for GetEnvConfigurationsInputOptions or GetEnvConfigurationsOutputOptions
// This method does validate the input parameters GetEnvConfigurationsInputOptions or GetEnvConfigurationsOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetEnvConfigurationsInputOptions or GetEnvConfigurationsOutputOptions definition(and EnvOpts struct, which is embedded)
func NewGetEnvConfigurationsCommand(inputOpts *GetEnvConfigurationsInputOptions, outputOpts *GetEnvConfigurationsOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetEnvConfigurationsAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the GetEnvConfigurations API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.Envs != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		bytes, err := yaml.Marshal(outputOpts.Envs)
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
