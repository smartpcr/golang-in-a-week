package store

import (
	"database/sql"

	"webapi/models"
)

type UserRepository struct {
	db *sql.DB
}

var _ Repository[models.User] = &UserRepository{}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) List() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Get(id string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Create(item models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Update(item models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
