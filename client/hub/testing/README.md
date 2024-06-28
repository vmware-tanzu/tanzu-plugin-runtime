# Mock TanzuHub Server for Writing Unit Tests

This package allows plugin authors to easily write unit tests using a mock GraphQL server endpoint for plugins that interact with the Tanzu Hub endpoint.

**Note:** This feature is EXPERIMENTAL. The API signature and implementation are subject to change or removal if an alternative means of providing equivalent functionality is introduced.

## Using the Mock Server to Write Unit Tests

To mock a response for GraphQL queries, plugin authors can follow these simple steps:

1. Start the mock server:

```go
import hubtesting "github.com/vmware-tanzu/tanzu-plugin-runtime/client/hub/testing"

hubMockServer := hubtesting.NewServer(t, WithQuery(...))
defer hubMockServer.Close()
```

1. Create a `mockHubClient` with the newly started server:

```go
import "github.com/vmware-tanzu/tanzu-plugin-runtime/client/hub"

mockHubClient, err = hub.CreateHubClient(c.CurrentContext, hub.WithEndpoint(hubMockServer.URL), hub.WithAccessToken("fake-token"))
```

1. To reuse a server across multiple unit tests, use `hubMockServer.Reset()` to clean up any previously registered queries, mutations, or errors.

1. To register mock queries or mutations, use the `hubMockServer.RegisterQuery` or `hubMockServer.RegisterMutation` APIs, which accept Operation objects.

## Defining an `Operation` Object to Mock a Response

Let's assume your GraphQL query looks like the following:

```gql
query QueryAllProjects {
  applicationEngineQuery {
    queryProjects(first: 1000) {
      projects {
        json
      }
    }
  }
}
```

The following Operation type defines what values should be provided to the object:

```go
// Operation is a general type that encompasses the Operation type and Response which
// is of the same type, but with data.
type Operation struct {
    // opType denotes whether the operation is a query or a mutation, using the opQuery
    // and opMutation constants. This is unexported as it is set by the *Server.RegisterQuery
    // and *Server.RegisterMutation functions, respectively.
    opType int

    // Identifier helps identify the operation in a request when coming through the Server.
    // For example, if your operation looks like this:
    //
    //   query {
    //     myOperation(foo: $foo) {
    //       fieldOne
    //        fieldTwo
    //     }
    //  }
    //
    // Then this field should be set to myOperation. It can also be more specific, a simple
    // strings.Contains check occurs to match operations. A more specific example of a
    // valid Identifier for the same operation given above would be myOperation(foo: $foo).
    Identifier string

    // Variables represents the map of variables that should be passed along with the
    // operation whenever it is invoked on the Server.
    Variables map[string]interface{}

    // Response represents the response that should be returned whenever the server makes
    // a match on Operation.opType, Operation.Name, and Operation.Variables.
    Response interface{}
}
```

If you have already generated the Go stubs using make tanzu-hub-stub-generate, all the Go type definitions necessary for writing the sample response should be available to you. These definitions might look like this based on the defined schema:

```go
// QueryAllProjectsApplicationEngineQuery includes the requested fields of the GraphQL type ApplicationEngineQuery.
type QueryAllProjectsApplicationEngineQuery struct {
    QueryProjects QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection `json:"queryProjects"`
}

// QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection includes the requested fields of the GraphQL type KubernetesKindProjectConnection.
type QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection struct {
    Projects []QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject `json:"projects"`
}

// QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject includes the requested fields of the GraphQL type KubernetesKindProject.
type QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject struct {
    // Raw JSON response as produced by KRM API
    Json json.RawMessage `json:"json"`
}
```

To create a testing.Operation object, we will reuse these generated type definitions to create a mock response object:

```go
import hubtesting "github.com/vmware-tanzu/tanzu-plugin-runtime/client/hub/testing"

type QueryResponse struct {
    ApplicationEngineQuery hub.QueryAllProjectsApplicationEngineQuery
}

mockResponse := hubtesting.Operation{
    Identifier: "QueryAllProjects",
    Response: QueryResponse{
        hub.QueryAllProjectsApplicationEngineQuery{
            QueryProjects: hub.QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnection{
                Projects: []hub.QueryAllProjectsApplicationEngineQueryQueryProjectsKubernetesKindProjectConnectionProjectsKubernetesKindProject{
                    {
                        Json: jsonBytesForProject1,
                    },
                    {
                        Json: jsonBytesForProject2,
                    },
                },
            },
        },
    },
}

hubMockServer.RegisterQuery(mockResponse)
```
