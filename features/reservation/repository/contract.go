package repository

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	GetAllReservation(offset int, limit int, search string, start_date string, end_date string) ([]dto.ReservationResponse, int, error)
	CheckReservation(id int) (*dto.ReservationCheck, error)
	CreateReservation(reservation *model.Reservation) (int, error)
	IsAvaible(roomID int64, startTime time.Time, endTime time.Time) (*model.Reservation, string, error)
	UpdateReservation(reservation *model.Reservation, id int) error
	DeleteReservation(id int) error
	InsertTrackReservation(id int)
}

type ReservationRepositoryImpl struct {
	DB      *gorm.DB
	MongoDB *mongo.Client
}

func NewReservationRepository(db *gorm.DB, mongo *mongo.Client) ReservationRepository {
	return &ReservationRepositoryImpl{
		DB:      db,
		MongoDB: mongo,
	}
}
