// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework"
)

var _ = Describe("Context API", func() {

	Context("Context API Set and Get", func() {

		It("SetContextAPIName v1.0.0 then GetContextAPIName v0.28.0", func() {

			// Input Parameters for Runtime SetContextAPIName API
			setContextInputOptions := framework.SetContextInputOptions{
				RuntimeAPIVersion: framework.RuntimeAPIVersion{
					RuntimeVersion: framework.Version100,
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
			var setContextOutputOptions framework.SetContextOutputOptions

			// Input Parameters for Runtime GetContextAPIName API
			getContextInputOptions := framework.GetContextInputOptions{
				RuntimeAPIVersion: framework.RuntimeAPIVersion{
					RuntimeVersion: framework.Version0280,
				},
				ContextName: "context-one",
			}

			// Output Parameters for Runtime GetContextAPIName API
			getContextOutputOptions := framework.GetContextOutputOptions{
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
			testCase := framework.NewTestCase().Add(setContextCommand).Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
			errs := testCase.Execute()
			Expect(errs).To(BeNil())
		})

		It("SetContextAPIName v0.25.4 then GetContextAPIName v0.28.0", func() {

			// Input Parameters for Runtime SetContextAPIName API
			setContextInputOptions := framework.SetContextInputOptions{
				RuntimeAPIVersion: framework.RuntimeAPIVersion{
					RuntimeVersion: framework.Version0254,
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
			var setContextOutputOptions framework.SetContextOutputOptions

			// Input Parameters for Runtime GetContextAPIName API
			getContextInputOptions := framework.GetContextInputOptions{
				RuntimeAPIVersion: framework.RuntimeAPIVersion{
					RuntimeVersion: framework.Version0280,
				},
				ContextName: "context-one",
			}

			// Output Parameters for Runtime GetContextAPIName API
			getContextOutputOptions := framework.GetContextOutputOptions{
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
			testCase := framework.NewTestCase().Add(setContextCommand).Add(getContextCommand) // re-named from NewTestCommands

			// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
			errs := testCase.Execute()
			Expect(errs).To(BeNil())
		})
	})
})
