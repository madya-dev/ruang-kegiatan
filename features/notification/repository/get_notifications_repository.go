package repository

import (
	"madyasantosa/ruangkegiatan/model"
)

func (r *NotificationRepositoryImpl) GetNotifications(offset int, limit int, username string) ([]model.Notification, error) {
	notif := []model.Notification{}

	result := r.DB.Where("username = ?", username).Offset(offset).Limit(limit).Find(&notif)

	if result.Error != nil {
		return nil, result.Error
	}
	return notif, nil
}

func (r *NotificationRepositoryImpl) GetAllNotifications(offset int, limit int) ([]model.Notification, int, error) {
	notif := []model.Notification{}

	var total int64

	result := r.DB.Offset(offset).Limit(limit).Find(&notif)

	result.Count(&total)

	if result.Error != nil {
		return nil, int(total), result.Error
	}
	return notif, int(total), nil
}
