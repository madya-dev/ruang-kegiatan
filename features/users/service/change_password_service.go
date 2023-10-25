package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

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
