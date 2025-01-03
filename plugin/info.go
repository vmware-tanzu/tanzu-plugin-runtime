// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"encoding/json"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
)

const (
	PluginRuntimeModulePath = "github.com/vmware-tanzu/tanzu-plugin-runtime"
)

// pluginInfo describes a plugin information. This is a super set of PluginDescriptor
// It includes some additional metadata that plugin runtime configures
type pluginInfo struct {
	// PluginDescriptor describes a plugin binary.
	PluginDescriptor `json:",inline" yaml:",inline"`

	// PluginRuntimeVersion of the plugin. Must be a valid semantic version https://semver.org/
	// This version specifies the version of Plugin Runtime that was used to build the plugin
	PluginRuntimeVersion string `json:"pluginRuntimeVersion" yaml:"pluginRuntimeVersion"`

	// The machine architecture of the plugin binary.
	// This information can prove useful on Darwin (MacOS) ARM64 machine
	// which can also execute AMD64 binaries in the Rosetta emulator.
	BinaryArch string `json:"binaryArch" yaml:"binaryArch"`
}

func newInfoCmd(desc *PluginDescriptor) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "info",
		Short:  "Plugin info",
		Hidden: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			updateCommandLevelMapEntries(cmd, desc)
			pi := pluginInfo{
				PluginDescriptor:     *desc,
				PluginRuntimeVersion: getPluginRuntimeVersion(),
				BinaryArch:           runtime.GOARCH,
			}
			b, err := json.Marshal(pi)
			if err != nil {
				return err
			}
			fmt.Println(string(b))
			return nil
		},
	}

	return cmd
}

func getSourceCommandByPath(rootCmd *cobra.Command, cmdPath string) *cobra.Command {
	hierarchy := strings.Fields(cmdPath)
	cmd, _, err := rootCmd.Find(hierarchy)
	if err == nil {
		return cmd
	}
	return nil
}

func updateCommandLevelMapEntries(cmd *cobra.Command, desc *PluginDescriptor) {
	if desc == nil || len(desc.CommandMap) == 0 {
		return
	}

	rootCmd := cmd.Parent()

	for i := range desc.CommandMap {
		mapEntry := &desc.CommandMap[i]
		// not subcommand mapping
		if mapEntry.SourceCommandPath == "" {
			continue
		}

		srcCmd := getSourceCommandByPath(rootCmd, mapEntry.SourceCommandPath)
		if srcCmd != nil {
			if len(mapEntry.Aliases) == 0 {
				mapEntry.Aliases = srcCmd.Aliases
			}

			if mapEntry.Description == "" {
				mapEntry.Description = srcCmd.Short
			}
		}
	}
}

func getPluginRuntimeVersion() string {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		panic("Can't read BuildInfo")
	}

	for _, dep := range buildInfo.Deps {
		if dep.Path == PluginRuntimeModulePath {
			return dep.Version
		}
	}
	return ""
}
