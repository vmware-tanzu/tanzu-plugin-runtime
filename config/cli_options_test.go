// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEdition(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name         string
		value        string
		errStr       string
		errStrForGet string
	}{
		{
			name:         "should return error for empty val",
			value:        "",
			errStr:       "value cannot be empty",
			errStrForGet: "edition not found",
		},
		{
			name:  "should persist tanzu when empty client config",
			value: "tanzu",
		},
		{
			name:  "should update and persist update-tanzu",
			value: "update-tanzu",
		},
		{
			name:  "should not persist same value update-tanzu",
			value: "update-tanzu",
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetEdition(spec.value)
			if spec.errStr != "" {
				assert.Equal(t, spec.errStr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			c, err := GetEdition()
			if spec.errStrForGet != "" {
				assert.Equal(t, spec.errStrForGet, err.Error())
			} else {
				assert.Equal(t, spec.value, c)
				assert.NoError(t, err)
			}
		})
	}
}
