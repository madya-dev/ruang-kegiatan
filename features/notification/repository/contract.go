package repository

import (
	"madyasantosa/ruangkegiatan/model"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	GetNotifications(offset int, limit int, username string) ([]model.Notification, error)
	GetAllNotifications(offset int, limit int) ([]model.Notification, int, error)
	CreateNotification(notif *model.Notification) error
	ReadNotification(id int) error
}

type NotificationRepositoryImpl struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &NotificationRepositoryImpl{
		DB: db,
	}
}
