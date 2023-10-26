package repository

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
)

func (r *ReservationRepositoryImpl) GetAllReservation(offset int, limit int, search string) ([]dto.Reservation, int, error) {
	reservation := []dto.Reservation{}
	var total int64

	result := r.DB.Model(&model.Reservation{}).
		Select("reservations.id as id, room_name, activity, start_time, end_time, reservations.study_program as study_program, class_of, document, pic, phone").
		Joins("left join users on users.username = reservations.pic").
		Joins("left join rooms on rooms.id = reservations.room_id").
		Where("room_name LIKE ?", "%"+search+"%").
		Find(&reservation)

	result.Count(&total)

	if result.Error != nil {
		return nil, int(total), result.Error
	}
	return reservation, int(total), nil
}
