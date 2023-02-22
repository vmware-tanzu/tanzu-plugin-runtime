package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidRuntimeVersion(t *testing.T) {
	version100 := &RuntimeAPIVersion{RuntimeVersion: "v1.0.0"}
	actual, err := ValidateRuntimeVersion(version100)
	assert.Nil(t, err)
	assert.Equal(t, true, actual)
}

func TestInvalidRuntimeVersion(t *testing.T) {
	version101 := &RuntimeAPIVersion{RuntimeVersion: "v1.0.1"}
	actual, err := ValidateRuntimeVersion(version101)
	assert.Equal(t, "runtime version v1.0.1 is not supported", err.Error())
	assert.Equal(t, false, actual)
}
