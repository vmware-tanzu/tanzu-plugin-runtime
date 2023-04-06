# Tanzu CLI Test Framework

This is a Go package that provides a test framework for Tanzu CLI plugins.
It includes functions and structures for creating a main test and individual subtests.
The tests run CLI commands and report on their success or failure.

## Usage

To use the framework package in your CLI plugin tests, import it as follows:

``` go
import "github.com/vmware-tanzu/tanzu-plugin-runtime/test/framework"
```

Create a new main test using the `NewMain` function, providing the main test name and a `CleanupFunc` function to be executed at the end of the test (if necessary).

``` go
  m := framework.NewMain("my-command", myCommand, myCleanupFunc)
```

- `my-command` is the name of your test suite.
- `myCommand` is a cobra.Command object that represents the CLI command you want to test.
- `myCleanupFunc` is a cleanup function that is executed at the end of the test (if specified).

The `NewMain` function takes a pointer to a `cobra.Command` as a parameter, so you need to have created this command before calling `NewMain`.

You can also pass in two optional bool parameters, `printReport` and `deferDelete`, which control whether a report should be printed and whether resource deletion should be deferred, respectively.

Create tests using the `NewTest` function, providing the test name, a slice of arguments to pass to the CLI command, and a boolean indicating whether the command should succeed or fail.

Add the tests to the main test using the `AddTest` function.

Run the tests using the `Run` method of the main test.

Here is an example usage of how to use the framework:

- This code is a simple CLI (Command Line Interface) written in Go, using the Cobra library. The CLI provides a single command, `builder`, which has a sub-command called init.

- When builder init is executed with the argument `myrepo --dry-run`, the CLI is expected to create a new repository named `myrepo`. However, since the `--dry-run` flag is specified, no actual action will be taken, and the CLI will simply print out what it would have done.

- To test this functionality, the test function is defined, which creates a new instance of `clitest.Main`, sets up the test, and executes it. The test verifies that the `myrepo` string is contained in the output of the builder init command by calling `t.ExecContainsString("myrepo")`.

- The `Cleanup` function is defined to clean up any resources that may have been created during the test, although in this case it is not used.

- The `main` function sets up a new plugin using the `plugin.NewPlugin` function, passing in the `clitest.NewTestFor("builder")` object as a descriptor. The `p.Cmd.RunE` field is set to the test function, and the plugin is executed using `p.Execute()`.

``` go
package main

import (
    "log"
    "os"

    "github.com/spf13/cobra"

    "github.com/vmware-tanzu/tanzu-framework/cli/runtime/plugin"
    clitest "github.com/vmware-tanzu/tanzu-framework/cli/runtime/test"
)

var descriptor = clitest.NewTestFor("builder")

func main() {
    p, err := plugin.NewPlugin(descriptor)
    if err != nil {
        log.Fatal(err)
    }
    p.Cmd.RunE = test
    if err := p.Execute(); err != nil {
        os.Exit(1)
    }
}

func test(c *cobra.Command, _ []string) error {
    m := clitest.NewMain("cluster", c, Cleanup)
    defer m.Finish()

    err := m.RunTest(
        "create a repo",
        "builder init myrepo --dry-run",
        func(t *clitest.Test) error {
            err := t.ExecContainsString("myrepo")
            if err != nil {
                return err
            }
            return nil
        },
    )
    if err != nil {
        return err
    }
    return nil
}

// Cleanup the test.
func Cleanup() error {
    return nil
}

```
