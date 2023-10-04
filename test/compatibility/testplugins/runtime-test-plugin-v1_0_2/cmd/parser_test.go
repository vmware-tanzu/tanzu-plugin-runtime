// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestParseContext(t *testing.T) {
	tests := []struct {
		name            string
		ctxStr          string
		err             string
		expectedContext *configtypes.Context
	}{
		{
			name: "Parse valid context str",
			ctxStr: `name: context-one
target: kubernetes
globalOpts:
  endpoint: test-endpoint
`,
			expectedContext: &configtypes.Context{
				Name:   "context-one",
				Target: "kubernetes",
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
		},
		{
			name:   "Failed to parse invalid string",
			ctxStr: `name`,
			err:    "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `name` into types.Context",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseContext(tt.ctxStr)
			if tt.err != "" || err != nil {
				assert.Equal(t, tt.err, err.Error())
			} else {
				assert.Equal(t, tt.expectedContext, actual)
			}
		})
	}
}

func TestParseServer(t *testing.T) {
	tests := []struct {
		name           string
		serverStr      string
		err            string
		expectedServer *configtypes.Server
	}{
		{
			name: "Parse valid context str",
			serverStr: `name: compatibility-test-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
			expectedServer: &configtypes.Server{
				Name: "compatibility-test-one",
				Type: configtypes.ManagementClusterServerType,
				GlobalOpts: &configtypes.GlobalServer{
					Endpoint: "default-compatibility-test-endpoint",
				},
			},
		},
		{
			name:      "Failed to parse invalid string",
			serverStr: `name`,
			err:       "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `name` into types.Server",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseServer(tt.serverStr)
			if tt.err != "" || err != nil {
				assert.Equal(t, tt.err, err.Error())
			} else {
				assert.Equal(t, tt.expectedServer, actual)
			}
		})
	}
}
