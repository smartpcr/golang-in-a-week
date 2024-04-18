package store

import (
	"database/sql"

	"webapi/models"
)

type ProjectRepository struct {
	db *sql.DB
}

func (p *ProjectRepository) List() ([]models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectRepository) Get(id string) (models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectRepository) Create(item models.Project) (models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectRepository) Update(item models.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

var _ Repository[models.Project] = &ProjectRepository{}
