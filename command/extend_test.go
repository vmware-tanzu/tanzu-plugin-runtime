// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package command

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"
)

func TestSequence(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		items  []func(cmd *cobra.Command, args []string) error
		output string
		err    error
	}{{
		name: "empty",
	}, {
		name: "single item",
		args: []string{"a", "b", "c"},
		items: []func(cmd *cobra.Command, args []string) error{
			func(cmd *cobra.Command, args []string) error {
				fmt.Fprintf(cmd.OutOrStdout(), "step %v\n", args)
				return nil
			},
		},
		output: `
step [a b c]
`,
	}, {
		name: "multiple items",
		items: []func(cmd *cobra.Command, _ []string) error{
			func(cmd *cobra.Command, _ []string) error {
				fmt.Fprintln(cmd.OutOrStdout(), "step 1")
				return nil
			},
			func(cmd *cobra.Command, _ []string) error {
				fmt.Fprintln(cmd.OutOrStdout(), "step 2")
				return nil
			},
			func(cmd *cobra.Command, _ []string) error {
				fmt.Fprintln(cmd.OutOrStdout(), "step 3")
				return nil
			},
		},
		output: `
step 1
step 2
step 3
`,
	}, {
		name: "stops on error",
		items: []func(cmd *cobra.Command, _ []string) error{
			func(cmd *cobra.Command, _ []string) error {
				fmt.Fprintln(cmd.OutOrStdout(), "step 1")
				return nil
			},
			func(cmd *cobra.Command, _ []string) error {
				fmt.Fprintln(cmd.OutOrStdout(), "step 2")
				return fmt.Errorf("test error")
			},
			func(cmd *cobra.Command, _ []string) error {
				fmt.Fprintln(cmd.OutOrStdout(), "step 3")
				return nil
			},
		},
		output: `
step 1
step 2
`,
		err: fmt.Errorf("test error"),
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			cmd := &cobra.Command{}
			cmd.SetOutput(output)

			err := Sequence(test.items...)(cmd, test.args)

			if expected, actual := fmt.Sprintf("%s", test.err), fmt.Sprintf("%s", err); expected != actual {
				t.Errorf("Expected error %q, actually %q", expected, actual)
			}
			if diff := cmp.Diff(strings.TrimSpace(test.output), strings.TrimSpace(output.String())); diff != "" {
				t.Errorf("Unexpected output (-expected, +actual): %s", diff)
			}
		})
	}
}

func TestVisit(t *testing.T) {
	tests := []struct {
		name    string
		cmd     func() *cobra.Command
		visitor func(*cobra.Command) error
		err     error
	}{{
		name: "single command",
		cmd: func() *cobra.Command {
			return &cobra.Command{Use: "root"}
		},
		visitor: func(_ *cobra.Command) error {
			return nil
		},
	}, {
		name: "parent-child",
		cmd: func() *cobra.Command {
			root := &cobra.Command{Use: "root"}
			root.AddCommand(&cobra.Command{Use: "child"})
			return root
		},
		visitor: func(_ *cobra.Command) error {
			return nil
		},
	}, {
		name: "error",
		cmd: func() *cobra.Command {
			return &cobra.Command{Use: "root"}
		},
		visitor: func(cmd *cobra.Command) error {
			return fmt.Errorf("%s", cmd.Name())
		},
		err: fmt.Errorf("root"),
	}, {
		name: "child error",
		cmd: func() *cobra.Command {
			root := &cobra.Command{Use: "root"}
			root.AddCommand(&cobra.Command{Use: "child"})
			return root
		},
		visitor: func(cmd *cobra.Command) error {
			if cmd.Name() == "child" {
				return fmt.Errorf("%s", cmd.Name())
			}
			return nil
		},
		err: fmt.Errorf("child"),
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Visit(test.cmd(), test.visitor)
			if expected, actual := fmt.Sprintf("%s", test.err), fmt.Sprintf("%s", err); expected != actual {
				t.Errorf("Expected error %q, actually %q", expected, actual)
			}
		})
	}
}

func TestCommandFromContext_WithCommand(t *testing.T) {
	cmd := &cobra.Command{}
	parentCtx := context.Background()
	childCtx := ContextWithCommand(parentCtx, cmd)

	if expected, actual := (*cobra.Command)(nil), CommandFromContext(parentCtx); expected != actual {
		t.Errorf("expected command %v, actually %v", expected, actual)
	}
	if expected, actual := cmd, CommandFromContext(childCtx); expected != actual {
		t.Errorf("expected command %v, actually %v", expected, actual)
	}
}
