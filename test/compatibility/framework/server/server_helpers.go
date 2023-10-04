// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"

	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

type CfgServerArgs struct {
	RuntimeAPIVersion     *core.RuntimeAPIVersion
	ServerName            string // required
	Type                  types.ServerType
	SetCurrentServer      bool // required
	GlobalOpts            *types.GlobalServerOpts
	ManagementClusterOpts *types.ManagementClusterServerOpts
	DiscoverySources      []types.PluginDiscoveryOpts
	ValidationStrategy    core.ValidationStrategy
	Error                 bool
}

type CfgServerArgsOption func(args *CfgServerArgs)

func WithServerName(name string) CfgServerArgsOption {
	return func(c *CfgServerArgs) {
		c.ServerName = name
	}
}

func WithRuntimeVersion(version core.RuntimeVersion) CfgServerArgsOption {
	return func(c *CfgServerArgs) {
		c.RuntimeAPIVersion = &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		}
	}
}

func WithError() CfgServerArgsOption {
	return func(c *CfgServerArgs) {
		c.Error = true
	}
}

func SetServerCommand(opts ...CfgServerArgsOption) *core.Command {
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	var inputOpts *SetServerInputOptions
	var outputOpts *SetServerOutputOptions

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		inputOpts = &SetServerInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ServerOpts: &types.ServerOpts{
				Name:       args.ServerName,
				Type:       args.Type,
				GlobalOpts: args.GlobalOpts,
			},
		}
	}

	cmd, err := NewSetServerCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func GetServerCommand(opts ...CfgServerArgsOption) *core.Command {
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ServerName: common.CompatibilityTestOne,

		Type: types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
		Error: false,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	var outputOpts *GetServerOutputOptions

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
		if args.Error {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find server \"%v\"", args.ServerName),
			}
		} else {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ServerOpts: &types.ServerOpts{
					Name:       args.ServerName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}

	case core.Version0254, core.Version0116:
		if args.Error {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find server \"%v\"", args.ServerName),
			}
		} else {
			outputOpts = &GetServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ServerOpts: &types.ServerOpts{
					Name:       args.ServerName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	cmd, err := NewGetServerCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

//nolint:dupl
func DeleteServerCommand(opts ...CfgServerArgsOption) *core.Command {
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &DeleteServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	var outputOpts *DeleteServerOutputOptions

	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.VersionLatest, core.Version102, core.Version090, core.Version0280, core.Version0254, core.Version0116:
			outputOpts = &DeleteServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ServerName),
			}
		}
	}

	cmd, err := NewDeleteServerCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

//nolint:dupl
func RemoveCurrentServerCommand(opts ...CfgServerArgsOption) *core.Command {
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &RemoveCurrentServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	var outputOpts *RemoveCurrentServerOutputOptions

	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.VersionLatest, core.Version102, core.Version090, core.Version0280, core.Version0254, core.Version0116:
			outputOpts = &RemoveCurrentServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ServerName),
			}
		}
	}

	cmd, err := NewRemoveCurrentServerCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func GetCurrentServerCommand(opts ...CfgServerArgsOption) *core.Command {
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ServerName: common.CompatibilityTestOne,

		Type: types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	var inputOpts *GetCurrentServerInputOptions
	var outputOpts *GetCurrentServerOutputOptions

	inputOpts = &GetCurrentServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
	}

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.VersionLatest, core.Version102, core.Version090, core.Version0280, core.Version0254, core.Version0116:
		if args.Error {
			outputOpts = &GetCurrentServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             "current server \"\" not found in tanzu config",
			}
		} else {
			outputOpts = &GetCurrentServerOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ServerOpts: &types.ServerOpts{
					Name:       args.ServerName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	cmd, err := NewGetCurrentServerCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func SetCurrentServerCommand(opts ...CfgServerArgsOption) *core.Command {
	args := &CfgServerArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ServerName: common.CompatibilityTestOne,
		Type:       types.ManagementClusterServerType,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetCurrentServerInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ServerName:        args.ServerName,
	}

	var outputOpts *SetCurrentServerOutputOptions

	cmd, err := NewSetCurrentServerCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}
