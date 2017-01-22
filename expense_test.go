package harvest

import (
	"fmt"
	"github.com/kr/pretty"
	"os"
	"testing"
)

func TestExpenseCategories(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "expense_test:") //Used for debugging

	expenseCategories, err := apiClient.ExpenseCategory.List()
	if err != nil {
		t.Fatalf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n", pretty.Formatter(expenseCategories))
		category, err := apiClient.ExpenseCategory.Find(expenseCategories[0].ID)
		if err != nil {
			t.Fatalf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n", pretty.Formatter(category))
		}
	}
}
