// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

var _ = Describe("Context API", func() {

	BeforeEach(func() {
		core.SetupTempCfgFiles()
	})

	Context("Context API Set and Get", func() {

		It("Runtime V100 SetContext API", func() {
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
			Expect(err).To(BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(setContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
			framework.Execute(testCase)
		})

		It("Runtime V0280 GetContext API", func() {

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
			Expect(err).To(BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
			framework.Execute(testCase)
		})

		It("Runtime V100 SetContext API and Runtime V0280 GetContext API", func() {
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
				ContextOpts: &framework.ContextOpts{
					Name:   "context-one",
					Target: framework.TargetK8s,
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Create SetContext Command
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, &setContextOutputOptions)
			Expect(err).To(BeNil())
			Expect(setContextCommand).NotTo(BeNil())

			// Create GetContext Command
			getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
			Expect(err).To(BeNil())
			Expect(getContextCommand).NotTo(BeNil())

			// Construct series of commands to execute

			testCase := core.NewTestCase().Add(setContextCommand).Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
			framework.Execute(testCase)
		})

		It("SetContext v1.0.0 then GetContext v0.28.0", func() {

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
			Expect(err).To(BeNil())

			// Create GetContextAPIName Command
			getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
			Expect(err).To(BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(setContextCommand).Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output
			framework.Execute(testCase)
		})

		It("SetContextAPIName v0.25.4 then GetContextAPIName v0.28.0", func() {

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
				IsCurrentContext: false,
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
				ContextOpts: &framework.ContextOpts{
					Name:   "context-one",
					Target: "",
					GlobalOpts: &framework.GlobalServerOpts{
						Endpoint: "test-endpoint",
					},
				},
			}

			// Create SetContextAPIName Command
			setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
			Expect(err).To(BeNil())

			// Create GetContextAPIName Command
			getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
			Expect(err).To(BeNil())

			// Construct series of commands to execute
			testCase := core.NewTestCase().Add(setContextCommand).Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output
			framework.Execute(testCase)
		})

		//// Sample Test Failure test case
		//It("SetContext v0.28.0 SetContext v0.28.0(Unsetting ClusterOpts.Endpoint) GetContext v0.28.0", func() {
		//	// Input Parameters for Runtime SetContext API V0.28.0
		//	setContextInputOptions := &framework.SetContextInputOptions{
		//		RuntimeAPIVersion: &core.RuntimeAPIVersion{
		//			RuntimeVersion: core.Version0280,
		//		},
		//		ContextOpts: &framework.ContextOpts{
		//			Name:   "context-one",
		//			Target: framework.TargetTMC,
		//			ClusterOpts: &framework.ClusterServerOpts{
		//				Endpoint: "test-endpoint",
		//				Path:     "test-path",
		//			},
		//		},
		//	}
		//
		//	// Input Parameters for Runtime SetContext API V0.28.0 With ClusterOpts.Endpoint unset
		//	setContextInputOptionsWithEndpointUnset := &framework.SetContextInputOptions{
		//		RuntimeAPIVersion: &core.RuntimeAPIVersion{
		//			RuntimeVersion: core.Version0280,
		//		},
		//		ContextOpts: &framework.ContextOpts{
		//			Name:   "context-one",
		//			Target: framework.TargetTMC,
		//			ClusterOpts: &framework.ClusterServerOpts{
		//				Endpoint: "",
		//				Path:     "test-path",
		//			},
		//		},
		//	}
		//	// Output Parameters for Runtime SetContext API
		//	var setContextOutputOptions framework.SetContextOutputOptions
		//
		//	// Create SetContext Command 1
		//	setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, &setContextOutputOptions)
		//	Expect(err).To(BeNil())
		//	Expect(setContextCommand).NotTo(BeNil())
		//
		//	// Create SetContext Command WithEndpointUnset
		//	setContextCommandWithEndpointUnset, err := framework.NewSetContextCommand(setContextInputOptionsWithEndpointUnset, &setContextOutputOptions)
		//	Expect(err.Error()).To(Equal("invalid set context input options for the specified runtime version v0.28.0"))
		//	Expect(setContextCommandWithEndpointUnset).NotTo(BeNil())
		//
		//	// Input Parameters for Runtime GetContext API
		//	getContextInputOptions := &framework.GetContextInputOptions{
		//		RuntimeAPIVersion: &core.RuntimeAPIVersion{
		//			RuntimeVersion: core.Version0280,
		//		},
		//		ContextName: "context-one",
		//	}
		//
		//	// Output Parameters for Runtime GetContext API
		//	getContextOutputOptions := &framework.GetContextOutputOptions{
		//		RuntimeAPIVersion: &core.RuntimeAPIVersion{
		//			RuntimeVersion: core.Version0280,
		//		},
		//		ContextOpts: &framework.ContextOpts{
		//			Name:   "context-one",
		//			Target: framework.TargetTMC,
		//			ClusterOpts: &framework.ClusterServerOpts{
		//				Endpoint: "test-endpoint",
		//				Path:     "test-path",
		//			},
		//		},
		//	}
		//
		//	getContextOutputOptionsWithEndpointNotExpected := &framework.GetContextOutputOptions{
		//		RuntimeAPIVersion: &core.RuntimeAPIVersion{
		//			RuntimeVersion: core.Version0280,
		//		},
		//		ContextOpts: &framework.ContextOpts{
		//			Name:   "context-one",
		//			Target: framework.TargetTMC,
		//			ClusterOpts: &framework.ClusterServerOpts{
		//				Path: "test-path",
		//			},
		//		},
		//	}
		//
		//	// Create GetContextAPIName Command
		//	getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
		//	Expect(err).To(BeNil())
		//	Expect(getContextCommand).NotTo(BeNil())
		//
		//	// Create GetContextAPIName Command WithEndpointNotExpected
		//	getContextCommandWithEndpointNotExpected, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptionsWithEndpointNotExpected)
		//	Expect(err).To(BeNil())
		//	Expect(getContextCommandWithEndpointNotExpected).NotTo(BeNil())
		//
		//	// Construct series of commands to execute
		//	testCase1 := core.NewTestCase().Add(setContextCommand).Add(setContextCommandWithEndpointUnset).Add(getContextCommand)
		//	// Executes the commands from the list and validates the expected output with actual output
		//	framework.Execute(testCase1) // This execution fails since ClusterOpts.Endpoint is unset in setContextCommandWithEndpointUnset but expected in getContextCommand
		//
		//	testCase2 := core.NewTestCase().Add(setContextCommand).Add(setContextCommandWithEndpointUnset).Add(getContextCommandWithEndpointNotExpected)
		//	// Executes the commands from the list and validates the expected output with actual output
		//	framework.Execute(testCase2) // This execution succeeds since ClusterOpts.Endpoint is unset in setContextCommandWithEndpointUnset and not expected in getContextCommandWithEndpointNotExpected
		//
		//})

	})
})
