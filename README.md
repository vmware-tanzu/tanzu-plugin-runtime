# Tanzu Plugin Runtime

:warning: NOTE: This repository is still under early development. We plan to
make releases available for evaluation in the second quarter of 2023.  Please
watch this page for further updates.

## Overview

The Tanzu CLI is based on a plugin architecture. This architecture enables
teams to build, own, and release their own piece of functionality as well as
enable external partners to integrate with the system.

The Tanzu Plugin Runtime is a library that provides functionality and helper
methods to develop Tanzu CLI plugins.

Developers begin plugin development by using the `builder` plugin to bootstrap
a new plugin project. The code generated in the project relies on the runtime
to provide some functionality common to all plugins. For more information about
the development process, see the (VVV update link) [Tanzu CLI Plugin Development guide](https://github.com/vuil/tanzu-cli/blob/docs-draft/docs/dev/main.md)

## The library

This Tanzu Plugin Runtime broadly consists of:

1. CLI UX Component library
2. Configuration library
3. Plugin integration
4. Command helpers
5. Test helpers

### CLI UX Component Library

This package implements reusable CLI user interface components, including:

- output writers (table, listtable, json, yaml, spinner)
- prompt
- selector
- question

### Configuration Library

This package implements helper functions to read, write and update various
Tanzu CLI configuration objects like Contexts, DiscoverySources, CLI
Features and environment settings.

For more details about the design and APIs go to [Configuration API](docs/config.md)

### Plugin integration

This package implements helper functions for new plugin creation. This is one
of the main packages that each and every plugin will need to import to
integrate with the Tanzu CLI. For more information about
the development process, see the (VVV update link) [Tanzu CLI Plugin Development guide](https://github.com/vuil/tanzu-cli/blob/docs-draft/docs/dev/main.md)

### Command Helpers

This package implements command specific helper functions like command deprecation, etc.

### Test Helpers

This package implements helper functions to develop test plugins and Cross-version Configuration Library APIs compatibility testing

Besides unit and integration tests, the runtime APIs are also being tested
in the presence of other versions runtime client code as part of cross-version API compatibility testing.
These tests are important in ensuring interoperability among runtime clients (and hence
different generations of CLI plugins)

#### Cross Version API Compatibility Testing

Cross Version API Compatibility testing ensures that [Config
APIs](docs/config.md) of different Tanzu Plugin Runtime versions work as
expected as long as those Tanzu Plugin Runtime versions are supported.

Cross Version API Compatibility testing simulates the interaction of Configuration library APIs
provided by different versions of plugin runtimes to ensure that they can
interoperate with one another.

For more details go to [Cross-version API Compatibility](test/compatibility/docs/cross-version-api-compatibility.md)
