package repository

import "madyasantosa/ruangkegiatan/model"

func (r *NotificationRepositoryImpl) CreateNotification(notif *model.Notification) error {
	result := r.DB.Create(&notif)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
