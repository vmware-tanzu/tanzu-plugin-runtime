// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package hub provides functions to create Tanzu Hub client for specific context
package hub

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config"
)

const (
	EnvTanzuHubGraphQLEndpoint = "TANZU_HUB_GRAPHQL_ENDPOINT"
)

// HubClient client to talk to Tanzu Hub through GraphQL apis
// It includes authenticated GraphQL client from github.com/Khan/genqlient
// that can be used to do GraphQL queries
type HubClient struct {
	// ContextName is Tanzu CLI context name
	ContextName string
	// GraphQLClient can be used to do graphql queries
	GraphQLClient graphql.Client
}

// CreateHubClient returns an authenticated Tanzu Hub client for the specified
// tanzu context. This client includes an authenticated GraphQLClient from github.com/Khan/genqlient
// that can be used to do GraphQL queries.
// Internally it configures the client with CSP access token for each request
//
// EXPERIMENTAL: Both the function's signature and implementation are subjected to change/removal
// if an alternative means to provide equivalent functionality can be introduced.
func CreateHubClient(contextName string) (*HubClient, error) {
	accessToken, err := config.GetTanzuContextAccessToken(contextName)
	if err != nil {
		return nil, err
	}

	tanzuHubEndpoint, err := getTanzuHubEndpointFromContext(contextName)
	if err != nil {
		return nil, err
	}

	httpClient := http.Client{
		Transport: &authTransport{
			accessToken: accessToken,
			wrapped:     http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient(fmt.Sprintf("%s/graphql", tanzuHubEndpoint), &httpClient)
	return &HubClient{
		ContextName:   contextName,
		GraphQLClient: graphqlClient,
	}, nil
}

func getTanzuHubEndpointFromContext(contextName string) (string, error) {
	// If `TANZU_HUB_GRAPHQL_ENDPOINT` environment variable is configured use that
	if endpoint := os.Getenv(EnvTanzuHubGraphQLEndpoint); endpoint != "" {
		return endpoint, nil
	}

	// Try to fetch the endpoint from the context metadata
	tzCtx, err := config.GetContext(contextName)
	if err != nil {
		return "", err
	}

	tanzuHubEndpoint := tzCtx.AdditionalMetadata[config.TanzuHubEndpointKey]
	if tanzuHubEndpoint == "" {
		return "", errors.Errorf("%q has not been configured for the %q context", config.TanzuHubEndpointKey, tzCtx.Name)
	}
	return tanzuHubEndpoint.(string), nil
}

// Configure the auth Transport to include authorization token when invoking GraphQL requests
type authTransport struct {
	accessToken string
	wrapped     http.RoundTripper
}

// Sets authentication bearer token to each http request
func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("authorization", "Bearer "+t.accessToken)
	return t.wrapped.RoundTrip(req)
}
