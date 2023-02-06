// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
)

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
