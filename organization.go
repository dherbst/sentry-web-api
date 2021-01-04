package sentry

import (
	"encoding/json"
	"fmt"
)

// Organization represents an org in sentry.
type Organization struct {
	// Name is the name of the org.
	Name string `json:"name"`
	// Slug is the org's unique name.
	Slug string `json:"slug"`
	// ID is the id of the org.
	ID string
}

// String is formatted version of the Organization struct.
func (o Organization) String() string {
	result := fmt.Sprintf("Name: %v, ID: %v, Slug: %v\n", o.Name, o.ID, o.Slug)
	return result
}

// List returns the list of Organizations available to the authenticated session.
func (c *Client) OrganizationList(owner bool, cursor string) ([]Organization, error) {
	response, err := c.APIGet("organizations/")
	if err != nil {
		return nil, err
	}

	orgs := make([]Organization, 1)
	err = json.Unmarshal(response, &orgs)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}
