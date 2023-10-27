package repository

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
	"time"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	GetAllReservation(offset int, limit int, search string) ([]dto.ReservationResponse, int, error)
	CheckReservation(id int) (*dto.ReservationCheck, error)
	CreateReservation(reservation *model.Reservation) error
	IsAvaible(roomID int64, startTime time.Time, endTime time.Time) (*model.Reservation, error)
	// UpdateReservation(Reservation *model.Reservation, id int) error
	// FindReservationById(id int) (*model.Reservation, error)
	// FindReservationByName(name string) (*model.Reservation, error)
	DeleteReservation(id int) error
}

type ReservationRepositoryImpl struct {
	DB *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &ReservationRepositoryImpl{
		DB: db,
	}
}
