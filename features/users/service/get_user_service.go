package service

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

func (s *UserServiceImpl) GetUserByUsername(ctx echo.Context) (*dto.UserResponse, error) {
	user, err := s.UserRepository.GetUserByUsername(ctx.Param("username"))
	if err != nil {
		return nil, err
	}
	userResponse := helper.ConvertToUserResponse(user)
	return userResponse, nil
}
