# AccessGrid SDK

A Go SDK for interacting with the [AccessGrid.com](https://www.accessgrid.com) API. This SDK provides a simple interface for managing NFC key cards and enterprise templates. Full docs at https://www.accessgrid.com/docs

## Installation

```bash
go get github.com/Access-Grid/accessgrid-go
```

## Quick Start

```go
package main

import (\n    "context"
    "fmt"
    "os"
    "github.com/Access-Grid/accessgrid-go"
)

func main() {
    accountID := os.Getenv("ACCOUNT_ID")
    secretKey := os.Getenv("SECRET_KEY")

    client, err := accessgrid.NewClient(accountID, secretKey)
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }
}
```

## API Reference

### Access Cards

#### Provision a new card

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/Access-Grid/accessgrid-go"
)

func main() {
    accountID := os.Getenv("ACCOUNT_ID")
    secretKey := os.Getenv("SECRET_KEY")

    client, err := accessgrid.NewClient(accountID, secretKey)
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }

    params := accessgrid.ProvisionParams{
        CardTemplateID:         "0xd3adb00b5",
        EmployeeID:             "123456789",
        TagID:                  "DDEADB33FB00B5",
        AllowOnMultipleDevices: true,
        FullName:               "Employee name",
        Email:                  "employee@yourwebsite.com",
        PhoneNumber:            "+19547212241",
        Classification:         "full_time",
        Department:             "Engineering",
        Location:               "San Francisco",
        SiteName:               "HQ Building A",
        Workstation:            "4F-207",
        MailStop:               "MS-401",
        CompanyAddress:         "123 Main St, San Francisco, CA 94105",
        StartDate:              time.Now().UTC(),
        ExpirationDate:         time.Now().UTC().AddDate(0, 3, 0),
        EmployeePhoto:          "[image_in_base64_encoded_format]",
        Title:                  "Engineering Manager",
        Metadata: map[string]interface{}{
            "department": "engineering",
            "badge_type": "contractor",
        },
    }

    ctx := context.Background()
    card, err := client.AccessCards.Provision(ctx, params)
    if err != nil {
        fmt.Printf("Error provisioning card: %v\n", err)
        return
    }

    fmt.Printf("Install URL: %s\n", card.URL)
}
```

#### Get a card

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/Access-Grid/accessgrid-go"
)

func main() {
    accountID := os.Getenv("ACCOUNT_ID")
    secretKey := os.Getenv("SECRET_KEY")

    client, err := accessgrid.NewClient(accountID, secretKey)
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }

    ctx := context.Background()
    card, err := client.AccessCards.Get(ctx, "0xc4rd1d")
    if err != nil {
        fmt.Printf("Error retrieving card: %v\n", err)
        return
    }

    fmt.Printf("Card ID: %s\n", card.ID)
    fmt.Printf("State: %s\n", card.State)
    fmt.Printf("Full Name: %s\n", card.FullName)
    fmt.Printf("Install URL: %s\n", card.InstallURL)
    fmt.Printf("Expiration Date: %s\n", card.ExpirationDate)
    fmt.Printf("Card Number: %s\n", card.CardNumber)
    fmt.Printf("Site Code: %s\n", card.SiteCode)
    fmt.Printf("Devices: %d\n", len(card.Devices))
    fmt.Printf("Metadata: %v\n", card.Metadata)
}
```

#### Update a card

```go
package main

import (\n    "context"
   "fmt"
   "os"
   "time"
   "github.com/Access-Grid/accessgrid-go"
)

func main() {
   accountID := os.Getenv("ACCOUNT_ID")
   secretKey := os.Getenv("SECRET_KEY")

   client, err := accessgrid.NewClient(accountID, secretKey)
   if err != nil {
       fmt.Printf("Error creating client: %v\n", err)
       return
   }

   params := accessgrid.UpdateParams{
       CardID:         "0xc4rd1d",
       EmployeeID:     "987654321",
       FullName:       "Updated Employee Name",
       Classification: "contractor",
       ExpirationDate: &time.Time{}, // In actual code: expirationDate := time.Now().UTC().AddDate(0, 3, 0); params.ExpirationDate = &expirationDate
       EmployeePhoto:  "[image_in_base64_encoded_format]",
   }

   ctx := context.Background()
   card, err := client.AccessCards.Update(ctx, params)
   if err != nil {
       fmt.Printf("Error updating card: %v\n", err)
       return
   }

   fmt.Println("Card updated successfully")
}
```

#### List NFC keys / Access passes

```go
package main

import (\n    "context"
    "fmt"
    "os"
    "github.com/Access-Grid/accessgrid-go"
)

func main() {
    accountID := os.Getenv("ACCOUNT_ID")
    secretKey := os.Getenv("SECRET_KEY")

    client, err := accessgrid.NewClient(accountID, secretKey)
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }

    // Get filtered keys by template
    templateFilter := accessgrid.ListKeysParams{
        TemplateID: "0xd3adb00b5",
    }
    ctx := context.Background()
    templateKeys, err := client.AccessCards.List(ctx, &templateFilter)
    if err != nil {
        fmt.Printf("Error listing cards: %v\n", err)
        return
    }

    // Get filtered keys by state
    stateFilter := accessgrid.ListKeysParams{
        State: "active",
    }
    activeKeys, err := client.AccessCards.List(ctx, &stateFilter)
    if err != nil {
        fmt.Printf("Error listing cards: %v\n", err)
        return
    }

    // Print keys
    for _, key := range templateKeys {
        fmt.Printf("Key ID: %s, Name: %s, State: %s\n", key.ID, key.FullName, key.State)
    }
}
```

#### Manage card states

```go
// Suspend a card
ctx := context.Background()
err = client.AccessCards.Suspend(ctx, "0xc4rd1d")
if err != nil {
    fmt.Printf("Error suspending card: %v\n", err)
    return
}

// Resume a card
err = client.AccessCards.Resume(ctx, "0xc4rd1d")
if err != nil {
    fmt.Printf("Error resuming card: %v\n", err)
    return
}

// Unlink a card
err = client.AccessCards.Unlink(ctx, "0xc4rd1d")
if err != nil {
    fmt.Printf("Error unlinking card: %v\n", err)
    return
}

// Delete a card
err = client.AccessCards.Delete(ctx, "0xc4rd1d")
if err != nil {
    fmt.Printf("Error deleting card: %v\n", err)
    return
}
```

### Enterprise Console

#### Create a template

```go
package main

import (\n    "context"
   "fmt"
   "os"
   "github.com/Access-Grid/accessgrid-go"
)

func main() {
   accountID := os.Getenv("ACCOUNT_ID")
   secretKey := os.Getenv("SECRET_KEY")

   client, err := accessgrid.NewClient(accountID, secretKey)
   if err != nil {
       fmt.Printf("Error creating client: %v\n", err)
       return
   }

   params := accessgrid.CreateTemplateParams{
       Name:                   "Employee Access Pass",
       Platform:               "apple",
       UseCase:                "employee_badge",
       Protocol:               "desfire",
       AllowOnMultipleDevices: true,
       WatchCount:             2,
       IPhoneCount:            3,
       BackgroundColor:        "#FFFFFF",
       LabelColor:             "#000000",
       LabelSecondaryColor:    "#333333",
       SupportURL:             "https://help.yourcompany.com",
       SupportPhoneNumber:     "+1-555-123-4567",
       SupportEmail:           "support@yourcompany.com",
       PrivacyPolicyURL:       "https://yourcompany.com/privacy",
       TermsAndConditionsURL:  "https://yourcompany.com/terms",
       Metadata: map[string]interface{}{
           "version":         "2.1",
           "approval_status": "approved",
       },
   }

   ctx := context.Background()
   template, err := client.Console.CreateTemplate(ctx, params)
   if err != nil {
       fmt.Printf("Error creating template: %v\n", err)
       return
   }

   fmt.Printf("Template created successfully: %s\n", template.ID)
}
```

#### Update a template

```go
package main

import (\n    "context"
   "fmt"
   "os"
   "github.com/Access-Grid/accessgrid-go"
)

func main() {
   accountID := os.Getenv("ACCOUNT_ID")
   secretKey := os.Getenv("SECRET_KEY")

   client, err := accessgrid.NewClient(accountID, secretKey)
   if err != nil {
       fmt.Printf("Error creating client: %v\n", err)
       return
   }

   allowMulti := true
   params := accessgrid.UpdateTemplateParams{
       CardTemplateID:         "0xd3adb00b5",
       Name:                   "Updated Employee Access Pass",
       AllowOnMultipleDevices: &allowMulti,
       WatchCount:             2,
       IPhoneCount:            3,
       BackgroundColor:        "#FFFFFF",
       LabelColor:             "#000000",
       LabelSecondaryColor:    "#333333",
       SupportURL:             "https://help.yourcompany.com",
       SupportPhoneNumber:     "+1-555-123-4567",
       SupportEmail:           "support@yourcompany.com",
       PrivacyPolicyURL:       "https://yourcompany.com/privacy",
       TermsAndConditionsURL:  "https://yourcompany.com/terms",
   }

   ctx := context.Background()
   template, err := client.Console.UpdateTemplate(ctx, params)
   if err != nil {
       fmt.Printf("Error updating template: %v\n", err)
       return
   }

   fmt.Printf("Template updated successfully: %s\n", template.ID)
}
```

#### Read a template

```go
package main

import (\n    "context"
   "fmt"
   "os"
   "github.com/Access-Grid/accessgrid-go"
)

func main() {
   accountID := os.Getenv("ACCOUNT_ID")
   secretKey := os.Getenv("SECRET_KEY")

   client, err := accessgrid.NewClient(accountID, secretKey)
   if err != nil {
       fmt.Printf("Error creating client: %v\n", err)
       return
   }

   ctx := context.Background()
   template, err := client.Console.ReadTemplate(ctx, "0xd3adb00b5")
   if err != nil {
       fmt.Printf("Error reading template: %v\n", err)
       return
   }

   fmt.Printf("Template ID: %s\n", template.ID)
   fmt.Printf("Name: %s\n", template.Name)
   fmt.Printf("Platform: %s\n", template.Platform)
   fmt.Printf("Protocol: %s\n", template.Protocol)
   fmt.Printf("Multi-device: %v\n", template.AllowOnMultipleDevices)
}
```

#### Get event logs

```go
package main

import (\n    "context"
   "fmt"
   "os"
   "time"
   "github.com/Access-Grid/accessgrid-go"
)

func main() {
   accountID := os.Getenv("ACCOUNT_ID")
   secretKey := os.Getenv("SECRET_KEY")

   client, err := accessgrid.NewClient(accountID, secretKey)
   if err != nil {
       fmt.Printf("Error creating client: %v\n", err)
       return
   }

   startDate := time.Now().AddDate(0, 0, -30).UTC()
   endDate := time.Now().UTC()
   filters := accessgrid.EventLogFilters{
       Device:    "mobile",
       StartDate: &startDate,
       EndDate:   &endDate,
       EventType: "install",
   }

   ctx := context.Background()
   events, err := client.Console.EventLog(ctx, "0xd3adb00b5", filters)
   if err != nil {
       fmt.Printf("Error fetching event log: %v\n", err)
       return
   }

   for _, event := range events {
       fmt.Printf("Event: %s at %s by %s\n", event.Type, event.Timestamp, event.UserID)
   }
}
```

### HID Organizations

#### Create an HID org

```go
ctx := context.Background()
org, err := client.Console.HID.Orgs.Create(ctx, &accessgrid.CreateHIDOrgParams{
    Name:        "My Org",
    FullAddress: "1 Main St, NY NY",
    Phone:       "+1-555-0000",
    FirstName:   "Ada",
    LastName:    "Lovelace",
})
if err != nil {
    fmt.Printf("Error creating org: %v\n", err)
    return
}

fmt.Printf("Created org: %s (ID: %s)\n", org.Name, org.ID)
fmt.Printf("Slug: %s\n", org.Slug)
```

#### List HID orgs

```go
ctx := context.Background()
orgs, err := client.Console.HID.Orgs.List(ctx)
if err != nil {
    fmt.Printf("Error listing orgs: %v\n", err)
    return
}

for _, org := range orgs {
    fmt.Printf("Org ID: %s, Name: %s, Slug: %s\n", org.ID, org.Name, org.Slug)
}
```

#### Activate an HID org

```go
ctx := context.Background()
result, err := client.Console.HID.Orgs.Activate(ctx, &accessgrid.CompleteHIDOrgParams{
    Email:    "admin@example.com",
    Password: "hid-password-123",
})
if err != nil {
    fmt.Printf("Error completing registration: %v\n", err)
    return
}

fmt.Printf("Completed registration for org: %s\n", result.Name)
fmt.Printf("Status: %s\n", result.Status)
```

### Landing Pages

#### List landing pages

```go
ctx := context.Background()
landingPages, err := client.Console.ListLandingPages(ctx)
if err != nil {
    fmt.Printf("Error listing landing pages: %v\n", err)
    return
}

for _, page := range landingPages {
    fmt.Printf("ID: %s, Name: %s, Kind: %s\n", page.ID, page.Name, page.Kind)
    fmt.Printf("  Password Protected: %v\n", page.PasswordProtected)
    if page.LogoURL != "" {
        fmt.Printf("  Logo URL: %s\n", page.LogoURL)
    }
}
```

#### Create a landing page

```go
ctx := context.Background()
params := accessgrid.CreateLandingPageParams{
    Name:                   "Miami Office Access Pass",
    Kind:                   "universal",
    AdditionalText:         "Welcome to the Miami Office",
    BgColor:                "#f1f5f9",
    AllowImmediateDownload: true,
}

landingPage, err := client.Console.CreateLandingPage(ctx, params)
if err != nil {
    fmt.Printf("Error creating landing page: %v\n", err)
    return
}

fmt.Printf("Landing page created: %s\n", landingPage.ID)
fmt.Printf("Name: %s, Kind: %s\n", landingPage.Name, landingPage.Kind)
```

#### Update a landing page

```go
ctx := context.Background()
params := accessgrid.UpdateLandingPageParams{
    LandingPageID:  "0xlandingpage1d",
    Name:           "Updated Miami Office Access Pass",
    AdditionalText: "Welcome! Tap below to get your access pass.",
    BgColor:        "#e2e8f0",
}

landingPage, err := client.Console.UpdateLandingPage(ctx, params)
if err != nil {
    fmt.Printf("Error updating landing page: %v\n", err)
    return
}

fmt.Printf("Landing page updated: %s\n", landingPage.ID)
fmt.Printf("Name: %s\n", landingPage.Name)
```

### Credential Profiles

#### List credential profiles

```go
ctx := context.Background()
profiles, err := client.Console.CredentialProfiles.List(ctx)
if err != nil {
    fmt.Printf("Error listing credential profiles: %v\n", err)
    return
}

for _, profile := range profiles {
    fmt.Printf("ID: %s, Name: %s, AID: %s\n", profile.ID, profile.Name, profile.AID)
}
```

#### Create a credential profile

```go
ctx := context.Background()
params := accessgrid.CreateCredentialProfileParams{
    Name:    "Main Office Profile",
    AppName: "KEY-ID-main",
    Keys: []accessgrid.KeyParam{
        {Value: "your_32_char_hex_master_key_here"},
        {Value: "your_32_char_hex__read_key__here"},
    },
}

profile, err := client.Console.CredentialProfiles.Create(ctx, params)
if err != nil {
    fmt.Printf("Error creating credential profile: %v\n", err)
    return
}

fmt.Printf("Profile created: %s\n", profile.ID)
fmt.Printf("AID: %s\n", profile.AID)
```

## Configuration

The SDK can be configured with custom options:

```go
client, err := accessgrid.NewClient(accountID, secretKey)
if err != nil {
    fmt.Printf("Error creating client: %v\n", err)
    return
}
```

## Error Handling

The SDK throws errors for various scenarios including:
- Missing required credentials
- API request failures
- Invalid parameters
- Server errors

Example error handling:

```go
params := accessgrid.ProvisionParams{
    // ... parameters
}

card, err := client.AccessCards.Provision(params)
if err != nil {
    fmt.Printf("Error provisioning card: %v\n", err)
    return
}
```

## Requirements

- Go 1.18 or higher

## Security

The SDK automatically handles:
- Request signing using HMAC-SHA256
- Secure payload encoding
- Authentication headers
- HTTPS communication

Never expose your `secretKey` in source code. Always use environment variables or a secure configuration management system.

## Feature Matrix

| Endpoint | Method | Supported |
|---|---|:---:|
| POST /v1/key-cards | `AccessCards.Provision()` | Y |
| GET /v1/key-cards/{id} | `AccessCards.Get()` | Y |
| PATCH /v1/key-cards/{id} | `AccessCards.Update()` | Y |
| GET /v1/key-cards | `AccessCards.List()` | Y |
| POST /v1/key-cards/{id}/suspend | `AccessCards.Suspend()` | Y |
| POST /v1/key-cards/{id}/resume | `AccessCards.Resume()` | Y |
| POST /v1/key-cards/{id}/unlink | `AccessCards.Unlink()` | Y |
| POST /v1/key-cards/{id}/delete | `AccessCards.Delete()` | Y |
| POST /v1/console/card-templates | `Console.CreateTemplate()` | Y |
| PUT /v1/console/card-templates/{id} | `Console.UpdateTemplate()` | Y |
| GET /v1/console/card-templates/{id} | `Console.ReadTemplate()` | Y |
| GET /v1/console/card-templates/{id}/logs | `Console.EventLog()` | Y |
| GET /v1/console/pass-template-pairs | `Console.ListPassTemplatePairs()` | Y |
| POST /v1/console/card-templates/{id}/ios_preflight | `Console.IosPreflight()` | Y |
| GET /v1/console/ledger-items | `Console.ListLedgerItems()` | Y |
| GET /v1/console/webhooks | `Console.Webhooks.List()` | Y |
| POST /v1/console/webhooks | `Console.Webhooks.Create()` | Y |
| DELETE /v1/console/webhooks/{id} | `Console.Webhooks.Delete()` | Y |
| GET /v1/console/landing-pages | `Console.ListLandingPages()` | Y |
| POST /v1/console/landing-pages | `Console.CreateLandingPage()` | Y |
| PUT /v1/console/landing-pages/{id} | `Console.UpdateLandingPage()` | Y |
| GET /v1/console/credential-profiles | `Console.CredentialProfiles.List()` | Y |
| POST /v1/console/credential-profiles | `Console.CredentialProfiles.Create()` | Y |
| POST /v1/console/hid/orgs | `Console.HID.Orgs.Create()` | Y |
| POST /v1/console/hid/orgs/activate | `Console.HID.Orgs.Activate()` | Y |
| GET /v1/console/hid/orgs | `Console.HID.Orgs.List()` | Y |

## License

MIT License - See LICENSE file for details.