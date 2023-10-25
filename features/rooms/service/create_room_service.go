package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

func (s *RoomServiceImpl) CreateRoom(ctx echo.Context, r dto.RoomRequest) error {
	err := s.Validate.Struct(r)
	if err != nil {
		return helper.ValidationError(ctx, err)
	}

	existingRoom, _ := s.RoomRespository.FindRoomByName(r.RoomName)
	if existingRoom != nil {
		return fmt.Errorf("Room name already exists")
	}

	room := helper.RoomRequestToRoomModel(r)

	err = s.RoomRespository.CreateRoom(room)

	if err != nil {
		return fmt.Errorf("Error when creating room %s:", err.Error())
	}
	return nil
}
