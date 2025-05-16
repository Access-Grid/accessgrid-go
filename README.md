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

import (\n    "context"
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
        CardTemplateID:          "0xd3adb00b5",
        EmployeeID:             "123456789",
        CardNumber:             "14563",
        SiteCode:               "42",
        AllowOnMultipleDevices: true,
        FullName:              "Employee name",
        Email:                 "employee@yourwebsite.com",
        PhoneNumber:           "+19547212241",
        Classification:        "full_time",
        StartDate:            time.Now().UTC(),
        ExpirationDate:       time.Now().UTC().AddDate(1, 0, 0),
        EmployeePhoto:        "[image_in_base64_encoded_format]",
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

   design := accessgrid.TemplateDesign{
       BackgroundColor:     "#FFFFFF",
       LabelColor:         "#000000",
       LabelSecondaryColor: "#333333",
       BackgroundImage:     "[image_in_base64_encoded_format]",
       LogoImage:          "[image_in_base64_encoded_format]",
       IconImage:          "[image_in_base64_encoded_format]",
   }

   supportInfo := accessgrid.SupportInfo{
       SupportURL:           "https://help.yourcompany.com",
       SupportPhoneNumber:   "+1-555-123-4567",
       SupportEmail:         "support@yourcompany.com",
       PrivacyPolicyURL:     "https://yourcompany.com/privacy",
       TermsAndConditionsURL: "https://yourcompany.com/terms",
   }

   params := accessgrid.CreateTemplateParams{
       Name:                 "Employee NFC key",
       Platform:            "apple",
       UseCase:             "employee_badge",
       Protocol:            "desfire",
       AllowOnMultipleDevices: true,
       WatchCount:          2,
       IPhoneCount:         3,
       Design:              design,
       SupportInfo:         supportInfo,
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

   supportInfo := accessgrid.SupportInfo{
       SupportURL:           "https://help.yourcompany.com",
       SupportPhoneNumber:   "+1-555-123-4567",
       SupportEmail:         "support@yourcompany.com",
       PrivacyPolicyURL:     "https://yourcompany.com/privacy",
       TermsAndConditionsURL: "https://yourcompany.com/terms",
   }

   params := accessgrid.UpdateTemplateParams{
       CardTemplateID:       "0xd3adb00b5",
       Name:                "Updated Employee NFC key",
       AllowOnMultipleDevices: true,
       WatchCount:          2,
       IPhoneCount:         3,
       SupportInfo:         supportInfo,
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

## License

MIT License - See LICENSE file for details.