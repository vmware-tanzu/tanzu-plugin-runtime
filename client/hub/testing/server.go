// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// NOTE: A portion of this file is adapted from github.com/getoutreach/goql
// and some modifications were made on top of the original file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package testing exports a GraphQL Mock Server that facilitates
// the testing of client.
package testing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

type Server struct {
	URL string

	subscriptions []Operation
	mutations     []Operation
	queries       []Operation
	errors        []OperationError

	t      *testing.T
	server *httptest.Server
}

type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type ResponseError struct {
	Message    string      `json:"message"`
	Path       []string    `json:"path"`
	Extensions interface{} `json:"extensions"`
}

type Response struct {
	Data   interface{}     `json:"data"`
	Errors []ResponseError `json:"errors,omitempty"`
}

type ServerOptions func(o *Server)

// WithQuery registers mock Query operations to the server
func WithQuery(operations ...Operation) ServerOptions {
	return func(s *Server) {
		for _, o := range operations {
			s.RegisterQuery(o)
		}
	}
}

// WithMutation registers mock Mutation operations to the server
func WithMutation(operations ...Operation) ServerOptions {
	return func(s *Server) {
		for _, o := range operations {
			s.RegisterMutation(o)
		}
	}
}

// WithSubscriptions registers mock Subscriptions operations to the server
func WithSubscriptions(operations ...Operation) ServerOptions {
	return func(s *Server) {
		for _, o := range operations {
			s.RegisterSubscription(o)
		}
	}
}

// WithErrors registers mock OperationError to the server
func WithErrors(operations []OperationError) ServerOptions {
	return func(s *Server) {
		for _, o := range operations {
			s.RegisterError(o)
		}
	}
}

// NewServer returns a Mock Server object. The server object returned
// contains a closing function which should be immediately registered using t.Cleanup
// after calling NewServer, example:
//
//	ts := testing.NewServer(t)
//	t.Cleanup(ts.Close)
//
// If you want to reuse a server across multiple unit tests than use ts.Reset()
// to clean up any already registered queries, mutations or errors
func NewServer(t *testing.T, opts ...ServerOptions) *Server { //nolint:gocyclo
	s := Server{
		t: t,
	}

	for _, o := range opts {
		o(&s)
	}

	var mux http.ServeMux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var reqBody Request
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			s.respondError(w, http.StatusInternalServerError, errors.Wrap(err, "decode request body"), nil)
			return
		}

		switch {
		case strings.HasPrefix(strings.TrimSpace(reqBody.Query), "mutation"):
			for i := range s.mutations {
				if strings.Contains(reqBody.Query, s.mutations[i].Identifier) {
					if s.equalVariables(s.mutations[i].Variables, reqBody.Variables) {
						s.respond(w, http.StatusOK, s.mutations[i].Response)
						return
					}
				}
			}
		case strings.HasPrefix(strings.TrimSpace(reqBody.Query), "query"):
			for i := range s.queries {
				if strings.Contains(reqBody.Query, s.queries[i].Identifier) {
					if s.equalVariables(s.queries[i].Variables, reqBody.Variables) {
						s.respond(w, http.StatusOK, s.queries[i].Response)
						return
					}
				}
			}
		case strings.HasPrefix(strings.TrimSpace(reqBody.Query), "subscription"):
			for i := range s.subscriptions {
				if strings.Contains(reqBody.Query, s.subscriptions[i].Identifier) {
					if s.equalVariables(s.subscriptions[i].Variables, reqBody.Variables) {
						flusher, ok := w.(http.Flusher)
						if !ok {
							http.Error(w, "SSE not supported", http.StatusInternalServerError)
							return
						}

						w.Header().Set("Content-Type", "text/event-stream")

						respChan := make(chan Response)
						go s.subscriptions[i].EventGenerator(r.Context(), respChan)

						for eventResp := range respChan {
							event, err := formatServerSentEvent("update", eventResp)
							if err != nil {
								fmt.Println(err)
								s.respond(w, http.StatusInternalServerError, err.Error())
								break
							}

							_, err = fmt.Fprint(w, event)
							if err != nil {
								fmt.Println(err)
								s.respond(w, http.StatusInternalServerError, err.Error())
								break
							}

							flusher.Flush()
						}
						s.respond(w, http.StatusOK, s.subscriptions[i].Response)
						return
					}
				}
			}
		case strings.HasPrefix(strings.TrimSpace(reqBody.Query), "error"):
			for i := range s.errors {
				if strings.Contains(reqBody.Query, s.errors[i].Identifier) {
					s.respondError(w, s.errors[i].Status, s.errors[i].Error, s.errors[i].Extensions)
					return
				}
			}
		}

		s.respondError(w, http.StatusNotFound, errors.New("operation not found"), nil)
	})

	s.server = httptest.NewServer(&mux)
	s.URL = s.server.URL

	return &s
}

// Close closes the underlying httptest.Server.
func (s *Server) Close() {
	s.server.Close()
}

// Mutations returns the registered mutations that the server will accept and respond
// to.
func (s *Server) Mutations() []Operation {
	return s.mutations
}

// Queries returns the registered queries that the server will accept and respond to.
func (s *Server) Queries() []Operation {
	return s.queries
}

// Subscriptions returns the registered subscriptions that the server will accept and respond to.
func (s *Server) Subscriptions() []Operation {
	return s.subscriptions
}

// RegisterQuery registers an Operation as a query that the server will recognize and
// respond to.
func (s *Server) RegisterQuery(operations ...Operation) {
	for _, o := range operations {
		o.opType = opQuery
		s.queries = append(s.queries, o)
	}
}

// RegisterMutation registers an Operation as a mutation that the server will recognize
// and respond to.
func (s *Server) RegisterMutation(operations ...Operation) {
	for _, o := range operations {
		o.opType = opMutation
		s.mutations = append(s.mutations, o)
	}
}

// RegisterSubscription registers an Operation as a subscription that the server will recognize
// and respond to.
func (s *Server) RegisterSubscription(operations ...Operation) {
	for _, o := range operations {
		o.opType = opSubscription
		s.subscriptions = append(s.subscriptions, o)
	}
}

// RegisterError registers an OperationError as an error that the server will recognize
// and respond to.
func (s *Server) RegisterError(operation OperationError) {
	s.errors = append(s.errors, operation)
}

// Reset resets the existing mocked responses that are already registered with the server
func (s *Server) Reset() {
	s.queries = []Operation{}
	s.mutations = []Operation{}
	s.subscriptions = []Operation{}
	s.errors = []OperationError{}
}

// Do takes a Request, performs it using the underlying httptest.Server, and returns a
// Response.
func (s *Server) Do(r Request) Response {
	s.t.Helper()

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(r); err != nil {
		s.t.Fatalf("encode graphql request body: %v", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, s.URL, &buf)
	if err != nil {
		s.t.Fatalf("create graphql request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		s.t.Errorf("do graphql request: %v", err)
	}
	defer res.Body.Close()

	var resBody Response
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		s.t.Errorf("decode graphql response body: %v", err)
	}

	return resBody
}

// equalVariables takes two variables and makes sure they are equal in length and
// each contain the same keys. The values of the keys are not checked.
func (s *Server) equalVariables(x, y map[string]interface{}) bool {
	if len(x) != len(y) {
		return false
	}

	for k := range x {
		if _, exists := y[k]; !exists {
			return false
		}
	}

	for k := range y {
		if _, exists := x[k]; !exists {
			return false
		}
	}

	return true
}

func (s *Server) respondError(w http.ResponseWriter, status int, err error, extensions interface{}) {
	s.t.Helper()

	res := Response{
		Data: nil,
	}

	res.Errors = append(res.Errors, ResponseError{
		Message:    err.Error(),
		Extensions: extensions,
	})

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		s.t.Errorf("encode graphql error response: %v", err)
	}
}

func (s *Server) respond(w http.ResponseWriter, status int, data interface{}) {
	s.t.Helper()

	res := Response{
		Data:   data,
		Errors: nil,
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		s.t.Errorf("encode graphql response: %v", err)
	}
}

func formatServerSentEvent(event string, data any) (string, error) {
	buff := bytes.NewBuffer([]byte{})

	encoder := json.NewEncoder(buff)

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", buff.String()))

	return sb.String(), nil
}
