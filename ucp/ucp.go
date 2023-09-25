// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package ucp provides APIs specific to ucp
package ucp

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/ucp/internal/kubeconfig"
)

const (
	OrgID = "ucpOrgID"
)

// GetKubeconfigForContext returns the kubeconfig for any arbitrary UCP resource in the UCP object hierarchy
// referred by the UCP context
// Pre-reqs: project and space names should be valid
//
// Notes:
// If projectName and spaceName is empty string the kubeconfig generated would be pointing to UCP org
//
//	ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid
//
// If projectName is valid projectName and spaceName is empty string the kubeconfig generated would be pointing to UCP project
//
//	ex: kubeconfig's cluster.server URL : https://endpoint/org/orgid/project/<projectName>
//
// similarly if both project and space names are valid names the kubeconfig generated would be pointing to UCP space
//
//	ex: kubeconfig's cluster.server URL:  https://endpoint/org/orgid/project/<projectName>/space/<spaceName>
func GetKubeconfigForContext(contextName, projectName, spaceName string) ([]byte, error) {
	ctx, err := config.GetContext(contextName)
	if err != nil {
		return nil, err
	}
	if ctx.Target != configtypes.TargetUCP {
		return nil, errors.Errorf("context must be of type: %s", configtypes.TargetUCP)
	}

	kc, err := kubeconfig.ReadKubeConfig(ctx.ClusterOpts.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read the UCP context kubeconfig")
	}

	kc, err = kubeconfig.MinifyKubeConfig(kc, ctx.ClusterOpts.Context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to minify the kubeconfig")
	}
	updateKubeconfigServerURL(kc, ctx, projectName, spaceName)

	kubeconfigBytes, err := yaml.Marshal(kc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal the kubeconfig")
	}
	return kubeconfigBytes, nil
}

func prepareClusterServerURL(context *configtypes.Context, projectName, spaceName string) string {
	serverURL := context.ClusterOpts.Endpoint
	if projectName == "" {
		return serverURL
	}
	serverURL = serverURL + "/project/" + projectName

	if spaceName == "" {
		return serverURL
	}
	return serverURL + "/space/" + spaceName
}

func updateKubeconfigServerURL(kc *kubeconfig.Config, cliContext *configtypes.Context, projectName, spaceName string) {
	currentContextName := kc.CurrentContext
	context := kubeconfig.GetContext(kc, currentContextName)
	cluster := kubeconfig.GetCluster(kc, context.Context.Cluster)
	cluster.Cluster.Server = prepareClusterServerURL(cliContext, projectName, spaceName)
}
