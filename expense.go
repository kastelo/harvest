package harvest

import (
	"fmt"
	"time"
)

type ExpenseService struct {
	Service
}

type Expense struct {
	ID                int     `json:"id,omitempty"`
	Notes             string  `json:"notes"`
	TotalCost         float64 `json:"total_cost,omitempty"`
	Units             float64 `json:"units,omitempty"`
	ProjectID         int     `json:"project_id"`
	ExpenseCategoryID int     `json:"expense_category_id"`
	Billable          bool    `json:"billable"`
	SpentAt           Date    `json:"spent_at"`
	IsLocked          bool    `json:"is_locked"`
}

type ExpenseResponse struct {
	Expense `json:"expense"`
}

// List requests list of expenses and returns response
func (c *ExpenseService) List() ([]ExpenseResponse, error) {
	resourceURL := "/expenses.json"
	var expenses []ExpenseResponse
	if err := c.get(resourceURL, &expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

// Find requests expense information for specified expense
// and returns response
func (c *ExpenseService) Find(expID int) (ExpenseResponse, error) {
	resourceURL := fmt.Sprintf("/expenses/%v.json", expID)
	var resp ExpenseResponse
	err := c.get(resourceURL, &resp)
	return resp, err
}

func (c *ExpenseService) ForProject(projID int, from, to time.Time) ([]ExpenseResponse, error) {
	resourceURL := fmt.Sprintf("/projects/%d/expenses.json?from=%s&to=%s", projID, from.Format("20060102"), to.Format("20060102"))
	var expenses []ExpenseResponse
	if err := c.get(resourceURL, &expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

func (c *ExpenseService) Update(e Expense) error {
	return c.Service.post("/expenses.json", ExpenseResponse{e})
}
