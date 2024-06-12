
# InviDNS

InviDNS is a Go package that provides a DNS provider implementation using the [libdns](https://github.com/libdns/libdns) interface. It offers a convenient way to manage DNS records for domains using the Duck DNS service.

## Features

- **Flexible Configuration**: Easily configure the provider with your Duck DNS API token and override domain.
- **Record Management**: Retrieve, update, and clear DNS records for your domain.
- **Support for Multiple Record Types**: Manage A, AAAA, and TXT records with ease.
- **Concurrency-Safe Operations**: The package ensures safe access to shared resources using mutexes.
- **Context Handling**: Proper propagation of context for cancellation and timeouts during HTTP requests.

## Installation

To use InviDNS in your Go project, you can install it using `go get`:

```bash
go get github.com/aayush-makhija/invidns
```

## Usage

Here's a simple example of how to use InviDNS to retrieve DNS records for a domain:

```go
package main

import (
	"context"
	"fmt"
	"github.com/aayush-makhija/invidns"
)

func main() {
	// Create a new instance of the InviDNS provider
	provider := &invidns.Provider{
		URL:            "https://yourapi.example.com/update",
		APIToken:       "your-api-token",
		OverrideDomain: "example.com",
	}

	// Retrieve DNS records for a domain
	records, err := provider.GetRecords(context.Background(), "example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print retrieved DNS records
	fmt.Println("Retrieved DNS records:", records)
}
```
