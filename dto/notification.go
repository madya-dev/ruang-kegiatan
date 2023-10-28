package dto

type NotificationResponse struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	IsRead   bool   `json:"is_read"`
	Username string `json:"username"`
}
type NotificationRequest struct {
	Title    string `json:"title" validate:"required"`
	Message  string `json:"message" validate:"required"`
	IsRead   bool   `json:"is_read"`
	Username string `json:"username" validate:"required"`
}
