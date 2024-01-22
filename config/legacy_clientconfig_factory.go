// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// StoreClientConfig stores the config in the local directory.
// Make sure to Acquire and Release tanzu lock when reading/writing to the
// tanzu client configuration
// Deprecated: StoreClientConfig is deprecated. Avoid using this method for Delete operations. Use New Config API methods.
func StoreClientConfig(cfg *configtypes.ClientConfig) error {
	// new plugins would be setting only contexts, so populate servers for backwards compatibility
	populateServers(cfg)
	// old plugins would be setting only servers, so populate contexts for forwards compatibility
	PopulateContexts(cfg)

	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	err = setServers(node, cfg.KnownServers)
	if err != nil {
		return err
	}
	if cfg.CurrentServer != "" {
		_, err = setCurrentServer(node, cfg.CurrentServer)
		if err != nil {
			return err
		}
	}
	err = setContexts(node, cfg.KnownContexts)
	if err != nil {
		return err
	}
	err = clientConfigSetCurrentContext(cfg, node)
	if err != nil {
		return err
	}
	err = clientConfigSetClientOptions(cfg, node)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

func clientConfigSetClientOptions(cfg *configtypes.ClientConfig, node *yaml.Node) error {
	if cfg.ClientOptions != nil {
		err := clientConfigSetFeatures(cfg, node)
		if err != nil {
			return err
		}
		err = clientConfigSetEnvs(cfg, node)
		if err != nil {
			return err
		}
		err = clientConfigSetCLI(cfg, node)
		if err != nil {
			return err
		}
	}
	return nil
}

// Deprecated: This method is deprecated
func clientConfigSetCLI(cfg *configtypes.ClientConfig, node *yaml.Node) (err error) {
	if cfg.ClientOptions.CLI != nil {
		err = clientConfigSetCLIRepositories(cfg, node)
		if err != nil {
			return err
		}
		if cfg.ClientOptions.CLI.UnstableVersionSelector != "" {
			setUnstableVersionSelector(node, string(cfg.ClientOptions.CLI.UnstableVersionSelector))
		}
		if cfg.ClientOptions.CLI.Edition != "" {
			setEdition(node, string(cfg.ClientOptions.CLI.Edition))
		}
		if cfg.ClientOptions.CLI.BOMRepo != "" {
			setBomRepo(node, cfg.ClientOptions.CLI.BOMRepo)
		}
		if cfg.ClientOptions.CLI.CompatibilityFilePath != "" {
			setCompatibilityFilePath(node, cfg.ClientOptions.CLI.CompatibilityFilePath)
		}
	}
	return nil
}

// Deprecated: This method is deprecated
func clientConfigSetCLIRepositories(cfg *configtypes.ClientConfig, node *yaml.Node) error {
	if cfg.ClientOptions.CLI.Repositories != nil && len(cfg.ClientOptions.CLI.Repositories) != 0 {
		err := setCLIRepositories(node, cfg.ClientOptions.CLI.Repositories)
		if err != nil {
			return err
		}
	}
	return nil
}

func clientConfigSetEnvs(cfg *configtypes.ClientConfig, node *yaml.Node) error {
	if cfg.ClientOptions.Env != nil {
		for key, value := range cfg.ClientOptions.Env {
			_, err := setEnv(node, key, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func clientConfigSetFeatures(cfg *configtypes.ClientConfig, node *yaml.Node) error {
	if cfg.ClientOptions.Features != nil {
		for plugin := range cfg.ClientOptions.Features {
			for key, value := range cfg.ClientOptions.Features[plugin] {
				_, err := setFeature(node, plugin, key, value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func clientConfigSetCurrentContext(cfg *configtypes.ClientConfig, node *yaml.Node) error {
	if cfg.CurrentContext != nil {
		for _, contextName := range cfg.CurrentContext {
			ctx, contextErr := cfg.GetContext(contextName)
			if contextErr != nil {
				return contextErr
			}
			_, err := setCurrentContext(node, ctx.Name, ctx.ContextType)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteClientConfig deletes the config yaml from the local directory.
func DeleteClientConfig() error {
	cfgPath, err := ClientConfigPath()
	if err != nil {
		return err
	}
	err = os.Remove(cfgPath)
	if err != nil {
		return errors.Wrap(err, "could not remove config")
	}
	return nil
}

// DeleteClientConfigNextGen deletes the config-ng yaml from the local directory.
func DeleteClientConfigNextGen() error {
	cfgPath, err := ClientConfigNextGenPath()
	if err != nil {
		return err
	}
	err = os.Remove(cfgPath)
	if err != nil {
		return errors.Wrap(err, "could not remove config-ng")
	}
	return nil
}
