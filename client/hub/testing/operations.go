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

package testing

import (
	"context"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/client/hub"
)

// Operation Type Constants
const (
	opQuery = iota + 1
	opMutation
	opSubscription
)

// Operation is a general type that encompasses the Operation type and Response which
// is of the same type, but with data.
type Operation struct {
	// opType denotes whether the operation is a query, a mutation or a subscription, using the opQuery,
	// opMutation and opSubscription constants. This is unexported as it is set by the *Server.RegisterQuery,
	// *Server.RegisterMutation and *Server.RegisterSubscription functions, respectively.
	opType int

	// Identifier helps identify the operation in a request when coming through the Server.
	// For example, if your operation looks like this:
	//
	//	query {
	//		myOperation(foo: $foo) {
	//			fieldOne
	//			fieldTwo
	//		}
	//	}
	//
	// then this field should be set to myOperation. It can also be more specific, a simple
	// strings.Contains check occurs to match operations. A more specific example of a
	// valid Identifier for the same operation given above would be myOperation(foo: $foo).
	Identifier string

	// Variables represents the map of variables that should be passed along with the
	// operation whenever it is invoked on the Server.
	Variables map[string]interface{}

	// Response represents the response that should be returned whenever the server makes
	// a match on Operation.opType, Operation.Identifier, and Operation.Variables.
	// Response is to be used for Query and Mutation operations only.
	// Note: User can define either `Response` or implement `Responder` function but should
	// not be defining both.
	Response hub.Response

	// Responder implements the function that based on some operation parameters should respond
	// differently.
	// Tests that do not need flexibility in returning different responses based on the Operation
	// should just configure the `Response` field instead.
	// Responder is to be used for Query and Mutation operations only.
	// Note: User can define either `Response` or implement `Responder` function but should
	// not be defining both.
	Responder Responder

	// EventGenerator should implement a eventData generator for testing and
	// send mock event response to the `eventData` channel. To suggest end of
	// the event responses from server side, you can close the eventData channel
	// Note: This is only to be used for the Subscription where you will need to
	// mock the generation of the events. This should not be used with Query or Mutation.
	EventGenerator EventGenerator
}

// OperationError is a special type that brings together the properties that a
// response error can include.
type OperationError struct {
	// Identifier helps identify the operation error in a request when coming through the Server.
	// For example, if your operation looks like this:
	//
	//	error {
	//		myOperation(foo: $foo) {
	//			fieldOne
	//			fieldTwo
	//		}
	//	}
	//
	// Then this field should be set to myOperation. It can also be more specific, a simple
	// strings.Contains check occurs to match operations. A more specific example of a
	// valid Identifier for the same operation given above would be myOperation(foo: $foo).
	Identifier string

	// Status represents the http status code that should be returned in the response
	// whenever the server makes a match on OperationError.Identifier
	Status int

	// Error represents the error that should be returned in the response whenever
	// the server makes a match on OperationError.Identifier
	Error error

	// Extensions represents the object that should be returned in the response
	// as part of the api error whenever the server makes a match on OperationError.Extensions
	Extensions interface{}
}

// Responder implements the function that based on some operation parameters should respond
// differently. This type of Responder implementation is more useful when you want
// to implement a generic function that returns data based on the operation
type Responder func(ctx context.Context, op Operation) hub.Response

// EventGenerator should implement a eventData generator for testing and
// send mock event response to the `eventData` channel. To suggest end of
// the event responses from server side, you can close the eventData channel
type EventGenerator func(ctx context.Context, op Operation, eventData chan<- Response)
