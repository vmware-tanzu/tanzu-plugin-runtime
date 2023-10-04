# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL := /usr/bin/env bash

ROOT_DIR := $(shell git rev-parse --show-toplevel)

# Golang specific variables
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOHOSTOS ?= $(shell go env GOHOSTOS)
GOHOSTARCH ?= $(shell go env GOHOSTARCH)
GOTEST_VERBOSE ?= -v

# By Default verbose is turned off for compatibility tests; set the value to -v to turn on
COMPATIBILITY_TEST_VERBOSE ?= # -v

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif
GO := go
GINKGO := ginkgo

# Directories
TOOLS_DIR := $(abspath $(ROOT_DIR)/hack/tools)
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
GO_MODULES=$(shell find . -path "*/go.mod" | xargs -I _ dirname _)

# Add tooling binaries here and in hack/tools/Makefile
GOIMPORTS          := $(TOOLS_BIN_DIR)/goimports
GOLANGCI_LINT      := $(TOOLS_BIN_DIR)/golangci-lint
VALE               := $(TOOLS_BIN_DIR)/vale
MISSPELL           := $(TOOLS_BIN_DIR)/misspell
GINKGO             := $(TOOLS_BIN_DIR)/ginkgo
TOOLING_BINARIES   := $(GOIMPORTS) $(GOLANGCI_LINT) $(VALE) $(MISSPELL) $(GINKGO)

## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help (default)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-28s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m\033[32m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## All
## --------------------------------------

.PHONY: all
all: modules test lint ## ## Run all major targets (lint, test, modules)

## --------------------------------------
## Testing
## --------------------------------------

.PHONY: test
test: fmt ## Run Tests
	make -C test/plugins all
	${GO} test ./... -timeout 60m -race -coverprofile coverage.txt

.PHONY: fmt
fmt: $(GOIMPORTS) ## Run goimports
	$(GOIMPORTS) -w -local github.com/vmware-tanzu ./

lint: tools go-lint doc-lint misspell yamllint ## Run linting and misspell checks
	# Check licenses in shell scripts and Makefiles
	hack/check/check-license.sh

misspell: $(MISSPELL)
	hack/check/misspell.sh

yamllint:
	hack/check/check-yaml.sh

go-lint: $(GOLANGCI_LINT)  ## Run linting of go source
	$(GOLANGCI_LINT) --version
	@for i in $(GO_MODULES); do \
		echo "-- Linting $$i --"; \
		pushd $${i}; \
		$(GOLANGCI_LINT) run -v --timeout=10m || exit 1; \
		popd; \
	done


	# Prevent use of deprecated ioutils module
	@CHECK=$$(grep -r --include="*.go"  --exclude="zz_generated*" ioutil .); \
	if [ -n "$${CHECK}" ]; then \
		echo "ioutil is deprecated, use io or os replacements"; \
		echo "https://go.dev/doc/go1.16#ioutil"; \
		echo "$${CHECK}"; \
		exit 1; \
	fi

doc-lint: $(VALE) ## Run linting checks for docs
	$(VALE) --config=.vale/config.ini --glob='*.md' ./
	# mdlint rules with possible errors and fixes can be found here:
	# https://github.com/DavidAnson/markdownlint/blob/main/doc/Rules.md
	# Additional configuration can be found in the .markdownlintrc file.
	hack/check/check-mdlint.sh

.PHONY: modules
modules: ## Runs go mod to ensure modules are up to date.
	@for i in $(GO_MODULES); do \
		echo "-- Tidying $$i --"; \
		pushd $${i}; \
		$(GO) mod tidy || exit 1; \
		popd; \
	done

## --------------------------------------
## Tooling Binaries
## --------------------------------------

tools: $(TOOLING_BINARIES) ## Build tooling binaries
.PHONY: $(TOOLING_BINARIES)
$(TOOLING_BINARIES):
	make -C $(TOOLS_DIR) $(@F)

.PHONY: clean
clean: ## Remove all generated binaries
	make -C test/plugins clean

GO_MODULES=$(shell find . -path "*/go.mod" | xargs -I _ dirname _)

## --------------------------------------
## Compatibility Testing
## --------------------------------------

.PHONY: compatibility-tests
compatibility-tests: tools build-compatibility-test-plugins run-compatibility-tests ## Build and Run Compatibility tests

.PHONY: build-compatibility-test-plugins
build-compatibility-test-plugins: ## Builds all runtime compatibility test plugins
	cd ./test/compatibility/testplugins && mkdir -p bin
	cd ./test/compatibility/testplugins/runtime-test-plugin-v0_11_6 && ${GO} mod tidy && GOOS=$(OS) GOARCH=$(ARCH) ${GO} build -o ../bin
	cd ./test/compatibility/testplugins/runtime-test-plugin-v0_25_4 && ${GO} mod tidy && GOOS=$(OS) GOARCH=$(ARCH) ${GO} build -o ../bin
	cd ./test/compatibility/testplugins/runtime-test-plugin-v0_28_0 && ${GO} mod tidy && GOOS=$(OS) GOARCH=$(ARCH) ${GO} build -o ../bin
	cd ./test/compatibility/testplugins/runtime-test-plugin-v0_90 && ${GO} mod tidy && GOOS=$(OS) GOARCH=$(ARCH) ${GO} build -o ../bin
	cd ./test/compatibility/testplugins/runtime-test-plugin-v1_0_2 && ${GO} mod tidy && GOOS=$(OS) GOARCH=$(ARCH) ${GO} build -o ../bin
	cd ./test/compatibility/testplugins/runtime-test-plugin-latest && ${GO} mod tidy && GOOS=$(OS) GOARCH=$(ARCH) ${GO} build -o ../bin

# The below command runs the compatibility tests using ginkgo and filter the logs as per regex and exit code with '0' representing success and non-zero values indicating test failures
.PHONY: run-compatibility-tests
run-compatibility-tests: ## Run Compatibility tests
	${GINKGO} --keep-going --output-dir testresults --json-report=compatibility-tests.json --race ${GOTEST_VERBOSE} -r test/compatibility/framework/compatibilitytests  --randomize-all --trace > /tmp/out && { cat /tmp/out | grep -Ev 'STEP:|seconds|.go:'; rm /tmp/out; } || { exit_code=$$?; cat /tmp/out | grep -Ev 'STEP:|seconds|.go:'; rm /tmp/out; exit $$exit_code; }
