package repository

import (
	"madyasantosa/ruangkegiatan/model"
)

func (r *ReservationRepositoryImpl) CreateReservation(reservation *model.Reservation) (int, error) {
	result := r.DB.Create(&reservation)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(reservation.ID), nil
}
