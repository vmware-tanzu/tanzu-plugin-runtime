// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package clidiscoverysources provides api command helpers and validators to write compatibility tests for CLI Discovery sources apis
package clidiscoverysources

import (
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewSetCLIDiscoverySourceCommand constructs a command to make a call to specific runtime version SetCLIDiscoverySource API
// Input Parameter inputOpts has all input parameters which are required for Runtime SetCLIDiscoverySource API
// Input Parameter: outputOpts has details about expected output from Runtime SetCLIDiscoverySource API call
// Return: command to execute or error if any validations fails for SetCLIDiscoverySourceInputOptions or SetCLIDiscoverySourceOutputOptions
// This method does validate the input parameters SetCLIDiscoverySourceInputOptions or SetCLIDiscoverySourceOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetCLIDiscoverySourceInputOptions or SetCLIDiscoverySourceOutputOptions definition(and CLIDiscoverySourceOpts struct, which is embedded)
func NewSetCLIDiscoverySourceCommand(inputOpts *SetCLIDiscoverySourceInputOptions, outputOpts *SetCLIDiscoverySourceOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.SetCLIDiscoverySourceAPI

	// Validate the SetCLIDiscoverySource input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetCLIDiscoverySource API arguments
	bytes, err := yaml.Marshal(inputOpts.PluginDiscoveryOpts)
	if err != nil {
		return nil, err
	}

	api.Arguments = map[core.APIArgumentType]interface{}{
		core.DiscoverySource: string(bytes),
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

// NewGetCLIDiscoverySourceCommand constructs a command to make a call to specific runtime version GetCLIDiscoverySource API
// Input Parameter inputOpts has all input parameters which are required for Runtime GetCLIDiscoverySource API
// Input Parameter: outputOpts has details about expected output from Runtime GetCLIDiscoverySource API call
// Return: command to execute or error if any validations fails for GetCLIDiscoverySourceInputOptions or GetCLIDiscoverySourceOutputOptions
// This method does validate the input parameters GetCLIDiscoverySourceInputOptions or GetCLIDiscoverySourceOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetCLIDiscoverySourceInputOptions or GetCLIDiscoverySourceOutputOptions definition(and CLIDiscoverySourceOpts struct, which is embedded)
func NewGetCLIDiscoverySourceCommand(inputOpts *GetCLIDiscoverySourceInputOptions, outputOpts *GetCLIDiscoverySourceOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetCLIDiscoverySourceAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the GetCLIDiscoverySource API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Name: inputOpts.DiscoverySourceName,
	}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts != nil {
		if outputOpts.Error != "" {
			res = core.Failed
			content = outputOpts.Error
		} else if outputOpts.PluginDiscoveryOpts != nil {
			// Validate the Output Options
			_, err = outputOpts.Validate()
			if err != nil {
				return nil, err
			}

			// Construct get server output server opts
			bytes, err := yaml.Marshal(outputOpts.PluginDiscoveryOpts)
			if err != nil {
				return nil, err
			}

			content = string(bytes)
			res = core.Success
		}
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

// NewDeleteCLIDiscoverySourceCommand constructs a command to make a call to specific runtime version DeleteCLIDiscoverySource API
// Input Parameter inputOpts has all input parameters which are required for Runtime DeleteCLIDiscoverySource API
// Input Parameter: outputOpts has details about expected output from Runtime DeleteCLIDiscoverySource API call
// Return: command to execute or error if any validations fails for DeleteCLIDiscoverySourceInputOptions or DeleteCLIDiscoverySourceOutputOptions
// This method does validate the input parameters DeleteCLIDiscoverySourceInputOptions or DeleteCLIDiscoverySourceOutputOptions based on Runtime API Version
// For more details about supported parameters refer to DeleteCLIDiscoverySourceInputOptions or DeleteCLIDiscoverySourceOutputOptions definition(and CLIDiscoverySourceOpts struct, which is embedded)
func NewDeleteCLIDiscoverySourceCommand(inputOpts *DeleteCLIDiscoverySourceInputOptions, outputOpts *DeleteCLIDiscoverySourceOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.DeleteCLIDiscoverySourceAPI

	// Validate the input options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the server api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Name: inputOpts.DiscoverySourceName,
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
