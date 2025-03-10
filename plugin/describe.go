// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDescribeCmd(description string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "describe",
		Short:  "Describes the plugin",
		Hidden: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			fmt.Println(description)
			return nil
		},
	}

	return cmd
}
