package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (nh *NotificationHandlerImpl) ReadNotification(ctx echo.Context) error {
	_, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return helper.StatusBadRequest(ctx, fmt.Errorf("Notification ID not valid!"))
	}

	err = nh.NotificationService.ReadNotification(ctx)

	if err != nil {
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to update notification"))
	}

	return helper.StatusNoContent(ctx, "Success to get data")

}
