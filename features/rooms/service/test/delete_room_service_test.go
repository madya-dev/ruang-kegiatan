package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/features/rooms/service"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoomServiceImpl_DeleteRoom(t *testing.T) {
	roomService := &service.RoomServiceImpl{
		RoomRespository: &MockRoomRepository{
			DeleteRoomFunc: func(id int) error {
				return nil
			},
		},
	}

	e := echo.New()

	roomID := 42
	req := httptest.NewRequest(echo.DELETE, fmt.Sprintf("/your-api-endpoint/%d", roomID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(roomID))

	err := roomService.DeleteRoom(c)

	assert.NoError(t, err)
}

func TestRoomServiceImpl_DeleteRoomInvalidID(t *testing.T) {
	roomService := &service.RoomServiceImpl{}

	e := echo.New()

	invalidID := "notanumber"
	req := httptest.NewRequest(echo.DELETE, fmt.Sprintf("/your-api-endpoint/%s", invalidID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(invalidID)

	err := roomService.DeleteRoom(c)

	assert.Error(t, err)
}
