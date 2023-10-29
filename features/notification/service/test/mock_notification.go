package test

import "madyasantosa/ruangkegiatan/model"

type MockNotificationRepository struct {
	GetNotificationsFunc    func(offset int, limit int, username string) ([]model.Notification, error)
	GetAllNotificationsFunc func(offset int, limit int) ([]model.Notification, int, error)
	CreateNotificationFunc  func(notif *model.Notification) error
	ReadNotificationFunc    func(id int) error
}

func (m *MockNotificationRepository) GetNotifications(offset int, limit int, username string) ([]model.Notification, error) {
	return m.GetNotificationsFunc(offset, limit, username)
}

func (m *MockNotificationRepository) GetAllNotifications(offset int, limit int) ([]model.Notification, int, error) {
	return m.GetAllNotificationsFunc(offset, limit)
}

func (m *MockNotificationRepository) CreateNotification(notif *model.Notification) error {
	return m.CreateNotificationFunc(notif)
}

func (m *MockNotificationRepository) ReadNotification(id int) error {
	return m.ReadNotificationFunc(id)

}
