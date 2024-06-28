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

	accessToken      string
	tanzuHubEndpoint string
	httpClient       *http.Client
}

type ClientOptions func(o *HubClient)

// WithAccessToken creates the HubClient using the specified Access Token
func WithAccessToken(token string) ClientOptions {
	return func(c *HubClient) {
		c.accessToken = token
	}
}

// WithEndpoint creates the HubClient using the specified Endpoint
func WithEndpoint(endpoint string) ClientOptions {
	return func(c *HubClient) {
		c.tanzuHubEndpoint = endpoint
	}
}

// WithHTTPClient creates the HubClient using the specified HttpClient
func WithHTTPClient(httpClient *http.Client) ClientOptions {
	return func(c *HubClient) {
		c.httpClient = httpClient
	}
}

// CreateHubClient returns an authenticated Tanzu Hub client for the specified
// tanzu context. This client includes an authenticated GraphQLClient from github.com/Khan/genqlient
// that can be used to do GraphQL queries.
// Internally it configures the client with CSP access token for each request
//
// Note that the authenticated client is assured to have at least 30 min access to the GraphQL endpoint.
// If you want a long running client beyond this period, recommendation is to reinitialize your client.
//
// EXPERIMENTAL: Both the function's signature and implementation are subjected to change/removal
// if an alternative means to provide equivalent functionality can be introduced.
func CreateHubClient(contextName string, opts ...ClientOptions) (*HubClient, error) {
	hc := &HubClient{
		ContextName: contextName,
	}

	// configure all options for the HubClient
	for _, o := range opts {
		o(hc)
	}

	httpClient, err := hc.getHTTPClient(contextName)
	if err != nil {
		return nil, err
	}
	hc.GraphQLClient = graphql.NewClient(fmt.Sprintf("%s/graphql", hc.tanzuHubEndpoint), httpClient)
	return hc, nil
}

func (hc *HubClient) getHTTPClient(contextName string) (*http.Client, error) {
	var err error
	if hc.httpClient != nil {
		return hc.httpClient, nil
	}

	if hc.accessToken == "" {
		hc.accessToken, err = config.GetTanzuContextAccessToken(contextName)
		if err != nil {
			return nil, err
		}
	}

	if hc.tanzuHubEndpoint == "" {
		hc.tanzuHubEndpoint, err = getTanzuHubEndpointFromContext(contextName)
		if err != nil {
			return nil, err
		}
	}

	return &http.Client{
		Transport: &authTransport{
			accessToken: hc.accessToken,
			wrapped:     http.DefaultTransport,
		},
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
