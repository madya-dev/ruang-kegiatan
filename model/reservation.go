package model

import "time"

type Reservation struct {
	ID           int64 `gorm:"primaryKey;autoIncrement:true;unique"`
	RoomID       int64
	Activity     string    `gorm:"type:varchar(255);not null"`
	StartTime    time.Time `gorm:"not null"`
	EndTime      time.Time `gorm:"not null"`
	StudyProgram string    `gorm:"type:varchar(100);not null"`
	ClassOf      string    `gorm:"type:varchar(100);not null"`
	Document     string    `gorm:"type:varchar(255)"`
	PIC          string    `gorm:"not null;index;type:varchar(20)"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	User         User      `gorm:"foreignKey:PIC;references:Username;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Room         Room      `gorm:"foreignKey:RoomID;references:ID"`
}
