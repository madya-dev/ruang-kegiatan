package service

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/reservation/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ReservationService interface {
	GetAllReservation(ctx echo.Context) ([]dto.Reservation, int, error)
	DeleteReservation(ctx echo.Context, username string) error
	// CreateReservation(ctx echo.Context, r dto.RoomRequest) error
	// UpdateReservation(ctx echo.Context, r dto.RoomRequest) error
}

type ReservationServiceImpl struct {
	ReservationRepository repository.ReservationRepository
	Validate              *validator.Validate
}

func NewReservationService(rr repository.ReservationRepository, validate *validator.Validate) *ReservationServiceImpl {
	return &ReservationServiceImpl{
		ReservationRepository: rr,
		Validate:              validate,
	}
}
