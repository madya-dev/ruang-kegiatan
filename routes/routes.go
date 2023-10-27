package routes

import (
	reservationHandler "madyasantosa/ruangkegiatan/features/reservation/handler"
	roomHandler "madyasantosa/ruangkegiatan/features/rooms/handler"
	userHandler "madyasantosa/ruangkegiatan/features/users/handler"
	"madyasantosa/ruangkegiatan/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uh userHandler.UserHandler) {
	e.POST("/login", uh.UserLogin)
	e.POST("/register", uh.RegisterUser)
	e.PATCH("/users/change-password", uh.ChangePassword, middleware.AuthMiddleware)
	e.GET("/users", uh.GetAllUsers, middleware.AuthMiddlewareAdmin)
	e.GET("/users/:username", uh.GetUserByUsername, middleware.AuthMiddleware)
	e.PUT("/users/:username", uh.UpdateUser, middleware.AuthMiddleware)
	e.DELETE("/users/:username", uh.DeleteUser, middleware.AuthMiddlewareAdmin)
	e.PATCH("/users/:username/role", uh.UpdateRoleUser, middleware.AuthMiddlewareAdmin)
}

func RoomRoutes(e *echo.Echo, rh roomHandler.RoomHandler) {
	e.GET("/rooms", rh.GetAllRooms)
	e.POST("/rooms", rh.CreateRoom, middleware.AuthMiddlewareAdmin)
	e.PUT("/rooms/:id", rh.UpdateRoom, middleware.AuthMiddlewareAdmin)
	e.DELETE("/rooms/:id", rh.DeleteRoom, middleware.AuthMiddlewareAdmin)
}

func ReservationRoutes(e *echo.Echo, rh reservationHandler.ReservationHanlder) {
	e.GET("/reservations", rh.GetAllReservation)
	e.DELETE("/reservations/:id", rh.DeleteReservation, middleware.AuthMiddleware)
	e.POST("/reservations", rh.CreateReservation, middleware.AuthMiddleware)
	e.PUT("/reservations/:id", rh.UpdateReservation, middleware.AuthMiddleware)
}
