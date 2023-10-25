package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *UserRepositoryImpl) GetUserByUsername(username string) (*model.User, error) {
	user := model.User{}
	result := r.DB.Where("username = ?", username).First(&user)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
