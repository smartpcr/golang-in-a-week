package math

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserValid(t *testing.T) {
	user := User{
		Name:  "John",
		Age:   30,
		Email: "john@email.com",
	}
	err := ValidateUser(&user)
	assert.NoError(t, err, "User should be valid")
}

func TestUserInvalid(t *testing.T) {
	user := User{
		Name:  "",
		Age:   30,
		Email: "",
	}
	err := ValidateUser(&user)
	assert.Error(t, err, "User should be invalid")
	assert.Len(t, err, 2, "User should have 2 errors")
	for _, validationErr := range err.(validator.ValidationErrors) {
		switch validationErr.Field() {
		case "Name":
			assert.Contains(t, validationErr.Error(), "required")
		case "Email":
			assert.Contains(t, validationErr.Error(), "required")
		}
	}
}
