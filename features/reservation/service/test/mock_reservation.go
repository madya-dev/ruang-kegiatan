package test

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
	"time"
)

type MockReservationRepository struct {
	GetAllReservationFunc func(offset int, limit int, search string) ([]dto.ReservationResponse, int, error)
	CheckReservationFunc  func(id int) (*dto.ReservationCheck, error)
	CreateReservationFunc func(reservation *model.Reservation) error
	IsAvaibleFunc         func(roomID int64, startTime time.Time, endTime time.Time) (*model.Reservation, string, error)
	UpdateReservationFunc func(reservation *model.Reservation, id int) error
	DeleteReservationFunc func(id int) error
}

func (m *MockReservationRepository) GetAllReservation(offset int, limit int, search string) ([]dto.ReservationResponse, int, error) {
	return m.GetAllReservationFunc(offset, limit, search)
}

func (m *MockReservationRepository) CheckReservation(id int) (*dto.ReservationCheck, error) {
	return m.CheckReservationFunc(id)
}

func (m *MockReservationRepository) CreateReservation(reservation *model.Reservation) error {
	return m.CreateReservationFunc(reservation)
}

func (m *MockReservationRepository) IsAvaible(roomID int64, startTime time.Time, endTime time.Time) (*model.Reservation, string, error) {
	return m.IsAvaibleFunc(roomID, startTime, endTime)
}

func (m *MockReservationRepository) UpdateReservation(reservation *model.Reservation, id int) error {
	return m.UpdateReservationFunc(reservation, id)
}

func (m *MockReservationRepository) DeleteReservation(id int) error {
	return m.DeleteReservationFunc(id)
}
