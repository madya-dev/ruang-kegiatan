package service

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *RoomServiceImpl) DeleteRoom(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("id not valid: %v", err)
	}

	if err := s.RoomRespository.DeleteRoom(id); err != nil {
		return err
	}

	return nil
}
