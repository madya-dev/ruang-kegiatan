package repository

import (
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
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
