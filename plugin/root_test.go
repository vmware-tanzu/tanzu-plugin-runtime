// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newRootCmd(t *testing.T) {
	assert := assert.New(t)

	descriptor := PluginDescriptor{
		Name:            "Test Plugin",
		Description:     "Description of the plugin",
		Version:         "1.2.3",
		BuildSHA:        "cafecafe",
		Group:           "TestGroup",
		DocURL:          "https://docs.example.com",
		Hidden:          false,
		PostInstallHook: func() error { return nil },
	}

	cmd := newRootCmd(&descriptor)
	assert.Equal("Test Plugin", cmd.Use)
	assert.Equal(("Description of the plugin"), cmd.Short)
}
