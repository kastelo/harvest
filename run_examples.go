package main

import (
	"os"

	"./examples"
	"./harvest"
)

func main() {
	apiClient := harvest.NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	examples.ExerciseClients(apiClient)
	examples.ExerciseInvoices(apiClient)
	examples.ExercisePeople(apiClient)
	examples.ExerciseProjects(apiClient)
	examples.ExerciseAccounts(apiClient)
	examples.ExerciseTasks(apiClient)
	examples.ExerciseContacts(apiClient)
	examples.ExerciseExpenseCategories(apiClient)
	examples.ExercisePersonEntries(apiClient)
	examples.ExerciseProjectEntries(apiClient)
}
