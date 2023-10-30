package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *UserRepositoryImpl) UpdateUser(user *model.User, username string) (*model.User, error) {
	users := model.User{}
	result := r.DB.Model(&users).Where("username = ?", username).Updates(model.User{FullName: user.FullName, Username: user.Username, StudyProgram: user.StudyProgram, Phone: user.Phone, RegistrationToken: user.RegistrationToken})
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
