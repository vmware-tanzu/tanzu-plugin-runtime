// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"
)

// EULAStatus is the user's EULA acceptance status
type EULAStatus string

const (
	// User is shown the EULA, but has not accepted it.
	EULAStatusShown EULAStatus = "shown"
	// User has accepted EULA.
	EULAStatusAccepted EULAStatus = "accepted"
	// Acceptance state is not set
	EULAStatusUnset EULAStatus = ""
)

// GetCEIPOptIn retrieves ClientOptions ceipOptIn
func GetCEIPOptIn() (string, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return "", err
	}
	return getCEIPOptIn(node)
}

// SetCEIPOptIn adds or updates ceipOptIn value
func SetCEIPOptIn(val string) (err error) {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Add or Update ceipOptIn in the yaml node
	persist := setCLIOptionsString(node, KeyCEIPOptIn, val)

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}
	return err
}

func getCEIPOptIn(node *yaml.Node) (string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}
	if cfg != nil && cfg.CoreCliOptions != nil {
		return cfg.CoreCliOptions.CEIPOptIn, nil
	}
	return "", errors.New("ceipOptIn not found")
}

// GetEULAStatus retrieves EULA status
func GetEULAStatus() (EULAStatus, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return "", err
	}
	return getEULAStatus(node)
}

// SetEULAStatus adds or updates the EULA status
func SetEULAStatus(val EULAStatus) (err error) {
	if val != EULAStatusShown && val != EULAStatusUnset && val != EULAStatusAccepted {
		return errors.New("invalid eula status")
	}

	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Add or update EULA acceptance status in the yaml node
	persist := setCLIOptionsString(node, KeyEULAStatus, string(val))

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}
	return err
}

func getEULAStatus(node *yaml.Node) (EULAStatus, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}
	if cfg != nil && cfg.CoreCliOptions != nil {
		if cfg.CoreCliOptions.EULAStatus == "" {
			return EULAStatusUnset, nil
		}
		return EULAStatus(cfg.CoreCliOptions.EULAStatus), nil
	}
	return "", errors.New("eulaStatus not found")
}

// SetEULAAcceptedVersions updates the list of EULA versions accepted
func SetEULAAcceptedVersions(acceptedVersions []string) (err error) {
	for _, v := range acceptedVersions {
		if !semver.IsValid(v) {
			return errors.Errorf("invalid eula version: %v", v)
		}
	}

	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	var valueToSet string
	if len(acceptedVersions) > 0 {
		semver.Sort(acceptedVersions)
		valueToSet = strings.Join(acceptedVersions, ",")
	}

	// Add or update EULA accepted versions list in the yaml node
	persist := setCLIOptionsString(node, KeyEULAVersions, valueToSet)

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}
	return err
}

func getEULAAcceptedVersions(node *yaml.Node) ([]string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	if cfg != nil && cfg.CoreCliOptions != nil {
		if cfg.CoreCliOptions.EULAAcceptedVersions == "" {
			return []string{}, nil
		}
		return strings.Split(cfg.CoreCliOptions.EULAAcceptedVersions, ","), nil
	}
	return nil, errors.New("unable to find accepted eula versions")
}

// GetEULAAcceptedVersions returns the list of EULA versions accepted
func GetEULAAcceptedVersions() ([]string, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getEULAAcceptedVersions(node)
}

// GetCLIId retrieves cliId
func GetCLIId() (string, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return "", err
	}
	return getCLIId(node)
}

// SetCLIId adds or updates cliId value
func SetCLIId(val string) (err error) {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Add or Update cliId in the yaml node
	persist := setCLIOptionsString(node, KeyCLIId, val)

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}
	return err
}

func getCLIId(node *yaml.Node) (string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}
	if cfg != nil && cfg.CoreCliOptions != nil {
		return cfg.CoreCliOptions.CliID, nil
	}
	return "", errors.New("cliId not found")
}

// getNGCLIOptionsChildNode parses the yaml node and returns the matched node based on configOptions
func getNGCLIOptionsChildNode(key string, node *yaml.Node) *yaml.Node {
	configOptions := func(c *nodeutils.CfgNode) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyCLI, Type: yaml.MappingNode},
			{Name: key, Type: yaml.ScalarNode, Value: ""},
		}
	}
	keyNode := nodeutils.FindNode(node.Content[0], configOptions)
	return keyNode
}

func setCLIOptionsString(node *yaml.Node, key, val string) (persist bool) {
	cliOptionNode := getNGCLIOptionsChildNode(key, node)
	if cliOptionNode != nil && cliOptionNode.Value != val {
		cliOptionNode.Value = val
		persist = true
	}
	return persist
}
