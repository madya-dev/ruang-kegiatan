package dto

import "time"

type Reservation struct {
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
}
type ReservationCheck struct {
	ID  int    `json:"id"`
	PIC string `json:"pic"`
}
