package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/reservation/service"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllReservation(t *testing.T) {
	reservationRepository := &MockReservationRepository{}
	reservationService := &service.ReservationServiceImpl{
		ReservationRepository: reservationRepository,
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/your-api-endpoint", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("ValidQueryParametersWithData", func(t *testing.T) {
		reservationRepository.GetAllReservationFunc = func(offset, limit int, search string) ([]dto.ReservationResponse, int, error) {
			return []dto.ReservationResponse{
				{ID: 1, RoomName: "Reservation 1"},
				{ID: 2, RoomName: "Reservation 2"},
			}, 2, nil
		}

		c.QueryParams().Set("s", "search")
		c.QueryParams().Set("limit", "10")
		c.QueryParams().Set("offset", "0")

		reservations, total, err := reservationService.GetAllReservation(c)

		assert.NoError(t, err)
		assert.Equal(t, 2, total)
		assert.Len(t, reservations, 2)
	})

	t.Run("InvalidLimitParameter", func(t *testing.T) {
		c.QueryParams().Set("s", "search")
		c.QueryParams().Set("limit", "invalid")
		c.QueryParams().Set("offset", "0")

		_, _, err := reservationService.GetAllReservation(c)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Params limit not valid")
	})

	t.Run("InternalServerError", func(t *testing.T) {
		reservationRepository.GetAllReservationFunc = func(offset, limit int, search string) ([]dto.ReservationResponse, int, error) {
			return nil, 0, fmt.Errorf("Internal Server Error")
		}

		c.QueryParams().Set("s", "search")
		c.QueryParams().Set("limit", "10")
		c.QueryParams().Set("offset", "0")

		_, _, err := reservationService.GetAllReservation(c)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Internal Server Error")
	})

	t.Run("NoReservationsFound", func(t *testing.T) {
		reservationRepository.GetAllReservationFunc = func(offset, limit int, search string) ([]dto.ReservationResponse, int, error) {
			return []dto.ReservationResponse{}, 0, nil
		}

		c.QueryParams().Set("s", "s")
		c.QueryParams().Set("limit", "10")
		c.QueryParams().Set("offset", "0")

		_, _, err := reservationService.GetAllReservation(c)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Reservation not found")
	})
}
