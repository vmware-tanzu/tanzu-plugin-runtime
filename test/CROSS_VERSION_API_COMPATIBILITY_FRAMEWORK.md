# Cross Version API Compatibility Framework

## Summary

This section describes on how to write a go test for cross version API compatibility and how the test case is processed by the framework to then invoke various Runtime APIs and perform validations.
Each test case command is transformed and written into temporary files which is consumed by the `runtime-test-plugins-vX.XX.X` binaries build with specific versions of the runtime.

## Framework

Each Test is of below type

``` go
package framework

// TestCase represents the list of commands to execute as part of test case
type TestCase struct {
   Commands []*Command `json:"commands" yaml:"commands"`
}

// Command represents the list of apis to execute as part of command execution
type Command struct {
   APIs []*API `json:"apis" yaml:"apis"`
}

// API represents the runtime api to execute
type API struct {
   Name      RuntimeAPIName         `json:"name" yaml:"name"`
   Version   RuntimeVersion         `json:"version" yaml:"version"`
   Arguments map[string]interface{} `json:"arguments" yaml:"arguments"`
   Output    *Output                `json:"output" yaml:"output"`
}

// Output represents the runtime api expected output for validation
type Output struct {
   Result  Result `json:"result" yaml:"result"`
   Content string `json:"content" yaml:"content"`
}

type Result string

const (
   Success Result = "success"
   Failed         = "failed"
)

// RuntimeAPIVersion represents the runtime library version
type RuntimeAPIVersion struct {
   RuntimeVersion RuntimeVersion `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`
}

// RuntimeVersion Runtime library versions
type RuntimeVersion string

const (
   Version0116 RuntimeVersion = "v0.11.6"
   Version0254                = "v0.25.4"
   Version0280                = "v0.28.0"
   Version100                 = "v1.0.0"
)

// NewTestCase creates an instance of TestCase
func NewTestCase() *TestCase {
   return &TestCase{}
}

// Add series of commands to test case to be executed in sequence
func (t *TestCase) Add(command ...*Command) *TestCase {
   t.Commands = append(t.Commands, command...)
   return t
}

// APILog represents the logs/output/errors returned from runtime test plugin binaries
type APILog struct {
   APIResponse *APIResponse `json:"apiResponse" yaml:"apiResponse"`
   APIError    string       `json:"error" yaml:"error"`
}

// APIResponse represents the output response returned from runtime test plugin binaries
type APIResponse struct {
   ResponseType ResponseType `json:"responseType" yaml:"responseType"`
   ResponseBody interface{}  `json:"responseBody" yaml:"responseBody"`
}

type ResponseType string

const (
   MapResponse     ResponseType = "map"
   BooleanResponse              = "bool"
   StringResponse               = "str"
   IntegerResponse              = "int"
)
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

Running the  *testcase_1*, the framework generates below commands internally to trigger specific runtime version libraries.

## Internal Implementation

### Below command is generated to run Runtime lib v1.0.0 SetContext API method as per testcase-1

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

### Below command is generated to run Runtime lib v0.28.0 GetContext API method as per testcase-1

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

### Combination / Co-existence Testing

Below table describes all possible combination tests for all supported Runtime APIs.

|  _  |  V  | ALL | ALL | ALL | ALL |
|:---:|:---:|:---:|:---:|:---:|:---:|
|  V  |  _  | V1  | V28 | V25 | V11 |
| ALL | V1  |  Y  |  Y  |  Y  |  Y  |
|  ^  | V28 |  Y  |  Y  |  Y  |  Y  |
|  ^  | V25 |  Y  |  Y  |  Y  |  Y  |
|  ^  | V11 |  Y  |  Y  |  Y  |  Y  |

| Symbol | Description                                                                                                                                                                                                                                                                                                                                                                                                                        |
|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| V      | Runtime Library Version                                                                                                                                                                                                                                                                                                                                                                                                            |
| V1     | Runtime Library Version v1.0.0                                                                                                                                                                                                                                                                                                                                                                                                     |
| V28    | Runtime Library Version v0.28.0                                                                                                                                                                                                                                                                                                                                                                                                    |
| V25    | Runtime Library Version v0.25.4                                                                                                                                                                                                                                                                                                                                                                                                    |
| V11    | Runtime Library Version v0.11.6                                                                                                                                                                                                                                                                                                                                                                                                    |
| ALL    | SetContext, GetContext, DeleteContext, GetCurrentContext, SetCurrentContext, DeleteCurrentContext, GetServer, SetServer, DeleteServer, GetCurrentServer, SetCurrentServer, DeleteCurrentServer, IsFeatureEnabled, DeleteFeature, SetFeature, ConfigureDefaultFeatureFlagsIfMissing, GetEnv, SetEnv, DeleteEnv, GetCLIPluginDiscoverySources, SetCLIPluginDiscoverySources, DeleteCLIPluginDiscoverySources, GetEdition, SetEdition |
