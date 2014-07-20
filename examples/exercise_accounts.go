package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseAccounts(apiClient *harvest.APIClient) {
	account, err := apiClient.Account.Find()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%v\n", account)
	}
}
