package model

import "time"

type Room struct {
	ID        int64  `gorm:"primaryKey;autoIncrement:true;unique"`
	RoomName  string `gorm:"type:varchar(50);not null"`
	Capacity  int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
