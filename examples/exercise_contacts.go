package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseContacts(apiClient *harvest.APIClient) {
	err, contacts := apiClient.Contact.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", contacts)
		err, contact := apiClient.Contact.Find(contacts[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", contact)
		}
	}
}
