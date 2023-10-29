package test

import (
	"madyasantosa/ruangkegiatan/features/notification/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNotificationServiceImpl_GetAllNotification(t *testing.T) {
	t.Run("NotificationGetSuccessFully", func(t *testing.T) {
		notificationService := &service.NotificationServiceImpl{
			NotificationRepository: &MockNotificationRepository{
				GetAllNotificationsFunc: func(offset, limit int) ([]model.Notification, int, error) {
					return []model.Notification{}, 0, nil
				},
			},
			Validate: validator.New(),
		}

		e := echo.New()

		req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.QueryParams().Set("limit", "10")
		c.QueryParams().Set("offset", "0")

		res, total, err := notificationService.GetAllNotifications(c)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.NotNil(t, total)
	})
}
