package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/reservation/service"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteReservation(t *testing.T) {
	reservationService := &service.ReservationServiceImpl{
		ReservationRepository: &MockReservationRepository{},
	}

	t.Run("ReservationNotFound", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}
		reservationRepository.CheckReservationFunc = func(id int) (*dto.ReservationCheck, error) {
			return nil, fmt.Errorf("Reservation not found")
		}

		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/your-api-endpoint/123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("123")

		err := reservationService.DeleteReservation(c, "testuser")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Reservation not found")
	})
	t.Run("ReservationForbidden", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}
		reservationRepository.CheckReservationFunc = func(id int) (*dto.ReservationCheck, error) {
			return &dto.ReservationCheck{
				PIC: "otheruser",
			}, nil
		}

		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/your-api-endpoint/123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("123")

		err := reservationService.DeleteReservation(c, "testuser")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Forbidden!")
	})
	t.Run("ReservationDeleteSuccessfuly", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}
		reservationRepository.CheckReservationFunc = func(id int) (*dto.ReservationCheck, error) {
			return &dto.ReservationCheck{
				PIC:       "testuser",
				ID:        1,
				StartTime: time.Now(),
				Document:  "",
			}, nil
		}
		reservationRepository.DeleteReservationFunc = func(id int) error {
			return nil
		}

		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/your-api-endpoint/123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("123")

		err := reservationService.DeleteReservation(c, "testuser")

		assert.NoError(t, err)
	})
}
