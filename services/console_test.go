package services

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Access-Grid/accessgrid-go/client"
	"github.com/Access-Grid/accessgrid-go/models"
)

func setupConsoleTestServer() (*httptest.Server, *ConsoleService) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		switch r.URL.Path {
		case "/v1/console/card-templates":
			if r.Method == http.MethodPost {
				// Create Template
				w.Write([]byte(`{
					"id": "0xd3adb00b5",
					"name": "Employee NFC key",
					"platform": "apple",
					"use_case": "employee_badge",
					"protocol": "desfire",
					"watch_count": 2,
					"iphone_count": 3
				}`))
			} else if r.Method == http.MethodGet {
				// List Templates
				w.Write([]byte(`[
					{
						"id": "0xd3adb00b5",
						"name": "Employee NFC key",
						"platform": "apple",
						"protocol": "desfire"
					}
				]`))
			}
		case "/v1/console/card-templates/0xd3adb00b5":
			if r.Method == http.MethodPut {
				// Update Template
				w.Write([]byte(`{
					"id": "0xd3adb00b5",
					"name": "Updated Employee NFC key",
					"platform": "apple",
					"protocol": "desfire"
				}`))
			} else if r.Method == http.MethodGet {
				// Read Template
				w.Write([]byte(`{
					"id": "0xd3adb00b5",
					"name": "Employee NFC key",
					"platform": "apple",
					"protocol": "desfire",
					"watch_count": 2,
					"iphone_count": 3
				}`))
			} else if r.Method == http.MethodDelete {
				// Delete Template
				w.Write([]byte(`{}`))
			}
		case "/v1/console/card-templates/0xd3adb00b5/logs":
			// Event Log
			w.Write([]byte(`[
				{
					"id": "evt_123",
					"type": "install",
					"user_id": "usr_456",
					"card_id": "0xc4rd1d",
					"template_id": "0xd3adb00b5",
					"device": "mobile",
					"timestamp": "2023-01-01T12:00:00Z"
				}
			]`))
		case "/v1/console/pass-template-pairs":
			w.Write([]byte(`{
				"pass_template_pairs": [
					{
						"id": "pair_1",
						"name": "Employee Badge Pair",
						"created_at": "2025-01-01T00:00:00Z",
						"ios_template": {"id": "tmpl_ios_1", "name": "iOS Badge", "platform": "apple"},
						"android_template": {"id": "tmpl_android_1", "name": "Android Badge", "platform": "android"}
					},
					{
						"id": "pair_2",
						"name": "Contractor Badge Pair",
						"created_at": "2025-01-02T00:00:00Z",
						"ios_template": {"id": "tmpl_ios_2", "name": "iOS Contractor", "platform": "apple"},
						"android_template": null
					}
				],
				"pagination": {
					"current_page": 1,
					"total_pages": 1
				}
			}`))
		case "/v1/console/ledger-items":
			w.Write([]byte(`{
				"ledger_items": [
					{
						"created_at": "2025-06-15T14:30:00Z",
						"amount": -1.50,
						"id": "li_abc123",
						"kind": "access_pass_debit",
						"metadata": {
							"access_pass_ex_id": "ap_xyz",
							"pass_template_ex_id": "pt_456"
						},
						"access_pass": {
							"id": "ap_xyz",
							"full_name": "Jane Doe",
							"state": "active",
							"metadata": {"department": "Engineering"},
							"unified_access_pass_ex_id": "uap_789",
							"pass_template": {
								"id": "pt_456",
								"name": "Employee Badge",
								"protocol": "desfire",
								"platform": "apple",
								"use_case": "employee_badge"
							}
						}
					},
					{
						"created_at": "2025-06-14T08:15:00Z",
						"amount": 500.00,
						"id": "li_def456",
						"kind": "credit",
						"metadata": {},
						"access_pass": null
					}
				],
				"pagination": {
					"current_page": 1,
					"per_page": 50,
					"total_pages": 3,
					"total_count": 125
				}
			}`))
		}
	}))

	c, _ := client.NewClient("test-account", "test-secret", client.WithBaseURL(server.URL))
	service := NewConsoleService(c)

	return server, service
}

func TestConsoleService_CreateTemplate(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	params := models.CreateTemplateParams{
		Name:                   "Employee NFC key",
		Platform:               "apple",
		UseCase:                "employee_badge",
		Protocol:               "desfire",
		AllowOnMultipleDevices: true,
		WatchCount:             2,
		IPhoneCount:            3,
		BackgroundColor:        "#FFFFFF",
		LabelColor:             "#000000",
		LabelSecondaryColor:    "#333333",
		SupportURL:             "https://help.example.com",
		SupportPhoneNumber:     "+1-555-123-4567",
		SupportEmail:           "support@example.com",
		PrivacyPolicyURL:       "https://example.com/privacy",
		TermsAndConditionsURL:  "https://example.com/terms",
		Metadata: map[string]interface{}{
			"version":         "2.1",
			"approval_status": "approved",
		},
	}

	ctx := context.Background()
	template, err := service.CreateTemplate(ctx, params)
	if err != nil {
		t.Fatalf("CreateTemplate() error = %v", err)
	}

	if template.ID != "0xd3adb00b5" {
		t.Errorf("CreateTemplate() template.ID = %v, want %v", template.ID, "0xd3adb00b5")
	}
	if template.Name != "Employee NFC key" {
		t.Errorf("CreateTemplate() template.Name = %v, want %v", template.Name, "Employee NFC key")
	}
	if template.Platform != "apple" {
		t.Errorf("CreateTemplate() template.Platform = %v, want %v", template.Platform, "apple")
	}
}

func TestConsoleService_UpdateTemplate(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	supportInfo := models.SupportInfo{
		SupportURL:         "https://help.example.com",
		SupportPhoneNumber: "+1-555-123-4567",
		SupportEmail:       "support@example.com",
	}

	params := models.UpdateTemplateParams{
		CardTemplateID: "0xd3adb00b5",
		Name:           "Updated Employee NFC key",
		WatchCount:     2,
		IPhoneCount:    3,
		SupportInfo:    &supportInfo,
	}

	ctx := context.Background()
	template, err := service.UpdateTemplate(ctx, params)
	if err != nil {
		t.Fatalf("UpdateTemplate() error = %v", err)
	}

	if template.ID != "0xd3adb00b5" {
		t.Errorf("UpdateTemplate() template.ID = %v, want %v", template.ID, "0xd3adb00b5")
	}
	if template.Name != "Updated Employee NFC key" {
		t.Errorf("UpdateTemplate() template.Name = %v, want %v", template.Name, "Updated Employee NFC key")
	}
}

func TestConsoleService_ReadTemplate(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	ctx := context.Background()
	template, err := service.ReadTemplate(ctx, "0xd3adb00b5")
	if err != nil {
		t.Fatalf("ReadTemplate() error = %v", err)
	}

	if template.ID != "0xd3adb00b5" {
		t.Errorf("ReadTemplate() template.ID = %v, want %v", template.ID, "0xd3adb00b5")
	}
	if template.Name != "Employee NFC key" {
		t.Errorf("ReadTemplate() template.Name = %v, want %v", template.Name, "Employee NFC key")
	}
	if template.WatchCount != 2 {
		t.Errorf("ReadTemplate() template.WatchCount = %v, want %v", template.WatchCount, 2)
	}
}

func TestConsoleService_ListTemplates(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	ctx := context.Background()
	templates, err := service.ListTemplates(ctx)
	if err != nil {
		t.Fatalf("ListTemplates() error = %v", err)
	}

	if len(templates) != 1 {
		t.Fatalf("ListTemplates() got %v templates, want %v", len(templates), 1)
	}

	if templates[0].ID != "0xd3adb00b5" {
		t.Errorf("ListTemplates() templates[0].ID = %v, want %v", templates[0].ID, "0xd3adb00b5")
	}
}

func TestConsoleService_DeleteTemplate(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	ctx := context.Background()
	err := service.DeleteTemplate(ctx, "0xd3adb00b5")
	if err != nil {
		t.Errorf("DeleteTemplate() error = %v", err)
	}
}

func TestConsoleService_EventLog(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	startDate, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endDate, _ := time.Parse(time.RFC3339, "2023-01-31T23:59:59Z")

	filters := models.EventLogFilters{
		Device:    "mobile",
		StartDate: &startDate,
		EndDate:   &endDate,
		EventType: "install",
	}

	ctx := context.Background()
	events, err := service.EventLog(ctx, "0xd3adb00b5", filters)
	if err != nil {
		t.Fatalf("EventLog() error = %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("EventLog() got %v events, want %v", len(events), 1)
	}

	if events[0].Type != "install" {
		t.Errorf("EventLog() events[0].Type = %v, want %v", events[0].Type, "install")
	}
	if events[0].CardID != "0xc4rd1d" {
		t.Errorf("EventLog() events[0].CardID = %v, want %v", events[0].CardID, "0xc4rd1d")
	}
}

// --- Pass Template Pairs ---

func TestConsoleService_ListPassTemplatePairs(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	ctx := context.Background()
	response, err := service.ListPassTemplatePairs(ctx, models.ListPassTemplatePairsParams{})
	if err != nil {
		t.Fatalf("ListPassTemplatePairs() error = %v", err)
	}

	if len(response.PassTemplatePairs) != 2 {
		t.Fatalf("got %d pairs, want 2", len(response.PassTemplatePairs))
	}

	// First pair: both templates present
	pair := response.PassTemplatePairs[0]
	if pair.ID != "pair_1" {
		t.Errorf("pair.ID = %v, want pair_1", pair.ID)
	}
	if pair.Name != "Employee Badge Pair" {
		t.Errorf("pair.Name = %v, want Employee Badge Pair", pair.Name)
	}
	expectedTime, _ := time.Parse(time.RFC3339, "2025-01-01T00:00:00Z")
	if !pair.CreatedAt.Equal(expectedTime) {
		t.Errorf("pair.CreatedAt = %v, want %v", pair.CreatedAt, expectedTime)
	}
	if pair.IOSTemplate == nil {
		t.Fatal("pair.IOSTemplate is nil, want non-nil")
	}
	if pair.IOSTemplate.ID != "tmpl_ios_1" {
		t.Errorf("IOSTemplate.ID = %v, want tmpl_ios_1", pair.IOSTemplate.ID)
	}
	if pair.IOSTemplate.Name != "iOS Badge" {
		t.Errorf("IOSTemplate.Name = %v, want iOS Badge", pair.IOSTemplate.Name)
	}
	if pair.IOSTemplate.Platform != "apple" {
		t.Errorf("IOSTemplate.Platform = %v, want apple", pair.IOSTemplate.Platform)
	}
	if pair.AndroidTemplate == nil {
		t.Fatal("pair.AndroidTemplate is nil, want non-nil")
	}
	if pair.AndroidTemplate.ID != "tmpl_android_1" {
		t.Errorf("AndroidTemplate.ID = %v, want tmpl_android_1", pair.AndroidTemplate.ID)
	}

	// Second pair: null android_template
	pair2 := response.PassTemplatePairs[1]
	if pair2.AndroidTemplate != nil {
		t.Errorf("pair2.AndroidTemplate = %v, want nil", pair2.AndroidTemplate)
	}
	if pair2.IOSTemplate == nil {
		t.Fatal("pair2.IOSTemplate is nil, want non-nil")
	}

	// Pagination
	if response.Pagination.CurrentPage != 1 {
		t.Errorf("Pagination.CurrentPage = %v, want 1", response.Pagination.CurrentPage)
	}
	if response.Pagination.TotalPages != 1 {
		t.Errorf("Pagination.TotalPages = %v, want 1", response.Pagination.TotalPages)
	}
}

func TestConsoleService_ListPassTemplatePairs_WithPagination(t *testing.T) {
	var capturedURL string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedURL = r.URL.String()
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"pass_template_pairs": [], "pagination": {"current_page": 2, "total_pages": 5}}`))
	}))
	defer server.Close()

	c, _ := client.NewClient("test-account", "test-secret", client.WithBaseURL(server.URL))
	service := NewConsoleService(c)

	ctx := context.Background()
	response, err := service.ListPassTemplatePairs(ctx, models.ListPassTemplatePairsParams{
		Page:    2,
		PerPage: 10,
	})
	if err != nil {
		t.Fatalf("ListPassTemplatePairs() error = %v", err)
	}

	if len(response.PassTemplatePairs) != 0 {
		t.Errorf("got %d pairs, want 0", len(response.PassTemplatePairs))
	}
	if response.Pagination.CurrentPage != 2 {
		t.Errorf("Pagination.CurrentPage = %v, want 2", response.Pagination.CurrentPage)
	}
	if response.Pagination.TotalPages != 5 {
		t.Errorf("Pagination.TotalPages = %v, want 5", response.Pagination.TotalPages)
	}

	if !strings.Contains(capturedURL, "page=2") {
		t.Errorf("expected page=2 in URL, got %s", capturedURL)
	}
	if !strings.Contains(capturedURL, "per_page=10") {
		t.Errorf("expected per_page=10 in URL, got %s", capturedURL)
	}
}

// --- Ledger Items ---

func TestConsoleService_ListLedgerItems(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	ctx := context.Background()
	response, err := service.ListLedgerItems(ctx, models.ListLedgerItemsParams{})
	if err != nil {
		t.Fatalf("ListLedgerItems() error = %v", err)
	}

	if len(response.LedgerItems) != 2 {
		t.Fatalf("got %d items, want 2", len(response.LedgerItems))
	}

	// First item: full nested structure
	item := response.LedgerItems[0]
	expectedTime, _ := time.Parse(time.RFC3339, "2025-06-15T14:30:00Z")
	if !item.CreatedAt.Equal(expectedTime) {
		t.Errorf("item.CreatedAt = %v, want %v", item.CreatedAt, expectedTime)
	}
	if item.Amount != -1.50 {
		t.Errorf("item.Amount = %v, want -1.50", item.Amount)
	}
	if item.ID != "li_abc123" {
		t.Errorf("item.ID = %v, want li_abc123", item.ID)
	}
	if item.Kind != "access_pass_debit" {
		t.Errorf("item.Kind = %v, want access_pass_debit", item.Kind)
	}
	if item.Metadata["access_pass_ex_id"] != "ap_xyz" {
		t.Errorf("item.Metadata[access_pass_ex_id] = %v, want ap_xyz", item.Metadata["access_pass_ex_id"])
	}

	// Nested access_pass
	ap := item.AccessPass
	if ap == nil {
		t.Fatal("item.AccessPass is nil, want non-nil")
	}
	if ap.ID != "ap_xyz" {
		t.Errorf("AccessPass.ID = %v, want ap_xyz", ap.ID)
	}
	if ap.FullName != "Jane Doe" {
		t.Errorf("AccessPass.FullName = %v, want Jane Doe", ap.FullName)
	}
	if ap.State != "active" {
		t.Errorf("AccessPass.State = %v, want active", ap.State)
	}
	if ap.Metadata["department"] != "Engineering" {
		t.Errorf("AccessPass.Metadata[department] = %v, want Engineering", ap.Metadata["department"])
	}
	if ap.UnifiedAccessPassExID != "uap_789" {
		t.Errorf("AccessPass.UnifiedAccessPassExID = %v, want uap_789", ap.UnifiedAccessPassExID)
	}

	// Nested pass_template
	pt := ap.PassTemplate
	if pt == nil {
		t.Fatal("AccessPass.PassTemplate is nil, want non-nil")
	}
	if pt.ID != "pt_456" {
		t.Errorf("PassTemplate.ID = %v, want pt_456", pt.ID)
	}
	if pt.Name != "Employee Badge" {
		t.Errorf("PassTemplate.Name = %v, want Employee Badge", pt.Name)
	}
	if pt.Protocol != "desfire" {
		t.Errorf("PassTemplate.Protocol = %v, want desfire", pt.Protocol)
	}
	if pt.Platform != "apple" {
		t.Errorf("PassTemplate.Platform = %v, want apple", pt.Platform)
	}
	if pt.UseCase != "employee_badge" {
		t.Errorf("PassTemplate.UseCase = %v, want employee_badge", pt.UseCase)
	}

	// Second item: null access_pass
	item2 := response.LedgerItems[1]
	if item2.Kind != "credit" {
		t.Errorf("item2.Kind = %v, want credit", item2.Kind)
	}
	if item2.AccessPass != nil {
		t.Errorf("item2.AccessPass = %v, want nil", item2.AccessPass)
	}

	// Pagination
	if response.Pagination.CurrentPage != 1 {
		t.Errorf("Pagination.CurrentPage = %v, want 1", response.Pagination.CurrentPage)
	}
	if response.Pagination.PerPage != 50 {
		t.Errorf("Pagination.PerPage = %v, want 50", response.Pagination.PerPage)
	}
	if response.Pagination.TotalPages != 3 {
		t.Errorf("Pagination.TotalPages = %v, want 3", response.Pagination.TotalPages)
	}
	if response.Pagination.TotalCount != 125 {
		t.Errorf("Pagination.TotalCount = %v, want 125", response.Pagination.TotalCount)
	}
}

func TestConsoleService_ListLedgerItems_MissingPassTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"ledger_items": [
				{
					"created_at": "2025-06-15T14:30:00Z",
					"amount": -1.50,
					"id": "li_no_pt",
					"kind": "access_pass_debit",
					"metadata": {},
					"access_pass": {
						"id": "ap_orphan",
						"full_name": "John Smith",
						"state": "suspended",
						"metadata": {},
						"unified_access_pass_ex_id": null
					}
				}
			],
			"pagination": {"current_page": 1, "per_page": 50, "total_pages": 1, "total_count": 1}
		}`))
	}))
	defer server.Close()

	c, _ := client.NewClient("test-account", "test-secret", client.WithBaseURL(server.URL))
	service := NewConsoleService(c)

	ctx := context.Background()
	response, err := service.ListLedgerItems(ctx, models.ListLedgerItemsParams{})
	if err != nil {
		t.Fatalf("ListLedgerItems() error = %v", err)
	}

	ap := response.LedgerItems[0].AccessPass
	if ap == nil {
		t.Fatal("AccessPass is nil, want non-nil")
	}
	if ap.FullName != "John Smith" {
		t.Errorf("AccessPass.FullName = %v, want John Smith", ap.FullName)
	}
	if ap.UnifiedAccessPassExID != "" {
		t.Errorf("AccessPass.UnifiedAccessPassExID = %v, want empty", ap.UnifiedAccessPassExID)
	}
	if ap.PassTemplate != nil {
		t.Errorf("AccessPass.PassTemplate = %v, want nil", ap.PassTemplate)
	}
}

func TestConsoleService_ListLedgerItems_WithFilters(t *testing.T) {
	var capturedURL string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedURL = r.URL.String()
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ledger_items": [], "pagination": {"current_page": 2, "per_page": 20, "total_pages": 1, "total_count": 0}}`))
	}))
	defer server.Close()

	c, _ := client.NewClient("test-account", "test-secret", client.WithBaseURL(server.URL))
	service := NewConsoleService(c)

	startDate, _ := time.Parse(time.RFC3339, "2025-01-01T00:00:00Z")
	endDate, _ := time.Parse(time.RFC3339, "2025-06-30T23:59:59Z")

	ctx := context.Background()
	response, err := service.ListLedgerItems(ctx, models.ListLedgerItemsParams{
		Page:      2,
		PerPage:   20,
		StartDate: &startDate,
		EndDate:   &endDate,
	})
	if err != nil {
		t.Fatalf("ListLedgerItems() error = %v", err)
	}

	if len(response.LedgerItems) != 0 {
		t.Errorf("got %d items, want 0", len(response.LedgerItems))
	}
	if response.Pagination.CurrentPage != 2 {
		t.Errorf("Pagination.CurrentPage = %v, want 2", response.Pagination.CurrentPage)
	}

	if !strings.Contains(capturedURL, "page=2") {
		t.Errorf("expected page=2 in URL, got %s", capturedURL)
	}
	if !strings.Contains(capturedURL, "per_page=20") {
		t.Errorf("expected per_page=20 in URL, got %s", capturedURL)
	}
	if !strings.Contains(capturedURL, "start_date=") {
		t.Errorf("expected start_date in URL, got %s", capturedURL)
	}
	if !strings.Contains(capturedURL, "end_date=") {
		t.Errorf("expected end_date in URL, got %s", capturedURL)
	}
}

func TestConsoleService_ErrorPropagation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Resource not found"}`))
	}))
	defer server.Close()

	c, _ := client.NewClient("test-account", "test-secret", client.WithBaseURL(server.URL))
	service := NewConsoleService(c)

	ctx := context.Background()
	_, err := service.ReadTemplate(ctx, "nonexistent-id")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Verify the service wraps the error with context
	if !strings.Contains(err.Error(), "error reading template") {
		t.Errorf("expected wrapped message, got: %s", err.Error())
	}

	// Verify the underlying APIError is still accessible
	var apiErr *client.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected unwrappable *client.APIError, got %T", err)
	}

	if apiErr.StatusCode != 404 {
		t.Errorf("StatusCode = %d, want 404", apiErr.StatusCode)
	}

	if apiErr.Message != "Resource not found" {
		t.Errorf("Message = %q, want %q", apiErr.Message, "Resource not found")
	}
}
