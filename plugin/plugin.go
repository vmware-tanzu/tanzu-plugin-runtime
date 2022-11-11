// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/multierr"
	"golang.org/x/mod/semver"
)

// Plugin is a Tanzu CLI plugin.
type Plugin struct {
	Cmd *cobra.Command
}

// NewPlugin creates an instance of Plugin.
func NewPlugin(descriptor *PluginDescriptor) (*Plugin, error) {
	ApplyDefaultConfig(descriptor)
	err := ValidatePlugin(descriptor)
	if err != nil {
		return nil, err
	}
	p := &Plugin{
		Cmd: newRootCmd(descriptor),
	}
	// TODO(prkalle): To be uncommented when lint command is added along with the component package dependency
	// p.Cmd.AddCommand(lintCmd)
	p.Cmd.AddCommand(genDocsCmd)
	p.Cmd.AddCommand(newPostInstallCmd(descriptor))
	return p, nil
}

// AddCommands adds commands to the plugin.
func (p *Plugin) AddCommands(commands ...*cobra.Command) {
	p.Cmd.AddCommand(commands...)
}

// Execute executes the plugin.
func (p *Plugin) Execute() error {
	return p.Cmd.Execute()
}

// ApplyDefaultConfig applies default configurations to plugin descriptor.
func ApplyDefaultConfig(p *PluginDescriptor) {
	if p.PostInstallHook == nil {
		p.PostInstallHook = func() error {
			return nil
		}
	}
}

// ValidatePlugin validates the plugin descriptor.
func ValidatePlugin(p *PluginDescriptor) (err error) {
	// skip builder plugin for bootstrapping
	if p.Name == "builder" {
		return nil
	}
	if p.Name == "" {
		err = multierr.Append(err, fmt.Errorf("plugin %q name cannot be empty", p.Name))
	}
	if p.Version == "" {
		err = multierr.Append(err, fmt.Errorf("plugin %q version cannot be empty", p.Name))
	}
	if !semver.IsValid(p.Version) && p.Version != "dev" {
		err = multierr.Append(err, fmt.Errorf("version %q %q is not a valid semantic version", p.Name, p.Version))
	}
	if p.Description == "" {
		err = multierr.Append(err, fmt.Errorf("plugin %q description cannot be empty", p.Name))
	}
	if p.Group == "" {
		err = multierr.Append(err, fmt.Errorf("plugin %q group cannot be empty", p.Name))
	}
	return
}
