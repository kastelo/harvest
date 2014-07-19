package harvest

import (
	"fmt"
)

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
	Task Task
}

func (c *TaskService) List() (err error, tasks []Task) {
	resourceURL := "/tasks.json"
	var taskResponse []TaskResponse
	err = c.list(resourceURL, &taskResponse)
	if err != nil {
		return
	}

	for _, element := range taskResponse {
		tasks = append(tasks, element.Task)
	}
	return
}

func (c *TaskService) Find(taskID int) (err error, task Task) {
	resourceURL := fmt.Sprintf("/tasks/%v.json", taskID)
	var taskResponse TaskResponse
	err = c.find(resourceURL, &taskResponse)
	if err != nil {
		return
	}
	task = taskResponse.Task
	return
}
