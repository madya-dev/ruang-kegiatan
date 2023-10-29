package test

import (
	"madyasantosa/ruangkegiatan/features/notification/service"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNotificationServiceImpl_ReadNotification(t *testing.T) {
	notificationService := &service.NotificationServiceImpl{
		NotificationRepository: &MockNotificationRepository{
			ReadNotificationFunc: func(id int) error {
				return nil
			},
		},
		Validate: validator.New(),
	}

	e := echo.New()

	req := httptest.NewRequest(echo.PUT, "/your-api-endpoint/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := notificationService.ReadNotification(c)

	assert.NoError(t, err)
}
