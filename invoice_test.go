package harvest

import (
	"fmt"
	"github.com/kr/pretty"
	"os"
	"testing"
)

func TestInvoices(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "invoice_test:") //Used for debugging

	invoices, err := apiClient.Invoice.List()
	if err != nil {
		t.Fatalf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n", pretty.Formatter(invoices))
		invoice, err := apiClient.Invoice.Find(invoices[0].ID)
		if err != nil {
			t.Fatalf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n", pretty.Formatter(invoice))
		}
	}
}
