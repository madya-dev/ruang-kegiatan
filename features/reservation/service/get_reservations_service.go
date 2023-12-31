package service

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *ReservationServiceImpl) GetAllReservation(ctx echo.Context) ([]dto.ReservationResponse, int, error) {
	params := ctx.QueryParams()
	search := params.Get("s")
	start_time := params.Get("start_time")
	end_time := params.Get("end_time")

	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return nil, 0, fmt.Errorf("Params limit not valid")
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return nil, 0, fmt.Errorf("Params offset not valid")
	}

	reservations, total, err := s.ReservationRepository.GetAllReservation(offset, limit, search, start_time, end_time)
	if err != nil {
		return nil, total, fmt.Errorf("Internal Server Error")
	}
	if len(search) > 0 && total <= 0 {
		return nil, total, fmt.Errorf("Reservation not found")
	}
	return reservations, total, nil
}
