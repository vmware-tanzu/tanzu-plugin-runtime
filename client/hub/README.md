# Tanzu Hub Client

This package enables the creation of authenticated Tanzu Hub clients which in turn
enables the interaction with Tanzu Hub endpoint through GraphQL queries for the
`tanzu` context.

**Note:** The Tanzu Hub Client feature is an EXPERIMENTAL feature. The API signature
and implementation are subjected to change/removal if an alternative means to provide
equivalent functionality can be introduced.

## Creating a Tanzu Hub Client

To create a Tanzu Hub client, use the `NewClient(contextName string)` API
by providing the `tanzu` context name. An authenticated Tanzu Hub client for the specified tanzu context will be returned.
Internally it configures the client with an access token for each request.
By default, it will get the Tanzu Hub endpoint from the specified context metadata. To specify any custom Tanzu Hub
endpoint for testing please configure the `TANZU_HUB_ENDPOINT` environment variable.

Note that the authenticated client is assured to have at least 30 min access to the GraphQL endpoint.
If you want a long running client beyond this period, recommendation is to reinitialize your client.

## Examples

### Query/Mutation

```golang
const QueryAllProjects_Operation = `
query QueryAllProjects {
    applicationEngineQuery {
        queryProjects(first: 1000) {
            projects {
                name
            }
        }
    }
}`

// getProjects is a wrapper of an `QueryAllProjects“ API call to fetch project names
func getProjects(contextName string) ([]string, error) {
    hc, err := NewClient(contextName )

    req := &hub.Request{
        OpName: "QueryAllProjects",
        Query:  QueryAllProjects_Operation,
    }
    var responseData QueryAllProjectsResponse // Assuming the response type is already defined
    err := hc.Request(context.Background(), req, &responseData)
    if err != nil {
        return nil, err
    }

    // Process the response
    projects := []string{}
    for _, p := range responseData.ApplicationEngineQuery.QueryProjects.Projects {
        projects = append(projects, p.Name)
    }

    return projects, nil
}
```

### Subscriptions

```golang
const SubscribeAppLogs_Operation = `
subscription appLogs($appEntityId: EntityId!) {
  kubernetesAppLogs(appEntityId: $appEntityId, logParams: {includeTimestamps: true, tailLines: 50, includePrevious: false}) {
    value
    timestamp
  }
}`

func subscribeAppLogs(contextName, appEntityId string) ([]string, error) {
    hc, err := NewClient(contextName )

    req := &hub.Request{
        OpName: "SubscribeAppLogs",
        Query:  SubscribeAppLogs_Operation,
        Variables: map[string]string{"appEntityId": appEntityId}
    }

    err := hc.Subscribe(context.Background(), req, logEventHandler)
    if err != nil {
        return nil, err
    }

    return nil
}

func logEventHandler(eventResponse EventResponse) {
    respData := eventResponse.Data
    fmt.Println(respData)
}
```

## Generating the golang stub for the GraphQL queries

There are many golang client libraries for the GraphQL, however, Tanzu Plugin Runtime recommends
using `github.com/Khan/genqlient` for generating type definitions for your GraphQL queries.

Note: Client does not require users to use this library for generating type definitions and users can always define the Golang type definitions
for the specified query by themselves.
