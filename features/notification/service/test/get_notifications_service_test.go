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

func TestNotificationServiceImpl_GetNotification(t *testing.T) {
	t.Run("NotificationGetSuccessFully", func(t *testing.T) {
		notificationService := &service.NotificationServiceImpl{
			NotificationRepository: &MockNotificationRepository{
				GetNotificationsFunc: func(offset, limit int, username string) ([]model.Notification, error) {
					return []model.Notification{}, nil
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
		c.SetParamNames("username")
		c.SetParamValues("usertest")

		res, err := notificationService.GetNotifications(c)

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}
