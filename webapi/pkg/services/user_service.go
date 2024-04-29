package services

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"webapi/pkg/store"
	"webapi/schema/v1"
)

type UserService struct {
	Repo store.Repository[v1.User]
}

func (u *UserService) RegisterRoutes(a *fiber.App) {
	a.Get("/users", u.HandleList)
	a.Post("/users", u.HandleCreate)
	a.Get("/users/:id", u.HandleGet)
	a.Put("/users/:id", u.HandleUpdate)
	a.Delete("/users/:id", u.HandleDelete)
}

func (u *UserService) HandleList(c fiber.Ctx) error {
	users, err := u.Repo.List(c.Context())
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(users)
}

func (u *UserService) HandleCreate(c fiber.Ctx) error {
	user := new(v1.User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var validate = validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	createdUser, err := u.Repo.Create(c.Context(), user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(createdUser)
}

func (u *UserService) HandleGet(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	user, err := u.Repo.Get(c.Context(), uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(user)
}

func (u *UserService) HandleUpdate(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	user := new(v1.User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	user.ID = uint(id)
	var validate = validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = u.Repo.Update(c.Context(), user)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (u UserService) HandleDelete(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = u.Repo.Delete(c.Context(), uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

var _ Service = &UserService{}
