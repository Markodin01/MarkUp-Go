package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"

	"github.com/spf13/cobra"
)

type Task struct {
	gorm.Model
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

var rootCmd = &cobra.Command{
	Use:   "taskmanager",
	Short: "A simple task manager CLI",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for creating a task
		taskName := args[0]
		err := invokeCreateTaskAPI(taskName)
		if err != nil {
			log.Fatalf("Failed to create task: %v", err)
		}
		fmt.Printf("Task created successfully: %s\n", taskName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Define flags using Viper
	createCmd.Flags().StringP("category", "c", "", "Set the task category")

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func invokeCreateTaskAPI(taskName string) error {
	// Define the task payload
	task := Task{
		ID:          11,
		Title:       taskName,
		Description: "sample description",
		Status:      "on hold",
		DueDate:     time.Now(),
	}

	// Marshal the task payload into JSON
	body, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %v", err)
	}

	// Make an HTTP POST request to create the task
	resp, err := http.Post("http://localhost:8080/tasks", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create task: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create task. Status: %s", resp.Status)
	}

	// Task created successfully
	return nil
}
