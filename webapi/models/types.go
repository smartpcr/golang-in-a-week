package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Project struct {
	gorm.Model
	Name        string
	Description string
	OwnerId     uint
}

type Task struct {
	gorm.Model
	Name        string
	Description string
	ProjectId   string
	AssigneeId  string
	Status      string
}
