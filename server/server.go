package server

import (
	"fmt"
	"madyasantosa/ruangkegiatan/config"
	userHandlerPkg "madyasantosa/ruangkegiatan/features/users/handler"
	userRepositoryPkg "madyasantosa/ruangkegiatan/features/users/repository"
	userServicePkg "madyasantosa/ruangkegiatan/features/users/service"
	"madyasantosa/ruangkegiatan/pkg/database"
	"madyasantosa/ruangkegiatan/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() error {
	config := *config.InitConfig()

	validate := validator.New()

	db := database.InitDB(config)
	database.Migrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userRepository := userRepositoryPkg.NewUserRepository(db)
	userService := userServicePkg.NewUserService(userRepository, validate)
	userHandler := userHandlerPkg.NewUserHandler(userService)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "PORT: "+config.AppPort)
	})

	routes.UserRoutes(e, userHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppPort)))

	return nil
}
