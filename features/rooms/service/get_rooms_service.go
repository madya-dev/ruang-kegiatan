package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *RoomServiceImpl) GetAllRooms(ctx echo.Context) ([]model.Room, int, error) {
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

	rooms, total, err := s.RoomRespository.GetAllRooms(offset, limit, search)
	if err != nil {
		return nil, total, fmt.Errorf("Internal Server Error")
	}
	if len(search) > 0 && total <= 0 {
		return nil, total, fmt.Errorf("room not found")
	}
	return rooms, total, nil
}
