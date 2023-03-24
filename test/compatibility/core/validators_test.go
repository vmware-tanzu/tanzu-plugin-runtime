// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidRuntimeVersion(t *testing.T) {
	versionLatest := &RuntimeAPIVersion{RuntimeVersion: "latest"}
	actual, err := versionLatest.Validate()
	assert.Nil(t, err)
	assert.Equal(t, true, actual)
}

func TestInvalidRuntimeVersion(t *testing.T) {
	version101 := &RuntimeAPIVersion{RuntimeVersion: "v1.0.1"}
	actual, err := version101.Validate()
	assert.Equal(t, "runtime version v1.0.1 is not supported", err.Error())
	assert.Equal(t, false, actual)
}
