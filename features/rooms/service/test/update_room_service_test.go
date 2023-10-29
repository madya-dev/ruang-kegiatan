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

func TestRoomServiceImpl_UpdateRoom(t *testing.T) {
	roomService := &service.RoomServiceImpl{
		RoomRespository: &MockRoomRepository{
			FindRoomByIdFunc: func(id int) (*model.Room, error) {
				if id == 1 {
					return &model.Room{
						ID:       1,
						RoomName: "testroom",
						Capacity: 100,
					}, nil
				}
				return nil, nil
			},
			UpdateRoomFunc: func(room *model.Room, id int) error {
				return nil
			},
		},
		Validate: validator.New(),
	}

	e := echo.New()

	req := httptest.NewRequest(echo.PUT, "/your-api-endpoint/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	updateRoomRequest := dto.RoomRequest{
		RoomName: "updatedroom",
		Capacity: 150,
	}

	err := roomService.UpdateRoom(c, updateRoomRequest)

	assert.NoError(t, err)
}
