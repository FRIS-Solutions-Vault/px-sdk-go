package px

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// GenerateRequest is the API generation request schema.
type GenerateRequest struct {
	// ApiKey is the API key to use when generating px cookies.
	//
	// This is required for all generation requests.
	ApiKey string `json:"apiKey"`

	// UserAgent is the user agent to use when generating px cookies.
	//
	// Current restrictions apply to this preference. The user agent
	// must be a Google Chrome v114 or v119 user agent.
	// Callers can use any platform they like, however it is highly
	// recommended to use Windows.
	UserAgent string `json:"ua"`

	// PageURL is the URL of the page to generate sensor data for.
	PageURL string `json:"pageUrl"`

	// Proxy is the proxy to use when generating px cookies.
	Proxy string `json:"proxy"`

	// Data is the px data to use when solving holdcaptcha challenges.
	//
	// This is optional for the generation request, but required for holdcaptcha
	Data string `json:"data,omitempty"`

	// _pxhd is the pxhd value
	//
	// This is optional for the generation request, but required for holdcaptcha
	PxHd string `json:"_pxhd,omitempty"`
}

// GenerateResponse is the API generation response schema.
type GenerateResponse struct {
	// Cookie is the px cookie
	Cookie string `json:"cookie"`

	// Cts is the px cts value
	Cts string `json:"cts"`

	// Vid is the px vid value
	Vid string `json:"vid"`

	// Headers is the headers value
	Headers map[string]interface{} `json:"headers"`

	// Success is the success value
	Success bool `json:"success"`

	// Flagged is the flagged value
	Flagged bool `json:"flagged"`

	// Data is the data value
	//
	// Only returned for holdcaptcha
	Data string `json:"data,omitempty"`
}

// GeneratePerimeterXCookie generates perimeterx cookie
func (session Session) GeneratePerimeterXCookie(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error) {
	req.ApiKey = session.apiKey
	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", "https://api.frisapi.dev/pxweb/init", bytes.NewBuffer(encoded))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := session.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, ApiOperationError{
			StatusCode: response.StatusCode,
			Message:    GetMessageFromErrorResponse(body),
		}
	}

	var resp GenerateResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

var (
	// ErrInvalidPageURL is an error caused by Session.Generate if the provided page URL is not
	// a valid or absolute URL. An absolute URL must contain a scheme and host.
	ErrInvalidPageURL = errors.New("px-sdk-go: invalid page URL")
)

// SolveHoldCaptcha solves holdcaptcha
func (session Session) SolveHoldCaptcha(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error) {
	req.ApiKey = session.apiKey
	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", "https://api.frisapi.dev/pxweb/holdcap", bytes.NewBuffer(encoded))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := session.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, ApiOperationError{
			StatusCode: response.StatusCode,
			Message:    GetMessageFromErrorResponse(body),
		}
	}

	var resp GenerateResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
