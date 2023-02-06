package cmd

import (
	. "github.com/onsi/gomega"
	configtypes "github.com/vmware-tanzu/tanzu-framework/apis/config/v1alpha1"
	configlib "github.com/vmware-tanzu/tanzu-framework/pkg/v1/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
	"gopkg.in/yaml.v3"
)

// triggerServerAPIs trigger context related runtime apis and construct logs
func triggerServerAPIs(api *compatibilitytestingframework.API, logs map[string][]compatibilitytestingframework.APILog) {
	if api.Name == compatibilitytestingframework.AddServerAPIName {
		log := triggerAddServerAPI(api)
		logs[compatibilitytestingframework.AddServerAPIName] = append(logs[compatibilitytestingframework.AddServerAPIName], log)
	}
	if api.Name == compatibilitytestingframework.GetServerAPIName {
		log := triggerGetServerAPI(api)
		logs[compatibilitytestingframework.GetServerAPIName] = append(logs[compatibilitytestingframework.GetServerAPIName], log)
	}
}

// triggerGetServerAPI trigger get context runtime api
func triggerGetServerAPI(api *compatibilitytestingframework.API) compatibilitytestingframework.APILog {
	// Parse arguments needed to trigger the runtime api
	serverName, err := helpers.ParseStr(api.Arguments["serverName"])
	Expect(err).To(BeNil())
	//Call runtime GetServer API
	server, err := configlib.GetServer(serverName)

	// Construct logging
	log := compatibilitytestingframework.APILog{}
	if err != nil {
		log.APIError = err.Error()
	}
	log.APIResponse = &compatibilitytestingframework.APIResponse{
		ResponseBody: server,
		ResponseType: compatibilitytestingframework.MapResponse,
	}
	return log
}

// triggerAddServerAPI trigger set context runtime api
func triggerAddServerAPI(api *compatibilitytestingframework.API) compatibilitytestingframework.APILog {
	// Parse arguments needed to trigger the runtime api
	server, err := parseServer(api.Arguments["server"].(string))
	Expect(err).To(BeNil())
	isCurrent := api.Arguments["isCurrent"].(bool)

	// Call the runtime SetServer API
	err = configlib.AddServer(server, isCurrent)

	// Construct logging
	log := compatibilitytestingframework.APILog{}
	if err != nil {
		log.APIError = err.Error()
	}
	log.APIResponse = &compatibilitytestingframework.APIResponse{
		ResponseBody: "",
		ResponseType: compatibilitytestingframework.StringResponse,
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
