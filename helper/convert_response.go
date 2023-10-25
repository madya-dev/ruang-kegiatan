package helper

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
)

func ConvertToUserResponse(r *model.User) *dto.UserResponse {
	if r.Role == "" {
		r.Role = "user"
	}
	return &dto.UserResponse{
		Username:     r.Username,
		FullName:     r.FullName,
		StudyProgram: r.StudyProgram,
		Phone:        r.Phone,
		Role:         string(r.Role),
	}
}
