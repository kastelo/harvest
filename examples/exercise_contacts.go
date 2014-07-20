package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseContacts(apiClient *harvest.APIClient) {
	contacts, err := apiClient.Contact.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", contacts)
		contact, err := apiClient.Contact.Find(contacts[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", contact)
		}
	}
}
