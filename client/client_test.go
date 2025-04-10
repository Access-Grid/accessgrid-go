package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name      string
		accountID string
		secretKey string
		wantErr   bool
	}{
		{
			name:      "Valid credentials",
			accountID: "test-account",
			secretKey: "test-secret",
			wantErr:   false,
		},
		{
			name:      "Missing accountID",
			accountID: "",
			secretKey: "test-secret",
			wantErr:   true,
		},
		{
			name:      "Missing secretKey",
			accountID: "test-account",
			secretKey: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.accountID, tt.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientOptions(t *testing.T) {
	accountID := "test-account"
	secretKey := "test-secret"
	customURL := "https://custom.api.example.com"
	customTimeout := 5 * time.Second

	client, err := NewClient(
		accountID,
		secretKey,
		WithBaseURL(customURL),
		WithHTTPClient(&http.Client{Timeout: customTimeout}),
	)

	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if client.BaseURL != customURL {
		t.Errorf("WithBaseURL() got = %v, want %v", client.BaseURL, customURL)
	}

	if client.HTTPClient.Timeout != customTimeout {
		t.Errorf("WithHTTPClient() got timeout = %v, want %v", client.HTTPClient.Timeout, customTimeout)
	}
}

func TestClientRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be application/json")
		}
		if r.Header.Get("X-AccessGrid-Account-ID") == "" {
			t.Errorf("Expected X-AccessGrid-Account-ID header to be set")
		}
		if r.Header.Get("X-AccessGrid-Timestamp") == "" {
			t.Errorf("Expected X-AccessGrid-Timestamp header to be set")
		}
		if r.Header.Get("X-AccessGrid-Signature") == "" {
			t.Errorf("Expected X-AccessGrid-Signature header to be set")
		}

		// Return a simple JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "response"}`))
	}))
	defer server.Close()

	// Create client with test server URL
	client, err := NewClient("test-account", "test-secret", WithBaseURL(server.URL))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Test request
	var response map[string]string
	err = client.Request("GET", "/test-path", nil, &response)
	if err != nil {
		t.Fatalf("client.Request() error = %v", err)
	}

	// Verify response
	if response["test"] != "response" {
		t.Errorf("Expected response[\"test\"] = %v, got %v", "response", response["test"])
	}
}