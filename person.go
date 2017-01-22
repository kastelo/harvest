package harvest

import (
	"fmt"
	"time"
)

type PersonService struct {
	Service
}

type Person struct {
	ID                           int `json:"id"`
	DefaultExpenseCategoryID     int
	DefaultExpenseProjectID      int
	DefaultTaskID                int
	DefaultTimeProjectID         int
	DefaultHourlyRate            float32 `json:"default_hourly_rate"`
	FirstName                    string  `json:"first_name"`
	LastName                     string  `json:"last_name"`
	Email                        string  `json:"email"`
	IdentityURL                  string
	OpensocialIdentifier         string
	Telephone                    string  `json:"telephone"`
	Timezone                     string  `json:"timezone"`
	CostRate                     float32 `json:"cost_rate"`
	WeeklyDigestSentOn           string
	Department                   string `json:"department"`
	IsContractor                 bool   `json:"is_contractor"`
	IsAdmin                      bool   `json:"is_admin"`
	IsActive                     bool   `json:"is_active"`
	HasAccessToAllFutureProjects bool   `json:"has_access_to_all_future_projects"`
	WantsNewsletter              bool   `json:"wants_newsletter"`
	WantsWeeklyDigest            bool
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
}

type PersonResponse struct {
	Person `json:"user"`
}

// List requests list of people and returns response
func (c *PersonService) List() ([]PersonResponse, error) {
	resourceURL := "/people.json"
	var people []PersonResponse
	if err := c.get(resourceURL, &people); err != nil {
		return nil, err
	}
	return people, nil
}

// Find requests people information for specified person and returns response
func (c *PersonService) Find(personID int) (PersonResponse, error) {
	resourceURL := fmt.Sprintf("/people/%v.json", personID)
	var resp PersonResponse
	err := c.get(resourceURL, &resp)
	return resp, err
}
