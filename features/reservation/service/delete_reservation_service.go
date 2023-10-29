package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/pkg/s3"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *ReservationServiceImpl) DeleteReservation(ctx echo.Context, username string) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return fmt.Errorf("id not valid")
	}

	res, err := s.ReservationRepository.CheckReservation(id)

	if res == nil {
		return fmt.Errorf("Reservation not found")
	}

	if res.PIC != username {
		return fmt.Errorf("Forbidden!")
	}

	err = s.ReservationRepository.DeleteReservation(id)

	if err != nil {
		return err
	}

	if res.Document != "" {
		err = s3.DeleteFileS3(res.PIC, strconv.FormatInt(time.Time(res.StartTime).Unix(), 10))
	}

	return nil
}
