package v1

import (
	"reflect"

	"gorm.io/gorm"
)

type Entity interface {
	GetID() uint
	GetTableName() string
}

type User struct {
	gorm.Model
	Name  string `gorm:"unique,not null,size:255"`
	Email string `gorm:"size:255"`
}

func (u User) GetID() uint {
	return u.ID
}

func (u User) GetTableName() string {
	return "users"
}

type Project struct {
	gorm.Model
	Name        string `gorm:"unique,not null,size:255"`
	Description string `gorm:"size:1023"`
	OwnerId     uint   `gorm:"not null,index"`
}

func (p Project) GetID() uint {
	return p.ID
}

func (p Project) GetTableName() string {
	return "projects"
}

type Task struct {
	gorm.Model
	Name        string `gorm:"unique,not null,size:255"`
	Description string `gorm:"size:1023"`
	ProjectId   uint   `gorm:"not null,index"`
	AssigneeId  uint   `gorm:"not null,index"`
	Status      string
}

func (t Task) GetID() uint {
	return t.ID
}

func (t Task) GetTableName() string {
	return "tasks"
}

var _ Entity = &User{}
var _ Entity = &Project{}
var _ Entity = &Task{}

func GetTableName[T Entity]() string {
	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	switch typeName {
	case "User":
		user := new(User)
		return user.GetTableName()
	case "Project":
		project := new(Project)
		return project.GetTableName()
	case "Task":
		task := new(Task)
		return task.GetTableName()
	default:
		return typeName
	}
}
