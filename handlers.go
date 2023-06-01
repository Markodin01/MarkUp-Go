package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// TaskHandler represents the HTTP handlers for tasks.
type TaskHandler struct {
	db *gorm.DB
}

// NewTaskHandler creates a new instance of TaskHandler with the given database connection.
func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

// GetTasks handles the GET /tasks endpoint and returns all tasks.
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	result := h.db.Find(&tasks)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, tasks, http.StatusOK)
}

// CreateTask handles the POST /tasks endpoint and creates a new task.
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result := h.db.Create(&task)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, task, http.StatusCreated)
}

// GetTask handles the GET /tasks/{id} endpoint and returns a specific task.
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var task Task
	result := h.db.First(&task, id)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, task, http.StatusOK)
}

// UpdateTask handles the PUT /tasks/{id} endpoint and updates a specific task.
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var task Task
	result := h.db.First(&task, id)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result = h.db.Save(&task)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, task, http.StatusOK)
}

// DeleteTask handles the DELETE /tasks/{id} endpoint and deletes a specific task.
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var task Task
	result := h.db.First(&task, id)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	result = h.db.Delete(&task)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// jsonResponse writes the given data as JSON response with the specified status code.
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

