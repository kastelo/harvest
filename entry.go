package harvest

import (
	"fmt"
	"time"
)

type EntryService struct {
	Service
}

type Entry struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TaskID    int       `json:"task_id"`
	InvoiceID int       `json:"invoice_id"`
	ProjectID int       `json:"project_id"`
	Hours     float32   `json:"hours"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	SpentAt   Date      `json:"spent_at"`
	// `timer_started_at` is in Harvest's JSON response, but in entries examined
	// always set to null, and is creating JSON marshalling problems, so we're
	// ignoring it for now.
	//TimerStartedAt time.Time   `json:"timer_started_at"`
	IsBilled bool   `json:"is_billed"`
	IsClosed bool   `json:"is_closed"`
	Notes    string `json:"notes"`
}

type EntryResponse struct {
	Entry `json:"day_entry"`
}

// Returns all entries associated with a project
func (c *EntryService) ListProject(projectID int, from time.Time, to time.Time) ([]EntryResponse, error) {
	return c.entryList("projects", projectID, from, to)
}

// Returns all entries associated with a person
func (c *EntryService) ListPerson(personID int, from time.Time, to time.Time) ([]EntryResponse, error) {
	return c.entryList("people", personID, from, to)
}

// entryList returns entries
func (c *EntryService) entryList(resource string, projectID int, from time.Time, to time.Time) ([]EntryResponse, error) {
	// TODO check that `to` does not precede `from`
	fromStr := from.Format("20060102")
	toStr := to.Format("20060102")
	var entries []EntryResponse
	resourceURL := fmt.Sprintf("/%s/%v/entries.json?from=%v&to=%v", resource, projectID, fromStr, toStr)

	if err := c.get(resourceURL, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}
