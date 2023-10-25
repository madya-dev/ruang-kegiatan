package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"madyasantosa/ruangkegiatan/model"
	"madyasantosa/ruangkegiatan/repository"
	"strconv"

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

func (s *UserServiceImpl) GetAllUsers(ctx echo.Context) ([]model.User, int, error) {
	params := ctx.QueryParams()
	search := params.Get("s")
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return nil, 0, fmt.Errorf("Params limit not valid")
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return nil, 0, fmt.Errorf("Params offset not valid")
	}

	users, total, err := s.UserRepository.GetAllUsers(offset, limit, search)
	if err != nil {
		return nil, total, fmt.Errorf("Internal Server Error")
	}
	if len(search) > 0 && total <= 0 {
		return nil, total, fmt.Errorf("Users not found")
	}
	return users, total, nil
}

func (s *UserServiceImpl) GetUserByUsername(ctx echo.Context) (*dto.UserResponse, error) {
	user, err := s.UserRepository.GetUserByUsername(ctx.Param("username"))
	if err != nil {
		return nil, err
	}
	userResponse := helper.ConvertToUserResponse(user)
	return userResponse, nil
}
func (s *UserServiceImpl) DeleteUser(ctx echo.Context) error {
	err := s.UserRepository.DeleteUser(ctx.Param("username"))
	if err != nil {
		return err
	}

	return nil
}
func (s *UserServiceImpl) CreateUser(ctx echo.Context, r dto.UserCreateRequest) (*dto.UserResponse, error) {
	err := s.Validate.Struct(r)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.GetUserByUsername(r.Username)
	if existingUser != nil {
		return nil, fmt.Errorf("Username Already Exists")
	}

	user := helper.UserCreateRequestToUserModel(r)
	user.Password = helper.HashPassword(user.Password)

	res, err := s.UserRepository.CreateUser(user)

	if err != nil {
		return nil, fmt.Errorf("Error When Creating User %s:", err.Error())
	}
	userResponse := helper.ConvertToUserResponse(res)
	return userResponse, nil
}
func (s *UserServiceImpl) UpdateUser(ctx echo.Context, r dto.UserUpdateRequest) (*dto.UserResponse, error) {
	err := s.Validate.Struct(r)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.GetUserByUsername(ctx.Param("username"))
	if existingUser == nil {
		return nil, fmt.Errorf("User not found")
	}

	user := helper.UserUpdateRequestToUserModel(r)

	res, err := s.UserRepository.UpdateUser(user, ctx.Param("username"))
	if err != nil {
		return nil, fmt.Errorf("Error when updating user: %s", err.Error())
	}
	userResponse := helper.ConvertToUserResponse(res)

	return userResponse, nil
}
func (s *UserServiceImpl) UpdateRoleUser(ctx echo.Context, r dto.UserRoleUpdateRequest) (*dto.UserResponse, error) {
	err := s.Validate.Struct(r)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.GetUserByUsername(ctx.Param("username"))
	if existingUser == nil {
		return nil, fmt.Errorf("User not found")
	}

	user := helper.UserRoleUpdateRequestToUserModel(r)

	res, err := s.UserRepository.UpdateRoleUser(user, ctx.Param("username"))
	if err != nil {
		return nil, fmt.Errorf("Error when updating user: %s", err.Error())
	}
	existingUser.Role = res.Role
	userResponse := helper.ConvertToUserResponse(existingUser)

	return userResponse, nil
}

func (s *UserServiceImpl) ChangePassword(ctx echo.Context, r dto.ChangePasswordRequest, username string) error {
	err := s.Validate.Struct(r)
	if err != nil {
		return helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.GetUserByUsername(username)
	if existingUser == nil {
		return fmt.Errorf("User not found")
	}

	err = helper.ComparePassword(existingUser.Password, r.OldPassword)

	if err != nil {
		return fmt.Errorf("Invalid old password")
	}

	existingUser.Password = helper.HashPassword(r.NewPassword)
	err = s.UserRepository.ChangePassword(existingUser, username)
	if err != nil {
		return fmt.Errorf("Error when change password user: %s", err.Error())
	}

	return nil
}
func (s *UserServiceImpl) UserLogin(ctx echo.Context, r dto.UserLoginRequest) (*dto.UserResponse, error) {
	err := s.Validate.Struct(r)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	fmt.Println(r, r.Username)
	existingUser, _ := s.UserRepository.GetUserByUsername(r.Username)
	if existingUser == nil {
		return nil, fmt.Errorf("Invalid username or password")
	}

	err = helper.ComparePassword(existingUser.Password, r.Password)

	if err != nil {
		return nil, fmt.Errorf("Invalid username or password")
	}

	user := helper.ConvertToUserResponse(existingUser)

	return user, nil
}
