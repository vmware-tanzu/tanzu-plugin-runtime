// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func TestNewSetServerCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *SetServerInputOptions
		outputOpts *SetServerOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&SetServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version090,
				},
				ServerOpts: &types.ServerOpts{
					Name: "compatibility-one",
					Type: types.ManagementClusterServerType,
					GlobalOpts: &types.GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.SetServerAPI,
						Version: core.Version090,
						Arguments: map[core.APIArgumentType]interface{}{
							core.Server: `name: compatibility-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
							"setCurrent": false,
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
		{
			&SetServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ServerOpts: &types.ServerOpts{
					Name: "compatibility-one",
					Type: types.ManagementClusterServerType,
					GlobalOpts: &types.GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.SetServerAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.Server: `name: compatibility-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
							"setCurrent": false,
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
		cmd, err := NewSetServerCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewGetServerCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *GetServerInputOptions
		outputOpts *GetServerOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&GetServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version090,
				},
				ServerName: "compatibility-one",
			}, &GetServerOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ServerOpts: &types.ServerOpts{
					Name: "compatibility-one",
					Type: types.ManagementClusterServerType,
					GlobalOpts: &types.GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			},
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.GetServerAPI,
						Version: core.Version090,
						Arguments: map[core.APIArgumentType]interface{}{
							core.ServerName: "compatibility-one",
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content: `name: compatibility-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
						},
					},
				},
			}, "",
		},
		{
			&GetServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ServerName: "compatibility-one",
			}, &GetServerOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ServerOpts: &types.ServerOpts{
					Name: "compatibility-one",
					Type: types.ManagementClusterServerType,
					GlobalOpts: &types.GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			},
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.GetServerAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.ServerName: "compatibility-one",
						},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content: `name: compatibility-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewGetServerCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewDeleteServerCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *DeleteServerInputOptions
		outputOpts *DeleteServerOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&DeleteServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version090,
				},
				ServerName: "compatibility-one",
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.DeleteServerAPI,
						Version: core.Version090,
						Arguments: map[core.APIArgumentType]interface{}{
							core.ServerName: "compatibility-one",
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
		{
			&DeleteServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ServerName: "compatibility-one",
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.DeleteServerAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.ServerName: "compatibility-one",
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
		cmd, err := NewDeleteServerCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewSetCurrentServerCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *SetCurrentServerInputOptions
		outputOpts *SetCurrentServerOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&SetCurrentServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version090,
				},
				ServerName: "compatibility-one",
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.SetCurrentServerAPI,
						Version: core.Version090,
						Arguments: map[core.APIArgumentType]interface{}{
							core.ServerName: "compatibility-one",
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
		{
			&SetCurrentServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ServerName: "compatibility-one",
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.SetCurrentServerAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.ServerName: "compatibility-one",
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
		cmd, err := NewSetCurrentServerCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewGetCurrentServerCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *GetCurrentServerInputOptions
		outputOpts *GetCurrentServerOutputOptions
		cmd        *core.Command
		err        string
	}{
		{
			&GetCurrentServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version090,
				},
			}, &GetCurrentServerOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ServerOpts: &types.ServerOpts{
					Name: "compatibility-one",
					Type: types.ManagementClusterServerType,
					GlobalOpts: &types.GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			},
			&core.Command{
				APIs: []*core.API{
					{
						Name:      core.GetCurrentServerAPI,
						Version:   core.Version090,
						Arguments: map[core.APIArgumentType]interface{}{},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content: `name: compatibility-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
						},
					},
				},
			}, "",
		},
		{
			&GetCurrentServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
			}, &GetCurrentServerOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ServerOpts: &types.ServerOpts{
					Name: "compatibility-one",
					Type: types.ManagementClusterServerType,
					GlobalOpts: &types.GlobalServerOpts{
						Endpoint: "default-compatibility-test-endpoint",
					},
				},
			},
			&core.Command{
				APIs: []*core.API{
					{
						Name:      core.GetCurrentServerAPI,
						Version:   core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{},
						Output: &core.Output{
							ValidationStrategy: "",
							Result:             core.Success,
							Content: `name: compatibility-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
						},
					},
				},
			}, "",
		},
	}

	for _, tt := range tests {
		cmd, err := NewGetCurrentServerCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}

func TestNewRemoveCurrentServerCommand(t *testing.T) {
	tests := []struct {
		inputOpts  *RemoveCurrentServerInputOptions
		outputOpts *RemoveCurrentServerOutputOptions
		cmd        *core.Command
		err        string
	}{{
		&RemoveCurrentServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version090,
			},
			ServerName: "compatibility-one",
		}, nil,
		&core.Command{
			APIs: []*core.API{
				{
					Name:    core.RemoveCurrentServerAPI,
					Version: core.Version090,
					Arguments: map[core.APIArgumentType]interface{}{
						core.ServerName: "compatibility-one",
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
		{
			&RemoveCurrentServerInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.VersionLatest,
				},
				ServerName: "compatibility-one",
			}, nil,
			&core.Command{
				APIs: []*core.API{
					{
						Name:    core.RemoveCurrentServerAPI,
						Version: core.VersionLatest,
						Arguments: map[core.APIArgumentType]interface{}{
							core.ServerName: "compatibility-one",
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
		cmd, err := NewRemoveCurrentServerCommand(tt.inputOpts, tt.outputOpts)
		if tt.err != "" {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Equal(t, tt.cmd, cmd)
		}
	}
}
