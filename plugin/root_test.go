// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

// TestDeadcodeElimination checks that a simple program using the tanzu-plugin-runtime
// is linked taking full advantage of the linker's deadcode elimination step.
//
// If reflect.Value.MethodByName/reflect.Value.Method are reachable the
// linker will not always be able to prove that exported methods are
// unreachable, making deadcode elimination less effective. Using
// text/template and html/template makes reflect.Value.MethodByName
// reachable.
//
// This test checks that those function can be proven to be unreachable by
// the linker.
//
// Taken from https://github.com/spf13/cobra/blob/f98cf4216d3cb5235e6e0cd00ee00959deb1dc65/cobra_test.go#L245
// See also: https://github.com/spf13/cobra/pull/1956
func TestDeadcodeElimination(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("go tool nm fails on windows")
	}

	// check that a simple program using tanzu-plugin-runtime is
	// linked with deadcode elimination enabled.
	const (
		dirname  = "test_deadcode"
		progname = "test_deadcode_elimination"
	)
	_ = os.Mkdir(dirname, 0770)
	defer os.RemoveAll(dirname)
	filename := filepath.Join(dirname, progname+".go")
	err := os.WriteFile(filename, []byte(`package main

import "github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"

func main() {
	p, _ := plugin.NewPlugin(&plugin.PluginDescriptor{
		Name:        "deadcode-check",
		Description: "some desc",
		Group:       "Manage",
		Target:      "global",
		Version:     "v0.0.0",
	})
	_ = p.Execute()
}
`), 0600)
	if err != nil {
		t.Fatalf("could not write test program: %v", err)
	}
	buf, err := exec.Command("go", "build", filename).CombinedOutput()
	if err != nil {
		t.Fatalf("could not compile test program: %s", string(buf))
	}
	defer os.Remove(progname)
	buf, err = exec.Command("go", "tool", "nm", progname).CombinedOutput()
	if err != nil {
		t.Fatalf("could not run go tool nm: %v", err)
	}
	if strings.Contains(string(buf), "MethodByName") {
		t.Error("compiled programs contains MethodByName symbol")
	}
}
