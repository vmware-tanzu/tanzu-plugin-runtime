# Tanzu Hub Client

This package enables the creation of authenticated Tanzu Hub clients which in turn
enables the interaction with Tanzu Hub endpoint through GraphQL queries for the
`tanzu` context.

**Note:** The Tanzu Hub Client feature is an EXPERIMENTAL feature. The API signature
and implementation are subjected to change/removal if an alternative means to provide
equivalent functionality can be introduced.

## Creating a Tanzu Hub Client

To create a Tanzu Hub client, use the `CreateHubClient(contextName string)` API
by providing the `tanzu` context name. An authenticated Tanzu Hub client for the specified tanzu context will be returned.
Internally it configures the client with an access token for each request.
By default, it will get the Tanzu Hub endpoint from the specified context metadata. To specify any custom Tanzu Hub
endpoint for testing please configure the `TANZU_HUB_ENDPOINT` environment variable.

Note that the authenticated client is assured to have at least 30 min access to the GraphQL endpoint.
If you want a long running client beyond this period, recommendation is to reinitialize your client.

## Generating the golang stub to invoke graphQL queries

There are many golang client libraries for the graphQL, however, Tanzu Plugin Runtime uses `github.com/Khan/genqlient` and
also returns the corresponding graphQL client as part of the Tanzu Hub client creation.

github.com/Khan/genqlient is a Go library to easily generate type-safe code to query a GraphQL API.

To help plugin authors generate the stub for the Tanzu Hub endpoint, a [tanzuhub.mk](../../hack/hub/tanzuhub.mk) has been provided.
This makefile provides an easy means for the plugin authors to initialize a `hub` package and also generate the stub from the graphQL queries.
To use this library plugin authors can follow the below steps:

1. Copy the [tanzuhub.mk](../../hack/hub/tanzuhub.mk) to your project and import it to your `Makefile` with `include ./tanzuhub.mk`
2. Configure the `TANZU_HUB_SCHEMA_FILE_URL` environment variable to the `schema.graphql` of the Tanzu Hub
3. Run `make tanzu-hub-stub-init` to initialize a `hub` package. This will create the following files under the `hub` package:
    * `genqlient.yaml`: Configuration file for generating golang code from GraphQL query with `github.com/Khan/genqlient`
    * `queries.graphql`: File to write all graphQL queries
    * `main.go`: A golang file with necessary imports to easily run `go generate` to generate stub code
4. Once the initialization is done, you can add your GraphQL queries to the `queries.graphql` file
5. After adding new graphQL queries or updating an existing query, run `make tanzu-hub-stub-generate` to generate a golang stub for the GraphQL queries
    * This will create a `generate.go` file under the `hub` package with golang APIs that can be consumed directly by other packages by passing the GraphQLClient available with TanzuHub client

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