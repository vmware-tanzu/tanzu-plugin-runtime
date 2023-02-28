// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	pluginRoot = "../compatibility-test-plugins/bin"

	pluginV0116 = "runtime-test-plugin-v0_11_6"
	pluginV0254 = "runtime-test-plugin-v0_25_4"
	pluginV0280 = "runtime-test-plugin-v0_28_0"
	pluginV100  = "runtime-test-plugin-v1_0_0"
)

const (
	testPluginFilePathArgument = " --file "
)

// ConstructTestPluginCmd constructs the specific runtime test plugin command as per runtime version and apis
func ConstructTestPluginCmd(version RuntimeVersion, apis []*API) (string, error) {
	// Create root command for the specified runtime version
	pluginCommand := makeRootCommand(version)

	// Create a temp file with apis to execute
	fileName, err := writeAPIsToTempFile(apis)
	if err != nil {
		return "", err
	}

	pluginCommand += testPluginFilePathArgument + fileName

	return pluginCommand, nil
}

// makeRootCommand construct the root runtime test plugin command as per runtime version specified
func makeRootCommand(version RuntimeVersion) string {
	switch version {
	case Version0116:
		return pluginRoot + "/" + pluginV0116 + " test"
	case Version0254:
		return pluginRoot + "/" + pluginV0254 + " test"
	case Version0280:
		return pluginRoot + "/" + pluginV0280 + " test"
	case Version100:
		return pluginRoot + "/" + pluginV100 + " test"
	default:
		return pluginRoot + "/" + pluginV100 + " test"
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

// StrToMap convert str represented struct data to map
func StrToMap(s string) map[string]interface{} {
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

// ParseRuntimeAPIsFromFile reads the filepath and unmarshalls the file content into array of APIs
func ParseRuntimeAPIsFromFile(filePath string) ([]API, error) {
	var apis []API
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &apis)
	if err != nil {
		return nil, err
	}

	return apis, nil
}

// ParseStr converts interface{} to string type
func ParseStr(val interface{}) (string, error) {
	var value string

	data, err := yaml.Marshal(val)
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(data, &value)
	if err != nil {
		return "", err
	}

	return value, nil
}

// ValidateMaps recursive equality check on map structs
func ValidateMaps(actual, expected map[string]interface{}) bool {
	for k, v := range expected {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			ValidateMaps(actual[k].(map[string]interface{}), v.(map[string]interface{}))
		} else if !reflect.DeepEqual(actual[k], v) {
			return false
		}
	}
	return true
}

// ParseStdout convert the string represented std out log into map structure
func ParseStdout(s string) (map[RuntimeAPIName][]APILog, error) {
	var logs map[RuntimeAPIName][]APILog
	err := yaml.Unmarshal([]byte(s), &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
