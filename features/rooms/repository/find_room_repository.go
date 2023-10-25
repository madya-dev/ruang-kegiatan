package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *RoomRespositoryImpl) FindRoomById(id int) (*model.Room, error) {
	room := model.Room{}
	result := r.DB.Where("id = ?", id).First(&room)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &room, nil
}

func (r *RoomRespositoryImpl) FindRoomByName(name string) (*model.Room, error) {
	room := model.Room{}
	result := r.DB.Where("room_name = ?", name).First(&room)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &room, nil
}
