package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"

	"github.com/labstack/echo/v4"
)

func (s *NotificationServiceImpl) CreateNotification(ctx echo.Context, r dto.NotificationRequest) error {
	err := s.Validate.Struct(r)
	if err != nil {
		return helper.ValidationError(ctx, err)
	}

	notif := helper.NotificationCreateRequestToNotificationModel(r)

	err = s.NotificationRepository.CreateNotification(notif)

	if err != nil {
		return fmt.Errorf("Error when creating notification %s:", err.Error())
	}

	return nil
}
