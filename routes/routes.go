package routes

import (
	"madyasantosa/ruangkegiatan/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController) {
	e.PATCH("/users/change-password", uc.ChangePassword)
	e.POST("/login", uc.UserLogin)
	e.POST("/register", uc.RegisterUser)
	e.GET("/users", uc.GetAllUsers)
	e.GET("/users/:username", uc.GetUserByUsername)
	e.PUT("/users/:username", uc.UpdateUser)
	e.DELETE("/users/:username", uc.DeleteUser)
	e.PATCH("/users/:username/role", uc.UpdateRoleUser)
}
