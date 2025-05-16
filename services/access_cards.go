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
func (s *AccessCardsService) Provision(ctx context.Context, params models.ProvisionParams) (*models.Card, error) {
	var card models.Card
	err := s.client.Request(ctx, http.MethodPost, "/v1/key-cards", params, &card)
	if err != nil {
		return nil, fmt.Errorf("error provisioning card: %w", err)
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
	err := s.client.Request(ctx, http.MethodGet, "/v1/key-cards", params, &response)
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
