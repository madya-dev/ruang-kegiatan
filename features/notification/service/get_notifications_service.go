package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *NotificationServiceImpl) GetNotifications(ctx echo.Context) ([]model.Notification, error) {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))
	username := ctx.Param("username")

	if err != nil {
		return nil, fmt.Errorf("Username not valid")
	}

	if err != nil {
		return nil, fmt.Errorf("Params limit not valid")
	}

	if err != nil {
		return nil, fmt.Errorf("Params limit not valid")
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return nil, fmt.Errorf("Params offset not valid")
	}

	notifications, err := s.NotificationRepository.GetNotifications(offset, limit, username)
	if err != nil {
		return nil, fmt.Errorf("Internal Server Error")
	}
	return notifications, nil
}
