// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package types

// MetadataOpts struct type to store config related metadata
type MetadataOpts struct {
	// ConfigMetadata to store any config related metadata or settings
	ConfigMetadata *ConfigMetadataOpts `json:"configMetadata,omitempty" yaml:"configMetadata,omitempty" mapstructure:"configMetadata,omitempty"`
}

// ConfigMetadataOpts to store any config related metadata or settings
type ConfigMetadataOpts struct {
	// PatchStrategy patch strategy to determine merge of nodes in config file. Two ways of patch strategies are merge and replace
	PatchStrategy map[string]string `json:"patchStrategy,omitempty" yaml:"patchStrategy,omitempty" mapstructure:"patchStrategy,omitempty"`
	// Settings related to config
	Settings map[string]string `json:"settings,omitempty" yaml:"settings,omitempty" mapstructure:"settings,omitempty"`
}
