// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"

	"gopkg.in/yaml.v3"
)

// parseContext unmarshalls string to Context struct
func parseContext(context string) (*configtypes.Context, error) {
	var ctx configtypes.Context
	err := yaml.Unmarshal([]byte(context), &ctx)
	if err != nil {
		return nil, err
	}
	return &ctx, nil
}

// parseServer unmarshalls string to Server struct
func parseServer(server string) (*configtypes.Server, error) { //nolint:staticcheck // Deprecated
	var s configtypes.Server //nolint:staticcheck // Deprecated
	err := yaml.Unmarshal([]byte(server), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
