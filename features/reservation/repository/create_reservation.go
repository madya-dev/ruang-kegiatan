package repository

import "madyasantosa/ruangkegiatan/model"

func (r *ReservationRepositoryImpl) CreateReservation(reservation *model.Reservation) error {
	result := r.DB.Create(&reservation)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
