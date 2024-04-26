// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package plugin provides functions to create new CLI plugins.
package plugin

import (
	"fmt"
	"os"
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

func generateDocsBasic(cmd *cobra.Command, args []string) error {
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
		Use: "tanzu",
	}
	// Necessary to generate correct output
	tanzuCmd.AddCommand(cmd.Parent())
	if err := doc.GenMarkdownTreeCustom(cmd.Parent(), docsDir, emptyStr, identity); err != nil {
		return fmt.Errorf("error generating docs %q", err)
	}

	return nil
}

func getGenDocFn(desc *PluginDescriptor) func(cmd *cobra.Command, args []string) error {
	return generateDocsBasic
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
