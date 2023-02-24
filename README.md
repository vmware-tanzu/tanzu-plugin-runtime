# Tanzu Plugin Runtime

:warning: NOTE: This repository is still under early development. We hope to be
make Alpha releases available for evaluation in March 2023.  Please watch this
page for further updates.

## Overview

The Tanzu CLI is based on a plugin architecture. This architecture enables teams to build, own, and release their own piece of functionality as well as enable external partners to integrate with the system. The Tanzu Plugin Runtime provides functionality and helper methods to develop Tanzu CLI plugins.

Developers can use the `Builder` admin plugin to bootstrap a new plugin which can then use tooling and functionality available within the plugin runtime to implement its own features.

## Config API

Tanzu Plugin Runtime provides various config API methods to perform CRUD operations on Contexts, Servers, DiscoverySources, Features, Envs etc.

For more details about the design and APIs go to [Config API](docs/config.md)

## Cross Version API Compatibility Testing

### Overview

Cross Version API Compatibility testing ensures that [Config APIs](docs/config.md) of different Tanzu Plugin Runtime versions work as expected as long as those Tanzu Plugin Runtime versions are supported.
Cross Version API Compatibility testing simulates the interaction of APIs provided by different versions of plugin runtimes to ensure that they can interoperate with one another.

For more details go to [Cross-version API Compatibility](./test/cross-version-api-compatibility.md)
