// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mattn/go-isatty"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
)

// OutputWriterSpinner is OutputWriter augmented with a spinner.
type OutputWriterSpinner interface {
	OutputWriter
	// RenderWithSpinner will stop spinner and render the output
	// Deprecated: RenderWithSpinner is being deprecated in favor of Render.
	RenderWithSpinner()
	// StartSpinner start the spinner instance and renders the spinnerText
	StartSpinner()
	// StopSpinner stops the running spinner instance and renders FinalText if set
	StopSpinner()
	// SetText sets the spinner text
	SetText(text string)
	// SetFinalText sets the spinner final text and prefix
	// log indicator (log.LogTypeOUTPUT can be used for no prefix)
	SetFinalText(finalText string, prefix log.LogType)
}

// outputwriterspinner is our internal implementation.
type outputwriterspinner struct {
	outputwriter
	spinnerText        string
	spinnerFinalText   string
	startSpinnerOnInit bool
	spinner            *spinner.Spinner
}

// OutputWriterSpinnerOption is an option to configure outputwriterspinner
type OutputWriterSpinnerOption func(*outputwriterspinner)

// WithSpinnerFinalText sets the spinner final text and prefix log indicator
// (log.LogTypeOUTPUT can be used for no prefix)
func WithSpinnerFinalText(finalText string, prefix log.LogType) OutputWriterSpinnerOption {
	return func(ows *outputwriterspinner) {
		ows.spinnerFinalText = fmt.Sprintf("%s%s", log.GetLogTypeIndicator(prefix), finalText)
	}
}

// WithOutputWriterOptions configures OutputWriterOptions to the spinner
func WithOutputWriterOptions(opts ...OutputWriterOption) OutputWriterSpinnerOption {
	return func(ow *outputwriterspinner) {
		ow.applyOptions(opts)
	}
}

// WithSpinnerText sets the spinner text
func WithSpinnerText(text string) OutputWriterSpinnerOption {
	return func(ows *outputwriterspinner) {
		ows.spinnerText = text
	}
}

// WithSpinnerStarted starts the spinner
func WithSpinnerStarted() OutputWriterSpinnerOption {
	return func(ows *outputwriterspinner) {
		ows.startSpinnerOnInit = true
	}
}

// WithHeaders sets key headers
func WithHeaders(headers ...string) OutputWriterSpinnerOption {
	return func(ows *outputwriterspinner) {
		ows.keys = headers
	}
}

// WithOutputFormat sets output format for the OutputWriterSpinner component
func WithOutputFormat(outputFormat OutputType) OutputWriterSpinnerOption {
	return func(ows *outputwriterspinner) {
		ows.outputFormat = outputFormat
	}
}

// WithOutputStream sets the output stream for the OutputWriterSpinner component
func WithOutputStream(writer io.Writer) OutputWriterSpinnerOption {
	return func(ows *outputwriterspinner) {
		ows.out = writer
	}
}

// NewOutputWriterWithSpinner returns implementation of OutputWriterSpinner.
//
// Deprecated: NewOutputWriterWithSpinner is being deprecated in favor of
// NewOutputWriterSpinner.
// Until it is removed, it will retain the existing behavior of converting
// incoming row values to their golang string representation for backward
// compatibility reasons
func NewOutputWriterWithSpinner(output io.Writer, outputFormat, spinnerText string, startSpinner bool, headers ...string) (OutputWriterSpinner, error) {
	opts := []OutputWriterOption{WithAutoStringify()}
	return NewOutputWriterSpinnerWithOptions(output, outputFormat, spinnerText, startSpinner, opts, headers...)
}

// NewOutputWriterSpinnerWithOptions returns implementation of OutputWriterSpinner.
//
// Deprecated: NewOutputWriterSpinnerWithOptions is being deprecated in favor of
// NewOutputWriterSpinner.
func NewOutputWriterSpinnerWithOptions(output io.Writer, outputFormat, spinnerText string, startSpinner bool, opts []OutputWriterOption, headers ...string) (OutputWriterSpinner, error) {
	ows := &outputwriterspinner{}
	ows.out = output
	ows.outputFormat = OutputType(outputFormat)
	ows.keys = headers
	ows.applyOptions(opts)
	ows.spinnerText = spinnerText
	ows.startSpinnerOnInit = startSpinner
	return initializeSpinner(ows), nil
}

// NewOutputWriterSpinner returns implementation of OutputWriterSpinner
func NewOutputWriterSpinner(opts ...OutputWriterSpinnerOption) (OutputWriterSpinner, error) {
	ows := &outputwriterspinner{}
	ows.out = os.Stdout
	ows.applySpinnerOptions(opts)
	// Note: We are returning the 'nil' error always to protect against the possible API
	// enhancement which might require throwing an error in future
	return initializeSpinner(ows), nil
}

// initializeSpinner initializes the spinner
func initializeSpinner(ows *outputwriterspinner) OutputWriterSpinner {
	if ows.outputFormat != JSONOutputType && ows.outputFormat != YAMLOutputType {
		ows.spinner = spinner.New(spinner.CharSets[9], 100*time.Millisecond,
			spinner.WithWriter(ows.out),
			spinner.WithFinalMSG(ows.spinnerFinalText),
			spinner.WithSuffix(fmt.Sprintf(" %s", ows.spinnerText)),
		)

		// If the colors cannot be assigned to the spinner, we do not want to throw and error
		// as spinner will use default coloring. Note this color scheme is only applicable to spinner object
		// and not to the text we display along with the spinner object
		_ = ows.spinner.Color("bgBlack", "bold", "fgWhite")

		// Start the spinner only if attached to terminal
		attachedToTerminal := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
		if ows.startSpinnerOnInit && attachedToTerminal {
			ows.spinner.Start()
		}
	}
	return ows
}

// RenderWithSpinner will stop spinner and render the output
//
// Deprecated: RenderWithSpinner is being deprecated in favor of Render.
func (ows *outputwriterspinner) RenderWithSpinner() {
	ows.Render()
}

// Render will stop spinner and render the output
func (ows *outputwriterspinner) Render() {
	ows.StopSpinner()
	ows.outputwriter.Render()
}

// StartSpinner start the spinner instance and renders the spinnerText
func (ows *outputwriterspinner) StartSpinner() {
	attachedToTerminal := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	if ows.spinner != nil && !ows.spinner.Active() && attachedToTerminal {
		ows.spinner.Start()
	}
}

// StopSpinner stops the running spinner instance and renders FinalText if set
func (ows *outputwriterspinner) StopSpinner() {
	if ows.spinner != nil && ows.spinner.Active() {
		ows.spinner.Stop()
		fmt.Fprintln(ows.out)
	}
}

// SetFinalText sets the spinner final text and prefix log indicator
// (log.LogTypeOUTPUT can be used for no prefix)
func (ows *outputwriterspinner) SetFinalText(finalText string, prefix log.LogType) {
	if ows.spinner != nil {
		ows.spinnerFinalText = fmt.Sprintf("%s%s", log.GetLogTypeIndicator(prefix), finalText)
		spinner.WithFinalMSG(ows.spinnerFinalText)(ows.spinner)
	}
}

// SetText sets the spinner text
func (ows *outputwriterspinner) SetText(text string) {
	if ows.spinner != nil {
		ows.spinnerText = text
		ows.spinner.Suffix = fmt.Sprintf(" %s", text)
	}
}

// applySpinnerOptions applies the options to the outputwriterspinner
func (ows *outputwriterspinner) applySpinnerOptions(spinnerOpts []OutputWriterSpinnerOption) {
	for i := range spinnerOpts {
		spinnerOpts[i](ows)
	}
}
