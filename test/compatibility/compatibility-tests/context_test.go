// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package compatibility_tests_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

var _ = ginkgo.Describe("Context API Tests", func() {

	ginkgo.BeforeEach(func() {
		ginkgo.By("Setup mock temporary config files for testing")
		_, cleanup := core.SetupTempCfgFiles()

		ginkgo.DeferCleanup(func() {
			cleanup()
		})
	})

	ginkgo.Context("SetContext and GetContext APIs", func() {

		ginkgo.It("Run SetContext API of Runtime Library v1.0.0 then GetContext API of Runtime Library v0.28.0", func() {
			ginkgo.By("Construct Input and Output parameters for SetContext API of Runtime Library v1.0.0")
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   "context-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			var setContextOutputOptions *framework.SetContextOutputOptions

			ginkgo.By("Construct Input and Output parameters for GetContext API of Runtime Library v0.28.0")
			getContextInputOptions := &framework.GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextName: "context-one",
			}

			getContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   "context-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}
			ginkgo.By("Create SetContext API Command with Input and Output Parameters")
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Create GetContext API Command with Input and Output Parameters")
			getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Build TestCase to run SetContext API Command of Runtime Library v1.0.0 then GetContext API Command of Runtime Library v0.28.0 then GetContext API Command of Runtime Library v0.25.4")
			testCase := core.NewTestCase().Add(setContextCommand).Add(getContextCommand)

			ginkgo.By("Execute and validate the TestCase")
			framework.Execute(testCase)
		})

		ginkgo.It("Run SetContext API of Runtime Library v0.25.4 then GetContext API of Runtime Library v0.28.0", func() {
			ginkgo.By("Construct Input and Output parameters for SetContext API of Runtime Library v0.25.4")
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: "context-one",
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				SetCurrentContext: false,
			}

			var setContextOutputOptions *framework.SetContextOutputOptions

			ginkgo.By("Construct Input and Output parameters for GetContext API of Runtime Library v0.28.0")
			getContextInputOptionsForVersion0280 := &framework.GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextName: "context-one",
			}
			getContextOutputOptionsForVersion0280 := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Error: "context context-one not found",
			}

			ginkgo.By("Construct Input and Output parameters for GetContext API of Runtime Library v0.25.4")
			getContextInputOptionsForVersion0254 := &framework.GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextName: "context-one",
			}
			getContextOutputOptionsForVersion0254 := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},

				ContextOpts: &framework.ContextOpts{
					Name: "context-one",
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			ginkgo.By("Create SetContext API Command of Runtime Library v0.25.4 with Input and Output Parameters")
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Create GetContext API Command of Runtime Library v0.28.0 with Input and Output Parameters")
			getContextCommandForVersion0280, err := framework.NewGetContextCommand(getContextInputOptionsForVersion0280, getContextOutputOptionsForVersion0280)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Create GetContext API Command of Runtime Library v0.25.4 with Input and Output Parameters")
			getContextCommandForVersion0254, err := framework.NewGetContextCommand(getContextInputOptionsForVersion0254, getContextOutputOptionsForVersion0254)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Build TestCase to run SetContext API Command of Runtime v0.25.4 then GetContext API Command of Runtime v0.28.0 and then GetContext API Command of Runtime v0.25.4")
			testCase := core.NewTestCase().Add(setContextCommand).Add(getContextCommandForVersion0280).Add(getContextCommandForVersion0254) // re-named from NewTestCommands

			ginkgo.By("Executes the commands from the list and validates the expected output with actual output")
			framework.Execute(testCase)
		})
	})
})
