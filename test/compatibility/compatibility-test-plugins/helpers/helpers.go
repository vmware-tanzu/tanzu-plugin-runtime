// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"fmt"
	"os"

	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"gopkg.in/yaml.v3"
)

// GetTestData reads the filepath and unmarshalls the file content into array of APIs
func GetTestData(filePath string) ([]compatibilitytestingframework.API, error) {
	var apis []compatibilitytestingframework.API
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, &apis)
	if err != nil {
		return nil, err
	}

	return apis, nil
}

//ParseStr converts interface{} to string type
func ParseStr(val interface{}) (string, error) {
	var value string

	bytes, err := yaml.Marshal(val)
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(bytes, &value)
	if err != nil {
		return "", err
	}

	return value, nil
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
