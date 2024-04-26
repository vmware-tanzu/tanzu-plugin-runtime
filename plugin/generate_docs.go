// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package plugin provides functions to create new CLI plugins.
package plugin

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// DefaultDocsDir is the base docs directory
const DefaultDocsDir = "docs/cli/commands"
const ErrorDocsOutputFolderNotExists = "error reading docs output directory '%v', make sure directory exists or provide docs output directory as input value to '--docs-dir' flag"

var (
	docsDir string
)

func hierarchyFromPath(cmdPath string) []string {
	return strings.Fields(cmdPath)
}

// mapCommand inserts a subcommand at the position of the tanzu command tree as
// specified by mapEntry, ensuring that any ancestor commands are created as
// well, if necessary.
func mapCommand(tanzuCmd, subCommand *cobra.Command, mapEntry *CommandMapEntry) {
	if mapEntry == nil || mapEntry.DestinationCommandPath == "" {
		return
	}

	dstHierarchy := hierarchyFromPath(mapEntry.DestinationCommandPath)
	numParts := len(dstHierarchy)

	for i := 1; i <= numParts; i++ {
		cmd, cmdParent := findSubCommandByHierarchy(tanzuCmd, dstHierarchy[:i], matchOnCommandName)
		if cmdParent == nil {
			break
		}

		if cmd == nil {
			// dealing with actual command now
			if i == numParts {
				cmd = subCommand
				cmd.Hidden = false
				if mapEntry.Description != "" {
					cmd.Short = mapEntry.Description
				}
				if len(mapEntry.Aliases) != 0 {
					cmd.Aliases = append([]string{dstHierarchy[numParts-1]}, mapEntry.Aliases...)
				}
				subCommand.Use = dstHierarchy[numParts-1]
			} else {
				// create missing intermediate command
				cmd = &cobra.Command{
					Use: dstHierarchy[i-1],
				}
			}
			cmdParent.AddCommand(cmd)
		}
	}
}

// rebuildTanzuCommandTree reconstructs the tanzu CLI's placement of the
// plugin's commands after accounting for the plugin descriptor's CommandMap.
// It is primarily used to prep a Command tree for cobra's doc generation APIs
// warning: commands will be mutated as a side effect
func rebuildTanzuCommandTree(tanzuCmd, pluginRootCmd *cobra.Command, desc *PluginDescriptor) {
	tanzuCmd.AddCommand(pluginRootCmd)

	// straightfoward when no command mapping to account for
	if desc == nil || len(desc.CommandMap) == 0 {
		return
	}

	// reconstruct command tree by relocating all mapped commands, then
	cmap := desc.CommandMap
	sort.Slice(cmap, func(i, j int) bool { return cmap[i].SourceCommandPath > cmap[j].SourceCommandPath })

	for _, v := range cmap {
		mapEntry := v
		cmd, _ := findSubCommandByHierarchy(pluginRootCmd, hierarchyFromPath(mapEntry.SourceCommandPath), matchOnCommandName)
		mapCommand(tanzuCmd, cmd, &mapEntry)
	}
}

func getGenDocFn(desc *PluginDescriptor) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if docsDir == "" {
			docsDir = DefaultDocsDir
		}
		if dir, err := os.Stat(docsDir); err != nil || !dir.IsDir() {
			return errors.Wrap(err, fmt.Sprintf(ErrorDocsOutputFolderNotExists, docsDir))
		}
		identity := func(s string) string {
			if !strings.HasPrefix(s, "tanzu") {
				return fmt.Sprintf("tanzu_%s", s)
			}
			return s
		}
		emptyStr := func(s string) string { return "" }

		tanzuCmd := cobra.Command{
			Use:   "tanzu",
			Short: "The main Tanzu CLI",
		}

		rebuildTanzuCommandTree(&tanzuCmd, cmd.Parent(), desc)

		if err := doc.GenMarkdownTreeCustom(&tanzuCmd, docsDir, emptyStr, identity); err != nil {
			return fmt.Errorf("error generating docs %q", err)
		}

		return nil
	}
}

func newGenDocsCmd(desc *PluginDescriptor) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "generate-docs",
		Short:  "Generate Cobra CLI docs for all subcommands",
		Hidden: true,
		RunE:   getGenDocFn(desc),
	}
	cmd.Flags().StringVarP(&docsDir, "docs-dir", "d", DefaultDocsDir, "destination for docs output")

	return cmd
}
