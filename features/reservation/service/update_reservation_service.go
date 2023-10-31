package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"madyasantosa/ruangkegiatan/pkg/firebase"
	"madyasantosa/ruangkegiatan/pkg/s3"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *ReservationServiceImpl) UpdateReservation(ctx echo.Context, r dto.ReservationRequest, username string) error {
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

	data, registrationToken, _ := s.ReservationRepository.IsAvaible(r.RoomID, r.StartTime, r.EndTime)

	if data != nil && data.ID != int64(id) {
		return fmt.Errorf("Reservation failed, room already book to %s by %s", data.Activity, data.PIC)
	}

	fileHeader, err := ctx.FormFile("document")

	if fileHeader != nil {
		if res.StartTime != r.StartTime {
			s3.DeleteFileS3(res.PIC, strconv.FormatInt(time.Time(r.StartTime).Unix(), 10))
		}
	}

	docsUrl, err := helper.UploadToS3(fileHeader, res.PIC, strconv.FormatInt(time.Time(r.StartTime).Unix(), 10))

	if err != nil {
		return fmt.Errorf("Error when upload file %s:", err.Error())
	}

	reservation := helper.ReservationRequestToReservationModel(r, docsUrl, res.PIC)

	err = s.ReservationRepository.UpdateReservation(reservation, id)

	if err != nil {
		return fmt.Errorf("Error when updating reservation %s:", err.Error())
	}

	notification := dto.NotificationRequest{
		Title:    "Reservation Updated!",
		Message:  "Your reservation for " + data.Activity + " updated. Please check your reservation!",
		Username: username,
	}

	if registrationToken != "" {
		firebase.SendNotification(registrationToken, notification.Title, notification.Message)
	}

	s.ReservationRepository.InsertTrackReservation(id)

	notif := helper.NotificationCreateRequestToNotificationModel(notification)
	_ = s.NotificationRepository.CreateNotification(notif)

	return nil
}
