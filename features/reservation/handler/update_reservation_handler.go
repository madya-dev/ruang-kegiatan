package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (rh *ReservationHanlderImpl) UpdateReservation(ctx echo.Context) error {
	authorization := ctx.Request().Header["Authorization"]
	userToken := strings.Split(authorization[0], " ")[1]
	data, _ := helper.ExtractToken(userToken)

	reservationUpdateRequest := dto.ReservationRequest{}

	err := ctx.Bind(&reservationUpdateRequest)

	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	err = rh.ReservationService.UpdateReservation(ctx, reservationUpdateRequest, data.Username)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}
		if strings.Contains(err.Error(), "room already book") {
			return helper.StatusBadRequest(ctx, err)
		}
		if strings.Contains(err.Error(), "notification") {
			return helper.StatusBadRequest(ctx, err)
		}
		if strings.Contains(err.Error(), "Reservation not found") {
			return helper.StatusNotFound(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to update reservation"))
	}

	return helper.StatusNoContent(ctx, "Success to update reservation")
}
