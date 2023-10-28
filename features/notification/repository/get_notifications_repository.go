package repository

import (
	"madyasantosa/ruangkegiatan/model"
)

func (r *NotificationRepositoryImpl) GetNotifications(offset int, limit int, username string) ([]model.Notification, int, error) {
	notif := []model.Notification{}

	var unReadTotal int64

	r.DB.Where("username = ? AND is_read = false", username).Find(&notif).Count(&unReadTotal)

	result := r.DB.Where("username = ? AND is_read = false", username).Offset(offset).Limit(limit).Find(&notif)

	if result.Error != nil {
		return nil, int(unReadTotal), result.Error
	}
	return notif, int(unReadTotal), nil
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
