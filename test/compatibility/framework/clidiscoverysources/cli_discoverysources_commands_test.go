// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clidiscoverysources

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

const source string = `oci:
    name: compatibility-tests-source-name
    image: compatibility-tests-source-image
`

func TestNewSetCLIDiscoverySourceCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *SetCLIDiscoverySourceInputOptions
		outputOpts *SetCLIDiscoverySourceOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&SetCLIDiscoverySourceInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				PluginDiscoveryOpts: &types.PluginDiscoveryOpts{
					OCI: &types.OCIDiscoveryOpts{
						Name:  CompatibilityTestsSourceName,
						Image: CompatibilityTestsSourceImage,
					},
				},
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.SetCLIDiscoverySourceAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.DiscoverySource: source,
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content:            "",
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewSetCLIDiscoverySourceCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewGetCLIDiscoverySourceCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *GetCLIDiscoverySourceInputOptions
		outputOpts *GetCLIDiscoverySourceOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&GetCLIDiscoverySourceInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				DiscoverySourceName: CompatibilityTestsSourceName,
			}, &GetCLIDiscoverySourceOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				PluginDiscoveryOpts: &types.PluginDiscoveryOpts{
					OCI: &types.OCIDiscoveryOpts{
						Name:  CompatibilityTestsSourceName,
						Image: CompatibilityTestsSourceImage,
					},
				},
			},
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.GetCLIDiscoverySourceAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.Name: CompatibilityTestsSourceName,
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content:            source,
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewGetCLIDiscoverySourceCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewDeleteCLIDiscoverySourceCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *DeleteCLIDiscoverySourceInputOptions
		outputOpts *DeleteCLIDiscoverySourceOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&DeleteCLIDiscoverySourceInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				DiscoverySourceName: CompatibilityTestsSourceName,
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.DeleteCLIDiscoverySourceAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.Name: CompatibilityTestsSourceName,
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content:            "",
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewDeleteCLIDiscoverySourceCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}
