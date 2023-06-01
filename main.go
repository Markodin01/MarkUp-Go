package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var storage *Storage

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
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))

}

// Define your endpoint handlers

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Retrieve tasks from your storage or database
	tasks, err := storage.GetAllTasks()

	// Set the response content type to JSON
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

func CreateTask(w http.ResponseWriter, r *http.Request) {
	// Handle the POST /tasks endpoint
	// ...
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	// Handle the GET /tasks/{id} endpoint
	// ...
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Handle the PUT /tasks/{id} endpoint
	// ...
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Handle the DELETE /tasks/{id} endpoint
	// ...
}