package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *ReservationServiceImpl) CreateReservation(ctx echo.Context, r dto.ReservationRequest, pic string) error {
	err := s.Validate.Struct(r)
	if err != nil {
		return helper.ValidationError(ctx, err)
	}

	res, _, _ := s.ReservationRepository.IsAvaible(r.RoomID, r.StartTime, r.EndTime)

	if res != nil {
		return fmt.Errorf("Reservation failed, room already booked to %s by %s", res.Activity, res.PIC)
	}

	fileHeader, err := ctx.FormFile("document")

	docsUrl, err := helper.UploadToS3(fileHeader, pic, strconv.FormatInt(r.StartTime.Unix(), 10))

	if err != nil {
		return fmt.Errorf("Error when uploading file: %s", err.Error())
	}

	reservation := helper.ReservationRequestToReservationModel(r, docsUrl, pic)

	reservationID, err := s.ReservationRepository.CreateReservation(reservation)

	if err != nil {
		return fmt.Errorf("Error when creating reservation: %s", err.Error())
	}

	s.ReservationRepository.InsertTrackReservation(reservationID)

	return nil
}
