package sentry

import (
	"encoding/json"
	"fmt"
	"time"
)

// Status is used for the status of an entity.
type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Avatar is used to display an avatar for an entity.
type Avatar struct {
	// Type is the type of avatar set for the entity.
	Type string `json:"avatarType"`
	// ID is a uuid or empty string.
	ID string `json:"avatarUuid"`
}

// Organization represents an org in sentry.
type Organization struct {
	// Name is the name of the org.
	Name string `json:"name"`
	// Slug is the org's unique name.
	Slug string `json:"slug"`
	// ID is the id of the org.
	ID string `json:"id"`

	// Avatar is used to display a picture for an Organziation
	Avatar Avatar `json:"avatar"`

	// DateCreated is the date the org was created.
	DateCreated time.Time `json:"dateCreated"`

	// IsEarlyAdopter flag if the org is under an early adopter account.
	IsEarlyAdopter bool `json:"isEarlyAdopter"`
}

// String is formatted version of the Organization struct.
func (o Organization) String() string {
	result := fmt.Sprintf("Name: %v, ID: %v, Slug: %v\n", o.Name, o.ID, o.Slug)
	return result
}

// OrganizationList returns the list of Organizations available to the authenticated session.
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
