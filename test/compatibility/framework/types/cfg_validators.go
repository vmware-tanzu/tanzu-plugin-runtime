// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"github.com/pkg/errors"
)

func (opts *PluginDiscoveryOpts) ValidPluginDiscovery() error {
	configsSet := 0
	if opts.GCP != nil {
		configsSet++
	}
	if opts.OCI != nil {
		configsSet++
	}
	if opts.REST != nil {
		configsSet++
	}
	if opts.Kubernetes != nil {
		configsSet++
	}
	if opts.Local != nil {
		configsSet++
	}
	if configsSet != 1 {
		return errors.New("one discovery config should be set")
	}
	return nil
}

func (opts *PluginDiscoveryOpts) ValidContextType() error {
	if opts.ContextType == CtxTypeK8s || opts.ContextType == CtxTypeTMC {
		return nil
	}
	return fmt.Errorf("invalid context type for plugin discovery")
}

func (opts *ContextOpts) ShouldNotIncludeTarget() bool {
	return opts.Target == ""
}

func (opts *ContextOpts) ShouldNotIncludeContextType() bool {
	return opts.Type == ""
}

func (opts *ContextOpts) ValidName() bool {
	return opts.Name != ""
}

func (opts *ContextOpts) ValidTarget() bool {
	return opts.Target != "" && (opts.Target == TargetK8s || opts.Target == TargetTMC)
}

// ValidType validates legacy context type
func (opts *ContextOpts) ValidType() bool {
	return opts.Type != "" && (opts.Type == CtxTypeK8s || opts.Type == CtxTypeTMC)
}
func (opts *ContextOpts) ValidContextType() bool {
	return opts.ContextType != "" && (opts.ContextType == ContextTypeK8s || opts.ContextType == ContextTypeTMC || opts.ContextType == ContextTypeTanzu)
}

func (opts *ContextOpts) ValidGlobalOptsOrClusterOpts() bool {
	return (opts.GlobalOpts != nil && opts.GlobalOpts.Endpoint != "") || (opts.ClusterOpts != nil && opts.ClusterOpts.Endpoint != "")
}

func (opts *ContextOpts) ValidDiscoverySources() bool {
	return opts.DiscoverySources != nil || len(opts.DiscoverySources) == 0
}

func (opts *ServerOpts) ValidName() bool {
	return opts.Name != ""
}

func (opts *ServerOpts) ValidServerType() bool {
	return opts.Type != "" && (opts.Type == ManagementClusterServerType || opts.Type == GlobalServerType)
}

func (opts *ServerOpts) ValidGlobalOptsOrManagementClusterOpts() bool {
	return (opts.GlobalOpts != nil && opts.GlobalOpts.Endpoint != "") || (opts.ManagementClusterOpts != nil && opts.ManagementClusterOpts.Endpoint != "")
}

func (opts *ServerOpts) ValidDiscoverySources() bool {
	return opts.DiscoverySources != nil || len(opts.DiscoverySources) == 0
}
