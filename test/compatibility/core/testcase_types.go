// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

// TestCase represents the list of commands to execute as part of test case
type TestCase struct {
	Commands []*Command `json:"commands" yaml:"commands"`
}

// Command represents the list of apis to execute as part of command execution
type Command struct {
	APIs []*API `json:"apis" yaml:"apis"`
}

// API represents the runtime api to execute
type API struct {
	Name      RuntimeAPIName                  `json:"name" yaml:"name"`
	Version   RuntimeVersion                  `json:"version" yaml:"version"`
	Arguments map[APIArgumentType]interface{} `json:"arguments" yaml:"arguments"`
	Output    *Output                         `json:"output" yaml:"output"`
}

// Output represents the runtime api expected output for validation
type Output struct {
	ValidationStrategy ValidationStrategy `json:"validationstrategy" yaml:"validationstrategy"`
	Result             Result             `json:"result" yaml:"result"`
	Content            string             `json:"content" yaml:"content"`
}

// RuntimeAPIVersion represents the runtime library versions used in XXXOpts structs
type RuntimeAPIVersion struct {
	RuntimeVersion RuntimeVersion `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`
}

// NewTestCase creates an instance of TestCase
func NewTestCase() *TestCase {
	return &TestCase{}
}

// Add series of commands to test case to be executed in sequence
func (t *TestCase) Add(command ...*Command) *TestCase {
	if command != nil {
		t.Commands = append(t.Commands, command...)
		return t
	}
	return t
}

// APILog represents the logs/output/errors returned from runtime apis in test plugins
type APILog struct {
	APIResponse *APIResponse `json:"apiResponse" yaml:"apiResponse"`
}

// APIResponse represents the output response returned from runtime apis
type APIResponse struct {
	ResponseType ResponseType `json:"responseType" yaml:"responseType"`
	ResponseBody interface{}  `json:"responseBody" yaml:"responseBody"`
}
