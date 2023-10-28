package dto

type NotificationResponse struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Title    string `json:"title" gorm:"type:varchar(20);not null"`
	Message  string `json:"message" gorm:"type:varchar(255);not null"`
	IsRead   bool   `json:"is_read" gorm:"default:false"`
	Username string `gorm:"not null;index;type:varchar(20)"`
}
