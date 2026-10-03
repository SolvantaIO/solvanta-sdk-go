# Solvanta Go SDK

Official Go client for the [Solvanta](https://solvanta.dev) solver API.

## Install

```bash
go get github.com/SolvantaIO/solvanta-sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    solvanta "github.com/SolvantaIO/solvanta-sdk-go"
)

func main() {
    client := solvanta.New("sk_...")

    resp, err := client.Solve(context.Background(), solvanta.ReCaptchaV3Task, map[string]interface{}{
        "url":     "https://example.com",
        "sitekey": "6Le...",
        "action":  "login",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Status:", resp.Status)
    fmt.Println("Token:", string(resp.Result))
    fmt.Println("Credits:", resp.CreditsUsed)
}
```

## Usage

### Create a Client

```go
// Default settings
client := solvanta.New("sk_...")

// Custom options
client := solvanta.New("sk_...",
    solvanta.WithBaseURL("https://custom.endpoint.com"),
    solvanta.WithTimeout(60 * time.Second),
    solvanta.WithHTTPClient(customHTTPClient),
)
```

### Solve

```go
// ReCaptcha V2
resp, err := client.Solve(ctx, solvanta.ReCaptchaV2Task, map[string]interface{}{
    "url":     "https://example.com",
    "sitekey": "6Le...",
})

// Cloudflare
resp, err := client.Solve(ctx, solvanta.CloudflareTask, map[string]interface{}{
    "url":   "https://example.com",
    "proxy": "http://user:pass@host:port",
})

// With idempotency key
resp, err := client.SolveWithKey(ctx, solvanta.ArkoseTask, map[string]interface{}{
    "url":     "https://example.com",
    "sitekey": "pk_...",
}, "order-12345-captcha")
```

### Retrieve a Solve

```go
resp, err := client.GetSolve(ctx, "solve-uuid")
fmt.Println(resp.Status)
```

### List Solve History

```go
solves, err := client.ListSolves(ctx, 50)
for _, s := range solves {
    fmt.Println(s.ID, s.Status)
}
```

### Discover Services

```go
catalog, err := client.Services(ctx)
for _, svc := range catalog.Services {
    fmt.Printf("%s — %d credits\n", svc.Task, svc.Credits)
}
```

### Dashboard

```go
dash, err := client.Dashboard(ctx)
fmt.Printf("Credits: %d\n", dash.Credits)
```

## Task Constants

```go
solvanta.ReCaptchaV2Task           // "ReCaptchaV2Task"
solvanta.ReCaptchaV2EnterpriseTask // "ReCaptchaV2EnterpriseTask"
solvanta.ReCaptchaV3Task           // "ReCaptchaV3Task"
solvanta.ReCaptchaV3EnterpriseTask // "ReCaptchaV3EnterpriseTask"
solvanta.CloudflareTask            // "CloudflareTask"
solvanta.ArkoseTask                // "ArkoseTask"
solvanta.ForterTask                // "ForterTask"
solvanta.NuDataTask                // "NuDataTask"
solvanta.PxmTask                   // "PxmTask"
solvanta.ThreatMetrixTask          // "ThreatMetrixTask"
solvanta.UeMsmTask                 // "UeMsmTask"
solvanta.Metadata1Task             // "Metadata1Task"
```

## Error Handling

```go
resp, err := client.Solve(ctx, solvanta.CloudflareTask, payload)
if err != nil {
    if apiErr, ok := err.(*solvanta.Error); ok {
        fmt.Println("Code:", apiErr.Code)           // "insufficient_credits"
        fmt.Println("Message:", apiErr.Message)
        fmt.Println("HTTP:", apiErr.StatusCode)      // 402
    }
}
```

## Requirements

- Go 1.21+
- Zero external dependencies (stdlib only)
