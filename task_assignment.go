package harvest

import (
	"fmt"
	"time"
)

type TaskAssignmentService struct {
	Service
}

type TaskAssignment struct {
	ID         int       `json:"id"`
	Billable   bool      `json:"billable"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"updated_at"`
	HourlyRate float64   `json:"hourly_rate"`
	Project    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"project"`
	Task struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"task"`
}

type TaskAssignmentResponse struct {
	TaskAssignment `json:"task_assignment"`
}

func (c *TaskAssignmentService) ListByProject(projectID int) ([]TaskAssignmentResponse, error) {
	resourceURL := fmt.Sprintf("/projects/%d/task_assignments.json", projectID)
	var tasks []TaskAssignmentResponse
	if err := c.get(resourceURL, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
