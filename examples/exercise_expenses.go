package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseExpenseCategories(apiClient *harvest.APIClient) {
	err, expenseCategories := apiClient.ExpenseCategory.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", expenseCategories)
		err, category := apiClient.ExpenseCategory.Find(expenseCategories[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", category)
		}
	}
}
