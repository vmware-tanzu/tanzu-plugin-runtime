// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestPopulateContexts(t *testing.T) {
	tcs := []struct {
		name  string
		ip    *configtypes.ClientConfig
		op    *configtypes.ClientConfig
		delta bool
	}{
		{
			name:  "empty",
			ip:    &configtypes.ClientConfig{},
			op:    &configtypes.ClientConfig{},
			delta: false,
		},
		{
			name: "no delta",
			ip: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-mc",
						Target: configtypes.TargetK8s,
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			op: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-mc",
						Target: configtypes.TargetK8s,
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			delta: false,
		},
		{
			name: "w/ delta",
			ip: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			op: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-mc",
						Target: configtypes.TargetK8s,
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			delta: true,
		},
		{
			name: "server of type 'tanzu' should be ignored for context population",
			ip: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-tanzu",
						Type: configtypes.ServerType("tanzu"),
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-tanzu-endpoint",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-tmc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			op: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-tanzu",
						Type: configtypes.ServerType("tanzu"),
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-tanzu-endpoint",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-tmc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			delta: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			delta := PopulateContexts(tc.ip)
			assert.Equal(t, tc.delta, delta)
			// ensure that the servers are not lost
			assert.Equal(t, len(tc.op.KnownServers), len(tc.ip.KnownServers))
			assert.Equal(t, tc.op.CurrentServer, tc.ip.CurrentServer)
			// ensure that the missing contexts are added
			assert.Equal(t, len(tc.op.KnownContexts), len(tc.ip.KnownContexts))
			assert.Equal(t, tc.op.CurrentContext, tc.ip.CurrentContext)
		})
	}
}

func TestPopulateServers(t *testing.T) {
	tcs := []struct {
		name string
		ip   *configtypes.ClientConfig
		op   *configtypes.ClientConfig
	}{
		{
			name: "empty",
			ip:   &configtypes.ClientConfig{},
			op:   &configtypes.ClientConfig{},
		},
		{
			name: "no delta",
			ip: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-mc",
						Target: configtypes.TargetK8s,
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			op: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-mc",
						Target: configtypes.TargetK8s,
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
		},
		{
			name: "w/ delta",
			ip: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-mc",
						Target: configtypes.TargetK8s,
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
			op: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-mc",
						Target: configtypes.TargetK8s,
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
					configtypes.ContextTypeTMC: "test-tmc",
				},
			},
		},
		{
			name: "context type of 'tanzu' should be ignored for server population",
			ip: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-tanzu",
						Target: configtypes.Target(configtypes.ContextTypeTanzu),
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test--tanzu-endpoint",
						},
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeTanzu: "test-tanzu",
					configtypes.ContextTypeTMC:   "test-tmc",
				},
			},
			op: &configtypes.ClientConfig{
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name: "test-tmc",
						Type: configtypes.GlobalServerType,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentServer: "test-mc",
				KnownContexts: []*configtypes.Context{
					{
						Name:   "test-tanzu",
						Target: configtypes.Target(configtypes.ContextTypeTanzu),
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test--tanzu-endpoint",
						},
						ClusterOpts: &configtypes.ClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
					},
					{
						Name:   "test-tmc",
						Target: configtypes.TargetTMC,
						GlobalOpts: &configtypes.GlobalServer{
							Endpoint: "test-endpoint",
						},
					},
				},
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeTanzu: "test-tanzu",
					configtypes.ContextTypeTMC:   "test-tmc",
				},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			populateServers(tc.ip)
			// ensure that the contexts are not lost
			assert.Equal(t, len(tc.op.KnownContexts), len(tc.ip.KnownContexts))
			assert.Equal(t, tc.op.CurrentContext, tc.ip.CurrentContext)
			// ensure that the missing servers are added
			assert.Equal(t, len(tc.op.KnownServers), len(tc.ip.KnownServers))
			assert.Equal(t, tc.op.CurrentServer, tc.ip.CurrentServer)
		})
	}
}
