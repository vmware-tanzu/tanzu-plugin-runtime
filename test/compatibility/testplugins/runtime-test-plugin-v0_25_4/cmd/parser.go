// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	configapi "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"

	"gopkg.in/yaml.v3"
)

// parseContext unmarshalls string to Context struct
func parseContext(context string) (*configapi.Context, error) {
	var ctx configapi.Context
	err := yaml.Unmarshal([]byte(context), &ctx)
	if err != nil {
		return nil, err
	}
	return &ctx, nil
}

// parseServer unmarshalls string to Server struct
func parseServer(server string) (*configapi.Server, error) {
	var s configapi.Server
	err := yaml.Unmarshal([]byte(server), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
