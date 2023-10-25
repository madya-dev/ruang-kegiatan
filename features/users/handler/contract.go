package handler

import (
	"madyasantosa/ruangkegiatan/features/users/service"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetAllUsers(ctx echo.Context) error
	GetUserByUsername(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	RegisterUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	UpdateRoleUser(ctx echo.Context) error
	ChangePassword(ctx echo.Context) error
	UserLogin(ctx echo.Context) error
}

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return &UserHandlerImpl{UserService: us}
}
