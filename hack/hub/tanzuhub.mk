# Paths and Directory information
ROOT_DIR := $(shell git rev-parse --show-toplevel)

## Tanzu Hub queries generatation

TANZU_HUB_PKG_BASE_DIR ?= $(ROOT_DIR)/pkg
TANZU_HUB_PKG_DIR=$(TANZU_HUB_PKG_BASE_DIR)/hub
TANZU_HUB_INIT_FILES := queries.graphql main.go genqlient.yaml
TANZU_HUB_INIT_PKG_URL ?= https://raw.githubusercontent.com/vmware-tanzu/tanzu-plugin-runtime/main/hack/hub/pkg
TANZU_HUB_SCHEMA_FILE_URL ?=

.PHONY: tanzu-hub-stub-init
tanzu-hub-stub-init: # Initialize a Tanzu Hub package with stub, schema.graphql and queries.graphql to generate graphQL client APIs
	mkdir -p $(TANZU_HUB_PKG_DIR)
	for filename in $(TANZU_HUB_INIT_FILES); do \
		[ -f $(TANZU_HUB_PKG_DIR)/$$filename ] && echo "Skipping $(TANZU_HUB_PKG_DIR)/$$filename (already exists)" || wget -O "$(TANZU_HUB_PKG_DIR)/$(filename)" $(TANZU_HUB_INIT_PKG_URL)/$$filename ; \
	done
	wget -O "$(TANZU_HUB_PKG_DIR)/schema.graphql" $(TANZU_HUB_SCHEMA_FILE_URL)
	go mod tidy

.PHONY: tanzu-hub-stub-generate
tanzu-hub-stub-generate: ## Generate golang stub for the Tanzu Hub queries specified under $(TANZU_HUB_PKG_BASE_DIR)/hub/queries.graphql
	go generate ./$(TANZU_HUB_PKG_DIR)/...
