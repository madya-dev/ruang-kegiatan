package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (rh *ReservationHanlderImpl) DeleteReservation(ctx echo.Context) error {
	authorization := ctx.Request().Header["Authorization"]
	userToken := strings.Split(authorization[0], " ")[1]
	data, _ := helper.ExtractToken(userToken)

	err := rh.ReservationService.DeleteReservation(ctx, data.Username)

	if err != nil {
		if strings.Contains(err.Error(), "Reservation not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("Reservation not found!"))
		}
		if strings.Contains(err.Error(), "Forbidden!") {
			return helper.StatusForbidden(ctx, fmt.Errorf("Forbidden!"))
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to delete reservation"))
	}

	return helper.StatusNoContent(ctx, "Success to delete data")
}
