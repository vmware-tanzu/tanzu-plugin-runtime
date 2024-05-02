// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestCliCmdSuite(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "config suite")
}

func TestSetGetDeleteContext(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	ctx1 := &configtypes.Context{
		Name:   "test1",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Path:                "test-path",
			Context:             "test-context",
			IsManagementCluster: true,
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "updated-test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
		},
		AdditionalMetadata: map[string]interface{}{
			"metaToken": "token1",
		},
	}

	ctx2 := &configtypes.Context{
		Name:   "test2",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Path:                "test-path",
			Context:             "test-context",
			IsManagementCluster: true,
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "updated-test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
		},
		AdditionalMetadata: map[string]interface{}{
			"metaToken": "token1",
		},
	}

	ctx, err := GetContext("test1")
	assert.Equal(t, "context test1 not found", err.Error())
	assert.Nil(t, ctx)

	err = SetContext(ctx1, true)
	assert.NoError(t, err)

	ctx, err = GetContext("test1")
	assert.Nil(t, err)
	assert.Equal(t, ctx1, ctx)

	ctx, err = GetCurrentContext(configtypes.TargetK8s)
	assert.Nil(t, err)
	assert.Equal(t, ctx1, ctx)

	err = SetContext(ctx2, false)
	assert.NoError(t, err)

	ctx, err = GetContext("test2")
	assert.Nil(t, err)
	assert.Equal(t, ctx2, ctx)

	ctx, err = GetCurrentContext(configtypes.TargetK8s)
	assert.Nil(t, err)
	assert.Equal(t, ctx1, ctx)

	err = DeleteContext("test")
	assert.Equal(t, "context test not found", err.Error())

	err = DeleteContext("test1")
	assert.Nil(t, err)

	ctx, err = GetContext("test1")
	assert.Nil(t, ctx)
	assert.Equal(t, "context test1 not found", err.Error())
}

func TestSetContextWithOldVersion(t *testing.T) {
	tanzuConfigBytes := `
currentContext:
    kubernetes: test-mc
contexts:
    - name: test-mc
      ctx-field: new-ctx-field
      optional: true
      target: kubernetes
      additionalMetadata:
        metaToken: token1
      clusterOpts:
        isManagementCluster: true
        endpoint: old-test-endpoint
        annotation: one
        required: true
        annotationStruct:
            one: one
      discoverySources:
        - gcp:
            name: test
            bucket: test-ctx-bucket
            manifestPath: test-ctx-manifest-path
            annotation: one
            required: true
        - gcp:
            name: test2
            bucket: test2-bucket
            manifestPath: test2-manifest-path
            annotation: one
            required: true
`

	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfgNextGen: tanzuConfigBytes})

	defer func() {
		cleanUp()
	}()

	ctx := &configtypes.Context{
		Name:   "test-mc",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Path:                "test-path",
			Context:             "test-context",
			IsManagementCluster: true,
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "updated-test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
		},
		AdditionalMetadata: map[string]interface{}{
			"metaToken": "token1",
		},
	}

	err := SetContext(ctx, false)
	assert.NoError(t, err)

	c, err := GetContext(ctx.Name)
	assert.NoError(t, err)
	assert.Equal(t, c.Name, ctx.Name)
	assert.Equal(t, c.ClusterOpts.Endpoint, "old-test-endpoint")
	assert.Equal(t, c.ClusterOpts.Path, ctx.ClusterOpts.Path)
	assert.Equal(t, c.ClusterOpts.Context, ctx.ClusterOpts.Context)
	assert.Equal(t, c.AdditionalMetadata, ctx.AdditionalMetadata)
}

func TestSetContextWithDiscoverySourceWithNewFields(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name    string
		src     *configtypes.ClientConfig
		ctx     *configtypes.Context
		current bool
		errStr  string
	}{
		{
			name: "should add new context with new discovery sources to empty client config",
			src:  &configtypes.ClientConfig{},
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
				DiscoverySources: []configtypes.PluginDiscovery{
					{
						GCP: &configtypes.GCPDiscovery{
							Name:         "test",
							Bucket:       "updated-test-bucket",
							ManifestPath: "test-manifest-path",
						},
					},
				},
			},
			current: true,
		},
		{
			name: "should update existing context",
			src: &configtypes.ClientConfig{
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
						DiscoverySources: []configtypes.PluginDiscovery{
							{
								GCP: &configtypes.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
							},
						},
					},
				},
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
						DiscoverySources: []configtypes.PluginDiscovery{
							{
								GCP: &configtypes.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
							},
						},
					},
				},
				CurrentServer: "test-mc",
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
				},
			},
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "updated-test-endpoint",
					Path:                "updated-test-path",
					Context:             "updated-test-context",
					IsManagementCluster: true,
				},
				DiscoverySources: []configtypes.PluginDiscovery{
					{
						GCP: &configtypes.GCPDiscovery{
							Name:         "test",
							Bucket:       "updated-test-bucket",
							ManifestPath: "updated-test-manifest-path",
						},
					},
				},
			},
			current: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SetContext(tc.ctx, tc.current)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			ok, err := ContextExists(tc.ctx.Name)
			assert.True(t, ok)
			assert.NoError(t, err)
		})
	}
}

func TestSetContextWithDiscoverySource(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name    string
		src     *configtypes.ClientConfig
		ctx     *configtypes.Context
		current bool
		errStr  string
	}{
		{
			name: "should add new context with new discovery sources to empty client config",
			src:  &configtypes.ClientConfig{},
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
				DiscoverySources: []configtypes.PluginDiscovery{
					{
						GCP: &configtypes.GCPDiscovery{
							Name:         "test",
							Bucket:       "updated-test-bucket",
							ManifestPath: "test-manifest-path",
						},
					},
				},
			},
			current: true,
		},
		{
			name: "should update existing context",
			src: &configtypes.ClientConfig{
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
						DiscoverySources: []configtypes.PluginDiscovery{
							{
								GCP: &configtypes.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
							},
						},
					},
				},
				KnownServers: []*configtypes.Server{
					{
						Name: "test-mc",
						Type: configtypes.ManagementClusterServerType,
						ManagementClusterOpts: &configtypes.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
						DiscoverySources: []configtypes.PluginDiscovery{
							{
								GCP: &configtypes.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
							},
						},
					},
				},
				CurrentServer: "test-mc",
				CurrentContext: map[configtypes.ContextType]string{
					configtypes.ContextTypeK8s: "test-mc",
				},
			},
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "updated-test-endpoint",
					Path:                "updated-test-path",
					Context:             "updated-test-context",
					IsManagementCluster: true,
				},
				DiscoverySources: []configtypes.PluginDiscovery{
					{
						GCP: &configtypes.GCPDiscovery{
							Name:         "test",
							Bucket:       "updated-test-bucket",
							ManifestPath: "updated-test-manifest-path",
						},
					},
				},
			},
			current: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := SetContext(tc.ctx, tc.current)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			ok, err := ContextExists(tc.ctx.Name)
			assert.True(t, ok)
			assert.NoError(t, err)
		})
	}
}

func setupForGetContext() error {
	// setup
	cfg := &configtypes.ClientConfig{
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
				Name:   "test-mc-2",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint-2",
					Path:                "test-path-2",
					Context:             "test-context-2",
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
			{
				Name:        "k8s-context",
				ContextType: configtypes.ContextTypeK8s,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "https://api.tanzu.cloud.vmware.com:443/org/fake-org-id",
					Path:     "test-path",
					Context:  "test-context",
				},
			},
			{
				Name:        "k8s-context",
				ContextType: configtypes.ContextTypeK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "https://api.tanzu.cloud.vmware.com:443/org/fake-org-id",
					Path:     "test-path",
					Context:  "test-context",
				},
			},
			{
				Name:        "test-tanzu",
				ContextType: configtypes.ContextTypeTanzu,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "https://api.tanzu.cloud.vmware.com:443/org/fake-org-id",
					Path:     "test-path",
					Context:  "test-context",
				},
				AdditionalMetadata: map[string]interface{}{
					OrgIDKey:   "fake-org-id",
					OrgNameKey: "fake-org-name",
				},
			},
		},
		CurrentContext: map[configtypes.ContextType]string{
			configtypes.ContextTypeK8s: "test-mc-2",
			configtypes.ContextTypeTMC: "test-tmc",
		},
	}
	return func() error {
		LocalDirName = TestLocalDirName
		err := StoreClientConfig(cfg)
		return err
	}()
}

func TestGetContext(t *testing.T) {
	err := setupForGetContext()
	assert.NoError(t, err)

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tcs := []struct {
		name    string
		ctxName string
		errStr  string
	}{
		{
			name:    "success k8s",
			ctxName: "test-mc",
		},
		{
			name:    "success tmc",
			ctxName: "test-tmc",
		},
		{
			name:    "success tanzu",
			ctxName: "test-tanzu",
		},
		{
			name:    "failure",
			ctxName: "test",
			errStr:  "context test not found",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c, err := GetContext(tc.ctxName)
			if tc.errStr == "" {
				assert.Equal(t, tc.ctxName, c.Name)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
		})
	}
}

func TestGetContextsByType(t *testing.T) {
	err := setupForGetContext()
	assert.NoError(t, err)

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tcs := []struct {
		name                    string
		contextType             configtypes.ContextType
		expectedNamesOfContexts []string
		contextToDelete         string
	}{
		{
			name:                    "get k8s contexts",
			expectedNamesOfContexts: []string{"test-mc", "test-mc-2", "k8s-context"},
			contextType:             configtypes.ContextTypeK8s,
		},
		{
			name:                    "get tmc contexts",
			expectedNamesOfContexts: []string{"test-tmc"},
			contextType:             configtypes.ContextTypeTMC,
		},
		{
			name:                    "get tanzu contexts",
			expectedNamesOfContexts: []string{"test-tanzu"},
			contextType:             configtypes.ContextTypeTanzu,
		},
		{
			name:                    "get tanzu contexts",
			contextToDelete:         "test-tanzu",
			expectedNamesOfContexts: []string(nil),
			contextType:             configtypes.ContextTypeTanzu,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			var ctxNames []string

			if tc.contextToDelete != "" {
				err := DeleteContext(tc.contextToDelete)
				assert.Nil(t, err)
			}
			ctxs, err := GetContextsByType(tc.contextType)
			assert.Nil(t, err)

			for _, ctx := range ctxs {
				ctxNames = append(ctxNames, ctx.Name)
			}
			assert.Equal(t, ctxNames, tc.expectedNamesOfContexts)
		})
	}
}

func TestContextExists(t *testing.T) {
	err := setupForGetContext()
	assert.NoError(t, err)

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tcs := []struct {
		name    string
		ctxName string
		ok      bool
	}{
		{
			name:    "success k8s",
			ctxName: "test-mc",
			ok:      true,
		},
		{
			name:    "success tmc",
			ctxName: "test-tmc",
			ok:      true,
		},
		{
			name:    "success tanzu",
			ctxName: "test-tanzu",
			ok:      true,
		},
		{
			name:    "failure",
			ctxName: "test",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := ContextExists(tc.ctxName)
			assert.Equal(t, tc.ok, ok)
			assert.NoError(t, err)
		})
	}
}

func TestSetContext(t *testing.T) {
	// setup
	func() {
		LocalDirName = TestLocalDirName
		// setup data
		node := &yaml.Node{
			Kind: yaml.DocumentNode,
			Content: []*yaml.Node{
				{
					Kind:    yaml.MappingNode,
					Content: []*yaml.Node{},
				},
			},
		}
		err := persistConfig(node)
		assert.NoError(t, err)
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tcs := []struct {
		name               string
		ctx                *configtypes.Context
		current            bool
		errStr             string
		shouldServerExists bool
	}{
		{
			name: "should add new context and set current on empty config",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
				AdditionalMetadata: map[string]interface{}{
					"metaToken": "token1",
				},
			},
			current:            true,
			shouldServerExists: true,
		},

		{
			name: "should add new context but not current and configure missing ContextType from Target",
			ctx: &configtypes.Context{
				Name:   "test-mc2",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
				AdditionalMetadata: map[string]interface{}{
					"metaToken": "token1",
				},
			},
			shouldServerExists: true,
		},
		{
			name: "should add new context and configure missing Target from ContextType",
			ctx: &configtypes.Context{
				Name:        "test-mc2",
				ContextType: configtypes.ContextTypeK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
				AdditionalMetadata: map[string]interface{}{
					"metaToken": "token1",
				},
			},
			shouldServerExists: true,
		},
		{
			name: "success tmc current",
			ctx: &configtypes.Context{
				Name:   "test-tmc1",
				Target: configtypes.TargetTMC,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
			current:            true,
			shouldServerExists: true,
		},
		{
			name: "success tmc not_current",
			ctx: &configtypes.Context{
				Name:        "test-tmc2",
				Target:      configtypes.TargetTMC,
				ContextType: configtypes.ContextTypeTMC,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
			shouldServerExists: true,
		},
		{
			name: "success update test-mc",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "good-test-endpoint",
					Path:                "updated-test-path",
					Context:             "updated-test-context",
					IsManagementCluster: true,
				},
				AdditionalMetadata: map[string]interface{}{
					"metaToken": "updated-token1",
				},
			},
			shouldServerExists: true,
		},
		{
			name: "success update tmc",
			ctx: &configtypes.Context{
				Name:   "test-tmc",
				Target: configtypes.TargetTMC,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "updated-test-endpoint",
				},
			},
			shouldServerExists: true,
		},
		{
			name: "success tanzu current",
			ctx: &configtypes.Context{
				Name:        "test-tanzu1",
				ContextType: configtypes.ContextTypeTanzu,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "test-endpoint",
					Path:     "test-path",
					Context:  "test-context",
				},
				AdditionalMetadata: map[string]interface{}{
					"org": "fake-org-1",
				},
			},
			current:            true,
			shouldServerExists: false,
		},
		{
			name: "success tanzu not_current",
			ctx: &configtypes.Context{
				Name:        "test-tanzu2",
				ContextType: configtypes.ContextTypeTanzu,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "test-endpoint",
					Path:     "test-path",
					Context:  "test-context",
				},
				AdditionalMetadata: map[string]interface{}{
					"org": "fake-org-2",
				},
			},
			shouldServerExists: false,
		},
		{
			name: "error target and contexttype does not match",
			ctx: &configtypes.Context{
				Name:        "test-error",
				Target:      configtypes.TargetTMC,
				ContextType: configtypes.ContextTypeK8s,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint: "test-endpoint",
					Path:     "test-path",
					Context:  "test-context",
				},
			},
			errStr: "error while validating the Context object: specified Target(mission-control) and ContextType(kubernetes) for the Context object does not match",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// perform test
			err := SetContext(tc.ctx, tc.current)
			if tc.errStr != "" {
				assert.EqualError(t, err, tc.errStr)
			} else {
				assert.NoError(t, err)
				ctx, err := GetContext(tc.ctx.Name)
				assert.NoError(t, err)
				assert.Equal(t, tc.ctx.Name, ctx.Name)
				assert.NotEmpty(t, string(ctx.Target))
				assert.NotEmpty(t, string(ctx.ContextType))
				// Verify that even though only Target or ContextType was provided when
				// setting context, retrieving the Context should have both set
				assert.Equal(t, string(ctx.Target), string(ctx.ContextType))

				s, err := GetServer(tc.ctx.Name)
				// Verify that for some context types(e.g "tanzu") the server entry in not populated
				if !tc.shouldServerExists {
					assert.Nil(t, s)
					assert.EqualError(t, err, fmt.Sprintf("could not find server %q", tc.ctx.Name))
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tc.ctx.Name, s.Name)
				}
			}
		})
	}
}

func TestRemoveContext(t *testing.T) {
	// setup
	err := setupForGetContext()
	assert.NoError(t, err)
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tcs := []struct {
		name    string
		ctxName string
		errStr  string
	}{
		{
			name:    "success k8s",
			ctxName: "test-mc",
		},
		{
			name:    "success tmc",
			ctxName: "test-tmc",
		},
		{
			name:    "success tanzu",
			ctxName: "test-tanzu",
		},
		{
			name:    "failure",
			ctxName: "test",
			errStr:  "context test not found",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.errStr == "" {
				ok, err := ContextExists(tc.ctxName)
				require.True(t, ok)
				require.NoError(t, err)
			}
			err := RemoveContext(tc.ctxName)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ok, err := ContextExists(tc.ctxName)
			assert.False(t, ok)
			assert.NoError(t, err)
			ok, err = ServerExists(tc.ctxName)
			assert.Nil(t, err)
			assert.False(t, ok)
		})
	}
}

func TestGetAllCurrentContexts(t *testing.T) {
	// setup
	err := setupForGetContext()
	assert.NoError(t, err)
	defer func() {
		cleanupDir(LocalDirName)
	}()

	currentContextMap, err := GetAllCurrentContextsMap()
	assert.NoError(t, err)
	assert.Equal(t, "test-mc-2", currentContextMap[configtypes.TargetK8s].Name)
	assert.Equal(t, "test-tmc", currentContextMap[configtypes.TargetTMC].Name)
	assert.Nil(t, currentContextMap[configtypes.Target(configtypes.ContextTypeTanzu)])

	activeContextMap, err := GetAllActiveContextsMap()
	assert.NoError(t, err)
	assert.Equal(t, "test-mc-2", activeContextMap[configtypes.ContextTypeK8s].Name)
	assert.Equal(t, "test-tmc", activeContextMap[configtypes.ContextTypeTMC].Name)
	assert.Nil(t, activeContextMap[configtypes.ContextTypeTanzu])

	currentContextsList, err := GetAllCurrentContextsList()
	assert.NoError(t, err)
	assert.Contains(t, currentContextsList, "test-mc-2")
	assert.Contains(t, currentContextsList, "test-tmc")
	assert.NotContains(t, currentContextsList, "test-tanzu")

	activeContextsList, err := GetAllActiveContextsList()
	assert.NoError(t, err)
	assert.Contains(t, activeContextsList, "test-mc-2")
	assert.Contains(t, activeContextsList, "test-tmc")
	assert.NotContains(t, activeContextsList, "test-tanzu")

	// set the tanzu context (k8s and tanzu current contexts are mutual exclusive)
	err = SetCurrentContext("test-tanzu")
	assert.NoError(t, err)
	// GetAllCurrentContextsMap does not return tanzu context
	currentContextMap, err = GetAllCurrentContextsMap()
	assert.NoError(t, err)
	assert.Nil(t, currentContextMap[configtypes.TargetK8s])
	assert.Equal(t, "test-tmc", currentContextMap[configtypes.TargetTMC].Name)
	assert.Nil(t, currentContextMap[configtypes.Target(configtypes.ContextTypeTanzu)])
	// GetAllActiveContextsMap should return tanzu context and should match
	activeContextMap, err = GetAllActiveContextsMap()
	assert.NoError(t, err)
	assert.Nil(t, activeContextMap[configtypes.ContextTypeK8s])
	assert.Equal(t, "test-tmc", activeContextMap[configtypes.ContextTypeTMC].Name)
	assert.NotNil(t, activeContextMap[configtypes.ContextTypeTanzu])
	assert.Equal(t, "test-tanzu", activeContextMap[configtypes.ContextTypeTanzu].Name)

	currentContextsList, err = GetAllCurrentContextsList()
	assert.NoError(t, err)
	assert.NotContains(t, currentContextsList, "test-mc2")
	assert.Contains(t, currentContextsList, "test-tmc")
	assert.Contains(t, currentContextsList, "test-tanzu")

	// remove the tanzu current context
	err = RemoveCurrentContext(configtypes.Target(configtypes.ContextTypeTanzu))
	assert.NoError(t, err)
	currentContextMap, err = GetAllCurrentContextsMap()
	assert.NoError(t, err)
	assert.Nil(t, currentContextMap[configtypes.TargetK8s])
	assert.Equal(t, "test-tmc", currentContextMap[configtypes.TargetTMC].Name)
	assert.Nil(t, currentContextMap[configtypes.Target(configtypes.ContextTypeTanzu)])

	currentContextsList, err = GetAllCurrentContextsList()
	assert.NoError(t, err)
	assert.NotContains(t, currentContextsList, "test-mc")
	assert.Contains(t, currentContextsList, "test-tmc")
	assert.NotContains(t, currentContextsList, "test-tanzu")
}

func TestRemoveCurrentContext(t *testing.T) {
	// setup
	err := setupForGetContext()
	assert.NoError(t, err)
	defer func() {
		cleanupDir(LocalDirName)
	}()

	err = RemoveCurrentContext(configtypes.TargetK8s)
	assert.NoError(t, err)

	currCtx, err := GetCurrentContext(configtypes.TargetK8s)
	assert.Equal(t, "no current context set for type \"kubernetes\"", err.Error())
	assert.Nil(t, currCtx)

	currSrv, err := GetCurrentServer()
	assert.Equal(t, "current server \"\" not found in tanzu config", err.Error())
	assert.Nil(t, currSrv)

	currCtx, err = GetCurrentContext(configtypes.TargetTMC)
	assert.NoError(t, err)
	assert.Equal(t, currCtx.Name, "test-tmc")
}

func TestSetSingleContext(t *testing.T) {
	// setup
	func() {
		LocalDirName = TestLocalDirName
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tcs := []struct {
		name    string
		ctx     *configtypes.Context
		current bool
		errStr  string
	}{
		{
			name: "success k8s current",
			ctx: &configtypes.Context{
				Name:   "test-mc",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := SetContext(tc.ctx, tc.current)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
			ok, err := ContextExists(tc.ctx.Name)
			assert.True(t, ok)
			assert.NoError(t, err)
			ok, err = ServerExists(tc.ctx.Name)
			assert.True(t, ok)
			assert.NoError(t, err)
		})
	}
}

func TestSetContextMultiFile(t *testing.T) {
	configBytes, configNextGenBytes := setupMultiCfgData()
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfg: configBytes, cfgNextGen: configNextGenBytes})

	defer func() {
		cleanUp()
	}()

	ctx := &configtypes.Context{
		Name:        "test-mc",
		Target:      configtypes.TargetK8s,
		ContextType: configtypes.ContextTypeK8s,
		ClusterOpts: &configtypes.ClusterServer{
			IsManagementCluster: true,
			Endpoint:            "test-endpoint",
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test2",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
		},
	}

	ctx2 := &configtypes.Context{
		Name:   "test-mc2",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Endpoint: "updated-test-endpoint",
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name: "test",
				},
			},
			{
				GCP: &configtypes.GCPDiscovery{
					Name: "test2",
				},
			},
		},
	}

	expectedCtx2 := &configtypes.Context{
		Name:        "test-mc2",
		Target:      configtypes.TargetK8s,
		ContextType: configtypes.ContextTypeK8s,
		ClusterOpts: &configtypes.ClusterServer{
			IsManagementCluster: true,
			Endpoint:            "updated-test-endpoint",
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test2",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
		},
	}

	c, err := GetCurrentContext(configtypes.TargetK8s)
	assert.NoError(t, err)
	assert.Equal(t, ctx, c)

	c, err = GetContext("test-mc")
	assert.NoError(t, err)
	assert.Equal(t, ctx, c)

	err = SetContext(ctx2, true)
	assert.NoError(t, err)

	c, err = GetContext(ctx2.Name)
	assert.NoError(t, err)
	assert.Equal(t, expectedCtx2, c)
}

func TestSetContextMultiFileAndMigrateToNewConfig(t *testing.T) {
	configBytes, configNextGenBytes := setupMultiCfgData()

	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{cfg: configBytes, cfgNextGen: configNextGenBytes, cfgMetadata: setupConfigMetadataWithMigrateToNewConfig()})

	defer func() {
		cleanUp()
	}()

	ctx := &configtypes.Context{
		Name:        "test-mc",
		Target:      configtypes.TargetK8s,
		ContextType: configtypes.ContextTypeK8s,
		ClusterOpts: &configtypes.ClusterServer{
			IsManagementCluster: true,
			Endpoint:            "test-endpoint",
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test2",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
		},
	}

	ctx2 := &configtypes.Context{
		Name:   "test-mc2",
		Target: configtypes.TargetK8s,
		ClusterOpts: &configtypes.ClusterServer{
			Endpoint: "updated-test-endpoint",
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name: "test",
				},
			},
			{
				GCP: &configtypes.GCPDiscovery{
					Name: "test2",
				},
			},
		},
	}

	expectedCtx2 := &configtypes.Context{
		Name:        "test-mc2",
		Target:      configtypes.TargetK8s,
		ContextType: configtypes.ContextTypeK8s,
		ClusterOpts: &configtypes.ClusterServer{
			IsManagementCluster: true,
			Endpoint:            "updated-test-endpoint",
		},
		DiscoverySources: []configtypes.PluginDiscovery{
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
			{
				GCP: &configtypes.GCPDiscovery{
					Name:         "test2",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
			},
		},
	}

	c, err := GetCurrentContext(configtypes.TargetK8s)
	assert.NoError(t, err)
	assert.Equal(t, ctx, c)

	c, err = GetContext("test-mc")
	assert.NoError(t, err)
	assert.Equal(t, ctx, c)

	err = SetContext(ctx2, true)
	assert.NoError(t, err)

	c, err = GetContext(ctx2.Name)
	assert.NoError(t, err)
	assert.Equal(t, expectedCtx2, c)
}

func TestSetContextWithUniquePermissions(t *testing.T) {
	// setup
	func() {
		LocalDirName = TestLocalDirName
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	ctx := &configtypes.Context{
		Name:   "test-mc",
		Target: configtypes.TargetTMC,
		GlobalOpts: &configtypes.GlobalServer{
			Endpoint: "test-endpoint",
			Auth: configtypes.GlobalServerAuth{
				IDToken: "",
				Issuer:  "https://console-stg.cloud.vmware.com/csp/gateway/am/api",
				Permissions: []string{
					"external/25834195-19aa-4ffd-8933-f5f20094ab24/service:member",
					"csp:org_owner",
					"external/f52d39b0-c298-4adf-9c6f-0a4a07351cd7/service:admin",
					"csp:org_member",
					"external/f52d39b0-c298-4adf-9c6f-0a4a07351cd7/service:member",
				},
				RefreshToken: "XXX",
				Type:         "api-token",
				UserName:     "tanzu-core",
			},
		},
	}

	ctx2 := &configtypes.Context{
		Name:   "test-mc",
		Target: configtypes.TargetTMC,
		GlobalOpts: &configtypes.GlobalServer{
			Endpoint: "test-endpoint-updated",
			Auth: configtypes.GlobalServerAuth{
				IDToken: "",
				Issuer:  "https://console-stg.cloud.vmware.com/csp/gateway/am/api",
				Permissions: []string{
					"csp:org_member2",
					"external/f52d39b0-c298-4adf-9c6f-0a4a07351cd7/service:member",
				},
				RefreshToken: "XXX",
				Type:         "api-token",
				UserName:     "tanzu-core",
			},
		},
	}

	ctx3 := &configtypes.Context{
		Name:   "test-mc",
		Target: configtypes.TargetTMC,
		GlobalOpts: &configtypes.GlobalServer{
			Endpoint: "test-endpoint-updated3",
			Auth: configtypes.GlobalServerAuth{
				IDToken: "",
				Issuer:  "https://console-stg.cloud.vmware.com/csp/gateway/am/api",
				Permissions: []string{
					"external/25834195-19aa-4ffd-8933-f5f20094ab24/service:member",
					"csp:org_owner3",
				},
				RefreshToken: "XXX",
				Type:         "api-token",
				UserName:     "tanzu-core",
			},
		},
	}

	for i := 1; i <= 100; i++ {
		err := SetContext(ctx, true)
		assert.NoError(t, err)
		err = SetContext(ctx2, true)
		assert.NoError(t, err)
		err = SetContext(ctx3, true)
		assert.NoError(t, err)
		err = SetContext(ctx, true)
		assert.NoError(t, err)
	}

	c, err := GetContext("test-mc")
	assert.NoError(t, err)
	assert.Equal(t, 7, len(c.GlobalOpts.Auth.Permissions))

	s, err := GetServer("test-mc")
	assert.NoError(t, err)
	assert.Equal(t, 7, len(s.GlobalOpts.Auth.Permissions))
}

func TestSetContextWithEmptyName(t *testing.T) {
	// setup
	func() {
		LocalDirName = TestLocalDirName
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()
	tcs := []struct {
		name    string
		ctx     *configtypes.Context
		current bool
		errStr  string
	}{
		{
			name: "success  current",
			ctx: &configtypes.Context{
				Name:   "",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
			errStr: "error while validating the Context object: context name cannot be empty",
		},
		{
			name: "success re empty current",
			ctx: &configtypes.Context{
				Name:   "",
				Target: configtypes.TargetK8s,
				ClusterOpts: &configtypes.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
			errStr: "error while validating the Context object: context name cannot be empty",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := SetContext(tc.ctx, tc.current)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
		})
	}
}

func TestSetCurrentContext(t *testing.T) {
	// setup
	func() {
		err := setupForGetContext()
		assert.NoError(t, err)
	}()
	defer func() {
		cleanupDir(LocalDirName)
	}()

	err := SetCurrentContext("test-tanzu")
	assert.NoError(t, err)
	validateActiveContextV2(t, configtypes.ContextTypeTanzu, "test-tanzu", false, "")

	err = SetCurrentContext("test-mc")
	assert.NoError(t, err)
	validateActiveContextV2(t, configtypes.ContextTypeK8s, "test-mc", true, "test-mc")

	_, err = GetCurrentContext(configtypes.Target(configtypes.ContextTypeTanzu))
	assert.Error(t, err)
	assert.ErrorContains(t, err, `no current context set for type "tanzu"`)
}

var _ = Describe("testing SetCurrentContext & SetActiveContext", func() {
	var (
		err error
	)

	BeforeEach(func() {
		// setup
		err := setupForGetContext()
		gomega.Expect(err).To(gomega.BeNil())
	})
	AfterEach(func() {
		cleanupDir(LocalDirName)
	})

	Context("tmc context as current context", func() {
		It("should set tmc context as current context successfully", func() {
			err = SetCurrentContext("test-tmc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTMC, "test-tmc", true, "test-mc-2")
		})
		It("should set tmc context as active context successfully", func() {
			err = RemoveActiveContext(configtypes.ContextTypeTMC)
			gomega.Expect(err).To(gomega.BeNil())
			err = SetActiveContext("test-tmc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTMC, "test-tmc", true, "test-mc-2")
		})
	})
	Context("k8s context as current context", func() {
		It("should set k8s context as current context successfully", func() {
			err = SetCurrentContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")
		})
		It("should set k8s context as active context successfully", func() {
			err = RemoveActiveContext(configtypes.ContextTypeK8s)
			gomega.Expect(err).To(gomega.BeNil())
			err = SetActiveContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")
		})
	})
	Context("tanzu context as current context", func() {
		It("should set tanzu context as current context successfully", func() {
			//Remove the k8s current context set during initial setup
			err = RemoveCurrentContext(configtypes.TargetK8s)
			gomega.Expect(err).To(gomega.BeNil())

			err = SetCurrentContext("test-tanzu")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTanzu, "test-tanzu", false, "")

		})
		It("should set tanzu context as active context successfully", func() {
			//Remove the k8s current context set during initial setup
			err = RemoveActiveContext(configtypes.ContextTypeK8s)
			gomega.Expect(err).To(gomega.BeNil())

			err = SetActiveContext("test-tanzu")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTanzu, "test-tanzu", false, "")

		})
	})
	Context("k8s context as current context after tmc context ", func() {
		It("should have k8s and tmc contexts as current for their respective targets", func() {
			err = SetCurrentContext("test-tmc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTMC, "test-tmc", true, "test-mc-2")

			err = SetCurrentContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")

		})
		It("should have k8s and tmc contexts as active for their respective context types", func() {
			err = SetActiveContext("test-tmc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTMC, "test-tmc", true, "test-mc-2")

			err = SetActiveContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")

		})
	})
	Context("k8s context as current context after tanzu context(mutual-exclusion test between k8s and tanzu) ", func() {
		It("should have only k8s as current context and tanzu context should be removed from the current context", func() {
			err = SetCurrentContext("test-tanzu")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTanzu, "test-tanzu", false, "")

			err = SetCurrentContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")

			_, err = GetCurrentContext(configtypes.Target(configtypes.ContextTypeTanzu))
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.ContainSubstring(`no current context set for type "tanzu"`))
		})
		It("should have only k8s as current context and tanzu context should be removed from the active context", func() {
			err = SetActiveContext("test-tanzu")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTanzu, "test-tanzu", false, "")

			err = SetActiveContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")

			_, err = GetActiveContext(configtypes.ContextTypeTanzu)
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.ContainSubstring(`no current context set for type "tanzu"`))
		})
	})
	Context("tanzu context as current context after k8s context(mutual-exclusion test between k8s and tanzu) ", func() {
		It("should have only tanzu as current context and k8s context should be removed from the current context", func() {
			err = SetCurrentContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")

			err = SetCurrentContext("test-tanzu")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTanzu, "test-tanzu", false, "")

			_, err = GetCurrentContext(configtypes.TargetK8s)
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.ContainSubstring(`no current context set for type "kubernetes"`))
		})
		It("should have only tanzu as current context and k8s context should be removed from the active context", func() {
			err = SetActiveContext("test-mc")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeK8s, "test-mc", true, "test-mc")

			err = SetActiveContext("test-tanzu")
			gomega.Expect(err).To(gomega.BeNil())
			validateActiveContext(configtypes.ContextTypeTanzu, "test-tanzu", false, "")

			_, err = GetActiveContext(configtypes.ContextTypeK8s)
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.ContainSubstring(`no current context set for type "kubernetes"`))
		})
	})
})

func validateActiveContext(contextType configtypes.ContextType, ctxName string, isServerExpected bool, serverName string) {
	c, err := GetActiveContext(contextType)
	gomega.Expect(err).To(gomega.BeNil())
	gomega.Expect(c.Name).To(gomega.Equal(ctxName))

	server, err := GetCurrentServer()
	if !isServerExpected {
		gomega.Expect(err).ToNot(gomega.BeNil())
	} else {
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(server.Name).To(gomega.Equal(serverName))
	}
}

func validateActiveContextV2(t *testing.T, contextType configtypes.ContextType, ctxName string, isServerExpected bool, serverName string) {
	c, err := GetActiveContext(contextType)
	assert.NoError(t, err)
	assert.Equal(t, c.Name, ctxName)

	server, err := GetCurrentServer()
	if !isServerExpected {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, server.Name, serverName)
	}
}
