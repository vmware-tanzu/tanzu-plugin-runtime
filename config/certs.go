// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/nodeutils"

	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

// GetCerts retrieves all the certs
func GetCerts() ([]*configtypes.Cert, error) {
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCerts(node)
}

// GetCert retrieves the cert configuration by host
func GetCert(host string) (*configtypes.Cert, error) {
	if host == "" {
		return nil, errors.New("host is empty")
	}
	// Retrieve client config node
	node, err := getClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getCert(node, host)
}

// SetCert add or update cert configuration
func SetCert(c *configtypes.Cert) error {
	if c == nil {
		return nil
	}
	if c.Host == "" {
		return errors.New("host is empty")
	}
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	// Add or update the cert
	persist, err := setCert(node, c)
	if err != nil {
		return err
	}
	if persist {
		err = persistConfig(node)
		if err != nil {
			return err
		}
	}
	return err
}

// DeleteCert delete a cert configuration by host
func DeleteCert(host string) error {
	if host == "" {
		return errors.New("host is empty")
	}
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}
	_, err = getCert(node, host)
	if err != nil {
		return err
	}
	err = removeCert(node, host)
	if err != nil {
		return err
	}
	return persistConfig(node)
}

// CertExists checks if cert config by host already exists
func CertExists(host string) (bool, error) {
	if host == "" {
		return false, errors.New("host is empty")
	}
	exists, _ := GetCert(host)
	return exists != nil, nil
}

// Pre-reqs: node != nil
func getCerts(node *yaml.Node) ([]*configtypes.Cert, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	if cfg.Certs != nil {
		return cfg.Certs, nil
	}
	// return empty list if the config doesn't have Certs
	return make([]*configtypes.Cert, 0), nil
}

// Pre-reqs: node != nil and host != ""
func getCert(node *yaml.Node, host string) (*configtypes.Cert, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	for _, cert := range cfg.Certs {
		if cert.Host == host {
			return cert, nil
		}
	}
	return nil, fmt.Errorf("cert configuration for %v not found", host)
}

// Pre-reqs: node != nil and cert != nil
func setCert(node *yaml.Node, cert *configtypes.Cert) (persist bool, err error) {
	// Get Patch Strategies from config metadata
	patchStrategies, err := GetConfigMetadataPatchStrategy()
	if err != nil {
		patchStrategies = make(map[string]string)
	}

	// Convert cert to node
	newCertNode, err := convertObjectToNode(cert)
	if err != nil {
		return persist, err
	}

	// Find the certs node from the root node
	keys := []nodeutils.Key{
		{Name: KeyCerts, Type: yaml.SequenceNode},
	}
	certsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithForceCreate(), nodeutils.WithKeys(keys))
	if certsNode == nil {
		return persist, err
	}

	exists := false
	var result []*yaml.Node
	// Skip duplicate for cert
	for _, certNode := range certsNode.Content {
		if index := nodeutils.GetNodeIndex(certNode.Content, "host"); index != -1 &&
			certNode.Content[index].Value == cert.Host {
			exists = true
			// replace the nodes as per patch strategy
			_, err = nodeutils.DeleteNodes(newCertNode.Content[0], certNode, nodeutils.WithPatchStrategyKey(KeyCerts), nodeutils.WithPatchStrategies(patchStrategies))
			if err != nil {
				return false, err
			}
			persist, err = nodeutils.MergeNodes(newCertNode.Content[0], certNode)
			if err != nil {
				return false, err
			}
			result = append(result, certNode)
			continue
		}
		result = append(result, certNode)
	}
	if !exists {
		result = append(result, newCertNode.Content[0])
		persist = true
	}
	certsNode.Content = result
	return persist, err
}

func removeCert(node *yaml.Node, host string) error {
	// Find the certs node in the yaml node
	keys := []nodeutils.Key{
		{Name: KeyCerts},
	}
	certsNode := nodeutils.FindNode(node.Content[0], nodeutils.WithKeys(keys))
	if certsNode == nil {
		return nil
	}
	var certs []*yaml.Node
	for _, certNode := range certsNode.Content {
		if index := nodeutils.GetNodeIndex(certNode.Content, "host"); index != -1 && certNode.Content[index].Value == host {
			continue
		}
		certs = append(certs, certNode)
	}
	certsNode.Content = certs
	return nil
}
