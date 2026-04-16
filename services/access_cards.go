package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Access-Grid/accessgrid-go/client"
	"github.com/Access-Grid/accessgrid-go/models"
)

// AccessCardsService handles operations related to NFC cards
type AccessCardsService struct {
	client *client.Client
}

// NewAccessCardsService creates a new AccessCardsService
func NewAccessCardsService(client *client.Client) *AccessCardsService {
	return &AccessCardsService{client: client}
}

// Provision creates a new NFC key/card
func (s *AccessCardsService) Provision(ctx context.Context, params models.ProvisionParams) (*models.CardProvisionResponse, error) {
	var res models.CardProvisionResponse
	err := s.client.Request(ctx, http.MethodPost, "/v1/key-cards", params, &res)
	if err != nil {
		return nil, fmt.Errorf("error provisioning card: %w", err)
	}
	return &res, nil
}

// Get retrieves a specific NFC key/card by ID
func (s *AccessCardsService) Get(ctx context.Context, cardID string) (*models.Card, error) {
	var card models.Card
	path := fmt.Sprintf("/v1/key-cards/%s", url.PathEscape(cardID))
	err := s.client.Request(ctx, http.MethodGet, path, nil, &card)
	if err != nil {
		return nil, fmt.Errorf("error getting card: %w", err)
	}
	return &card, nil
}

// Update updates an existing NFC key/card
func (s *AccessCardsService) Update(ctx context.Context, params models.UpdateParams) (*models.Card, error) {
	var card models.Card
	path := fmt.Sprintf("/v1/key-cards/%s", url.PathEscape(params.CardID))
	err := s.client.Request(ctx, http.MethodPatch, path, params, &card)
	if err != nil {
		return nil, fmt.Errorf("error updating card: %w", err)
	}
	return &card, nil
}

// List retrieves cards with optional filtering
func (s *AccessCardsService) List(ctx context.Context, params *models.ListKeysParams) ([]models.Card, error) {
	var response struct {
		Keys []models.Card `json:"keys"`
	}

	query := url.Values{}
	if params != nil {
		if params.TemplateID != "" {
			query.Set("template_id", params.TemplateID)
		}
		if params.State != "" {
			query.Set("state", params.State)
		}
		if params.EmployeeID != "" {
			query.Set("employee_id", params.EmployeeID)
		}
		if params.CardNumber != "" {
			query.Set("card_number", params.CardNumber)
		}
		if params.SiteCode != "" {
			query.Set("site_code", params.SiteCode)
		}
	}

	path := "/v1/key-cards"
	if len(query) > 0 {
		path = path + "?" + query.Encode()
	}

	err := s.client.Request(ctx, http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("error listing cards: %w", err)
	}
	return response.Keys, nil
}

// Suspend suspends a card
func (s *AccessCardsService) Suspend(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/v1/key-cards/%s/suspend", url.PathEscape(cardID))
	err := s.client.Request(ctx, http.MethodPost, path, map[string]string{}, nil)
	if err != nil {
		return fmt.Errorf("error suspending card: %w", err)
	}
	return nil
}

// Resume resumes a suspended card
func (s *AccessCardsService) Resume(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/v1/key-cards/%s/resume", url.PathEscape(cardID))
	err := s.client.Request(ctx, http.MethodPost, path, map[string]string{}, nil)
	if err != nil {
		return fmt.Errorf("error resuming card: %w", err)
	}
	return nil
}

// Unlink unlinks a card from a device
func (s *AccessCardsService) Unlink(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/v1/key-cards/%s/unlink", url.PathEscape(cardID))
	err := s.client.Request(ctx, http.MethodPost, path, map[string]string{}, nil)
	if err != nil {
		return fmt.Errorf("error unlinking card: %w", err)
	}
	return nil
}

// Delete deletes a card
func (s *AccessCardsService) Delete(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/v1/key-cards/%s/delete", url.PathEscape(cardID))
	err := s.client.Request(ctx, http.MethodPost, path, map[string]string{}, nil)
	if err != nil {
		return fmt.Errorf("error deleting card: %w", err)
	}
	return nil
}
