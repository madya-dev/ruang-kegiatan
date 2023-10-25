package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

func (s *UserServiceImpl) UserLogin(ctx echo.Context, r dto.UserLoginRequest) (*dto.UserResponse, error) {
	err := s.Validate.Struct(r)

	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

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
