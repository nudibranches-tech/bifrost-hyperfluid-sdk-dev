package sdk

import (
	"bifrost-for-developers/sdk/utils"
	"net/http"
)

// Client is the main entry point for the SDK.
type Client struct {
	config     utils.Configuration
	httpClient *http.Client
}

// NewClient creates a new Bifrost client.
func NewClient(config utils.Configuration) *Client {
	// we create a copy of the configuration to avoid side effects
	cfg := config
	return &Client{
		config: cfg,
		httpClient: utils.CreateHTTPClientWithSettings(
			cfg.SkipTLSVerify,
			cfg.RequestTimeout,
		),
	}
}

// GetCatalog retrieves a catalog by name.
// This is the legacy API - consider using the fluent API instead:
//
//	client.Catalog("name").Schema("schema").Table("table").Get(ctx)
func (c *Client) GetCatalog(name string) *Catalog {
	return &Catalog{
		Name:   name,
		client: c,
	}
}

// Query creates a new QueryBuilder for fluent query construction.
// Example:
//
//	resp, err := client.Query().
//	    Catalog("sales").
//	    Schema("public").
//	    Table("orders").
//	    Limit(10).
//	    Get(ctx)
func (c *Client) Query() *QueryBuilder {
	return newQueryBuilder(c)
}

// Catalog starts a new fluent query with the catalog name.
// This is a shortcut for client.Query().Catalog(name).
func (c *Client) Catalog(name string) *QueryBuilder {
	return newQueryBuilder(c).Catalog(name)
}

// Org starts a new fluent query with a specific organization ID.
// This overrides the default OrgID from the client configuration.
func (c *Client) Org(orgID string) *QueryBuilder {
	return newQueryBuilder(c).Org(orgID)
}
