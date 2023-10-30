package repository

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
)

func (r *ReservationRepositoryImpl) CheckReservation(id int) (*dto.ReservationCheck, error) {
	reservation := &dto.ReservationCheck{}

	result := r.DB.Model(&model.Reservation{}).
		Select("id, pic, start_time, document").Where("id = ?", id).
		First(&reservation)

	if result.Error != nil {
		return nil, result.Error
	}
	return reservation, nil
}
