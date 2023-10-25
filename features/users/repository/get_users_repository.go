package repository

import (
	"fmt"
	"madyasantosa/ruangkegiatan/model"
)

func (r *UserRepositoryImpl) GetAllUsers(offset int, limit int, search string) ([]model.User, int, error) {
	users := []model.User{}
	var total int64
	r.DB.Where("full_name LIKE ?", "%"+search+"%").Or("username LIKE ?", "%"+search+"%").Find(&users).Count(&total)

	result := r.DB.Where("full_name LIKE ?", "%"+search+"%").Or("username LIKE ?", "%"+search+"%").Offset(offset).Limit(limit).Find(&users)
	fmt.Println(result)
	if result.Error != nil {
		return nil, int(total), result.Error
	}
	return users, int(total), nil
}
