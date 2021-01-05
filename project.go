package sentry

import (
	"encoding/json"
	"time"
)

// Project is a sentry project.
type Project struct {
	Avatar       Avatar       `json:"avatar"`
	Color        string       `json:"color"`
	DateCreated  time.Time    `json:"dateCreated"`
	Features     []string     `json:"features"`
	FirstEvent   string       `json:"firstEvent"`
	HasAccess    bool         `json:"hasAccess"`
	ID           string       `json:"id"`
	IsBookmarked bool         `json:"isBookmarked"`
	IsInternal   bool         `json:"isInternal"`
	IsMember     bool         `json:"isMember"`
	IsPublic     bool         `json:"isPublic"`
	Name         string       `json:"name"`
	Organization Organization `json:"organization"`
	Platform     string       `json:"platform"`
	Slug         string       `json:"slug"`
	Status       string       `json:"status"`
}

// ProjectList returns the list of Projects.
func (c *Client) ProjectList(cursor string) ([]Project, error) {
	response, err := c.APIGet("projects/")
	if err != nil {
		return nil, err
	}
	projects := make([]Project, 1)
	err = json.Unmarshal(response, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
