package harvest

import (
	"github.com/kr/pretty"
	"testing"
	"os"
	"fmt"
)

func TestAccounts(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "account_test:")		//Used for debugging

	account, err := apiClient.Account.Find()
	if err != nil {
		t.Fatalf("\n%v\n", err)
		t.Fail()
	} else {
		// t.Logf("%# v", pretty.Formatter(account))
		fmt.Printf("\n%# v\n", pretty.Formatter(account))

	}
}
