package main

import (
	"backend/internal/handler"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/repository/db"
	"backend/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		_ = fmt.Errorf("error initializing configs: %s", err.Error())
	}

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

	models := []interface{}{
		&model.User{},
	}

	migrator := database.Migrator()
	_ = migrator.DropTable(models...)
	_ = database.AutoMigrate(models...)

	repos := repository.NewRepository(database)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	handlers.InitRoute(app)
	app.Listen(":5000")
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}
