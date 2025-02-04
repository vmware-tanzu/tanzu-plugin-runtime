// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package command

import (
	"context"

	"github.com/spf13/cobra"
)

func Sequence(items ...func(cmd *cobra.Command, args []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for i := range items {
			if err := items[i](cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}

func Visit(cmd *cobra.Command, f func(c *cobra.Command) error) error {
	err := f(cmd)
	if err != nil {
		return err
	}
	for _, c := range cmd.Commands() {
		err := Visit(c, f)
		if err != nil {
			return err
		}
	}
	return nil
}

type commandKey struct{}

func ContextWithCommand(ctx context.Context, cmd *cobra.Command) context.Context {
	return context.WithValue(ctx, commandKey{}, cmd)
}

func CommandFromContext(ctx context.Context) *cobra.Command {
	if cmd, ok := ctx.Value(commandKey{}).(*cobra.Command); ok {
		return cmd
	}
	return nil
}
