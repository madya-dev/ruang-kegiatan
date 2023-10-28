package handler

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/notification/service"

	"github.com/labstack/echo/v4"
)

type NotificationHandler interface {
	CreateNotification(ctx echo.Context, r dto.NotificationRequest) error
}

type NotificationHandlerImpl struct {
	NotificationService service.NotificationService
}

func NewNotificationHandler(us service.NotificationService) NotificationHandler {
	return &NotificationHandlerImpl{NotificationService: us}
}
