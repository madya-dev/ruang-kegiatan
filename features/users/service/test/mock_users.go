package test

import (
	"madyasantosa/ruangkegiatan/model"
)

type MockUsersRepository struct {
	GetAllUsersFunc       func(offset int, limit int, search string) ([]model.User, int, error)
	GetUserByUsernameFunc func(username string) (*model.User, error)
	CreateUserFunc        func(user *model.User) (*model.User, error)
	UpdateUserFunc        func(user *model.User, username string) (*model.User, error)
	ChangePasswordFunc    func(user *model.User, username string) error
	UpdateRoleUserFunc    func(user *model.User, username string) (*model.User, error)
	DeleteUserFunc        func(username string) error
}

func (m *MockUsersRepository) GetAllUsers(offset int, limit int, search string) ([]model.User, int, error) {
	return m.GetAllUsersFunc(offset, limit, search)
}

func (m *MockUsersRepository) GetUserByUsername(username string) (*model.User, error) {
	return m.GetUserByUsernameFunc(username)
}

func (m *MockUsersRepository) CreateUser(user *model.User) (*model.User, error) {
	return m.CreateUserFunc(user)
}

func (m *MockUsersRepository) UpdateUser(user *model.User, username string) (*model.User, error) {
	return m.UpdateUserFunc(user, username)
}

func (m *MockUsersRepository) ChangePassword(user *model.User, username string) error {
	return m.ChangePasswordFunc(user, username)
}

func (m *MockUsersRepository) UpdateRoleUser(user *model.User, username string) (*model.User, error) {
	return m.UpdateRoleUserFunc(user, username)
}

func (m *MockUsersRepository) DeleteUser(username string) error {
	return m.DeleteUserFunc(username)
}
