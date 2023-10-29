package test

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/reservation/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	reservationService := &service.ReservationServiceImpl{
		ReservationRepository: &MockReservationRepository{},
		Validate:              validator.New(),
	}
	reservationRequest := dto.ReservationRequest{
		RoomID:       1,
		Activity:     "Meeting",
		StudyProgram: "Informatics",
		ClassOf:      "2021",
		StartTime:    time.Now(),
		EndTime:      time.Now().Add(1 * time.Hour),
	}

	t.Run("ValidCreateReservation", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}

		reservationRepository.IsAvaibleFunc = func(roomID int64, startTime, endTime time.Time) (*model.Reservation, string, error) {
			return nil, "", nil
		}

		reservationRepository.CreateReservationFunc = func(reservation *model.Reservation) error {
			return nil
		}

		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := reservationService.CreateReservation(c, reservationRequest, "testuser")

		assert.NoError(t, err)
	})

	t.Run("InvalidReservationRequest", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}
		reservationRepository.IsAvaibleFunc = func(roomID int64, startTime, endTime time.Time) (*model.Reservation, string, error) {
			return nil, "", nil
		}

		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("username", "testuser")

		reservationRequest := dto.ReservationRequest{}

		err := reservationService.CreateReservation(c, reservationRequest, "testuser")

		assert.Error(t, err)
	})

	t.Run("RoomAlreadyBooked", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}
		reservationRepository.IsAvaibleFunc = func(roomID int64, startTime, endTime time.Time) (*model.Reservation, string, error) {
			return &model.Reservation{
				Activity: "Meeting",
				PIC:      "otheruser",
			}, "", nil
		}

		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("username", "testuser")

		err := reservationService.CreateReservation(c, reservationRequest, "testuser")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Reservation failed, room already booked")
	})
}
