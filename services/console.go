package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Access-Grid/accessgrid-go/client"
	"github.com/Access-Grid/accessgrid-go/models"
)

// ConsoleService handles operations related to the enterprise console
type ConsoleService struct {
	client             *client.Client
	HID                *HIDService
	Webhooks           *WebhooksService
	CredentialProfiles *CredentialProfilesService
}

// NewConsoleService creates a new ConsoleService
func NewConsoleService(c *client.Client) *ConsoleService {
	return &ConsoleService{
		client:             c,
		HID:                NewHIDService(c),
		Webhooks:           NewWebhooksService(c),
		CredentialProfiles: NewCredentialProfilesService(c),
	}
}

// IosPreflight retrieves iOS In-App Provisioning identifiers
func (s *ConsoleService) IosPreflight(ctx context.Context, params models.IosPreflightParams) (*models.IosPreflight, error) {
	var result models.IosPreflight
	path := fmt.Sprintf("/v1/console/card-templates/%s/ios_preflight", url.PathEscape(params.CardTemplateID))
	body := map[string]string{"access_pass_ex_id": params.AccessPassExID}
	err := s.client.Request(ctx, http.MethodPost, path, body, &result)
	if err != nil {
		return nil, fmt.Errorf("error fetching iOS preflight: %w", err)
	}
	return &result, nil
}

// WebhooksService handles webhook operations
type WebhooksService struct {
	client *client.Client
}

// NewWebhooksService creates a new WebhooksService
func NewWebhooksService(c *client.Client) *WebhooksService {
	return &WebhooksService{client: c}
}

// Create creates a new webhook
func (s *WebhooksService) Create(ctx context.Context, params models.CreateWebhookParams) (*models.Webhook, error) {
	if params.AuthMethod == "" {
		params.AuthMethod = "bearer_token"
	}
	var webhook models.Webhook
	err := s.client.Request(ctx, http.MethodPost, "/v1/console/webhooks", params, &webhook)
	if err != nil {
		return nil, fmt.Errorf("error creating webhook: %w", err)
	}
	return &webhook, nil
}

// List retrieves all webhooks
func (s *WebhooksService) List(ctx context.Context) (*models.WebhooksResponse, error) {
	var response models.WebhooksResponse
	err := s.client.Request(ctx, http.MethodGet, "/v1/console/webhooks", nil, &response)
	if err != nil {
		return nil, fmt.Errorf("error listing webhooks: %w", err)
	}
	return &response, nil
}

// Delete deletes a webhook by ID
func (s *WebhooksService) Delete(ctx context.Context, webhookID string) error {
	path := fmt.Sprintf("/v1/console/webhooks/%s", url.PathEscape(webhookID))
	err := s.client.Request(ctx, http.MethodDelete, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error deleting webhook: %w", err)
	}
	return nil
}

// HIDService provides access to HID-related services
type HIDService struct {
	Orgs *HIDOrgsService
}

// NewHIDService creates a new HIDService
func NewHIDService(c *client.Client) *HIDService {
	return &HIDService{
		Orgs: NewHIDOrgsService(c),
	}
}

// HIDOrgsService handles HID organization operations
type HIDOrgsService struct {
	client *client.Client
}

// NewHIDOrgsService creates a new HIDOrgsService
func NewHIDOrgsService(c *client.Client) *HIDOrgsService {
	return &HIDOrgsService{client: c}
}

// Create creates a new HID organization
func (s *HIDOrgsService) Create(ctx context.Context, params *models.CreateHIDOrgParams) (*models.HIDOrg, error) {
	var org models.HIDOrg
	err := s.client.Request(ctx, http.MethodPost, "/v1/console/hid/orgs", params, &org)
	if err != nil {
		return nil, fmt.Errorf("error creating HID org: %w", err)
	}
	return &org, nil
}

// List retrieves all HID organizations
func (s *HIDOrgsService) List(ctx context.Context) ([]models.HIDOrg, error) {
	var orgs []models.HIDOrg
	err := s.client.Request(ctx, http.MethodGet, "/v1/console/hid/orgs", nil, &orgs)
	if err != nil {
		return nil, fmt.Errorf("error listing HID orgs: %w", err)
	}
	return orgs, nil
}

// Activate completes HID org registration with credentials
func (s *HIDOrgsService) Activate(ctx context.Context, params *models.CompleteHIDOrgParams) (*models.HIDOrg, error) {
	var org models.HIDOrg
	err := s.client.Request(ctx, http.MethodPost, "/v1/console/hid/orgs/activate", params, &org)
	if err != nil {
		return nil, fmt.Errorf("error activating HID org: %w", err)
	}
	return &org, nil
}

// CreateTemplate creates a new card template
func (s *ConsoleService) CreateTemplate(ctx context.Context, params models.CreateTemplateParams) (*models.Template, error) {
	var template models.Template
	err := s.client.Request(ctx, http.MethodPost, "/v1/console/card-templates", params, &template)
	if err != nil {
		return nil, fmt.Errorf("error creating template: %w", err)
	}
	return &template, nil
}

// UpdateTemplate updates an existing card template
func (s *ConsoleService) UpdateTemplate(ctx context.Context, params models.UpdateTemplateParams) (*models.Template, error) {
	var template models.Template
	path := fmt.Sprintf("/v1/console/card-templates/%s", url.PathEscape(params.CardTemplateID))
	err := s.client.Request(ctx, http.MethodPut, path, params, &template)
	if err != nil {
		return nil, fmt.Errorf("error updating template: %w", err)
	}
	return &template, nil
}

// ReadTemplate retrieves a card template by ID
func (s *ConsoleService) ReadTemplate(ctx context.Context, templateID string) (*models.Template, error) {
	var template models.Template
	path := fmt.Sprintf("/v1/console/card-templates/%s", url.PathEscape(templateID))
	err := s.client.Request(ctx, http.MethodGet, path, nil, &template)
	if err != nil {
		return nil, fmt.Errorf("error reading template: %w", err)
	}
	return &template, nil
}

// ListTemplates retrieves all card templates
func (s *ConsoleService) ListTemplates(ctx context.Context) ([]models.Template, error) {
	var templates []models.Template
	err := s.client.Request(ctx, http.MethodGet, "/v1/console/card-templates", nil, &templates)
	if err != nil {
		return nil, fmt.Errorf("error listing templates: %w", err)
	}
	return templates, nil
}

// DeleteTemplate deletes a card template
func (s *ConsoleService) DeleteTemplate(ctx context.Context, templateID string) error {
	path := fmt.Sprintf("/v1/console/card-templates/%s", url.PathEscape(templateID))
	err := s.client.Request(ctx, http.MethodDelete, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error deleting template: %w", err)
	}
	return nil
}

// ListPassTemplatePairs retrieves pass template pairs
func (s *ConsoleService) ListPassTemplatePairs(ctx context.Context, params models.ListPassTemplatePairsParams) (*models.PassTemplatePairsResponse, error) {
	var response models.PassTemplatePairsResponse

	query := url.Values{}
	if params.Page > 0 {
		query.Add("page", fmt.Sprintf("%d", params.Page))
	}
	if params.PerPage > 0 {
		query.Add("per_page", fmt.Sprintf("%d", params.PerPage))
	}

	u := url.URL{Path: "/v1/console/pass-template-pairs"}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	err := s.client.Request(ctx, http.MethodGet, u.String(), nil, &response)
	if err != nil {
		return nil, fmt.Errorf("error listing pass template pairs: %w", err)
	}
	return &response, nil
}

// ListLedgerItems retrieves billing ledger items
func (s *ConsoleService) ListLedgerItems(ctx context.Context, params models.ListLedgerItemsParams) (*models.LedgerItemsResponse, error) {
	var response models.LedgerItemsResponse

	query := url.Values{}
	if params.Page > 0 {
		query.Add("page", fmt.Sprintf("%d", params.Page))
	}
	if params.PerPage > 0 {
		query.Add("per_page", fmt.Sprintf("%d", params.PerPage))
	}
	if params.StartDate != nil {
		query.Add("start_date", params.StartDate.Format(time.RFC3339))
	}
	if params.EndDate != nil {
		query.Add("end_date", params.EndDate.Format(time.RFC3339))
	}

	u := url.URL{Path: "/v1/console/ledger-items"}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	err := s.client.Request(ctx, http.MethodGet, u.String(), nil, &response)
	if err != nil {
		return nil, fmt.Errorf("error listing ledger items: %w", err)
	}
	return &response, nil
}

// ListLandingPages retrieves all landing pages
func (s *ConsoleService) ListLandingPages(ctx context.Context) ([]models.LandingPage, error) {
	var pages []models.LandingPage
	err := s.client.Request(ctx, http.MethodGet, "/v1/console/landing-pages", nil, &pages)
	if err != nil {
		return nil, fmt.Errorf("error listing landing pages: %w", err)
	}
	return pages, nil
}

// CreateLandingPage creates a new landing page
func (s *ConsoleService) CreateLandingPage(ctx context.Context, params models.CreateLandingPageParams) (*models.LandingPage, error) {
	var page models.LandingPage
	err := s.client.Request(ctx, http.MethodPost, "/v1/console/landing-pages", params, &page)
	if err != nil {
		return nil, fmt.Errorf("error creating landing page: %w", err)
	}
	return &page, nil
}

// UpdateLandingPage updates an existing landing page
func (s *ConsoleService) UpdateLandingPage(ctx context.Context, params models.UpdateLandingPageParams) (*models.LandingPage, error) {
	var page models.LandingPage
	path := fmt.Sprintf("/v1/console/landing-pages/%s", url.PathEscape(params.LandingPageID))
	err := s.client.Request(ctx, http.MethodPut, path, params, &page)
	if err != nil {
		return nil, fmt.Errorf("error updating landing page: %w", err)
	}
	return &page, nil
}

// CredentialProfilesService handles credential profile operations
type CredentialProfilesService struct {
	client *client.Client
}

// NewCredentialProfilesService creates a new CredentialProfilesService
func NewCredentialProfilesService(c *client.Client) *CredentialProfilesService {
	return &CredentialProfilesService{client: c}
}

// List retrieves all credential profiles
func (s *CredentialProfilesService) List(ctx context.Context) ([]models.CredentialProfile, error) {
	var profiles []models.CredentialProfile
	err := s.client.Request(ctx, http.MethodGet, "/v1/console/credential-profiles", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("error listing credential profiles: %w", err)
	}
	return profiles, nil
}

// Create creates a new credential profile
func (s *CredentialProfilesService) Create(ctx context.Context, params models.CreateCredentialProfileParams) (*models.CredentialProfile, error) {
	var profile models.CredentialProfile
	err := s.client.Request(ctx, http.MethodPost, "/v1/console/credential-profiles", params, &profile)
	if err != nil {
		return nil, fmt.Errorf("error creating credential profile: %w", err)
	}
	return &profile, nil
}

// EventLog retrieves event logs for a specific template
func (s *ConsoleService) EventLog(ctx context.Context, templateID string, filters models.EventLogFilters) ([]models.Event, error) {
	var events []models.Event

	// Build query parameters
	query := url.Values{}
	if filters.Device != "" {
		query.Add("device", filters.Device)
	}
	if filters.StartDate != nil {
		query.Add("start_date", filters.StartDate.Format(time.RFC3339))
	}
	if filters.EndDate != nil {
		query.Add("end_date", filters.EndDate.Format(time.RFC3339))
	}
	if filters.EventType != "" {
		query.Add("event_type", filters.EventType)
	}

	// Build the URL properly using url.URL
	u := url.URL{
		Path: fmt.Sprintf("/v1/console/card-templates/%s/logs", url.PathEscape(templateID)),
	}

	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	path := u.String()

	err := s.client.Request(ctx, http.MethodGet, path, nil, &events)
	if err != nil {
		return nil, fmt.Errorf("error fetching event log: %w", err)
	}

	return events, nil
}
