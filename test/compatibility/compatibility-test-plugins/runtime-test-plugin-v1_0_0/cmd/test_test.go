// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
)

var _ = ginkgo.Describe("Test Root method", func() {

	ginkgo.BeforeEach(func() {
		core.SetupTempCfgFiles()
	})

	ginkgo.Context("Test runAPIs", func() {
		ginkgo.It("runAPIs with SetContext api data", func() {
			apis, err := core.ParseRuntimeAPIsFromFile("/var/folders/gw/3kzrkntn5rzbs4xhktdvdfjr0000gq/T/runtime_compatibility_testing4086182328")
			gomega.Expect(err).To(gomega.BeNil())
			runAPIs(apis)
		})
	})

})
