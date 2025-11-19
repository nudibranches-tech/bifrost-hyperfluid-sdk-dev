package developpementtests

import (
	"bifrost-for-developers/sdk"
	"context"
	"net/url"
	"os"
	"testing"
)

func TestIntegration_GetData(t *testing.T) {
	if testing.Short() {
		t.Skip("⏭️  Skipping integration test in short mode")
	}

	config, err := getTestConfig()
	if err != nil {
		t.Fatalf("Failed to get test config: %v", err)
	}

	testCatalog := os.Getenv("BIFROST_TEST_CATALOG")
	testSchema := os.Getenv("BIFROST_TEST_SCHEMA")
	testTable := os.Getenv("BIFROST_TEST_TABLE")

	if testCatalog == "" || testSchema == "" || testTable == "" {
		t.Skip("⏭️  Skipping integration test because BIFROST_TEST_CATALOG, BIFROST_TEST_SCHEMA or BIFROST_TEST_TABLE are not set")
	}

	client := sdk.NewClient(config)
	table := client.GetCatalog(testCatalog).Table(testSchema, testTable)

	params := url.Values{}
	params.Add("_limit", "1")

	resp, err := table.GetData(context.Background(), params)
	if err != nil {
		t.Fatalf("GetData failed: %v", err)
	}

	if resp.Status != "ok" {
		t.Fatalf("Expected status 'ok', got '%s'. Error: %s", resp.Status, resp.Error)
	}

	data, ok := resp.Data.([]interface{})
	if !ok {
		t.Fatalf("Expected response data to be a slice, got %T", resp.Data)
	}

	if len(data) != 1 {
		t.Errorf("Expected 1 row, got %d", len(data))
	}

	t.Log("✅ Successfully retrieved data from the API")
}
