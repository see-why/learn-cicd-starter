package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "test-api-key" {
		t.Errorf("Expected API key 'test-api-key', got '%s'", apiKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKeyWithSpecialCharacters(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key@#$^&*()_+")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "test-api-key@#$%^&*()_+" {
		t.Errorf("Expected API key 'test-api-key@#$^&*()_+', got '%q'", apiKey)
	}
}
