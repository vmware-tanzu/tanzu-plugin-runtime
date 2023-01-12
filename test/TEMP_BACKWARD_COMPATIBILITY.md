# Backward Compatibility Testing Draft documentation

This testing ensures that plugins developed with different Runtime versions work as expected as long as those runtime versions are supported. Current testing verifies that plugin developed with new Tanzu Plugin Runtime (Ex :- v1.0.0) works along with the plugin developed with old Tanzu Plugin Runtime versions like v0.28.0, v0.25.0, v0.11.0.

### Summary

Compatibility Test Framework provides test helper functions to write compatibility test cases.

Test writer is responsible for writing the go tests using the framework provided helper methods.
As a test writer you should be knowledgeable on specific APIs and specific Runtime libraries  i.e. what arguments each API takes and what it returns.
Test writers should also be familiar with how to write combination tests with multiple APIs i.e when testing mutators and reader APIs together.

For Example:-  Running SetContext alone in the test will not be a good test as it does not validate the actual context that is set in the config yaml file. But when SetContext is combined with GetContext/ GetCurrentContext making the tests more solid.

When writing a test involving two versions of Runtime library APIs a test writer should be aware of what specific fields are added or removed from the API method for those 2 versions and design tests accordingly.

## API and Test cases details
### Test Suite API
The framework has `framework.NewTestSuite` to write a new test suite, and the test suite will have test cases, each test case is independent. Framework has `framework.NewTestCase` to write a specific test case, each test case consist sequence of commands, each command is specific to a runtime api call.
We need to use the above high level APIs to construct specific test command, the input data for each high level api's changes based on runtime version's its being tested, for more details look into the input parameter's type definition.

``` go
// Construct series of commands to execute
 testSuite := framework.NewTestSuite(
		framework.NewTestCase(
			setContextCommand,
			getContextCommand,
		),
	)
```
### High-level test API
To write compatibility test cases, the test writer should be aware of below high level framework Test APIs, each below api helps to test specific runtime api by constructing test command:

``` go

// NewSetContextCommand contructs a command to make a call to specific runtime version SetContext API.
// Input Parameter: setContextInputOptions has all input parameters which are required for Runtime SetContext API.
// Input Parameter: setContextOutputOptions has details about expected ouput from Runtime SetContext API call.
// Return:  command to execute or error if any validations fails for SetContextInputOptions or SetContextOutputOptions
// This method does validates the input parameters  SetContextInputOptions/SetContextOutputOptions based on Runtime API Version
// For more details about supported parameters refer to SetContextInputOptions/SetContextOutputOptions defintion(and CtxOptions struct, which is embedded).
func NewSetContextCommand(setContextInputOptions SetContextInputOptions, setContextOutputOptions SetContextOutputOptions) (*Command, error)

func NewGetContextCommand(getContextInputOptions GetContextInputOptions, getContextOutputOptions GetContextOutputOptions) (*Command, error)

func NewDeleteContextCommand(delContextInputOptions DeleteContextInputOptions, deleteContextOutputOptions DeleteContextOutputOptions) (*Command, error)
```

More details about input parameters for above high level api:

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
 RuntimeAPIVersion // required
 ContextName string // required
}

type GetContextOutputOptions struct{
 RuntimeAPIVersion // required
 CtxOptions // For specific version options look into CtxOptions definition
}

type SetContextInputOptions struct{
 RuntimeAPIVersion // required
 CtxOptions // required
}

type SetContextOutputOptions struct{
 error string // expected error message could be the sub string of actual error message
}

type DeleteContextInputOptions struct{
 RuntimeAPIVersion // required
 ContextName string // required
}

type DeleteContextOutputOptions struct{
 error string // expected error message could be the sub string of actual error message
}
```

### xxxOptions structs
- These structs are super set of parameters for respective configs like Context, Server, DiscoverySources etc. for all Runtime Versions.
- Based on the Runtime Version xxxOptions attributes may change(mandatory/optional/Not applicable).
- It is meant to be used as both inputOptions and outputOptions.
- Command Creation Helper functions (i.e. NewSetContextCommand) validate the supplied inputOptions as per RuntimeVersion and make sure all required attributes are set. If not Command Creation Helper functions will throw error and Test fails.
- Command Creation Helper functions validates the supplied outputOptions as per specified runtime version whether certain fields are supported for the runtime version used and compares the supplied data with what is returned from API method i.e. every non-null field in output struct will be checked for equality with returned result.
 If Optional fields are not set in the expected output validation will ignore checking that field and only validate the supplied data in the outputOptions. 
- The validation of “Input and expect output” happens in  Set/Get/DeleteXXXCommand creation (within framework.NewTestCase), so it's not part of test case execution. so its kind of test case setup. 
- Input struct: Framework validates only Input data (eg: GetContextInputOptions/SetContextInputOptions/DeleteContextInputOptions) as its required to make runtime api call, 
- Output struct: Framework DOES NOT validate expected output data at all (Eg: GetContextOutputOptions/SetContextOutputOptions/DeleteContextOutputOptions) as its test writer responsibility, the Output data is depends on the Input data (eg: if input is not valid, runtime api may return error) OR the sequence of calls made before the current call (eg: NewGetContextCommand() output depends on the previous SetContext** calls). If test writer is not passing valid expected output data, then it fails during the TEST CASE EXECUTION as test case fail. 
- So for the output struct, all filed are optional, there are NO MANDATORY fields!!! its test writer responsibility to identify the expected data!!

Example: For the Below NewSetContextCommand if the context arguments are passed incorrect

``` go
// Input Parameters for Runtime SetContext API
	setContextInputOptions := SetContextInputOptions{
	    RuntimeVersion: framework.Version0280,
	    CtxOptions: CtxOptions{
	        Name: "context-one",
	        Type: framework.CtxTypeK8s, // Invalid attribute sicne Type is not supported in v0.28.0
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
#### xxxOptions struct example below for Context config

### CtxOptions struct
CtxOptions is the super set of parameters for Context, for all Runtime Versions. Based on the Runtime Version CtxOptions attributes may change(mandatory/optional/Not applicable).
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
// Look below for each attirbute documentation for which version its been supported/mandatory/optional.
type CtxOptions struct {
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


## Sample test cases

``` go
// TestCase : Run Runtime V100 SetContext API and Runtime V0280 GetContext API
func TestSetContextV100AndGetContextV0280(){
	// Input Parameters for Runtime SetContext API
	setContextInputOptions := SetContextInputOptions{
	    RuntimeVersion: framework.Version100,
	    CtxOptions: CtxOptions{
	        Name: "context-one",
	        Target: framework.TargetK8s,
	        GlobalOpts: &framework.GlobalServer{
			    Endpoint: "test-endpoint",
		    }
	    }
	}
	
	// Output Parameters for Runtime SetContext API
	var setContextOutputOptions SetContextOutputOptions
	
	// Input Parameters for Runtime GetContext API
		getContextInputOptions := GetContextInputOptions{
	    RuntimeVersion: framework.Version100,
	    ContextName: "context-one",
	}
	
	// Output Parameters for Runtime GetContext API
		getContextOutputOptions := GetContextOutputOptions{
	    CtxOptions: CtxOptions{
	        Name: "context-one",
	        Target: framework.TargetK8s,
	        GlobalOpts: &framework.GlobalServer{
			    Endpoint: "test-endpoint",
		    }
	    }
	}
	
	// Create SetContext Command
	setContextCommand, err := framework.NewSetContextCommand(setContextInputOptions, setContextOutputOptions)
	Expect(err).To(BeNil())
	
	// Create GetContext Command
	getContextCommand, err := framework.NewGetContextCommand(getContextInputOptions, getContextOutputOptions)
	Expect(err).To(BeNil())

	// Construct series of commands to execute
	testSuite := framework.NewTestSuite(  // re-named from NewTestCommands
		framework.NewTestCase(	// renamed from WithCommands
			setContextCommand,
			getContextCommand,
		),
	)

	// Executes the commands from the list and validates the expected output with actual output and return err if output doesn't match
	errs := testSuite.Execute()
	Expect(err).To(BeNil())
}
```
