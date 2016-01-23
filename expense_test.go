package harvest

import (
	"github.com/kr/pretty"
	"fmt"
	"testing"
	"os"
)

func TestExpenseCategories(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "expense_test:")		//Used for debugging

	expenseCategories, err := apiClient.ExpenseCategory.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Logf("\n%# v\n", pretty.Formatter(expenseCategories))
		category, err := apiClient.ExpenseCategory.Find(expenseCategories[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Logf("\n%# v\n", pretty.Formatter(category))
		}
	}
}
