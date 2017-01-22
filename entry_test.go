package harvest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kr/pretty"
)

func TestProjectEntries(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "entry_test:") //Used for debugging

	//First argument defines format
	//Date format: yyyy-mm-dd
	fromTime, _ := time.Parse("2006-01-02", "2016-01-01")
	toTime, _ := time.Parse("2006-01-02", "2016-01-23")

	projects, err := apiClient.Project.List()
	if err != nil {
		t.Fatalf("\n%v\n", err)
	} else {
		for _, project := range projects {
			entries, err := apiClient.Entry.ListProject(project.ID, fromTime, toTime)
			if err != nil {
				t.Fatalf("Error: %v", err)
			} else {
				if len(entries) > 0 {
					//fmt.Printf("\n%v\n", entries[0])
					t.Logf("\n%# v\n", pretty.Formatter(entries[0]))
				}
			}
		}
	}
}

func TestPersonEntries(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	from_time, _ := time.Parse("2006-01-02", "2016-01-01")
	to_time, _ := time.Parse("2006-01-02", "2016-01-23")

	people, err := apiClient.People.List()
	if err != nil {
		t.Fatalf("\n%v\n", err)
	} else {
		for _, person := range people {
			entries, err := apiClient.Entry.ListPerson(person.ID, from_time, to_time)
			if err != nil {
				t.Fatalf("Error: %v", err)
			} else {
				if len(entries) > 0 {
					//fmt.Printf("\n%v\n", entries[0])
					t.Logf("\n%# v\n", pretty.Formatter(entries[0]))
				}
			}
		}
	}
}
