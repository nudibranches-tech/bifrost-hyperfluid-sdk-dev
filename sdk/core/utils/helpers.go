package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Environment variables handling
func GetEnvironmentVariable(key string, fallback string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return fallback
}
func GetEnvironmentVariableInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value != "" {
		parsedValue, conversionError := strconv.Atoi(value)
		if conversionError == nil {
			return parsedValue
		}
	}
	return fallback
}

// HTTP client handling
func CreateHTTPClientWithSettings(skipTLSVerification bool, timeoutDuration time.Duration) *http.Client {
	transport := &http.Transport{}
	if skipTLSVerification {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return &http.Client{Transport: transport, Timeout: timeoutDuration}
}

// Error handling
func (response *Response) HasError() bool {
	return response != nil && response.Error != ""
}
func ResponseError(message string) (*Response, error) {
	return &Response{Status: StatusError, Error: message}, fmt.Errorf("%s", message)
}

// Response handling
func (response *Response) IsOK() bool {
	return response != nil && response.Status == StatusOK
}
func ResponseSuccess(data any) *Response {
	return &Response{Status: StatusOK, Data: data}
}
func JsonMarshal(value any) []byte {
	encodedBytes, _ := json.Marshal(value)
	return encodedBytes
}
func (response *Response) GetDataAsSlice() ([]any, bool) {
	sliceValue, isSlice := response.Data.([]any)
	return sliceValue, isSlice
}
func (response *Response) GetDataAsMap() (map[string]any, bool) {
	mapValue, isMap := response.Data.(map[string]any)
	return mapValue, isMap
}
