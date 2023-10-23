package dto

import "time"

type User struct {
	Username     string `json:"username"`
	FullName     string `json:"fullname"`
	StudyProgram string `json:"study_program"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
}

type Token struct {
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Iat      time.Time `json:"iat"`
	Exp      time.Time `json:"exp"`
}
