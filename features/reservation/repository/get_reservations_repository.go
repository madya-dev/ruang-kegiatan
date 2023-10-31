package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
	"strings"
)

func (r *ReservationRepositoryImpl) GetAllReservation(offset int, limit int, search string, start_time string, end_time string) ([]dto.ReservationResponse, int, error) {
	reservation := []dto.ReservationResponse{}
	var total int64
	var conditions []string
	var values []interface{}

	conditions = append(conditions, "room_name LIKE ?")
	values = append(values, "%"+search+"%")

	fmt.Println(start_time)

	if start_time != "" {
		conditions = append(conditions, "start_time >= ? AND end_time <= ?")
		values = append(values, start_time, end_time)
	}

	condition := strings.Join(conditions, " AND ")

	result := r.DB.Model(&model.Reservation{}).
		Select("reservations.id as id, room_name, activity, start_time, end_time, reservations.study_program as study_program, class_of, document, pic, phone, reservations.updated_at").
		Joins("left join users on users.username = reservations.pic").
		Joins("left join rooms on rooms.id = reservations.room_id").
		Where(condition, values...).
		Offset(offset).
		Limit(limit).
		Find(&reservation)
	result.Count(&total)

	if result.Error != nil {
		return nil, int(total), result.Error
	}

	fmt.Println(reservation)
	fmt.Println(condition, values)

	return reservation, int(total), nil

}
