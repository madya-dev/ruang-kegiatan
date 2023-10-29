package service

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/users/repository"
	"madyasantosa/ruangkegiatan/model"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	GetAllUsers(ctx echo.Context) ([]model.User, int, error)
	GetUserByUsername(ctx echo.Context) (*dto.UserResponse, error)
	DeleteUser(ctx echo.Context) error
	ChangePassword(ctx echo.Context, r dto.ChangePasswordRequest, username string) error
	UserLogin(ctx echo.Context, r dto.UserLoginRequest) (*dto.UserResponse, error)
	CreateUser(ctx echo.Context, r dto.UserCreateRequest) (*dto.UserResponse, error)
	UpdateUser(ctx echo.Context, r dto.UserUpdateRequest) (*dto.UserResponse, error)
	UpdateRoleUser(ctx echo.Context, r dto.UserRoleUpdateRequest) (*dto.UserResponse, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(ur repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: ur,
		Validate:       validate,
	}
}
