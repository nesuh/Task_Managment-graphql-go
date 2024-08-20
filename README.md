# Tasks Management System

This project is a Tasks Management System that allows you to create, read, update, and delete tasks using a Go backend with Hasura Cloud for GraphQL operations. The system stores tasks in a PostgreSQL database managed through Hasura.

## Table of Contents

- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
- [Environment Variables](#environment-variables)
- [Running the Application](#running-the-application)
- [CRUD Operations](#crud-operations)
  - [Create a Task](#create-a-task)
  - [Read Tasks](#read-tasks)
  - [Update a Task](#update-a-task)
  - [Delete a Task](#delete-a-task)

## Project Structure

```bash
.
├── main.go                   # Entry point of the application
├── utils/
│   ├── tasks.go              # Contains functions to handle CRUD operations
│   └── utils.go              # Contains utility functions (e.g., making requests)
├── go.mod                    # Go module file
├── go.sum                    # Go dependencies
└── README.md                 # Project documentation

Set Up Hasura Cloud

Log in to your Hasura Cloud account.

Create a new project and connect it to your PostgreSQL database.

Create table Name Tasks
Environment Variables
Set up the following environment variables for your project. You can create a .env file in the root directory:

env
Copy code
HASURA_GRAPHQL_ENDPOINT=https://your-hasura-instance.hasura.app/v1/graphql
HASURA_ADMIN_SECRET=your-admin-secret
Replace your-hasura-instance and your-admin-secret with your actual Hasura Cloud instance URL and admin secret.

Running the Application
To run the Go application, execute:

bash
Copy code
go run main.go
CRUD Operations
Create a Task
To create a task, use Postman or any other API client to send the following request:

URL: https://your-hasura-instance.hasura.app/v1/graphql

Method: POST

Headers:

Content-Type: application/json
x-hasura-admin-secret: your-admin-secret
Body:

json
Copy code
{
  "query": "mutation InsertTask($task_name: String!, $description: String, $start_time: timestamptz, $end_time: timestamptz, $priority: String, $status: String) { insert_tasks_one(object: {task_name: $task_name, description: $description, start_time: $start_time, end_time: $end_time, priority: $priority, status: $status}) { id task_name description start_time end_time priority status } }",
  "variables": {
    "task_name": "Start Go Project",
    "description": "I am starting a Go project.",
    "start_time": "2024-08-17T09:00:00Z",
    "end_time": "2024-08-17T17:00:00Z",
    "priority": "High",
    "status": "Not Started"
  }
}
Read Tasks
To fetch all tasks, use the following query:

Body:

json
Copy code
{
  "query": "query { tasks { id task_name description start_time end_time priority status } }"
}
Update a Task
To update an existing task, use this mutation:

Body:

json
Copy code
{
  "query": "mutation UpdateTask($id: Int!, $status: String!) { update_tasks_by_pk(pk_columns: {id: $id}, _set: {status: $status}) { id task_name description start_time end_time priority status } }",
  "variables": {
    "id": 1,
    "status": "In Progress"
  }
}
Delete a Task
To delete a task, use this mutation:

Body:

json
Copy code
{
  "query": "mutation DeleteTask($id: Int!) { delete_tasks_by_pk(id: $id) { id } }",
  "variables": {
    "id": 1
  }
}


