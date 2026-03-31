package models

import "time"

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
	InstallURL            string                 `json:"install_url"`
	Details               interface{}            `json:"details,omitempty"`
	FileData              string                 `json:"file_data,omitempty"`
	DirectInstallURL      string                 `json:"direct_install_url,omitempty"`
	Temporary             bool                   `json:"temporary"`
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
	Temporary        bool      `json:"temporary"`
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
	Temporary      bool      `json:"temporary,omitempty"`
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
	TemplateID string `json:"template_id,omitempty"`
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
	Name                   string                 `json:"name"`
	Platform               string                 `json:"platform"`
	UseCase                string                 `json:"use_case"`
	Protocol               string                 `json:"protocol"`
	AllowOnMultipleDevices bool                   `json:"allow_on_multiple_devices,omitempty"`
	WatchCount             int                    `json:"watch_count"`
	IPhoneCount            int                    `json:"iphone_count"`
	BackgroundColor        string                 `json:"background_color,omitempty"`
	LabelColor             string                 `json:"label_color,omitempty"`
	LabelSecondaryColor    string                 `json:"label_secondary_color,omitempty"`
	BackgroundImage        string                 `json:"background_image,omitempty"`
	LogoImage              string                 `json:"logo_image,omitempty"`
	IconImage              string                 `json:"icon_image,omitempty"`
	SupportURL             string                 `json:"support_url,omitempty"`
	SupportPhoneNumber     string                 `json:"support_phone_number,omitempty"`
	SupportEmail           string                 `json:"support_email,omitempty"`
	PrivacyPolicyURL       string                 `json:"privacy_policy_url,omitempty"`
	TermsAndConditionsURL  string                 `json:"terms_and_conditions_url,omitempty"`
	Metadata               map[string]interface{} `json:"metadata,omitempty"`
}

// UpdateTemplateParams defines parameters for updating an existing template
type UpdateTemplateParams struct {
	CardTemplateID         string                 `json:"card_template_id"`
	Name                   string                 `json:"name,omitempty"`
	AllowOnMultipleDevices *bool                  `json:"allow_on_multiple_devices,omitempty"`
	WatchCount             int                    `json:"watch_count,omitempty"`
	IPhoneCount            int                    `json:"iphone_count,omitempty"`
	BackgroundColor        string                 `json:"background_color,omitempty"`
	LabelColor             string                 `json:"label_color,omitempty"`
	LabelSecondaryColor    string                 `json:"label_secondary_color,omitempty"`
	SupportURL             string                 `json:"support_url,omitempty"`
	SupportPhoneNumber     string                 `json:"support_phone_number,omitempty"`
	SupportEmail           string                 `json:"support_email,omitempty"`
	PrivacyPolicyURL       string                 `json:"privacy_policy_url,omitempty"`
	TermsAndConditionsURL  string                 `json:"terms_and_conditions_url,omitempty"`
	Metadata               map[string]interface{} `json:"metadata,omitempty"`
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

// Pagination represents pagination metadata in list responses
type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page,omitempty"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count,omitempty"`
}

// TemplateInfo represents minimal template info within a PassTemplatePair
type TemplateInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// PassTemplatePair represents a paired iOS and Android template configuration
type PassTemplatePair struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	CreatedAt       time.Time     `json:"created_at"`
	IOSTemplate     *TemplateInfo `json:"ios_template"`
	AndroidTemplate *TemplateInfo `json:"android_template"`
}

// PassTemplatePairsResponse represents the response from listing pass template pairs
type PassTemplatePairsResponse struct {
	PassTemplatePairs []PassTemplatePair `json:"pass_template_pairs"`
	Pagination        Pagination         `json:"pagination"`
}

// ListPassTemplatePairsParams defines parameters for listing pass template pairs
type ListPassTemplatePairsParams struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// LedgerItemPassTemplate represents a pass template reference within a ledger item's access pass
type LedgerItemPassTemplate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Platform string `json:"platform"`
	UseCase  string `json:"use_case"`
}

// LedgerItemAccessPass represents an access pass reference within a ledger item
type LedgerItemAccessPass struct {
	ID                     string                  `json:"id"`
	FullName               string                  `json:"full_name"`
	State                  string                  `json:"state"`
	Metadata               map[string]interface{}  `json:"metadata"`
	UnifiedAccessPassExID  string                  `json:"unified_access_pass_ex_id"`
	PassTemplate           *LedgerItemPassTemplate `json:"pass_template,omitempty"`
}

// LedgerItem represents a billing ledger item
type LedgerItem struct {
	CreatedAt  time.Time              `json:"created_at"`
	Amount     float64                `json:"amount"`
	ID         string                 `json:"id"`
	Kind       string                 `json:"kind"`
	Metadata   map[string]interface{} `json:"metadata"`
	AccessPass *LedgerItemAccessPass  `json:"access_pass"`
}

// LedgerItemsResponse represents the response from listing ledger items
type LedgerItemsResponse struct {
	LedgerItems []LedgerItem `json:"ledger_items"`
	Pagination  Pagination   `json:"pagination"`
}

// ListLedgerItemsParams defines parameters for listing ledger items
type ListLedgerItemsParams struct {
	Page      int        `json:"page,omitempty"`
	PerPage   int        `json:"per_page,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// HIDOrg represents an HID organization
type HIDOrg struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	FullAddress string `json:"full_address"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

// CreateHIDOrgParams defines parameters for creating an HID organization
type CreateHIDOrgParams struct {
	Name        string `json:"name"`
	FullAddress string `json:"full_address"`
	Phone       string `json:"phone"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

// CompleteHIDOrgParams defines parameters for completing HID org registration
type CompleteHIDOrgParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
