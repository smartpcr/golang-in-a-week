package v1

import (
	"reflect"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type Entity interface {
	GetID() uint
	GetTableName() string
}

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);unique;not null"`
	Email string `gorm:"type:varchar(255)"`
}

func GenerateTestUsers(count int, users []*User) {
	for i := 0; i < count; i++ {
		users = append(users, &User{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		})
	}
}

func (u User) GetID() uint {
	return u.ID
}

func (u User) GetTableName() string {
	return "users"
}

type Project struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);unique;not null"`
	Description string `gorm:"type:varchar(1023)"`
	OwnerId     uint   `gorm:"not null,index"`
}

func GenerateTestProjects(count int, existingUsers []*User, projects []*Project) {
	for i := 0; i < count; i++ {
		owner := existingUsers[gofakeit.Number(0, len(existingUsers)-1)]
		projects = append(projects, &Project{
			Name:        gofakeit.Name(),
			Description: gofakeit.Sentence(200),
			OwnerId:     owner.ID,
		})
	}
}

func (p Project) GetID() uint {
	return p.ID
}

func (p Project) GetTableName() string {
	return "projects"
}

type TaskStatus string

const (
	New        TaskStatus = "new"
	Active     TaskStatus = "active"
	InProgress TaskStatus = "in_progress"
	Done       TaskStatus = "done"
)

type Task struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(255);unique;not null"`
	Description string     `gorm:"type:varchar(1023)"`
	ProjectId   uint       `gorm:"not null,index"`
	AssigneeId  uint       `gorm:"not null,index"`
	Status      TaskStatus `gorm:"type:varchar(50);check:status_check,status in ('new', 'active', 'in_progress', 'done')"`
}

func GenerateTestTasks(count int, existingUsers []*User, existingProjects []*Project, tasks []*Task) {
	for i := 0; i < count; i++ {
		assignee := existingUsers[gofakeit.Number(0, len(existingUsers)-1)]
		project := existingProjects[gofakeit.Number(0, len(existingProjects)-1)]
		tasks = append(tasks, &Task{
			Name:        gofakeit.Name(),
			Description: gofakeit.Sentence(200),
			ProjectId:   project.ID,
			AssigneeId:  assignee.ID,
			Status:      New,
		})
	}
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
