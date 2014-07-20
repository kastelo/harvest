package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseTasks(apiClient *harvest.APIClient) {
	err, tasks := apiClient.Task.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", tasks)
		err, task := apiClient.Task.Find(tasks[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", task)
		}
	}
}
