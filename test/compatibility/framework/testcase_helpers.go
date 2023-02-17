// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

const (
	pluginRoot = "../compatibility-test-plugins"

	pluginV011 = "runtime-test-plugin-0-11"
	pluginV025 = "runtime-test-plugin-0-25"
	pluginV028 = "runtime-test-plugin-0-28"
	pluginV100 = "runtime-test-plugin-1-00"
)

// constructTestPluginCmd constructs the specific runtime test plugin command as per runtime version and apis
func constructTestPluginCmd(version RuntimeVersion, apis []*API) (string, error) {

	// Create root command for the specified runtime version
	pluginCommand := makeRootCommand(version)

	// Create a temp file with apis to execute
	fileName, err := writeAPIsToTempFile(apis)
	if err != nil {
		return "", err
	}

	pluginCommand += " --file " + fileName

	fmt.Println("Generated  cmd", pluginCommand)
	return pluginCommand, nil
}

// makeRootCommand construct the root runtime test plugin command as per runtime version specified
func makeRootCommand(version RuntimeVersion) string {
	switch version {
	case Version0116:
		return pluginRoot + "/" + pluginV011 + "/" + pluginV011 + " test"
	case Version0254:
		return pluginRoot + "/" + pluginV025 + "/" + pluginV025 + " test"
	case Version0280:
		return pluginRoot + "/" + pluginV028 + "/" + pluginV028 + " test"
	case Version100:
		return pluginRoot + "/" + pluginV100 + "/" + pluginV100 + " test"
	default:
		return pluginRoot + "/" + pluginV100 + "/" + pluginV100 + " test"
	}
}

// writeAPIsToTempFile create a temp file with all the api data
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

// strToMap convert str represented struct data to map
func strToMap(s string) map[string]interface{} {
	var mapper map[string]interface{}

	err := yaml.Unmarshal([]byte(s), &mapper)
	if err != nil {
		fmt.Println(err)
	}
	return mapper
}

// SetupTempCfgFiles mock runtime config files
func SetupTempCfgFiles() (files []*os.File, cleanup func()) {
	// Setup config data
	cfgFile, err := os.CreateTemp("", "tanzu_config")

	err = os.WriteFile(cfgFile.Name(), []byte{}, 0644)

	err = os.Setenv("TANZU_CONFIG", cfgFile.Name())

	cfgNextGenFile, err := os.CreateTemp("", "tanzu_config_ng")

	err = os.WriteFile(cfgNextGenFile.Name(), []byte{}, 0644)

	err = os.Setenv("TANZU_CONFIG_NEXT_GEN", cfgNextGenFile.Name())

	cfgMetadataFile, err := os.CreateTemp("", "tanzu_config_metadata")

	err = os.WriteFile(cfgMetadataFile.Name(), []byte{}, 0644)

	err = os.Setenv("TANZU_CONFIG_METADATA", cfgMetadataFile.Name())

	cleanup = func() {
		err = os.Remove(cfgFile.Name())
		Expect(err).To(BeNil())

		err = os.Remove(cfgNextGenFile.Name())
		Expect(err).To(BeNil())

		err = os.Remove(cfgMetadataFile.Name())
		Expect(err).To(BeNil())
	}

	return []*os.File{cfgFile, cfgNextGenFile, cfgMetadataFile}, cleanup
}
