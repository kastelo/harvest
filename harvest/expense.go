package harvest

import (
	"fmt"
	"time"
)

type ExpenseCategoryService struct {
	Service
}

type ExpenseCategory struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	UnitName    string    `json:"unit_name"`
	UnitPrice   float32   `json:"unit_price"`
	Deactivated bool      `json:"deactivated"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type ExpenseCategoryResponse struct {
	ExpenseCategory ExpenseCategory `json:"expense_category"`
}

func (c *ExpenseCategoryService) List() (err error, expenseCategories []ExpenseCategory) {
	resourceURL := "/expense_categories.json"
	var expenseCategoryResponse []ExpenseCategoryResponse
	err = c.list(resourceURL, &expenseCategoryResponse)
	if err != nil {
		return
	}

	for _, element := range expenseCategoryResponse {
		expenseCategories = append(expenseCategories, element.ExpenseCategory)
	}
	return
}

func (c *ExpenseCategoryService) Find(expenseCategoryID int) (err error, expense ExpenseCategory) {
	resourceURL := fmt.Sprintf("/expense_categories/%v.json", expenseCategoryID)
	var expenseCategoryResponse ExpenseCategoryResponse
	err = c.find(resourceURL, &expenseCategoryResponse)
	if err != nil {
		return
	}
	expense = expenseCategoryResponse.ExpenseCategory
	return
}
