package harvest

import "fmt"

type TaskService struct {
	Service
}

type Task struct {
	Id                int         `json:"id"`
	Name              string      `json:"name"`
	BillableByDefault bool        `json:"billable-by-default"`
	Deactivated       bool        `json:"deactivated"`
	CreatedAt         HarvestDate `json:"created_at"`
	UpdateAt          HarvestDate `json:"updated_at"`
	DefaultHourlyRate float32     `json:"default-hourly-rate"`
	IsDefault         bool        `json:"is-default"`
}

type TaskResponse struct {
	Task `json:"task"`
}

func (c *TaskService) List() (err error, tasks []TaskResponse) {
	resourceURL := "/tasks.json"
	var resp []TaskResponse
	err = c.list(resourceURL, &resp)
	if err != nil {
		return err, resp
	}
	for _, element := range resp {
		tasks = append(tasks, element)
	}
	return err, tasks
}

func (c *TaskService) Find(taskID int) (err error, task TaskResponse) {
	resourceURL := fmt.Sprintf("/tasks/%v.json", taskID)
	var resp TaskResponse
	err = c.find(resourceURL, &resp)
	if err != nil {
		return
	}
	return err, resp
}
