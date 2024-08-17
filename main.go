package main

import (
	"TaskManagement/utils" // Adjust this to your actual module path
	"fmt"
	"log"
)

func main() {
	// Create a new task
	task := utils.Task{
		TaskName:    "Finish Go Project",
		Description: "Complete the Go project by end of the day",
		StartTime:   "2024-08-17T09:00:00Z",
		EndTime:     "2024-08-17T17:00:00Z",
		Priority:    "High",
		Status:      "Not Started",
	}

	err := utils.InsertTask(task)
	if err != nil {
		log.Fatalf("Failed to insert task: %v", err)
	}
	fmt.Println("Task inserted successfully.")

	// Read tasks
	tasks, err := utils.GetTasks()
	if err != nil {
		log.Fatalf("Failed to get tasks: %v", err)
	}
	fmt.Println("Tasks:", tasks)

	// Update a task
	taskToUpdate := tasks[0]
	taskToUpdate.Status = "In Progress"
	err = utils.UpdateTask(taskToUpdate.ID, taskToUpdate)
	if err != nil {
		log.Fatalf("Failed to update task: %v", err)
	}
	fmt.Println("Task updated successfully.")

	// Delete a task
	err = utils.DeleteTask(taskToUpdate.ID)
	if err != nil {
		log.Fatalf("Failed to delete task: %v", err)
	}
	fmt.Println("Task deleted successfully.")
}
