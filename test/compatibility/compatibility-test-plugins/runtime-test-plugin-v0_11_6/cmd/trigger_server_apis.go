// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"

	configtypes "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	compatibilitytestingcore "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerServerAPIs trigger context related runtime apis and construct logs
func triggerServerAPIs(api *compatibilitytestingcore.API, logs map[compatibilitytestingcore.RuntimeAPIName][]compatibilitytestingcore.APILog) {
	if api.Name == compatibilitytestingcore.AddServerAPIName {
		log := triggerAddServerAPI(api)
		logs[compatibilitytestingcore.AddServerAPIName] = append(logs[compatibilitytestingcore.AddServerAPIName], log)
	}
	if api.Name == compatibilitytestingcore.GetServerAPIName {
		log := triggerGetServerAPI(api)
		logs[compatibilitytestingcore.GetServerAPIName] = append(logs[compatibilitytestingcore.GetServerAPIName], log)
	}
}

// triggerGetServerAPI trigger get context runtime api
func triggerGetServerAPI(api *compatibilitytestingcore.API) compatibilitytestingcore.APILog {
	// Parse arguments needed to trigger the runtime api
	serverName, err := compatibilitytestingcore.ParseStr(api.Arguments["serverName"])
	Expect(err).To(BeNil())
	//Call runtime GetServer API
	server, err := configlib.GetServer(serverName)

	// Construct logging
	log := compatibilitytestingcore.APILog{}
	if err != nil {
		log.APIError = err.Error()
	}
	log.APIResponse = &compatibilitytestingcore.APIResponse{
		ResponseBody: server,
		ResponseType: compatibilitytestingcore.MapResponse,
	}
	return log
}

// triggerAddServerAPI trigger set context runtime api
func triggerAddServerAPI(api *compatibilitytestingcore.API) compatibilitytestingcore.APILog {
	// Parse arguments needed to trigger the runtime api
	server, err := parseServer(api.Arguments["server"].(string))
	Expect(err).To(BeNil())
	isCurrent := api.Arguments["isCurrent"].(bool)

	// Call the runtime SetServer API
	err = configlib.AddServer(server, isCurrent)

	// Construct logging
	log := compatibilitytestingcore.APILog{}
	if err != nil {
		log.APIError = err.Error()
	}
	log.APIResponse = &compatibilitytestingcore.APIResponse{
		ResponseBody: "",
		ResponseType: compatibilitytestingcore.StringResponse,
	}
	return log
}

// parseServer unmarshalls string to Context struct
func parseServer(server string) (*configtypes.Server, error) {
	var s configtypes.Server
	err := yaml.Unmarshal([]byte(server), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
