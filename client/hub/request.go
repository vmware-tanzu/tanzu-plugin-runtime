// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package hub provides functions to create Tanzu Hub client for specific context
package hub

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Request Sends a GraphQL request to the Tanzu Hub endpoint
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
	reader := bufio.NewReader(httpResp.Body)

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

func waitForEvents(reader *bufio.Reader, eventChan chan EventResponse) error {
	ev := EventResponse{}
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return errors.Errorf("Error during resp.Body read: %s", err.Error())
		}

		switch {
		case hasPrefix(line, ":"):
			// Comment, do nothing
		case hasPrefix(line, "retry:"):
			// Retry, do nothing for now
		case hasPrefix(line, "event:"):
			// Retry, do nothing for now

		// event data
		case hasPrefix(line, "data: "):
			buf.Write(line[6:])
		case hasPrefix(line, "data:"):
			buf.Write(line[5:])

		// end of event
		case bytes.Equal(line, []byte("\n")):
			b := buf.Bytes()

			if hasPrefix(b, "{") {
				var resp Response
				err := json.Unmarshal(b, &resp)
				if err != nil {
					return errors.Errorf("Error unmarshaling the event response. Error:%s", err.Error())
				}
				ev.Data = resp
				buf.Reset()
				eventChan <- ev
				ev = EventResponse{}
			}

		default:
			return errors.Errorf("Error: len:%d\n%s", len(line), line)
		}
	}
}

func hasPrefix(s []byte, prefix string) bool {
	return bytes.HasPrefix(s, []byte(prefix))
}
