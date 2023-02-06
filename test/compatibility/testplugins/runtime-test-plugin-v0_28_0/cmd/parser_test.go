// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
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
