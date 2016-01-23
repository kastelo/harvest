package harvest

import (
	"github.com/kr/pretty"
	"fmt"
	"testing"
	"os"
)

func TestInvoices(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "invoice_test:")		//Used for debugging

	invoices, err := apiClient.Invoice.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n", pretty.Formatter(invoices))
		invoice, err := apiClient.Invoice.Find(invoices[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n", pretty.Formatter(invoice))
		}
	}
}
