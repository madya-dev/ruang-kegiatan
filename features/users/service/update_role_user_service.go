package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

func (s *UserServiceImpl) UpdateRoleUser(ctx echo.Context, r dto.UserRoleUpdateRequest) error {
	err := s.Validate.Struct(r)
	if err != nil {
		return helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.GetUserByUsername(ctx.Param("username"))
	if existingUser == nil {
		return fmt.Errorf("User not found")
	}

	user := helper.UserRoleUpdateRequestToUserModel(r)

	err = s.UserRepository.UpdateRoleUser(user, ctx.Param("username"))
	if err != nil {
		return fmt.Errorf("Error when updating user: %s", err.Error())
	}

	return nil
}
