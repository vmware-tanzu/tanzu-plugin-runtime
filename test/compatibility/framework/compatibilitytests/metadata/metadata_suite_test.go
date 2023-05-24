// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package metadata_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfigMetadata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cross-version API Compatibility Test Suite for Config Metadata APIs for version latest, v0.28.0")
}
