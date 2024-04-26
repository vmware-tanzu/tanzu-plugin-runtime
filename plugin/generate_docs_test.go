// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

func TestGenDocsCmd(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "test",
		Short: aurora.Bold(`Test plugin command`).String(),
	}

	descriptor := PluginDescriptor{
		Name:        "test",
		Target:      types.TargetGlobal,
		Description: "test plugin",
		Version:     "v1.2.3",
		BuildSHA:    "cafecafe",
		Group:       "TestGroup",
		DocURL:      "https://docs.example.com",
		Hidden:      false,
	}

	assert := assert.New(t)
	docsDir, err := os.MkdirTemp("", "docs-gen")
	assert.Nil(err)
	defer os.RemoveAll(docsDir)

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

	docsCmd := newGenDocsCmd(&descriptor)
	cmd.AddCommand(docsCmd)
	args := []string{}
	args = append(args, "generate-docs")
	args = append(args, fmt.Sprintf("--docs-dir=%s", docsDir))
	cmd.SetArgs(args)
	err = docsCmd.Execute()
	assert.Nil(err)
	w.Close()

	got := <-c
	assert.Equal("", string(got))

	assert.Nil(checkDirectoryHasMDFiles(docsDir))
}

func checkDirectoryHasMDFiles(dirName string) error {
	filesInfos, err := os.ReadDir(dirName)
	if err != nil {
		return errors.Errorf("failed to read the files in dir '%s", dirName)
	}
	for _, fileInfo := range filesInfos {
		if fileInfo.IsDir() {
			continue
		}
		matched, err := filepath.Match("*.md", fileInfo.Name())
		if err != nil {
			return errors.Errorf("file matching error for file '%s'", fileInfo.Name())
		}
		if matched {
			return nil
		}
	}
	return errors.Errorf("directory is expected to have MD files")
}
