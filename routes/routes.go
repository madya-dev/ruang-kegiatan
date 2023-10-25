package routes

import (
	"madyasantosa/ruangkegiatan/controller"
	"madyasantosa/ruangkegiatan/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController) {
	e.POST("/login", uc.UserLogin)
	e.POST("/register", uc.RegisterUser)
	e.PATCH("/users/change-password", uc.ChangePassword, middleware.AuthMiddleware)
	e.GET("/users", uc.GetAllUsers, middleware.AuthMiddlewareAdmin)
	e.GET("/users/:username", uc.GetUserByUsername, middleware.AuthMiddleware)
	e.PUT("/users/:username", uc.UpdateUser, middleware.AuthMiddleware)
	e.DELETE("/users/:username", uc.DeleteUser, middleware.AuthMiddlewareAdmin)
	e.PATCH("/users/:username/role", uc.UpdateRoleUser, middleware.AuthMiddlewareAdmin)
}
