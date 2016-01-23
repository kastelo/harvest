package harvest

import (
	"github.com/kr/pretty"
	"fmt"
	"testing"
	"os"
)

func TestProjects(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "project_test:")		//Used for debugging

	projects, err := apiClient.Project.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n", pretty.Formatter(projects))
		project, err := apiClient.Project.Find(projects[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n", pretty.Formatter(project))
		}
	}
}
