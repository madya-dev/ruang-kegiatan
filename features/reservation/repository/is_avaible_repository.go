package repository

import (
	"madyasantosa/ruangkegiatan/model"
	"time"
)

func (r *ReservationRepositoryImpl) IsAvaible(roomID int64, startTime time.Time, endTime time.Time) (*model.Reservation, string, error) {
	reservation := model.Reservation{}
	result := r.DB.Where("room_id = ? AND ((? >= start_time AND ? < end_time) OR (? < start_time AND ? <= end_time AND ? > start_time) OR (? < start_time AND ? >= end_time AND ? > start_time))", roomID, startTime, startTime, startTime, endTime, endTime, startTime, endTime, endTime).First(&reservation)
	if result.Error != nil {
		return nil, "", result.Error
	}
	var registrationToken string

	r.DB.Model(&model.User{}).Select("registration_token").Where("username = ?", reservation.PIC).First(&registrationToken)

	return &reservation, registrationToken, nil
}
