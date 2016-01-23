package harvest

import (
	"github.com/kr/pretty"
	"fmt"
	"time"
	"testing"
	"os"
)

func TestProjectEntries(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "entry_test:")		//Used for debugging

	from_time, _ := time.Parse("2006-01-02", "2014-01-01")
	to_time, _ := time.Parse("2006-01-02", "2014-02-01")

	projects, err := apiClient.Project.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		for _, project := range projects {
			entries, err := apiClient.Entry.ListProject(project.Id, from_time, to_time)
			if err != nil {
				fmt.Printf("Error: %v", err)
			} else {
				if len(entries) > 0 {
					fmt.Printf("\n%# v\n", pretty.Formatter(entries[0]))
				}
			}
		}
	}
}

func ExercisePersonEntries(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	from_time, _ := time.Parse("2006-01-02", "2014-01-01")
	to_time, _ := time.Parse("2006-01-02", "2014-02-01")

  people, err := apiClient.People.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		for _, person := range people {
			entries, err := apiClient.Entry.ListPerson(person.Id, from_time, to_time)
			if err != nil {
				fmt.Printf("Error: %v", err)
			} else {
				if len(entries) > 0 {
					fmt.Printf("\n%# v\n", pretty.Formatter(entries[0])
				}
			}
		}
	}
}
