package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	handlers.InitRoute(app)

	app.Listen(":5000")
}
