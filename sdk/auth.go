package sdk

import (
	"bifrost-for-developers/sdk/utils"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) isKeycloakConfigured() bool {
	return c.config.KeycloakUsername != "" && c.config.KeycloakPassword != ""
}

func (c *Client) refreshToken(ctx context.Context) (string, error) {
	if !c.isKeycloakConfigured() {
		return "", utils.ErrInvalidConfiguration
	}

	form := url.Values{
		"grant_type": {"password"},
		"client_id":  {c.config.KeycloakClientID},
		"username":   {c.config.KeycloakUsername},
		"password":   {c.config.KeycloakPassword},
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", c.config.KeycloakBaseURL, c.config.KeycloakRealm),
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", fmt.Errorf("%w: cannot create keycloak request: %w", utils.ErrInvalidRequest, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a new HTTP client for the token refresh to avoid using the client's timeout settings
	client := &http.Client{}
	if c.config.SkipTLSVerify {
		client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: cannot reach keycloak: %w", utils.ErrAuthenticationFailed, err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: keycloak auth failed (%d): %s", utils.ErrAuthenticationFailed, resp.StatusCode, body)
	}

	var parsed map[string]any
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", fmt.Errorf("%w: invalid keycloak response: %w", utils.ErrAuthenticationFailed, err)
	}
	token, ok := parsed["access_token"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("%w: missing access_token in keycloak response", utils.ErrAuthenticationFailed)
	}

	c.config.Token = token
	return token, nil
}
