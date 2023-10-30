package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *UserRepositoryImpl) UpdateRoleUser(user *model.User, username string) error {
	users := model.User{}
	result := r.DB.Model(&users).Where("username = ?", username).Updates(model.User{Role: user.Role})
	fmt.Println(result)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
