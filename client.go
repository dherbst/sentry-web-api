package sentry

import "time"

// Client is used to make calls to the sentry web api.
type Client struct {
	// Domain is the base domain for the api, including https://
	Domain string

	// Timeout in seconds before an api call times out.
	Timeout time.Duration
}

// NewClient creates a client with the given domain to make api calls.
func NewClient(domain string, seconds int) *Client {
	if domain == "" {
		domain = "https://sentry.io/api/0/"
	}
	if seconds == 0 {
		seconds = 3
	}

	c := &Client{Domain: domain, Timeout: time.Duration(seconds * time.Second)}
	return c
}
