package developpementtests

import (
	"bifrost-for-developers/sdk/utils"
	"os"
	"strconv"
	"time"
)

func getTestConfig() (utils.Configuration, error) {
	orgID := os.Getenv("HYPERFLUID_ORG_ID")
	token := os.Getenv("HYPERFLUID_TOKEN")

	config := utils.Configuration{
		BaseURL:        os.Getenv("HYPERFLUID_BASE_URL"),
		OrgID:          orgID,
		Token:          token,
		RequestTimeout: time.Duration(getEnvInt("HYPERFLUID_REQUEST_TIMEOUT", 30)) * time.Second,
	}

	if config.Token == "" {
		config.KeycloakUsername = os.Getenv("KEYCLOAK_USERNAME")
		config.KeycloakPassword = os.Getenv("KEYCLOAK_PASSWORD")
		config.KeycloakBaseURL = os.Getenv("KEYCLOAK_BASE_URL")
		config.KeycloakClientID = os.Getenv("KEYCLOAK_CLIENT_ID")
		config.KeycloakRealm = os.Getenv("KEYCLOAK_REALM")
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}
