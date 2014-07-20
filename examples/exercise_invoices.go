package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseInvoices(apiClient *harvest.APIClient) {
	invoices, err := apiClient.Invoice.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%+v\n", invoices)
		invoice, err := apiClient.Invoice.Find(invoices[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("\n%v\n", invoice)
		}
	}
}
