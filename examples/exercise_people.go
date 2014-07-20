package examples

import (
	"fmt"

	"../harvest"
)

func ExercisePeople(apiClient *harvest.APIClient) {
	people, err := apiClient.People.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n\n\n", people)
		person, err := apiClient.People.Find(people[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n\n\n", person)
		}
	}
}
