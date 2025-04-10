package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/access_grid/accessgrid-go/client"
	"github.com/access_grid/accessgrid-go/models"
)

func setupAccessCardsTestServer() (*httptest.Server, *AccessCardsService) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		switch r.URL.Path {
		case "/cards":
			if r.Method == http.MethodPost {
				// Provision
				w.Write([]byte(`{
					"id": "0xc4rd1d",
					"card_template_id": "0xd3adb00b5",
					"full_name": "Employee name",
					"state": "active",
					"url": "https://accessgrid.com/install/0xc4rd1d"
				}`))
			} else if r.Method == http.MethodGet {
				// List
				w.Write([]byte(`[
					{
						"id": "0xc4rd1d",
						"card_template_id": "0xd3adb00b5",
						"full_name": "Employee name",
						"state": "active"
					}
				]`))
			}
		case "/cards/0xc4rd1d":
			if r.Method == http.MethodPut {
				// Update
				w.Write([]byte(`{
					"id": "0xc4rd1d",
					"card_template_id": "0xd3adb00b5",
					"full_name": "Updated Employee Name",
					"state": "active"
				}`))
			} else if r.Method == http.MethodDelete {
				// Delete
				w.Write([]byte(`{}`))
			}
		case "/cards/0xc4rd1d/suspend":
			// Suspend
			w.Write([]byte(`{}`))
		case "/cards/0xc4rd1d/resume":
			// Resume
			w.Write([]byte(`{}`))
		case "/cards/0xc4rd1d/unlink":
			// Unlink
			w.Write([]byte(`{}`))
		}
	}))

	c, _ := client.NewClient("test-account", "test-secret", client.WithBaseURL(server.URL))
	service := NewAccessCardsService(c)

	return server, service
}

func TestAccessCardsService_Provision(t *testing.T) {
	server, service := setupAccessCardsTestServer()
	defer server.Close()

	params := models.ProvisionParams{
		CardTemplateID:         "0xd3adb00b5",
		EmployeeID:             "123456789",
		TagID:                  "DDEADB33FB00B5",
		AllowOnMultipleDevices: true,
		FullName:              "Employee name",
		Email:                 "employee@example.com",
		PhoneNumber:           "+19547212241",
		Classification:        "full_time",
		StartDate:            "2023-01-01T00:00:00Z",
		ExpirationDate:       "2025-01-01T00:00:00Z",
	}

	card, err := service.Provision(params)
	if err != nil {
		t.Fatalf("Provision() error = %v", err)
	}

	if card.ID != "0xc4rd1d" {
		t.Errorf("Provision() card.ID = %v, want %v", card.ID, "0xc4rd1d")
	}
	if card.CardTemplateID != "0xd3adb00b5" {
		t.Errorf("Provision() card.CardTemplateID = %v, want %v", card.CardTemplateID, "0xd3adb00b5")
	}
	if card.FullName != "Employee name" {
		t.Errorf("Provision() card.FullName = %v, want %v", card.FullName, "Employee name")
	}
	if card.URL != "https://accessgrid.com/install/0xc4rd1d" {
		t.Errorf("Provision() card.URL = %v, want %v", card.URL, "https://accessgrid.com/install/0xc4rd1d")
	}
}

func TestAccessCardsService_Update(t *testing.T) {
	server, service := setupAccessCardsTestServer()
	defer server.Close()

	params := models.UpdateParams{
		CardID:         "0xc4rd1d",
		EmployeeID:     "987654321",
		FullName:       "Updated Employee Name",
		Classification: "contractor",
	}

	card, err := service.Update(params)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if card.ID != "0xc4rd1d" {
		t.Errorf("Update() card.ID = %v, want %v", card.ID, "0xc4rd1d")
	}
	if card.FullName != "Updated Employee Name" {
		t.Errorf("Update() card.FullName = %v, want %v", card.FullName, "Updated Employee Name")
	}
}

func TestAccessCardsService_List(t *testing.T) {
	server, service := setupAccessCardsTestServer()
	defer server.Close()

	params := &models.ListKeysParams{
		TemplateID: "0xd3adb00b5",
	}

	cards, err := service.List(params)
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if len(cards) != 1 {
		t.Fatalf("List() got %v cards, want %v", len(cards), 1)
	}

	if cards[0].ID != "0xc4rd1d" {
		t.Errorf("List() cards[0].ID = %v, want %v", cards[0].ID, "0xc4rd1d")
	}
}

func TestAccessCardsService_CardStateOperations(t *testing.T) {
	server, service := setupAccessCardsTestServer()
	defer server.Close()

	// Test Suspend
	err := service.Suspend("0xc4rd1d")
	if err != nil {
		t.Errorf("Suspend() error = %v", err)
	}

	// Test Resume
	err = service.Resume("0xc4rd1d")
	if err != nil {
		t.Errorf("Resume() error = %v", err)
	}

	// Test Unlink
	err = service.Unlink("0xc4rd1d")
	if err != nil {
		t.Errorf("Unlink() error = %v", err)
	}

	// Test Delete
	err = service.Delete("0xc4rd1d")
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}