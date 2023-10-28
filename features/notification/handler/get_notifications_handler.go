package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

func (nh *NotificationHandlerImpl) GetNotifications(ctx echo.Context) error {
	res, err := nh.NotificationService.GetNotifications(ctx)

	if err != nil {
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to get notification %s", ctx.Param("username")))
	}

	return helper.StatusOK(ctx, "Success to get data", res)
}
