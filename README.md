# Bifrost SDK

Go SDK for Hyperfluid data access.

## Quick Start

```bash
# Install
go get bifrost-for-developers/sdk
```

## Usage

```go
import (
    "fmt"
    "net/url"
    "bifrost-for-developers/sdk"
    "bifrost-for-developers/sdk/utils"
)

func main() {
    // Configure the client
    config := utils.Configuration{
        BaseURL: "https://bifrost.hyperfluid.cloud",
        OrgID:   "your-org-id",
        Token:   "your-token",
        // or use Keycloak for token management
    }

    // Create a new client
    client := sdk.NewClient(config)

    // Get a catalog
    catalog := client.GetCatalog("my_catalog")

    // Get a table from the catalog
    table := catalog.Table("my_schema", "my_table")

    // Prepare query parameters
    params := url.Values{}
    params.Add("_limit", "10")
    params.Add("select", "col1,col2")

    // Get data from the table
    response, err := table.GetData(params)
    if err != nil {
        // Handle error
    }

    fmt.Println(response.Data)
}
```

## Configuration

### Required
- `HYPERFLUID_ORG_ID` - Your organization ID
- `HYPERFLUID_TOKEN` - API token (or use Keycloak)

### Optional
- `HYPERFLUID_BASE_URL` - API endpoint (default: `https://bifrost.hyperfluid.cloud`)

### Keycloak (alternative to token)
- `KEYCLOAK_USERNAME` - Your username
- `KEYCLOAK_PASSWORD` - Your password
- `KEYCLOAK_BASE_URL` - Keycloak server
- `KEYCLOAK_CLIENT_ID` - Client ID
- `KEYCLOAK_REALM` - Realm name

## Project Structure

```
sdk/
  client.go     # Client object and public API
  domain.go     # Domain objects (Catalog, Table)
  utils/        # Utility functions and types
```

## Error Handling

```go
response, err := table.GetData(params)
if err != nil {
    // Handle request error (e.g., network error, authentication error)
    log.Fatalf("Request failed: %v", err)
}

if response.Status != utils.StatusOK {
    // Handle API error (e.g., table not found)
    log.Printf("API error: %s", response.Error)
}
```

## License

Private SDK for internal use.
