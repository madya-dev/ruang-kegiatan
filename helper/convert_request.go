package helper

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
)

func UserCreateRequestToUserModel(r dto.UserCreateRequest) *model.User {
	return &model.User{
		FullName:     r.FullName,
		Username:     r.Username,
		Password:     r.Password,
		StudyProgram: r.StudyProgram,
		Phone:        r.Phone,
	}
}

func UserLoginRequestToUserModel(r dto.UserLoginRequest) *model.User {
	return &model.User{
		Username: r.Username,
		Password: r.Password,
	}
}

func UserUpdateRequestToUserModel(r dto.UserUpdateRequest) *model.User {
	return &model.User{
		FullName:     r.FullName,
		Username:     r.Username,
		StudyProgram: r.StudyProgram,
		Phone:        r.Phone,
	}
}
func UserRoleUpdateRequestToUserModel(r dto.UserRoleUpdateRequest) *model.User {
	return &model.User{
		Role: model.Role(r.Role),
	}
}
