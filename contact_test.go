package harvest

import (
	"fmt"
	"github.com/kr/pretty"
	"os"
	"testing"
)

func TestContacts(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "contact_test:") //Used for debugging

	contacts, err := apiClient.Contact.List()
	if err != nil {
		t.Fatalf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n", pretty.Formatter(contacts))
		contact, err := apiClient.Contact.Find(contacts[0].ID)
		if err != nil {
			t.Fatalf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n", pretty.Formatter(contact))
		}
	}
}
