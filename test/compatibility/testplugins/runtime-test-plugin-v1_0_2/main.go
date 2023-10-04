// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package main contains test cli plugin to trigger various runtime APIs
package main

import (
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/testplugins/runtime-test-plugin-v1_0_2/cmd"
)

func main() {
	cmd.Execute()
}
