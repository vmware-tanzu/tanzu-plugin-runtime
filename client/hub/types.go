// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package hub provides functions to create Tanzu Hub client for specific context
package hub

import "github.com/vektah/gqlparser/v2/gqlerror"

// Request contains all the values required to build queries executed by
// the Client.
//
// Typically, GraphQL APIs will accept a JSON payload of the form
//
//	{"query": "query myQuery { ... }", "variables": {...}}`
//
// and Request marshals to this format.
type Request struct {
	// The literal string representing the GraphQL query, e.g.
	// `query myQuery { myField }`.
	Query string `json:"query"`
	// A JSON-marshalable value containing the variables to be sent
	// along with the query, or nil if there are none.
	Variables interface{} `json:"variables,omitempty"`
	// The GraphQL operation name. The server typically doesn't
	// require this unless there are multiple queries in the
	// document.
	OpName string `json:"operationName,omitempty"`
}

// Response that contains data returned by the GraphQL API.
//
// Typically, GraphQL APIs will return a JSON payload of the form
//
//	{"data": {...}, "errors": {...}}
//
// It may additionally contain a key named "extensions", that
// might hold GraphQL protocol extensions. Extensions and Errors
// are optional, depending on the values returned by the server.
type Response struct {
	Data       interface{}            `json:"data"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	Errors     gqlerror.List          `json:"errors,omitempty"`
}

// EventResponse represents a Server-Sent event response
type EventResponse struct {
	// Name contains the value of the "event:" header from the event stream.
	Name string
	// ID contains the value of the "id:" header from the event stream.
	ID string
	// RawData contains the concatenated payload from the "data:" headers received as part of the event stream.
	RawData []byte
	// ResponseData contains the parsed GraphQL response object if the RawData can be successfully parsed to Response object, otherwise it is nil.
	ResponseData *Response
	// Retry contains the value of the "retry:" header from the event stream.
	Retry string
}

// EventResponseHandler represents a Subscription event handler function
// that will be passed to the `Subscribe` method
type EventResponseHandler func(eventResponse EventResponse)
