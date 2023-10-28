package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (un *NotificationHandlerImpl) CreateNotification(ctx echo.Context, r dto.NotificationRequest) error {

	err := un.NotificationService.CreateNotification(ctx, r)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to create notification"))
	}

	return helper.StatusCreatedNoContent(ctx, "Success to created notification")
}
