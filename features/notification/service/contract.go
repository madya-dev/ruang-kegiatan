package service

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/notification/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type NotificationService interface {
	// GetAllNotifications(ctx echo.Context) ([]dto.NotificationResponse, int, error)
	// GetNotifications(ctx echo.Context, username string) ([]dto.NotificationResponse, int, error)
	CreateNotification(ctx echo.Context, r dto.NotificationRequest) error
	// ReadNotification(ctx echo.Context, id int) error
}

type NotificationServiceImpl struct {
	NotificationRepository repository.NotificationRepository
	Validate               *validator.Validate
}

func NewNotificationService(rr repository.NotificationRepository, validate *validator.Validate) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		NotificationRepository: rr,
		Validate:               validate,
	}
}
