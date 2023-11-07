// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package context contains all the cross version api compatibility tests for context apis
package context

import (
	"fmt"

	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/types"
)

func SetContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		ContextType: types.ContextTypeK8s,
		Type:        types.CtxTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	var inputOpts *SetContextInputOptions
	var outputOpts *SetContextOutputOptions

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.VersionLatest:
		inputOpts = &SetContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextOpts: &types.ContextOpts{
				Name:        args.ContextName,
				Target:      args.Target,
				ContextType: args.ContextType,
				GlobalOpts:  args.GlobalOpts,
			},
		}
	case core.Version102, core.Version090, core.Version0280:
		inputOpts = &SetContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextOpts: &types.ContextOpts{
				Name:       args.ContextName,
				Target:     args.Target,
				GlobalOpts: args.GlobalOpts,
			},
		}
	case core.Version0254:
		inputOpts = &SetContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextOpts: &types.ContextOpts{
				Name:       args.ContextName,
				Type:       args.Type,
				GlobalOpts: args.GlobalOpts,
			},
		}
	}

	cmd, err := NewSetContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func GetContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		Type:        types.CtxTypeK8s,
		ContextType: types.ContextTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
		Error: false,
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ContextName:       args.ContextName,
	}

	var outputOpts *GetContextOutputOptions

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.VersionLatest:
		if args.Error {
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ContextName),
			}
		} else {
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:        args.ContextName,
					Target:      args.Target,
					ContextType: args.ContextType,
					GlobalOpts:  args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}
	case core.Version102, core.Version090, core.Version0280:
		if args.Error {
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ContextName),
			}
		} else {
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Target:     args.Target,
					GlobalOpts: args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}

	case core.Version0254:
		if args.Error {
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find context \"%v\"", args.ContextName),
			}
		} else {
			outputOpts = &GetContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	cmd, err := NewGetContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func DeleteContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		Type:        types.CtxTypeK8s,
		ContextType: types.ContextTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &DeleteContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ContextName:       args.ContextName,
	}

	var outputOpts *DeleteContextOutputOptions

	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.VersionLatest, core.Version102, core.Version090, core.Version0280:
			outputOpts = &DeleteContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("context %v not found", args.ContextName),
			}
		case core.Version0254:
			outputOpts = &DeleteContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("could not find context \"%v\"", args.ContextName),
			}
		}
	}

	cmd, err := NewDeleteContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func RemoveCurrentContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		Type:        types.CtxTypeK8s,
		ContextType: types.ContextTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &RemoveCurrentContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		Target:            args.Target,
	}

	var outputOpts *RemoveCurrentContextOutputOptions

	if args.Error {
		switch args.RuntimeAPIVersion.RuntimeVersion {
		case core.VersionLatest:
			outputOpts = &RemoveCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for type \"%v\"", args.Target),
			}
		case core.Version102, core.Version090, core.Version0280:
			outputOpts = &RemoveCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for target \"%v\"", args.Target),
			}
		case core.Version0254:
			outputOpts = &RemoveCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for type \"%v\"", args.Type),
			}
		}
	}

	cmd, err := NewRemoveCurrentContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func GetCurrentContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		Type:        types.CtxTypeK8s,
		ContextType: types.ContextTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	var inputOpts *GetCurrentContextInputOptions
	var outputOpts *GetCurrentContextOutputOptions

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.VersionLatest:
		inputOpts = &GetCurrentContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			Target:            args.Target,
		}
		if args.Error {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for type \"%v\"", args.Target),
			}
		} else {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:        args.ContextName,
					Target:      args.Target,
					ContextType: args.ContextType,
					GlobalOpts:  args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}
	case core.Version102, core.Version090, core.Version0280:
		inputOpts = &GetCurrentContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			Target:            args.Target,
		}

		if args.Error {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for target \"%v\"", args.Target),
			}
		} else {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Target:     args.Target,
					GlobalOpts: args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}

	case core.Version0254:
		inputOpts = &GetCurrentContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextType:       args.Type,
		}

		if args.Error {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for type \"%v\"", args.Type),
			}
		} else {
			outputOpts = &GetCurrentContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:       args.ContextName,
					Type:       args.Type,
					GlobalOpts: args.GlobalOpts,
				},
			}
		}
	}

	cmd, err := NewGetCurrentContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func SetCurrentContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		Type:        types.CtxTypeK8s,
		ContextType: types.ContextTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &SetCurrentContextInputOptions{
		RuntimeAPIVersion: args.RuntimeAPIVersion,
		ContextName:       args.ContextName,
	}

	var outputOpts *SetCurrentContextOutputOptions

	cmd, err := NewSetCurrentContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}

func GetActiveContextCommand(opts ...CfgContextArgsOption) *core.Command {
	args := &CfgContextArgs{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: core.VersionLatest,
		},
		ContextName: common.CompatibilityTestOne,
		Target:      types.TargetK8s,
		ContextType: types.ContextTypeK8s,
		GlobalOpts: &types.GlobalServerOpts{
			Endpoint: common.DefaultEndpoint,
		},
	}

	for _, opt := range opts {
		opt(args)
	}

	var inputOpts *GetActiveContextInputOptions
	var outputOpts *GetActiveContextOutputOptions

	switch args.RuntimeAPIVersion.RuntimeVersion {
	case core.VersionLatest:
		inputOpts = &GetActiveContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextType:       args.ContextType,
		}
		if args.Error {
			outputOpts = &GetActiveContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				Error:             fmt.Sprintf("no current context set for type \"%v\"", args.ContextType),
			}
		} else {
			outputOpts = &GetActiveContextOutputOptions{
				RuntimeAPIVersion: args.RuntimeAPIVersion,
				ContextOpts: &types.ContextOpts{
					Name:        args.ContextName,
					Target:      types.Target(args.ContextType),
					ContextType: args.ContextType,
					GlobalOpts:  args.GlobalOpts,
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}
		}
	default:
		// add runtime version and context type for unsupported versions
		inputOpts = &GetActiveContextInputOptions{
			RuntimeAPIVersion: args.RuntimeAPIVersion,
			ContextType:       args.ContextType,
		}
	}

	cmd, err := NewGetActiveContextCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())

	return cmd
}
