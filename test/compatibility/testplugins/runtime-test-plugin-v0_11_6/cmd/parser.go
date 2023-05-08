// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	configtypes "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"

	"gopkg.in/yaml.v3"
)

// parseServer unmarshalls string to Server struct
func parseServer(server string) (*configtypes.Server, error) {
	var s configtypes.Server
	err := yaml.Unmarshal([]byte(server), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
