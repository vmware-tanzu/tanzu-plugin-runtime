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

	ginkgo.Context("Test SetContext and SetCurrentContext with all versions of GetContext and GetCurrentContext", func() {

		ginkgo.It("SetContext, SetCurrentContext v1.0.0  then GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 then DeleteContext, RemoveCurrentContext v0.28.0 then  GetContext, GetCurrentContext v0.25.4, v0.28.0, v1.0.0 ", func() {
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

		ginkgo.It("Run SetContext v1.0.0 then SetCurrentContext v1.0.0 then GetContext v0.25.4, v0.28.0, v1.0.0 then GetCurrentContext v0.25.4, v0.28.0, v1.0.0", func() {
			// Input Parameters for SetContext v1.0.0
			setContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityOne,
					Target: framework.TargetK8s, // Target is required
					//	Type: framework.CtxTypeK8s, // Type is not supported
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
					Name: framework.CtxCompatibilityOne,
					// Name: framework.CtxCompatibilityTwo,
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

			// Input and Output Params for GetContext v1.0.0
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

			setContextWithGetContextOnAllVersionsTest := core.NewTestCase().Add(setContextCommand).Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

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

			// Input and Output Params for GetContext v1.0.0
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

			setContextWithGetContextOnAllVersionsTest := core.NewTestCase().Add(setContextCommand).Add(ctxV100GetContextCmd).Add(ctxV0280GetContextCmd).Add(ctxV0254GetContextCmd)

			framework.Execute(setContextWithGetContextOnAllVersionsTest)
		})
	})

	ginkgo.Context("Test SetContext, GetContext, DeleteContext, GetContext", func() {
		ginkgo.It("SetContext v0.28.0, GetContext v1.0.0, SetContext 2 v0.28.0, SetCurrentContext 2 v0.28.0, GetCurrentContext 2 v1.0.0, DeleteCurrentContext 2 v0.28.0, DeleteContext 2 v0.28.0, GetContext v0.28.0, GetContext 2 v0.28.0, GetCurrentContext v0.28.0, GetCurrentContext 2 v0.28.0", func() {
			// Input Parameters for Runtime SetContext v0.28.0 API
			contextOneSetContextInputOptions := &framework.SetContextInputOptions{
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

			contextTwoSetContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityTwo,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input and Output Parameters for Runtime GetContext API
			contextOneGetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			contextTwoGetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityTwo)

			contextOneGetContextOutputOptions := &framework.GetContextOutputOptions{
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

			contextTwoGetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityTwo,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			NoContextGetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Error: "context compatibility-one not found",
			}

			// Input and Output Options for SetCurrentContext
			contextOneSetCurrentContextInputOptions := framework.MakeSetCurrentContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			contextTwoSetCurrentContextInputOptions := framework.MakeSetCurrentContextInputOptions(core.Version0280, framework.CtxCompatibilityTwo)

			// Input and Output Options for GetCurrentContext
			kubernetesGetCurrentContextInputOptions := framework.MakeGetCurrentContextInputOptions(core.Version100, framework.TargetK8s)
			contextOneGetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
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

			contextTwoGetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityTwo,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			noContextGetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Error: "no current context set for target \"kubernetes\"",
			}

			// Input and Output Options for RemoveCurrentContext
			kubernetesRemoveCurrentContextInputOptions := framework.MakeRemoveCurrentContextInputOptions(core.Version0280, framework.TargetK8s)

			// Input and Output Options for Delete Context
			contextOneDeleteContextInputOptions := framework.MakeDeleteContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)

			// Create SetContext Commands
			contextOneSetContextCommand, err := framework.NewSetContextCommand(contextOneSetContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())
			contextTwoSetContextCommand, err := framework.NewSetContextCommand(contextTwoSetContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands
			contextOneGetContextCmd, err := framework.NewGetContextCommand(contextOneGetContextInputOptions, contextOneGetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())
			contextTwoGetContextCmd, err := framework.NewGetContextCommand(contextTwoGetContextInputOptions, contextTwoGetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			noContextOneGetContextCmd, err := framework.NewGetContextCommand(contextOneGetContextInputOptions, NoContextGetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext Commands
			contextOneSetCurrentContextCmd, err := framework.NewSetCurrentContextCommand(contextOneSetCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			contextTwoSetCurrentContextCmd, err := framework.NewSetCurrentContextCommand(contextTwoSetCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			kubernetesContextOneGetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(kubernetesGetCurrentContextInputOptions, contextOneGetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())
			KubernetesContextTwoGetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(kubernetesGetCurrentContextInputOptions, contextTwoGetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			kubernetesNoContextGetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(kubernetesGetCurrentContextInputOptions, noContextGetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentContext Commands
			KubernetesRemoveCurrentContextCmd, err := framework.NewRemoveCurrentContextCommand(kubernetesRemoveCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext Commands
			contextOneDeleteContextCmd, err := framework.NewDeleteContextCommand(contextOneDeleteContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(contextOneSetContextCommand).Add(contextTwoSetContextCommand).Add(contextOneGetContextCmd).Add(contextTwoGetContextCmd)

			testCase.Add(contextTwoSetCurrentContextCmd).Add(KubernetesContextTwoGetCurrentContextCmd).Add(contextOneSetCurrentContextCmd).Add(kubernetesContextOneGetCurrentContextCmd)
			testCase.Add(KubernetesRemoveCurrentContextCmd).Add(kubernetesNoContextGetCurrentContextCmd)
			testCase.Add(contextOneDeleteContextCmd).Add(noContextOneGetContextCmd).Add(contextTwoGetContextCmd)

			// Executes the commands from the list and validates the expected output with actual output
			framework.Execute(testCase)
		})

		ginkgo.It("SetContext v1.0.0 then GetContext v1.0.0, v0.28.0, v0.35.4 then DeleteContext v0.28.0  GetContext v1.0.0, v0.28.0, v0.35.4 ", func() {
			// Input Parameters for Runtime SetContext v0.28.0 API
			contextOneSetContextInputOptions := &framework.SetContextInputOptions{
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

			contextTwoSetContextInputOptions := &framework.SetContextInputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version0280,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityTwo,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Input and Output Parameters for Runtime GetContext API
			contextOneGetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityOne)
			contextTwoGetContextInputOptions := framework.MakeGetContextInputOptions(core.Version100, framework.CtxCompatibilityTwo)

			contextOneGetContextOutputOptions := &framework.GetContextOutputOptions{
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

			contextTwoGetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityTwo,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			NoContextGetContextOutputOptions := &framework.GetContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Error: "context compatibility-one not found",
			}

			// Input and Output Options for SetCurrentContext
			contextOneSetCurrentContextInputOptions := framework.MakeSetCurrentContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)
			contextTwoSetCurrentContextInputOptions := framework.MakeSetCurrentContextInputOptions(core.Version0280, framework.CtxCompatibilityTwo)

			// Input and Output Options for GetCurrentContext
			kubernetesGetCurrentContextInputOptions := framework.MakeGetCurrentContextInputOptions(core.Version100, framework.TargetK8s)
			contextOneGetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
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

			contextTwoGetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				ContextOpts: &framework.ContextOpts{
					Name:   framework.CtxCompatibilityTwo,
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			noContextGetCurrentContextOutputOptions := &framework.GetCurrentContextOutputOptions{
				RuntimeAPIVersion: &core.RuntimeAPIVersion{
					RuntimeVersion: core.Version100,
				},
				Error: "no current context set for target \"kubernetes\"",
			}

			// Input and Output Options for RemoveCurrentContext
			kubernetesRemoveCurrentContextInputOptions := framework.MakeRemoveCurrentContextInputOptions(core.Version0280, framework.TargetK8s)

			// Input and Output Options for Delete Context
			contextOneDeleteContextInputOptions := framework.MakeDeleteContextInputOptions(core.Version0280, framework.CtxCompatibilityOne)

			// Create SetContext Commands
			contextOneSetContextCommand, err := framework.NewSetContextCommand(contextOneSetContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())
			contextTwoSetContextCommand, err := framework.NewSetContextCommand(contextTwoSetContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetContext Commands
			contextOneGetContextCmd, err := framework.NewGetContextCommand(contextOneGetContextInputOptions, contextOneGetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())
			contextTwoGetContextCmd, err := framework.NewGetContextCommand(contextTwoGetContextInputOptions, contextTwoGetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			noContextOneGetContextCmd, err := framework.NewGetContextCommand(contextOneGetContextInputOptions, NoContextGetContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Create SetCurrentContext Commands
			contextOneSetCurrentContextCmd, err := framework.NewSetCurrentContextCommand(contextOneSetCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			contextTwoSetCurrentContextCmd, err := framework.NewSetCurrentContextCommand(contextTwoSetCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create GetCurrentContext Commands
			kubernetesContextOneGetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(kubernetesGetCurrentContextInputOptions, contextOneGetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())
			KubernetesContextTwoGetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(kubernetesGetCurrentContextInputOptions, contextTwoGetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			kubernetesNoContextGetCurrentContextCmd, err := framework.NewGetCurrentContextCommand(kubernetesGetCurrentContextInputOptions, noContextGetCurrentContextOutputOptions)
			gomega.Expect(err).To(gomega.BeNil())

			// Create RemoveCurrentContext Commands
			KubernetesRemoveCurrentContextCmd, err := framework.NewRemoveCurrentContextCommand(kubernetesRemoveCurrentContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Create DeleteContext Commands
			contextOneDeleteContextCmd, err := framework.NewDeleteContextCommand(contextOneDeleteContextInputOptions, nil)
			gomega.Expect(err).To(gomega.BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(contextOneSetContextCommand).Add(contextTwoSetContextCommand).Add(contextOneGetContextCmd).Add(contextTwoGetContextCmd)

			testCase.Add(contextTwoSetCurrentContextCmd).Add(KubernetesContextTwoGetCurrentContextCmd).Add(contextOneSetCurrentContextCmd).Add(kubernetesContextOneGetCurrentContextCmd)
			testCase.Add(KubernetesRemoveCurrentContextCmd).Add(kubernetesNoContextGetCurrentContextCmd)
			testCase.Add(contextOneDeleteContextCmd).Add(noContextOneGetContextCmd).Add(contextTwoGetContextCmd)

			// Executes the commands from the list and validates the expected output with actual output
			framework.Execute(testCase)
		})

	})
})
