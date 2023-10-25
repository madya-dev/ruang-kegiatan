package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *RoomServiceImpl) UpdateRoom(ctx echo.Context, r dto.RoomRequest) error {
	err := s.Validate.Struct(r)
	if err != nil {
		return helper.ValidationError(ctx, err)
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		return fmt.Errorf("id not valid")
	}

	existingRoom, _ := s.RoomRespository.FindRoomById(id)

	if existingRoom == nil {
		return fmt.Errorf("Room not found")
	}

	room := helper.RoomRequestToRoomModel(r)

	err = s.RoomRespository.UpdateRoom(room, id)

	if err != nil {
		return fmt.Errorf("Error when updating user: %s", err.Error())
	}

	return nil
}
