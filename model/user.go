package model

import "time"

type Role string

const (
	UserNormal Role = "user"
	Admin      Role = "admin"
)

type User struct {
	Username     string    `gorm:"index:unique;primaryKey;type:varchar(20);not null"`
	FullName     string    `gorm:"type:varchar(255);not null"`
	StudyProgram string    `gorm:"type:varchar(100);not null"`
	Phone        string    `gorm:"type:varchar(15);not null"`
	Role         Role      `gorm:"type:enum('user','admin');default:'user'"`
	Password     string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
