package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers"
)

var _ = Describe("Test Root method", func() {

	BeforeEach(func() {
		helpers.SetupTempCfgFiles()
	})

	Context("Test runAPIs", func() {

		It("runAPIs with SetContext api data", func() {

			apis, err := helpers.GetTestData("/var/folders/gw/3kzrkntn5rzbs4xhktdvdfjr0000gq/T/runtime_compatibility_testing4086182328")

			Expect(err).To(BeNil())

			runAPIs(apis)
		})
	})

})
