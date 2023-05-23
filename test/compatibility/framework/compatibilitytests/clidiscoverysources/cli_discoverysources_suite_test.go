// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clidiscoverysources_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCLIDiscoverySources(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cross-version API Compatibility Test Suite for CLI Discovery Sources")
}
