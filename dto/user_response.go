package dto

import "time"

type UserResponse struct {
	Username     string `json:"username"`
	FullName     string `json:"fullname"`
	StudyProgram string `json:"study_program"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
}
type UserLoginResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type UserDataForNotification struct {
	Username          string `json:"username"`
	RegistrationToken string `json:"registration_token"`
}

type Token struct {
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Iat      time.Time `json:"iat"`
	Exp      time.Time `json:"exp"`
}
