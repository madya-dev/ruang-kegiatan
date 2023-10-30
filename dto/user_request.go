package dto

type UserCreateRequest struct {
	Username          string `json:"username" validate:"required,min=1,max=20"`
	FullName          string `json:"fullname" validate:"required,min=1,max=255"`
	StudyProgram      string `json:"study_program" validate:"required,min=1,max=100"`
	Phone             string `json:"phone" validate:"required,min=1,max=15"`
	Password          string `json:"password" validate:"required,min=8,max=255"`
	RegistrationToken string `json:"registration_token"`
}

type UserUpdateRequest struct {
	Username          string `json:"username" validate:"required,min=1,max=20"`
	FullName          string `json:"fullname" validate:"required,min=1,max=255"`
	StudyProgram      string `json:"study_program" validate:"required,min=1,max=100"`
	Phone             string `json:"phone" validate:"required,min=1,max=15"`
	RegistrationToken string `json:"registration_token"`
}

type UserRoleUpdateRequest struct {
	Role string `json:"role" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=255"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
