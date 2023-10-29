package repository

import (
	"madyasantosa/ruangkegiatan/model"
)

func (r *ReservationRepositoryImpl) UpdateReservation(reservation *model.Reservation, id int) error {
	reservations := &model.Reservation{}
	result := r.DB.Model(reservations).Where("id = ?", id).
		Updates(model.Reservation{
			RoomID:       reservation.RoomID,
			Activity:     reservation.Activity,
			StartTime:    reservation.StartTime,
			EndTime:      reservation.EndTime,
			StudyProgram: reservation.StudyProgram,
			ClassOf:      reservation.ClassOf,
			Document:     reservation.Document,
		})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
