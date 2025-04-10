package models

import "time"

// Card represents an NFC key or access pass
type Card struct {
	ID                    string    `json:"id"`
	CardTemplateID        string    `json:"card_template_id"`
	EmployeeID            string    `json:"employee_id"`
	TagID                 string    `json:"tag_id"`
	AllowOnMultipleDevices bool      `json:"allow_on_multiple_devices"`
	FullName              string    `json:"full_name"`
	Email                 string    `json:"email"`
	PhoneNumber           string    `json:"phone_number"`
	Classification        string    `json:"classification"`
	StartDate             time.Time `json:"start_date"`
	ExpirationDate        time.Time `json:"expiration_date"`
	EmployeePhoto         string    `json:"employee_photo"`
	State                 string    `json:"state"`
	URL                   string    `json:"url"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// ProvisionParams defines parameters for provisioning a new card
type ProvisionParams struct {
	CardTemplateID          string    `json:"card_template_id"`
	EmployeeID              string    `json:"employee_id"`
	TagID                   string    `json:"tag_id"`
	AllowOnMultipleDevices  bool      `json:"allow_on_multiple_devices"`
	FullName                string    `json:"full_name"`
	Email                   string    `json:"email"`
	PhoneNumber             string    `json:"phone_number"`
	Classification          string    `json:"classification"`
	StartDate               time.Time `json:"start_date"`
	ExpirationDate          time.Time `json:"expiration_date"`
	EmployeePhoto           string    `json:"employee_photo"`
}

// UpdateParams defines parameters for updating an existing card
type UpdateParams struct {
	CardID          string     `json:"card_id"`
	EmployeeID      string     `json:"employee_id,omitempty"`
	FullName        string     `json:"full_name,omitempty"`
	Email           string     `json:"email,omitempty"`
	PhoneNumber     string     `json:"phone_number,omitempty"`
	Classification  string     `json:"classification,omitempty"`
	ExpirationDate  *time.Time `json:"expiration_date,omitempty"`
	EmployeePhoto   string     `json:"employee_photo,omitempty"`
}

// ListKeysParams defines parameters for filtering cards
type ListKeysParams struct {
	TemplateID string `json:"template_id,omitempty"`
	State      string `json:"state,omitempty"`
	EmployeeID string `json:"employee_id,omitempty"`
	TagID      string `json:"tag_id,omitempty"`
}

// Template represents a card template
type Template struct {
	ID                    string       `json:"id"`
	Name                  string       `json:"name"`
	Platform              string       `json:"platform"`
	UseCase               string       `json:"use_case"`
	Protocol              string       `json:"protocol"`
	AllowOnMultipleDevices bool         `json:"allow_on_multiple_devices"`
	WatchCount            int          `json:"watch_count"`
	IPhoneCount           int          `json:"iphone_count"`
	Design                TemplateDesign `json:"design"`
	SupportInfo           SupportInfo  `json:"support_info"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
}

// TemplateDesign represents the design elements of a card template
type TemplateDesign struct {
	BackgroundColor     string `json:"background_color"`
	LabelColor          string `json:"label_color"`
	LabelSecondaryColor string `json:"label_secondary_color"`
	BackgroundImage     string `json:"background_image"`
	LogoImage           string `json:"logo_image"`
	IconImage           string `json:"icon_image"`
}

// SupportInfo represents support information for a card template
type SupportInfo struct {
	SupportURL            string `json:"support_url"`
	SupportPhoneNumber    string `json:"support_phone_number"`
	SupportEmail          string `json:"support_email"`
	PrivacyPolicyURL      string `json:"privacy_policy_url"`
	TermsAndConditionsURL string `json:"terms_and_conditions_url"`
}

// CreateTemplateParams defines parameters for creating a new template
type CreateTemplateParams struct {
	Name                  string        `json:"name"`
	Platform              string        `json:"platform"`
	UseCase               string        `json:"use_case"`
	Protocol              string        `json:"protocol"`
	AllowOnMultipleDevices bool          `json:"allow_on_multiple_devices"`
	WatchCount            int           `json:"watch_count"`
	IPhoneCount           int           `json:"iphone_count"`
	Design                TemplateDesign `json:"design"`
	SupportInfo           SupportInfo   `json:"support_info"`
}

// UpdateTemplateParams defines parameters for updating an existing template
type UpdateTemplateParams struct {
	CardTemplateID         string       `json:"card_template_id"`
	Name                   string       `json:"name,omitempty"`
	AllowOnMultipleDevices bool         `json:"allow_on_multiple_devices,omitempty"`
	WatchCount             int          `json:"watch_count,omitempty"`
	IPhoneCount            int          `json:"iphone_count,omitempty"`
	Design                 *TemplateDesign `json:"design,omitempty"`
	SupportInfo            *SupportInfo `json:"support_info,omitempty"`
}

// EventLogFilters defines parameters for filtering event logs
type EventLogFilters struct {
	Device    string     `json:"device,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	EventType string     `json:"event_type,omitempty"`
}

// Event represents an event in the event log
type Event struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	UserID     string    `json:"user_id"`
	CardID     string    `json:"card_id"`
	TemplateID string    `json:"template_id"`
	Device     string    `json:"device"`
	Timestamp  time.Time `json:"timestamp"`
	Details    string    `json:"details"`
}