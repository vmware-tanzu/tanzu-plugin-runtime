// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewPlugin(t *testing.T) {
	assert := assert.New(t)

	descriptor := PluginDescriptor{
		Name:            "Test Plugin",
		Description:     "Description of the plugin",
		Version:         "v1.2.3",
		BuildSHA:        "cafecafe",
		Group:           "TestGroup",
		DocURL:          "https://docs.example.com",
		Hidden:          false,
		PostInstallHook: func() error { return nil },
	}

	cmd, err := NewPlugin(&descriptor)
	if err != nil {
		t.Error(err)
	}
	assert.Equal("Test Plugin", cmd.Cmd.Use)
	assert.Equal(("Description of the plugin"), cmd.Cmd.Short)
}

func TestAddCommands(t *testing.T) {
	assert := assert.New(t)

	descriptor := PluginDescriptor{
		Name:        "Test Plugin",
		Description: "Description of the plugin",
		Version:     "v1.2.3",
		BuildSHA:    "cafecafe",
		Group:       "TestGroup",
		DocURL:      "https://docs.example.com",
		Hidden:      false,
	}

	cmd, err := NewPlugin(&descriptor)
	if err != nil {
		t.Error(err)
	}

	subCmd := &cobra.Command{
		Use:   "Sub1",
		Short: "Sub1 description",
	}
	cmd.AddCommands(subCmd)

	// Plugin gets 6 commands by default (describe, info, version, lint, post-install, generate-docs), ours should make 7.
	assert.Equal(7, len(cmd.Cmd.Commands()))
}

func TestExecute(t *testing.T) {
	assert := assert.New(t)

	descriptor := PluginDescriptor{
		Name:        "Test Plugin",
		Description: "Description of the plugin",
		Version:     "v1.2.3",
		BuildSHA:    "cafecafe",
		Group:       "TestGroup",
		DocURL:      "https://docs.example.com",
		Hidden:      false,
	}

	cmd, err := NewPlugin(&descriptor)
	if err != nil {
		t.Error(err)
	}

	assert.Nil(cmd.Execute())
}
