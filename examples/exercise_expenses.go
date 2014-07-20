package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseExpenseCategories(apiClient *harvest.APIClient) {
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
