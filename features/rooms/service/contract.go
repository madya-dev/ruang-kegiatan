package service

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/rooms/repository"
	"madyasantosa/ruangkegiatan/model"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RoomService interface {
	GetAllRooms(ctx echo.Context) ([]model.Room, int, error)
	DeleteRoom(ctx echo.Context) error
	CreateRoom(ctx echo.Context, r dto.RoomRequest) error
	UpdateRoom(ctx echo.Context, r dto.RoomRequest) error
}

type RoomServiceImpl struct {
	RoomRespository repository.RoomRespository
	Validate        *validator.Validate
}

func NewRoomService(rr repository.RoomRespository, validate *validator.Validate) *RoomServiceImpl {
	return &RoomServiceImpl{
		RoomRespository: rr,
		Validate:        validate,
	}
}
