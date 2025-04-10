// Package accessgrid provides a client for the AccessGrid API
package accessgrid

import (
	"net/http"
	
	"github.com/access_grid/accessgrid-go/client"
	"github.com/access_grid/accessgrid-go/models"
	"github.com/access_grid/accessgrid-go/services"
)

// Client is the main entry point for the AccessGrid API
type Client struct {
	client      *client.Client
	AccessCards *services.AccessCardsService
	Console     *services.ConsoleService
}

// NewClient creates a new AccessGrid API client
func NewClient(accountID, secretKey string, options ...client.ClientOption) (*Client, error) {
	c, err := client.NewClient(accountID, secretKey, options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:      c,
		AccessCards: services.NewAccessCardsService(c),
		Console:     services.NewConsoleService(c),
	}, nil
}

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(url string) client.ClientOption {
	return client.WithBaseURL(url)
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) client.ClientOption {
	return client.WithHTTPClient(httpClient)
}

// Export model types for easy access
type (
	// Card represents an NFC key or access pass
	Card = models.Card
	
	// ProvisionParams defines parameters for provisioning a new card
	ProvisionParams = models.ProvisionParams
	
	// UpdateParams defines parameters for updating an existing card
	UpdateParams = models.UpdateParams
	
	// ListKeysParams defines parameters for filtering cards
	ListKeysParams = models.ListKeysParams
	
	// Template represents a card template
	Template = models.Template
	
	// TemplateDesign represents the design elements of a card template
	TemplateDesign = models.TemplateDesign
	
	// SupportInfo represents support information for a card template
	SupportInfo = models.SupportInfo
	
	// CreateTemplateParams defines parameters for creating a new template
	CreateTemplateParams = models.CreateTemplateParams
	
	// UpdateTemplateParams defines parameters for updating an existing template
	UpdateTemplateParams = models.UpdateTemplateParams
	
	// EventLogFilters defines parameters for filtering event logs
	EventLogFilters = models.EventLogFilters
	
	// Event represents an event in the event log
	Event = models.Event
)