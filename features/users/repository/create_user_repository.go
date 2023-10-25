package repository

import "madyasantosa/ruangkegiatan/model"

func (r *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	result := r.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
