package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (rh *RoomHandlerImpl) UpdateRoom(ctx echo.Context) error {
	roomUpdateRequest := dto.RoomRequest{}

	err := ctx.Bind(&roomUpdateRequest)

	fmt.Println(roomUpdateRequest)

	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	err = rh.RoomService.UpdateRoom(ctx, roomUpdateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "Room not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("Room not found!"))
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed update room"))
	}
	return helper.StatusNoContent(ctx, "Success to update user")

}
