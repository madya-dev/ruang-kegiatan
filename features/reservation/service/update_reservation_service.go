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

	data, _ := s.ReservationRepository.IsAvaible(r.RoomID, r.StartTime, r.EndTime)

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

	message := "Your reservation for " + data.Activity + " updated. Please check your reservation!"

	firebase.SendNotification("cVtcoEkNhF038xc5rKD0u4:APA91bFiEqmgD7XE0uksikpl4eC8rqca4MkH1cE87T1qsdZSSNPtCvw9UZsucVD6EggpidNQ_OfzkvuaYEZTJ63yPvdtt_cP0D-IQObziFzGVaSeCGXs2gGBhg2cICWnBbKW2Ay8ONs-", "Reservation Updated!", message)

	return nil
}
