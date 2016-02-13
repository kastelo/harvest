package harvest

import (
	"fmt"
	"time"
)

type PersonService struct {
	Service
}

type Person struct {
	Id                           int		`json:"id"`
	DefaultExpenseCategoryId     int
	DefaultExpenseProjectId      int
	DefaultTaskId                int
	DefaultTimeProjectId         int
	DefaultHourlyRate            float32	`json:"default_hourly_rate"`
	FirstName                    string		`json:"first_name"`
	LastName                     string		`json:"last_name"`
	Email                        string		`json:"email"`
	IdentityUrl                  string
	OpensocialIdentifier         string
	Telephone                    string		`json:"telephone"`
	Timezone                     string		`json:"timezone"`
	CostRate                     string		`json:"cost_rate"`
	WeeklyDigestSentOn           string
	Department                   string		`json:"department"`
	IsContractor                 bool		`json:"is_contractor"`
	IsAdmin                      bool		`json:"is_admin"`
	IsActive                     bool		`json:"is_active"`
	HasAccessToAllFutureProjects bool		`json:"has_access_to_all_future_projects"`
	WantsNewsletter              bool		`json:"wants_newsletter"`
	WantsWeeklyDigest            bool
	CreatedAt                    time.Time	`json:"created_at"`
	UpdatedAt                    time.Time	`json:"updated_at"`
}

type PersonResponse struct {
	Person `json:"user"`
}

// List requests list of people and returns response
func (c *PersonService) List() (people []PersonResponse, err error) {
	resourceURL := "/people.json"
	var resp []PersonResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		people = append(people, element)
	}
	return people, err
}

// Find requests people information for specified person and returns response
func (c *PersonService) Find(personID int) (person PersonResponse, err error) {
	resourceURL := fmt.Sprintf("/people/%v.json", personID)
	var resp PersonResponse
	err = c.find(resourceURL, &resp)
	return resp, err
}
