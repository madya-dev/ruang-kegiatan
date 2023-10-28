package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *NotificationServiceImpl) GetAllNotifications(ctx echo.Context) ([]model.Notification, int, error) {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return nil, 0, fmt.Errorf("Params limit not valid")
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return nil, 0, fmt.Errorf("Params offset not valid")
	}

	notifications, total, err := s.NotificationRepository.GetAllNotifications(offset, limit)
	if err != nil {
		return nil, total, fmt.Errorf("Internal Server Error")
	}
	return notifications, total, nil
}
