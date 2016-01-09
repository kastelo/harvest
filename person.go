package harvest

import (
	"fmt"
	"time"
)

type PersonService struct {
	Service
}

type Person struct {
	Id                           int
	DefaultExpenseCategoryId     int
	DefaultExpenseProjectId      int
	DefaultTaskId                int
	DefaultTimeProjectId         int
	DefaultHourlyRate            float32
	FirstName                    string
	LastName                     string
	Email                        string   `json:"email"`
	IdentityUrl                  string
	OpensocialIdentifier         string
	Telephone                    string
	Timezone                     string
	CostRate                     string
	WeeklyDigestSentOn           string
	Department                   string
	IsContractor                 bool
	IsAdmin                      bool
	IsActive                     bool
	HasAccessToAllFutureProjects bool
	WantsNewsletter              bool
	WantsWeeklyDigest            bool
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
}

type PersonResponse struct {
	Person `json:"user"`
}

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

func (c *PersonService) Find(personID int) (person PersonResponse, err error) {
	resourceURL := fmt.Sprintf("/people/%v.json", personID)
	var resp PersonResponse
	err = c.find(resourceURL, &resp)
	return resp, err
}
