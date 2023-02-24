package framework

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

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
