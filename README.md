# Tanzu Plugin Runtime

:warning: NOTE: This repository is still under early development. We hope to be
make Alpha releases available for evaluation in March 2023.  Please watch this
page for further updates.

## Overview

The Tanzu CLI is based on a plugin architecture. This architecture enables teams to build, own, and release their own piece of functionality as well as enable external partners to integrate with the system. The Tanzu Plugin Runtime provides functionality and helper methods to develop Tanzu CLI plugins.

Developers can use the `Builder` admin plugin to bootstrap a new plugin which can then use tooling and functionality available within the plugin runtime to implement its own features.

## Config API

Tanzu Plugin Runtime provides various config api methods to perform CRUD operations on Contexts, Servers, DiscoverySources, Features, Envs etc.

For more details about the design and apis go to [Config API](docs/config.md)

## Backward Compatibility Testing

### Overview

This testing ensures that plugins developed with different Runtime versions work as expected as long as those runtime versions are supported.
Current testing verifies that plugin developed with new Tanzu Plugin Runtime (Ex :- v1.0.0) works along with the plugin developed with old Tanzu Plugin Runtime versions like v0.28.0, v0.25.0, v0.11.0.

For more details go to [Backward Compatibility](BACKWARD_COMPATIBILITY.md)
