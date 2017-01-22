package harvest

import (
	"fmt"
	"github.com/kr/pretty"
	"os"
	"testing"
)

func TestTasks(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "task_test:") //Used for debugging

	tasks, err := apiClient.Task.List()
	if err != nil {
		t.Fatalf("\n%v\n", err)
	} else {
		t.Logf("\n%# v\n", pretty.Formatter(tasks))
		task, err := apiClient.Task.Find(tasks[0].ID)
		if err != nil {
			t.Fatalf("\n%v\n", err)
		} else {
			t.Logf("\n%# v\n", pretty.Formatter(task))
		}
	}
}
