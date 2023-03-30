// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCEIPOptIn(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "should persist ceipOptIn value as true when empty client config",
			value: "true",
		},
		{
			name:  "should update and persist eipOptIn value as false",
			value: "false",
		},
		{
			name:  "should not persist same value false",
			value: "false",
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCEIPOptIn(spec.value)
			assert.NoError(t, err)
			c, err := GetCEIPOptIn()
			assert.Equal(t, spec.value, c)
			assert.NoError(t, err)
		})
	}
}
