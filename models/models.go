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
	ID               string                 `json:"id"`
	CardTemplateID   string                 `json:"card_template_id"`
	EmployeeID       string                 `json:"employee_id"`
	OrganizationName string                 `json:"organization_name,omitempty"`
	CardNumber       string                 `json:"card_number"`
	SiteCode         string                 `json:"site_code,omitempty"`
	FullName         string                 `json:"full_name"`
	Email            string                 `json:"email"`
	PhoneNumber      string                 `json:"phone_number"`
	Classification   string                 `json:"classification"`
	Title            string                 `json:"title,omitempty"`
	StartDate        time.Time              `json:"start_date"`
	ExpirationDate   time.Time              `json:"expiration_date"`
	EmployeePhoto    string                 `json:"employee_photo"`
	State            string                 `json:"state"`
	URL              string                 `json:"url,omitempty"`
	InstallURL       string                 `json:"install_url"`
	Details          interface{}            `json:"details,omitempty"`
	FileData         string                 `json:"file_data,omitempty"`
	DirectInstallURL string                 `json:"direct_install_url,omitempty"`
	Temporary        bool                   `json:"temporary"`
	Devices          []Device               `json:"devices,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
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
	CardTemplateID         string                 `json:"card_template_id"`
	EmployeeID             string                 `json:"employee_id,omitempty"`
	TagID                  string                 `json:"tag_id,omitempty"`
	AllowOnMultipleDevices bool                   `json:"allow_on_multiple_devices,omitempty"`
	CardNumber             string                 `json:"card_number,omitempty"`
	SiteCode               string                 `json:"site_code,omitempty"`
	FullName               string                 `json:"full_name,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	PhoneNumber            string                 `json:"phone_number,omitempty"`
	Classification         string                 `json:"classification,omitempty"`
	Title                  string                 `json:"title,omitempty"`
	Department             string                 `json:"department,omitempty"`
	Location               string                 `json:"location,omitempty"`
	SiteName               string                 `json:"site_name,omitempty"`
	Workstation            string                 `json:"workstation,omitempty"`
	MailStop               string                 `json:"mail_stop,omitempty"`
	CompanyAddress         string                 `json:"company_address,omitempty"`
	StartDate              time.Time              `json:"start_date"`
	ExpirationDate         time.Time              `json:"expiration_date"`
	EmployeePhoto          string                 `json:"employee_photo,omitempty"`
	Temporary              bool                   `json:"temporary,omitempty"`
	Metadata               map[string]interface{} `json:"metadata,omitempty"`
}

// UpdateParams defines parameters for updating an existing card
type UpdateParams struct {
	CardID         string     `json:"card_id"`
	EmployeeID     string     `json:"employee_id,omitempty"`
	FullName       string     `json:"full_name,omitempty"`
	Email          string     `json:"email,omitempty"`
	PhoneNumber    string     `json:"phone_number,omitempty"`
	Classification string     `json:"classification,omitempty"`
	Title          string     `json:"title,omitempty"`
	Department     string     `json:"department,omitempty"`
	Location       string     `json:"location,omitempty"`
	SiteName       string     `json:"site_name,omitempty"`
	Workstation    string     `json:"workstation,omitempty"`
	MailStop       string     `json:"mail_stop,omitempty"`
	CompanyAddress string     `json:"company_address,omitempty"`
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
	Logo                   string                 `json:"logo,omitempty"`
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
	ID         interface{} `json:"id"`
	Event      string      `json:"event"`
	Type       string      `json:"type"`
	UserID     string      `json:"user_id"`
	CardID     string      `json:"card_id"`
	TemplateID string      `json:"template_id"`
	Device     string      `json:"device"`
	Timestamp  time.Time   `json:"timestamp"`
	CreatedAt  string      `json:"created_at"`
	Details    string      `json:"details"`
	IPAddress  string      `json:"ip_address"`
	UserAgent  string      `json:"user_agent"`
	Metadata   interface{} `json:"metadata"`
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
	ExID     string `json:"ex_id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// PassTemplatePair represents a paired iOS and Android template configuration
type PassTemplatePair struct {
	ID              string        `json:"id"`
	ExID            string        `json:"ex_id"`
	Name            string        `json:"name"`
	CreatedAt       time.Time     `json:"created_at"`
	IOSTemplate     *TemplateInfo `json:"ios_template"`
	AndroidTemplate *TemplateInfo `json:"android_template"`
}

// PassTemplatePairsResponse represents the response from listing pass template pairs.
// The upstream JSON key is "card_template_pairs"; the Go field name is preserved
// for backward compatibility.
type PassTemplatePairsResponse struct {
	PassTemplatePairs []PassTemplatePair `json:"card_template_pairs"`
	Pagination        Pagination         `json:"pagination"`
}

// ListPassTemplatePairsParams defines parameters for listing pass template pairs
type ListPassTemplatePairsParams struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// CreatePassTemplatePairParams defines parameters for creating a pass template pair.
// Both referenced card templates must be published (status: ready) and use the same
// protocol. AppleCardTemplateID must reference an Apple (iOS) template and
// GoogleCardTemplateID must reference a Google (Android) template.
type CreatePassTemplatePairParams struct {
	Name                 string `json:"name"`
	AppleCardTemplateID  string `json:"apple_card_template_id"`
	GoogleCardTemplateID string `json:"google_card_template_id"`
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
	ID                    string                  `json:"id"`
	FullName              string                  `json:"full_name"`
	State                 string                  `json:"state"`
	Metadata              map[string]interface{}  `json:"metadata"`
	UnifiedAccessPassExID string                  `json:"unified_access_pass_ex_id"`
	PassTemplate          *LedgerItemPassTemplate `json:"pass_template,omitempty"`
}

// LedgerItem represents a billing ledger item
type LedgerItem struct {
	CreatedAt  time.Time              `json:"created_at"`
	Amount     interface{}            `json:"amount"`
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

// IosPreflight represents an iOS In-App Provisioning preflight response
type IosPreflight struct {
	ProvisioningCredentialIdentifier string `json:"provisioningCredentialIdentifier"`
	SharingInstanceIdentifier        string `json:"sharingInstanceIdentifier"`
	CardTemplateIdentifier           string `json:"cardTemplateIdentifier"`
	EnvironmentIdentifier            string `json:"environmentIdentifier"`
}

// IosPreflightParams defines parameters for iOS preflight
type IosPreflightParams struct {
	CardTemplateID string `json:"card_template_id"`
	AccessPassExID string `json:"access_pass_ex_id"`
}

// Webhook represents a webhook configuration
type Webhook struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	URL              string   `json:"url"`
	AuthMethod       string   `json:"auth_method"`
	SubscribedEvents []string `json:"subscribed_events"`
	CreatedAt        string   `json:"created_at"`
	PrivateKey       string   `json:"private_key,omitempty"`
	ClientCert       string   `json:"client_cert,omitempty"`
	CertExpiresAt    string   `json:"cert_expires_at,omitempty"`
}

// WebhooksResponse represents the response from listing webhooks
type WebhooksResponse struct {
	Webhooks   []Webhook  `json:"webhooks"`
	Pagination Pagination `json:"pagination"`
}

// CreateWebhookParams defines parameters for creating a webhook
type CreateWebhookParams struct {
	Name             string   `json:"name"`
	URL              string   `json:"url"`
	SubscribedEvents []string `json:"subscribed_events"`
	AuthMethod       string   `json:"auth_method,omitempty"`
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

// LandingPage represents a landing page configuration
type LandingPage struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Kind              string `json:"kind"`
	CreatedAt         string `json:"created_at"`
	PasswordProtected bool   `json:"password_protected"`
	LogoURL           string `json:"logo_url,omitempty"`
}

// CreateLandingPageParams defines parameters for creating a landing page
type CreateLandingPageParams struct {
	Name                   string `json:"name"`
	Kind                   string `json:"kind"`
	AdditionalText         string `json:"additional_text,omitempty"`
	BgColor                string `json:"bg_color,omitempty"`
	AllowImmediateDownload bool   `json:"allow_immediate_download,omitempty"`
	Password               string `json:"password,omitempty"`
	Is2FAEnabled           bool   `json:"is_2fa_enabled,omitempty"`
	Logo                   string `json:"logo,omitempty"`
}

// UpdateLandingPageParams defines parameters for updating a landing page
type UpdateLandingPageParams struct {
	LandingPageID          string `json:"landing_page_id"`
	Name                   string `json:"name,omitempty"`
	AdditionalText         string `json:"additional_text,omitempty"`
	BgColor                string `json:"bg_color,omitempty"`
	AllowImmediateDownload *bool  `json:"allow_immediate_download,omitempty"`
	Password               string `json:"password,omitempty"`
	Is2FAEnabled           *bool  `json:"is_2fa_enabled,omitempty"`
	Logo                   string `json:"logo,omitempty"`
}

// CredentialProfile represents a credential profile
type CredentialProfile struct {
	ID          string                 `json:"id"`
	AID         string                 `json:"aid"`
	Name        string                 `json:"name"`
	AppleID     string                 `json:"apple_id,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	CardStorage interface{} `json:"card_storage,omitempty"`
	Keys        []interface{}          `json:"keys,omitempty"`
	Files       []interface{}          `json:"files,omitempty"`
}

// KeyParam represents a key parameter for credential profile creation
type KeyParam struct {
	Value string `json:"value"`
}

// CreateCredentialProfileParams defines parameters for creating a credential profile
type CreateCredentialProfileParams struct {
	Name    string     `json:"name"`
	AppName string     `json:"app_name"`
	Keys    []KeyParam `json:"keys,omitempty"`
	FileID  string     `json:"file_id,omitempty"`
}
