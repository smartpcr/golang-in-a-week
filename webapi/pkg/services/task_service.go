package services

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"webapi/pkg/store"
	"webapi/schema/v1"
)

type TaskService struct {
	Repo store.Repository[v1.Task]
}

func (t *TaskService) RegisterRoutes(a *fiber.App) {
	a.Get("/tasks", t.HandleList)
	a.Post("/tasks", t.HandleCreate)
	a.Get("/tasks/:id", t.HandleGet)
	a.Put("/tasks/:id", t.HandleUpdate)
	a.Delete("/tasks/:id", t.HandleDelete)
}

func (t *TaskService) HandleList(c fiber.Ctx) error {
	tasks, err := t.Repo.List(c.Context())
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(tasks)
}

func (t *TaskService) HandleCreate(c fiber.Ctx) error {
	task := new(v1.Task)
	if err := c.Bind().Body(task); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var validate = validator.New()
	if err := validate.Struct(task); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	createdTask, err := t.Repo.Create(c.Context(), task)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(createdTask)
}

func (t *TaskService) HandleGet(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	task, err := t.Repo.Get(c.Context(), uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(task)
}

func (t *TaskService) HandleUpdate(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	task := new(v1.Task)
	if err = c.Bind().Body(task); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	task.ID = uint(id)
	var validate = validator.New()
	if err := validate.Struct(task); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = t.Repo.Update(c.Context(), task)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (t *TaskService) HandleDelete(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = t.Repo.Delete(c.Context(), uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

var _ Service = &TaskService{}
