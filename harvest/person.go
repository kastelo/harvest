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
	Email                        string
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

func (c *PersonService) List() (err error, people []PersonResponse) {
	resourceURL := "/people.json"
	var resp []PersonResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return err, resp
	}
	for _, element := range resp {
		people = append(people, element)
	}
	return err, people
}

func (c *PersonService) Find(personID int) (err error, person PersonResponse) {
	resourceURL := fmt.Sprintf("/people/%v.json", personID)
	var resp PersonResponse
	err = c.find(resourceURL, &resp)
	return err, resp
}
