// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:            "helloworld-test",
	Target:          types.TargetGlobal,
	Description:     "Hello world test plugin",
	Group:           plugin.AdminCmdGroup,
	Version:         "v0.0.1",
	BuildSHA:        buildinfo.SHA,
	PostInstallHook: postInstallHook,
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "")
	}

	p.AddCommands(
		PrintCmd,
	)

	if err := p.Execute(); err != nil {
		log.Fatal(err, "")
	}
}

func postInstallHook() error {
	fmt.Println("Hello world - post install method")
	return nil
}
