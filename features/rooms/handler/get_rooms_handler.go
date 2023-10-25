package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (rh *RoomHandlerImpl) GetAllRooms(ctx echo.Context) error {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		fmt.Println(err)
		return helper.StatusBadRequest(ctx, fmt.Errorf("Params limit not valid"))
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		fmt.Println(err)
		return helper.StatusBadRequest(ctx, fmt.Errorf("Params offset not valid"))
	}
	res, total, err := rh.RoomService.GetAllRooms(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Rooms not found") {
			return helper.StatusNotFound(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusOKWithPagination(ctx, "Success to Get Data", res, offset, limit, total)
}
