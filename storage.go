package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(connStr string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AutoMigrate() error {
	err := s.db.AutoMigrate(&Task{}, &User{}, &Category{}, &Comment{}, &Project{}, &Admin{})
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := s.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (s *Storage) CreateTask(task *Task) error {
	result := s.db.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) GetTask(id int) (Task, error) {
	var task Task
	result := s.db.First(&task, id)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}
