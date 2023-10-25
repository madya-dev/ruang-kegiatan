package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (rh *RoomHandlerImpl) DeleteRoom(ctx echo.Context) error {
	err := rh.RoomService.DeleteRoom(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Room not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("Room not found!"))
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to delete room"))
	}

	return helper.StatusNoContent(ctx, "Success to delete data")
}
