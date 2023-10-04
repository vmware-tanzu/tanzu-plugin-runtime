// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package core contains compatibility testing framework core types and functions
package core

// RuntimeAPIName describes all the Runtime APIs
type RuntimeAPIName string

const (
	SetContextAPI           RuntimeAPIName = "SetContext"
	AddContextAPI           RuntimeAPIName = "AddContext"
	GetContextAPIName       RuntimeAPIName = "GetContext"
	DeleteContextAPI        RuntimeAPIName = "DeleteContext"
	RemoveContextAPI        RuntimeAPIName = "RemoveContext"
	SetCurrentContextAPI    RuntimeAPIName = "SetCurrentContext"
	GetCurrentContextAPI    RuntimeAPIName = "GetCurrentContext"
	RemoveCurrentContextAPI RuntimeAPIName = "RemoveCurrentContext"

	SetServerAPI           RuntimeAPIName = "SetServer"
	AddServerAPI           RuntimeAPIName = "AddServer"
	PutServerAPI           RuntimeAPIName = "PutServer"
	GetServerAPI           RuntimeAPIName = "GetServer"
	DeleteServerAPI        RuntimeAPIName = "DeleteServer"
	RemoveServerAPI        RuntimeAPIName = "RemoveServer"
	SetCurrentServerAPI    RuntimeAPIName = "SetCurrentServer"
	GetCurrentServerAPI    RuntimeAPIName = "GetCurrentServer"
	RemoveCurrentServerAPI RuntimeAPIName = "RemoveCurrentServer"

	SetFeatureAPI         RuntimeAPIName = "SetFeature"
	IsFeatureEnabledAPI   RuntimeAPIName = "IsFeatureEnabled"
	IsFeatureActivatedAPI RuntimeAPIName = "IsFeatureActivated"
	DeleteFeatureAPI      RuntimeAPIName = "DeleteFeature"

	SetEnvAPI               RuntimeAPIName = "SetEnv"
	GetEnvAPI               RuntimeAPIName = "GetEnv"
	DeleteEnvAPI            RuntimeAPIName = "DeleteEnv"
	GetEnvConfigurationsAPI RuntimeAPIName = "GetEnvConfigurations"

	SetCLIDiscoverySourceAPI    RuntimeAPIName = "SetCLIDiscoverySource"
	GetCLIDiscoverySourceAPI    RuntimeAPIName = "GetCLIDiscoverySource"
	DeleteCLIDiscoverySourceAPI RuntimeAPIName = "DeleteCLIDiscoverySource"

	GetMetadataAPI                     RuntimeAPIName = "GetMetadata"
	GetConfigMetadataAPI               RuntimeAPIName = "GetConfigMetadata"
	GetConfigMetadataPatchStrategyAPI  RuntimeAPIName = "GetConfigMetadataPatchStrategy"
	SetConfigMetadataPatchStrategyAPI  RuntimeAPIName = "SetConfigMetadataPatchStrategy"
	GetConfigMetadataSettingsAPI       RuntimeAPIName = "GetConfigMetadataSettings"
	GetConfigMetadataSettingAPI        RuntimeAPIName = "GetConfigMetadataSetting"
	IsConfigMetadataSettingsEnabledAPI RuntimeAPIName = "IsConfigMetadataSettingsEnabled"
	UseUnifiedConfigAPI                RuntimeAPIName = "UseUnifiedConfig"
	DeleteConfigMetadataSettingAPI     RuntimeAPIName = "DeleteConfigMetadataSetting"
	SetConfigMetadataSettingAPI        RuntimeAPIName = "SetConfigMetadataSetting"

	GetClientConfigAPI   RuntimeAPIName = "GetClientConfig"
	StoreClientConfigAPI RuntimeAPIName = "StoreClientConfig"
)

// APIArgumentType describes all the arguments types required for Runtime APIs
type APIArgumentType string

const (
	ClientConfig    APIArgumentType = "clientConfig"
	Context         APIArgumentType = "context"
	ContextName     APIArgumentType = "contextName"
	SetCurrent      APIArgumentType = "setCurrent"
	Server          APIArgumentType = "server"
	ServerName      APIArgumentType = "serverName"
	Name            APIArgumentType = "name"
	Target          APIArgumentType = "target"
	ContextType     APIArgumentType = "contextType"
	Feature         APIArgumentType = "feature"
	Plugin          APIArgumentType = "plugin"
	Key             APIArgumentType = "key"
	Value           APIArgumentType = "value"
	DiscoverySource APIArgumentType = "discoverySource"
)

type Result string

const (
	Success Result = "success"
	Failed  Result = "failed"
)

// ValidationStrategy used to describe the validation strategy. default is non-strict
type ValidationStrategy string

// ValidationStrategyStrict used to compare deep equality of objects
const ValidationStrategyStrict ValidationStrategy = "strict"

// RuntimeVersion Runtime library versions
type RuntimeVersion string

const (
	Version0116   RuntimeVersion = "v0.11.6"
	Version0254   RuntimeVersion = "v0.25.4"
	Version0280   RuntimeVersion = "v0.28.0"
	Version090    RuntimeVersion = "v0.90.0"
	Version102    RuntimeVersion = "v1.0.2"
	VersionLatest RuntimeVersion = "latest"
)

// SupportedRuntimeVersions Current supported runtime library versions
var SupportedRuntimeVersions = []RuntimeVersion{
	Version0116,
	Version0254,
	Version0280,
	Version090,
	Version102,
	VersionLatest,
}

type ResponseType string

const (
	MapResponse     ResponseType = "map"
	BooleanResponse ResponseType = "bool"
	StringResponse  ResponseType = "str"
	IntegerResponse ResponseType = "int"
	ErrorResponse   ResponseType = "err"
)
