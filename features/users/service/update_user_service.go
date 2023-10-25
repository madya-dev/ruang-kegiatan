package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

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
