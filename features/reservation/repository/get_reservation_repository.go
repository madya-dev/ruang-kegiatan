package repository

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
)

func (r *ReservationRepositoryImpl) CheckReservation(id int) (*dto.ReservationCheck, error) {
	reservation := &dto.ReservationCheck{}

	result := r.DB.Model(&model.Reservation{}).
		Select("id, pic").Where("id = ?", id).
		Find(&reservation)

	if result.Error != nil {
		return nil, result.Error
	}
	return reservation, nil
}
