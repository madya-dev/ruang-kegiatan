package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

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
