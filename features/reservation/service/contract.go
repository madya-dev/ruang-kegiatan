package service

import (
	"madyasantosa/ruangkegiatan/dto"
	notifRepo "madyasantosa/ruangkegiatan/features/notification/repository"
	"madyasantosa/ruangkegiatan/features/reservation/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ReservationService interface {
	GetAllReservation(ctx echo.Context) ([]dto.ReservationResponse, int, error)
	DeleteReservation(ctx echo.Context, username string) error
	CreateReservation(ctx echo.Context, r dto.ReservationRequest, pic string) error
	UpdateReservation(ctx echo.Context, r dto.ReservationRequest, pic string) error
}

type ReservationServiceImpl struct {
	ReservationRepository  repository.ReservationRepository
	Validate               *validator.Validate
	NotificationRepository notifRepo.NotificationRepository
}

func NewReservationService(rr repository.ReservationRepository, validate *validator.Validate, nr notifRepo.NotificationRepository) *ReservationServiceImpl {
	return &ReservationServiceImpl{
		ReservationRepository:  rr,
		Validate:               validate,
		NotificationRepository: nr,
	}
}
