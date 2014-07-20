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
	ExpenseCategory `json:"expense_category"`
}

func (c *ExpenseCategoryService) List() (err error, expenseCategories []ExpenseCategoryResponse) {
	resourceURL := "/expense_categories.json"
	var resp []ExpenseCategoryResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return err, resp
	}
	for _, element := range resp {
		expenseCategories = append(expenseCategories, element)
	}
	return err, expenseCategories
}

func (c *ExpenseCategoryService) Find(catId int) (err error, expense ExpenseCategoryResponse) {
	resourceURL := fmt.Sprintf("/expense_categories/%v.json", catId)
	var resp ExpenseCategoryResponse
	err = c.find(resourceURL, &resp)
	return err, resp
}
