package repository

import (
	"madyasantosa/ruangkegiatan/model"

	"gorm.io/gorm"
)

type RoomRespository interface {
	GetAllRooms(offset int, limit int, search string) ([]model.Room, int, error)
	CreateRoom(room *model.Room) error
	UpdateRoom(room *model.Room, id int) error
	FindRoomById(id int) (*model.Room, error)
	FindRoomByName(name string) (*model.Room, error)
	DeleteRoom(id int) error
}

type RoomRespositoryImpl struct {
	DB *gorm.DB
}

func NewRoomRespository(db *gorm.DB) RoomRespository {
	return &RoomRespositoryImpl{
		DB: db,
	}
}
