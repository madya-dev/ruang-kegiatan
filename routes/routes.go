package routes

import (
	"madyasantosa/ruangkegiatan/features/users/handler"
	"madyasantosa/ruangkegiatan/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uh handler.UserHandler) {
	e.POST("/login", uh.UserLogin)
	e.POST("/register", uh.RegisterUser)
	e.PATCH("/users/change-password", uh.ChangePassword, middleware.AuthMiddleware)
	e.GET("/users", uh.GetAllUsers, middleware.AuthMiddlewareAdmin)
	e.GET("/users/:username", uh.GetUserByUsername, middleware.AuthMiddleware)
	e.PUT("/users/:username", uh.UpdateUser, middleware.AuthMiddleware)
	e.DELETE("/users/:username", uh.DeleteUser, middleware.AuthMiddlewareAdmin)
	e.PATCH("/users/:username/role", uh.UpdateRoleUser, middleware.AuthMiddlewareAdmin)
}
