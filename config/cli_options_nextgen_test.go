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
			name:  "should update and persist ceipOptIn value as false",
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

func TestSetEULAStatus(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name        string
		value       EULAStatus
		expectError bool
	}{
		{
			name:        "should persist eulaStatus value as accepted when empty client config",
			value:       EULAStatusAccepted,
			expectError: false,
		},
		{
			name:        "should update and persist shown eulaStatus value",
			value:       EULAStatusShown,
			expectError: false,
		},
		{
			name:        "should update and persist unset eulaStatus value",
			value:       EULAStatusUnset,
			expectError: false,
		},
		{
			name:        "should error on invalid eulaStatus value",
			value:       EULAStatus("invalidinvalid"),
			expectError: true,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetEULAStatus(spec.value)
			if spec.expectError {
				assert.Equal(t, "invalid eula status", err.Error())
			} else {
				assert.NoError(t, err)
				val, err := GetEULAStatus()
				assert.Equal(t, spec.value, val)
				assert.NoError(t, err)
			}
		})
	}
}

func TestSetEULAAcceptedVersions(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name                  string
		acceptedVersionsToSet []string
		expectedVersions      []string
		expectSetError        bool
		expectGetError        bool
	}{
		{
			name:                  "should persist accepted versions on valid input",
			acceptedVersionsToSet: []string{"v1.0.0"},
			expectedVersions:      []string{"v1.0.0"},
		},
		{
			name:                  "should persist accepted versions on valid input in sorted semver order",
			acceptedVersionsToSet: []string{"v1.0.0", "v0.9.0", "v1.10.0", "v1.9.1"},
			expectedVersions:      []string{"v0.9.0", "v1.0.0", "v1.9.1", "v1.10.0"},
		},
		{
			name:                  "does not set accepted versions with given nil list",
			acceptedVersionsToSet: nil,
			expectedVersions:      nil,
			expectGetError:        true,
		},
		{
			name:                  "does not set accepted versions with given empty list",
			acceptedVersionsToSet: []string{},
			expectedVersions:      nil,
			expectGetError:        true,
		},
		{
			name:                  "error when version is not valid semver",
			acceptedVersionsToSet: []string{"1.0.0"},
			expectSetError:        true,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetEULAAcceptedVersions(spec.acceptedVersionsToSet)
			if spec.expectSetError {
				assert.Contains(t, err.Error(), "invalid eula version")
			} else {
				assert.NoError(t, err)
				vers, err := GetEULAAcceptedVersions()
				if !spec.expectGetError {
					assert.Equal(t, spec.expectedVersions, vers)
					assert.NoError(t, err)
				}
			}
		})
	}
}

func TestSetCLIId(t *testing.T) {
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
			name:  "should persist cliId value when empty client config",
			value: "fake-cli-id",
		},
		{
			name:  "should update and persist cliId value",
			value: "fake-cli-id-updated",
		},
		{
			name:  "should not persist same value false",
			value: "fake-cli-id-updated",
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetCLIId(spec.value)
			assert.NoError(t, err)
			c, err := GetCLIId()
			assert.Equal(t, spec.value, c)
			assert.NoError(t, err)
		})
	}
}
