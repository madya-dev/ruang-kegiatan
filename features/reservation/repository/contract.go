package repository

import (
	"madyasantosa/ruangkegiatan/dto"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	GetAllReservation(offset int, limit int, search string) ([]dto.Reservation, int, error)
	// CreateReservation(Reservation *model.Reservation) error
	// UpdateReservation(Reservation *model.Reservation, id int) error
	// FindReservationById(id int) (*model.Reservation, error)
	// FindReservationByName(name string) (*model.Reservation, error)
	// DeleteReservation(id int) error
}

type ReservationRepositoryImpl struct {
	DB *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &ReservationRepositoryImpl{
		DB: db,
	}
}
