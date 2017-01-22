package harvest

import (
	"fmt"
	"time"
)

type ExpenseCategoryService struct {
	Service
}

type ExpenseCategory struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	UnitName    string    `json:"unit_name"`
	UnitPrice   float32   `json:"unit_price"`
	Deactivated bool      `json:"deactivated"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type CategoryResponse struct {
	ExpenseCategory `json:"expense_category"`
}

// List requests list of expense categories and returns response
func (c *ExpenseCategoryService) List() (expenseCategories []CategoryResponse, err error) {
	resourceURL := "/expense_categories.json"
	var resp []CategoryResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		expenseCategories = append(expenseCategories, element)
	}
	return expenseCategories, err
}

// Find requests expense category information for specified expense category
// and returns response
func (c *ExpenseCategoryService) Find(catID int) (expense CategoryResponse, err error) {
	resourceURL := fmt.Sprintf("/expense_categories/%v.json", catID)
	var resp CategoryResponse
	err = c.find(resourceURL, &resp)
	return resp, err
}
