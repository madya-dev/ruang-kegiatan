package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (rh *RoomHandlerImpl) CreateRoom(ctx echo.Context) error {
	roomCreateRequest := dto.RoomRequest{}

	err := ctx.Bind(&roomCreateRequest)

	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	err = rh.RoomService.CreateRoom(ctx, roomCreateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to create room"))
	}

	return helper.StatusCreatedNoContent(ctx, "Success to created room")
}
