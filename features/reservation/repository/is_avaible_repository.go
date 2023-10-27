package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
	"time"
)

func (r *ReservationRepositoryImpl) IsAvaible(roomID int64, startTime time.Time, endTime time.Time) (*model.Reservation, error) {
	reservation := model.Reservation{}
	result := r.DB.Where("room_id = ? AND ((? >= start_time AND ? < end_time) OR (? < start_time AND ? <= end_time AND ? > start_time) OR (? < start_time AND ? >= end_time AND ? > start_time))", roomID, startTime, startTime, startTime, endTime, endTime, startTime, endTime, endTime).First(&reservation)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reservation, nil
}
