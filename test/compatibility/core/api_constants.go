// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package core contains compatibility testing framework core types and functions
package core

// RuntimeAPIName describes all the runtime api functions
type RuntimeAPIName string

const (
	SetContextAPIName RuntimeAPIName = "SetContext"
	GetContextAPIName RuntimeAPIName = "GetContext"
	AddServerAPIName  RuntimeAPIName = "AddServer"
	GetServerAPIName  RuntimeAPIName = "GetServer"
)

// APIArgumentType describes all the arguments types for runtime api functions
type APIArgumentType string

const (
	Context     APIArgumentType = "context"
	ContextName APIArgumentType = "contextName"
	SetCurrent  APIArgumentType = "setCurrent"
	Server      APIArgumentType = "server"
	ServerName  APIArgumentType = "serverName"
)

type Result string

const (
	Success Result = "success"
	Failed  Result = "failed"
)

// ValidationStrategy used to describe the validation strategy. default is non-strict
type ValidationStrategy string

const ValidationStrategyStrict ValidationStrategy = "strict"

// RuntimeVersion Runtime library versions
type RuntimeVersion string

const (
	Version0116 RuntimeVersion = "v0.11.6"
	Version0254 RuntimeVersion = "v0.25.4"
	Version0280 RuntimeVersion = "v0.28.0"
	Version100  RuntimeVersion = "v1.0.0"
)

var SupportedRuntimeVersions = []RuntimeVersion{
	Version0116,
	Version0254,
	Version0280,
	Version100,
}

type ResponseType string

const (
	MapResponse     ResponseType = "map"
	BooleanResponse ResponseType = "bool"
	StringResponse  ResponseType = "str"
	IntegerResponse ResponseType = "int"
	ErrorResponse   ResponseType = "err"
)
