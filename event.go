package sentry

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Event is a sentry event related to a project.
type Event struct {
	EventID     string    `json:"eventID"`
	Tags        []Tag     `json:"tags"`
	DateCreated time.Time `json:"dateCreated"`
	User        string    `json:"user"`
	Message     string    `json:"message"`
	Title       string    `json:"title"`
	ID          string    `json:"id"`
	Platform    string    `json:"platform"`
	EventType   string    `json:"event.type"`
	GroupID     string    `json:"groupID"`
}

// EventList returns a list for the given organization and project.
func (c *Client) EventList(orgSlug string, projectSlug string, full bool, cursor string) ([]Event, error) {
	url := fmt.Sprintf("projects/%v/%v/events/", url.PathEscape(orgSlug), url.PathEscape(projectSlug))

	response, err := c.APIGet(url)
	if err != nil {
		return nil, err
	}
	events := make([]Event, 1)
	err = json.Unmarshal(response, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
