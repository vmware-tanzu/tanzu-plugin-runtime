// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func setupBasicCommand(cmdName string) *cobra.Command {
	return &cobra.Command{
		Use:   cmdName,
		Short: fmt.Sprintf("%s command", cmdName),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s : %s", cmd.Use, args[0])
		},
		Example: fmt.Sprintf("example usage of the %s command", cmdName),
	}
}

// geneate a Plugin with with customizable command map and command visibility
//
// plugin root command
//   - foo
//   - bar
//   - deeper
//     -- baz
//
// with supplied command map and visibility list affect how and if each command is invocable
func genDocsTestPlugin(t *testing.T, mapEntries []CommandMapEntry, hiddenCommands []string) *Plugin {
	var fooCmd = setupBasicCommand("foo")
	var barCmd = setupBasicCommand("bar")
	var bazCmd = setupBasicCommand("baz")
	var deeperCmd = &cobra.Command{
		Use:   "deeper",
		Short: "deeper node",
	}
	deeperCmd.AddCommand(bazCmd)

	var descriptor = PluginDescriptor{
		Name:        "plug",
		Target:      types.TargetGlobal,
		Aliases:     []string{"t"},
		Description: "Test the CLI generate-docs",
		Group:       AdminCmdGroup,
		Version:     "v1.1.0",
		BuildSHA:    "1234567",
	}

	var value string
	var bvalue string
	fooCmd.Flags().StringVarP(&value, "value", "v", "", "value to pass")
	bazCmd.Flags().StringVarP(&bvalue, "bvalue", "b", "", "bvalue to pass")

	descriptor.CommandMap = mapEntries
	p, err := NewPlugin(&descriptor)
	assert.Nil(t, err)

	p.AddCommands(
		fooCmd,
		barCmd,
		deeperCmd,
	)

	for _, commandPath := range hiddenCommands {
		hierarchy := strings.Fields(commandPath)
		cmd, _ := findSubCommandByHierarchy(p.Cmd, hierarchy, matchOnCommandName)
		assert.NotNil(t, cmd)
		cmd.Hidden = true
	}

	return p
}

type fileContent struct {
	contains []string
	omits    []string
}

type markdownState struct {
	expectedFiles []string
	filesContent  map[string]fileContent
}

func checkMarkdownState(t *testing.T, docsDir string, ms markdownState) {
	filesInfos, err := os.ReadDir(docsDir)
	errorMsg := fmt.Sprintf("will processing doc output in %s\n", docsDir)

	assert.Nil(t, err, errorMsg)

	var foundFiles []string
	for _, fileInfo := range filesInfos {
		foundFiles = append(foundFiles, fileInfo.Name())
	}
	assert.ElementsMatch(t, ms.expectedFiles, foundFiles, errorMsg)

	for _, fileName := range foundFiles {
		b, err := os.ReadFile(filepath.Join(docsDir, fileName))
		assert.Nil(t, err, errorMsg)
		content := string(b)
		for _, expectedString := range ms.filesContent[fileName].contains {
			assert.Contains(t, content, expectedString, errorMsg)
		}
		for _, unexpectedString := range ms.filesContent[fileName].omits {
			assert.NotContains(t, content, unexpectedString, errorMsg)
		}
	}
}

func TestGenerateDocsWithPlugin(t *testing.T) {
	tests := []struct {
		test           string
		commandMap     []CommandMapEntry
		hiddenCommands []string
		expected       markdownState
	}{
		{
			test:           "baseline:no command mapping and all commands visible",
			commandMap:     []CommandMapEntry{},
			hiddenCommands: []string{},
			expected: markdownState{
				expectedFiles: []string{"tanzu.md", "tanzu_plug.md", "tanzu_plug_bar.md", "tanzu_plug_foo.md", "tanzu_plug_deeper.md", "tanzu_plug_deeper_baz.md"},
				filesContent: map[string]fileContent{
					"tanzu.md": fileContent{
						contains: []string{"tanzu", "[tanzu plug](tanzu_plug.md)"},
					},
					"tanzu_plug.md": fileContent{
						contains: []string{
							"[tanzu plug foo](tanzu_plug_foo.md)",
							"[tanzu plug bar](tanzu_plug_bar.md)",
							"[tanzu plug deeper](tanzu_plug_deeper.md)"},
						omits: []string{"tanzu plug baz"},
					},
					"tanzu_plug_deeper.md": fileContent{
						contains: []string{
							"[tanzu plug](tanzu_plug.md)",
							"- Test the CLI generate-docs",
							"[tanzu plug deeper baz](tanzu_plug_deeper_baz.md)",
							"- baz command"},
					},
					"tanzu_plug_foo.md": fileContent{
						contains: []string{
							"foo command",
							"tanzu plug foo [flags]",
							"example usage of the foo command",
							"-v, --value string",
							"### SEE ALSO",
							"[tanzu plug](tanzu_plug.md)"},
						omits: []string{"tanzu plug baz"},
					},
					"tanzu_plug_deep_baz.md": fileContent{
						contains: []string{
							"baz command",
							"tanzu plug deeper baz [flags]",
							"example usage of the baz command",
							"-b, -b-value string",
							"### SEE ALSO",
							"[tanzu plug deeper](tanzu_plug_deeper.md)"},
						omits: []string{"tanzu plug baz"},
					},
				},
			},
		},
		{
			test:           "base: no command mapping some hidden commands",
			commandMap:     []CommandMapEntry{},
			hiddenCommands: []string{"bar", "deeper baz"},
			expected: markdownState{
				expectedFiles: []string{"tanzu.md", "tanzu_plug.md", "tanzu_plug_foo.md"},
				filesContent: map[string]fileContent{
					"tanzu.md": fileContent{
						contains: []string{"tanzu", "[tanzu plug](tanzu_plug.md)"},
					},
					"tanzu_plug.md": fileContent{
						contains: []string{"[tanzu plug foo](tanzu_plug_foo.md)"},
						omits: []string{
							"tanzu plug baz",
							"[tanzu plug bar](tanzu_plug_bar.md)",
							"[tanzu plug deeper](tanzu_plug_deeper.md)"},
					},
					"tanzu_plug_foo.md": fileContent{
						contains: []string{"foo command"},
						omits:    []string{"tanzu plug baz"},
					},
				},
			},
		},
		{
			test: "top level plugin command mapped to top level of CLI",
			commandMap: []CommandMapEntry{
				{
					SourceCommandPath:      "foo",
					DestinationCommandPath: "foo",
				},
			},
			hiddenCommands: []string{"foo", "deeper baz"},
			expected: markdownState{
				expectedFiles: []string{"tanzu.md", "tanzu_plug.md", "tanzu_foo.md", "tanzu_plug_bar.md"},
				filesContent: map[string]fileContent{
					"tanzu.md": fileContent{
						contains: []string{"tanzu", "[tanzu plug](tanzu_plug.md)", "[tanzu foo](tanzu_foo.md)"},
					},
					"tanzu_plug.md": fileContent{
						contains: []string{"[tanzu plug bar](tanzu_plug_bar.md)"},
						omits: []string{
							"tanzu plug deeper",
							"tanzu plug foo",
						},
					},
					"tanzu_foo.md": fileContent{
						contains: []string{"foo command", "[tanzu](tanzu.md)"},
						omits:    []string{"tanzu plug foo"},
					},
				},
			},
		},
		{
			test: "command and plugin level mapping",
			commandMap: []CommandMapEntry{
				{
					SourceCommandPath:      "",
					DestinationCommandPath: "pi",
				},
				{
					SourceCommandPath:      "foo",
					DestinationCommandPath: "foo",
				},
			},
			hiddenCommands: []string{"foo", "deeper baz"},
			expected: markdownState{
				expectedFiles: []string{"tanzu.md", "tanzu_pi.md", "tanzu_foo.md", "tanzu_pi_bar.md"},
				filesContent: map[string]fileContent{
					"tanzu.md": fileContent{
						contains: []string{"tanzu", "[tanzu pi](tanzu_pi.md)", "[tanzu foo](tanzu_foo.md)"},
					},
					"tanzu_pi.md": fileContent{
						contains: []string{"[tanzu pi bar](tanzu_pi_bar.md)"},
						omits: []string{
							"tanzu pi deeper",
							"tanzu pi foo",
						},
					},
					"tanzu_foo.md": fileContent{
						contains: []string{"foo command", "[tanzu](tanzu.md)"},
						omits:    []string{"tanzu pi foo"},
					},
				},
			},
		},
		{
			test: "plugin level mapping to multiple destination paths",
			commandMap: []CommandMapEntry{
				{
					SourceCommandPath:      "",
					DestinationCommandPath: "pi",
				},
				{
					SourceCommandPath:      "",
					DestinationCommandPath: "hi",
				},
			},
			hiddenCommands: []string{"foo", "deeper baz"},
			expected: markdownState{
				expectedFiles: []string{"tanzu.md", "tanzu_pi.md", "tanzu_hi.md", "tanzu_pi_bar.md", "tanzu_hi_bar.md"},
				filesContent: map[string]fileContent{
					"tanzu.md": fileContent{
						contains: []string{"tanzu", "[tanzu pi](tanzu_pi.md)", "[tanzu hi](tanzu_hi.md)"},
						omits:    []string{"plug"},
					},
					"tanzu_pi.md": fileContent{
						contains: []string{"[tanzu pi bar](tanzu_pi_bar.md)"},
						omits:    []string{"tanzu pi deeper", "tanzu pi foo"},
					},
					"tanzu_hi.md": fileContent{
						contains: []string{"[tanzu hi bar](tanzu_hi_bar.md)", "[tanzu](tanzu.md)"},
						omits:    []string{"tanzu pi"},
					},
				},
			},
		},
	}

	for _, spec := range tests {
		t.Run(spec.test, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "gendoctest")
			if err != nil {
				t.Fatalf("Failed to create tmpdir: %v", err)
			}

			p := genDocsTestPlugin(t, spec.commandMap, spec.hiddenCommands)

			p.Cmd.SetArgs([]string{"generate-docs", "--docs-dir", tmpDir})
			err = p.Execute()
			assert.Nil(t, err)

			checkMarkdownState(t, tmpDir, spec.expected)
			if !t.Failed() {
				os.RemoveAll(tmpDir)
			}
		})
	}
}
