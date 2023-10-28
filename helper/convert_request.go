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

func RoomRequestToRoomModel(r dto.RoomRequest) *model.Room {
	return &model.Room{
		RoomName: r.RoomName,
		Capacity: r.Capacity,
	}
}

func ReservationRequestToReservationModel(r dto.ReservationRequest, document string, username string) *model.Reservation {
	return &model.Reservation{
		RoomID:       r.RoomID,
		Activity:     r.Activity,
		StartTime:    r.StartTime,
		EndTime:      r.EndTime,
		StudyProgram: r.StudyProgram,
		ClassOf:      r.ClassOf,
		Document:     document,
		PIC:          username,
	}
}

func NotificationCreateRequestToNotificationModel(r dto.NotificationRequest) *model.Notification {
	return &model.Notification{
		Title:   r.Title,
		Message: r.Message,
		PIC:     r.Username,
	}
}
