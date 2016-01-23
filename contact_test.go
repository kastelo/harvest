package harvest

import (
	"github.com/kr/pretty"
	"fmt"
	"testing"
	"os"
)

func TestContacts(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "contact_test:")		//Used for debugging

	contacts, err := apiClient.Contact.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n", pretty.Formatter(contacts))
		contact, err := apiClient.Contact.Find(contacts[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n", pretty.Formatter(contact))
		}
	}
}
