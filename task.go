package harvest

import "fmt"

type TaskService struct {
	Service
}

type Task struct {
	Id                int         `json:"id"`
	Name              string      `json:"name"`
	BillableByDefault bool        `json:"billable_by_default"`
	Deactivated       bool        `json:"deactivated"`
	CreatedAt         HarvestDate `json:"created_at"`
	UpdateAt          HarvestDate `json:"updated_at"`
	DefaultHourlyRate float32     `json:"default_hourly_rate"`
	IsDefault         bool        `json:"is_default"`
}

type TaskResponse struct {
	Task `json:"task"`
}

// List requests list of tasks and returns response
func (c *TaskService) List() (tasks []TaskResponse, err error) {
	resourceURL := "/tasks.json"
	var resp []TaskResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return resp, err
	}
	for _, element := range resp {
		tasks = append(tasks, element)
	}
	return tasks, err
}

// Find requests task information for specified task and returns response
func (c *TaskService) Find(taskID int) (task TaskResponse, err error) {
	resourceURL := fmt.Sprintf("/tasks/%v.json", taskID)
	var resp TaskResponse
	err = c.find(resourceURL, &resp)
	if err != nil {
		return
	}
	return resp, err
}
