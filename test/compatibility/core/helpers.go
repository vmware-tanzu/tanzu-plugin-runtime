// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

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

// ParseStr converts interface{} to string type
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
