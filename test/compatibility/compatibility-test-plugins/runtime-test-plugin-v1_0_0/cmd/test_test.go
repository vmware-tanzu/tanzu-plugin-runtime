// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	compatibilitytestingframework "github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = Describe("Test Root method", func() {

	BeforeEach(func() {
		compatibilitytestingframework.SetupTempCfgFiles()
	})

	Context("Test runAPIs", func() {

		It("runAPIs with SetContext api data", func() {

			apis, err := compatibilitytestingframework.ParseRuntimeAPIsFromFile("/var/folders/gw/3kzrkntn5rzbs4xhktdvdfjr0000gq/T/runtime_compatibility_testing4086182328")

			Expect(err).To(BeNil())

			runAPIs(apis)
		})
	})

})
