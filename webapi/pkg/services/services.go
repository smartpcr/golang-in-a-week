package services

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Service interface {
	RegisterRoutes(a *fiber.App)
	HandleList(c fiber.Ctx) error
	HandleCreate(c fiber.Ctx) error
	HandleGet(c fiber.Ctx) error
	HandleUpdate(c fiber.Ctx) error
	HandleDelete(c fiber.Ctx) error
}

func handleError(c fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		return c.Status(e.Code).SendString(e.Message)
	} else {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
}
