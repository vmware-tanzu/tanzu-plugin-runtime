// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package command

import (
	"errors"
	"fmt"
	"testing"
)

func TestSilenceError(t *testing.T) {
	err := fmt.Errorf("test error")
	silentErr := SilenceError(err)

	if errors.Is(err, SilentError) {
		t.Errorf("expected error to not be silent, got %#v", err)
	}
	if !errors.Is(silentErr, SilentError) {
		t.Errorf("expected error to be silent, got %#v", err)
	}
	if expected, actual := err, errors.Unwrap(silentErr); expected != actual {
		t.Errorf("errors expected to match, expected %v, actually %v", expected, actual)
	}
	if expected, actual := err.Error(), silentErr.Error(); expected != actual {
		t.Errorf("errors expected to match, expected %q, actually %q", expected, actual)
	}
}
