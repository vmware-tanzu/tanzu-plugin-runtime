// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

func MakeSetContextInputOptions(version core.RuntimeVersion, contextName string) (*SetContextInputOptions, error) {
	switch version {
	case core.Version100:
		return &SetContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version100,
			},
			ContextOpts: &ContextOpts{
				Name:   contextName,
				Target: TargetK8s,
				GlobalOpts: &GlobalServerOpts{
					Endpoint: "test-endpoint",
				},
			},
		}, nil
	case core.Version0280:
		return &SetContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0280,
			},
			ContextOpts: &ContextOpts{
				Name:   contextName,
				Target: TargetK8s,
				GlobalOpts: &GlobalServerOpts{
					Endpoint: "test-endpoint",
				},
			},
		}, nil
	case core.Version0254:
		return &SetContextInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0254,
			},
			ContextOpts: &ContextOpts{
				Name: contextName,
				Type: CtxTypeK8s,
				GlobalOpts: &GlobalServerOpts{
					Endpoint: "test-endpoint",
				},
			},
		}, nil
	default:
		return nil, fmt.Errorf("context for runtime version %v is not supported", version)
	}

}

func MakeGetContextInputOptions(version core.RuntimeVersion, contextName string) *GetContextInputOptions {
	return &GetContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ContextName: contextName,
	}
}

func MakeDeleteContextInputOptions(version core.RuntimeVersion, contextName string) *DeleteContextInputOptions {
	return &DeleteContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ContextName: contextName,
	}
}

func MakeGetCurrentContextInputOptions(version core.RuntimeVersion, target Target) *GetCurrentContextInputOptions {
	return &GetCurrentContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Target: target,
	}
}

func MakeSetCurrentContextInputOptions(version core.RuntimeVersion, contextName string) *SetCurrentContextInputOptions {
	return &SetCurrentContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ContextName: contextName,
	}
}

func MakeRemoveCurrentContextInputOptions(version core.RuntimeVersion, target Target) *RemoveCurrentContextInputOptions {
	return &RemoveCurrentContextInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		Target: target,
	}
}

const (
	CtxCompatibilityOne string = "compatibility-one"
	CtxCompatibilityTwo string = "compatibility-two"
)
