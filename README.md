# px-sdk-go
API SDK for usage with the FRIS Solutions PerimeterX API.

## Example usage
This example uses the `Session.GeneratePerimeterXCookie` function, which automatically generates cookies for you.
<br>
It also uses the `Session.SolveHoldCaptcha` function, which automatically solves holdcaptcha for you.
```go
package main

import (
	"context"
	"github.com/FRIS-Solutions-Vault/px-sdk-go"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func main() {
	// Create an API session; this can be re-used across multiple tasks
	session := px.NewSession(os.Getenv("FRIS_API_KEY"))
	
    // The user agent to use
    const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"

	req := &GenerateRequest{
		UserAgent: userAgent,
		PageURL:   "https://auth.ticketmaster.com/",
		Proxy:     "",
	}

    // Generate px cookie
	if sensor, err := session.GeneratePerimeterXCookie(context.Background(), req); err != nil {
		panic(err)
	} else {
		log.Println("PX Response: ", sensor)
	}

	// HoldCaptcha
	req = &GenerateRequest{
		UserAgent: userAgent,
		PageURL:   "https://auth.ticketmaster.com/",
		Proxy:     "",
		Data:      sensor.Data,
	}

	if sensor, err := session.GeneratePerimeterXCookie(context.Background(), req); err != nil {
		panic(err)
	} else {
		log.Println("PX HoldCaptcha Response: ", sensor)
	}
}
```