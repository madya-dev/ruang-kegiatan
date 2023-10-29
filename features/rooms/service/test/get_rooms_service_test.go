package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/features/rooms/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoomServiceImpl_GetAllRooms(t *testing.T) {
	roomService := &service.RoomServiceImpl{
		RoomRespository: &MockRoomRepository{
			GetAllRoomsFunc: func(offset int, limit int, search string) ([]model.Room, int, error) {
				if search == "existingroom" {
					return []model.Room{
						{ID: 1, RoomName: "existingroom", Capacity: 50},
					}, 1, nil
				} else if search == "emptysearch" {
					return []model.Room{}, 0, nil
				} else if search == "invalidlimit" {
					return nil, 0, fmt.Errorf("Params limit not valid")
				} else if search == "invalidoffset" {
					return nil, 0, fmt.Errorf("Params offset not valid")
				} else if search == "internalerror" {
					return nil, 0, fmt.Errorf("Internal Server Error")
				}
				return nil, 0, nil
			},
		},
	}

	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/your-api-endpoint?s=existingroom&limit=10&offset=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	rooms, total, err := roomService.GetAllRooms(c)

	assert.NoError(t, err)
	assert.NotNil(t, rooms)
	assert.Equal(t, 1, total)
	assert.Len(t, rooms, 1)
	assert.Equal(t, "existingroom", rooms[0].RoomName)
}
