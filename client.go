package sentry

import (
	"io/ioutil"
	"net/http"
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

// APIGet calls the API passing in the bearer token.
func (c *Client) APIGet(path string) ([]byte, error) {
	url := c.Domain + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if c.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}
	client := http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
