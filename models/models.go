package models

import "time"

// Union is an interface representing the base type for access pass responses.
// Both Card and UnifiedAccessPass implement this interface.
type Union interface {
	GetID() string
	GetURL() string
	GetState() string
	isUnion()
}

// Device represents a device associated with an access pass
type Device struct {
	ID         string    `json:"id"`
	Platform   string    `json:"platform"`
	DeviceType string    `json:"device_type"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Card represents an NFC key or access pass
type Card struct {
	ID                    string                 `json:"id"`
	CardTemplateID        string                 `json:"card_template_id"`
	EmployeeID            string                 `json:"employee_id"`
	CardNumber            string                 `json:"card_number"`
	SiteCode              string                 `json:"site_code,omitempty"`
	FullName              string                 `json:"full_name"`
	Email                 string                 `json:"email"`
	PhoneNumber           string                 `json:"phone_number"`
	Classification        string                 `json:"classification"`
	StartDate             time.Time              `json:"start_date"`
	ExpirationDate        time.Time              `json:"expiration_date"`
	EmployeePhoto         string                 `json:"employee_photo"`
	State                 string                 `json:"state"`
	URL                   string                 `json:"install_url"`
	Details               interface{}            `json:"details,omitempty"`
	FileData              string                 `json:"file_data,omitempty"`
	DirectInstallURL      string                 `json:"direct_install_url,omitempty"`
	Devices               []Device               `json:"devices,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// CardProvisionResponse represents the response from provisioning a card
type CardProvisionResponse struct {
	ID               string    `json:"id"`
	CardTemplateID   string    `json:"card_template_id"`
	EmployeeID       string    `json:"employee_id"`
	CardNumber       string    `json:"card_number"`
	SiteCode         string    `json:"site_code,omitempty"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	Classification   string    `json:"classification"`
	StartDate        time.Time `json:"start_date"`
	ExpirationDate   time.Time `json:"expiration_date"`
	EmployeePhoto    string    `json:"employee_photo"`
	State            string    `json:"state"`
	URL              string    `json:"install_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DirectInstallUrl string    `json:"direct_install_url"`
	Details          []Card    `json:"details"`
}

// ProvisionParams defines parameters for provisioning a new card
type ProvisionParams struct {
	CardTemplateID string    `json:"card_template_id"`
	EmployeeID     string    `json:"employee_id"`
	CardNumber     string    `json:"card_number"`
	SiteCode       string    `json:"site_code,omitempty"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Classification string    `json:"classification"`
	Title          string    `json:"title,omitempty"`
	StartDate      time.Time `json:"start_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	EmployeePhoto  string    `json:"employee_photo"`
}

// UpdateParams defines parameters for updating an existing card
type UpdateParams struct {
	CardID         string     `json:"card_id"`
	EmployeeID     string     `json:"employee_id,omitempty"`
	FullName       string     `json:"full_name,omitempty"`
	Email          string     `json:"email,omitempty"`
	PhoneNumber    string     `json:"phone_number,omitempty"`
	Classification string     `json:"classification,omitempty"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	EmployeePhoto  string     `json:"employee_photo,omitempty"`
}

// ListKeysParams defines parameters for filtering cards
type ListKeysParams struct {
	TemplateID string `json:"card_template_id,omitempty"`
	State      string `json:"state,omitempty"`
	EmployeeID string `json:"employee_id,omitempty"`
	CardNumber string `json:"card_number,omitempty"`
	SiteCode   string `json:"site_code,omitempty"`
}

// Template represents a card template
type Template struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Platform    string         `json:"platform"`
	UseCase     string         `json:"use_case"`
	Protocol    string         `json:"protocol"`
	WatchCount  int            `json:"watch_count"`
	IPhoneCount int            `json:"iphone_count"`
	Design      TemplateDesign `json:"design"`
	SupportInfo SupportInfo    `json:"support_info"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
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
	Name        string         `json:"name"`
	Platform    string         `json:"platform"`
	UseCase     string         `json:"use_case"`
	Protocol    string         `json:"protocol"`
	WatchCount  int            `json:"watch_count"`
	IPhoneCount int            `json:"iphone_count"`
	Design      TemplateDesign `json:"design"`
	SupportInfo SupportInfo    `json:"support_info"`
}

// UpdateTemplateParams defines parameters for updating an existing template
type UpdateTemplateParams struct {
	CardTemplateID string          `json:"card_template_id"`
	Name           string          `json:"name,omitempty"`
	WatchCount     int             `json:"watch_count,omitempty"`
	IPhoneCount    int             `json:"iphone_count,omitempty"`
	Design         *TemplateDesign `json:"design,omitempty"`
	SupportInfo    *SupportInfo    `json:"support_info,omitempty"`
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

type UnifiedAccessPass struct {
	ID      string `json:"id"`
	URL     string `json:"install_url"`
	State   string `json:"state"`
	Status  string `json:"status"`
	Details []Card `json:"details"`
}

func (u *UnifiedAccessPass) GetID() string    { return u.ID }
func (u *UnifiedAccessPass) GetURL() string   { return u.URL }
func (u *UnifiedAccessPass) GetState() string { return u.State }
func (u *UnifiedAccessPass) isUnion()         {}

func (c *Card) GetID() string    { return c.ID }
func (c *Card) GetURL() string   { return c.URL }
func (c *Card) GetState() string { return c.State }
func (c *Card) isUnion()         {}
