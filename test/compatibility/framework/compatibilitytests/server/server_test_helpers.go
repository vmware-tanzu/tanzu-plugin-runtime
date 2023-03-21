// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package server provides cross-version Server API compatibility tests
package server

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework/compatibilitytests/common"
)

const ServerNotFound = "current server \"\" not found in tanzu config"

// DefaultSetServerInputOptions helper method to construct SetServer API input options
func DefaultSetServerInputOptions(version core.RuntimeVersion, serverName string) *framework.SetServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.SetServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}

	return nil
}

// DefaultGetServerInputOptions helper method to construct GetServer API input options
func DefaultGetServerInputOptions(version core.RuntimeVersion, serverName string) *framework.GetServerInputOptions {
	return &framework.GetServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultGetServerOutputOptions helper method to construct GetServer API output options
//nolint: dupl
func DefaultGetServerOutputOptions(version core.RuntimeVersion, serverName string) *framework.GetServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &framework.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0254, core.Version0116:
		return &framework.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
		}
	}
	return nil
}

// DefaultGetServerOutputOptionsWithError helper method to construct GetServer API output options with error
//nolint: dupl
func DefaultGetServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *framework.GetServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.GetServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.VersionLatest,
			},
			Error: fmt.Sprintf("could not find server \"%v\"", serverName),
		}
	}
	return nil
}

// DefaultSetCurrentServerInputOptions helper method to construct SetCurrentServer API input options
func DefaultSetCurrentServerInputOptions(version core.RuntimeVersion, serverName string) *framework.SetCurrentServerInputOptions {
	return &framework.SetCurrentServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultGetCurrentServerInputOptions helper method to construct GetCurrentServer API input options
func DefaultGetCurrentServerInputOptions(version core.RuntimeVersion) *framework.GetCurrentServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.GetCurrentServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.VersionLatest,
			},
		}
	}
	return nil
}

// DefaultGetCurrentServerOutputOptions helper method to construct GetCurrentServer API output options
//nolint: dupl
func DefaultGetCurrentServerOutputOptions(version core.RuntimeVersion, serverName string) *framework.GetCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0254, core.Version0116:
		return &framework.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	case core.Version0280:
		return &framework.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: core.Version0280,
			},
			ServerOpts: &framework.ServerOpts{
				Name: serverName,
				Type: framework.ManagementClusterServerType,
				GlobalOpts: &framework.GlobalServerOpts{
					Endpoint: common.DefaultEndpoint,
				},
			},
			ValidationStrategy: core.ValidationStrategyStrict,
		}
	}
	return nil
}

// DefaultGetCurrentServerOutputOptionsWithError helper method to construct GetCurrentServer API output options with error
func DefaultGetCurrentServerOutputOptionsWithError(version core.RuntimeVersion) *framework.GetCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.GetCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: ServerNotFound,
		}
	}
	return nil
}

// DefaultRemoveCurrentServerInputOptions helper method to construct RemoveCurrentServer API input options
func DefaultRemoveCurrentServerInputOptions(version core.RuntimeVersion) *framework.RemoveCurrentServerInputOptions {
	switch version {
	case core.VersionLatest, core.Version0280:
		return &framework.RemoveCurrentServerInputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			ServerName: common.CtxCompatibilityOne,
		}
	}
	return nil
}

// DefaultRemoveCurrentServerOutputOptionsWithError helper method to construct RemoveCurrentServer API output option
func DefaultRemoveCurrentServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *framework.RemoveCurrentServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.RemoveCurrentServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", serverName),
		}
	}
	return nil
}

// DefaultDeleteServerInputOptions helper method to construct DeleteServer API input options
func DefaultDeleteServerInputOptions(version core.RuntimeVersion, serverName string) *framework.DeleteServerInputOptions {
	return &framework.DeleteServerInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ServerName: serverName,
	}
}

// DefaultDeleteServerOutputOptionsWithError helper method to construct DeleteServer API output options
//nolint: dupl
func DefaultDeleteServerOutputOptionsWithError(version core.RuntimeVersion, serverName string) *framework.DeleteServerOutputOptions {
	switch version {
	case core.VersionLatest, core.Version0280, core.Version0254, core.Version0116:
		return &framework.DeleteServerOutputOptions{
			RuntimeAPIVersion: &core.RuntimeAPIVersion{
				RuntimeVersion: version,
			},
			Error: fmt.Sprintf("context %v not found", serverName),
		}
	}
	return nil
}
