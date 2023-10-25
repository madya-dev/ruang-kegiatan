package handler

import (
	"madyasantosa/ruangkegiatan/features/rooms/service"

	"github.com/labstack/echo/v4"
)

type RoomHandler interface {
	GetAllRooms(ctx echo.Context) error
	DeleteRoom(ctx echo.Context) error
	CreateRoom(ctx echo.Context) error
	UpdateRoom(ctx echo.Context) error
}

type RoomHandlerImpl struct {
	RoomService service.RoomService
}

func NewRoomHandler(rs service.RoomService) RoomHandler {
	return &RoomHandlerImpl{RoomService: rs}
}
