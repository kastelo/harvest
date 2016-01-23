package harvest

import (
	"github.com/kr/pretty"
	"fmt"
	"testing"
	"os"
)

func TestTasks(t *testing.T) {
	apiClient := NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	fmt.Printf("\n%v", "task_test:")		//Used for debugging

	tasks, err := apiClient.Task.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Logf("\n%# v\n", pretty.Formatter(tasks))
		task, err := apiClient.Task.Find(tasks[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Logf("\n%# v\n", pretty.Formatter(task))
		}
	}
}
