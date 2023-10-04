// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

// runtime test plugins binaries directory path
const (
	pluginRoot = "../../../testplugins/bin"

	pluginV0116  = "runtime-test-plugin-v0_11_6"
	pluginV0254  = "runtime-test-plugin-v0_25_4"
	pluginV0280  = "runtime-test-plugin-v0_28_0"
	pluginV090   = "runtime-test-plugin-v0_90"
	pluginV102   = "runtime-test-plugin-v1_0_2"
	pluginLatest = "runtime-test-plugin-latest"
)

const (
	testPluginFilePathArgument = " --file "
)

// ConstructTestPluginCmd constructs the specific runtime test plugin command as per runtime version and apis
func ConstructTestPluginCmd(version RuntimeVersion, apis []*API) (string, error) {
	// Create root command for the specified runtime version
	pluginCommand := makeRuntimeTestPluginCommand(version)

	// Create a temp file with apis to execute
	fileName, err := writeAPIsToTempFile(apis)
	if err != nil {
		return "", err
	}

	pluginCommand += testPluginFilePathArgument + fileName

	return pluginCommand, nil
}

// makeRuntimeTestPluginCommand construct the root runtime test plugin command as per runtime version specified
func makeRuntimeTestPluginCommand(version RuntimeVersion) string {
	switch version {
	case Version0116:
		return fmt.Sprintf("%v/%v test", pluginRoot, pluginV0116)
	case Version0254:
		return fmt.Sprintf("%v/%v test", pluginRoot, pluginV0254)
	case Version0280:
		return fmt.Sprintf("%v/%v test", pluginRoot, pluginV0280)
	case Version090:
		return fmt.Sprintf("%v/%v test", pluginRoot, pluginV090)
	case Version102:
		return fmt.Sprintf("%v/%v test", pluginRoot, pluginV102)
	case VersionLatest:
		return fmt.Sprintf("%v/%v test", pluginRoot, pluginLatest)
	default:
		return fmt.Sprintf("%v/%v test", pluginRoot, pluginLatest)
	}
}

// writeAPIsToTempFile create a temp file with all the api data that is sent to runtime-test-plugins-vX.XX.XX binaries
func writeAPIsToTempFile(apis []*API) (string, error) {
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

// SetupTempCfgFiles mock runtime config files
func SetupTempCfgFiles() (files []*os.File, cleanup func()) {
	// Setup config data
	cfgFile, err := os.CreateTemp("", "tanzu_config")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(cfgFile.Name(), []byte{}, 0644)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("TANZU_CONFIG", cfgFile.Name())
	if err != nil {
		fmt.Println(err)
	}
	cfgNextGenFile, err := os.CreateTemp("", "tanzu_config_ng")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(cfgNextGenFile.Name(), []byte{}, 0644)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("TANZU_CONFIG_NEXT_GEN", cfgNextGenFile.Name())
	if err != nil {
		fmt.Println(err)
	}
	cfgMetadataFile, err := os.CreateTemp("", "tanzu_config_metadata")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(cfgMetadataFile.Name(), []byte{}, 0644)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("TANZU_CONFIG_METADATA", cfgMetadataFile.Name())
	if err != nil {
		fmt.Println(err)
	}

	cleanup = func() {
		err = os.Remove(cfgFile.Name())
		if err != nil {
			fmt.Println(err)
		}

		err = os.Remove(cfgNextGenFile.Name())
		if err != nil {
			fmt.Println(err)
		}

		err = os.Remove(cfgMetadataFile.Name())
		if err != nil {
			fmt.Println(err)
		}
	}

	return []*os.File{cfgFile, cfgNextGenFile, cfgMetadataFile}, cleanup
}
