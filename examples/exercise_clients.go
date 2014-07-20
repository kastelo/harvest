package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseClients(apiClient *harvest.APIClient) {
	clients, err := apiClient.Client.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%+v\n", clients)
		client, err := apiClient.Client.Find(clients[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("\n%v\n", client)
		}
	}
}
