package handler

import (
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute(app *fiber.App) fiber.Handler {

	app.Get("", func(c *fiber.Ctx) error {
		return nil
	})

	return nil
}
