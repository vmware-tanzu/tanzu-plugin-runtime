// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"
	"strconv"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewSetFeatureCommand constructs a command to make a call to specific runtime version SetFeature API
// Input Parameter inputOpts has all input parameters which are required for Runtime SetFeature API
// Input Parameter: outputOpts has details about expected output from Runtime SetFeature API call
// Return: command to execute or error if any validations fails for SetFeatureInputOptions or SetFeatureOutputOptions
// This method does validate the input parameters  SetFeatureInputOptions or SetFeatureOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetFeatureInputOptions or SetFeatureOutputOptions definition(and FeatureOpts struct, which is embedded)
func NewSetFeatureCommand(inputOpts *SetFeatureInputOptions, outputOpts *SetFeatureOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.SetFeatureAPI

	// Validate the SetFeature input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetFeature API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.PluginName: inputOpts.PluginName,
		core.KeyName:    inputOpts.KeyName,
		core.ValueName:  inputOpts.ValueName,
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

// NewIsFeatureEnabledCommand constructs a command to make a call to specific runtime version IsFeatureEnabled API
// Input Parameter inputOpts has all input parameters which are required for Runtime IsFeatureEnabled API
// Input Parameter: outputOpts has details about expected output from Runtime IsFeatureEnabled API call
// Return: command to execute or error if any validations fails for IsFeatureEnabledInputOptions or IsFeatureEnabledOutputOptions
// This method does validate the input parameters  IsFeatureEnabledInputOptions or IsFeatureEnabledOutputOptions based on Runtime API Version
// For more details about supported parameters refer to IsFeatureEnabledInputOptions or IsFeatureEnabledOutputOptions definition(and FeatureOpts struct, which is embedded)
func NewIsFeatureEnabledCommand(inputOpts *IsFeatureEnabledInputOptions, outputOpts *IsFeatureEnabledOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.IsFeatureEnabledAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the IsFeatureEnabled API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.PluginName: inputOpts.PluginName,
		core.KeyName:    inputOpts.KeyName,
		core.Feature:    fmt.Sprintf("features.%v.%v", inputOpts.PluginName, inputOpts.KeyName),
	}

	// Construct Output parameters
	var res core.Result
	var content string

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}

		content = strconv.FormatBool(outputOpts.FeatureEnabled)
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

// NewDeleteFeatureCommand constructs a command to make a call to specific runtime version DeleteFeature API
// Input Parameter inputOpts has all input parameters which are required for Runtime DeleteFeature API
// Input Parameter: outputOpts has details about expected output from Runtime DeleteFeature API call
// Return: command to execute or error if any validations fails for DeleteFeatureInputOptions or DeleteFeatureOutputOptions
// This method does validate the input parameters  DeleteFeatureInputOptions or DeleteFeatureOutputOptions based on Runtime API Version
// For more details about supported parameters refer to DeleteFeatureInputOptions or DeleteFeatureOutputOptions definition(and FeatureOpts struct, which is embedded)
func NewDeleteFeatureCommand(inputOpts *DeleteFeatureInputOptions, outputOpts *DeleteFeatureOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.DeleteFeatureAPI

	// Validate the input options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the context api arguments and output
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.PluginName: inputOpts.PluginName,
		core.KeyName:    inputOpts.KeyName,
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
