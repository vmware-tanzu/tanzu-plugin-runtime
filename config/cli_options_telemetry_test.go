// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestSetGetTelemetryOptions(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	telOptions := &configtypes.TelemetryOptions{
		Source: "/fake/path",
	}

	telOptionsUpdate := &configtypes.TelemetryOptions{
		Source: "/fake/path/updated",
	}

	// get telemetry options when the config file is empty
	gotTelemetryoptions, err := GetCLITelemetryOptions()
	assert.Equal(t, "telemetry not found", err.Error())
	assert.Nil(t, gotTelemetryoptions)

	// When the telemetry fields are configured
	err = SetCLITelemetryOptions(telOptions)
	assert.NoError(t, err)

	gotTelemetryoptions, err = GetCLITelemetryOptions()
	assert.Nil(t, err)
	assert.Equal(t, telOptions, gotTelemetryoptions)

	// update telemetry options
	err = SetCLITelemetryOptions(telOptionsUpdate)
	assert.NoError(t, err)

	gotTelemetryoptions, err = GetCLITelemetryOptions()
	assert.Nil(t, err)
	assert.Equal(t, gotTelemetryoptions, telOptionsUpdate)

	// test configuring with nil
	err = SetCLITelemetryOptions(nil)
	assert.NoError(t, err)
}

func TestDeleteTelemetryOptions(t *testing.T) {
	// Setup config data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	telOptions := &configtypes.TelemetryOptions{
		Source: "/fake/path",
	}

	// delete telemetry options when the config file is empty should not return error
	err := DeleteTelemetryOptions()
	assert.NoError(t, err)

	// When the telemetry fields are configured, delete telemetry operation should delete the
	// telemetryOperations node, and it's child nodes
	err = SetCLITelemetryOptions(telOptions)
	assert.NoError(t, err)

	err = DeleteTelemetryOptions()
	assert.NoError(t, err)

	// When telemetryOptions are deleted the subsequent get operation should return error
	gotTelemetryoptions, err := GetCLITelemetryOptions()
	assert.Equal(t, "telemetry not found", err.Error())
	assert.Nil(t, gotTelemetryoptions)
}
