// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/onsi/gomega"
	"gopkg.in/yaml.v3"

	configtypes "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

// triggerServerAPIs trigger server related runtime apis and construct response logs
func triggerServerAPIs(api *core.API, logs map[core.RuntimeAPIName][]core.APILog) {
	if api.Name == core.AddServerAPIName {
		log := triggerAddServerAPI(api)
		logs[core.AddServerAPIName] = append(logs[core.AddServerAPIName], log)
	}
	if api.Name == core.GetServerAPIName {
		log := triggerGetServerAPI(api)
		logs[core.GetServerAPIName] = append(logs[core.GetServerAPIName], log)
	}
}

// triggerGetServerAPI trigger get server runtime api
func triggerGetServerAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	serverName, err := core.ParseStr(api.Arguments[core.ServerName])
	gomega.Expect(err).To(gomega.BeNil())

	// Call runtime GetServer API
	server, err := configlib.GetServer(serverName)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	}

	if server != nil {
		log.APIResponse = &core.APIResponse{
			ResponseBody: server,
			ResponseType: core.MapResponse,
		}
	}

	return log
}

// triggerAddServerAPI trigger add server runtime api
func triggerAddServerAPI(api *core.API) core.APILog {
	// Parse arguments needed to trigger the runtime api
	server, err := parseServer(api.Arguments[core.Server].(string))
	gomega.Expect(err).To(gomega.BeNil())
	setCurrent := api.Arguments[core.SetCurrent].(bool)

	// Call runtime SetServer API
	err = configlib.AddServer(server, setCurrent)

	// Construct logging
	log := core.APILog{}
	if err != nil {
		log.APIResponse = &core.APIResponse{
			ResponseType: core.ErrorResponse,
			ResponseBody: err.Error(),
		}
	} else {
		log.APIResponse = &core.APIResponse{
			ResponseBody: "",
			ResponseType: core.StringResponse,
		}
	}
	return log
}

// parseServer unmarshalls string to Server struct
func parseServer(server string) (*configtypes.Server, error) {
	var s configtypes.Server
	err := yaml.Unmarshal([]byte(server), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
