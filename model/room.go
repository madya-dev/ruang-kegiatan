package model

import "time"

type Room struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	RoomName  string    `json:"room_name" gorm:"type:varchar(50);not null"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
