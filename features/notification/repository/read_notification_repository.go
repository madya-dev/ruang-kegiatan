package repository

import (
	"madyasantosa/ruangkegiatan/model"
)

func (r *NotificationRepositoryImpl) ReadNotification(id int) error {
	notif := model.Notification{}
	result := r.DB.Model(&notif).Where("id = ?", id).Updates(model.Notification{IsRead: true})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
