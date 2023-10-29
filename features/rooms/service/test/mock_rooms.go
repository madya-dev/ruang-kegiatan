package test

import (
	"madyasantosa/ruangkegiatan/model"
)

type MockRoomRepository struct {
	GetAllRoomsFunc    func(offset int, limit int, search string) ([]model.Room, int, error)
	CreateRoomFunc     func(room *model.Room) error
	UpdateRoomFunc     func(room *model.Room, id int) error
	FindRoomByIdFunc   func(id int) (*model.Room, error)
	FindRoomByNameFunc func(name string) (*model.Room, error)
	DeleteRoomFunc     func(id int) error
}

func (m *MockRoomRepository) GetAllRooms(offset int, limit int, search string) ([]model.Room, int, error) {
	return m.GetAllRoomsFunc(offset, limit, search)
}

func (m *MockRoomRepository) CreateRoom(room *model.Room) error {
	return m.CreateRoomFunc(room)
}

func (m *MockRoomRepository) UpdateRoom(room *model.Room, id int) error {
	return m.UpdateRoomFunc(room, id)
}

func (m *MockRoomRepository) FindRoomById(id int) (*model.Room, error) {
	return m.FindRoomByIdFunc(id)
}

func (m *MockRoomRepository) FindRoomByName(name string) (*model.Room, error) {
	return m.FindRoomByNameFunc(name)
}

func (m *MockRoomRepository) DeleteRoom(id int) error {
	return m.DeleteRoomFunc(id)
}
