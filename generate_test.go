package px

import (
	"context"
	"log"
	"os"
	"testing"
)

// TestGenerate tests the Session.GeneratePerimeterXCookie method
//
// Users wishing to run this test need to use their own FRIS Solutions API key by setting the FRIS_API_KEY
// environment variable when running the test.
func TestGenerate(t *testing.T) {
	session := NewSession(os.Getenv("FRIS_API_KEY"))

	const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"

	req := &GenerateRequest{
		UserAgent: userAgent,
		PageURL:   "https://auth.ticketmaster.com/",
		Proxy:     "",
	}

	var sensor *GenerateResponse

	if sensor, err := session.GeneratePerimeterXCookie(context.Background(), req); err != nil {
		panic(err)
	} else {
		log.Println("PX Response: ", sensor)
	}

	// HoldCaptcha
	hcReq := &GenerateHoldCapRequest{
		UserAgent: userAgent,
		PageURL:   "https://auth.ticketmaster.com/",
		Proxy:     "",
		Data:      sensor.Headers["data"].(string),
		PxHd:      "",
	}

	if sensor, err := session.SolveHoldCaptcha(context.Background(), hcReq); err != nil {
		panic(err)
	} else {
		log.Println("PX HoldCaptcha Response: ", sensor)
	}
}
