// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configapi "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
)

func TestParseContext(t *testing.T) {
	tests := []struct {
		name            string
		ctxStr          string
		err             string
		expectedContext *configapi.Context
	}{
		{
			name: "Parse valid context str",
			ctxStr: `name: context-one
type: k8s
globalOpts:
  endpoint: test-endpoint
`,
			expectedContext: &configapi.Context{
				Name: "context-one",
				Type: "k8s",
				GlobalOpts: &configapi.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
		},
		{
			name:   "Failed to parse invalid string",
			ctxStr: `name`,
			err:    "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `name` into v1alpha1.Context",
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
		expectedServer *configapi.Server
	}{
		{
			name: "Parse valid context str",
			serverStr: `name: compatibility-test-one
type: managementcluster
globalOpts:
    endpoint: default-compatibility-test-endpoint
`,
			expectedServer: &configapi.Server{
				Name: "compatibility-test-one",
				Type: configapi.ManagementClusterServerType,
				GlobalOpts: &configapi.GlobalServer{
					Endpoint: "default-compatibility-test-endpoint",
				},
			},
		},
		{
			name:      "Failed to parse invalid string",
			serverStr: `name`,
			err:       "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `name` into v1alpha1.Server",
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
