package store

import (
	"database/sql"

	"webapi/models"
)

type TaskRepository struct {
	db *sql.DB
}

func (t *TaskRepository) List() ([]models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) Get(id string) (models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) Create(item models.Task) (models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) Update(item models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

var _ Repository[models.Task] = &TaskRepository{}
