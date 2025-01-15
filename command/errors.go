// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package command

var SilentError = &silentError{}

type silentError struct {
	err error
}

func (e *silentError) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *silentError) Unwrap() error {
	return e.err
}

func (e *silentError) Is(err error) bool {
	_, ok := err.(*silentError)
	return ok
}

func SilenceError(err error) error {
	return &silentError{err: err}
}
