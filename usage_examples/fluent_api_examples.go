package main

import (
	"bifrost-for-developers/sdk"
	"context"
	"fmt"
)

// This file demonstrates the new fluent API for the Bifrost SDK.
// The fluent API provides a more intuitive and user-friendly way to interact with the SDK.

func runFluentAPISimpleExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Fluent API Example 1: Simple Query")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	config := getConfig()
	client := sdk.NewClient(config)

	testCatalog := getEnv("BIFROST_TEST_CATALOG", "")
	testSchema := getEnv("BIFROST_TEST_SCHEMA", "")
	testTable := getEnv("BIFROST_TEST_TABLE", "")

	if testCatalog == "" || testSchema == "" || testTable == "" {
		fmt.Println("⚠️  Skipping: BIFROST_TEST_CATALOG, BIFROST_TEST_SCHEMA, or BIFROST_TEST_TABLE not set")
		fmt.Println()
		return
	}

	fmt.Printf("📝 Fluent query: client.Catalog(%q).Schema(%q).Table(%q).Limit(5).Get(ctx)\n",
		testCatalog, testSchema, testTable)

	// NEW FLUENT API - Simple and intuitive!
	resp, err := client.
		Catalog(testCatalog).
		Schema(testSchema).
		Table(testTable).
		Limit(5).
		Get(context.Background())

	handleResponse(resp, err)
	fmt.Println()
}

func runFluentAPIWithSelectExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Fluent API Example 2: Query with SELECT")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	config := getConfig()
	client := sdk.NewClient(config)

	testCatalog := getEnv("BIFROST_TEST_CATALOG", "")
	testSchema := getEnv("BIFROST_TEST_SCHEMA", "")
	testTable := getEnv("BIFROST_TEST_TABLE", "")
	testColumns := getEnv("BIFROST_TEST_COLUMNS", "")

	if testCatalog == "" || testSchema == "" || testTable == "" {
		fmt.Println("⚠️  Skipping: Test environment variables not set")
		fmt.Println()
		return
	}

	if testColumns == "" {
		fmt.Println("⚠️  Skipping: BIFROST_TEST_COLUMNS not set")
		fmt.Println()
		return
	}

	fmt.Printf("📝 Fluent query with SELECT: .Select(%q).Limit(10).Get(ctx)\n", testColumns)

	// Select specific columns (comma-separated string to variadic args)
	cols := splitColumns(testColumns)

	resp, err := client.
		Catalog(testCatalog).
		Schema(testSchema).
		Table(testTable).
		Select(cols...).
		Limit(10).
		Get(context.Background())

	handleResponse(resp, err)
	fmt.Println()
}

func runFluentAPIComplexExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Fluent API Example 3: Complex Query")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	config := getConfig()
	client := sdk.NewClient(config)

	testCatalog := getEnv("BIFROST_TEST_CATALOG", "")
	testSchema := getEnv("BIFROST_TEST_SCHEMA", "")
	testTable := getEnv("BIFROST_TEST_TABLE", "")

	if testCatalog == "" || testSchema == "" || testTable == "" {
		fmt.Println("⚠️  Skipping: Test environment variables not set")
		fmt.Println()
		return
	}

	fmt.Println("📝 Complex fluent query with:")
	fmt.Println("   - Multiple SELECT columns")
	fmt.Println("   - WHERE filters")
	fmt.Println("   - ORDER BY")
	fmt.Println("   - Pagination (LIMIT + OFFSET)")

	// Complex query with all features
	resp, err := client.
		Catalog(testCatalog).
		Schema(testSchema).
		Table(testTable).
		Select("id", "name", "created_at").
		Where("status", "=", "active").
		Where("amount", ">", 100).
		OrderBy("created_at", "DESC").
		Limit(20).
		Offset(0).
		Get(context.Background())

	handleResponse(resp, err)
	fmt.Println()
}

func runFluentAPICustomOrgExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Fluent API Example 4: Custom Org ID")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	config := getConfig()
	client := sdk.NewClient(config)

	customOrgID := getEnv("HYPERFLUID_CUSTOM_ORG_ID", "")
	if customOrgID == "" {
		fmt.Println("⚠️  Using default org from config")
		customOrgID = config.OrgID
	}

	testCatalog := getEnv("BIFROST_TEST_CATALOG", "")
	testSchema := getEnv("BIFROST_TEST_SCHEMA", "")
	testTable := getEnv("BIFROST_TEST_TABLE", "")

	if testCatalog == "" || testSchema == "" || testTable == "" {
		fmt.Println("⚠️  Skipping: Test environment variables not set")
		fmt.Println()
		return
	}

	fmt.Printf("📝 Query with custom org ID: .Org(%q)\n", customOrgID)

	// Override org ID for this specific query
	resp, err := client.
		Org(customOrgID).
		Catalog(testCatalog).
		Schema(testSchema).
		Table(testTable).
		Limit(5).
		Get(context.Background())

	handleResponse(resp, err)
	fmt.Println()
}

func runFluentAPIMultipleChainsExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Fluent API Example 5: Multiple Chained Calls")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	config := getConfig()
	client := sdk.NewClient(config)

	testCatalog := getEnv("BIFROST_TEST_CATALOG", "")
	testSchema := getEnv("BIFROST_TEST_SCHEMA", "")
	testTable := getEnv("BIFROST_TEST_TABLE", "")

	if testCatalog == "" || testSchema == "" || testTable == "" {
		fmt.Println("⚠️  Skipping: Test environment variables not set")
		fmt.Println()
		return
	}

	fmt.Println("📝 Building query step by step:")

	// You can also build the query in steps
	query := client.
		Catalog(testCatalog).
		Schema(testSchema).
		Table(testTable)

	fmt.Println("   1. Base query created")

	// Add select columns
	query = query.Select("id", "name")
	fmt.Println("   2. Added SELECT columns")

	// Add filters
	query = query.Where("status", "=", "active")
	fmt.Println("   3. Added WHERE filter")

	// Add ordering
	query = query.OrderBy("id", "ASC")
	fmt.Println("   4. Added ORDER BY")

	// Add pagination
	query = query.Limit(10)
	fmt.Println("   5. Added LIMIT")

	// Execute
	fmt.Println("   6. Executing query...")
	resp, err := query.Get(context.Background())

	handleResponse(resp, err)
	fmt.Println()
}

func runFluentAPIComparisonExample() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎯 Fluent API Example 6: Old vs New API Comparison")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	config := getConfig()
	client := sdk.NewClient(config)

	testCatalog := getEnv("BIFROST_TEST_CATALOG", "")
	testSchema := getEnv("BIFROST_TEST_SCHEMA", "")
	testTable := getEnv("BIFROST_TEST_TABLE", "")

	if testCatalog == "" || testSchema == "" || testTable == "" {
		fmt.Println("⚠️  Skipping: Test environment variables not set")
		fmt.Println()
		return
	}

	fmt.Println("📝 OLD API (still supported):")
	fmt.Println("   catalog := client.GetCatalog(\"catalog\")")
	fmt.Println("   table := catalog.Table(\"schema\", \"table\")")
	fmt.Println("   params := url.Values{}")
	fmt.Println("   params.Add(\"_limit\", \"10\")")
	fmt.Println("   resp, err := table.GetData(ctx, params)")
	fmt.Println()

	fmt.Println("📝 NEW FLUENT API (recommended):")
	fmt.Println("   resp, err := client.")
	fmt.Println("       Catalog(\"catalog\").")
	fmt.Println("       Schema(\"schema\").")
	fmt.Println("       Table(\"table\").")
	fmt.Println("       Limit(10).")
	fmt.Println("       Get(ctx)")
	fmt.Println()

	fmt.Println("🚀 Running new fluent API...")

	resp, err := client.
		Catalog(testCatalog).
		Schema(testSchema).
		Table(testTable).
		Limit(10).
		Get(context.Background())

	handleResponse(resp, err)
	fmt.Println()
}

// Helper functions

func splitColumns(cols string) []string {
	if cols == "" {
		return []string{}
	}
	// Simple split by comma
	var result []string
	current := ""
	for _, c := range cols {
		if c == ',' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else if c != ' ' {
			current += string(c)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
