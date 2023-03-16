// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"os"

	"gopkg.in/yaml.v3"
)

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

// ParseStdout convert the string represented std out log into map structure
func ParseStdout(s string) (map[RuntimeAPIName][]APILog, error) {
	var logs map[RuntimeAPIName][]APILog
	err := yaml.Unmarshal([]byte(s), &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
