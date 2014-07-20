package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseProjects(apiClient *harvest.APIClient) {
	projects, err := apiClient.Project.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", projects)
		project, err := apiClient.Project.Find(projects[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", project)
		}
	}
}
