# Tanzu Hub Client

This package enables the creation of authenticated Tanzu Hub clients which in turn
enables the interaction with Tanzu Hub endpoint through GraphQL queries for the
`tanzu` context.

## Creating a Tanzu Hub Client

To create a Tanzu Hub client, use the `CreateHubClient(contextName string)` API
by providing the `tanzu` context name. An authenticated Tanzu Hub client for the specified tanzu context will be returned.
This client includes an authenticated GraphQLClient from the `github.com/Khan/genqlient` package
that can be used to do GraphQL queries. Internally it configures the client with an access token for each request.
By default, it will get the Tanzu Hub endpoint from the specified context metadata. To specify any custom Tanzu Hub
endpoint for testing please configure the `TANZU_HUB_GRAPHQL_ENDPOINT` environment variable.

Note that the authenticated client is assured to have at least 30 min access to the GraphQL endpoint.
If you want a long running client beyond this period, recommendation is to reinitialize your client.

## Generating the golang stub to invoke graphQL queries

There are many golang client libraries for the graphQL, however, Tanzu Plugin Runtime uses `github.com/Khan/genqlient` and
also returns the corresponding graphQL client as part of the Tanzu Hub client creation.

github.com/Khan/genqlient is a Go library to easily generate type-safe code to query a GraphQL API.

To help plugin authors generate the stub for the Tanzu Hub endpoint, a [tanzuhub.mk](../../hack/hub/tanzuhub.mk) has been provided.
This makefile provides an easy means for the plugin authors to initialize a `hub` package and also generate the stub from the graphQL queries.
To use this library plugin authors can follow below steps:

1. Import [tanzuhub.mk](../../hack/hub/tanzuhub.mk) to your project's `Makefile`
2. Configure the `TANZU_HUB_SCHEMA_FILE_URL` environment variable to the `schema.graphql` of the Tanzu Hub
3. Run `make tanzu-hub-stub-init` to initialize a `hub` package. This will create the following files under the `hub` package:
    * `genqlient.yaml`: Configuration file for generating golang code from GraphQL query with `github.com/Khan/genqlient`
    * `queries.graphql`: File to write all graphQL queries
    * `main.go`: A golang file with necessary imports to easily run `go generate` to generate stub code
4. Once the initialization is done, you can add your GraphQL queries to the `queries.graphql` file
5. After adding new graphQL queries or updating an existing queries, run `make tanzu-hub-stub-generate` to generate a golang stub for the GraphQL queries
    * This will create a `generate.go` file under the `hub` package with golang APIs that can be consumed directly by other packages by passing the GraphQLClient available with TanzuHub client
