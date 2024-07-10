// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package hub provides functions to create Tanzu Hub client for specific context
package hub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/r3labs/sse/v2"
)

// Request sends a GraphQL request to the Tanzu Hub endpoint
//
//	ctx context.Context: The context for the request. If provided, it will be used to cancel the request if the context is canceled.
//	req *Request: The GraphQL request to be sent.
//	responseData interface{}: The interface to store the response data. The response data will be unmarshaled into this interface.
func (c *HubClient) Request(ctx context.Context, req *Request, responseData interface{}) error {
	resp := &Response{Data: responseData}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/graphql", c.tanzuHubEndpoint), bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		var respBody []byte
		respBody, err = io.ReadAll(httpResp.Body)
		if err != nil {
			respBody = []byte(fmt.Sprintf("<unreadable: %v>", err))
		}
		return fmt.Errorf("returned error %v: %s", httpResp.Status, respBody)
	}

	err = json.NewDecoder(httpResp.Body).Decode(resp)
	if err != nil {
		return err
	}
	if len(resp.Errors) > 0 {
		return resp.Errors
	}
	return nil
}

// Subscribe to a GraphQL endpoint and streams events to the provided handler
//
//	ctx context.Context: The context for the subscription. If provided, it will be used to cancel the subscription if the context is canceled.
//	req *Request: The GraphQL subscription request to be sent.
//	handler EventResponseHandler: The handler function to process incoming events.
func (c *HubClient) Subscribe(ctx context.Context, req *Request, handler EventResponseHandler) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/subscriptions", c.tanzuHubEndpoint), bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept-Encoding", "gzip, deflate, br")
	httpReq.Header.Set("Accept", "text/event-stream")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		var respBody []byte
		respBody, err = io.ReadAll(httpResp.Body)
		if err != nil {
			respBody = []byte(fmt.Sprintf("<unreadable: %v>", err))
		}
		return fmt.Errorf("returned error %v: %s", httpResp.Status, respBody)
	}

	eventChan := make(chan EventResponse)
	errChan := make(chan error)
	reader := sse.NewEventStreamReader(httpResp.Body, 1024)

	go func() {
		errChan <- waitForEvents(reader, eventChan)
	}()

	for {
		select {
		case e := <-eventChan:
			handler(e)
		case err := <-errChan:
			if err != nil {
				return err
			}
			// goroutine has finished, exit loop
			return nil
		case <-ctx.Done():
			// context was canceled, exit loop
			return ctx.Err()
		}
	}
}

func waitForEvents(reader *sse.EventStreamReader, eventChan chan EventResponse) error {
	for {
		// Read each new line and process the type of event
		event, err := reader.ReadEvent()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		// If we get an error, ignore it.
		var eventMsg *EventResponse
		if eventMsg, err = processEvent(event); err == nil {
			// Send downstream if the event has something useful
			if eventMsg != nil {
				// Try to unMarshal the event Data into Response format
				if bytes.HasPrefix(eventMsg.RawData, []byte("{")) {
					var resp Response
					err := json.Unmarshal(eventMsg.RawData, &resp)
					if err == nil {
						eventMsg.ResponseData = &resp
					}
				}
				eventChan <- *eventMsg
			}
		}
	}
}

var (
	headerID    = []byte("id:")
	headerData  = []byte("data:")
	headerEvent = []byte("event:")
	headerRetry = []byte("retry:")
)

func processEvent(msg []byte) (event *EventResponse, err error) {
	var e EventResponse

	if len(msg) < 1 {
		return nil, errors.New("event message was empty")
	}

	// Normalize the crlf to lf to make it easier to split the lines.
	// Split the line by "\n" or "\r", per the spec.
	for _, line := range bytes.FieldsFunc(msg, func(r rune) bool { return r == '\n' || r == '\r' }) {
		switch {
		case bytes.HasPrefix(line, headerID):
			e.ID = string(trimHeader(len(headerID), line))
		case bytes.HasPrefix(line, headerData):
			// The spec allows for multiple data fields per event, concatenated them with "\n".
			e.RawData = append(e.RawData, append(trimHeader(len(headerData), line), byte('\n'))...)
		// The spec says that a line that simply contains the string "data" should be treated as a data field with an empty body.
		case bytes.Equal(line, bytes.TrimSuffix(headerData, []byte(":"))):
			e.RawData = append(e.RawData, byte('\n'))
		case bytes.HasPrefix(line, headerEvent):
			e.Name = string(trimHeader(len(headerEvent), line))
		case bytes.HasPrefix(line, headerRetry):
			e.Retry = string(trimHeader(len(headerRetry), line))
		default:
			// Ignore any garbage that doesn't match what we're looking for.
		}
	}

	// Trim the last "\n" per the spec.
	e.RawData = bytes.TrimSuffix(e.RawData, []byte("\n"))

	return &e, err
}

func trimHeader(size int, data []byte) []byte {
	if data == nil || len(data) < size {
		return data
	}

	data = data[size:]
	// Remove optional leading whitespace
	if len(data) > 0 && data[0] == 32 {
		data = data[1:]
	}
	// Remove trailing new line
	if len(data) > 0 && data[len(data)-1] == 10 {
		data = data[:len(data)-1]
	}
	return data
}
