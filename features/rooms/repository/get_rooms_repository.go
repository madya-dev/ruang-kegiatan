package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *RoomRespositoryImpl) GetAllRooms(offset int, limit int, search string) ([]model.Room, int, error) {
	rooms := []model.Room{}
	var total int64
	r.DB.Where("room_name LIKE ?", "%"+search+"%").Find(&rooms).Count(&total)

	result := r.DB.Where("room_name LIKE ?", "%"+search+"%").Offset(offset).Limit(limit).Find(&rooms)
	fmt.Println(result)
	if result.Error != nil {
		return nil, int(total), result.Error
	}
	return rooms, int(total), nil
}
