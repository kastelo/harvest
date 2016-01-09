package harvest

import (
	"fmt"
	"time"
)

type EntryService struct {
	Service
}

type Entry struct {
	Id        int         `json:"id"`
	UserId    int         `json:"user_id"`
	TaskId    int         `json:"task_id"`
	InoviceId int         `json:"invoice_id"`
	ProjectId int         `json:"project_id"`
	Hours     float32     `json:"hours"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"created_at"`
	SpentAt   HarvestDate `json:"spent_at"`
	// `timer_started_at` is in Harvest's JSON response, but in entries examined
	// always set to null, and is creating JSON marshalling problems, so we're
	// ignoring it for now.
	//TimerStartedAt time.Time   `json:"timer_started_at"`
	IsBilled bool   `json:"is_billed"`
	IsClosed bool   `json:"created_at"`
	Notes    string `json:"notes"`
}

type EntryResponse struct {
	Entry `json:"day_entry"`
}

// Returns all entries associated with a project
func (c *EntryService) ListProject(projectID int, from time.Time, to time.Time) (entries []EntryResponse, err error) {
	return c.entryList("projects", projectID, from, to)
}

// Returns all entries associated with a person
func (c *EntryService) ListPerson(personID int, from time.Time, to time.Time) (entries []EntryResponse, err error) {
	return c.entryList("people", personID, from, to)
}

func (c *EntryService) entryList(resource string, projectID int, from time.Time, to time.Time) (entries []EntryResponse, err error) {
	// TODO check that `to` does not precede `from`
	from_format := from.Format("20060102")
	to_format := to.Format("20060102")
	var resp []EntryResponse
	resourceURL := fmt.Sprintf("/%s/%v/entries.json?from=%v&to=%v", resource, projectID, from_format, to_format)

	err = c.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		entries = append(entries, element)
	}
	return entries, err
}
