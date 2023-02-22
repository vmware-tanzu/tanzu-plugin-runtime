// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

const (
	pluginRoot = "../compatibility-test-plugins"

	pluginV0116 = "runtime-test-plugin-v0_11_6"
	pluginV0254 = "runtime-test-plugin-v0_25_4"
	pluginV0280 = "runtime-test-plugin-v0_28_0"
	pluginV100  = "runtime-test-plugin-v1_0_0"
)

// ConstructTestPluginCmd constructs the specific runtime test plugin command as per runtime version and apis
func ConstructTestPluginCmd(version core.RuntimeVersion, apis []*core.API) (string, error) {
	// Create root command for the specified runtime version
	pluginCommand := makeRootCommand(version)

	// Create a temp file with apis to execute
	fileName, err := writeAPIsToTempFile(apis)
	if err != nil {
		return "", err
	}

	pluginCommand += " --file " + fileName

	return pluginCommand, nil
}

// makeRootCommand construct the root runtime test plugin command as per runtime version specified
func makeRootCommand(version core.RuntimeVersion) string {
	switch version {
	case core.Version0116:
		return pluginRoot + "/" + pluginV0116 + "/" + pluginV0116 + " test"
	case core.Version0254:
		return pluginRoot + "/" + pluginV0254 + "/" + pluginV0254 + " test"
	case core.Version0280:
		return pluginRoot + "/" + pluginV0280 + "/" + pluginV0280 + " test"
	case core.Version100:
		return pluginRoot + "/" + pluginV100 + "/" + pluginV100 + " test"
	default:
		return pluginRoot + "/" + pluginV100 + "/" + pluginV100 + " test"
	}
}

// writeAPIsToTempFile create a temp file with all the api data
func writeAPIsToTempFile(apis []*core.API) (string, error) {
	b, err := yaml.Marshal(apis)
	if err != nil {
		return "", err
	}

	f, err := os.CreateTemp("", "runtime_compatibility_testing")
	if err != nil {
		return "", err
	}

	err = os.WriteFile(f.Name(), b, 0644)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

// Exec the command, exit on error
func Exec(command string) (stdOut, stdErr *bytes.Buffer, err error) {
	cmdInput := strings.Split(command, " ")
	cmdName := cmdInput[0]
	cmdArgs := cmdInput[1:]

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return &stdout, &stderr, fmt.Errorf(fmt.Sprintf("error while running %s", command), err)
	}
	return &stdout, &stderr, nil
}

// StrToMap convert str represented struct data to map
func StrToMap(s string) map[string]interface{} {
	var mapper map[string]interface{}

	err := yaml.Unmarshal([]byte(s), &mapper)
	if err != nil {
		fmt.Println(err)
	}
	return mapper
}
