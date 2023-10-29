package dto

import "time"

type ReservationResponse struct {
	ID           int       `json:"id"`
	RoomName     string    `json:"room_name"`
	Activity     string    `json:"activity"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	StudyProgram string    `json:"study_program"`
	ClassOf      string    `json:"class_of"`
	Document     string    `json:"document"`
	PIC          string    `json:"pic"`
	Phone        string    `json:"phone"`
	UpdatedAt    time.Time `json:"timestamp"`
}
type ReservationCheck struct {
	ID        int       `json:"id"`
	PIC       string    `json:"pic"`
	StartTime time.Time `json:"start_time"`
	Document  string    `json:"document"`
}

type ReservationRequest struct {
	RoomID       int64     `form:"room_id" validate:"required"`
	Activity     string    `form:"activity" `
	StartTime    time.Time `form:"start_time" validate:"required"`
	EndTime      time.Time `form:"end_time" validate:"required"`
	StudyProgram string    `form:"study_program" validate:"required"`
	ClassOf      string    `form:"class_of" validate:"required"`
	Document     []byte    `form:"document"`
}
