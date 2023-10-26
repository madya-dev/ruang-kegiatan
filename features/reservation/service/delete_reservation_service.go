package service

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *ReservationServiceImpl) DeleteReservation(ctx echo.Context, username string) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return fmt.Errorf("id not valid")
	}

	res, err := s.ReservationRepository.CheckReservation(id)

	if err != nil {
		return fmt.Errorf("Reservation not found")
	}

	if res.PIC != username {
		return fmt.Errorf("Forbidden!")
	}

	err = s.ReservationRepository.DeleteReservation(id)

	if err != nil {
		return err
	}

	return nil
}
