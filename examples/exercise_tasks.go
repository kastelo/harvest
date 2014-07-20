package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseTasks(apiClient *harvest.APIClient) {
	tasks, err := apiClient.Task.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", tasks)
		task, err := apiClient.Task.Find(tasks[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", task)
		}
	}
}
