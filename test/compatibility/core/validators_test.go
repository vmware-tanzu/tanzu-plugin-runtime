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

func TestValidRuntimeVersionV090(t *testing.T) {
	version090 := &RuntimeAPIVersion{RuntimeVersion: "v0.90.0"}
	actual, err := version090.Validate()
	assert.Nil(t, err)
	assert.Equal(t, true, actual)
}

func TestInvalidRuntimeVersion(t *testing.T) {
	Version103 := &RuntimeAPIVersion{RuntimeVersion: "v1.0.3"}
	actual, err := Version103.Validate()
	assert.Equal(t, "runtime version v1.0.3 is not supported", err.Error())
	assert.Equal(t, false, actual)
}
