package test

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/notification/service/test"
	"madyasantosa/ruangkegiatan/features/reservation/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdateReservation(t *testing.T) {
	reservationService := &service.ReservationServiceImpl{
		ReservationRepository:  &MockReservationRepository{},
		Validate:               validator.New(),
		NotificationRepository: &test.MockNotificationRepository{},
	}

	reservationRequest := dto.ReservationRequest{
		RoomID:    1,
		Activity:  "Meeting",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	t.Run("ValidUpdateReservation", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}
		reservationRepository.CheckReservationFunc = func(id int) (*dto.ReservationCheck, error) {
			return &dto.ReservationCheck{
				PIC:       "testuser",
				ID:        1,
				StartTime: time.Now(),
				Document:  "",
			}, nil
		}
		reservationRepository.IsAvaibleFunc = func(roomID int64, startTime, endTime time.Time) (*model.Reservation, string, error) {
			return &model.Reservation{
				Activity:  "Meeting",
				PIC:       "testuser",
				ID:        1,
				StartTime: time.Now(),
				Document:  "",
			}, "", nil
		}
		reservationRepository.UpdateReservationFunc = func(reservation *model.Reservation, id int) error {
			return nil
		}
		notificationRepository := &test.MockNotificationRepository{}
		notificationRepository.CreateNotificationFunc = func(notif *model.Notification) error {
			return nil
		}

		reservationService.NotificationRepository = notificationRepository
		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/your-api-endpoint/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := reservationService.UpdateReservation(c, reservationRequest, "testuser")

		assert.NoError(t, err)
	})

	t.Run("InvalidReservationRequest", func(t *testing.T) {
		reservationRepository := &MockReservationRepository{}
		reservationRepository.CheckReservationFunc = func(id int) (*dto.ReservationCheck, error) {
			return &dto.ReservationCheck{
				PIC:       "testuser",
				ID:        1,
				StartTime: time.Now(),
				Document:  "",
			}, nil
		}
		reservationRepository.IsAvaibleFunc = func(roomID int64, startTime, endTime time.Time) (*model.Reservation, string, error) {
			return &model.Reservation{
				Activity:  "Meeting",
				PIC:       "testuser",
				ID:        1,
				StartTime: time.Now(),
				Document:  "",
			}, "", nil
		}
		reservationRepository.UpdateReservationFunc = func(reservation *model.Reservation, id int) error {
			return nil
		}
		notificationRepository := &test.MockNotificationRepository{}
		notificationRepository.CreateNotificationFunc = func(notif *model.Notification) error {
			return nil
		}

		reservationService.NotificationRepository = notificationRepository
		reservationService.ReservationRepository = reservationRepository

		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/your-api-endpoint/2", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		err := reservationService.UpdateReservation(c, reservationRequest, "testuser")

		assert.Error(t, err)
	})
}
