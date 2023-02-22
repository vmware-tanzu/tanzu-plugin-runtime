// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package framework_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestFramework(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Framework Suite")
}
