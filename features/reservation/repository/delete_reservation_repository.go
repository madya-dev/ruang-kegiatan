package repository

import (
	"errors"
	"fmt"
	"madyasantosa/ruangkegiatan/model"

	"gorm.io/gorm"
)

func (r *ReservationRepositoryImpl) DeleteReservation(id int) error {
	reservation := model.Reservation{}

	result := r.DB.Where("id = ?", id).First(&reservation)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Reservation not found")
		}
		return fmt.Errorf("Internal Server Error")
	}

	if err := r.DB.Delete(&reservation).Error; err != nil {
		return err
	}
	return nil
}
