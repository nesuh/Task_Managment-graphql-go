package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Task struct {
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

func InsertTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Your logic to insert the task into the database goes here
	// For example:
	// err := insertTaskIntoDB(task)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Task added successfully"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/task", InsertTask)

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"}, // Allow only your frontend origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
