package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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
		if r.Header.Get("X-ACCT-ID") == "" {
			t.Errorf("Expected X-ACCT-ID header to be set")
		}
		if r.Header.Get("X-PAYLOAD-SIG") == "" {
			t.Errorf("Expected X-PAYLOAD-SIG header to be set")
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
	ctx := context.Background()
	err = client.Request(ctx, "GET", "/test-path", nil, &response)
	if err != nil {
		t.Fatalf("client.Request() error = %v", err)
	}

	// Verify response
	if response["test"] != "response" {
		t.Errorf("Expected response[\"test\"] = %v, got %v", "response", response["test"])
	}
}

func TestClientRequest_ErrorResponses(t *testing.T) {
	tests := []struct {
		name          string
		statusCode    int
		contentType   string
		responseBody  string
		requestID     string // X-Request-ID header value
		wantMessage   string
		wantRequestID string
		wantRawBody   string
	}{
		{
			name:         "401 with message field",
			statusCode:   401,
			contentType:  "application/json",
			responseBody: `{"message":"Invalid credentials"}`,
			wantMessage:  "Invalid credentials",
			wantRawBody:  `{"message":"Invalid credentials"}`,
		},
		{
			name:         "404 with message field",
			statusCode:   404,
			contentType:  "application/json",
			responseBody: `{"message":"Resource not found"}`,
			wantMessage:  "Resource not found",
		},
		{
			name:         "422 validation error",
			statusCode:   422,
			contentType:  "application/json",
			responseBody: `{"message":"Invalid card_template_id"}`,
			wantMessage:  "Invalid card_template_id",
		},
		{
			name:         "500 with message field",
			statusCode:   500,
			contentType:  "application/json",
			responseBody: `{"message":"Internal server error"}`,
			wantMessage:  "Internal server error",
		},
		{
			name:         "500 with error field instead of message",
			statusCode:   500,
			contentType:  "application/json",
			responseBody: `{"error":"Something went wrong"}`,
			wantMessage:  "Something went wrong",
		},
		{
			name:         "500 with empty body",
			statusCode:   500,
			contentType:  "application/json",
			responseBody: "",
			wantMessage:  "",
		},
		{
			name:         "503 with non-JSON body",
			statusCode:   503,
			contentType:  "text/plain",
			responseBody: "Service Unavailable",
			wantMessage:  "Service Unavailable",
		},
		{
			name:          "500 with X-Request-ID header",
			statusCode:    500,
			contentType:   "application/json",
			responseBody:  `{"message":"Server error"}`,
			requestID:     "req-abc-123",
			wantMessage:   "Server error",
			wantRequestID: "req-abc-123",
		},
		{
			name:          "500 with request_id in body overrides header",
			statusCode:    500,
			contentType:   "application/json",
			responseBody:  `{"message":"Server error","request_id":"req-from-body"}`,
			requestID:     "req-from-header",
			wantMessage:   "Server error",
			wantRequestID: "req-from-body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.requestID != "" {
					w.Header().Set("X-Request-ID", tt.requestID)
				}
				w.Header().Set("Content-Type", tt.contentType)
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			c, err := NewClient("test-account", "test-secret", WithBaseURL(server.URL))
			if err != nil {
				t.Fatalf("NewClient() error = %v", err)
			}

			var result map[string]interface{}
			err = c.Request(context.Background(), "GET", "/test", nil, &result)

			if err == nil {
				t.Fatal("expected error, got nil")
			}

			var apiErr *APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("expected *APIError, got %T: %v", err, err)
			}

			if apiErr.StatusCode != tt.statusCode {
				t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, tt.statusCode)
			}

			if apiErr.Message != tt.wantMessage {
				t.Errorf("Message = %q, want %q", apiErr.Message, tt.wantMessage)
			}

			if tt.wantRequestID != "" && apiErr.RequestID != tt.wantRequestID {
				t.Errorf("RequestID = %q, want %q", apiErr.RequestID, tt.wantRequestID)
			}

			if tt.wantRawBody != "" && apiErr.RawBody != tt.wantRawBody {
				t.Errorf("RawBody = %q, want %q", apiErr.RawBody, tt.wantRawBody)
			}

			// Verify Error() string includes status code
			errStr := apiErr.Error()
			if !strings.Contains(errStr, fmt.Sprintf("status %d", tt.statusCode)) {
				t.Errorf("Error() = %q, expected it to contain status %d", errStr, tt.statusCode)
			}
		})
	}
}
