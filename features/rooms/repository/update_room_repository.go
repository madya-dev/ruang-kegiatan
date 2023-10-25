package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *RoomRespositoryImpl) UpdateRoom(room *model.Room, id int) (*model.Room, error) {
	rooms := model.Room{}
	result := r.DB.Model(&rooms).Where("id = ?", id).Updates(model.Room{RoomName: room.RoomName, Capacity: room.Capacity})
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}

	return room, nil
}
