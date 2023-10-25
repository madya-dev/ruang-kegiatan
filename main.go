package main

import (
	"fmt"
	"madyasantosa/ruangkegiatan/config"
	"madyasantosa/ruangkegiatan/controller"
	"madyasantosa/ruangkegiatan/pkg/database"
	"madyasantosa/ruangkegiatan/repository"
	"madyasantosa/ruangkegiatan/routes"
	"madyasantosa/ruangkegiatan/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := *config.InitConfig()

	validate := validator.New()

	db := database.InitDB(config)
	database.Migrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "PORT: "+config.AppPort)
	})

	routes.UserRoutes(e, userController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppPort)))
}
