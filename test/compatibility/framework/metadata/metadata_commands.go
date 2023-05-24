// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package metadata provides api command helpers and validators to write compatibility tests for metadata apis
package metadata

import (
	"strconv"

	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// NewSetConfigMetadataPatchStrategyCommand constructs a command to make a call to specific runtime version SetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime SetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime SetConfigMetadata API call
// Return: command to execute or error if any validations fails for SetConfigMetadataInputOptions or SetConfigMetadataOutputOptions
// This method does validate the input parameters SetConfigMetadataInputOptions or SetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetConfigMetadataInputOptions or SetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
//
//nolint:dupl
func NewSetConfigMetadataPatchStrategyCommand(inputOpts *SetConfigMetadataPatchStrategyInputOptions, outputOpts *SetConfigMetadataPatchStrategyOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.SetConfigMetadataPatchStrategyAPI

	// Validate the SetConfigMetadata input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetConfigMetadata API arguments
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

// NewSetConfigMetadataSettingCommand constructs a command to make a call to specific runtime version SetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime SetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime SetConfigMetadata API call
// Return: command to execute or error if any validations fails for SetConfigMetadataInputOptions or SetConfigMetadataOutputOptions
// This method does validate the input parameters SetConfigMetadataInputOptions or SetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetConfigMetadataInputOptions or SetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
//
//nolint:dupl
func NewSetConfigMetadataSettingCommand(inputOpts *SetConfigMetadataSettingInputOptions, outputOpts *SetConfigMetadataSettingOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.SetConfigMetadataSettingAPI

	// Validate the SetConfigMetadata input arguments
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetConfigMetadata API arguments
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

// NewDeleteConfigMetadataSettingCommand constructs a command to make a call to specific runtime version DeleteConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime DeleteConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime DeleteConfigMetadata API call
// Return: command to execute or error if any validations fails for DeleteConfigMetadataInputOptions or DeleteConfigMetadataOutputOptions
// This method does validate the input parameters DeleteConfigMetadataInputOptions or DeleteConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to DeleteConfigMetadataInputOptions or DeleteConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
func NewDeleteConfigMetadataSettingCommand(inputOpts *DeleteConfigMetadataSettingInputOptions, outputOpts *DeleteConfigMetadataSettingOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.DeleteConfigMetadataSettingAPI

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

// NewGetMetadataCommand constructs a command to make a call to specific runtime version GetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime GetConfigMetadata API call
// Return: command to execute or error if any validations fails for GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions
// This method does validate the input parameters GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
//
//nolint:dupl
func NewGetMetadataCommand(inputOpts *GetMetadataInputOptions, outputOpts *GetMetadataOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetMetadataAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.MetadataOpts != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}
		// Construct get context output context opts
		bytes, err := yaml.Marshal(outputOpts.MetadataOpts)
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

// NewGetConfigMetadataCommand constructs a command to make a call to specific runtime version GetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime GetConfigMetadata API call
// Return: command to execute or error if any validations fails for GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions
// This method does validate the input parameters GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
//
//nolint:dupl
func NewGetConfigMetadataCommand(inputOpts *GetConfigMetadataInputOptions, outputOpts *GetConfigMetadataOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetConfigMetadataAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.ConfigMetadataOpts != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}
		// Construct get context output context opts
		bytes, err := yaml.Marshal(outputOpts.ConfigMetadataOpts)
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

// NewGetConfigMetadataPatchStrategyCommand constructs a command to make a call to specific runtime version GetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime GetConfigMetadata API call
// Return: command to execute or error if any validations fails for GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions
// This method does validate the input parameters GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
//
//nolint:dupl
func NewGetConfigMetadataPatchStrategyCommand(inputOpts *GetConfigMetadataPatchStrategyInputOptions, outputOpts *GetConfigMetadataPatchStrategyOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetConfigMetadataPatchStrategyAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.PatchStrategy != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}
		// Construct get context output context opts
		bytes, err := yaml.Marshal(outputOpts.PatchStrategy)
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

// NewGetConfigMetadataSettingsCommand constructs a command to make a call to specific runtime version GetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime GetConfigMetadata API call
// Return: command to execute or error if any validations fails for GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions
// This method does validate the input parameters GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
//
//nolint:dupl
func NewGetConfigMetadataSettingsCommand(inputOpts *GetConfigMetadataSettingsInputOptions, outputOpts *GetConfigMetadataSettingsOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetConfigMetadataSettingsAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct Output parameters
	var res = core.Success
	var content = ""

	if outputOpts.Error != "" {
		res = core.Failed
		content = outputOpts.Error
	} else if outputOpts.MetadataSettings != nil {
		// Validate the Output Options
		_, err = outputOpts.Validate()
		if err != nil {
			return nil, err
		}
		// Construct get context output context opts
		bytes, err := yaml.Marshal(outputOpts.MetadataSettings)
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

// NewGetConfigMetadataSettingCommand constructs a command to make a call to specific runtime version GetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime GetConfigMetadata API call
// Return: command to execute or error if any validations fails for GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions
// This method does validate the input parameters GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
func NewGetConfigMetadataSettingCommand(inputOpts *GetConfigMetadataSettingInputOptions, outputOpts *GetConfigMetadataSettingOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.GetConfigMetadataSettingAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetConfigMetadata API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Key: inputOpts.Key,
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

// NewIsConfigMetadataSettingsEnabledCommand constructs a command to make a call to specific runtime version GetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime GetConfigMetadata API call
// Return: command to execute or error if any validations fails for GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions
// This method does validate the input parameters GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
func NewIsConfigMetadataSettingsEnabledCommand(inputOpts *IsConfigMetadataSettingsEnabledInputOptions, outputOpts *IsConfigMetadataSettingsEnabledOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.IsConfigMetadataSettingsEnabledAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

	// Construct the SetConfigMetadata API arguments
	api.Arguments = map[core.APIArgumentType]interface{}{
		core.Key: inputOpts.Key,
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

		content = strconv.FormatBool(outputOpts.Enabled)
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

// NewUseUnifiedConfigCommand constructs a command to make a call to specific runtime version GetConfigMetadata API
// Input Parameter: inputOpts has all input parameters which are required for Runtime GetConfigMetadata API
// Input Parameter: outputOpts has details about expected output from Runtime GetConfigMetadata API call
// Return: command to execute or error if any validations fails for GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions
// This method does validate the input parameters GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions based on Runtime API Version
// For more details about supported parameters refer to GetConfigMetadataInputOptions or GetConfigMetadataOutputOptions definition (and ConfigMetadataOpts struct, which is embedded)
func NewUseUnifiedConfigCommand(inputOpts *UseUnifiedConfigInputOptions, outputOpts *UseUnifiedConfigOutputOptions) (*core.Command, error) {
	// Init the Command object
	c := &core.Command{}

	// Init the API object
	api := &core.API{}

	// Set API name
	api.Name = core.UseUnifiedConfigAPI

	// Validate the Input Options
	_, err := inputOpts.Validate()
	if err != nil {
		return nil, err
	}

	// Set API version
	api.Version = inputOpts.RuntimeVersion

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

		content = strconv.FormatBool(outputOpts.Enabled)
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
