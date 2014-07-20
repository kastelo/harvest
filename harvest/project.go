package harvest

import "fmt"

type ProjectService struct {
	Service
}

type Project struct {
	Id                               int         `json:"id"`
	ClientId                         int         `json:"client_id"`
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
}

// Wrapper for simple unmarshalling of JSON data
type ProjectResponse struct {
	Project `json:"project"`
}

func (c *ProjectService) List() (err error, projects []ProjectResponse) {
	resourceURL := "/projects.json"
	var resp []ProjectResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return err, resp
	}
	for _, element := range resp {
		projects = append(projects, element)
	}
	return err, projects
}

func (c *ProjectService) Find(projectID int) (err error, project ProjectResponse) {
	resourceURL := fmt.Sprintf("/projects/%v.json", projectID)
	var resp ProjectResponse
	err = c.find(resourceURL, &resp)
	return err, resp
}
