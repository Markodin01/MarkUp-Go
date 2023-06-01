package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
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

