// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	executil "github.com/vmware-tanzu/tanzu-plugin-runtime/test/exec"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plugin test suite")
}

var _ = Describe("Plugin tests", func() {
	var (
		command    *executil.Command
		pluginName string
	)

	Describe("Test plugin contract of minimal viable plugin", func() {
		BeforeEach(func() {
			cwd, err := os.Getwd()
			Expect(err).To(BeNil())
			pluginName = filepath.Join(cwd, "plugins", "bin", "helloworld")
		})
		Context("When minimal viable plugin is created", func() {
			It("Should have 'info' command ", func() {
				command = executil.NewCommand(
					executil.WithCommand(pluginName),
					executil.WithArgs("info"),
				)
				stdout, stderr, err := command.Run(context.Background())
				Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("failed to run plugin, stdout - %v \n stderr -  %v\n", string(stdout), string(stderr)))
				Expect(stderr).To(BeEmpty())
				Expect(string(stdout)).To(ContainSubstring(`"name":"helloworld-test"`))
				Expect(string(stdout)).To(ContainSubstring(`"version":"v0.0.1"`))
				Expect(string(stdout)).To(ContainSubstring(`"pluginRuntimeVersion"`))
			})
			It("Should have 'version' command ", func() {
				command = executil.NewCommand(
					executil.WithCommand(pluginName),
					executil.WithArgs("version"),
				)
				stdout, stderr, err := command.Run(context.Background())
				Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("failed to run plugin, stdout - %v \n stderr -  %v\n", string(stdout), string(stderr)))
				Expect(stderr).To(BeEmpty())
				Expect(string(stdout)).To(ContainSubstring(`v0.0.1`))
			})
			It("Should have 'describe' command ", func() {
				command = executil.NewCommand(
					executil.WithCommand(pluginName),
					executil.WithArgs("describe"),
				)
				stdout, stderr, err := command.Run(context.Background())
				Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("failed to run plugin, stdout - %v \n stderr -  %v\n", string(stdout), string(stderr)))
				Expect(stderr).To(BeEmpty())
				Expect(string(stdout)).To(ContainSubstring(`Hello world test plugin`))
			})
			It("Should have 'post-install' command and should run successfully", func() {
				command = executil.NewCommand(
					executil.WithCommand(pluginName),
					executil.WithArgs("post-install"),
				)
				stdout, stderr, err := command.Run(context.Background())
				Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("failed to run plugin, stdout - %v \n stderr -  %v\n", string(stdout), string(stderr)))
				Expect(stderr).To(BeEmpty())
				Expect(string(stdout)).To(ContainSubstring(`Hello world - post install method`))
			})
			It("Should have run additionally added command successfully", func() {
				command = executil.NewCommand(
					executil.WithCommand(pluginName),
					executil.WithArgs("print"),
				)
				stdout, stderr, err := command.Run(context.Background())
				Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("failed to run plugin, stdout - %v \n stderr -  %v\n", string(stdout), string(stderr)))
				Expect(stderr).To(BeEmpty())
				Expect(string(stdout)).To(ContainSubstring(`Hello world from test plugin`))
			})

		})
	})
})
