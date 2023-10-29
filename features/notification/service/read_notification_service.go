package service

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *NotificationServiceImpl) ReadNotification(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return fmt.Errorf("id not valid")
	}

	err = s.NotificationRepository.ReadNotification(id)

	if err != nil {
		return err
	}

	return nil
}
