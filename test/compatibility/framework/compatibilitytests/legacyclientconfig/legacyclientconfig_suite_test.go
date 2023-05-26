// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package legacyclientconfig_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLegacyClientConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cross-version API Compatibility Test Suite for Legacy Client Config APIs on supported Runtime libraries v0.28.0, v0.25.4, v0.11.6 and latest")
}
