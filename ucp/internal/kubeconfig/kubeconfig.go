// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package kubeconfig provides the kubeconfig file access functionality
package kubeconfig

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// ReadKubeConfig reads the kubeconfig file and returns the Config
func ReadKubeConfig(path string) (*Config, error) {
	kubeconfig, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(kubeconfig, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// MinifyKubeConfig removes all information not used by the provided kubecontext name
func MinifyKubeConfig(kubeconfig *Config, kubeContextName string) (*Config, error) {
	context := GetContext(kubeconfig, kubeContextName)
	if context == nil {
		return nil, errors.Errorf("context %q missing in the kubeconfig", kubeContextName)
	}

	clusterName := context.Context.Cluster
	userName := context.Context.AuthInfo

	cluster := GetCluster(kubeconfig, clusterName)
	if cluster == nil {
		return nil, errors.Errorf("cluster %q missing in the kubeconfig", clusterName)
	}
	user := GetAuthInfo(kubeconfig, userName)
	if user == nil {
		return nil, errors.Errorf("user %q missing in the kubeconfig", userName)
	}

	return &Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       []*Cluster{{Name: clusterName, Cluster: cluster.Cluster}},
		AuthInfos:      []*AuthInfo{{Name: userName, AuthInfo: user.AuthInfo}},
		Contexts:       []*Context{{Name: kubeContextName, Context: context.Context}},
		CurrentContext: kubeContextName,
	}, nil
}

// GetContext returns the Context information from the provided config
func GetContext(kubeconfig *Config, kubeContextName string) *Context {
	for idx := range kubeconfig.Contexts {
		if kubeconfig.Contexts[idx].Name == kubeContextName {
			return kubeconfig.Contexts[idx]
		}
	}
	return nil
}

// GetCluster returns the Cluster information from the provided config
func GetCluster(kubeconfig *Config, clusterName string) *Cluster {
	for idx := range kubeconfig.Clusters {
		if kubeconfig.Clusters[idx].Name == clusterName {
			return kubeconfig.Clusters[idx]
		}
	}
	return nil
}

// GetAuthInfo returns the AuthInfo information from the provided config
func GetAuthInfo(kubeconfig *Config, userName string) *AuthInfo {
	for idx := range kubeconfig.AuthInfos {
		if kubeconfig.AuthInfos[idx].Name == userName {
			return kubeconfig.AuthInfos[idx]
		}
	}
	return nil
}
