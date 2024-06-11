// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/json"
	"os"
	"runtime"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	c := make(chan []byte)
	go readOutput(t, r, c)

	// Set up for our test
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	descriptor := PluginDescriptor{
		Name:                 "Foo Plugin",
		Target:               types.TargetK8s,
		Description:          "Description of the plugin",
		Version:              "v1.2.3",
		BuildSHA:             "cafecafe",
		Group:                "TestGroup",
		DocURL:               "https://docs.example.com",
		Hidden:               false,
		SupportedContextType: []types.ContextType{types.ContextTypeTanzu},
		CommandMap: []CommandMapEntry{
			{
				SourceCommandPath:      "subber",
				DestinationCommandPath: "subber",
				Overrides:              "somecommand",
				RequiredContextType:    []types.ContextType{types.ContextTypeTanzu},
			},
		},
	}

	expectedInfo := pluginInfo{
		PluginDescriptor: descriptor,
		BinaryArch:       runtime.GOARCH,
	}

	p, err := NewPlugin(&descriptor)
	if err != nil {
		t.Error(err)
	}

	p.Cmd.SetArgs([]string{"info"})

	subCmd := &cobra.Command{
		Use:     "subber",
		Short:   "subcommand description",
		Aliases: []string{"sub"},
	}
	p.AddCommands(subCmd)

	err = p.Cmd.Execute()
	w.Close()
	assert.Nil(err)

	got := <-c

	gotInfo := &pluginInfo{}
	err = json.Unmarshal(got, gotInfo)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(expectedInfo.Name, gotInfo.Name)
	assert.Equal(expectedInfo.Description, gotInfo.Description)
	assert.Equal(expectedInfo.Version, gotInfo.Version)
	assert.Equal(expectedInfo.BuildSHA, gotInfo.BuildSHA)
	assert.Equal(expectedInfo.DocURL, gotInfo.DocURL)
	assert.Equal(expectedInfo.Hidden, gotInfo.Hidden)
	assert.Equal(expectedInfo.SupportedContextType, gotInfo.SupportedContextType)
	assert.Equal("somecommand", gotInfo.CommandMap[0].Overrides)
	assert.Equal([]types.ContextType{types.ContextTypeTanzu}, gotInfo.CommandMap[0].RequiredContextType)
	assert.Equal(subCmd.Aliases, gotInfo.CommandMap[0].Aliases)
	assert.Equal(subCmd.Short, gotInfo.CommandMap[0].Description)

	assert.Equal(expectedInfo.BinaryArch, gotInfo.BinaryArch)
	assert.Empty(gotInfo.PluginRuntimeVersion, "Should be empty since unit tests doesn't have the self (tanzu-plugin-runtime) module dependency")
}
