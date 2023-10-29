package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/users/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceImpl_UpdateUser(t *testing.T) {
	userService := &service.UserServiceImpl{
		UserRepository: &MockUsersRepository{
			GetUserByUsernameFunc: func(username string) (*model.User, error) {
				if username == "testuser" {
					return &model.User{
						Username:     "testuser",
						FullName:     "John Doe",
						StudyProgram: "Computer Science",
						Phone:        "1234567890",
					}, nil
				}
				return nil, fmt.Errorf("User not found")
			},
			UpdateUserFunc: func(user *model.User, username string) (*model.User, error) {
				return user, nil
			},
		},
		Validate: validator.New(),
	}

	e := echo.New()

	req := httptest.NewRequest(echo.PUT, "/your-api-endpoint", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("username")
	c.SetParamValues("testuser")

	updateUserRequest := dto.UserUpdateRequest{
		Username:     "testuser",
		FullName:     "Updated John Doe",
		StudyProgram: "Updated Computer Science",
		Phone:        "9876543210",
	}

	t.Run("ValidUser", func(t *testing.T) {
		userResponse, err := userService.UpdateUser(c, updateUserRequest)

		assert.NoError(t, err)
		assert.NotNil(t, userResponse)
		assert.Equal(t, "testuser", userResponse.Username)
		assert.Equal(t, "Updated John Doe", userResponse.FullName)
		assert.Equal(t, "Updated Computer Science", userResponse.StudyProgram)
		assert.Equal(t, "9876543210", userResponse.Phone)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		c.SetParamNames("username")
		c.SetParamValues("testuser1")
		updateUserRequest := dto.UserUpdateRequest{
			Username:     "nonexistentuser",
			FullName:     "Updated John Doe",
			StudyProgram: "Updated Computer Science",
			Phone:        "9876543210",
		}

		userResponse, err := userService.UpdateUser(c, updateUserRequest)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Contains(t, err.Error(), "User not found")
	})
}
