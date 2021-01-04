package sentry

import (
	"os"
	"time"
)

var (
	// GitHash is the sha hash of the current commit used to build this code.
	GitHash string
	// Version is the semantic version according to the release.
	Version string
)

// Client is used to make calls to the sentry web api.
type Client struct {
	// AuthToken is the bearer token for use in the API.
	AuthToken string

	// Domain is the base domain for the api, including https://
	Domain string

	// Timeout in seconds before an api call times out.
	Timeout time.Duration
}

// NewClient creates a client with the given domain to make api calls.
func NewClient(domain string, seconds int, authToken string) *Client {
	if domain == "" {
		domain = "https://sentry.io/api/0/"
	}
	if seconds == 0 {
		seconds = 3
	}
	if authToken == "" {
		authToken = os.Getenv("SENTRY_AUTH_TOKEN")
	}

	c := &Client{
		Domain:    domain,
		Timeout:   time.Duration(seconds) * time.Second,
		AuthToken: authToken,
	}
	return c
}
