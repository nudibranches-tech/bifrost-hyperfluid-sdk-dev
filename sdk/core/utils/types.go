package utils

import (
	"errors"
	"time"
)

type Configuration struct {
	BaseURL string
	OrgID   string
	Token   string

	SkipTLSVerify  bool
	RequestTimeout time.Duration
	MaxRetries     int

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresDatabase string

	KeycloakBaseURL  string
	KeycloakRealm    string
	KeycloakClientID string
	KeycloakUsername string
	KeycloakPassword string

	TestCatalog string
	TestSchema  string
	TestTable   string
	TestColumns string
}

type RequestType string

const (
	RequestGraphQL  RequestType = "graphql"
	RequestOpenAPI  RequestType = "openapi"
	RequestPostgres RequestType = "postgres"
)

type BifrostRequest struct {
	Type            RequestType
	GraphQLPayload  *GraphQLPayload
	OpenAPIPayload  *OpenAPIPayload
	PostgresPayload *PostgresPayload
}

type GraphQLPayload struct {
	Query     string
	Variables map[string]any
}

type OpenAPIPayload struct {
	Catalog string
	Schema  string
	Table   string
	Method  string
	Params  map[string]string
}

type PostgresPayload struct {
	SQL string
}

type Response struct {
	Status   string
	Data     any
	Error    string
	HTTPCode int
}

type Result struct {
	Response *Response
	Error    error
}

const (
	StatusOK    = "ok"
	StatusError = "error"
)

var (
	ErrorUninitialized          = errors.New("SDK is not initialized, call InitFromEnv() first")
	ErrorInvalidRequest         = errors.New("invalid request payload")
	ErrorUnsupportedRequestType = errors.New("unsupported request type")
)
