package main

import (
    "gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Category struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	gorm.Model
	ID        int       `json:"id"`
	TaskID    int       `json:"task_id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Project struct {
	gorm.Model
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

type Admin struct{
	gorm.Model
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`			
}
