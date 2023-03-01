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

	ginkgo.Context("Test SetContext and SetCurrentContext on specific Runtime version and all Runtime versions of GetContext and GetCurrentContext", func() {

		ginkgo.It("SetContext, SetCurrentContext v1.0.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0", func() {
			ginkgo.By("Construct Input parameters for SetContext API of Runtime Library v1.0.0")

			setContextInputOptions, err := framework.MakeSetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Construct Input and Output parameters for GetContext API of Runtime Library v1.0.0")
			ctxV100GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			ctxV100GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   "compatibility-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ginkgo.By("Construct Input and Output parameters for GetContext API of Runtime Library v0.28.0")
			ctxV0280GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			ctxV0280GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   "compatibility-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ginkgo.By("Construct Input and Output parameters for GetContext API of Runtime Library v0.11.6")
			ctxV0254GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0254, framework.CtxCompatibilityOne)
			ctxV0254GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: "compatibility-one",
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			ginkgo.By("Create SetContext API Command with Input and Output Parameters for Runtime Library v1.0.0")
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Create GetContext API Command with Input and Output Parameters for Runtime Library v1.0.0")
			ctxV100GetContextCmd, err := framework.NewGetContextCommand(ctxV100GetContextInputOptions, ctxV100GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Create GetContext API Command with Input and Output Parameters for Runtime Library v0.28.0")
			ctxV0280GetContextCmd, err := framework.NewGetContextCommand(ctxV0280GetContextInputOptions, ctxV0280GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Create GetContext API Command with Input and Output Parameters for Runtime Library v0.25.4")
			ctxV0254GetContextCmd, err := framework.NewGetContextCommand(ctxV0254GetContextInputOptions, ctxV0254GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ginkgo.By("Build TestCase to run SetContext API Command of Runtime Library v1.0.0 then GetContext API Command of Runtime Library v1.0.0, v0.28.0 and v0.25.4")
			setContextWithGetContextOnAllVersionsTest := core.NewTestCase().Add(setContextCommand).Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

			ginkgo.By("Execute and validate the TestCase")
			framework.Execute(setContextWithGetContextOnAllVersionsTest)
		})

		ginkgo.It("SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0", func() {
			// Input Parameters for SetContext v0.28.0
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   "compatibility-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input and Output Params for GetContext v1.0.0, v0.28.0, v0.25.4
			ctxV100GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			ctxV0280GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			ctxV0254GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0254, framework.CtxCompatibilityOne)

			ctxV100GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   "compatibility-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0280GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   "compatibility-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0254GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: "compatibility-one",
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Create SetContext Command with input and output options
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			ctxV100GetContextCmd, err := framework.NewGetContextCommand(ctxV100GetContextInputOptions, ctxV100GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetContextCmd, err := framework.NewGetContextCommand(ctxV0280GetContextInputOptions, ctxV0280GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetContextCmd, err := framework.NewGetContextCommand(ctxV0254GetContextInputOptions, ctxV0254GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Build the test case with series of commands
			setContextWithGetContextOnAllVersionsTest := core.NewTestCase().Add(setContextCommand).Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

			//Run the test case
			framework.Execute(setContextWithGetContextOnAllVersionsTest)
		})

		ginkgo.It("SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0", func() {
			// Input Parameters for SetContext v0.25.4
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: "compatibility-one",
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input and Output Parameters for GetContext v1.0.0, v0.28.0, v0.25.4
			ctxV100GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			ctxV0280GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			ctxV0254GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0254, framework.CtxCompatibilityOne)

			ctxV100GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Error: "context compatibility-one not found",
			}

			ctxV0280GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Error: "context compatibility-one not found",
			}

			ctxV0254GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: "compatibility-one",
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Create SetContext Command with input and output options
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			ctxV100GetContextCmd, err := framework.NewGetContextCommand(ctxV100GetContextInputOptions, ctxV100GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetContextCmd, err := framework.NewGetContextCommand(ctxV0280GetContextInputOptions, ctxV0280GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetContextCmd, err := framework.NewGetContextCommand(ctxV0254GetContextInputOptions, ctxV0254GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Build the TestCase with series of commands
			setContextWithGetContextOnAllVersionsTest := core.NewTestCase().Add(setContextCommand).Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

			// Run the test case
			framework.Execute(setContextWithGetContextOnAllVersionsTest)
		})

	})

	ginkgo.Context("Test SetContext, SetCurrentContext, GetContext, GetCurrentContext, DeleteContext, RemoveCurrentContext on respective Runtime API versions", func() {

		ginkgo.It("SetContext, SetCurrentContext v1.0.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 then DeleteContext, RemoveCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 ", func() {
			// Input Parameters for SetContext v1.0.0
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input Parameters for SetCurrentContext v1.0.0
			setCurrentContextInputOptions := framework.MakeSetCurrentContextInputOptions(core.Version100, framework.CtxCompatibilityOne)

			// Input Parameters for GetCurrentContext
			V100GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Target: framework.TargetK8s,
			}

			V0280GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Target: framework.TargetK8s,
			}

			V0254GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextType: framework.CtxTypeK8s,
			}

			// Output Params for GetCurrentContext
			ctxV100GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0280GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0254GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: framework.CtxCompatibilityOne,
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input and Output Params for GetContext
			ctxV100GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			ctxV0280GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			ctxV0254GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0254, framework.CtxCompatibilityOne)

			ctxV100GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0280GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0254GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: framework.CtxCompatibilityOne,
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Create SetContext Command with input and output options
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext Command with input and output options
			setCurrentContextCommand, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			ctxV100GetContextCmd, err := framework.NewGetContextCommand(ctxV100GetContextInputOptions, ctxV100GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetContextCmd, err := framework.NewGetContextCommand(ctxV0280GetContextInputOptions, ctxV0280GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetContextCmd, err := framework.NewGetContextCommand(ctxV0254GetContextInputOptions, ctxV0254GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV100GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V100GetCurrentContextInputOptions, ctxV100GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V0280GetCurrentContextInputOptions, ctxV0280GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V0254GetCurrentContextInputOptions, ctxV0254GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCommand).Add(setCurrentContextCommand)

			// Add GetContext v1.0.0, v0.28.0, v0.25.4 Commands
			testCase.Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

			// Add GetCurrentContext v1.0.0, v0.28.0, v0.25.4 Commands
			testCase.Add(ctxV100GetCurrentContextCmd).Add(ctxV0280GetCurrentContextCmd).Add(ctxV0254GetCurrentContextCmd)

			// Run all the commands
			framework.Execute(testCase)
		})

		ginkgo.It("SetContext, SetCurrentContext v0.28.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 then DeleteContext, RemoveCurrentContext v0.25.4 then  GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 ", func() {
			// Input Parameters for SetContext v0.28.0
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input Parameters for SetCurrentContext v0.28.0
			setCurrentContextInputOptions := framework.MakeSetCurrentContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)

			// Input Parameters for GetCurrentContext
			V100GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Target: framework.TargetK8s,
			}

			V0280GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Target: framework.TargetK8s,
			}

			V0254GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextType: framework.CtxTypeK8s,
			}

			// Output Params for GetCurrentContext
			ctxV100GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0280GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0254GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: framework.CtxCompatibilityOne,
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input and Output Params for GetContext
			ctxV100GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			ctxV0280GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			ctxV0254GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0254, framework.CtxCompatibilityOne)

			ctxV100GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0280GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
				ValidationStrategy: core.ValidationStrategyStrict,
			}

			ctxV0254GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: framework.CtxCompatibilityOne,
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Create SetContext Command with input and output options
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext Command with input and output options
			setCurrentContextCommand, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			ctxV100GetContextCmd, err := framework.NewGetContextCommand(ctxV100GetContextInputOptions, ctxV100GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetContextCmd, err := framework.NewGetContextCommand(ctxV0280GetContextInputOptions, ctxV0280GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetContextCmd, err := framework.NewGetContextCommand(ctxV0254GetContextInputOptions, ctxV0254GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV100GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V100GetCurrentContextInputOptions, ctxV100GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V0280GetCurrentContextInputOptions, ctxV0280GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V0254GetCurrentContextInputOptions, ctxV0254GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Add SetContext and SetCurrentContext Commands on v0.28.0
			testCase := core.NewTestCase().Add(setContextCommand).Add(setCurrentContextCommand)

			// Add GetContext v1.0.0, v0.28.0, v0.25.4 Commands
			testCase.Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

			// Add GetCurrentContext v1.0.0, v0.28.0, v0.25.4 Commands
			testCase.Add(ctxV100GetCurrentContextCmd).Add(ctxV0280GetCurrentContextCmd).Add(ctxV0254GetCurrentContextCmd)

			// Run all the commands
			framework.Execute(testCase)
		})

		ginkgo.It("SetContext, SetCurrentContext v0.25.4 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 then DeleteContext, RemoveCurrentContext v1.0.0 then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 ", func() {
			// Input Parameters for SetContext v0.25.4
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: framework.CtxCompatibilityOne,
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input Parameters for SetCurrentContext v0.25.4
			setCurrentContextInputOptions := framework.MakeSetCurrentContextInputOptions(core.Version0254, framework.CtxCompatibilityOne)

			// Input Parameters for GetCurrentContext
			V100GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Target: framework.TargetK8s,
			}

			V0280GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Target: framework.TargetK8s,
			}

			V0254GetCurrentContextInputOptions := &framework.GetCurrentContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextType: framework.CtxTypeK8s,
			}

			// Output Params for GetCurrentContext
			ctxV100GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Error: "no current context set for target \"kubernetes\"",
			}

			ctxV0280GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Error: "no current context set for target \"kubernetes\"",
			}

			ctxV0254GetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: framework.CtxCompatibilityOne,
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input and Output Params for GetContext
			ctxV100GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			ctxV0280GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			ctxV0254GetContextInputOptions := framework.MakeGetContextInputOptions(core.Version0254, framework.CtxCompatibilityOne)

			ctxV100GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Error: "context compatibility-one not found",
			}

			ctxV0280GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				Error: "context compatibility-one not found",
			}

			ctxV0254GetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0254,
				},
				ContextOpts: &framework.ContextOpts{
					Name: framework.CtxCompatibilityOne,
					Type: framework.CtxTypeK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Create SetContext Command with input and output options
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext Command with input and output options
			setCurrentContextCommand, err := framework.NewSetCurrentContextCommand(setCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands with input and output options
			ctxV100GetContextCmd, err := framework.NewGetContextCommand(ctxV100GetContextInputOptions, ctxV100GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetContextCmd, err := framework.NewGetContextCommand(ctxV0280GetContextInputOptions, ctxV0280GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetContextCmd, err := framework.NewGetContextCommand(ctxV0254GetContextInputOptions, ctxV0254GetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV100GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V100GetCurrentContextInputOptions, ctxV100GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0280GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V0280GetCurrentContextInputOptions, ctxV0280GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			ctxV0254GetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(V0254GetCurrentContextInputOptions, ctxV0254GetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Add SetContext and SetCurrentContext Commands
			testCase := core.NewTestCase().Add(setContextCommand).Add(setCurrentContextCommand)

			// Add GetContext v1.0.0, v0.28.0, v0.25.4 Commands
			testCase.Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

			// Add GetCurrentContext v1.0.0, v0.28.0, v0.25.4 Commands
			testCase.Add(ctxV100GetCurrentContextCmd).Add(ctxV0280GetCurrentContextCmd).Add(ctxV0254GetCurrentContextCmd)

			// Run all the commands
			framework.Execute(testCase)
		})

	})

	ginkgo.Context("Test SetContext and GetContext APIs on specific Runtime Library version v1.0.0, v0.28.0, v0.25.4", func() {

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
