package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *NotificationRepositoryImpl) CreateNotification(notif *model.Notification) error {
	fmt.Println("Teststststststst")
	result := r.DB.Create(&notif)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
