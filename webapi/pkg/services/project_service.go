package services

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"webapi/pkg/store"
	"webapi/schema/v1"
)

type ProjectService struct {
	Repo store.Repository[v1.Project]
}

func (p *ProjectService) RegisterRoutes(a *fiber.App) {
	a.Get("/projects", p.HandleList)
	a.Post("/projects", p.HandleCreate)
	a.Get("/projects/:id", p.HandleGet)
	a.Put("/projects/:id", p.HandleUpdate)
	a.Delete("/projects/:id", p.HandleDelete)
}

func (p *ProjectService) HandleList(c fiber.Ctx) error {
	projects, err := p.Repo.List(c.Context())
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(projects)
}

func (p *ProjectService) HandleCreate(c fiber.Ctx) error {
	project := new(v1.Project)
	if err := c.Bind().Body(project); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var validate = validator.New()
	if err := validate.Struct(project); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	createdProject, err := p.Repo.Create(c.Context(), project)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(createdProject)
}

func (p *ProjectService) HandleGet(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	project, err := p.Repo.Get(c.Context(), uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(project)
}

func (p *ProjectService) HandleUpdate(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	project := new(v1.Project)
	if err := c.Bind().Body(project); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	project.ID = uint(id)
	var validate = validator.New()
	if err := validate.Struct(project); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = p.Repo.Update(c.Context(), project)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (p *ProjectService) HandleDelete(c fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = p.Repo.Delete(c.Context(), uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

var _ Service = &ProjectService{}
