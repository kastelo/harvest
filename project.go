package harvest

import "fmt"

type ProjectService struct {
	Service
}

type Project struct {
	ID                               int         `json:"id"`
	ClientID                         int         `json:"client_id"`
	Name                             string      `json:"name"`
	Code                             string      `json:"code"`
	Notes                            string      `json:"notes"`
	BillBy                           string      `json:"bill_by"`
	BudgetBy                         string      `json:"budget_by"`
	Active                           bool        `json:"active"`
	CostBudgetIncludeExpenses        bool        `json:"cost_budget_include_expenses"`
	Billable                         bool        `json:"billable"`
	ShowBudgetToAll                  bool        `json:"show_budget_to_all"`
	CostBudget                       float32     `json:"cost_budget"`
	HourlyRate                       float32     `json:"hourly_rate"`
	Budget                           float32     `json:"budget"`
	NotifyWhenOverBudget             float32     `json:"notify_when_overbudget"`
	OverBudgetNotificationPercentage float32     `json:"over_budget_notification_percentage"`
	OverBudgetNotifiedAt             HarvestDate `json:"over_budget_notified_at"`
	CreatedAt                        HarvestDate `json:"created_at"`
	UpdateAt                         HarvestDate `json:"updated_at"`
	HintEarliestRecordAt             HarvestDate `json:"hint_earliest_record_at"`
	HintLatestRecordAt               HarvestDate `json:"hint_latest_record_at"`
	StartsOn                         HarvestDate `json:"starts_on"`
	EndsOn                           HarvestDate `json:"ends_on"`
	Estimate                         float32     `json:"estimate"`
	EstimateBy                       string      `json:"estimate_by"`
}

// Wrapper for simple unmarshalling of JSON data
type ProjectResponse struct {
	Project `json:"project"`
}

// List requests list of projects and returns response
func (c *ProjectService) List() (projects []ProjectResponse, err error) {
	resourceURL := "/projects.json"
	var resp []ProjectResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		projects = append(projects, element)
	}
	return projects, err
}

// Find requests project information for specified project
func (c *ProjectService) Find(projectID int) (project ProjectResponse, err error) {
	resourceURL := fmt.Sprintf("/projects/%v.json", projectID)
	var resp ProjectResponse
	err = c.find(resourceURL, &resp)
	return resp, err
}
