// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestGetEdition(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name   string
		in     *configtypes.ClientConfig
		out    string
		errStr string
	}{
		{
			name: "success k8s",
			in: &configtypes.ClientConfig{
				ClientOptions: &configtypes.ClientOptions{
					Env: map[string]string{
						"test": "test",
					},
				},
			},
			out: "test",
		},
	}
	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.in)
			assert.NoError(t, err)
			c, err := GetEnv("test")
			assert.NoError(t, err)
			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)
		})
	}
}

func TestSetEdition(t *testing.T) {
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
			assert.NoError(t, err)
			c, err := GetEdition()
			assert.Equal(t, spec.value, c)
			assert.NoError(t, err)
		})
	}
}
