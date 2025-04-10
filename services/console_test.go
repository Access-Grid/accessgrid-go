package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/access_grid/accessgrid-go/client"
	"github.com/access_grid/accessgrid-go/models"
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
		}
	}))

	c, _ := client.NewClient("test-account", "test-secret", client.WithBaseURL(server.URL))
	service := NewConsoleService(c)

	return server, service
}

func TestConsoleService_CreateTemplate(t *testing.T) {
	server, service := setupConsoleTestServer()
	defer server.Close()

	design := models.TemplateDesign{
		BackgroundColor:     "#FFFFFF",
		LabelColor:         "#000000",
		LabelSecondaryColor: "#333333",
	}

	supportInfo := models.SupportInfo{
		SupportURL:           "https://help.example.com",
		SupportPhoneNumber:   "+1-555-123-4567",
		SupportEmail:         "support@example.com",
		PrivacyPolicyURL:     "https://example.com/privacy",
		TermsAndConditionsURL: "https://example.com/terms",
	}

	params := models.CreateTemplateParams{
		Name:                 "Employee NFC key",
		Platform:            "apple",
		UseCase:             "employee_badge",
		Protocol:            "desfire",
		WatchCount:          2,
		IPhoneCount:         3,
		Design:              design,
		SupportInfo:         supportInfo,
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
		SupportURL:           "https://help.example.com",
		SupportPhoneNumber:   "+1-555-123-4567",
		SupportEmail:         "support@example.com",
	}

	params := models.UpdateTemplateParams{
		CardTemplateID:       "0xd3adb00b5",
		Name:                "Updated Employee NFC key",
		WatchCount:          2,
		IPhoneCount:         3,
		SupportInfo:         &supportInfo,
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