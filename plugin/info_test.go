// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/json"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
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
		Name:            "Test Plugin",
		Description:     "Description of the plugin",
		Version:         "1.2.3",
		BuildSHA:        "cafecafe",
		Group:           "TestGroup",
		DocURL:          "https://docs.example.com",
		Hidden:          false,
		PostInstallHook: func() error { return nil },
	}

	infoCmd := newInfoCmd(&descriptor)
	err = infoCmd.Execute()
	w.Close()
	assert.Nil(err)

	got := <-c

	expectedInfo := pluginInfo{
		PluginDescriptor: descriptor,
		BinaryArch:       runtime.GOARCH,
	}

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
	assert.Equal(expectedInfo.BinaryArch, gotInfo.BinaryArch)
	assert.Empty(gotInfo.PluginRuntimeVersion, "Should be empty since unit tests doesn't have the self (tanzu-plugin-runtime) module dependency")
}
