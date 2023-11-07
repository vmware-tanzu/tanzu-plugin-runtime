// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructTestPluginCmd(t *testing.T) {
	tests := []struct {
		version   RuntimeVersion
		apis      []*API
		pluginCmd string
		err       string
	}{
		{
			version: VersionLatest,
			apis: []*API{
				{
					Name:      SetContextAPI,
					Version:   VersionLatest,
					Arguments: map[APIArgumentType]interface{}{},
					Output:    nil,
				},
			},
			pluginCmd: pluginLatest,
			err:       "",
		},
		{
			version: Version102,
			apis: []*API{
				{
					Name:      SetContextAPI,
					Version:   Version102,
					Arguments: map[APIArgumentType]interface{}{},
					Output:    nil,
				},
			},
			pluginCmd: pluginV102,
			err:       "",
		},
		{
			version: Version090,
			apis: []*API{
				{
					Name:      SetContextAPI,
					Version:   Version090,
					Arguments: map[APIArgumentType]interface{}{},
					Output:    nil,
				},
			},
			pluginCmd: pluginV090,
			err:       "",
		},
	}

	for _, tt := range tests {
		pCmd, err := ConstructTestPluginCmd(tt.version, tt.apis)
		if err != nil {
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.Contains(t, pCmd, tt.pluginCmd)
		}
	}
}
