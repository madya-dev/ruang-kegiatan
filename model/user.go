package model

import "time"

type Role string

const (
	UserNormal Role = "user"
	Admin      Role = "admin"
)

type User struct {
	Username          string    `json:"username" gorm:"index:unique;primaryKey;type:varchar(20);not null"`
	FullName          string    `json:"fullname" gorm:"type:varchar(255);not null"`
	StudyProgram      string    `json:"study_program" gorm:"type:varchar(100);not null"`
	Phone             string    `json:"phone" gorm:"type:varchar(15);not null"`
	Role              Role      `json:"role" gorm:"type:enum('user','admin');default:'user'"`
	RegistrationToken string    `json:"registration_token" gorm:"type:varchar(255)"`
	Password          string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt         time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"-" gorm:"autoUpdateTime"`
}
type UserWithPass struct {
	Username     string    `json:"username" gorm:"index:unique;primaryKey;type:varchar(20);not null"`
	FullName     string    `json:"fullname" gorm:"type:varchar(255);not null"`
	StudyProgram string    `json:"study_program" gorm:"type:varchar(100);not null"`
	Phone        string    `json:"phone" gorm:"type:varchar(15);not null"`
	Role         Role      `json:"role" gorm:"type:enum('user','admin');default:'user'"`
	Password     string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"-" gorm:"autoUpdateTime"`
}
