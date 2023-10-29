package test

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/notification/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNotificationServiceImpl_CreateNotification(t *testing.T) {
	t.Run("NotificationCreatedSuccessFully", func(t *testing.T) {
		notificationService := &service.NotificationServiceImpl{
			NotificationRepository: &MockNotificationRepository{
				CreateNotificationFunc: func(notif *model.Notification) error {
					return nil
				},
			},
			Validate: validator.New(),
		}

		e := echo.New()

		req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		CreateNotification := dto.NotificationRequest{
			Title:    "Update",
			Message:  "updet",
			Username: "usertest",
		}

		err := notificationService.CreateNotification(c, CreateNotification)

		assert.NoError(t, err)
	})
}
