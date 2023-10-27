package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (rh *ReservationHanlderImpl) CreateReservation(ctx echo.Context) error {
	authorization := ctx.Request().Header["Authorization"]
	userToken := strings.Split(authorization[0], " ")[1]
	data, _ := helper.ExtractToken(userToken)

	reservationCreateRequest := dto.ReservationRequest{}

	err := ctx.Bind(&reservationCreateRequest)

	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	err = rh.ReservationService.CreateReservation(ctx, reservationCreateRequest, data.Username)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}
		if strings.Contains(err.Error(), "room already book") {
			return helper.StatusBadRequest(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to create reservation"))
	}

	return helper.StatusCreatedNoContent(ctx, "Success to created reservation")
}
