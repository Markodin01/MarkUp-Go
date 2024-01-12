package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=marekgaj sslmode=disable"

	storage, err := NewStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	err = storage.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router using the Gorilla Mux router
	router := mux.NewRouter()

	// Define your endpoints
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllTasks(w, r, storage)
		case http.MethodPost:
			CreateTask(w, r, storage)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods("GET", "POST")

	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetTask(w, r, storage)
		case http.MethodPut:
			UpdateTask(w, r, storage)
		case http.MethodDelete:
			DeleteTask(w, r, storage)

		}
	})

	// Start the server on port 8080
	print("The sever started")
	log.Fatal(http.ListenAndServe(":8080", router))

}

// Defining endpoint handlers

func UpdateTask(w http.ResponseWriter, r *http.Request, storage *Storage) {
	// Parse the request body to get the task data
	var task Task
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Perform any necessary validation on the task data
	if task.Title == "" {
		http.Error(w, "Task title is required", http.StatusBadRequest)
		return
	}

	// Save the task to the storage or database
	err = storage.UpdateTask(&task, taskID)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	// Set the response status to 201 Created
	w.WriteHeader(http.StatusCreated)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request, storage *Storage) {
	// Retrieve tasks from your storage or database
	tasks, err := storage.GetAllTasks()
	if err != nil {
		// Handle the error and return an appropriate response
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Convert tasks to JSON
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		// Handle the error and return an appropriate response
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonData)
}

func CreateTask(w http.ResponseWriter, r *http.Request, storage *Storage) {
	// Parse the request body to get the task data
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Perform any necessary validation on the task data
	if task.Title == "" {
		http.Error(w, "Task title is required", http.StatusBadRequest)
		return
	}

	// Save the task to the storage or database
	err = storage.CreateTask(&task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	// Set the response status to 201 Created
	w.WriteHeader(http.StatusCreated)
}

func GetTask(w http.ResponseWriter, r *http.Request, storage *Storage) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := storage.GetTask(taskID)
	if err != nil {
		http.Error(w, "Failed to retrieve task", http.StatusInternalServerError)
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Convert task to JSON
	jsonData, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Failed to serialize task", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonData)
	fmt.Printf("GET /tasks/%d - Status: %d\n", taskID, http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request, storage *Storage) {

	var task Task

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = storage.DeleteTask(&task, taskID)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
