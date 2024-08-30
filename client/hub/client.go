// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package hub provides functions to create Tanzu Hub client for specific context
package hub

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
)

const (
	EnvTanzuHubEndpoint = "TANZU_HUB_ENDPOINT"
)

// Client is an interface for the Tanzu Hub Client
type Client interface {
	// Request sends a GraphQL request to the Tanzu Hub endpoint
	//
	//	ctx context.Context: The context for the request. If provided, it will be used to cancel the request if the context is canceled.
	//	req *Request: The GraphQL request to be sent.
	//	responseData interface{}: The interface to store the response data. The response data will be unmarshaled into this interface.
	Request(ctx context.Context, req *Request, responseData interface{}) error

	// Subscribe to a GraphQL endpoint and streams events to the provided handler
	//
	//	ctx context.Context: The context for the subscription. If provided, it will be used to cancel the subscription if the context is canceled.
	//	req *Request: The GraphQL subscription request to be sent.
	//	handler EventResponseHandler: The handler function to process incoming events.
	Subscribe(ctx context.Context, req *Request, handler EventResponseHandler) error
}

// hubClient client to talk to Tanzu Hub through GraphQL APIs
type hubClient struct {
	// contextName is Tanzu CLI context name
	contextName string

	accessToken      string
	tanzuHubEndpoint string
	httpClient       *http.Client
}

type ClientOptions func(o *hubClient)

// WithAccessToken creates the Client using the specified Access Token
func WithAccessToken(token string) ClientOptions {
	return func(c *hubClient) {
		c.accessToken = token
	}
}

// WithEndpoint creates the Client using the specified Endpoint
func WithEndpoint(endpoint string) ClientOptions {
	return func(c *hubClient) {
		c.tanzuHubEndpoint = endpoint
	}
}

// WithHTTPClient creates the Client using the specified HttpClient
func WithHTTPClient(httpClient *http.Client) ClientOptions {
	return func(c *hubClient) {
		c.httpClient = httpClient
	}
}

// NewClient returns an authenticated Tanzu Hub client for the specified
// tanzu context. Internally it configures the client with CSP access token for each request
//
// Note that the authenticated client is assured to have at least 30 min access to the GraphQL endpoint.
// If you want a long running client beyond this period, recommendation is to reinitialize your client.
//
// EXPERIMENTAL: Both the function's signature and implementation are subjected to change/removal
// if an alternative means to provide equivalent functionality can be introduced.
func NewClient(contextName string, opts ...ClientOptions) (Client, error) {
	c := &hubClient{
		contextName: contextName,
	}

	// configure all options for the HubClient
	for _, o := range opts {
		o(c)
	}

	err := c.initializeClient(contextName)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *hubClient) initializeClient(contextName string) error {
	var err error

	// Set accessToken if it is not already set
	if c.accessToken == "" {
		c.accessToken, err = config.GetTanzuContextAccessToken(contextName)
		if err != nil {
			return err
		}
	}

	// Set tanzuHubEndpoint if it is not already set
	if c.tanzuHubEndpoint == "" {
		c.tanzuHubEndpoint, err = getTanzuHubEndpointFromContext(contextName)
		if err != nil {
			return err
		}
	}

	transport := http.DefaultTransport.(*http.Transport)
	transport.TLSClientConfig = c.getTLSConfig()

	// Set httpClient if it is not already set
	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Transport: &authTransport{
				accessToken: c.accessToken,
				wrapped:     transport,
			},
			Timeout: 0,
		}
	}

	return nil
}

func getTanzuHubEndpointFromContext(contextName string) (string, error) {
	// If `TANZU_HUB_ENDPOINT` environment variable is configured use that
	if endpoint := os.Getenv(EnvTanzuHubEndpoint); endpoint != "" {
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

func (c *hubClient) getTLSConfig() *tls.Config {
	// If the certificate information is found for the hub endpoint
	// then configure TLSConfig and return it
	certData, err := config.GetCert(c.tanzuHubEndpoint)
	if err != nil || certData == nil {
		return nil
	}

	// If CACertData is present use it
	if certData.CACertData != "" {
		var pool *x509.CertPool

		decodedCACertData, err := base64.StdEncoding.DecodeString(certData.CACertData)
		if err != nil {
			log.Infof("unable to use custom cert for '%s' endpoint. Error: %s", c.tanzuHubEndpoint, err.Error())
			return nil
		}

		pool, err = x509.SystemCertPool()
		if err != nil || pool == nil {
			pool = x509.NewCertPool()
		}

		if ok := pool.AppendCertsFromPEM(decodedCACertData); !ok {
			log.Infof("unable to use custom cert for %s endpoint", c.tanzuHubEndpoint)
			return nil
		}
		return &tls.Config{RootCAs: pool, MinVersion: tls.VersionTLS12}
	}

	skipCertVerify, _ := strconv.ParseBool(certData.SkipCertVerify)
	if skipCertVerify {
		//nolint:gosec
		// skipTLSVerify: true is only possible if the user has explicitly enabled it
		return &tls.Config{InsecureSkipVerify: skipCertVerify, MinVersion: tls.VersionTLS12}
	}

	return nil
}

// Configure the auth Transport to include authorization token when invoking GraphQL requests
type authTransport struct {
	accessToken string
	wrapped     *http.Transport
}

// Sets authentication bearer token to each http request
func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("authorization", "Bearer "+t.accessToken)
	return t.wrapped.RoundTrip(req)
}
