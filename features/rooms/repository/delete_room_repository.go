package repository

import (
	"errors"
	"fmt"
	"madyasantosa/ruangkegiatan/model"

	"gorm.io/gorm"
)

func (r *RoomRespositoryImpl) DeleteRoom(id int) error {
	room := model.Room{}
	result := r.DB.Where("id = ?", id).First(&room)
	fmt.Println(result)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Room not found")
		}
		return fmt.Errorf("Internal Server Error")
	}

	if err := r.DB.Delete(&room).Error; err != nil {
		return err
	}
	return nil
}
