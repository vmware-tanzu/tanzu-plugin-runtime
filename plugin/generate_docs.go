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

func cloneCommand(sourceCommand *cobra.Command) *cobra.Command {
	// TODO(vuil): before cloning the command, we need the inherited flags
	// (persistent flags from command's ancestors) to be captured in the
	// command because once the clone is remapped and has a different
	// ancestry, some of the inherited flags may not be dynamically
	// discoverable anymore. Somewhat dubiously, until there is a means to
	// better dictate what a cobra Command's inherited flags are, we rely
	// on the side-effect of InheritedFlags() to populate the discovered
	// inherited flags in Command.iFlags
	_ = sourceCommand.InheritedFlags()

	clone := *sourceCommand
	return &clone
}

func cloneChildCommands(parentCloneCmd *cobra.Command) {
	childCommands := parentCloneCmd.Commands()
	for _, child := range childCommands {
		newChild := cloneCommand(child)

		cloneChildCommands(newChild)
		childParent := child.Parent()
		// child's parent pointer is wrong (still point to childParent, the
		// source that parentCloneCmd is cloned from), so remove said child
		parentCloneCmd.RemoveCommand(child)
		// and add the properly updated clone of child instead
		parentCloneCmd.AddCommand(newChild)
		// however, the removal has the side-effect of wiping the child's
		// parent link, so remove then re-add back as child of the clone source
		childParent.RemoveCommand(child)
		childParent.AddCommand(child)

		// there is currently no way to obtain the current help command so
		// check the .Use field instead
		// failure to update the help command to the clone will result in md
		// generated for the clone (tanzu_xxx_help.md), something that is
		// extraneous but not incorrect.
		if child.Use == "help [command]" {
			parentCloneCmd.SetHelpCommand(newChild)
		}
	}
}

func deepCopy(cmd *cobra.Command) *cobra.Command {
	clone := cloneCommand(cmd)
	cloneChildCommands(clone)
	return clone
}

// rebuildTanzuCommandTree reconstructs the tanzu CLI's placement of the
// plugin's commands after accounting for the plugin descriptor's CommandMap.
// It is primarily used to prep a Command tree for cobra's doc generation APIs
// warning: commands could be mutated as a side effect
func rebuildTanzuCommandTree(tanzuCmd, pluginRootCmd *cobra.Command, desc *PluginDescriptor) {
	tanzuCmd.AddCommand(pluginRootCmd)

	// straightfoward when no command mapping to account for
	if desc == nil || len(desc.CommandMap) == 0 {
		return
	}

	// reconstruct command tree by relocating all mapped commands
	cmap := desc.CommandMap
	sort.Slice(cmap, func(i, j int) bool { return cmap[i].SourceCommandPath > cmap[j].SourceCommandPath })

	for _, v := range cmap {
		mapEntry := v
		cmd, _ := findSubCommandByHierarchy(pluginRootCmd, hierarchyFromPath(mapEntry.SourceCommandPath), matchOnCommandName)
		copyOfCmd := deepCopy(cmd)
		mapCommand(tanzuCmd, copyOfCmd, &mapEntry)
		if cmd == pluginRootCmd {
			// if this is a plugin level mapping, assume that the original
			// plugin's command tree should not be made available anymore
			cmd.Hidden = true
		}
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
			Use:               "tanzu",
			Short:             "The main Tanzu CLI",
			DisableAutoGenTag: true,
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
