package models

type User struct {
	Id    string
	Name  string
	Email string
}

type Project struct {
	Id          string
	Name        string
	Description string
	OwnerId     string
}

type Task struct {
	Id          string
	Name        string
	Description string
	ProjectId   string
	AssigneeId  string
	Status      string
}
