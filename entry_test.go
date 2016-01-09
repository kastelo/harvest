package examples

import (
	"fmt"
	"time"

	"../harvest"
)

func TestProjectEntries(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

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
					fmt.Printf("\n%v\n", entries[0])
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
					fmt.Printf("\n%v\n", entries[0])
				}
			}
		}
	}
}