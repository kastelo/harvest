package harvest

import "fmt"

type TaskService struct {
	Service
}

type Task struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	BillableByDefault bool    `json:"billable_by_default"`
	Deactivated       bool    `json:"deactivated"`
	CreatedAt         Date    `json:"created_at"`
	UpdateAt          Date    `json:"updated_at"`
	DefaultHourlyRate float64 `json:"default_hourly_rate"`
	IsDefault         bool    `json:"is_default"`
}

type TaskResponse struct {
	Task `json:"task"`
}

// List requests list of tasks and returns response
func (c *TaskService) List() ([]TaskResponse, error) {
	resourceURL := "/tasks.json"
	var tasks []TaskResponse
	if err := c.get(resourceURL, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// Find requests task information for specified task and returns response
func (c *TaskService) Find(taskID int) (TaskResponse, error) {
	resourceURL := fmt.Sprintf("/tasks/%v.json", taskID)
	var resp TaskResponse
	err := c.get(resourceURL, &resp)
	return resp, err
}
