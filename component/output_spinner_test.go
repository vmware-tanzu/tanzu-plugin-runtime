// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
)

const loading = "Loading..."

func TestNewOutputWriterWithSpinner(t *testing.T) {
	output := bytes.Buffer{}
	spinnerText := loading
	headers := []string{"Name", "Age"}

	// Test creating an OutputWriterSpinner with a spinner
	ows, err := NewOutputWriterWithSpinner(&output, "table", spinnerText, true, headers...)
	assert.NoError(t, err)
	assert.NotNil(t, ows)

	// Test creating an OutputWriterSpinner without a spinner
	ows, err = NewOutputWriterWithSpinner(&output, "table", spinnerText, false, headers...)
	assert.NoError(t, err)
	assert.NotNil(t, ows)

	// Test creating an OutputWriterSpinner with unsupported output format
	ows, err = NewOutputWriterWithSpinner(&output, "unsupported", spinnerText, true, headers...)
	assert.NoError(t, err)
	assert.NotNil(t, ows)
}

func TestNewOutputWriterSpinnerWithOptions(t *testing.T) {
	output := bytes.Buffer{}
	spinnerText := loading
	headers := []string{"Name", "Age"}

	// Test creating an OutputWriterSpinner with options and a spinner
	opts := []OutputWriterOption{WithAutoStringify()}
	ows, err := NewOutputWriterSpinnerWithOptions(&output, "table", spinnerText, true, opts, headers...)
	assert.NoError(t, err)
	assert.NotNil(t, ows)

	// Test creating an OutputWriterSpinner with options without a spinner
	opts = []OutputWriterOption{WithAutoStringify()}
	ows, err = NewOutputWriterSpinnerWithOptions(&output, "table", spinnerText, false, opts, headers...)
	assert.NoError(t, err)
	assert.NotNil(t, ows)

	// Test creating an OutputWriterSpinner with unsupported output format
	opts = []OutputWriterOption{WithAutoStringify()}
	ows, err = NewOutputWriterSpinnerWithOptions(&output, "unsupported", spinnerText, true, opts, headers...)
	assert.NoError(t, err)
	assert.NotNil(t, ows)
}

func TestNewOutputWriterSpinner(t *testing.T) {
	output := bytes.Buffer{}
	spinnerText := loading
	headers := []string{"Name", "Age"}

	ows := NewOutputWriterSpinner(WithOutputStream(&output),
		WithOutputFormat(TableOutputType),
		WithSpinnerText(spinnerText),
		WithSpinnerStarted(),
		WithOutputWriterOptions(WithAutoStringify()),
		WithHeaders(headers...),
		WithSpinnerFinalText("Done!", log.LogTypeSUCCESS))
	assert.NotNil(t, ows)

	ows = NewOutputWriterSpinner(WithOutputStream(&output),
		WithOutputFormat(TableOutputType),
		WithSpinnerText(spinnerText),
		WithHeaders(headers...),
		WithSpinnerFinalText("Done!", log.LogTypeSUCCESS))
	assert.NotNil(t, ows)

	ows = NewOutputWriterSpinner(WithOutputStream(&output),
		WithOutputFormat("unsupported"),
		WithSpinnerText(spinnerText),
		WithSpinnerStarted(),
		WithHeaders(headers...))
	assert.NotNil(t, ows)
}

func TestOutputWriterSpinnerRenderWithSpinner(t *testing.T) {
	output := bytes.Buffer{}
	spinnerText := loading
	headers := []string{"Name", "Age"}

	// Create an OutputWriterSpinner with a spinner
	ows, err := NewOutputWriterWithSpinner(&output, "table", spinnerText, true, headers...)
	ows.AddRow(map[string]interface{}{"Name": "John", "Age": 30})
	assert.NoError(t, err)
	assert.NotNil(t, ows)

	// Render with spinner
	ows.RenderWithSpinner()
	assert.Contains(t, output.String(), "NAME")
	assert.Contains(t, output.String(), "John")
	assert.Contains(t, output.String(), "30")
}
