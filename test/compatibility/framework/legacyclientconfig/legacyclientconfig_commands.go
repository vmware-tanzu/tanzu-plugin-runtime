// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package legacyclientconfig provides api command helpers and validators to write compatibility tests for legacy client config apis
package legacyclientconfig

import (
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewStoreClientConfigCommand constructs a command to make a call to StoreClientConfig API
// Input Parameter: inputOpts has all input parameters which are required for Runtime StoreClientConfig API
// Input Parameter: outputOpts has details about expected output from Runtime StoreClientConfig API call
// Return: command to execute or error if any validations fails for StoreClientConfigInputOptions or StoreClientConfigOutputOptions
// This method does validate the input parameters StoreClientConfigInputOptions or StoreClientConfigOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetContextInputOptions or StoreClientConfigOutputOptions definition (and ClientConfigOpts struct, which is embedded)
func NewStoreClientConfigCommand(inputOpts *StoreClientConfigInputOptions, outputOpts *StoreClientConfigOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.StoreClientConfigAPI

	// Validate the StoreClientConfig input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the StoreClientConfig API arguments
	bytes, err := yaml.Marshal(inputOpts.ClientConfigOpts)
	if err != nil {
		return nil, err
	}

	api.Arguments = map[core.APIArgumentType]interface{}{
		core.ClientConfig: string(bytes),
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

// NewGetClientConfigCommand constructs a command to make a call to GetClientConfig API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetClientConfig API
// Input Parameter: outputOpts has details about expected output from Runtime GetClientConfig API call
// Return: command to execute or error if any validations fails for GetClientConfigInputOptions or GetClientConfigOutputOptions
// This method does validate the input parameters GetClientConfigInputOptions or GetClientConfigOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetClientConfigInputOptions or GetClientConfigOutputOptions definition (and ClientConfigOpts struct, which is embedded)
func NewGetClientConfigCommand(inputOpts *GetClientConfigInputOptions, outputOpts *GetClientConfigOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetClientConfigAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the GetClientConfig API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{}

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.ClientConfigOpts != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		// Construct get context output context opts
		bytes, err := yaml.Marshal(outputOpts.ClientConfigOpts)
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
