package px

import (
	"net/http"
)

// Session is an API session that allows interaction with the FRIS Solutions PerimeterX API.
// Sessions should be created with one of the utility functions, like NewSession.
//
// Callers should re-use the same Session as much as possible, even across different tasks.
// This will provide the best performance.
type Session struct {
	// The API key to use when authorizing with the FRIS Solutions API.
	apiKey string

	// The http.Client to use when making API requests.
	client *http.Client
}

// NewSessionWithClient creates a new Session with the given API key and HTTP client.
// The given client is responsible for making requests to the FRIS Solutions API.
//
// NewSessionWithClient panics if client == nil.
func NewSessionWithClient(apiKey string, client *http.Client) Session {
	if client == nil {
		panic("px-sdk-go: nil client passed to NewSessionWithClient")
	}

	return Session{
		apiKey: apiKey,
		client: client,
	}
}

// NewSession creates a new Session with the given API key.
// It uses the default client to make requests to the FRIS Solutions API.
func NewSession(apiKey string) Session {
	return NewSessionWithClient(apiKey, http.DefaultClient)
}
