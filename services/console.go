package services

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/access_grid/accessgrid-go/client"
	"github.com/access_grid/accessgrid-go/models"
)

// ConsoleService handles operations related to the enterprise console
type ConsoleService struct {
	client *client.Client
}

// NewConsoleService creates a new ConsoleService
func NewConsoleService(client *client.Client) *ConsoleService {
	return &ConsoleService{client: client}
}

// CreateTemplate creates a new card template
func (s *ConsoleService) CreateTemplate(params models.CreateTemplateParams) (*models.Template, error) {
	var template models.Template
	err := s.client.Request(http.MethodPost, "/templates", params, &template)
	if err != nil {
		return nil, fmt.Errorf("error creating template: %w", err)
	}
	return &template, nil
}

// UpdateTemplate updates an existing card template
func (s *ConsoleService) UpdateTemplate(params models.UpdateTemplateParams) (*models.Template, error) {
	var template models.Template
	path := fmt.Sprintf("/templates/%s", params.CardTemplateID)
	err := s.client.Request(http.MethodPut, path, params, &template)
	if err != nil {
		return nil, fmt.Errorf("error updating template: %w", err)
	}
	return &template, nil
}

// ReadTemplate retrieves a card template by ID
func (s *ConsoleService) ReadTemplate(templateID string) (*models.Template, error) {
	var template models.Template
	path := fmt.Sprintf("/templates/%s", templateID)
	err := s.client.Request(http.MethodGet, path, nil, &template)
	if err != nil {
		return nil, fmt.Errorf("error reading template: %w", err)
	}
	return &template, nil
}

// ListTemplates retrieves all card templates
func (s *ConsoleService) ListTemplates() ([]models.Template, error) {
	var templates []models.Template
	err := s.client.Request(http.MethodGet, "/templates", nil, &templates)
	if err != nil {
		return nil, fmt.Errorf("error listing templates: %w", err)
	}
	return templates, nil
}

// DeleteTemplate deletes a card template
func (s *ConsoleService) DeleteTemplate(templateID string) error {
	path := fmt.Sprintf("/templates/%s", templateID)
	err := s.client.Request(http.MethodDelete, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error deleting template: %w", err)
	}
	return nil
}

// EventLog retrieves event logs for a specific template
func (s *ConsoleService) EventLog(templateID string, filters models.EventLogFilters) ([]models.Event, error) {
	var events []models.Event
	
	// Build query parameters
	query := url.Values{}
	if filters.Device != "" {
		query.Add("device", filters.Device)
	}
	if filters.StartDate != "" {
		query.Add("start_date", filters.StartDate)
	}
	if filters.EndDate != "" {
		query.Add("end_date", filters.EndDate)
	}
	if filters.EventType != "" {
		query.Add("event_type", filters.EventType)
	}
	
	path := fmt.Sprintf("/templates/%s/events", templateID)
	if len(query) > 0 {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}
	
	err := s.client.Request(http.MethodGet, path, nil, &events)
	if err != nil {
		return nil, fmt.Errorf("error fetching event log: %w", err)
	}
	
	return events, nil
}