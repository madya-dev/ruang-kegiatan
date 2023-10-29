package handler

import (
	"madyasantosa/ruangkegiatan/features/reservation/service"

	"github.com/labstack/echo/v4"
)

type ReservationHanlder interface {
	GetAllReservation(ctx echo.Context) error
	DeleteReservation(ctx echo.Context) error
	CreateReservation(ctx echo.Context) error
	UpdateReservation(ctx echo.Context) error
}

type ReservationHanlderImpl struct {
	ReservationService service.ReservationService
}

func NewReservationHanlder(rs service.ReservationService) ReservationHanlder {
	return &ReservationHanlderImpl{ReservationService: rs}
}
