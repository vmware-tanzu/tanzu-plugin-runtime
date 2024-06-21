// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package hub provides functions to create Tanzu Hub client for specific context
package hub_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

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

// getProjects is a wrapper of an `QueryAllProjects“ API call to fetch project names
// Note: This is only for the testing the MockServer with HubClient
func getProjects(hc hub.Client) ([]string, error) {
	req := &hub.Request{
		OpName: "QueryAllProjects",
		Query:  QueryAllProjects_Operation,
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

func TestTanzuHubClient(t *testing.T) {
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
					Response: QueryAllProjectsResponse{
						ApplicationEngineQuery: QueryAllProjectsApplicationEngineQuery{
							QueryProjects: QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection{
								Projects: []QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject{},
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
					Response: QueryAllProjectsResponse{
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
			expectedOutput: []string{"project1", "project2", "project3"},
		},
		{
			name:              "when the query is not registered with the server or incorrect query is used",
			mockResponses:     []hubtesting.Operation{},
			expectedOutput:    []string{},
			expectedErrString: "operation not found",
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
