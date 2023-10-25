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
type UserCreateRequest struct {
	Username     string `json:"username" validate:"required,min=1,max=20"`
	FullName     string `json:"fullname" validate:"required,min=1,max=255"`
	StudyProgram string `json:"study_program" validate:"required,min=1,max=100"`
	Phone        string `json:"phone" validate:"required,min=1,max=15"`
	Password     string `json:"password" validate:"required,min=8,max=255"`
}

type UserUpdateRequest struct {
	Username     string `json:"username" validate:"required,min=1,max=20"`
	FullName     string `json:"fullname" validate:"required,min=1,max=255"`
	StudyProgram string `json:"study_program" validate:"required,min=1,max=100"`
	Phone        string `json:"phone" validate:"required,min=1,max=15"`
}
type UserRoleUpdateRequest struct {
	Role string `json:"role" validate:"required"`
}
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=255"`
}
type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Iat      time.Time `json:"iat"`
	Exp      time.Time `json:"exp"`
}
