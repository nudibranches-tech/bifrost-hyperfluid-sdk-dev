package main

import (
	bifrost "bifrost-for-developers/sdk"
	"fmt"
	"strings"
)

func handleResult(result *bifrost.Result) {
	if result.Error != nil {
		fmt.Printf("❌ Error: %s\n", error.Error(result.Error))
		return
	}
	if !result.Response.IsOK() {
		fmt.Printf("❌ Error: %s\n", result.Response.Error)
		return
	}
	fmt.Println("✅ Success!")
	if dataSlice, isSlice := result.Response.GetDataAsSlice(); isSlice {
		fmt.Printf("📦 %d records", len(dataSlice))
		if len(dataSlice) > 0 {
			fmt.Printf(" | First: %v", dataSlice[0])
		}
		fmt.Println()
	} else if dataMap, isMap := result.Response.GetDataAsMap(); isMap {
		fmt.Printf("📦 Data: %v\n", dataMap)
	}
}

func runPostgresExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Example 1: PostgreSQL Query")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	bifrost.Init()
	var globalConfiguration = bifrost.GetGlobalConfiguration()

	sqlQuery := fmt.Sprintf("SELECT * FROM %s.%s.%s LIMIT 5", globalConfiguration.TestCatalog, globalConfiguration.TestSchema, globalConfiguration.TestTable)
	fmt.Printf("📝 %s\n", sqlQuery)

	result := <-bifrost.Request(bifrost.BifrostRequest{Type: bifrost.RequestPostgres, PostgresPayload: &bifrost.PostgresPayload{SQL: sqlQuery}})
	handleResult(&result)
	fmt.Println()
}

func runGraphQLExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Example 2: GraphQL Query")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	bifrost.Init()
	var globalConfiguration = bifrost.GetGlobalConfiguration()

	graphqlQuery := fmt.Sprintf("{ %s { %s { %s (limit: 5){ %s }}}}", globalConfiguration.TestCatalog, globalConfiguration.TestSchema, globalConfiguration.TestTable, strings.ReplaceAll(globalConfiguration.TestColumns, ",", " "))
	fmt.Printf("📝 %s\n", graphqlQuery)

	result := <-bifrost.Request(bifrost.BifrostRequest{Type: bifrost.RequestGraphQL, GraphQLPayload: &bifrost.GraphQLPayload{Query: graphqlQuery}})
	handleResult(&result)
	fmt.Println()
}

func runOpenAPIExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Example 3: OpenAPI (REST) Query")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	bifrost.Init()
	var globalConfiguration = bifrost.GetGlobalConfiguration()

	fmt.Printf("📝 GET /%s/%s/%s?_limit=10&select=%s\n", globalConfiguration.TestCatalog, globalConfiguration.TestSchema, globalConfiguration.TestTable, globalConfiguration.TestColumns)

	result := <-bifrost.Request(bifrost.BifrostRequest{Type: bifrost.RequestOpenAPI, OpenAPIPayload: &bifrost.OpenAPIPayload{
		Catalog: globalConfiguration.TestCatalog,
		Schema:  globalConfiguration.TestSchema,
		Table:   globalConfiguration.TestTable,
		Method:  "GET",
		Params:  map[string]string{"select": globalConfiguration.TestColumns},
	}})
	handleResult(&result)
	fmt.Println()
}
