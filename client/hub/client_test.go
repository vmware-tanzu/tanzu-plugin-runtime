// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package hub provides functions to create Tanzu Hub client for specific context
package hub_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/client/hub"
	hubtesting "github.com/vmware-tanzu/tanzu-plugin-runtime/client/hub/testing"
)

// The code in file `client_generated_test.go` contains the generated code for the
// following GraphQL query that we will be using for our tests.
//
//   query QueryAllProjects {
//     applicationEngineQuery {
//       queryProjects(first: 1000) {
//         projects {
//           name
//         }
//       }
//     }
//   }

// A set of variables that are sent with the query.
// This allows us to test that the mock server returns these
// variables in the response.
var defaultVariables = map[string]interface{}{
	"name":  "john",
	"lines": 100,
}

// getProjects is a wrapper of an `QueryAllProjects“ API call to fetch project names
// Note: This is only for the testing the MockServer with HubClient
func getProjects(hc hub.Client) ([]string, error) {
	req := &hub.Request{
		OpName:    "QueryAllProjects",
		Query:     QueryAllProjects_Operation,
		Variables: defaultVariables,
	}
	var responseData QueryAllProjectsResponse
	err := hc.Request(context.Background(), req, &responseData)
	if err != nil {
		return nil, err
	}

	projects := []string{}
	for _, p := range responseData.ApplicationEngineQuery.QueryProjects.Projects {
		projects = append(projects, p.Name)
	}

	return projects, nil
}

func TestQueryWithTanzuHubClient(t *testing.T) {
	// Start Mock GraphQL Server
	mockServer := hubtesting.NewServer(t)
	if mockServer == nil {
		t.Fatalf("error while starting a mock graphql server")
	}
	defer mockServer.Close()

	// Create the Hub Client using the above mock server
	hc, err := hub.NewClient("fake-context", hub.WithEndpoint(mockServer.URL), hub.WithAccessToken("fake-token"))
	if err != nil {
		t.Fatalf("error while creating hub client. %s", err.Error())
	}

	var tests = []struct {
		name              string
		mockResponses     []hubtesting.Operation
		expectedOutput    []string
		expectedErrString string
	}{
		{
			name: "no projects found",
			mockResponses: []hubtesting.Operation{
				{
					Identifier: "QueryAllProjects",
					Variables:  defaultVariables,
					Response: hub.Response{
						Data: QueryAllProjectsResponse{
							ApplicationEngineQuery: QueryAllProjectsApplicationEngineQuery{
								QueryProjects: QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection{
									Projects: []QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject{},
								},
							},
						},
					},
				},
			},
			expectedOutput: []string{},
		},
		{
			name: "when projects found and query returns response",
			mockResponses: []hubtesting.Operation{
				{
					Identifier: "QueryAllProjects",
					Variables:  defaultVariables,
					Response: hub.Response{
						Data: QueryAllProjectsResponse{
							ApplicationEngineQuery: QueryAllProjectsApplicationEngineQuery{
								QueryProjects: QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection{
									Projects: []QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject{
										{
											Name: "project1",
										},
										{
											Name: "project2",
										},
										{
											Name: "project3",
										},
									},
								},
							},
						},
					},
				},
			},
			expectedOutput: []string{"project1", "project2", "project3"},
		},
		{
			name: "when projects found and query returns response - use responder implementation",
			mockResponses: []hubtesting.Operation{
				{
					Identifier: "QueryAllProjects",
					// Change the value of the default variables
					// to make sure the mock server returns the variables that were
					// received in the query and NOT the ones that are registered here.
					// Note that the variable keys must match the ones in the query, but
					// the values can be different.
					Variables: map[string]interface{}{
						"name":  "notjohn",
						"lines": 0,
					},
					Responder: func(_ context.Context, receivedReq hubtesting.Request) hub.Response {
						return hub.Response{
							Data: QueryAllProjectsResponse{
								ApplicationEngineQuery: QueryAllProjectsApplicationEngineQuery{
									QueryProjects: QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection{
										Projects: []QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject{
											{
												Name: "project1",
											},
											{
												Name: "project2",
											},
											{
												// Check that the mock server returns the received query to the Responder
												// by putting the received query as a project name.
												Name: fmt.Sprintf("%v", receivedReq),
											},
										},
									},
								},
							},
						}
					},
				},
			},
			expectedOutput: []string{"project1", "project2", `{
query QueryAllProjects {
	applicationEngineQuery {
		queryProjects(first: 1000) {
			projects {
				name
			}
		}
	}
}
 map[lines:100 name:john]}`},
		},
		{
			name: "when query returns error response",
			mockResponses: []hubtesting.Operation{
				{
					Identifier: "QueryAllProjects",
					Variables:  defaultVariables,
					Response: hub.Response{
						Errors: gqlerror.List{{Message: "fake-error-message"}},
					},
				},
			},
			expectedErrString: "fake-error-message",
		},
		{
			name: "when query returns error response - use responder implementation",
			mockResponses: []hubtesting.Operation{
				{
					Identifier: "QueryAllProjects",
					// Change the value of the default variables
					// to make sure the mock server returns the variables that were
					// received in the query and NOT the ones that are registered here.
					// Note that the variable keys must match the ones in the query, but
					// the values can be different.
					Variables: map[string]interface{}{
						"name":  "notjohn",
						"lines": 0,
					},
					Responder: func(_ context.Context, receivedReq hubtesting.Request) hub.Response {
						return hub.Response{
							Errors: gqlerror.List{{Message: fmt.Sprintf("operation failed with error and received %v", receivedReq)}},
						}
					},
				},
			},
			expectedErrString: `operation failed with error and received {
query QueryAllProjects {
	applicationEngineQuery {
		queryProjects(first: 1000) {
			projects {
				name
			}
		}
	}
}
 map[lines:100 name:john]}
`,
		},
		{
			name:              "when the query is not registered with the server or incorrect query is used",
			mockResponses:     []hubtesting.Operation{},
			expectedOutput:    []string{},
			expectedErrString: "operation not found",
		},
		{
			name: "when the query is registered but has variables that use different keys",
			mockResponses: []hubtesting.Operation{
				{
					Identifier: "QueryAllProjects",
					Variables: map[string]interface{}{
						"differentkey": "anothervalue",
						"name":         "john",
						"lines":        100,
					},
					Response: hub.Response{
						Errors: gqlerror.List{{Message: "fake-error-message"}},
					},
				},
			},
			expectedOutput: []string{},
			// The error should not be the error returned by the registered query
			// because we are testing that the registered query does not match
			// because it uses different variable keys that the real query.
			expectedErrString: "operation not found",
		},
		{
			name: "when the query is registered but has variables that use different values",
			mockResponses: []hubtesting.Operation{
				{
					Identifier: "QueryAllProjects",
					// Change the value of the default variables but use the same keys
					// This should still match the registered query and return the error
					Variables: map[string]interface{}{
						"name":  "notjohn",
						"lines": 0,
					},
					Response: hub.Response{
						Errors: gqlerror.List{{Message: "fake-error-message"}},
					},
				},
			},
			expectedErrString: "fake-error-message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset all the registered queries before running a new test
			mockServer.Reset()

			// Register all the mock responses for this test
			mockServer.RegisterQuery(tt.mockResponses...)

			// Try to get projects using the hub client
			projects, err := getProjects(hc)

			// Compare the results
			if tt.expectedErrString == "" {
				assert.Equal(t, tt.expectedOutput, projects)
			} else {
				assert.Contains(t, err.Error(), tt.expectedErrString)
			}
		})
	}
}

// The query or mutation executed by QueryAllProjects.
const SubscriptionLogsOperation = `
subscription SubscribeAppLogs {
	appLogs {
		value
	}
}
`

type Log struct {
	Value string
}

type AppLogs struct {
	AppLog Log `json:"appLogs"`
}

// subscribeAppLogs is a wrapper of an `SubscribeAppLogs API call to fetch logs
// Note: This is only for the testing the MockServer with HubClient
func subscribeAppLogs(hc hub.Client) string {
	req := &hub.Request{
		OpName: "SubscribeAppLogs",
		Query:  SubscriptionLogsOperation,
	}

	logs := ""
	logProcessor := func(eventResponse hub.EventResponse) {
		resp := eventResponse.ResponseData
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}

		data := AppLogs{}
		responseTyped := &hub.Response{Data: &data}
		err = json.Unmarshal(b, responseTyped)
		if err != nil {
			return
		}
		logs += fmt.Sprintln(data.AppLog.Value)
	}

	// ctxSubscription, _ := context.WithCancel(context.Background())
	ctxSubscription := context.Background()

	_ = hc.Subscribe(ctxSubscription, req, logProcessor)
	// TODO: Figure how errors should be handled
	//  1. if server closes the connection this will always return error.
	//  2. if client closes the connection by closing ctxSubscription context

	return logs
}

func TestSubscriptionWithTanzuHubClient(t *testing.T) {
	// Start Mock GraphQL Server
	mockServer := hubtesting.NewServer(t)
	if mockServer == nil {
		t.Fatalf("error while starting a mock graphql server")
	}
	defer mockServer.Close()

	// Create the Hub Client using the above mock server
	hc, err := hub.NewClient("fake-context", hub.WithEndpoint(mockServer.URL), hub.WithAccessToken("fake-token"))
	if err != nil {
		t.Fatalf("error while creating hub client. %s", err.Error())
	}

	var tests = []struct {
		name          string
		mockResponses []hubtesting.Operation
		expectedLogs  string
	}{
		{
			name: "app logs",
			mockResponses: []hubtesting.Operation{
				{
					Identifier:     "SubscribeAppLogs",
					EventGenerator: mockAppLogGenerator,
				},
			},
			expectedLogs: `log 0
log 1
log 2
log 3
log 4

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset all the registered queries before running a new test
			mockServer.Reset()

			// Register all the mock responses for this test
			mockServer.RegisterSubscription(tt.mockResponses...)

			// Try to get projects using the hub client
			logs := subscribeAppLogs(hc)

			// Compare the results
			assert.Equal(t, tt.expectedLogs, logs)
		})
	}
}

func mockAppLogGenerator(_ context.Context, _ hubtesting.Request, eventData chan<- hubtesting.Response) {
	i := 0
	for i < 5 {
		time.Sleep(1 * time.Second)
		eventData <- hubtesting.Response{
			Data: AppLogs{AppLog: Log{Value: fmt.Sprintf("log %v", i)}},
		}
		i++
	}
	close(eventData)
}
