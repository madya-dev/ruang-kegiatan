package database

import (
	"madyasantosa/ruangkegiatan/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Room{}, &model.Reservation{}, &model.Notification{})
}
