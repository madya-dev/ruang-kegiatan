package dto

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	Capacity int    `json:"capacity"`
}
type RoomResponse struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	Capacity int    `json:"capacity"`
}
type RoomRequest struct {
	RoomName string `json:"room_name" validate:"required,min=1,max=50"`
	Capacity int    `json:"capacity" validate:"required"`
}
