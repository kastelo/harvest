package harvest

import (
	"fmt"
	"testing"
	"os"
)

func TestClients(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "client_test:")		//Used for debugging

	clients, err := apiClient.Client.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%+v\n", clients)
		client, err := apiClient.Client.Find(clients[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("\n%v", "Client test:")
			fmt.Printf("\n%v\n", client)
		}
	}
}
