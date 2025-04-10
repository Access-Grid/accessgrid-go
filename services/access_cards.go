package services

import (
	"fmt"
	"net/http"

	"github.com/access_grid/accessgrid-go/client"
	"github.com/access_grid/accessgrid-go/models"
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
func (s *AccessCardsService) Provision(params models.ProvisionParams) (*models.Card, error) {
	var card models.Card
	err := s.client.Request(http.MethodPost, "/cards", params, &card)
	if err != nil {
		return nil, fmt.Errorf("error provisioning card: %w", err)
	}
	return &card, nil
}

// Update updates an existing NFC key/card
func (s *AccessCardsService) Update(params models.UpdateParams) (*models.Card, error) {
	var card models.Card
	path := fmt.Sprintf("/cards/%s", params.CardID)
	err := s.client.Request(http.MethodPut, path, params, &card)
	if err != nil {
		return nil, fmt.Errorf("error updating card: %w", err)
	}
	return &card, nil
}

// List retrieves cards with optional filtering
func (s *AccessCardsService) List(params *models.ListKeysParams) ([]models.Card, error) {
	var cards []models.Card
	err := s.client.Request(http.MethodGet, "/cards", params, &cards)
	if err != nil {
		return nil, fmt.Errorf("error listing cards: %w", err)
	}
	return cards, nil
}

// Suspend suspends a card
func (s *AccessCardsService) Suspend(cardID string) error {
	path := fmt.Sprintf("/cards/%s/suspend", cardID)
	err := s.client.Request(http.MethodPost, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error suspending card: %w", err)
	}
	return nil
}

// Resume resumes a suspended card
func (s *AccessCardsService) Resume(cardID string) error {
	path := fmt.Sprintf("/cards/%s/resume", cardID)
	err := s.client.Request(http.MethodPost, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error resuming card: %w", err)
	}
	return nil
}

// Unlink unlinks a card from a device
func (s *AccessCardsService) Unlink(cardID string) error {
	path := fmt.Sprintf("/cards/%s/unlink", cardID)
	err := s.client.Request(http.MethodPost, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error unlinking card: %w", err)
	}
	return nil
}

// Delete deletes a card
func (s *AccessCardsService) Delete(cardID string) error {
	path := fmt.Sprintf("/cards/%s", cardID)
	err := s.client.Request(http.MethodDelete, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error deleting card: %w", err)
	}
	return nil
}