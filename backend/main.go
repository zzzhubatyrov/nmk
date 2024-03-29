package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/repository/db"
	"backend/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	database, err := db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		_ = fmt.Errorf("failed to initialize db: %s", err.Error())
	}

	app := fiber.New()
	repos := repository.NewRepository(database)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	handlers.InitRoute(app)

	app.Listen(":5000")
}
