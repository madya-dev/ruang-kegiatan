package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/users/service"
	"madyasantosa/ruangkegiatan/helper"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceImpl_UserLogin(t *testing.T) {
	userService := &service.UserServiceImpl{
		UserRepository: &MockUsersRepository{
			GetUserByUsernameFunc: func(username string) (*model.User, error) {
				if username == "testuser" {
					return &model.User{
						Username: "testuser",
						Password: helper.HashPassword("password123"),
					}, nil
				}
				return nil, fmt.Errorf("User not found")
			},
		},
		Validate: validator.New(),
	}

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("SuccessfulLogin", func(t *testing.T) {
		loginRequest := dto.UserLoginRequest{
			Username: "testuser",
			Password: "password123",
		}

		userResponse, err := userService.UserLogin(c, loginRequest)

		assert.NoError(t, err)
		assert.NotNil(t, userResponse)
		assert.Equal(t, "testuser", userResponse.Username)
	})

	t.Run("InvalidUsername", func(t *testing.T) {
		loginRequest := dto.UserLoginRequest{
			Username: "nonexistentuser",
			Password: "password123",
		}

		userResponse, err := userService.UserLogin(c, loginRequest)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Contains(t, err.Error(), "Invalid username or password")
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		loginRequest := dto.UserLoginRequest{
			Username: "testuser",
			Password: "wrongpassword",
		}

		userResponse, err := userService.UserLogin(c, loginRequest)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Contains(t, err.Error(), "Invalid username or password")
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		loginRequest := dto.UserLoginRequest{}

		userResponse, err := userService.UserLogin(c, loginRequest)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Contains(t, err.Error(), "Invalid username or password")
	})
}
