package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

const graphqlEndpoint = "https://obliging-crawdad-72.hasura.app/v1/graphql"
const adminSecret = "YioKNtyq13pmTY61gxHW1FToNW9Pscz1Wgjv0NRscLsYvIe6QLV2eS7A7ZEzWSY4" // Replace with your Hasura admin secret

// InsertTask inserts a new task into the database
func InsertTask(task Task) error {
	query := `mutation InsertTask($task_name: String!, $description: String, $start_time: timestamptz, $end_time: timestamptz, $priority: String, $status: String) {
		insert_tasks_one(object: {task_name: $task_name, description: $description, start_time: $start_time, end_time: $end_time, priority: $priority, status: $status}) {
			id
		}
	}`

	variables := map[string]interface{}{
		"task_name":   task.TaskName,
		"description": task.Description,
		"start_time":  task.StartTime,
		"end_time":    task.EndTime,
		"priority":    task.Priority,
		"status":      task.Status,
	}

	return graphqlQuery(query, variables)
}

// GetTasks retrieves all tasks from the database
func GetTasks() ([]Task, error) {
	query := `query GetTasks {
		tasks {
			id
			task_name
			description
			start_time
			end_time
			priority
			status
		}
	}`

	var result struct {
		Data struct {
			Tasks []Task `json:"tasks"`
		} `json:"data"`
	}

	err := graphqlQueryWithResponse(query, nil, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.Tasks, nil
}

// UpdateTask updates an existing task in the database
func UpdateTask(id int, task Task) error {
	query := `mutation UpdateTask($id: Int!, $task_name: String, $description: String, $start_time: timestamptz, $end_time: timestamptz, $priority: String, $status: String) {
		update_tasks_by_pk(pk_columns: {id: $id}, _set: {task_name: $task_name, description: $description, start_time: $start_time, end_time: $end_time, priority: $priority, status: $status}) {
			id
		}
	}`

	variables := map[string]interface{}{
		"id":          id,
		"task_name":   task.TaskName,
		"description": task.Description,
		"start_time":  task.StartTime,
		"end_time":    task.EndTime,
		"priority":    task.Priority,
		"status":      task.Status,
	}

	return graphqlQuery(query, variables)
}

// DeleteTask deletes a task from the database by ID
func DeleteTask(id int) error {
	query := `mutation DeleteTask($id: Int!) {
		delete_tasks_by_pk(id: $id) {
			id
		}
	}`

	variables := map[string]interface{}{
		"id": id,
	}

	return graphqlQuery(query, variables)
}

// graphqlQuery sends a query to the GraphQL endpoint without expecting a response
func graphqlQuery(query string, variables map[string]interface{}) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", graphqlEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", adminSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body)) // For debugging purposes
	return nil
}

// graphqlQueryWithResponse sends a query to the GraphQL endpoint and expects a response
func graphqlQueryWithResponse(query string, variables map[string]interface{}, result interface{}) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", graphqlEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", adminSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}

	return nil
}
