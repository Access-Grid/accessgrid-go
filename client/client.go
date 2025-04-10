package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.accessgrid.com/v1"
	defaultTimeout = 30 * time.Second
	version        = "0.1.0"
)

// Client is the main AccessGrid API client
type Client struct {
	AccountID  string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
}

// ClientOption allows for customizing the client
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.BaseURL = url
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// NewClient creates a new AccessGrid API client
func NewClient(accountID, secretKey string, options ...ClientOption) (*Client, error) {
	if accountID == "" {
		return nil, errors.New("accountID is required")
	}
	if secretKey == "" {
		return nil, errors.New("secretKey is required")
	}

	client := &Client{
		AccountID:  accountID,
		SecretKey:  secretKey,
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: defaultTimeout},
	}

	// Apply any custom options
	for _, option := range options {
		option(client)
	}

	return client, nil
}

// Request makes an authenticated API request
func (c *Client) Request(method, path string, body interface{}, result interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error marshaling request body: %w", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Set headers to match Python SDK
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ACCT-ID", c.AccountID)
	req.Header.Set("User-Agent", fmt.Sprintf("accessgrid.go @ v%s", version))

	// Generate signature
	signature, err := c.signRequest(reqBody)
	if err != nil {
		return fmt.Errorf("error signing request: %w", err)
	}
	req.Header.Set("X-PAYLOAD-SIG", signature)

	// Send the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// Check for API errors
	if resp.StatusCode >= 400 {
		var apiError struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(respBody, &apiError); err != nil {
			return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
		}
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, apiError.Error)
	}

	// Parse response into result
	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("error unmarshaling response: %w", err)
		}
	}

	return nil
}

// signRequest generates a signature matching the Python SDK implementation
func (c *Client) signRequest(payload []byte) (string, error) {
	var payloadStr string
	if payload != nil {
		payloadStr = string(payload)
	} else {
		payloadStr = "{}"
	}

	// Base64 encode the payload
	encodedPayload := base64.StdEncoding.EncodeToString([]byte(payloadStr))

	// Create HMAC using the shared secret as the key and the base64 encoded payload as the message
	h := hmac.New(sha256.New, []byte(c.SecretKey))
	_, err := h.Write([]byte(encodedPayload))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}