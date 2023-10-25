package repository

import "madyasantosa/ruangkegiatan/model"

func (r *RoomRespositoryImpl) CreateRoom(room *model.Room) (*model.Room, error) {
	result := r.DB.Create(&room)

	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}
