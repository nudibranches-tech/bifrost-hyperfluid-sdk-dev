package sdk

import (
	"strings"
	"testing"
)

func TestLoadServiceAccountFromJSON(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		wantErr     bool
		errContains string
		checkFunc   func(t *testing.T, sa *ServiceAccount)
	}{
		{
			name: "valid service account",
			json: `{
				"client_id": "hf-org-sa-12345",
				"client_secret": "secret123",
				"issuer": "https://auth.hyperfluid.cloud/realms/my-org",
				"auth_uri": "https://auth.hyperfluid.cloud/realms/my-org/protocol/openid-connect/auth",
				"token_uri": "https://auth.hyperfluid.cloud/realms/my-org/protocol/openid-connect/token"
			}`,
			wantErr: false,
			checkFunc: func(t *testing.T, sa *ServiceAccount) {
				if sa.ClientID != "hf-org-sa-12345" {
					t.Errorf("ClientID = %q, want %q", sa.ClientID, "hf-org-sa-12345")
				}
				if sa.ClientSecret != "secret123" {
					t.Errorf("ClientSecret = %q, want %q", sa.ClientSecret, "secret123")
				}
				if sa.Issuer != "https://auth.hyperfluid.cloud/realms/my-org" {
					t.Errorf("Issuer = %q, want %q", sa.Issuer, "https://auth.hyperfluid.cloud/realms/my-org")
				}
			},
		},
		{
			name: "valid with token_uri only (no issuer)",
			json: `{
				"client_id": "hf-org-sa-12345",
				"client_secret": "secret123",
				"token_uri": "https://auth.hyperfluid.cloud/realms/my-org/protocol/openid-connect/token"
			}`,
			wantErr: false,
		},
		{
			name:        "missing client_id",
			json:        `{"client_secret": "secret123", "issuer": "https://auth.hyperfluid.cloud/realms/my-org"}`,
			wantErr:     true,
			errContains: "client_id is required",
		},
		{
			name:        "missing client_secret",
			json:        `{"client_id": "hf-org-sa-12345", "issuer": "https://auth.hyperfluid.cloud/realms/my-org"}`,
			wantErr:     true,
			errContains: "client_secret is required",
		},
		{
			name:        "missing both issuer and token_uri",
			json:        `{"client_id": "hf-org-sa-12345", "client_secret": "secret123"}`,
			wantErr:     true,
			errContains: "either issuer or token_uri is required",
		},
		{
			name:        "invalid json",
			json:        `{"client_id": "hf-org-sa-12345"`,
			wantErr:     true,
			errContains: "failed to parse service account JSON",
		},
		{
			name:        "empty json",
			json:        `{}`,
			wantErr:     true,
			errContains: "client_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa, err := LoadServiceAccountFromJSON(tt.json)
			if tt.wantErr {
				if err == nil {
					t.Errorf("LoadServiceAccountFromJSON() error = nil, want error containing %q", tt.errContains)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("LoadServiceAccountFromJSON() error = %q, want error containing %q", err.Error(), tt.errContains)
				}
				return
			}
			if err != nil {
				t.Errorf("LoadServiceAccountFromJSON() unexpected error = %v", err)
				return
			}
			if tt.checkFunc != nil {
				tt.checkFunc(t, sa)
			}
		})
	}
}

func TestServiceAccount_ParseIssuer(t *testing.T) {
	tests := []struct {
		name        string
		issuer      string
		tokenURI    string
		wantBaseURL string
		wantRealm   string
		wantErr     bool
		errContains string
	}{
		{
			name:        "standard issuer URL",
			issuer:      "https://auth.hyperfluid.cloud/realms/my-org",
			wantBaseURL: "https://auth.hyperfluid.cloud",
			wantRealm:   "my-org",
			wantErr:     false,
		},
		{
			name:        "issuer with trailing slash",
			issuer:      "https://auth.hyperfluid.cloud/realms/my-org/",
			wantBaseURL: "https://auth.hyperfluid.cloud",
			wantRealm:   "my-org",
			wantErr:     false,
		},
		{
			name:        "issuer with complex realm name",
			issuer:      "https://auth.hyperfluid.cloud/realms/nudibranches-tech",
			wantBaseURL: "https://auth.hyperfluid.cloud",
			wantRealm:   "nudibranches-tech",
			wantErr:     false,
		},
		{
			name:        "fallback to token_uri when issuer empty",
			issuer:      "",
			tokenURI:    "https://auth.hyperfluid.cloud/realms/fallback-org/protocol/openid-connect/token",
			wantBaseURL: "https://auth.hyperfluid.cloud",
			wantRealm:   "fallback-org",
			wantErr:     false,
		},
		{
			name:        "invalid issuer URL - no realms path",
			issuer:      "https://auth.hyperfluid.cloud/other/path",
			wantErr:     true,
			errContains: "does not contain /realms/<realm> pattern",
		},
		{
			name:        "both empty",
			issuer:      "",
			tokenURI:    "",
			wantErr:     true,
			errContains: "issuer is empty and no token_uri available",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := &ServiceAccount{
				ClientID:     "test-client",
				ClientSecret: "test-secret",
				Issuer:       tt.issuer,
				TokenURI:     tt.tokenURI,
			}
			baseURL, realm, err := sa.ParseIssuer()
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseIssuer() error = nil, want error")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("ParseIssuer() error = %q, want error containing %q", err.Error(), tt.errContains)
				}
				return
			}
			if err != nil {
				t.Errorf("ParseIssuer() unexpected error = %v", err)
				return
			}
			if baseURL != tt.wantBaseURL {
				t.Errorf("ParseIssuer() baseURL = %q, want %q", baseURL, tt.wantBaseURL)
			}
			if realm != tt.wantRealm {
				t.Errorf("ParseIssuer() realm = %q, want %q", realm, tt.wantRealm)
			}
		})
	}
}

func TestServiceAccount_ToConfiguration(t *testing.T) {
	sa := &ServiceAccount{
		ClientID:     "hf-org-sa-12345",
		ClientSecret: "secret123",
		Issuer:       "https://auth.hyperfluid.cloud/realms/my-org",
		AuthURI:      "https://auth.hyperfluid.cloud/realms/my-org/protocol/openid-connect/auth",
		TokenURI:     "https://auth.hyperfluid.cloud/realms/my-org/protocol/openid-connect/token",
	}

	opts := ServiceAccountOptions{
		BaseURL:        "https://api.hyperfluid.cloud",
		OrgID:          "org-123",
		DataDockID:     "dd-456",
		SkipTLSVerify:  false,
		RequestTimeout: 60,
		MaxRetries:     5,
	}

	cfg, err := sa.ToConfiguration(opts)
	if err != nil {
		t.Fatalf("ToConfiguration() unexpected error = %v", err)
	}

	// Verify all fields are correctly mapped
	if cfg.BaseURL != "https://api.hyperfluid.cloud" {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, "https://api.hyperfluid.cloud")
	}
	if cfg.OrgID != "org-123" {
		t.Errorf("OrgID = %q, want %q", cfg.OrgID, "org-123")
	}
	if cfg.DataDockID != "dd-456" {
		t.Errorf("DataDockID = %q, want %q", cfg.DataDockID, "dd-456")
	}
	if cfg.KeycloakBaseURL != "https://auth.hyperfluid.cloud" {
		t.Errorf("KeycloakBaseURL = %q, want %q", cfg.KeycloakBaseURL, "https://auth.hyperfluid.cloud")
	}
	if cfg.KeycloakRealm != "my-org" {
		t.Errorf("KeycloakRealm = %q, want %q", cfg.KeycloakRealm, "my-org")
	}
	if cfg.KeycloakClientID != "hf-org-sa-12345" {
		t.Errorf("KeycloakClientID = %q, want %q", cfg.KeycloakClientID, "hf-org-sa-12345")
	}
	if cfg.KeycloakClientSecret != "secret123" {
		t.Errorf("KeycloakClientSecret = %q, want %q", cfg.KeycloakClientSecret, "secret123")
	}
	if cfg.MaxRetries != 5 {
		t.Errorf("MaxRetries = %d, want %d", cfg.MaxRetries, 5)
	}
}

func TestNewClientFromServiceAccount(t *testing.T) {
	tests := []struct {
		name        string
		sa          *ServiceAccount
		opts        ServiceAccountOptions
		wantErr     bool
		errContains string
	}{
		{
			name: "valid service account",
			sa: &ServiceAccount{
				ClientID:     "hf-org-sa-12345",
				ClientSecret: "secret123",
				Issuer:       "https://auth.hyperfluid.cloud/realms/my-org",
			},
			opts: ServiceAccountOptions{
				BaseURL: "https://api.hyperfluid.cloud",
			},
			wantErr: false,
		},
		{
			name:        "nil service account",
			sa:          nil,
			opts:        ServiceAccountOptions{BaseURL: "https://api.hyperfluid.cloud"},
			wantErr:     true,
			errContains: "service account is nil",
		},
		{
			name: "missing BaseURL",
			sa: &ServiceAccount{
				ClientID:     "hf-org-sa-12345",
				ClientSecret: "secret123",
				Issuer:       "https://auth.hyperfluid.cloud/realms/my-org",
			},
			opts:        ServiceAccountOptions{},
			wantErr:     true,
			errContains: "BaseURL is required",
		},
		{
			name: "invalid issuer",
			sa: &ServiceAccount{
				ClientID:     "hf-org-sa-12345",
				ClientSecret: "secret123",
				Issuer:       "https://invalid-url.com/no-realms",
			},
			opts: ServiceAccountOptions{
				BaseURL: "https://api.hyperfluid.cloud",
			},
			wantErr:     true,
			errContains: "failed to create configuration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClientFromServiceAccount(tt.sa, tt.opts)
			if tt.wantErr {
				if err == nil {
					t.Errorf("NewClientFromServiceAccount() error = nil, want error")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("NewClientFromServiceAccount() error = %q, want error containing %q", err.Error(), tt.errContains)
				}
				return
			}
			if err != nil {
				t.Errorf("NewClientFromServiceAccount() unexpected error = %v", err)
				return
			}
			if client == nil {
				t.Error("NewClientFromServiceAccount() returned nil client")
			}
		})
	}
}

func TestLoadServiceAccountFromReader(t *testing.T) {
	json := `{
		"client_id": "hf-org-sa-reader-test",
		"client_secret": "reader-secret",
		"issuer": "https://auth.hyperfluid.cloud/realms/reader-test"
	}`

	reader := strings.NewReader(json)
	sa, err := LoadServiceAccountFromReader(reader)
	if err != nil {
		t.Fatalf("LoadServiceAccountFromReader() unexpected error = %v", err)
	}

	if sa.ClientID != "hf-org-sa-reader-test" {
		t.Errorf("ClientID = %q, want %q", sa.ClientID, "hf-org-sa-reader-test")
	}
}
