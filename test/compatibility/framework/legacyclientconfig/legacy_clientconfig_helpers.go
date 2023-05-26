// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package legacyclientconfig

import (
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// DefaultGetClientConfigCommand creates a GetClientConfig Command with default input and output options
func DefaultGetClientConfigCommand(version core.RuntimeVersion, opts ...CfgClientConfigArgsOption) *core.Command {
	args := &CfgClientConfigArgs{}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &GetClientConfigInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
	}

	outputOpts := &GetClientConfigOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
		ClientConfigOpts:   args.ClientConfigOpts,
	}

	cmd, err := NewGetClientConfigCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}

// DefaultStoreClientConfigCommand creates a StoreClientConfig Command with default input and output options
func DefaultStoreClientConfigCommand(version core.RuntimeVersion, opts ...CfgClientConfigArgsOption) *core.Command {
	args := &CfgClientConfigArgs{}

	for _, opt := range opts {
		opt(args)
	}

	inputOpts := &StoreClientConfigInputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ClientConfigOpts: args.ClientConfigOpts,
	}

	outputOpts := &StoreClientConfigOutputOptions{
		RuntimeAPIVersion: &core.RuntimeAPIVersion{
			RuntimeVersion: version,
		},
		ValidationStrategy: args.ValidationStrategy,
		Error:              args.Error,
	}

	cmd, err := NewStoreClientConfigCommand(inputOpts, outputOpts)
	gomega.Expect(err).To(gomega.BeNil())
	return cmd
}
