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

func TestUserServiceImpl_ChangePassword(t *testing.T) {
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
			ChangePasswordFunc: func(user *model.User, username string) error {
				if username == "testuser" {
					return nil
				}
				return fmt.Errorf("Failed to change password")
			},
		},
		Validate: validator.New(),
	}

	t.Run("PasswordChangedSuccessfully", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues("testuser")

		changePasswordRequest := dto.ChangePasswordRequest{
			OldPassword: "password123",
			NewPassword: "newpassword123",
		}

		err := userService.ChangePassword(c, changePasswordRequest, "testuser")

		assert.NoError(t, err)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues("nonexistentuser")

		changePasswordRequest := dto.ChangePasswordRequest{
			OldPassword: "password123",
			NewPassword: "newpassword123",
		}

		err := userService.ChangePassword(c, changePasswordRequest, "nonexistentuser")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "User not found")
	})

	t.Run("InvalidOldPassword", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues("testuser")

		changePasswordRequest := dto.ChangePasswordRequest{
			OldPassword: "incorrectpassword",
			NewPassword: "newpassword123",
		}

		err := userService.ChangePassword(c, changePasswordRequest, "testuser")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Invalid old password")
	})

	t.Run("ChangePasswordError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues("testuser")

		changePasswordRequest := dto.ChangePasswordRequest{
			OldPassword: "password12333",
			NewPassword: "newpassword123",
		}

		err := userService.ChangePassword(c, changePasswordRequest, "testuser")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Invalid old password")
	})
}
