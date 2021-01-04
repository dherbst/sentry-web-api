package sentry

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c := NewClient("", 0, "secret")
	if c.Domain != "https://sentry.io/api/0/" {
		t.Fatalf("Did not get expected domain, got %v\n", c.Domain)
	}
	if c.Timeout != time.Duration(3*time.Second) {
		t.Fatalf("Did not get expected timeout, got %v\n", c.Timeout)
	}
	if c.AuthToken != "secret" {
		t.Fatalf("Did not get expected auth token.  Check why.")
	}
}
