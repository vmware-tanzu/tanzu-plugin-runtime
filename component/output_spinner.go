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
	RenderWithSpinner()
	StopSpinner()
	// SetFinalText sets the spinner final text and prefix
	// log indicator (log.LogTypeOUTPUT can be used for no prefix)
	SetFinalText(finalText string, prefix log.LogType)
}

// outputwriterspinner is our internal implementation.
type outputwriterspinner struct {
	outputwriter
	spinnerText      string
	spinnerFinalText string
	spinner          *spinner.Spinner
}

type OutputWriterSpinnerOptions struct {
	OutputWriterOptions []OutputWriterOption
	SpinnerOptions      []OutputWriterSpinnerOption
}

// OutputWriterSpinnerOption is an option for outputwriterspinner
type OutputWriterSpinnerOption func(*outputwriterspinner)

// WithSpinnerFinalText sets the spinner final text and prefix log indicator
// (log.LogTypeOUTPUT can be used for no prefix)
func WithSpinnerFinalText(finalText string, prefix log.LogType) OutputWriterSpinnerOption {
	finalText = fmt.Sprintf("%s%s", log.GetLogTypeIndicator(prefix), finalText)
	return func(ows *outputwriterspinner) {
		ows.spinnerFinalText = finalText
	}
}

// NewOutputWriterWithSpinner returns implementation of OutputWriterSpinner.
//
// Deprecated: NewOutputWriterWithSpinner is being deprecated in favor of
// NewOutputWriterSpinnerWithSpinnerOptions.
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
// NewOutputWriterSpinnerWithSpinnerOptions.
func NewOutputWriterSpinnerWithOptions(output io.Writer, outputFormat, spinnerText string, startSpinner bool, opts []OutputWriterOption, headers ...string) (OutputWriterSpinner, error) {
	ows := &outputwriterspinner{}
	ows.out = output
	ows.outputFormat = OutputType(outputFormat)
	ows.keys = headers
	ows.applyOptions(opts)

	return setAndInitializeSpinner(ows, spinnerText, startSpinner)
}

// NewOutputWriterSpinnerWithSpinnerOptions returns implementation of OutputWriterSpinner.
func NewOutputWriterSpinnerWithSpinnerOptions(output io.Writer, outputFormat OutputType, spinnerText string, startSpinner bool, opts OutputWriterSpinnerOptions, headers ...string) (OutputWriterSpinner, error) {
	ows := &outputwriterspinner{}
	ows.out = output
	ows.outputFormat = outputFormat
	ows.keys = headers
	ows.applyOptions(opts.OutputWriterOptions)
	ows.applyOutputWriterSpinnerOptions(opts.SpinnerOptions)
	return setAndInitializeSpinner(ows, spinnerText, startSpinner)
}

// setAndInitializeSpinner sets the spinner text and initializes the spinner
func setAndInitializeSpinner(ows *outputwriterspinner, spinnerText string, startSpinner bool) (OutputWriterSpinner, error) {
	if ows.outputFormat != JSONOutputType && ows.outputFormat != YAMLOutputType {
		ows.spinnerText = spinnerText
		ows.spinner = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		if err := ows.spinner.Color("bgBlack", "bold", "fgWhite"); err != nil {
			return nil, err
		}
		ows.spinner.Suffix = fmt.Sprintf(" %s", spinnerText)
		if ows.spinnerFinalText != "" {
			spinner.WithFinalMSG(ows.spinnerFinalText)(ows.spinner)
		}

		// Start the spinner only if attached to terminal
		attachedToTerminal := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
		if startSpinner && attachedToTerminal {
			ows.spinner.Start()
		}
	}
	return ows, nil
}

// RenderWithSpinner will stop spinner and render the output
func (ows *outputwriterspinner) RenderWithSpinner() {
	if ows.spinner != nil && ows.spinner.Active() {
		ows.spinner.Stop()
		fmt.Fprintln(ows.out)
	}
	ows.Render()
}

// stop spinner
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

// applyOutputWriterSpinnerOptions applies the options to the outputwriterspinner
func (ows *outputwriterspinner) applyOutputWriterSpinnerOptions(spinnerOpts []OutputWriterSpinnerOption) {
	for i := range spinnerOpts {
		spinnerOpts[i](ows)
	}
}
