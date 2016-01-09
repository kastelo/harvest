package harvest

import (
	"fmt"

	"../harvest"
)

func TestExpenseCategories(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	expenseCategories, err := apiClient.ExpenseCategory.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", expenseCategories)
		category, err := apiClient.ExpenseCategory.Find(expenseCategories[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", category)
		}
	}
}
