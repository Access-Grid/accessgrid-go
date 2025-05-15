package accessgrid

import (
	"net/http"
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
	customClient := &http.Client{Timeout: 5 * time.Second}

	client, err := NewClient(
		accountID,
		secretKey,
		WithBaseURL(customURL),
		WithHTTPClient(customClient),
	)

	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Verify service instances
	if client.AccessCards == nil {
		t.Error("Expected AccessCards service to be initialized")
	}

	if client.Console == nil {
		t.Error("Expected Console service to be initialized")
	}
}
