package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
	if len(search) > 0 && total <= 0 {
		return nil, total, fmt.Errorf("Users not found")
	}
	if err != nil {
		return nil, total, fmt.Errorf("Internal Server Error")
	}
	return users, total, nil
}
