// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package envflags_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEnvFlags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cross-version API Compatibility Test Suite for Env flags")
}
