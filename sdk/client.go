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
func (c *Client) GetCatalog(name string) *Catalog {
	return &Catalog{
		Name:   name,
		client: c,
	}
}
