package repository

import (
	"errors"
	"fmt"
	"madyasantosa/ruangkegiatan/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(offset int, limit int, search string) ([]model.User, int, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User, username string) (*model.User, error)
	ChangePassword(user *model.User, username string) error
	UpdateRoleUser(user *model.User, username string) (*model.User, error)
	DeleteUser(username string) error

	// ChangePassword(username string, oldPassword string, newPassword string) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

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

func (r *UserRepositoryImpl) GetUserByUsername(username string) (*model.User, error) {
	user := model.User{}
	result := r.DB.Where("username = ?", username).First(&user)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
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

func (r *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	result := r.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *model.User, username string) (*model.User, error) {
	users := model.User{}
	result := r.DB.Model(&users).Where("username = ?", username).Updates(model.User{FullName: user.FullName, Username: user.Username, StudyProgram: user.StudyProgram, Phone: user.Phone})
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
func (r *UserRepositoryImpl) UpdateRoleUser(user *model.User, username string) (*model.User, error) {
	users := model.User{}
	result := r.DB.Model(&users).Where("username = ?", username).Updates(model.User{Role: user.Role})
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
func (r *UserRepositoryImpl) ChangePassword(user *model.User, username string) error {
	users := model.User{}
	result := r.DB.Model(&users).Where("username = ?", username).Updates(model.User{Password: user.Password})
	fmt.Println(result)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
