package repository

import (
	"errors"
	"fmt"
	"madyasantosa/ruangkegiatan/model"

	"gorm.io/gorm"
)

func (r *UserRepositoryImpl) DeleteUser(username string) error {
	user := model.User{}
	result := r.DB.Where("username = ?", username).First(&user)
	fmt.Println(result)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("User not found")
		}
		return fmt.Errorf("Internal Server Error")
	}

	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
