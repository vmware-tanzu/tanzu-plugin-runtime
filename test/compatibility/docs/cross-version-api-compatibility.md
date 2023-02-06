# Cross-Version API Compatibility Testing documentation

## Summary

Cross-Version API Compatibility Test Framework provides test helper functions to write compatibility test cases.

Test writer is responsible for writing the go tests using the framework provided helper methods.

As a test writer you should be knowledgeable on specific APIs and specific Runtime libraries i.e. what arguments each API takes and what it returns.
Test writers should also be familiar with how to write combination tests with multiple APIs i.e. when testing mutators(Set***, Delete*** APIs) and reader(Get*** APIs) together.
When writing a test involving two versions of Runtime library APIs a test writer should be aware of what specific fields are added or removed from the API method for those 2 versions and design tests accordingly.

### Directory Structure

```shell
├── core # compatibility core module contains types and helpers methods
├── docs # compatibility testing docs on how to write a test and framework details
├── framework # compatibility framework contains helper methods to contruct API Commands, Test Case builder, executer, validator and actual tests
│   ├── compatibilitytests # compatibility test cases
│   │   ├── context # context api compatibility test cases
│   ├── executer # Functions to execute the testcase
│   │   └── testcase_executer.go
│   └── validators # validator functions to validate the input and output options of API Commands
│       └── testcase_validator.go
└── testplugins  # Specific Runtime version Test Plugin to trigger APIs
    ├── bin # Generated test plugin binaries
    ├── runtime-test-plugin-latest
    ├── runtime-test-plugin-v0_11_6
    ├── runtime-test-plugin-v0_25_4
    └── runtime-test-plugin-v0_28_0
```

### Commands

Build and Run compatibility tests

```shell
make compatibility-tests
```

Build all the runtime specific version test plugins. Binaries of all the test plugins needs to be build to run any compatibility test.

``` shell
make build-compatibility-test-plugins
```

Run all compatibility tests

``` shell
make run-compatibility-tests
```

### How are the Tests being run ?

The tests are executed as a GitHub runner CI pipeline `Tanzu Plugin Runtime Compatibility Tests`

### What will the tests cover ?

Interactions of different implementations of APIs across various versions of library works as expected.

### What will be the test results ?

GitHub CI runner pipeline include Logs with details on test cases from test-case that are executed and succeeded/failed.

## API and Test cases details

### Test Suite API

Framework has `framework.NewTestCase` to write a specific test case, each test case consists of a sequence of commands, with each command corresponding to an invocation of a API from a specific version of runtime.

``` go
// Construct series of commands to execute
 testCase := framework.NewTestCase().Add(setContextCommand).Add(getContextCommand)
```

### High-level test API

To write compatibility test cases, the test writer should be aware of below high level framework Test APIs, each below API helps to test specific runtime API by constructing test command:
We need to use the below test helper functions to construct specific test command, the input data for each high level apis changes based on runtime version's its being tested, for more details look into the input parameter's type definition.

``` go
func NewSetXXXCommand(setXXXInputOptions SetXXXInputOptions, setXXXOutputOptions SetXXXOutputOptions) (*Command, error)
func NewGetXXXCommand(getXXXInputOptions GetXXXInputOptions, getXXXOutputOptions GetXXXOutputOptions) (*Command, error)
```

``` go

// NewSetContextCommand contructs a command to make a call to specific runtime version SetContext API.
// Input Parameter: setContextInputOptions has all input parameters which are required for Runtime SetContext API.
// Input Parameter: setContextOutputOptions has details about expected output from Runtime SetContext API call.
// Return: A command to execute or error if any validations fails for SetContextInputOptions or SetContextOutputOptions
// This method does validates the input parameters  SetContextInputOptions/SetContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetContextInputOptions/SetContextOutputOptions definition(and CtxOptions struct, which is embedded).
func NewSetContextCommand(setContextInputOptions SetContextInputOptions, setContextOutputOptions SetContextOutputOptions) (*Command, error)

func NewGetContextCommand(getContextInputOptions GetContextInputOptions, getContextOutputOptions GetContextOutputOptions) (*Command, error)

func NewDeleteContextCommand(delContextInputOptions DeleteContextInputOptions, deleteContextOutputOptions DeleteContextOutputOptions) (*Command, error)
```

More details about input parameters for above high level API:

``` go
const (
  Version0110 RuntimeVersion = "v0.11.0"
  Version0250 RuntimeVersion = "v0.25.0"
  Version0280 RuntimeVersion = "v0.28.0"
  Version100 RuntimeVersion = "v1.0.0"
)

type RuntimeVersion string

type RuntimeAPIVersion struct{
 RuntimeVersion RuntimeVersion
}

type GetContextInputOptions struct {
 *core.RuntimeAPIVersion        // required
 ContextName             string // required
}

type GetContextOutputOptions struct {
 *core.RuntimeAPIVersion                         // required
 *ContextOpts                                    // For specific version options look into ContextOpts definition
 ValidationStrategy      core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
 Error                   string                  // expected error message could be the sub string of actual error message
}

type SetContextInputOptions struct {
 *core.RuntimeAPIVersion      // required
 *ContextOpts                 // required
 SetCurrentContext       bool // required
}

type SetContextOutputOptions struct {
 ValidationStrategy core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
 Error              string                  // expected error message could be the sub string of actual error message
}

type DeleteContextInputOptions struct {
 *core.RuntimeAPIVersion        // required
 ContextName             string // required
}

type DeleteContextOutputOptions struct {
 ValidationStrategy core.ValidationStrategy // Type of validation to be performed i.e. exact or partial. default is partial
 Error              string                  // expected error message could be the sub string of actual error message
}
```

### xxxOpts structs

- Each of these struct represents the super set of all fields for their respective config types (e.g .Context, Server, DiscoverySource) across for all Runtime versions in which the types are defined.
- The xxxOpts structs are used as inputs to the API invocation as well as data to verify against the output of the invocation. The implications of setting or not setting certain attributes changes depend on which scenario under which it is being used.
- Command Creation Helper functions (i.e. NewSetContextCommand) validate the supplied inputOptions as per RuntimeVersion and make sure all required attributes are set. If not Command Creation Helper functions will throw error and Test fails.
- The validation of “Input and expect output” happens in Set/Get/DeleteXXXCommand creation (within framework.NewTestCase), so it's not part of test case execution. so its kind of test case setup.
- Input struct: Framework validates the supplied Input data (eg: GetContextInputOptions/SetContextInputOptions/DeleteContextInputOptions) as per specified runtime version supported fields.
- Output struct: Framework validates the supplied outputOptions as per specified runtime version whether certain fields are supported for the runtime version used and compares the supplied data with what is returned from API method i.e. every non-null field in output struct will be checked for equality with returned result.
  (Eg: GetContextOutputOptions/SetContextOutputOptions/DeleteContextOutputOptions) as its test writer responsibility, the Output data is depends on the Input data (eg: if input is not valid, runtime API may return error) OR the sequence of calls made before the current call (eg: NewGetContextCommand() output depends on the previous SetContext** calls).
  If test writer is not passing valid expected output data, then it fails during the TEST CASE EXECUTION as test case fail.
- Test framework does validate the expected Output struct’s field values with runtime API response to its entirety.
- Test writer does not have to explicitly specify to check whether an attribute value is nil as all fields default values are nil or “”.
- Test writer will need to specify what exact object attribute values he expects out of GetXX API method (framework perform a full equality check and not partial equality, means if API returns attribute which not specified by test writer then test case will fail).
- So for the output struct, all filed are optional, there are NO MANDATORY fields!!! its test writer responsibility to identify the expected data!!

Example 1: For the Below NewSetContextCommand if the context arguments are passed incorrect

``` go
// Input Parameters for Runtime SetContext API
   setContextInputOptions := SetContextInputOptions{
        RuntimeAPIVersion: &core.RuntimeAPIVersion{
          RuntimeVersion: core.Version0280,
        },
       CtxOptions: CtxOptions{
           Name: "context-one",
           Type: framework.CtxTypeK8s, // Invalid attribute since Type is not supported in v0.28.0
           GlobalOpts: &framework.GlobalServer{
             Endpoint: "test-endpoint",
          }
       }
   }
   // Create SetContext Command
   setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
   Expect(err).To(BeNil()) // Error is thrown since Type is not supported in specified runtime version v0.28.0

Below Error is thrown when the test case with above command input options is run
    Expected
        <*errors.fundamental | 0x140000b8030>: {
            msg: "invalid set context input options for the specified runtime version v0.28.0",
            stack: [0x1051f2ac9, 0x1051f1bf0, 0x1051f4544, 0x10500a120, 0x105009c64, 0x105009474, 0x10500c478, 0x10500bec8, 0x1050285c0, 0x10502833c, 0x105027ad0, 0x105029b5c, 0x105034da0, 0x105034c44, 0x1051f4d38, 0x104e27fd0, 0x104d91814],
        }
    to be nil
```

Example 2  How framework validates the null fields and what is expected of test writer

SetContext V0.28.0 with ClusterOpts.Endpoint value set
then SetContext v0.28.0 with ClusterOpts.Endpoint value unset
then GetContext v0.28.0 with two different expectations and behaviour explained below

``` go
It("SetContext v0.28.0 SetContext v0.28.0(Unsetting ClusterOpts.Endpoint) GetContext v0.28.0", func() {
   // Input Parameters for Runtime SetContext API V0.28.0
   setContextInputOptions := &framework.SetContextInputOptions{
      RuntimeAPIVersion: &framework.RuntimeAPIVersion{
         RuntimeVersion: framework.Version0280,
      },
      ContextOpts: &framework.ContextOpts{
         Name:   "context-one",
         Target: framework.TargetTMC,
         ClusterOpts: &framework.ClusterServerOpts{
            Endpoint: "test-endpoint",
            Path:     "test-path",
         },
      },
   }

   // Input Parameters for Runtime SetContext API V0.28.0 With ClusterOpts.Endpoint unset
   setContextInputOptionsWithEndpointUnset := &framework.SetContextInputOptions{
      RuntimeAPIVersion: &framework.RuntimeAPIVersion{
         RuntimeVersion: framework.Version0280,
      },
      ContextOpts: &framework.ContextOpts{
         Name:   "context-one",
         Target: framework.TargetTMC,
         ClusterOpts: &framework.ClusterServerOpts{
            Endpoint: "",
            Path:     "test-path",
         },
      },
   }

   // Create SetContext Command 1
   setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, nil)
   Expect(err).To(BeNil())

   // Create SetContext Command WithEndpointUnset
   setContextCommandWithEndpointUnset, err := framework.NewSetContextCommand(setContextInputOptionsWithEndpointUnset, nil)
   Expect(err).To(BeNil())

   // Input Parameters for Runtime GetContext API
   getContextInputOptions := &framework.GetContextInputOptions{
      RuntimeAPIVersion: &framework.RuntimeAPIVersion{
         RuntimeVersion: framework.Version0280,
      },
      ContextName: "context-one",
   }

   // Output Parameters for Runtime GetContext API
   getContextOutputOptions := &framework.GetContextOutputOptions{
      RuntimeAPIVersion: &framework.RuntimeAPIVersion{
         RuntimeVersion: framework.Version0280,
      },
      ContextOpts: &framework.ContextOpts{
         Name:   "context-one",
         Target: framework.TargetTMC,
         ClusterOpts: &framework.ClusterServerOpts{
            Endpoint: "test-endpoint",
            Path:     "test-path",
         },
      },
   }

   getContextOutputOptionsWithEndpointNotExpected := &framework.GetContextOutputOptions{
      RuntimeAPIVersion: &framework.RuntimeAPIVersion{
         RuntimeVersion: framework.Version0280,
      },
      ContextOpts: &framework.ContextOpts{
         Name:   "context-one",
         Target: framework.TargetTMC,
         ClusterOpts: &framework.ClusterServerOpts{
            Path: "test-path",
         },
      },
   }

   // Create GetContextAPIName Command
   getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
   Expect(err).To(BeNil())

   // Create GetContextAPIName Command WithEndpointNotExpected
   getContextCommandWithEndpointNotExpected, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptionsWithEndpointNotExpected)
   Expect(err).To(BeNil())

   // Construct series of commands to execute
   testCase1 := framework.NewTestCase().Add(setContextCommand).Add(setContextCommandWithEndpointUnset).Add(getContextCommand)
   // Executes the commands from the list and validates the expected output with actual output
   testCase1.Execute() // This execution fails since ClusterOpts.Endpoint is unset in setContextCommandWithEndpointUnset but expected in getContextCommand

   testCase2 := framework.NewTestCase().Add(setContextCommand).Add(setContextCommandWithEndpointUnset).Add(getContextCommandWithEndpointNotExpected)
   // Executes the commands from the list and validates the expected output with actual output
   testCase2.Execute() // This execution succeeds since ClusterOpts.Endpoint is unset in setContextCommandWithEndpointUnset and not expected in getContextCommandWithEndpointNotExpected
})
```

#### xxxOpts struct example below for Context config

### ContextOpts struct

ContextOpts is the super set of parameters for Context, for all Runtime Versions. Based on the Runtime Version CtxOptions attributes may change(mandatory/optional/Not applicable).
Command Helper functions (i.e. NewSetContextCommand) validate the supplied inputOptions as per RuntimeVersion and make sure all required attributes are set.
If not Command Helper functions will throw error and Test fails.

Below table explains about each attribute of CtxOptions requirement based on Runtime Version.

#### Context Attributes Compatibility Matrix

|    Attribute     | v1.0.0 | v0.28.0 | v0.25.0 | v0.11.0 |
|:----------------:|:------:|:-------:|:-------:|:-------:|
|       Name       |   M    |    M    |    M    |    M    |
|      Target      |   M    |    M    |   N/A   |   N/A   |
|       Type       |  N/A   |    M    |    M    |    M    |
|    GlobalOpts    |   M    |    M    |    M    |    M    |
|    ServerOpts    |   M    |    M    |    M    |    M    |
| DiscoverySources |   O    |    O    |    O    |    O    |
| IsCurrentContext |   O    |    O    |    O    |    O    |

*Either one of GlobalOpts or ServerOpts is required.

- M: Mandatory
- O: Optional
- N/A: Not Applicable

#### Context struct Definition

``` go
// Look below for each attribute documentation for which version its been supported/mandatory/optional.
type ContextOpts struct {
 // Name of the context.
 // required for all runtime versions till v0.28.0
 Name string `json:"name,omitempty" yaml:"name,omitempty"`
  // Target of the context.
  // required for runtime versions from v0.28.0
 Target Target `json:"target,omitempty" yaml:"target,omitempty"`
 // Type of the context.
 // required for all runtime versions from v0.25.0 to v0.27.0
 Type ContextType `json:"type,omitempty" yaml:"type,omitempty""`
 // GlobalOpts if the context is a global control plane (e.g., TMC).
 // Either global opts or clusterOpts should be specified for all runtime versions till v0.28.0
 GlobalOpts *GlobalServer `json:"globalOpts,omitempty" yaml:"globalOpts,omitempty"`
 // ClusterOpts if the context is a kubernetes cluster.
 // Either global opts or clusterOpts should be specified for all runtime versions till v0.28.0
 ClusterOpts *ClusterServer `json:"clusterOpts,omitempty" yaml:"clusterOpts,omitempty"`
 // DiscoverySources determines from where to discover plugins associated with this context.
 // supported from v0.25.0 ; optional
 DiscoverySources []DiscoverySourceOptions `json:"discoverySources,omitempty" yaml:"discoverySources,omitempty""`
// supported from v0.25.0 ;optional; default is false
 IsCurrentContext bool `json:"isCurrentContext,omitempty" yaml:"isCurrentContext,omitempty"`
}
```

## Sample test case Referred as *testcase-1*

``` go
It("Run Runtime V100 SetContext API and Runtime V0280 GetContext API", func() {
// Input Parameters for Runtime SetContext API
setContextInputOptions := &framework.SetContextInputOptions{
 RuntimeAPIVersion: &framework.RuntimeAPIVersion{
   RuntimeVersion: framework.Version100,
   },
   ContextOpts: &framework.ContextOpts{
    Name:   "context-one",
    Target: framework.TargetK8s,
            GlobalOpts: &framework.GlobalServerOpts{
               Endpoint: "test-endpoint",
            },
         },
      }

      // Output Parameters for Runtime SetContext API
      var setContextOutputOptions *framework.SetContextOutputOptions

      // Input Parameters for Runtime GetContext API
      getContextInputOptions := &framework.GetContextInputOptions{
         RuntimeAPIVersion: &framework.RuntimeAPIVersion{
            RuntimeVersion: framework.Version100,
         },
         ContextName: "context-one",
      }

      // Output Parameters for Runtime GetContext API
      getContextOutputOptions := &framework.GetContextOutputOptions{
         ContextOpts: &framework.ContextOpts{
            Name:   "context-one",
            Target: framework.TargetK8s,
            GlobalOpts: &framework.GlobalServerOpts{
               Endpoint: "test-endpoint",
            },
         },
      }

      // Create SetContext Command
      setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
      Expect(err).To(BeNil())

      // Create GetContext Command
      getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
      Expect(err).To(BeNil())

      // Construct series of commands to execute

      testCase := framework.NewTestCase().Add(setContextCommand).Add(getContextCommand) // re-named from NewTestCommands

      // Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
      testCase.Execute()
 })

```

For more details on framework go to [Cross_Version_API Compatibility Framework](cross-version-api-compatibility-framework.md)
