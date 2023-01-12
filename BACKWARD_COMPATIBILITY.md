
# Backward Compatibility

Verify that plugin developed with new Tanzu Plugin Runtime (Ex v1.0.0) works along with plugin developed with old Tanzu Plugin Runtime(Ex v0.28.0).

## Framework

## Phase 1

### Summary

Compatibility Test Framework provides test helper functions to write `go test cases`.

Test writer is responsible for writing the go tests using the framework provided helper methods.
As a test writer you should be knowledgeable on specific APIs and specific Runtime libraries  i.e. what arguments each API takes and what it returns.
Test writers should also be familiar with how to write combination tests with multiple APIs i.e when testing mutators and reader APIs together.

For Example:-  Running SetContext alone in the test will not be a good test as it does not validate the actual context that is set in the config yaml file. But when SetContext is combined with GetContext/ GetCurrentContext making the tests more solid.

When writing a test involving two versions of Runtime library APIs a test writer should be aware of what specific fields are added or removed from the API method for those 2 versions and design tests accordingly.

### Commands

`make backward-compatibility-tests` - Will internally run all `**_test.go` files from compatibility-tests directory.

### How to write a test?

Each Test is of below type

``` go
type TestSuite struct {
    // Name of the test suite
  Name      string      `json:"name" yaml:"name"`
  TestCases []*TestCase `json:"testcases" yaml:"testcases"`
}
type TestCase struct {
    // Name of the test case
  Name     string     `json:"name" yaml:"name"`
  // Series of commands to execute
  Commands []*Command `json:"commands" yaml:"commands"`
}

type Command struct {
    // Versions of Runtime library to run the APIs
  Versions []string `json:"versions" yaml:"versions"`
  // Runtime library APIs to run
  APIs     []*API   `json:"apis" yaml:"apis"`
}

type API struct {
    // Name of the Runtime API method
  Name      string                 `json:"name" yaml:"name"`
  // Runtime API method parameters
  Arguments map[string]interface{} `json:"arguments" yaml:"arguments"`
  // Runtime API method response object for validation
  Output    *Output                `json:"output" yaml:"output"`
}

type Output struct {
    // Result of the Runtime API method Ex: success, failed
  Result  string `json:"result" yaml:"result"`
  // Context is the yaml representation of expected response of the specific Runtime API method
  Content string `json:"content" yaml:"content"`
}
```

**Arguments Type**
Map of key value pairs that needs to be specified for each API method. Each key is the name of the argument that the API accepts.
The content of arguments needs to be specified as per definition of API method.
The content of `context` should be the yaml representation of the Context object that SetContext v1.0.0 accepts.
`arguments` Keys will be predefined in the framework to help Test writers to construct the specific API test.
Possible keys for arguments map are :- context, server, isCurrent, name, discoverySources,val,repository,target,key, plugin,feature, patchStrategies etc.

Each Test Case accepts

- `name` Name of the Test Case specifying the various APIs and versions of runtime are being tested
- `commands` Each Command takes in array of versions and apis to be tested.
- Each `command` takes
  - `name` API Method Name .
  - `arguments` API Method Arguments.
  - `output` API Method Expected Output for validation.

Example:-  Runs Runtime V1.0.0 SetContext and RuntimeV0.28.0 GetContext API methods.

TestCase-1:

``` go
func TestContextAPI(t *testing.T) {
    // Build Test Suite
    testSuite := types.NewTestSuite(
        types.WithTestSuiteName("Context API"),
        types.WithTestCases(
            types.NewTestCase(
                types.WithTestCaseName("Runtime v1.0.0 SetContext and Runtime v0.28.0 GetContext"),
                types.WithTestCaseCommands(
                    types.NewCommand(
                        types.WithCommandVersions([]string{"v1.0.0"}),
                        types.WithCommandAPIs(
                            helpers.APIForSetContext("v1.0.0"),
                        ),
                    ),
                    types.NewCommand(
                        types.WithCommandVersions([]string{"v0.28.0"}),
                        types.WithCommandAPIs(
                            helpers.APIForGetContext("v1.0.0"),
                        ),
                    ),
                ),
            ),
        ),
    )
    // Execute Test Suite
    response := testSuite.Execute()
    // Perform Assertion
    assert.Equal(t, response, "success")
}
```

``` go
func (t *TestSuite) Execute() string
```

Execute() will parse the TestSuite and run each command and version as specified and validate the return response of each command execution.

``` go
// helpers.go
func APIForGetContext(version string, options ...types.APIOption) *API
func APIForSetContext(version string, options ...types.APIOption) *API
```

*options can accept arguments and outputContent - default values are used if not provided

Running the above *testcase-1* Internally the framework generates below commands.

#### Below command is generated to run Runtime lib v1.0.0 SetContext API method as per testcase-1

```shell
stdout, _, err := Exec("./runtime-test-plugin-1-00/runtime-test-plugin-1-00 test --test-file temp-tests-100.yaml")
```

*Exec() internally uses “os/exec” pkg

`temp-tests-100.yaml` will have the []API to run for that specific runtime version

```yaml
     apis:
       - name: SetContext
         arguments:
           context: |-
               name: test-mc
               target: kubernetes
               clusterOpts:
                 isManagementCluster: true
                 endpoint: old-test-endpoint
                 annotation: one
           isCurrent: false
         output:
           result: success
           content: ""
```

#### Below command is generated to run Runtime lib v0.28.0 GetContext API method as per testcase-1

```shell
stdout, _, err := Exec("./runtime-test-plugin-0-28/runtime-test-plugin-0-28 test --test-file temp-tests-028.yaml")
```

*Exec() internally uses “os/exec” pkg

`temp-tests-028.yaml` will have the []API to run for that specific runtime version

``` yaml
apis:
  - name: GetContext
    arguments:
      contextName: test-bc
    output:
      result: success
      content: |-
        name: test-mc
        target: kubernetes
        clusterOpts:
          isManagementCluster: true
          endpoint: old-test-endpoint
```

### Implementation details of runtime-test-plugin-{version}

``` shell
runtime-test-plugin-1-00 test –test-file temp-tests-100.yaml
```

- test command will parse the API method and arguments  from –test-file flag temp-tests.yaml then triggers the runtime lib APIs.
- Validations or Assertions are done based on passed Output arguments.

``` go
// Call runtime SetContext() API function
err := SetContext(ctx, isCurrent)
// Print the log to stdout
```

``` shell
runtime-test-plugin-0-28 test –test-file temp-tests-028.yaml
```

``` go
// Call runtime GetContext() API function
actualCtx := GetContext(ctxName) //i.e. test-bc

// Print the log to stdout
```

### Combination / Co - Existence Testing

|  X  |  V  | ALL | ALL | ALL | ALL |
|:---:|:---:|:---:|:---:|:---:|:---:|
|  V  |  X  | V1  | V28 | V25 | V11 |
| ALL | V1  |  Y  |  Y  |  Y  |  Y  |
|  ^  | V28 |  Y  |  Y  |  Y  |  Y  |
|  ^  | V25 |  Y  |  Y  |  Y  |  Y  |
|  ^  | V11 |  Y  |  Y  |  Y  |  Y  |

| Symbol | Description                                                                                                                                                                                                                                                                                                                                                                                                                        |
|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| V      | Runtime Library Version                                                                                                                                                                                                                                                                                                                                                                                                            |
| V1     | Runtime Library Version v1.0.0                                                                                                                                                                                                                                                                                                                                                                                                     |
| V28    | Runtime Library Version v0.28.0                                                                                                                                                                                                                                                                                                                                                                                                    |
| V25    | Runtime Library Version v0.25.0                                                                                                                                                                                                                                                                                                                                                                                                    |
| V11    | Runtime Library Version v0.11.0                                                                                                                                                                                                                                                                                                                                                                                                    |
| ALL    | SetContext, GetContext, DeleteContext, GetCurrentContext, SetCurrentContext, DeleteCurrentContext, GetServer, SetServer, DeleteServer, GetCurrentServer, SetCurrentServer, DeleteCurrentServer, IsFeatureEnabled, DeleteFeature, SetFeature, ConfigureDefaultFeatureFlagsIfMissing, GetEnv, SetEnv, DeleteEnv, GetCLIPluginDiscoverySources, SetCLIPluginDiscoverySources, DeleteCLIPluginDiscoverySources, GetEdition, SetEdition |
