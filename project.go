package harvest

import "fmt"

type ProjectService struct {
	Service
}

type Project struct {
	ID                               int     `json:"id"`
	ClientID                         int     `json:"client_id"`
	Name                             string  `json:"name"`
	Code                             string  `json:"code"`
	Notes                            string  `json:"notes"`
	BillBy                           string  `json:"bill_by"`
	BudgetBy                         string  `json:"budget_by"`
	Active                           bool    `json:"active"`
	CostBudgetIncludeExpenses        bool    `json:"cost_budget_include_expenses"`
	Billable                         bool    `json:"billable"`
	ShowBudgetToAll                  bool    `json:"show_budget_to_all"`
	CostBudget                       float64 `json:"cost_budget"`
	HourlyRate                       float64 `json:"hourly_rate"`
	Budget                           float64 `json:"budget"`
	NotifyWhenOverBudget             float64 `json:"notify_when_overbudget"`
	OverBudgetNotificationPercentage float64 `json:"over_budget_notification_percentage"`
	OverBudgetNotifiedAt             Date    `json:"over_budget_notified_at"`
	CreatedAt                        Date    `json:"created_at"`
	UpdatedAt                        Date    `json:"updated_at"`
	HintEarliestRecordAt             Date    `json:"hint_earliest_record_at"`
	HintLatestRecordAt               Date    `json:"hint_latest_record_at"`
	StartsOn                         Date    `json:"starts_on"`
	EndsOn                           Date    `json:"ends_on"`
	Estimate                         float64 `json:"estimate"`
	EstimateBy                       string  `json:"estimate_by"`
	IsFixedFee                       bool    `json:"is_fixed_fee"`
	Fee                              float64 `json:"fee"`
}

// Wrapper for simple unmarshalling of JSON data
type ProjectResponse struct {
	Project `json:"project"`
}

// List requests list of projects and returns response
func (c *ProjectService) List() ([]ProjectResponse, error) {
	resourceURL := "/projects.json"
	var projects []ProjectResponse
	if err := c.get(resourceURL, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// Find requests project information for specified project
func (c *ProjectService) Find(projectID int) (ProjectResponse, error) {
	resourceURL := fmt.Sprintf("/projects/%v.json", projectID)
	var resp ProjectResponse
	err := c.get(resourceURL, &resp)
	return resp, err
}
