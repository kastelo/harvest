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
	UnitPrice   float64   `json:"unit_price"`
	Deactivated bool      `json:"deactivated"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type CategoryResponse struct {
	ExpenseCategory `json:"expense_category"`
}

// List requests list of expense categories and returns response
func (c *ExpenseCategoryService) List() ([]CategoryResponse, error) {
	resourceURL := "/expense_categories.json"
	var categories []CategoryResponse
	if err := c.get(resourceURL, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

// Find requests expense category information for specified expense category
// and returns response
func (c *ExpenseCategoryService) Find(catID int) (CategoryResponse, error) {
	resourceURL := fmt.Sprintf("/expense_categories/%v.json", catID)
	var resp CategoryResponse
	err := c.get(resourceURL, &resp)
	return resp, err
}
