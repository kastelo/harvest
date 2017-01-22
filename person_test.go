package harvest

import (
	"fmt"
	"github.com/kr/pretty"
	"os"
	"testing"
)

func TestPeople(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "person_test:") //Used for debugging

	people, err := apiClient.People.List()
	if err != nil {
		t.Fatalf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n\n\n", pretty.Formatter(people))
		person, err := apiClient.People.Find(people[0].Id)
		if err != nil {
			t.Fatalf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n\n\n", pretty.Formatter(person))
		}
	}
}
