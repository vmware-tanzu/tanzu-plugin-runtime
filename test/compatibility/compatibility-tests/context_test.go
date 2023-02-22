// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package compatibility_tests_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

var _ = ginkgo.Describe("Context API", func() {

	ginkgo.BeforeEach(func() {
		core.SetupTempCfgFiles()
	})

	ginkgo.Context("Runtime Context API Set and Get", func() {

		ginkgo.It("Run Runtime v1.0.0 SetContext API", func() {
			// Input Parameters for Runtime SetContext API
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

			// Output Parameters for Runtime SetContext API
			var setContextOutputOptions framework.SetContextOutputOptions

			// Create SetContext Command
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, &setContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(setContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
			framework.Execute(testCase)
		})

		ginkgo.It("Run Runtime v0.28.0 GetContext API", func() {
			// Input Parameters for Runtime GetContext API
			getContextInputOptions := &framework.GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextName: "context-one",
			}

			// Output Parameters for Runtime GetContext API
			getContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Error: "context context-one not found",
			}

			// Create GetContext Command
			getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
			framework.Execute(testCase)
		})

		ginkgo.It("Run Runtime v1.0.0 SetContext API then Runtime v0.28.0 GetContext", func() {
			// Input Parameters for Runtime SetContextAPIName API
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

			// Output Parameters for Runtime SetContextAPIName API
			var setContextOutputOptions *framework.SetContextOutputOptions

			// Input Parameters for Runtime GetContextAPIName API
			getContextInputOptions := &framework.GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextName: "context-one",
			}

			// Output Parameters for Runtime GetContextAPIName API
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

			// Create SetContextAPIName Command
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContextAPIName Command
			getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(setContextCommand).Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output
			framework.Execute(testCase)
		})

		ginkgo.It("Run Runtime v0.25.4 SetContext API then Runtime v0.28.0 GetContext API", func() {
			// Input Parameters for Runtime SetContextAPIName API
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

			// Output Parameters for Runtime SetContextAPIName API
			var setContextOutputOptions *framework.SetContextOutputOptions

			// Input Parameters for Runtime GetContextAPIName API
			getContextInputOptionsForVersion0280 := &framework.GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextName: "context-one",
			}

			getContextInputOptionsForVersion0254 := &framework.GetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextName: "context-one",
			}

			// Output Parameters for Runtime GetContextAPIName API
			getContextOutputOptionsForVersion0280 := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Error: "context context-one not found",
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

			// Create SetContext API Command
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext API Command
			getContextCommandForVersion0280, err := framework.NewGetContextCommand(getContextInputOptionsForVersion0280, getContextOutputOptionsForVersion0280)
			gomega.Expect(err).To(gomega.BeNil())
			getContextCommandForVersion0254, err := framework.NewGetContextCommand(getContextInputOptionsForVersion0254, getContextOutputOptionsForVersion0254)
			gomega.Expect(err).To(gomega.BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(setContextCommand).Add(getContextCommandForVersion0280).Add(getContextCommandForVersion0254) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output
			framework.Execute(testCase)
		})

	})

})
