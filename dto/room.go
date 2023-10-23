package dto

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	Capacity int    `json:"capacity"`
}
