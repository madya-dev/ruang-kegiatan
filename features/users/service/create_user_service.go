package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

func (s *UserServiceImpl) CreateUser(ctx echo.Context, r dto.UserCreateRequest) (*dto.UserResponse, error) {
	err := s.Validate.Struct(r)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.GetUserByUsername(r.Username)
	if existingUser != nil {
		return nil, fmt.Errorf("User already exists")
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
