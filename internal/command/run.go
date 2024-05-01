// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package command provides functions to invoke tanzu cli commands
package command

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
)

// CmdOptions specifies the command options
type CmdOptions struct {
	outWriter io.Writer
	errWriter io.Writer
}

type CommandOptions func(o *CmdOptions)

// WithOutputWriter specifies the CommandOption for configuring Stdout
func WithOutputWriter(outWriter io.Writer) CommandOptions {
	return func(o *CmdOptions) {
		o.outWriter = outWriter
	}
}

// WithErrorWriter specifies the CommandOption for configuring Stderr
func WithErrorWriter(errWriter io.Writer) CommandOptions {
	return func(o *CmdOptions) {
		o.errWriter = errWriter
	}
}

// WithNoStdout specifies to ignore stdout
func WithNoStdout() CommandOptions {
	return func(o *CmdOptions) {
		o.outWriter = io.Discard
	}
}

// WithNoStderr specifies to ignore stderr
func WithNoStderr() CommandOptions {
	return func(o *CmdOptions) {
		o.errWriter = io.Discard
	}
}

// RunTanzuCommand invokes a Tanzu CLI command
func RunTanzuCommand(args []string, cmdOpts ...CommandOptions) (bytes.Buffer, bytes.Buffer, error) {
	cliPath := os.Getenv("TANZU_BIN")
	if cliPath == "" {
		return bytes.Buffer{}, bytes.Buffer{}, errors.New("the environment variable TANZU_BIN is not set")
	}

	opts := &CmdOptions{}
	for _, o := range cmdOpts {
		o(opts)
	}

	command := exec.Command(cliPath, args...)

	var stderr bytes.Buffer
	var stdout bytes.Buffer

	wout := io.MultiWriter(&stdout, os.Stdout)
	werr := io.MultiWriter(&stderr, os.Stderr)

	if opts.outWriter != nil {
		wout = io.MultiWriter(&stdout, opts.outWriter)
	}
	if opts.errWriter != nil {
		werr = io.MultiWriter(&stderr, opts.errWriter)
	}

	command.Stdout = wout
	command.Stderr = werr

	return stdout, stderr, command.Run()
}
