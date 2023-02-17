// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// NewSetContextCommand constructs a command to make a call to specific runtime version SetContextAPIName API.
// Input Parameter: setContextInputOptions has all input parameters which are required for Runtime SetContextAPIName API.
// Input Parameter: setContextOutputOptions has details about expected output from Runtime SetContextAPIName API call.
// Return:  command to execute or error if any validations fails for SetContextInputOptions or SetContextOutputOptions
// This method does validate the input parameters  SetContextInputOptions/SetContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetContextInputOptions/SetContextOutputOptions definition(and CtxOptions struct, which is embedded).
func NewSetContextCommand(setContextInputOptions *SetContextInputOptions, setContextOutputOptions *SetContextOutputOptions) (*Command, error) {
	// Init the Command object
	c := &Command{}

	// Init the API object
	api := &API{}
	api.Name = SetContextAPIName
	api.Version = setContextInputOptions.RuntimeVersion

	// Validate the SetContext Input Options
	_, err := ValidateSetContextInputOptionsAsPerRuntimeVersion(setContextInputOptions)
	Expect(err).To(BeNil())

	// Construct the set context api arguments and output
	bytes, err := yaml.Marshal(setContextInputOptions.ContextOpts)
	Expect(err).To(BeNil())

	// Construct the setCurrent Argument
	var setCurrent bool
	if setContextInputOptions.IsCurrentContext {
		setCurrent = true
	} else {
		setCurrent = false
	}

	api.Arguments = map[string]interface{}{
		"context":   string(bytes),
		"isCurrent": setCurrent,
	}

	// Construct Output parameters
	var res Result
	var content string

	if setContextOutputOptions != nil && setContextOutputOptions.Error != "" {
		res = Failed
		content = setContextOutputOptions.Error
	} else {
		res = Success
		content = ""
	}
	api.Output = &Output{
		Result:  res,
		Content: content,
	}

	c.APIs = append(c.APIs, api)

	return c, nil
}

// NewGetContextCommand creates a get context command object from inputOptions and outputOptions
// Creates the context specific command based on runtimeVersion passed in inputOptions also validates if the required input and outputOptions are passed
func NewGetContextCommand(getContextInputOptions *GetContextInputOptions, getContextOutputOptions *GetContextOutputOptions) (*Command, error) {
	// Init the Command object
	c := &Command{}
	// Init the API object
	api := &API{}
	api.Name = GetContextAPIName
	api.Version = getContextInputOptions.RuntimeVersion

	// Validate the Input Options
	if getContextInputOptions.ContextName == "" {
		return nil, errors.New("context name is required")
	}

	// Construct the context api arguments and output
	api.Arguments = map[string]interface{}{
		"contextName": getContextInputOptions.ContextName,
	}

	// Construct Output parameters
	var res Result
	var content string

	if getContextOutputOptions.ContextOpts != nil {
		// Validate the Output Options
		_, err := ValidateGetContextOutputOptionsAsPerRuntimeVersion(getContextOutputOptions)
		Expect(err).To(BeNil())

		// Construct get context output context opts
		bytes, err := yaml.Marshal(getContextOutputOptions.ContextOpts)
		Expect(err).To(BeNil())

		content = string(bytes)
		res = Success
	} else if getContextOutputOptions.Error != "" {
		res = Failed
		content = getContextOutputOptions.Error
	}

	api.Output = &Output{
		Result:  res,
		Content: content,
	}

	c.APIs = append(c.APIs, api)
	return c, nil
}
