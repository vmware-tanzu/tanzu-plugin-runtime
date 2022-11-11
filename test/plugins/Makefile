# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOHOSTOS ?= $(shell go env GOHOSTOS)
GOHOSTARCH ?= $(shell go env GOHOSTARCH)

ROOT_DIR := $(shell git rev-parse --show-toplevel)
BIN_DIR := bin

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

GO := go

NUL = /dev/null
ifeq ($(GOHOSTOS),windows)
	NUL = NUL
endif

BUILD_SHA ?= $$(git describe --match=$(git rev-parse --short HEAD) --always --dirty)
BUILD_DATE ?= $$(date -u +"%Y-%m-%d")
BUILD_VERSION ?= $(shell git describe --tags --abbrev=0 2>$(NUL))

ifeq ($(strip $(BUILD_VERSION)),)
BUILD_VERSION = dev
endif


ifdef DEBUG
LD_FLAGS = -s
GC_FLAGS = all=-N -l
else
LD_FLAGS = -s -w
GC_FLAGS =
endif

# Remove old package path for buildinfo vars when it's no longer in use by plugin binaries.
LD_FLAGS += -X 'github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo.Date=$(BUILD_DATE)'
LD_FLAGS += -X 'github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo.SHA=$(BUILD_SHA)'
LD_FLAGS += -X 'github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo.Version=$(BUILD_VERSION)'


HELLOWORLD_PLUGIN      := $(BIN_DIR)/helloworld

.PHONY: all
all: build-helloworld-plugin

build-helloworld-plugin: $(HELLOWORLD_PLUGIN)
$(HELLOWORLD_PLUGIN):
	mkdir -p bin
	cd helloworld && GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -tags "${BUILD_TAGS}" --ldflags "$(LD_FLAGS)"  -o ../$(HELLOWORLD_PLUGIN) .

## --------------------------------------
## Cleanup
## --------------------------------------

.PHONY: clean
clean: ## Remove all generated binaries
	rm -rf bin