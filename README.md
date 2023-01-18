# Tanzu Plugin Runtime

:warning: NOTE: This repository is still under early development. We hope to be make releases available for evaluation in early February 2023.
Please watch this page for further updates.

## Overview

The Tanzu CLI is based on a plugin architecture. This architecture enables teams to build, own, and release their own piece of functionality as well as enable external partners to integrate with the system. The Tanzu Plugin Runtime provides functionality and helper methods to develop Tanzu CLI plugins.

Developers can use the `Builder` admin plugin to bootstrap a new plugin which can then use tooling and functionality available within the plugin runtime to implement its own features.

## Config API

Tanzu Plugin Runtime provides various config api methods to perform CRUD operations on Contexts, Servers, DiscoverySources, Features, Envs etc.

For more details about the design and apis go to [Config API](docs/config.md)
