package test

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/rooms/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoomServiceImpl_CreateRoom(t *testing.T) {
	roomService := &service.RoomServiceImpl{
		RoomRespository: &MockRoomRepository{
			FindRoomByNameFunc: func(name string) (*model.Room, error) {
				if name == "existingroom" {
					return &model.Room{
						RoomName: "existingroom",
					}, nil
				}
				return nil, nil
			},
			CreateRoomFunc: func(room *model.Room) error {
				return nil
			},
		},
		Validate: validator.New(),
	}

	e := echo.New()

	req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	createRoomRequest := dto.RoomRequest{
		RoomName: "newroom",
		Capacity: 50,
	}

	err := roomService.CreateRoom(c, createRoomRequest)

	assert.NoError(t, err)
}

func TestRoomServiceImpl_CreateRoomExistingName(t *testing.T) {
	roomService := &service.RoomServiceImpl{
		RoomRespository: &MockRoomRepository{
			FindRoomByNameFunc: func(name string) (*model.Room, error) {
				if name == "existingroom" {
					return &model.Room{
						RoomName: "existingroom",
					}, nil
				}
				return nil, nil
			},
		},
		Validate: validator.New(),
	}

	e := echo.New()

	req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	createRoomRequest := dto.RoomRequest{
		RoomName: "existingroom",
		Capacity: 50,
	}

	err := roomService.CreateRoom(c, createRoomRequest)

	assert.Error(t, err)
}
