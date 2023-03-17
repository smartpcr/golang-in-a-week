package math

import "github.com/go-playground/validator/v10"

type User struct {
	Name  string `validate:"required"`
	Age   int    `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}

func ValidateUser(user *User) error {
	validate := validator.New()
	err := validate.Struct(user)
	return err
}
