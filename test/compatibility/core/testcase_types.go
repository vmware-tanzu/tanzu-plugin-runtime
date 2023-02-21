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
	Name      RuntimeAPIName         `json:"name" yaml:"name"`
	Version   RuntimeVersion         `json:"version" yaml:"version"`
	Arguments map[string]interface{} `json:"arguments" yaml:"arguments"`
	Output    *Output                `json:"output" yaml:"output"`
}

// Output represents the runtime api expected output for validation
type Output struct {
	Result  Result `json:"result" yaml:"result"`
	Content string `json:"content" yaml:"content"`
}

type Result string

const (
	Success Result = "success"
	Failed         = "failed"
)

// RuntimeAPIVersion represents the runtime library version
type RuntimeAPIVersion struct {
	RuntimeVersion RuntimeVersion `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`
}

// RuntimeVersion Runtime library versions
type RuntimeVersion string

const (
	Version0116 RuntimeVersion = "v0.11.6"
	Version0254                = "v0.25.4"
	Version0280                = "v0.28.0"
	Version100                 = "v1.0.0"
)

// NewTestCase creates an instance of TestCase
func NewTestCase() *TestCase {
	return &TestCase{}
}

// Add series of commands to test case to be executed in sequence
func (t *TestCase) Add(command ...*Command) *TestCase {
	t.Commands = append(t.Commands, command...)
	return t
}

// APILog represents the logs/output/errors returned from runtime apis
type APILog struct {
	APIResponse *APIResponse `json:"apiResponse" yaml:"apiResponse"`
	APIError    string       `json:"error" yaml:"error"`
}

// APIResponse represents the output response returned from runtime apis
type APIResponse struct {
	ResponseType ResponseType `json:"responseType" yaml:"responseType"`
	ResponseBody interface{}  `json:"responseBody" yaml:"responseBody"`
}

type ResponseType string

const (
	MapResponse     ResponseType = "map"
	BooleanResponse              = "bool"
	StringResponse               = "str"
	IntegerResponse              = "int"
	ErrorResponse                = "err"
)
