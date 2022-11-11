// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"github.com/spf13/cobra"
)

func newRootCmd(descriptor *PluginDescriptor) *cobra.Command {
	cmd := &cobra.Command{
		Use:     descriptor.Name,
		Short:   descriptor.Description,
		Aliases: descriptor.Aliases,
	}
	// TODO(prkalle): To be uncommented when CmdTemplate/TemplateFunc(usage.go) are added along with the component package dependency
	// cobra.AddTemplateFuncs(TemplateFuncs)
	// cmd.SetUsageTemplate(CmdTemplate)

	cmd.AddCommand(
		newDescribeCmd(descriptor.Description),
		newVersionCmd(descriptor.Version),
		newInfoCmd(descriptor),
	)

	return cmd
}
